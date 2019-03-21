package resolvers

import (
	"context"
	"strconv"

	gographql "github.com/graph-gophers/graphql-go"

	"github.com/honestbee/Zen/dataloader"
	"github.com/honestbee/Zen/inout"
	"github.com/honestbee/Zen/models"
)

// SectionsResolver defines resolver models.
type SectionsResolver struct {
	m *inout.GetSectionsOut
}

// Sections is the Sections's field categories.
func (r *SectionsResolver) Sections(ctx context.Context) *[]*SectionResolver {
	ret := make([]*SectionResolver, 0)
	for _, section := range r.m.Sections {
		ret = append(ret, &SectionResolver{m: section})
	}
	return &ret
}

// Page is the Sections's field page.
func (r *SectionsResolver) Page(ctx context.Context) int32 {
	return int32(r.m.Page)
}

// PerPage is the Sections's field per_page.
func (r *SectionsResolver) PerPage(ctx context.Context) int32 {
	return int32(r.m.PerPage)
}

// PageCount is the Sections's field page_count.
func (r *SectionsResolver) PageCount(ctx context.Context) int32 {
	return int32(r.m.PageCount)
}

// Count is the Sections's field count.
func (r *SectionsResolver) Count(ctx context.Context) int32 {
	return int32(r.m.Count)
}

// SectionResolver defines resolver models.
type SectionResolver struct {
	m *models.Section
}

// ID is the Section's field id.
func (r *SectionResolver) ID(ctx context.Context) gographql.ID {
	return gographql.ID(strconv.Itoa(r.m.ID))
}

// Position is the Section's field position.
func (r *SectionResolver) Position(ctx context.Context) int32 {
	return int32(r.m.Position)
}

// CreatedAt is the Section's field created_at.
func (r *SectionResolver) CreatedAt(ctx context.Context) gographql.Time {
	return gographql.Time{Time: r.m.CreatedAt}
}

// UpdatedAt is the Section's field updated_at.
func (r *SectionResolver) UpdatedAt(ctx context.Context) gographql.Time {
	return gographql.Time{Time: r.m.UpdatedAt}
}

// SourceLocale is the Section's field source_locale.
func (r *SectionResolver) SourceLocale(ctx context.Context) string {
	return r.m.SourceLocale
}

// Outdated is the Section's field outdated.
func (r *SectionResolver) Outdated(ctx context.Context) bool {
	return r.m.Outdated
}

// CountryCode is the Section's field country_code.
func (r *SectionResolver) CountryCode(ctx context.Context) string {
	return r.m.CountryCode
}

// URL is the Section's field url.
func (r *SectionResolver) URL(ctx context.Context) string {
	return r.m.URL
}

// HTMLURL is the Section's field html_url.
func (r *SectionResolver) HTMLURL(ctx context.Context) string {
	return r.m.HTMLURL
}

// Description is the Section's field description.
func (r *SectionResolver) Description(ctx context.Context) string {
	return r.m.Description
}

// Locale is the Section's field locale.
func (r *SectionResolver) Locale(ctx context.Context) string {
	return r.m.Locale
}

// Name is the Section's field name.
func (r *SectionResolver) Name(ctx context.Context) string {
	return r.m.Name
}

// CategoryConnection is the Section's field category.
func (r *SectionResolver) CategoryConnection(ctx context.Context) (*CategoryResolver, error) {
	id := gographql.ID(strconv.Itoa(r.m.CategoryID))
	data := inout.QueryCategoryIn{
		CategoryIDOrKeyName: id,
		CountryCode:         r.m.CountryCode,
		Locale:              r.m.Locale,
	}

	// Load categories.
	result, err := dataloader.LoadCategory(ctx, data)
	if err != nil {
		return nil, err
	}
	return &CategoryResolver{m: result}, nil
}

// ArticlesConnection is the Section's field articles.
func (r *SectionResolver) ArticlesConnection(ctx context.Context, data inout.QueryArticlesIn) (*ArticlesResolver, error) {
	// Process input params. Ingore error for connection since country code
	// and locale are snake_case already.
	_ = data.ProcessInputParams()

	id := gographql.ID(strconv.Itoa(r.m.ID))
	data.SectionID = &id
	data.CountryCode = r.m.CountryCode
	data.Locale = r.m.Locale

	// Load articles.
	result, err := dataloader.LoadArticles(ctx, data)
	if err != nil {
		return nil, err
	}
	return &ArticlesResolver{m: result}, nil
}
