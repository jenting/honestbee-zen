package models

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/honestbee/Zen/internal/db"
)

type articlesService interface {
	SyncWithArticles(ctx context.Context, zendeskArticles []*Article, countryCode, locale string) error
	SyncWithArticle(ctx context.Context, articleID int, zendeskArticle *Article, countryCode, locale string) error
	GetArticles(ctx context.Context, params *GetArticlesParams) ([]*Article, int, error)
	GetArticlesByCategoryID(ctx context.Context, params *GetArticlesParams, labels []string) ([]*Article, int, error)
	GetArticlesBySectionID(ctx context.Context, params *GetArticlesParams) ([]*Article, int, error)
	GetArticleByArticleID(ctx context.Context, articleID int, locale, countryCode string) (*Article, error)
	GetTopNArticles(ctx context.Context, topN uint64, locale, countryCode string) ([]*Article, error)
	PlusOneArticleClickCounter(ctx context.Context, articleID int, locale, countryCode string) error
}

// Article is the article model.
type Article struct {
	SectionID       int       `json:"section_id"`
	ID              int       `json:"id"`
	AuthorID        int       `json:"author_id"`
	CommentsDisable bool      `json:"comments_disable"`
	Draft           bool      `json:"draft"`
	Promoted        bool      `json:"promoted"`
	Position        int       `json:"position"`
	VoteSum         int       `json:"vote_sum"`
	VoteCount       int       `json:"vote_count"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	SourceLocale    string    `json:"source_locale"`
	Outdated        bool      `json:"outdated"`
	OutdatedLocales []string  `json:"outdated_locales"`
	EditedAt        time.Time `json:"edited_at"`
	LabelNames      []string  `json:"label_names"`
	CountryCode     string    `json:"country_code"`
	URL             string    `json:"url"`
	HTMLURL         string    `json:"html_url"`
	Name            string    `json:"name"`
	Title           string    `json:"title"`
	Body            string    `json:"body"`
	Locale          string    `json:"locale"`
}

// SearchArticle is the  search article model which contains two additional properties.
type SearchArticle struct {
	*Article
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
	Snippet      string `json:"snippet"`
}

type articlesOps struct {
	db db.Database
}

// GetArticlesParams is the params structure of requesting GetArticles method.
type GetArticlesParams struct {
	CategoryID  int
	SectionID   int
	Locale      string
	CountryCode string
	PerPage     int
	Page        int
	SortBy      string
	SortOrder   string
}

const (
	updateArticlesQuery = `
	UPDATE articles SET 
		section_id = :section_id,
		author_id = :author_id, 
		comments_disable = :comments_disable, 
		draft = :draft, 
		promoted = :promoted, 
		position = :position, 
		vote_sum = :vote_sum, 
		vote_count = :vote_count, 
		created_at = :created_at,
		updated_at = :updated_at,
		source_locale = :source_locale,
		outdated = :outdated,
		outdated_locales = :outdated_locales,
		edited_at = :edited_at,
		label_names = :label_names
	WHERE id = :id AND country_code = :country_code`

	updateArticleTranslatesQuery = `
	UPDATE article_translates SET
		url = :url,
		html_url = :html_url,
		name = :name,
		title = :title,
		body = :body
	WHERE article_id = :article_id AND locale = :locale`

	insertArticlesQuery = `
	INSERT INTO articles (
		section_id,
		id, 
		author_id, 
		comments_disable, 
		draft, 
		promoted, 
		position, 
		vote_sum, 
		vote_count, 
		created_at,
		updated_at,
		source_locale,
		outdated,
		outdated_locales,
		edited_at,
		label_names,
		country_code
	) 
	VALUES (
		:section_id,
		:id, 
		:author_id, 
		:comments_disable, 
		:draft, 
		:promoted, 
		:position, 
		:vote_sum, 
		:vote_count, 
		:created_at,
		:updated_at,
		:source_locale,
		:outdated,
		:outdated_locales,
		:edited_at,
		:label_names,
		:country_code
	)`

	insertArticleTranslatesQuery = `
	INSERT INTO article_translates (article_id, url, html_url, name, title, body, locale)
	VALUES (:article_id, :url, :html_url, :name, :title, :body, :locale)`

	deleteArticlesQuery          = "DELETE FROM articles WHERE id = :id AND country_code = :country_code"
	deleteArticleTranslatesQuery = "DELETE FROM article_translates WHERE article_id = :article_id AND locale = :locale"
)

// SyncWithArticles ensures the database data will be same as the input data.
func (a *articlesOps) SyncWithArticles(ctx context.Context, zendeskArticles []*Article, countryCode, locale string) error {
	tx, err := a.db.Begin()
	if err != nil {
		return errors.Wrapf(err, "models: [SyncWithArticles] db.Begin failed")
	}

	query := fmt.Sprintf(`SELECT id FROM articles WHERE country_code = '%s'`, countryCode)
	ids := make([]int, 0)
	tx.Select(&ids, query)

	dbIDs := make(map[int]struct{})
	for _, id := range ids {
		dbIDs[id] = struct{}{}
	}

	for _, zendeskArticle := range zendeskArticles {
		dbArticle := &db.Articles{
			SectionID:       zendeskArticle.SectionID,
			AuthorID:        zendeskArticle.AuthorID,
			CommentsDisable: zendeskArticle.CommentsDisable,
			CountryCode:     zendeskArticle.CountryCode,
			CreatedAt:       zendeskArticle.CreatedAt,
			Draft:           zendeskArticle.Draft,
			EditedAt:        zendeskArticle.EditedAt,
			ID:              zendeskArticle.ID,
			LabelNames:      zendeskArticle.LabelNames,
			Outdated:        zendeskArticle.Outdated,
			OutdatedLocales: zendeskArticle.OutdatedLocales,
			Position:        zendeskArticle.Position,
			Promoted:        zendeskArticle.Promoted,
			SourceLocale:    zendeskArticle.SourceLocale,
			UpdatedAt:       zendeskArticle.UpdatedAt,
			VoteCount:       zendeskArticle.VoteCount,
			VoteSum:         zendeskArticle.VoteSum,
		}
		translates := &db.ArticleTranslates{
			ArticleID: zendeskArticle.ID,
			Body:      zendeskArticle.Body,
			HTMLURL:   zendeskArticle.HTMLURL,
			Locale:    zendeskArticle.Locale,
			Name:      zendeskArticle.Name,
			Title:     zendeskArticle.Title,
			URL:       zendeskArticle.URL,
		}

		if _, exist := dbIDs[zendeskArticle.ID]; exist {
			tx.NamedExec(updateArticlesQuery, dbArticle)

			tranIDs := make([]int, 0)
			query = fmt.Sprintf(
				`SELECT article_id FROM article_translates WHERE locale = '%s' AND article_id = '%d'`,
				locale,
				zendeskArticle.ID,
			)
			tx.Select(&tranIDs, query)

			if len(tranIDs) > 0 {
				tx.NamedExec(updateArticleTranslatesQuery, translates)
			} else {
				tx.NamedExec(insertArticleTranslatesQuery, translates)
			}

			delete(dbIDs, zendeskArticle.ID)
		} else {
			tx.NamedExec(insertArticlesQuery, dbArticle)
			tx.NamedExec(insertArticleTranslatesQuery, translates)
		}
	}

	for id := range dbIDs {
		tx.NamedExec(deleteArticleTranslatesQuery, map[string]interface{}{"article_id": id, "locale": locale})

		// Check if there is any article_translates reference to articles.
		total := 0
		query = fmt.Sprintf(`SELECT COUNT(*) FROM article_translates WHERE article_id = '%d'`, id)
		tx.Get(&total, query)
		if total == 0 {
			tx.NamedExec(deleteArticlesQuery, map[string]interface{}{"id": id, "country_code": countryCode})
		}
	}
	tx.Commit()

	return errors.Wrapf(tx.Err(), "models: [SyncWithArticles] db transaction failed")
}

// SyncWithArticle ensures the database data will be same as the input data.
func (a *articlesOps) SyncWithArticle(ctx context.Context, articleID int, zendeskArticle *Article, countryCode, locale string) error {
	tx, err := a.db.Begin()
	if err != nil {
		return errors.Wrapf(err, "models: [SyncWithArticle] db.Begin failed")
	}

	query := fmt.Sprintf(`SELECT id FROM articles WHERE id = %d AND country_code = '%s'`, articleID, countryCode)
	ids := make([]int, 0)
	tx.Select(&ids, query)

	dbIDs := make(map[int]struct{})
	for _, id := range ids {
		dbIDs[id] = struct{}{}
	}

	if zendeskArticle != nil {
		dbArticle := &db.Articles{
			SectionID:       zendeskArticle.SectionID,
			AuthorID:        zendeskArticle.AuthorID,
			CommentsDisable: zendeskArticle.CommentsDisable,
			CountryCode:     zendeskArticle.CountryCode,
			CreatedAt:       zendeskArticle.CreatedAt,
			Draft:           zendeskArticle.Draft,
			EditedAt:        zendeskArticle.EditedAt,
			ID:              zendeskArticle.ID,
			LabelNames:      zendeskArticle.LabelNames,
			Outdated:        zendeskArticle.Outdated,
			OutdatedLocales: zendeskArticle.OutdatedLocales,
			Position:        zendeskArticle.Position,
			Promoted:        zendeskArticle.Promoted,
			SourceLocale:    zendeskArticle.SourceLocale,
			UpdatedAt:       zendeskArticle.UpdatedAt,
			VoteCount:       zendeskArticle.VoteCount,
			VoteSum:         zendeskArticle.VoteSum,
		}
		translates := &db.ArticleTranslates{
			ArticleID: zendeskArticle.ID,
			Body:      zendeskArticle.Body,
			HTMLURL:   zendeskArticle.HTMLURL,
			Locale:    zendeskArticle.Locale,
			Name:      zendeskArticle.Name,
			Title:     zendeskArticle.Title,
			URL:       zendeskArticle.URL,
		}

		if _, exist := dbIDs[zendeskArticle.ID]; exist {
			tx.NamedExec(updateArticlesQuery, dbArticle)

			tranIDs := make([]int, 0)
			query = fmt.Sprintf(
				`SELECT article_id FROM article_translates WHERE locale = '%s' AND article_id = '%d'`,
				locale,
				zendeskArticle.ID,
			)
			tx.Select(&tranIDs, query)

			if len(tranIDs) > 0 {
				tx.NamedExec(updateArticleTranslatesQuery, translates)
			} else {
				tx.NamedExec(insertArticleTranslatesQuery, translates)
			}

			delete(dbIDs, zendeskArticle.ID)
		} else {
			tx.NamedExec(insertArticlesQuery, dbArticle)
			tx.NamedExec(insertArticleTranslatesQuery, translates)
		}
	}

	for id := range dbIDs {
		tx.NamedExec(deleteArticleTranslatesQuery, map[string]interface{}{"article_id": id, "locale": locale})

		// Check if there is any article_translates reference to articles.
		total := 0
		query = fmt.Sprintf(`SELECT COUNT(*) FROM article_translates WHERE article_id = '%d'`, id)
		tx.Get(&total, query)
		if total == 0 {
			tx.NamedExec(deleteArticlesQuery, map[string]interface{}{"id": id, "country_code": countryCode})
		}
	}
	tx.Commit()

	return errors.Wrapf(tx.Err(), "models: [SyncWithArticle] db transaction failed")
}

func (a *articlesOps) GetArticles(ctx context.Context, params *GetArticlesParams) ([]*Article, int, error) {
	articles := make([]*db.Articles, 0)
	query := fmt.Sprintf(
		`SELECT section_id,id,author_id,comments_disable,draft,promoted,position,
		vote_sum,vote_count,created_at,updated_at,source_locale,outdated,
		outdated_locales,edited_at,label_names,country_code 
		FROM articles WHERE country_code = '%s' 
		ORDER BY %s %s, created_at DESC LIMIT %d OFFSET %d`,
		params.CountryCode,
		params.SortBy,
		params.SortOrder,
		params.PerPage,
		params.Page,
	)
	if err := a.db.Select(ctx, &articles, query); err != nil {
		return nil, 0, errors.Wrapf(err, "models: [GetArticles] db select articles failed")
	}

	ret := make([]*Article, 0)
	for _, article := range articles {
		translate := new(db.ArticleTranslates)
		query = fmt.Sprintf(
			`SELECT url,html_url,name,title,body,locale FROM article_translates WHERE article_id = %d AND locale = '%s'`,
			article.ID,
			params.Locale,
		)
		if err := a.db.Get(ctx, translate, query); err != nil {
			if err == db.ErrNoRows {
				continue
			}
			return nil, 0, errors.Wrapf(err, "models: [GetArticles] db get translates failed")
		}

		ret = append(ret, &Article{
			SectionID:       article.SectionID,
			ID:              article.ID,
			AuthorID:        article.AuthorID,
			CommentsDisable: article.CommentsDisable,
			Draft:           article.Draft,
			Promoted:        article.Promoted,
			Position:        article.Position,
			VoteSum:         article.VoteSum,
			VoteCount:       article.VoteCount,
			CreatedAt:       article.CreatedAt,
			UpdatedAt:       article.UpdatedAt,
			SourceLocale:    article.SourceLocale,
			Outdated:        article.Outdated,
			OutdatedLocales: article.OutdatedLocales,
			EditedAt:        article.EditedAt,
			LabelNames:      article.LabelNames,
			CountryCode:     article.CountryCode,
			URL:             translate.URL,
			HTMLURL:         translate.HTMLURL,
			Name:            translate.Name,
			Title:           translate.Title,
			Body:            translate.Body,
			Locale:          translate.Locale,
		})
	}

	total := 0
	query = fmt.Sprintf(
		`SELECT COUNT(*) FROM articles INNER JOIN article_translates 
		ON articles.id=article_translates.article_id 
		WHERE country_code = '%s' AND locale = '%s'`,
		params.CountryCode, params.Locale)
	if err := a.db.Get(ctx, &total, query); err != nil {
		return nil, 0, errors.Wrapf(err, "models: [GetArticles] db get total failed")
	}

	return ret, total, nil
}

func (c *categoriesOps) GetArticlesByCategoryID(ctx context.Context, params *GetArticlesParams, labels []string) ([]*Article, int, error) {
	//SELECT label_names FROM articles WHERE (label_names::text LIKE '{%confirmed%}'
	// AND label_names::text LIKE '{%preparing%}' AND label_names::text LIKE '{%ontheway%}'
	// AND label_names::text LIKE '{%delivered%}') AND country_code = 'sg';
	queryLabelNames := ""
	if len(labels) > 0 {
		queryLabelNames += "("
		for _, label := range labels {
			queryLabelNames += "label_names::text LIKE '{%" + label + "%}' OR "
		}
		queryLabelNames = strings.TrimRight(queryLabelNames, " OR ")
		queryLabelNames += ") AND"
	}

	articles := make([]*db.Articles, 0)
	query := fmt.Sprintf(
		`SELECT section_id,id,author_id,comments_disable,draft,promoted,position,
		vote_sum,vote_count,created_at,updated_at,source_locale,outdated,
		outdated_locales,edited_at,label_names,country_code 
		FROM articles WHERE %s country_code = '%s' 
		ORDER BY %s %s, created_at DESC LIMIT %d OFFSET %d`,
		queryLabelNames,
		params.CountryCode,
		params.SortBy,
		params.SortOrder,
		params.PerPage,
		params.Page,
	)
	if err := c.db.Select(ctx, &articles, query); err != nil {
		return nil, 0, errors.Wrapf(err, "models: [GetArticlesByCategoryID] db select articles failed")
	}

	ret := make([]*Article, 0)
	for _, article := range articles {
		translates := new(db.ArticleTranslates)
		query = fmt.Sprintf(
			`SELECT url,html_url,name,title,body,locale 
			FROM article_translates WHERE article_id = '%d' AND locale = '%s'`,
			article.ID,
			params.Locale,
		)
		if err := c.db.Get(ctx, translates, query); err != nil {
			if err == db.ErrNoRows {
				continue
			}
			return nil, 0, errors.Wrapf(err, "models: [GetArticlesByCategoryID] db get translates failed")
		}

		ret = append(ret, &Article{
			SectionID:       article.SectionID,
			AuthorID:        article.AuthorID,
			Body:            translates.Body,
			CommentsDisable: article.CommentsDisable,
			CountryCode:     article.CountryCode,
			CreatedAt:       article.CreatedAt,
			Draft:           article.Draft,
			EditedAt:        article.EditedAt,
			HTMLURL:         translates.HTMLURL,
			ID:              article.ID,
			LabelNames:      article.LabelNames,
			Locale:          translates.Locale,
			Name:            translates.Name,
			Outdated:        article.Outdated,
			OutdatedLocales: article.OutdatedLocales,
			Position:        article.Position,
			Promoted:        article.Promoted,
			SourceLocale:    article.SourceLocale,
			Title:           translates.Title,
			UpdatedAt:       article.UpdatedAt,
			URL:             translates.URL,
			VoteCount:       article.VoteCount,
			VoteSum:         article.VoteSum,
		})
	}

	total := 0
	query = fmt.Sprintf(
		`SELECT COUNT(*) FROM articles INNER JOIN article_translates 
		ON articles.id=article_translates.article_id 
		WHERE %s country_code = '%s' AND locale = '%s'`,
		queryLabelNames, params.CountryCode, params.Locale,
	)
	if err := c.db.Get(ctx, &total, query); err != nil {
		return nil, 0, errors.Wrapf(err, "models: [GetArticlesByCategoryID] db get total failed")
	}

	return ret, total, nil
}

// GetArticlesBySectionID get articles with params.
func (s *sectionsOps) GetArticlesBySectionID(ctx context.Context, params *GetArticlesParams) ([]*Article, int, error) {
	articles := make([]*db.Articles, 0)
	query := fmt.Sprintf(
		`SELECT section_id,id,author_id,comments_disable,draft,promoted,position,
		vote_sum,vote_count,created_at,updated_at,source_locale,outdated,
		outdated_locales,edited_at,label_names,country_code 
		FROM articles WHERE country_code = '%s' AND section_id = '%d' 
		ORDER BY %s %s, created_at DESC LIMIT %d OFFSET %d`,
		params.CountryCode,
		params.SectionID,
		params.SortBy,
		params.SortOrder,
		params.PerPage,
		params.Page,
	)
	if err := s.db.Select(ctx, &articles, query); err != nil {
		return nil, 0, errors.Wrapf(err, "models: [GetArticlesBySectionID] db select articles failed")
	}

	ret := make([]*Article, 0)
	for _, article := range articles {
		translates := new(db.ArticleTranslates)
		query = fmt.Sprintf(
			`SELECT url,html_url,name,title,body,locale 
			FROM article_translates WHERE article_id = '%d' AND locale = '%s'`,
			article.ID,
			params.Locale,
		)
		if err := s.db.Get(ctx, translates, query); err != nil {
			if err == db.ErrNoRows {
				continue
			}
			return nil, 0, errors.Wrapf(err, "models: [GetArticlesBySectionID] db get translates failed")
		}

		ret = append(ret, &Article{
			SectionID:       article.SectionID,
			AuthorID:        article.AuthorID,
			Body:            translates.Body,
			CommentsDisable: article.CommentsDisable,
			CountryCode:     article.CountryCode,
			CreatedAt:       article.CreatedAt,
			Draft:           article.Draft,
			EditedAt:        article.EditedAt,
			HTMLURL:         translates.HTMLURL,
			ID:              article.ID,
			LabelNames:      article.LabelNames,
			Locale:          translates.Locale,
			Name:            translates.Name,
			Outdated:        article.Outdated,
			OutdatedLocales: article.OutdatedLocales,
			Position:        article.Position,
			Promoted:        article.Promoted,
			SourceLocale:    article.SourceLocale,
			Title:           translates.Title,
			UpdatedAt:       article.UpdatedAt,
			URL:             translates.URL,
			VoteCount:       article.VoteCount,
			VoteSum:         article.VoteSum,
		})
	}

	total := 0
	query = fmt.Sprintf(
		`SELECT COUNT(*) FROM articles INNER JOIN article_translates 
		ON articles.id=article_translates.article_id 
		WHERE country_code = '%s' AND locale = '%s' AND section_id = '%d'`,
		params.CountryCode, params.Locale, params.SectionID,
	)
	if err := s.db.Get(ctx, &total, query); err != nil {
		return nil, 0, errors.Wrapf(err, "models: [GetArticlesBySectionID] db get total failed")
	}

	return ret, total, nil
}

func (a *articlesOps) GetArticleByArticleID(ctx context.Context, articleID int, locale, countryCode string) (*Article, error) {
	article := new(db.Articles)
	query := fmt.Sprintf(
		`SELECT section_id,id,author_id,comments_disable,draft,promoted,position,
		vote_sum,vote_count,created_at,updated_at,source_locale,outdated,
		outdated_locales,edited_at,label_names,country_code 
		FROM articles WHERE country_code = '%s' AND id = '%d'`,
		countryCode,
		articleID,
	)
	if err := a.db.Get(ctx, article, query); err != nil {
		switch err {
		case db.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, errors.Wrapf(err, "models: [GetArticle] db get article failed")
		}
	}

	translates := new(db.ArticleTranslates)
	query = fmt.Sprintf(
		`SELECT url,html_url,name,title,body,locale 
		FROM article_translates WHERE article_id = '%d' AND locale = '%s'`,
		article.ID,
		locale,
	)
	if err := a.db.Get(ctx, translates, query); err != nil {
		switch err {
		case db.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, errors.Wrapf(err, "models: [GetArticle] db get translates failed")
		}
	}

	ret := &Article{
		SectionID:       article.SectionID,
		AuthorID:        article.AuthorID,
		Body:            translates.Body,
		CommentsDisable: article.CommentsDisable,
		CountryCode:     article.CountryCode,
		CreatedAt:       article.CreatedAt,
		Draft:           article.Draft,
		EditedAt:        article.EditedAt,
		HTMLURL:         translates.HTMLURL,
		ID:              article.ID,
		LabelNames:      article.LabelNames,
		Locale:          translates.Locale,
		Name:            translates.Name,
		Outdated:        article.Outdated,
		OutdatedLocales: article.OutdatedLocales,
		Position:        article.Position,
		Promoted:        article.Promoted,
		SourceLocale:    article.SourceLocale,
		Title:           translates.Title,
		UpdatedAt:       article.UpdatedAt,
		URL:             translates.URL,
		VoteCount:       article.VoteCount,
		VoteSum:         article.VoteSum,
	}

	return ret, nil
}

// PlusOneArticleClickCounter plus one on article click count.
func (a *articlesOps) PlusOneArticleClickCounter(ctx context.Context, articleID int, locale, countryCode string) error {
	dbArticle := &db.Articles{
		CountryCode: countryCode,
		ID:          articleID,
	}

	_, err := a.db.NamedExec(ctx,
		`UPDATE articles SET click_count = click_count + 1 WHERE id = :id AND country_code = :country_code`,
		dbArticle)

	if err != nil {
		return errors.Wrapf(err, "models: [PlusOneArticleClickCounter] update db article click count failed")
	}
	return nil
}

// GetTopNArticles get topN articles.
func (a *articlesOps) GetTopNArticles(ctx context.Context, topN uint64, locale, countryCode string) ([]*Article, error) {
	articles := make([]*db.Articles, 0)
	// Sorts with promoted=t and click_count descend order, also limit topN.
	query := fmt.Sprintf(
		`SELECT section_id,id,author_id,comments_disable,draft,promoted,position,
		vote_sum,vote_count,created_at,updated_at,source_locale,outdated,outdated_locales,
		edited_at,label_names,country_code 
		FROM articles WHERE country_code = '%s' 
		ORDER BY promoted DESC, click_count DESC LIMIT %d`,
		countryCode,
		topN,
	)

	if err := a.db.Select(ctx, &articles, query); err != nil {
		return nil, errors.Wrapf(err, "models: [GetTopNArticles] db select articles failed")
	}

	ret := make([]*Article, 0)
	for _, article := range articles {
		translate := new(db.ArticleTranslates)
		query = fmt.Sprintf(
			`SELECT url,html_url,name,title,body,locale 
			FROM article_translates WHERE article_id = '%d' AND locale = '%s'`,
			article.ID,
			locale,
		)

		if err := a.db.Get(ctx, translate, query); err != nil {
			if err == db.ErrNoRows {
				continue
			}
			return nil, errors.Wrapf(err, "models: [GetTopNArticles] db get translate failed")
		}

		ret = append(ret, &Article{
			SectionID:       article.SectionID,
			AuthorID:        article.AuthorID,
			Body:            translate.Body,
			CommentsDisable: article.CommentsDisable,
			CountryCode:     article.CountryCode,
			CreatedAt:       article.CreatedAt,
			Draft:           article.Draft,
			EditedAt:        article.EditedAt,
			HTMLURL:         translate.HTMLURL,
			ID:              article.ID,
			LabelNames:      article.LabelNames,
			Locale:          translate.Locale,
			Name:            translate.Name,
			Outdated:        article.Outdated,
			OutdatedLocales: article.OutdatedLocales,
			Position:        article.Position,
			Promoted:        article.Promoted,
			SourceLocale:    article.SourceLocale,
			Title:           translate.Title,
			UpdatedAt:       article.UpdatedAt,
			URL:             translate.URL,
			VoteCount:       article.VoteCount,
			VoteSum:         article.VoteSum,
		})
	}

	return ret, nil
}
