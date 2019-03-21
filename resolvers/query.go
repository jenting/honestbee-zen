package resolvers

import (
	"context"

	"github.com/honestbee/Zen/dataloader"
	"github.com/honestbee/Zen/inout"
)

// AllCategories creates a new allCategories resolver.
func (r *Resolver) AllCategories(ctx context.Context, data inout.QueryCategoriesIn) (*CategoriesResolver, error) {
	// Process input params.
	if err := data.ProcessInputParams(); err != nil {
		return nil, err
	}

	// Load categories.
	result, err := dataloader.LoadCategories(ctx, data)
	if err != nil {
		return nil, err
	}

	return &CategoriesResolver{m: result}, nil
}

// OneCategory creates a new oneCategory resolver.
func (r *Resolver) OneCategory(ctx context.Context, data inout.QueryCategoryIn) (*CategoryResolver, error) {
	// Process input params.
	if err := data.ProcessInputParams(); err != nil {
		return nil, err
	}

	// Load category.
	result, err := dataloader.LoadCategory(ctx, data)
	if err != nil {
		return nil, err
	}

	return &CategoryResolver{m: result}, nil
}

// AllSections creates a new allSections resolver.
func (r *Resolver) AllSections(ctx context.Context, data inout.QuerySectionsIn) (*SectionsResolver, error) {
	// Process input params.
	if err := data.ProcessInputParams(); err != nil {
		return nil, err
	}

	// Load sections.
	result, err := dataloader.LoadSections(ctx, data)
	if err != nil {
		return nil, err
	}
	return &SectionsResolver{m: result}, nil
}

// OneSection creates a new oneSection resolver.
func (r *Resolver) OneSection(ctx context.Context, data inout.QuerySectionIn) (*SectionResolver, error) {
	// Process input params.
	if err := data.ProcessInputParams(); err != nil {
		return nil, err
	}

	// Load section.
	result, err := dataloader.LoadSection(ctx, data)
	if err != nil {
		return nil, err
	}

	return &SectionResolver{m: result}, nil
}

// AllArticles creates a new allArticles resolver.
func (r *Resolver) AllArticles(ctx context.Context, data inout.QueryArticlesIn) (*ArticlesResolver, error) {
	// Process input params.
	if err := data.ProcessInputParams(); err != nil {
		return nil, err
	}

	// Load articles.
	result, err := dataloader.LoadArticles(ctx, data)
	if err != nil {
		return nil, err
	}
	return &ArticlesResolver{m: result}, nil
}

// TopArticles creates a new topArticles resolver.
func (r *Resolver) TopArticles(ctx context.Context, data inout.QueryTopArticlesIn) (*[]*ArticleResolver, error) {
	// Process input params.
	if err := data.ProcessInputParams(); err != nil {
		return nil, err
	}

	// Load top articles.
	ret := make([]*ArticleResolver, 0)
	results, err := dataloader.LoadTopArticles(ctx, data)
	if err != nil {
		return nil, err
	}

	for _, result := range results {
		ret = append(ret, &ArticleResolver{m: result})
	}
	return &ret, nil
}

// OneArticle creates a new oneArticle resolver.
func (r *Resolver) OneArticle(ctx context.Context, data inout.QueryArticleIn) (*ArticleResolver, error) {
	// Process input params.
	if err := data.ProcessInputParams(); err != nil {
		return nil, err
	}

	// Load article.
	result, err := dataloader.LoadArticle(ctx, data)
	if err != nil {
		return nil, err
	}

	return &ArticleResolver{m: result}, nil
}

// OneTicketForm creates a new oneTicketForm resolver.
func (r *Resolver) OneTicketForm(ctx context.Context, data inout.QueryTicketFormIn) (*TicketFormResolver, error) {
	result, err := dataloader.LoadTicketForm(ctx, data)
	if err != nil {
		return nil, err
	}
	return &TicketFormResolver{m: result}, nil
}

// SearchTitleArticles creates a new searchTitleArticles resolver.
func (r *Resolver) SearchTitleArticles(ctx context.Context, data inout.QuerySearchTitleArticlesIn) (*[]*SearchTitleArticleResolver, error) {
	// Process input params.
	if err := data.ProcessInputParams(); err != nil {
		return nil, err
	}

	// Load search title articles.
	results, err := dataloader.LoadSearchTitleArticles(ctx, data)
	if err != nil {
		return nil, err
	}

	ret := make([]*SearchTitleArticleResolver, 0)
	for _, result := range results {
		ret = append(ret, &SearchTitleArticleResolver{m: result})
	}

	return &ret, nil
}

// SearchBodyArticles creates a new searchBodyArticles resolver.
func (r *Resolver) SearchBodyArticles(ctx context.Context, data inout.QuerySearchBodyArticlesIn) (*SearchBodyArticlesResolver, error) {
	// Process input params.
	if err := data.ProcessInputParams(); err != nil {
		return nil, err
	}

	// Load search body articles.
	result, err := dataloader.LoadSearchBodyArticles(ctx, data)
	if err != nil {
		return nil, err
	}
	return &SearchBodyArticlesResolver{m: result}, nil
}

// Status creates a new status resolver.
func (r *Resolver) Status(ctx context.Context) (*StatusResolver, error) {
	return &StatusResolver{}, nil
}
