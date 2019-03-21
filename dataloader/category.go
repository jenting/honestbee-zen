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
	"github.com/honestbee/Zen/examiner"
	"github.com/honestbee/Zen/inout"
	"github.com/honestbee/Zen/models"
)

// LoadCategories implements data loader.
func LoadCategories(ctx context.Context, params interface{}) (*inout.GetCategoriesOut, error) {
	b, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	ldr, err := extract(ctx, categoriesLoaderKey)
	if err != nil {
		return nil, err
	}

	data, err := ldr.Load(ctx, dataloader.StringKey(b))()
	if err != nil {
		return nil, err
	}

	categories, ok := data.(*inout.GetCategoriesOut)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", categories, data)
	}

	return categories, nil
}

type categoriesLoader struct {
	service  models.Service
	examiner *examiner.Examiner
}

func newCategoriesLoader(service models.Service, examiner *examiner.Examiner) dataloader.BatchFunc {
	return categoriesLoader{service: service, examiner: examiner}.loadBatch
}

func (l categoriesLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var (
		n       = len(keys)
		results = make([]*dataloader.Result, n)
		wg      sync.WaitGroup
	)

	wg.Add(n)

	for i, key := range keys {
		go func(i int, key dataloader.Key) {
			defer wg.Done()

			data := inout.QueryCategoriesIn{}
			if err := json.Unmarshal([]byte(key.String()), &data); err != nil {
				results[i] = &dataloader.Result{
					Error: errs.NewErr(
						errs.ServerInternalErrorCode,
						errors.Wrapf(err, "dataloader: [categoriesLoader] json unmarshal failed"),
					)}
				return
			}

			defer l.examiner.CheckCategories(ctx, data.CountryCode, data.Locale)

			// Get key-value from cache.
			value, exist := l.service.CategoriesCacheGet(ctx, key.String(), data.CountryCode, data.Locale)
			if exist {
				categoriesOut := &inout.GetCategoriesOut{}
				if err := json.Unmarshal([]byte(value), &categoriesOut); err != nil {
					results[i] = &dataloader.Result{
						Error: errs.NewErr(
							errs.ServerInternalErrorCode,
							errors.Wrapf(err, "dataloader: [categoriesLoader] json unmarshal failed"),
						)}
					return
				}
				results[i] = &dataloader.Result{Data: categoriesOut}
			} else {
				categories, total, err := l.service.GetCategories(ctx,
					&models.GetCategoriesParams{
						Locale:      data.Locale,
						CountryCode: data.CountryCode,
						PerPage:     int(data.PerPage),
						Page:        int(data.Page),
						SortBy:      data.SortBy,
						SortOrder:   data.SortOrder,
					})
				if err != nil {
					results[i] = &dataloader.Result{
						Error: errs.NewErr(
							errs.ServerInternalErrorCode,
							errors.Wrapf(err, "dataloader: [categoriesLoader] service.GetCategories failed"),
						)}
					return
				}

				categoriesOut := &inout.GetCategoriesOut{
					Categories: categories,
					BaseOut: &inout.BaseOut{
						Page:      int(math.Round(float64(data.Page)/float64(data.PerPage))) + 1,
						PerPage:   int(data.PerPage),
						PageCount: int(math.Ceil(float64(total) / float64(data.PerPage))),
						Count:     total,
					},
				}
				results[i] = &dataloader.Result{Data: categoriesOut}

				// Set key-value to cache.
				if b, err := json.Marshal(categoriesOut); err == nil {
					l.service.CategoriesCacheSet(ctx, key.String(), string(b), data.CountryCode, data.Locale)
				}
			}
		}(i, key)
	}

	wg.Wait()

	return results
}

// LoadCategory implements data loader.
func LoadCategory(ctx context.Context, params interface{}) (*models.Category, error) {
	b, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	ldr, err := extract(ctx, categoryLoaderKey)
	if err != nil {
		return nil, err
	}

	data, err := ldr.Load(ctx, dataloader.StringKey(b))()
	if err != nil {
		return nil, err
	}

	category, ok := data.(*models.Category)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", category, data)
	}

	return category, nil
}

type categoryLoader struct {
	service  models.Service
	examiner *examiner.Examiner
}

func newCategoryLoader(service models.Service, examiner *examiner.Examiner) dataloader.BatchFunc {
	return categoryLoader{service: service, examiner: examiner}.loadBatch
}

func (l categoryLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var (
		n       = len(keys)
		results = make([]*dataloader.Result, n)
		wg      sync.WaitGroup
	)

	wg.Add(n)

	for i, key := range keys {
		go func(i int, key dataloader.Key) {
			defer wg.Done()

			data := inout.QueryCategoryIn{}
			if err := json.Unmarshal([]byte(key.String()), &data); err != nil {
				results[i] = &dataloader.Result{
					Error: errs.NewErr(
						errs.ServerInternalErrorCode,
						errors.Wrapf(err, "dataloader: [categoryLoader] json unmarshal failed"),
					)}
				return
			}

			defer l.examiner.CheckCategories(ctx, data.CountryCode, data.Locale)

			// Get key-value from cache.
			value, exist := l.service.CategoriesCacheGet(ctx, key.String(), data.CountryCode, data.Locale)
			if exist {
				categoryOut := &models.Category{}
				if err := json.Unmarshal([]byte(value), &categoryOut); err != nil {
					results[i] = &dataloader.Result{
						Error: errs.NewErr(
							errs.ServerInternalErrorCode,
							errors.Wrapf(err, "dataloader: [categoryLoader] json unmarshal failed"),
						)}
					return
				}
				results[i] = &dataloader.Result{Data: categoryOut}
			} else {
				categoryOut, err := l.service.GetCategoryByCategoryIDOrKeyName(ctx, string(data.CategoryIDOrKeyName), data.Locale, data.CountryCode)
				if err != nil {
					switch err {
					case models.ErrNotFound:
						results[i] = &dataloader.Result{
							Error: errs.NewErr(
								errs.RecordNotFoundErrorCode,
								errors.Wrapf(err, "dataloader: [categoryLoader] service.GetCategoryByCategoryIDOrKeyName not found"),
							)}
					default:
						results[i] = &dataloader.Result{
							Error: errs.NewErr(
								errs.ServerInternalErrorCode,
								errors.Wrapf(err, "dataloader: [categoryLoader] service.GetCategoryByCategoryIDOrKeyName failed"),
							)}
					}
					return
				}
				results[i] = &dataloader.Result{Data: categoryOut}

				// Set key-value to cache.
				if b, err := json.Marshal(categoryOut); err == nil {
					l.service.CategoriesCacheSet(ctx, key.String(), string(b), data.CountryCode, data.Locale)
				}
			}
		}(i, key)
	}

	wg.Wait()

	return results
}
