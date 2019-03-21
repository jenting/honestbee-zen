package dataloader

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/graph-gophers/dataloader"
	"github.com/pkg/errors"

	"github.com/honestbee/Zen/errs"
	"github.com/honestbee/Zen/inout"
	"github.com/honestbee/Zen/zendesk"
)

// LoadSearchTitleArticles implements data loader.
func LoadSearchTitleArticles(ctx context.Context, params interface{}) ([]*zendesk.InstantSearchResult, error) {
	b, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	ldr, err := extract(ctx, searchTitleArticlesLoaderKey)
	if err != nil {
		return nil, err
	}

	data, err := ldr.Load(ctx, dataloader.StringKey(b))()
	if err != nil {
		return nil, err
	}

	instantSearch, ok := data.([]*zendesk.InstantSearchResult)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", instantSearch, data)
	}

	return instantSearch, nil
}

type searchTitleArticlesLoader struct {
	zend *zendesk.ZenDesk
}

func newSearchTitleArticlesLoader(zend *zendesk.ZenDesk) dataloader.BatchFunc {
	return searchTitleArticlesLoader{zend: zend}.loadBatch
}

func (l searchTitleArticlesLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var (
		n       = len(keys)
		results = make([]*dataloader.Result, n)
		wg      sync.WaitGroup
	)

	wg.Add(n)

	for i, key := range keys {
		go func(i int, key dataloader.Key) {
			defer wg.Done()

			data := inout.QuerySearchTitleArticlesIn{}
			if err := json.Unmarshal([]byte(key.String()), &data); err != nil {
				results[i] = &dataloader.Result{
					Error: errs.NewErr(
						errs.ServerInternalErrorCode,
						errors.Wrapf(err, "dataloader: [searchTitleArticlesLoader] json unmarshal failed"),
					)}
				return
			}

			zendeskInstantSearch, err := l.zend.InstantSearch(ctx, data.Query, data.CountryCode, data.Locale)
			if err != nil {
				results[i] = &dataloader.Result{
					Error: errs.NewErr(
						errs.ServerInternalErrorCode,
						errors.Wrapf(err, "dataloader: [searchTitleArticlesLoader] zend.InstantSearch failed"),
					)}
				return
			}

			searchResult := make([]*zendesk.InstantSearchResult, len(zendeskInstantSearch.Results))
			for i, result := range zendeskInstantSearch.Results {
				searchResult[i] = new(zendesk.InstantSearchResult)
				searchResult[i].Title = result.Title
				searchResult[i].CategoryTitle = result.CategoryTitle
				searchResult[i].URL = result.URL
			}

			results[i] = &dataloader.Result{Data: searchResult}
		}(i, key)
	}

	wg.Wait()

	return results
}
