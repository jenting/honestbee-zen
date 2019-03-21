package resolvers

import (
	"context"

	"github.com/honestbee/Zen/zendesk"
)

// SearchTitleArticleResolver defines resolver models.
type SearchTitleArticleResolver struct {
	m *zendesk.InstantSearchResult
}

// Title is the InstantSearch's field id.
func (r *SearchTitleArticleResolver) Title(ctx context.Context) string {
	return r.m.Title
}

// CategoryTitle is the InstantSearch's field author_id.
func (r *SearchTitleArticleResolver) CategoryTitle(ctx context.Context) string {
	return r.m.CategoryTitle
}

// URL is the InstantSearch's field comments_disable.
func (r *SearchTitleArticleResolver) URL(ctx context.Context) string {
	return r.m.URL
}
