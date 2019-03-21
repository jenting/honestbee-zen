package handlers

import (
	"context"
	"math"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"

	"github.com/honestbee/Zen/errs"
	"github.com/honestbee/Zen/inout"
	"github.com/honestbee/Zen/models"
	"github.com/honestbee/Zen/zendesk"
)

// GetInstantSearchDecompressor combines params from URL or FORM
// and returns params in a structure that GetSearchHandler needs.
func GetInstantSearchDecompressor(ps httprouter.Params, r *http.Request) (interface{}, error) {
	baseParams, err := inout.FetchBaseParams(r)
	if err != nil {
		return nil, errs.NewErr(
			errs.InvalidAttributeErrorCode,
			errors.Wrapf(err, "handlers: [GetInstantSearchDecompressor] inout.FetchBaseParams failed"),
		)
	}

	queryStr := r.FormValue("query")
	if queryStr == "" {
		return nil, errs.NewErr(
			errs.InvalidAttributeErrorCode,
			errors.Wrapf(err, "handlers: [GetInstantSearchDecompressor] parse query value failed"),
		)
	}

	return &inout.GetInstantSearchIn{
		Query:       queryStr,
		Locale:      baseParams.Locale,
		CountryCode: baseParams.CountryCode,
	}, nil

}

// GetInstantSearchHandler handles instant search request.
func GetInstantSearchHandler(ctx context.Context, e *Env, in interface{}) (interface{}, error) {
	data, ok := in.(*inout.GetInstantSearchIn)
	if !ok {
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Errorf("handlers: [GetInstantSearchHandler] cast %v into *GetInstantSearchIn failed", in),
		)
	}

	zendeskInstantSearch, err := e.ZenDesk.InstantSearch(ctx, data.Query, data.CountryCode, data.Locale)
	if err != nil {
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Wrapf(err, "handlers: [GetInstantSearchHandler] ZenDesk.InstantSearch failed"),
		)
	}

	searchResult := make([]*inout.InstantSearchResult, len(zendeskInstantSearch.Results))

	for i, result := range zendeskInstantSearch.Results {
		searchResult[i] = new(inout.InstantSearchResult)
		searchResult[i].Title = result.Title
		searchResult[i].CategoryTitle = result.CategoryTitle
		searchResult[i].URL = result.URL
	}

	return &inout.GetInstantSearchOut{
		Results: searchResult,
	}, nil
}

// GetSearchDecompressor combines params from URL or FORM
// and returns params in a structure that GetSearchHandler needs.
func GetSearchDecompressor(ps httprouter.Params, r *http.Request) (interface{}, error) {
	baseParams, err := inout.FetchBaseParams(r)
	if err != nil {
		return nil, errs.NewErr(
			errs.InvalidAttributeErrorCode,
			errors.Wrapf(err, "handlers: [GetSearchDecompressor] inout.FetchBaseParams failed"),
		)
	}

	queryStr := r.FormValue("query")
	if queryStr == "" {
		return nil, errs.NewErr(
			errs.InvalidAttributeErrorCode,
			errors.Wrapf(err, "handlers: [GetSearchDecompressor] parse query value failed"),
		)
	}

	return &inout.GetSearchIn{
		Query:  queryStr,
		BaseIn: baseParams,
	}, nil
}

// GetSearchHandler handles search request.
func GetSearchHandler(ctx context.Context, e *Env, in interface{}) (interface{}, error) {
	data, ok := in.(*inout.GetSearchIn)
	if !ok {
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Errorf("handlers: [GetSearchHandler] cast %v into *GetSearchIn failed", in),
		)
	}

	categoryIDs, err := e.Service.GetCategoriesID(ctx, data.CountryCode)
	if err != nil {
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Wrapf(err, "handlers: [GetCategoriesHandler] Service.GetCategories failed"),
		)
	}

	zendeskSearch, err := e.ZenDesk.Search(ctx, categoryIDs, data.Query, data.CountryCode, data.Locale,
		&zendesk.Pagination{
			PerPage:   data.PerPage,
			Page:      int(math.Round(float64(data.Page)/float64(data.PerPage))) + 1,
			SortOrder: data.SortOrder,
		},
	)

	if err != nil {
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Wrapf(err, "handlers: [GetSearchHandler] ZenDesk.Search failed"),
		)
	}
	articles := make([]*models.SearchArticle, 0)
	for _, zendeskArticle := range zendeskSearch.Articles {
		category, err := e.Service.GetCategoryByArticleID(ctx, zendeskArticle.ID, zendeskArticle.Locale)
		if err != nil {
			if err == models.ErrNotFound {
				// if article not found in locale db, ignore it.
				continue
			}
			return nil, errs.NewErr(
				errs.RecordNotFoundErrorCode,
				errors.Wrapf(err, "handlers: [GetSearchHandler] Service.GetCategoryByArticleID failed"),
			)
		}

		articles = append(articles, &models.SearchArticle{
			Article: &models.Article{
				SectionID:       zendeskArticle.Article.SectionID,
				ID:              zendeskArticle.ID,
				AuthorID:        zendeskArticle.Article.AuthorID,
				CommentsDisable: zendeskArticle.CommentsDisable,
				Draft:           zendeskArticle.Draft,
				Promoted:        zendeskArticle.Promoted,
				Position:        zendeskArticle.Position,
				VoteSum:         zendeskArticle.VoteSum,
				VoteCount:       zendeskArticle.VoteCount,
				CreatedAt:       zendeskArticle.CreatedAt,
				UpdatedAt:       zendeskArticle.UpdatedAt,
				SourceLocale:    zendeskArticle.SourceLocale,
				Outdated:        zendeskArticle.Outdated,
				OutdatedLocales: zendeskArticle.OutdatedLocales,
				EditedAt:        zendeskArticle.EditedAt,
				LabelNames:      zendeskArticle.LabelNames,
				CountryCode:     data.CountryCode,
				URL:             zendeskArticle.URL,
				HTMLURL:         zendeskArticle.HTMLURL,
				Name:            zendeskArticle.Name,
				Title:           zendeskArticle.Title,
				Body:            zendeskArticle.Body,
				Locale:          zendeskArticle.Locale,
			},
			Snippet:      zendeskArticle.Snippet,
			CategoryID:   category.ID,
			CategoryName: category.Name,
		})
	}

	return &inout.GetSearchOut{
		Articles: articles,
		BaseOut: &inout.BaseOut{
			Page:      zendeskSearch.Page,
			PerPage:   zendeskSearch.PerPage,
			PageCount: zendeskSearch.PageCount,
			Count:     zendeskSearch.Count,
		},
	}, nil
}
