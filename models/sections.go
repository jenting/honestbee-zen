package models

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"

	"github.com/honestbee/Zen/internal/db"
)

type sectionsService interface {
	SyncWithSections(ctx context.Context, zendeskSections []*Section, countryCode, locale string) error
	GetSections(ctx context.Context, params *GetSectionsParams) ([]*Section, int, error)
	GetSectionsByCategoryID(ctx context.Context, params *GetSectionsParams) ([]*Section, int, error)
	GetSectionBySectionID(ctx context.Context, sectionID int, locale, countryCode string) (*Section, error)
	GetSectionByArticleID(ctx context.Context, articleID int, locale, countryCode string) (*Section, error)
}

// Section is the section model.
type Section struct {
	CategoryID   int       `json:"category_id"`
	ID           int       `json:"id"`
	Position     int       `json:"position"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	SourceLocale string    `json:"source_locale"`
	Outdated     bool      `json:"outdated"`
	CountryCode  string    `json:"country_code"`
	URL          string    `json:"url"`
	HTMLURL      string    `json:"html_url"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Locale       string    `json:"locale"`
}

type sectionsOps struct {
	db db.Database
}

// GetSectionsParams is the params structure of requesting GetSections method.
type GetSectionsParams struct {
	CategoryID  int
	Locale      string
	CountryCode string
	PerPage     int
	Page        int
	SortBy      string
	SortOrder   string
}

const (
	updateSectionsQuery = `
	UPDATE sections SET 
		category_id = :category_id,
		position = :position,
		created_at = :created_at,
		updated_at = :updated_at,
		source_locale = :source_locale,
		outdated = :outdated
	WHERE id = :id AND country_code = :country_code`

	updateSectionTranslates = `
	UPDATE section_translates SET 
		url = :url,
		html_url = :html_url,
		name = :name,
		description = :description
	WHERE section_id = :section_id AND locale = :locale`

	insertSectionsQuery = `
	INSERT INTO sections (category_id, id, position, created_at, updated_at, source_locale, outdated, country_code) 
	VALUES (:category_id, :id, :position, :created_at, :updated_at, :source_locale, :outdated, :country_code)`

	insertSectionTranslatesQuery = `
	INSERT INTO section_translates (section_id, url, html_url, name, description, locale)
	VALUES (:section_id, :url, :html_url, :name, :description, :locale)`

	deleteSectionsQuery          = "DELETE FROM sections WHERE id = :id AND country_code = :country_code"
	deleteSectionTranslatesQuery = "DELETE FROM section_translates WHERE section_id = :section_id AND locale = :locale"
)

// SyncWithSections ensures the database data will be same as the input data.
func (s *sectionsOps) SyncWithSections(ctx context.Context, zendeskSections []*Section, countryCode, locale string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return errors.Wrapf(err, "models: [SyncWithSections] db.Begin failed")
	}

	ids := make([]int, 0)
	query := fmt.Sprintf(`SELECT id FROM sections WHERE country_code = '%s'`, countryCode)
	tx.Select(&ids, query)

	dbIDs := make(map[int]struct{})
	for _, id := range ids {
		dbIDs[id] = struct{}{}
	}

	for _, zendeskSection := range zendeskSections {
		dbSection := &db.Sections{
			CategoryID:   zendeskSection.CategoryID,
			ID:           zendeskSection.ID,
			CountryCode:  zendeskSection.CountryCode,
			CreatedAt:    zendeskSection.CreatedAt,
			Outdated:     zendeskSection.Outdated,
			Position:     zendeskSection.Position,
			SourceLocale: zendeskSection.SourceLocale,
			UpdatedAt:    zendeskSection.UpdatedAt,
		}
		translates := &db.SectionTranslates{
			SectionID:   zendeskSection.ID,
			Description: zendeskSection.Description,
			HTMLURL:     zendeskSection.HTMLURL,
			Locale:      zendeskSection.Locale,
			Name:        zendeskSection.Name,
			URL:         zendeskSection.URL,
		}

		if _, exist := dbIDs[zendeskSection.ID]; exist {
			tx.NamedExec(updateSectionsQuery, dbSection)

			tranIDs := make([]int, 0)
			query = fmt.Sprintf(
				`SELECT section_id FROM section_translates WHERE locale = '%s' AND section_id = '%d'`,
				locale,
				zendeskSection.ID,
			)
			tx.Select(&tranIDs, query)

			if len(tranIDs) > 0 {
				tx.NamedExec(updateSectionTranslates, translates)
			} else {
				tx.NamedExec(insertSectionTranslatesQuery, translates)
			}

			delete(dbIDs, zendeskSection.ID)
		} else {
			tx.NamedExec(insertSectionsQuery, dbSection)
			tx.NamedExec(insertSectionTranslatesQuery, translates)
		}
	}

	for id := range dbIDs {
		tx.NamedExec(deleteSectionTranslatesQuery, map[string]interface{}{"section_id": id, "locale": locale})

		// Check if there is any section_translates reference to sections.
		total := 0
		query = fmt.Sprintf(`SELECT COUNT(*) FROM section_translates WHERE section_id = '%d'`, id)
		tx.Get(&total, query)
		if total == 0 {
			tx.NamedExec(deleteSectionsQuery, map[string]interface{}{"id": id, "country_code": countryCode})
		}
	}
	tx.Commit()

	return errors.Wrapf(tx.Err(), "models: [SyncWithSections] db transaction failed")
}

func (c *categoriesOps) GetSections(ctx context.Context, params *GetSectionsParams) ([]*Section, int, error) {
	sections := make([]*db.Sections, 0)
	query := fmt.Sprintf(
		`SELECT category_id,id,position,created_at,updated_at,source_locale,outdated,country_code 
		FROM sections WHERE country_code = '%s' ORDER BY %s %s, created_at DESC LIMIT %d OFFSET %d`,
		params.CountryCode,
		params.SortBy,
		params.SortOrder,
		params.PerPage,
		params.Page,
	)
	if err := c.db.Select(ctx, &sections, query); err != nil {
		return nil, 0, errors.Wrapf(err, "models: [GetSections] db select secitons failed")
	}

	ret := make([]*Section, 0)
	for _, section := range sections {
		translates := new(db.SectionTranslates)
		query = fmt.Sprintf(
			`SELECT url,html_url,name,description,locale 
			FROM section_translates WHERE section_id = '%d' AND locale = '%s'`,
			section.ID,
			params.Locale,
		)
		if err := c.db.Get(ctx, translates, query); err != nil {
			if err == db.ErrNoRows {
				continue
			}
			return nil, 0, errors.Wrapf(err, "models: [GetSections] db get translates failed")
		}

		ret = append(ret, &Section{
			CategoryID:   section.CategoryID,
			CountryCode:  section.CountryCode,
			CreatedAt:    section.CreatedAt,
			Description:  translates.Description,
			ID:           section.ID,
			HTMLURL:      translates.HTMLURL,
			Locale:       translates.Locale,
			Name:         translates.Name,
			Outdated:     section.Outdated,
			Position:     section.Position,
			SourceLocale: section.SourceLocale,
			UpdatedAt:    section.UpdatedAt,
			URL:          translates.URL,
		})
	}

	total := 0
	query = fmt.Sprintf(
		`SELECT COUNT(*) FROM sections INNER JOIN section_translates 
		ON sections.id=section_translates.section_id 
		WHERE country_code = '%s' AND locale = '%s'`,
		params.CountryCode, params.Locale,
	)
	if err := c.db.Get(ctx, &total, query); err != nil {
		return nil, 0, errors.Wrapf(err, "models: [GetSections] db get total failed")
	}

	return ret, total, nil
}

func (c *categoriesOps) GetSectionsByCategoryID(ctx context.Context, params *GetSectionsParams) ([]*Section, int, error) {
	sections := make([]*db.Sections, 0)
	query := fmt.Sprintf(
		`SELECT category_id,id,position,created_at,updated_at,source_locale,outdated,country_code 
		FROM sections WHERE country_code = '%s' AND category_id = '%d' 
		ORDER BY %s %s, created_at DESC LIMIT %d OFFSET %d`,
		params.CountryCode,
		params.CategoryID,
		params.SortBy,
		params.SortOrder,
		params.PerPage,
		params.Page,
	)
	if err := c.db.Select(ctx, &sections, query); err != nil {
		return nil, 0, errors.Wrapf(err, "models: [GetSectionsByCategoryID] db select sections failed")
	}

	ret := make([]*Section, 0)
	for _, section := range sections {
		translates := new(db.SectionTranslates)
		query = fmt.Sprintf(
			`SELECT url,html_url,name,description,locale 
			FROM section_translates WHERE section_id = '%d' AND locale = '%s'`,
			section.ID,
			params.Locale,
		)
		if err := c.db.Get(ctx, translates, query); err != nil {
			if err == db.ErrNoRows {
				continue
			}
			return nil, 0, errors.Wrapf(err, "models: [GetSectionsByCategoryID] db get translates failed")
		}

		ret = append(ret, &Section{
			CategoryID:   section.CategoryID,
			CountryCode:  section.CountryCode,
			CreatedAt:    section.CreatedAt,
			Description:  translates.Description,
			ID:           section.ID,
			HTMLURL:      translates.HTMLURL,
			Locale:       translates.Locale,
			Name:         translates.Name,
			Outdated:     section.Outdated,
			Position:     section.Position,
			SourceLocale: section.SourceLocale,
			UpdatedAt:    section.UpdatedAt,
			URL:          translates.URL,
		})
	}

	total := 0
	query = fmt.Sprintf(
		`SELECT COUNT(*) FROM sections INNER JOIN section_translates 
		ON sections.id=section_translates.section_id 
		WHERE country_code = '%s' AND locale = '%s' AND category_id = '%d'`,
		params.CountryCode, params.Locale, params.CategoryID,
	)
	if err := c.db.Get(ctx, &total, query); err != nil {
		return nil, 0, errors.Wrapf(err, "models: [GetSectionsByCategoryID] db get total failed")
	}

	return ret, total, nil
}

func (s *sectionsOps) GetSectionBySectionID(ctx context.Context, sectionID int, locale, countryCode string) (*Section, error) {
	section := new(db.Sections)
	query := fmt.Sprintf(
		`SELECT category_id,id,position,created_at,updated_at,source_locale,outdated,country_code 
		FROM sections WHERE country_code = '%s' AND id = '%d'`,
		countryCode,
		sectionID,
	)
	if err := s.db.Get(ctx, section, query); err != nil {
		switch err {
		case db.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, errors.Wrapf(err, "models: [GetSection] db get sections failed")
		}
	}

	translates := new(db.SectionTranslates)
	query = fmt.Sprintf(
		`SELECT url,html_url,name,description,locale 
		FROM section_translates WHERE section_id = '%d' AND locale = '%s'`,
		section.ID,
		locale,
	)
	if err := s.db.Get(ctx, translates, query); err != nil {
		if err != db.ErrNoRows {
			return nil, errors.Wrapf(err, "models: [GetSection] db get translates failed")
		}
	}

	ret := &Section{
		ID:           section.ID,
		CategoryID:   section.CategoryID,
		CountryCode:  section.CountryCode,
		CreatedAt:    section.CreatedAt,
		Description:  translates.Description,
		Outdated:     section.Outdated,
		Position:     section.Position,
		SourceLocale: section.SourceLocale,
		UpdatedAt:    section.UpdatedAt,
		HTMLURL:      translates.HTMLURL,
		Locale:       translates.Locale,
		Name:         translates.Name,
		URL:          translates.URL,
	}

	return ret, nil
}

func (s *sectionsOps) GetSectionByArticleID(ctx context.Context, articleID int, locale, countryCode string) (*Section, error) {
	article := new(db.Articles)
	query := fmt.Sprintf(`SELECT section_id FROM articles WHERE id = %d`, articleID)
	if err := s.db.Get(ctx, article, query); err != nil {
		return nil, errors.Wrapf(err, "models: [GetSectionByArticleID] db get articles failed")
	}

	section := new(db.Sections)
	query = fmt.Sprintf(
		`SELECT category_id,id,position,created_at,updated_at,source_locale,outdated,country_code 
		FROM sections WHERE country_code = '%s' AND id = '%d'`,
		countryCode,
		article.SectionID,
	)
	if err := s.db.Get(ctx, section, query); err != nil {
		switch err {
		case db.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, errors.Wrapf(err, "models: [GetSectionByArticleID] db get sections failed")
		}
	}

	translates := new(db.SectionTranslates)
	query = fmt.Sprintf(
		`SELECT url,html_url,name,description,locale 
		FROM section_translates WHERE section_id = '%d' AND locale = '%s'`,
		section.ID,
		locale,
	)
	if err := s.db.Get(ctx, translates, query); err != nil {
		if err != db.ErrNoRows {
			return nil, errors.Wrapf(err, "models: [GetSectionByArticleID] db get translates failed")
		}
	}

	ret := &Section{
		ID:           section.ID,
		CategoryID:   section.CategoryID,
		CountryCode:  section.CountryCode,
		CreatedAt:    section.CreatedAt,
		Description:  translates.Description,
		Outdated:     section.Outdated,
		Position:     section.Position,
		SourceLocale: section.SourceLocale,
		UpdatedAt:    section.UpdatedAt,
		HTMLURL:      translates.HTMLURL,
		Locale:       translates.Locale,
		Name:         translates.Name,
		URL:          translates.URL,
	}

	return ret, nil
}
