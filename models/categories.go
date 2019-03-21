package models

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/honestbee/Zen/internal/db"
)

type categoriesService interface {
	SyncWithCategories(ctx context.Context, zendeskCategories []*Category, countryCode, locale string) error
	GetCategoriesID(ctx context.Context, countryCode string) ([]int, error)
	GetCategories(ctx context.Context, params *GetCategoriesParams) ([]*Category, int, error)
	GetCategoryKeyNameToID(ctx context.Context, keyName, countryCode string) (int, error)
	GetCategoryByArticleID(ctx context.Context, articleID int, locale string) (*Category, error)
	GetCategoryBySectionID(ctx context.Context, sectionID int, locale string) (*Category, error)
	GetCategoryByCategoryIDOrKeyName(ctx context.Context, idOrKeyName, locale, countryCode string) (*Category, error)
}

// Category is the Category model.
type Category struct {
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
	KeyName      string    `json:"key_name"`
}

type categoriesOps struct {
	db db.Database
}

// GetCategoriesParams is the params structure of requesting GetCategories method.
type GetCategoriesParams struct {
	Locale      string
	CountryCode string
	PerPage     int
	Page        int
	SortBy      string
	SortOrder   string
}

const (
	updateCategoriesQuery = `
	UPDATE categories SET 
		position = :position,
		created_at = :created_at,
		updated_at = :updated_at,
		source_locale = :source_locale,
		outdated = :outdated
	WHERE id = :id AND country_code = :country_code`

	updateCategoryTranslates = `
	UPDATE category_translates SET 
		url = :url,
		html_url = :html_url,
		name = :name,
		description = :description
	WHERE category_id = :category_id AND locale = :locale`

	insertCategoriesQuery = `
	INSERT INTO categories (id, position, created_at, updated_at, source_locale, outdated, country_code) 
	VALUES (:id, :position, :created_at, :updated_at, :source_locale, :outdated, :country_code)`

	insertCategoryTranslatesQuery = `
	INSERT INTO category_translates (category_id, url, html_url, name, description, locale)
	VALUES (:category_id, :url, :html_url, :name, :description, :locale)`

	deleteCategoriesQuery         = "DELETE FROM categories WHERE id = :id AND country_code = :country_code"
	deleteCategoryTranslatesQuery = "DELETE FROM category_translates WHERE category_id = :category_id AND locale = :locale"
)

// SyncWithCategories ensures the database data will be same as the input data.
func (c *categoriesOps) SyncWithCategories(ctx context.Context, zendeskCategories []*Category, countryCode, locale string) error {
	tx, err := c.db.Begin()
	if err != nil {
		return errors.Wrapf(err, "models: [SyncWithCategories] db.Begin failed")
	}

	ids := make([]int, 0)
	query := fmt.Sprintf(`SELECT id FROM categories WHERE country_code = '%s'`, countryCode)
	tx.Select(&ids, query)

	dbIDs := make(map[int]struct{})
	for _, id := range ids {
		dbIDs[id] = struct{}{}
	}

	for _, zendeskCategory := range zendeskCategories {
		dbCategory := &db.Categories{
			ID:           zendeskCategory.ID,
			CountryCode:  zendeskCategory.CountryCode,
			CreatedAt:    zendeskCategory.CreatedAt,
			Outdated:     zendeskCategory.Outdated,
			Position:     zendeskCategory.Position,
			SourceLocale: zendeskCategory.SourceLocale,
			UpdatedAt:    zendeskCategory.UpdatedAt,
		}
		translates := &db.CategoryTranslates{
			CategoryID:  zendeskCategory.ID,
			Description: zendeskCategory.Description,
			HTMLURL:     zendeskCategory.HTMLURL,
			Locale:      zendeskCategory.Locale,
			Name:        zendeskCategory.Name,
			URL:         zendeskCategory.URL,
		}

		if _, exist := dbIDs[zendeskCategory.ID]; exist {
			tx.NamedExec(updateCategoriesQuery, dbCategory)

			tranIDs := make([]int, 0)
			query = fmt.Sprintf(
				`SELECT category_id FROM category_translates WHERE locale = '%s' AND category_id = '%d'`,
				locale,
				zendeskCategory.ID,
			)
			tx.Select(&tranIDs, query)

			if len(tranIDs) > 0 {
				tx.NamedExec(updateCategoryTranslates, translates)
			} else {
				tx.NamedExec(insertCategoryTranslatesQuery, translates)
			}

			delete(dbIDs, zendeskCategory.ID)
		} else {
			tx.NamedExec(insertCategoriesQuery, dbCategory)
			tx.NamedExec(insertCategoryTranslatesQuery, translates)
		}
	}

	for id := range dbIDs {
		tx.NamedExec(deleteCategoryTranslatesQuery, map[string]interface{}{"category_id": id, "locale": locale})

		// Check if there is any category_translates reference to categories.
		total := 0
		query = fmt.Sprintf(`SELECT COUNT(*) FROM category_translates WHERE category_id = '%d'`, id)
		tx.Get(&total, query)
		if total == 0 {
			tx.NamedExec(deleteCategoriesQuery, map[string]interface{}{"id": id, "country_code": countryCode})
		}
	}
	tx.Commit()

	return errors.Wrapf(tx.Err(), "models: [SyncWithCategories] db transaction failed")
}

func (c *categoriesOps) GetCategoriesID(ctx context.Context, countryCode string) ([]int, error) {
	ids := make([]int, 0)
	query := fmt.Sprintf(`SELECT id FROM categories WHERE country_code = '%s'`, countryCode)
	if err := c.db.Select(ctx, &ids, query); err != nil {
		return nil, errors.Wrapf(err, "models: [GetCategoriesID] db select categories failed")
	}
	return ids, nil
}

func (c *categoriesOps) GetCategories(ctx context.Context, params *GetCategoriesParams) ([]*Category, int, error) {
	categories := make([]*db.Categories, 0)
	query := fmt.Sprintf(
		`SELECT id,position,created_at,updated_at,source_locale,outdated,country_code 
		FROM categories WHERE country_code = '%s' 
		ORDER BY %s %s, created_at DESC LIMIT %d OFFSET %d`,
		params.CountryCode,
		params.SortBy,
		params.SortOrder,
		params.PerPage,
		params.Page,
	)
	if err := c.db.Select(ctx, &categories, query); err != nil {
		return nil, 0, errors.Wrapf(err, "models: [GetCategories] db select categories failed")
	}

	ret := make([]*Category, 0)
	for _, category := range categories {
		translates := new(db.CategoryTranslates)
		query = fmt.Sprintf(
			`SELECT url,html_url,name,description,locale 
			FROM category_translates WHERE category_id = '%d' AND locale = '%s'`,
			category.ID,
			params.Locale,
		)
		if err := c.db.Get(ctx, translates, query); err != nil {
			if err == db.ErrNoRows {
				continue
			}
			return nil, 0, errors.Wrapf(err, "models: [GetCategories] db get category translates failed ")
		}

		categoryKey := new(db.CategoryKey)
		query = fmt.Sprintf(
			`SELECT key_name FROM category_key WHERE category_id = '%d'`,
			category.ID,
		)
		if err := c.db.Get(ctx, categoryKey, query); err != nil {
			if err != db.ErrNoRows {
				return nil, 0, errors.Wrapf(err, "models: [GetCategories] db get category key failed ")
			}
		}

		ret = append(ret, &Category{
			ID:           category.ID,
			Position:     category.Position,
			CreatedAt:    category.CreatedAt,
			UpdatedAt:    category.UpdatedAt,
			SourceLocale: category.SourceLocale,
			Outdated:     category.Outdated,
			CountryCode:  category.CountryCode,
			URL:          translates.URL,
			HTMLURL:      translates.HTMLURL,
			Name:         translates.Name,
			Description:  translates.Description,
			Locale:       translates.Locale,
			KeyName:      categoryKey.KeyName,
		})
	}

	total := 0
	query = fmt.Sprintf(
		`SELECT COUNT(*) FROM categories INNER JOIN category_translates 
		ON categories.id=category_translates.category_id 
		WHERE country_code = '%s' AND locale = '%s'`,
		params.CountryCode, params.Locale,
	)
	if err := c.db.Get(ctx, &total, query); err != nil {
		return nil, 0, errors.Wrapf(err, "models: [GetCategories] db get total failed")
	}

	return ret, total, nil
}

func (c *categoriesOps) GetCategoryKeyNameToID(ctx context.Context, keyName, countryCode string) (int, error) {
	category := new(db.CategoryKey)
	// case-insensitive query
	query := fmt.Sprintf(
		`SELECT category_id FROM category_key 
		WHERE LOWER(key_name)=LOWER('%s') and country_code = '%s'`,
		keyName,
		countryCode,
	)
	if err := c.db.Get(ctx, category, query); err != nil {
		switch err {
		case db.ErrNoRows:
			return 0, ErrNotFound
		default:
			return 0, errors.Wrapf(err, "models: [GetCategoryKeyNameToID] db get category_key failed")
		}
	}

	return category.CategoryID, nil
}

func (c *categoriesOps) GetCategoryByArticleID(ctx context.Context, articleID int, locale string) (*Category, error) {
	article := new(db.Articles)
	query := fmt.Sprintf(`SELECT section_id FROM articles WHERE id = %d`, articleID)
	if err := c.db.Get(ctx, article, query); err != nil {
		return nil, errors.Wrapf(err, "models: [GetCategory] db get articles failed")
	}

	section := new(db.Sections)
	query = fmt.Sprintf(`SELECT category_id FROM sections WHERE id = %d`, article.SectionID)
	if err := c.db.Get(ctx, section, query); err != nil {
		return nil, errors.Wrapf(err, "models: [GetCategory] db get sections failed")
	}

	category := new(db.Category)
	query = fmt.Sprintf(
		`SELECT id,position,created_at,updated_at,source_locale,outdated,country_code FROM categories WHERE id = %d`,
		section.CategoryID,
	)
	if err := c.db.Get(ctx, category, query); err != nil {
		return nil, errors.Wrapf(err, "models: [GetCategory] db get categories failed")
	}

	translate := new(db.CategoryTranslates)
	query = fmt.Sprintf(
		`SELECT url,html_url,name,description,locale FROM category_translates WHERE category_id = '%d' AND locale = '%s'`,
		category.ID,
		locale,
	)
	if err := c.db.Get(ctx, translate, query); err != nil {
		return nil, errors.Wrapf(err, "models: [GetCategory] db get category_translates failed")
	}

	return &Category{
		ID:           category.ID,
		Position:     category.Position,
		CreatedAt:    category.CreatedAt,
		UpdatedAt:    category.UpdatedAt,
		SourceLocale: category.SourceLocale,
		Outdated:     category.Outdated,
		CountryCode:  category.CountryCode,
		URL:          translate.URL,
		HTMLURL:      translate.HTMLURL,
		Name:         translate.Name,
		Description:  translate.Description,
		Locale:       translate.Locale,
	}, nil
}

func (c *categoriesOps) GetCategoryBySectionID(ctx context.Context, sectionID int, locale string) (*Category, error) {
	section := new(db.Sections)
	query := fmt.Sprintf(`SELECT category_id FROM sections WHERE id = %d`, sectionID)
	if err := c.db.Get(ctx, section, query); err != nil {
		return nil, errors.Wrapf(err, "models: [GetCategoryBySectionID] db get sections failed")
	}

	category := new(db.Category)
	query = fmt.Sprintf(
		`SELECT id,position,created_at,updated_at,source_locale,outdated,country_code FROM categories WHERE id = %d`,
		section.CategoryID,
	)
	if err := c.db.Get(ctx, category, query); err != nil {
		return nil, errors.Wrapf(err, "models: [GetCategoryBySectionID] db get categories failed")
	}

	translate := new(db.CategoryTranslates)
	query = fmt.Sprintf(
		`SELECT url,html_url,name,description,locale FROM category_translates WHERE category_id = '%d' AND locale = '%s'`,
		category.ID,
		locale,
	)
	if err := c.db.Get(ctx, translate, query); err != nil {
		return nil, errors.Wrapf(err, "models: [GetCategoryBySectionID] db get category_translates failed")
	}

	return &Category{
		ID:           category.ID,
		Position:     category.Position,
		CreatedAt:    category.CreatedAt,
		UpdatedAt:    category.UpdatedAt,
		SourceLocale: category.SourceLocale,
		Outdated:     category.Outdated,
		CountryCode:  category.CountryCode,
		URL:          translate.URL,
		HTMLURL:      translate.HTMLURL,
		Name:         translate.Name,
		Description:  translate.Description,
		Locale:       translate.Locale,
	}, nil
}

func (c *categoriesOps) GetCategoryByCategoryIDOrKeyName(ctx context.Context, idOrKeyName, locale, countryCode string) (*Category, error) {
	var query string

	// Select category_id from table category_key.
	categoryKey := new(db.CategoryKey)
	categoryID64, err := strconv.ParseInt(idOrKeyName, 10, 64)
	if err == nil {
		// it's a number.
		query = fmt.Sprintf(`SELECT category_id,key_name FROM category_key WHERE category_id = %d AND country_code = '%s'`, categoryID64, countryCode)
		if err := c.db.Get(ctx, categoryKey, query); err != nil {
			switch err {
			case db.ErrNoRows:
				return nil, ErrNotFound
			default:
				return nil, errors.Wrapf(err, "models: [GetCategoryByCategoryIDOrKeyName] db get category key failed")
			}
		}
	} else {
		// it's a string.
		query = fmt.Sprintf(`SELECT category_id,key_name FROM category_key WHERE LOWER(key_name)=LOWER('%s') AND country_code = '%s'`, idOrKeyName, countryCode)
		if err := c.db.Get(ctx, categoryKey, query); err != nil {
			switch err {
			case db.ErrNoRows:
				return nil, ErrNotFound
			default:
				return nil, errors.Wrapf(err, "models: [GetCategoryByCategoryIDOrKeyName] db get category key failed")
			}
		}
	}

	// Select table categories.
	category := new(db.Categories)
	query = fmt.Sprintf(`SELECT id,position,created_at,updated_at,source_locale,outdated,country_code FROM categories WHERE id = %d AND country_code = '%s'`, categoryKey.CategoryID, countryCode)
	if err := c.db.Get(ctx, category, query); err != nil {
		switch err {
		case db.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, errors.Wrapf(err, "models: [GetCategoryByCategoryIDOrKeyName] db get categories failed")
		}
	}

	// Select table category_translates.
	categoryTranslate := new(db.CategoryTranslates)
	query = fmt.Sprintf(`SELECT url,html_url,name,description,locale FROM category_translates WHERE category_id = %d AND locale = '%s'`, categoryKey.CategoryID, locale)
	if err := c.db.Get(ctx, categoryTranslate, query); err != nil {
		switch err {
		case db.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, errors.Wrapf(err, "models: [GetCategoryByCategoryIDOrKeyName] db get category_translates failed")
		}
	}

	return &Category{
		ID:           category.ID,
		Position:     category.Position,
		CreatedAt:    category.CreatedAt,
		UpdatedAt:    category.UpdatedAt,
		SourceLocale: category.SourceLocale,
		Outdated:     category.Outdated,
		CountryCode:  category.CountryCode,
		KeyName:      categoryKey.KeyName,
		URL:          categoryTranslate.URL,
		HTMLURL:      categoryTranslate.HTMLURL,
		Name:         categoryTranslate.Name,
		Description:  categoryTranslate.Description,
		Locale:       categoryTranslate.Locale,
	}, nil
}
