package resolvers

import (
	"context"
	"strconv"

	gographql "github.com/graph-gophers/graphql-go"

	"github.com/honestbee/Zen/dataloader"
	"github.com/honestbee/Zen/inout"
	"github.com/honestbee/Zen/models"
)

// SearchBodyArticlesResolver defines resolver models.
type SearchBodyArticlesResolver struct {
	m *inout.GetSearchOut
}

// Articles is the SearchBodyArticles's field articles.
func (r *SearchBodyArticlesResolver) Articles(ctx context.Context) *[]*SearchBodyArticleResolver {
	ret := make([]*SearchBodyArticleResolver, 0)
	for _, article := range r.m.Articles {
		ret = append(ret, &SearchBodyArticleResolver{m: article})
	}
	return &ret
}

// Page is the SearchBodyArticles's field page.
func (r *SearchBodyArticlesResolver) Page(ctx context.Context) int32 {
	return int32(r.m.Page)
}

// PerPage is the SearchBodyArticles's field per_page.
func (r *SearchBodyArticlesResolver) PerPage(ctx context.Context) int32 {
	return int32(r.m.PerPage)
}

// PageCount is the SearchBodyArticles's field page_count.
func (r *SearchBodyArticlesResolver) PageCount(ctx context.Context) int32 {
	return int32(r.m.PageCount)
}

// Count is the SearchBodyArticles's field count.
func (r *SearchBodyArticlesResolver) Count(ctx context.Context) int32 {
	return int32(r.m.Count)
}

// SearchBodyArticleResolver defines resolver models.
type SearchBodyArticleResolver struct {
	m *models.SearchArticle
}

// ID is the SearchArticle's field id.
func (r *SearchBodyArticleResolver) ID(ctx context.Context) gographql.ID {
	return gographql.ID(strconv.Itoa(r.m.ID))
}

// AuthorID is the SearchArticle's field author_id.
func (r *SearchBodyArticleResolver) AuthorID(ctx context.Context) string {
	return strconv.Itoa(r.m.AuthorID)
}

// CommentsDisable is the SearchArticle's field comments_disable.
func (r *SearchBodyArticleResolver) CommentsDisable(ctx context.Context) bool {
	return r.m.CommentsDisable
}

// Draft is the SearchArticle's field draft.
func (r *SearchBodyArticleResolver) Draft(ctx context.Context) bool {
	return r.m.Draft
}

// Promoted is the SearchArticle's field promoted.
func (r *SearchBodyArticleResolver) Promoted(ctx context.Context) bool {
	return r.m.Promoted
}

// Position is the SearchArticle's field position.
func (r *SearchBodyArticleResolver) Position(ctx context.Context) int32 {
	return int32(r.m.Position)
}

// VoteSum is the SearchArticle's field vote_sum.
func (r *SearchBodyArticleResolver) VoteSum(ctx context.Context) int32 {
	return int32(r.m.VoteSum)
}

// VoteCount is the SearchArticle's field vote_count.
func (r *SearchBodyArticleResolver) VoteCount(ctx context.Context) int32 {
	return int32(r.m.VoteCount)
}

// CreatedAt is the SearchArticle's field created_at.
func (r *SearchBodyArticleResolver) CreatedAt(ctx context.Context) gographql.Time {
	return gographql.Time{Time: r.m.CreatedAt}
}

// UpdatedAt is the SearchArticle's field updated_at.
func (r *SearchBodyArticleResolver) UpdatedAt(ctx context.Context) gographql.Time {
	return gographql.Time{Time: r.m.UpdatedAt}
}

// SourceLocale is the SearchArticle's field source_locale.
func (r *SearchBodyArticleResolver) SourceLocale(ctx context.Context) string {
	return r.m.SourceLocale
}

// Outdated is the SearchArticle's field outdated.
func (r *SearchBodyArticleResolver) Outdated(ctx context.Context) bool {
	return r.m.Outdated
}

// OutdatedLocales is the SearchArticle's field outdated_locales.
func (r *SearchBodyArticleResolver) OutdatedLocales(ctx context.Context) []string {
	return r.m.OutdatedLocales
}

// EditedAt is the SearchArticle's field edited_at.
func (r *SearchBodyArticleResolver) EditedAt(ctx context.Context) gographql.Time {
	return gographql.Time{Time: r.m.EditedAt}
}

// LabelNames is the SearchArticle's field label_names.
func (r *SearchBodyArticleResolver) LabelNames(ctx context.Context) []string {
	return r.m.LabelNames
}

// CountryCode is the SearchArticle's field country_code.
func (r *SearchBodyArticleResolver) CountryCode(ctx context.Context) string {
	return r.m.CountryCode
}

// URL is the SearchArticle's field url.
func (r *SearchBodyArticleResolver) URL(ctx context.Context) string {
	return r.m.URL
}

// HTMLURL is the SearchArticle's field html_url.
func (r *SearchBodyArticleResolver) HTMLURL(ctx context.Context) string {
	return r.m.HTMLURL
}

// Name is the SearchArticle's field name.
func (r *SearchBodyArticleResolver) Name(ctx context.Context) string {
	return r.m.Name
}

// Title is the SearchArticle's field title.
func (r *SearchBodyArticleResolver) Title(ctx context.Context) string {
	return r.m.Title
}

// Body is the SearchArticle's field body.
func (r *SearchBodyArticleResolver) Body(ctx context.Context) string {
	return r.m.Body
}

// Locale is the SearchArticle's field locale.
func (r *SearchBodyArticleResolver) Locale(ctx context.Context) string {
	return r.m.Locale
}

// Snippet is the SearchArticle's field snippet.
func (r *SearchBodyArticleResolver) Snippet(ctx context.Context) string {
	return r.m.Snippet
}

// CategoryConnection is the SearchArticle's field category.
func (r *SearchBodyArticleResolver) CategoryConnection(ctx context.Context) (*CategoryResolver, error) {
	data := inout.QueryCategoryIn{
		CategoryIDOrKeyName: gographql.ID(strconv.Itoa(r.m.CategoryID)),
		CountryCode:         r.m.CountryCode,
		Locale:              r.m.Locale,
	}

	// Load category.
	result, err := dataloader.LoadCategory(ctx, data)
	if err != nil {
		return nil, err
	}
	return &CategoryResolver{m: result}, nil
}

// SectionConnection is the SearchArticle's field section.
func (r *SearchBodyArticleResolver) SectionConnection(ctx context.Context) (*SectionResolver, error) {
	data := inout.QuerySectionIn{
		SectionID:   gographql.ID(strconv.Itoa(r.m.SectionID)),
		CountryCode: r.m.CountryCode,
		Locale:      r.m.Locale,
	}

	// Load sections.
	result, err := dataloader.LoadSection(ctx, data)
	if err != nil {
		return nil, err
	}
	return &SectionResolver{m: result}, nil
}
