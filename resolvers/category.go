package resolvers

import (
	"context"
	"strconv"

	gographql "github.com/graph-gophers/graphql-go"

	"github.com/honestbee/Zen/dataloader"
	"github.com/honestbee/Zen/inout"
	"github.com/honestbee/Zen/models"
)

// CategoriesResolver defines resolver models.
type CategoriesResolver struct {
	m *inout.GetCategoriesOut
}

// Categories is the Categories's field categories.
func (r *CategoriesResolver) Categories(ctx context.Context) *[]*CategoryResolver {
	ret := make([]*CategoryResolver, 0)
	for _, category := range r.m.Categories {
		ret = append(ret, &CategoryResolver{m: category})
	}
	return &ret
}

// Page is the Categories's field page.
func (r *CategoriesResolver) Page(ctx context.Context) int32 {
	return int32(r.m.Page)
}

// PerPage is the Categories's field per_page.
func (r *CategoriesResolver) PerPage(ctx context.Context) int32 {
	return int32(r.m.PerPage)
}

// PageCount is the Categories's field page_count.
func (r *CategoriesResolver) PageCount(ctx context.Context) int32 {
	return int32(r.m.PageCount)
}

// Count is the Categories's field count.
func (r *CategoriesResolver) Count(ctx context.Context) int32 {
	return int32(r.m.Count)
}

// CategoryResolver defines resolver models.
type CategoryResolver struct {
	m *models.Category
}

// ID is the Category's field id.
func (r *CategoryResolver) ID(ctx context.Context) gographql.ID {
	return gographql.ID(strconv.Itoa(r.m.ID))
}

// Position is the Category's field position.
func (r *CategoryResolver) Position(ctx context.Context) int32 {
	return int32(r.m.Position)
}

// CreatedAt is the Category's field created_at.
func (r *CategoryResolver) CreatedAt(ctx context.Context) gographql.Time {
	return gographql.Time{Time: r.m.CreatedAt}
}

// UpdatedAt is the Category's field updated_at.
func (r *CategoryResolver) UpdatedAt(ctx context.Context) gographql.Time {
	return gographql.Time{Time: r.m.UpdatedAt}
}

// SourceLocale is the Category's field source_locale.
func (r *CategoryResolver) SourceLocale(ctx context.Context) string {
	return r.m.SourceLocale
}

// Outdated is the Category's field outdated.
func (r *CategoryResolver) Outdated(ctx context.Context) bool {
	return r.m.Outdated
}

// CountryCode is the Category's field country_code.
func (r *CategoryResolver) CountryCode(ctx context.Context) string {
	return r.m.CountryCode
}

// KeyName is the Category's field key_name.
func (r *CategoryResolver) KeyName(ctx context.Context) string {
	return r.m.KeyName
}

// URL is the Category's field url.
func (r *CategoryResolver) URL(ctx context.Context) string {
	return r.m.URL
}

// HTMLURL is the Category's field html_url.
func (r *CategoryResolver) HTMLURL(ctx context.Context) string {
	return r.m.HTMLURL
}

// Name is the Category's field name.
func (r *CategoryResolver) Name(ctx context.Context) string {
	return r.m.Name
}

// Description is the Category's field description.
func (r *CategoryResolver) Description(ctx context.Context) string {
	return r.m.Description
}

// Locale is the Category's field locale.
func (r *CategoryResolver) Locale(ctx context.Context) string {
	return r.m.Locale
}

// SectionsConnection is the Category's field sections.
func (r *CategoryResolver) SectionsConnection(ctx context.Context, data inout.QuerySectionsIn) (*SectionsResolver, error) {
	// Process input params. Ingore error for connection since country code
	// and locale are snake_case already.
	_ = data.ProcessInputParams()

	id := gographql.ID(strconv.Itoa(r.m.ID))
	data.CategoryID = &id
	data.CountryCode = r.m.CountryCode
	data.Locale = r.m.Locale

	// Load sections.
	result, err := dataloader.LoadSections(ctx, data)
	if err != nil {
		return nil, err
	}
	return &SectionsResolver{m: result}, nil
}

// ArticlesConnection is the Category's field articles.
func (r *CategoryResolver) ArticlesConnection(ctx context.Context, data inout.QueryArticlesIn) (*ArticlesResolver, error) {
	// Process input params. Ingore error for connection since country code
	// and locale are snake_case already.
	_ = data.ProcessInputParams()

	id := gographql.ID(strconv.Itoa(r.m.ID))
	data.CategoryID = &id
	data.CountryCode = r.m.CountryCode
	data.Locale = r.m.Locale

	// Load articles.
	result, err := dataloader.LoadArticles(ctx, data)
	if err != nil {
		return nil, err
	}
	return &ArticlesResolver{m: result}, nil
}
