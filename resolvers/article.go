package resolvers

import (
	"context"
	"strconv"

	gographql "github.com/graph-gophers/graphql-go"

	"github.com/honestbee/Zen/dataloader"
	"github.com/honestbee/Zen/inout"
	"github.com/honestbee/Zen/models"
)

// ArticlesResolver defines resolver models.
type ArticlesResolver struct {
	m *inout.GetArticlesOut
}

// Articles is the Articles's field categories.
func (r *ArticlesResolver) Articles(ctx context.Context) *[]*ArticleResolver {
	ret := make([]*ArticleResolver, 0)
	for _, category := range r.m.Articles {
		ret = append(ret, &ArticleResolver{m: category})
	}
	return &ret
}

// Page is the Articles's field page.
func (r *ArticlesResolver) Page(ctx context.Context) int32 {
	return int32(r.m.Page)
}

// PerPage is the Articles's field per_page.
func (r *ArticlesResolver) PerPage(ctx context.Context) int32 {
	return int32(r.m.PerPage)
}

// PageCount is the Articles's field page_count.
func (r *ArticlesResolver) PageCount(ctx context.Context) int32 {
	return int32(r.m.PageCount)
}

// Count is the Articles's field count.
func (r *ArticlesResolver) Count(ctx context.Context) int32 {
	return int32(r.m.Count)
}

// ArticleResolver defines resolver models.
type ArticleResolver struct {
	m *models.Article
}

// ID is the Article's field id.
func (r *ArticleResolver) ID(ctx context.Context) gographql.ID {
	return gographql.ID(strconv.Itoa(r.m.ID))
}

// AuthorID is the Article's field author_id.
func (r *ArticleResolver) AuthorID(ctx context.Context) string {
	return strconv.Itoa(r.m.AuthorID)
}

// CommentsDisable is the Article's field comments_disable.
func (r *ArticleResolver) CommentsDisable(ctx context.Context) bool {
	return r.m.CommentsDisable
}

// Draft is the Article's field draft.
func (r *ArticleResolver) Draft(ctx context.Context) bool {
	return r.m.Draft
}

// Promoted is the Article's field promoted.
func (r *ArticleResolver) Promoted(ctx context.Context) bool {
	return r.m.Promoted
}

// Position is the Article's field position.
func (r *ArticleResolver) Position(ctx context.Context) int32 {
	return int32(r.m.Position)
}

// VoteSum is the Article's field vote_sum.
func (r *ArticleResolver) VoteSum(ctx context.Context) int32 {
	return int32(r.m.VoteSum)
}

// VoteCount is the Article's field vote_count.
func (r *ArticleResolver) VoteCount(ctx context.Context) int32 {
	return int32(r.m.VoteCount)
}

// CreatedAt is the Article's field created_at.
func (r *ArticleResolver) CreatedAt(ctx context.Context) gographql.Time {
	return gographql.Time{Time: r.m.CreatedAt}
}

// UpdatedAt is the Article's field updated_at.
func (r *ArticleResolver) UpdatedAt(ctx context.Context) gographql.Time {
	return gographql.Time{Time: r.m.UpdatedAt}
}

// SourceLocale is the Article's field source_locale.
func (r *ArticleResolver) SourceLocale(ctx context.Context) string {
	return r.m.SourceLocale
}

// Outdated is the Article's field outdated.
func (r *ArticleResolver) Outdated(ctx context.Context) bool {
	return r.m.Outdated
}

// OutdatedLocales is the Article's field outdated_locales.
func (r *ArticleResolver) OutdatedLocales(ctx context.Context) []string {
	return r.m.OutdatedLocales
}

// EditedAt is the Article's field edited_at.
func (r *ArticleResolver) EditedAt(ctx context.Context) gographql.Time {
	return gographql.Time{Time: r.m.EditedAt}
}

// LabelNames is the Article's field label_names.
func (r *ArticleResolver) LabelNames(ctx context.Context) []string {
	return r.m.LabelNames
}

// CountryCode is the Article's field country_code.
func (r *ArticleResolver) CountryCode(ctx context.Context) string {
	return r.m.CountryCode
}

// URL is the Article's field url.
func (r *ArticleResolver) URL(ctx context.Context) string {
	return r.m.URL
}

// HTMLURL is the Article's field html_url.
func (r *ArticleResolver) HTMLURL(ctx context.Context) string {
	return r.m.HTMLURL
}

// Name is the Article's field name.
func (r *ArticleResolver) Name(ctx context.Context) string {
	return r.m.Name
}

// Title is the Article's field title.
func (r *ArticleResolver) Title(ctx context.Context) string {
	return r.m.Title
}

// Body is the Article's field body.
func (r *ArticleResolver) Body(ctx context.Context) string {
	return r.m.Body
}

// Locale is the Article's field locale.
func (r *ArticleResolver) Locale(ctx context.Context) string {
	return r.m.Locale
}

// CategoryConnection is the Article's field category.
func (r *ArticleResolver) CategoryConnection(ctx context.Context) (*CategoryResolver, error) {
	sectionArgs := inout.QuerySectionIn{
		SectionID:   gographql.ID(strconv.Itoa(r.m.SectionID)),
		CountryCode: r.m.CountryCode,
		Locale:      r.m.Locale,
	}

	// Load section.
	loadSection, err := dataloader.LoadSection(ctx, sectionArgs)
	if err != nil {
		return nil, err
	}

	data := inout.QueryCategoryIn{
		CategoryIDOrKeyName: gographql.ID(strconv.Itoa(loadSection.CategoryID)),
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

// SectionConnection is the Article's field section.
func (r *ArticleResolver) SectionConnection(ctx context.Context) (*SectionResolver, error) {
	data := inout.QuerySectionIn{
		SectionID:   gographql.ID(strconv.Itoa(r.m.SectionID)),
		CountryCode: r.m.CountryCode,
		Locale:      r.m.Locale,
	}

	// Load section.
	result, err := dataloader.LoadSection(ctx, data)
	if err != nil {
		return nil, err
	}
	return &SectionResolver{m: result}, nil
}
