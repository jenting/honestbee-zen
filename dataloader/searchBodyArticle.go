package dataloader

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"sync"

	"github.com/graph-gophers/dataloader"
	"github.com/pkg/errors"

	"github.com/honestbee/Zen/errs"
	"github.com/honestbee/Zen/inout"
	"github.com/honestbee/Zen/models"
	"github.com/honestbee/Zen/zendesk"
)

// LoadSearchBodyArticles implements data loader.
func LoadSearchBodyArticles(ctx context.Context, params interface{}) (*inout.GetSearchOut, error) {
	b, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	ldr, err := extract(ctx, searchBodyArticlesLoaderKey)
	if err != nil {
		return nil, err
	}

	data, err := ldr.Load(ctx, dataloader.StringKey(b))()
	if err != nil {
		return nil, err
	}

	searchArticles, ok := data.(*inout.GetSearchOut)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", searchArticles, data)
	}

	return searchArticles, nil
}

type searchBodyArticlesLoader struct {
	service models.Service
	zend    *zendesk.ZenDesk
}

func newSearchBodyArticlesLoader(service models.Service, zend *zendesk.ZenDesk) dataloader.BatchFunc {
	return searchBodyArticlesLoader{service: service, zend: zend}.loadBatch
}

func (l searchBodyArticlesLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var (
		n       = len(keys)
		results = make([]*dataloader.Result, n)
		wg      sync.WaitGroup
	)

	wg.Add(n)

	for i, key := range keys {
		go func(i int, key dataloader.Key) {
			defer wg.Done()

			data := inout.QuerySearchBodyArticlesIn{}
			if err := json.Unmarshal([]byte(key.String()), &data); err != nil {
				results[i] = &dataloader.Result{
					Error: errs.NewErr(
						errs.ServerInternalErrorCode,
						errors.Wrapf(err, "dataloader: [searchBodyArticlesLoader] json unmarshal failed"),
					)}
				return
			}

			categoryIDs, err := l.service.GetCategoriesID(ctx, data.CountryCode)
			if err != nil {
				results[i] = &dataloader.Result{
					Error: errs.NewErr(
						errs.ServerInternalErrorCode,
						errors.Wrapf(err, "dataloader: [searchBodyArticlesLoader] Service.GetCategories failed"),
					)}
				return
			}

			zendeskSearch, err := l.zend.Search(ctx, categoryIDs, data.Query, data.CountryCode, data.Locale,
				&zendesk.Pagination{
					PerPage:   int(data.PerPage),
					Page:      int(math.Round(float64(data.Page)/float64(data.PerPage))) + 1,
					SortOrder: data.SortOrder,
				})
			if err != nil {
				results[i] = &dataloader.Result{
					Error: errs.NewErr(
						errs.ServerInternalErrorCode,
						errors.Wrapf(err, "dataloader: [searchBodyArticlesLoader] zend.Search failed"),
					)}
				return
			}

			articles := make([]*models.SearchArticle, 0)
			for _, zendeskArticle := range zendeskSearch.Articles {
				category, err := l.service.GetCategoryByArticleID(ctx, zendeskArticle.ID, zendeskArticle.Locale)
				if err != nil {
					if err == models.ErrNotFound {
						// if article not found in locale db, ignore it.
						continue
					}
					results[i] = &dataloader.Result{
						Error: errs.NewErr(
							errs.RecordNotFoundErrorCode,
							errors.Wrapf(err, "dataloader: [searchBodyArticlesLoader] service.GetCategoryByArticleID failed"),
						)}
					return
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

			articlesOut := &inout.GetSearchOut{
				Articles: articles,
				BaseOut: &inout.BaseOut{
					Page:      zendeskSearch.Page,
					PerPage:   zendeskSearch.PerPage,
					PageCount: zendeskSearch.PageCount,
					Count:     zendeskSearch.Count,
				},
			}
			results[i] = &dataloader.Result{Data: articlesOut}
		}(i, key)
	}

	wg.Wait()

	return results
}
