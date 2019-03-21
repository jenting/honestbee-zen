package dataloader

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"sync"

	"github.com/graph-gophers/dataloader"
	"github.com/pkg/errors"

	"github.com/honestbee/Zen/errs"
	"github.com/honestbee/Zen/examiner"
	"github.com/honestbee/Zen/inout"
	"github.com/honestbee/Zen/models"
)

// LoadArticles implements data loader.
func LoadArticles(ctx context.Context, params interface{}) (*inout.GetArticlesOut, error) {
	b, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	ldr, err := extract(ctx, articlesLoaderKey)
	if err != nil {
		return nil, err
	}

	data, err := ldr.Load(ctx, dataloader.StringKey(b))()
	if err != nil {
		return nil, err
	}

	articles, ok := data.(*inout.GetArticlesOut)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", articles, data)
	}

	return articles, nil
}

type articlesLoader struct {
	service  models.Service
	examiner *examiner.Examiner
}

func newArticlesLoader(service models.Service, examiner *examiner.Examiner) dataloader.BatchFunc {
	return articlesLoader{service: service, examiner: examiner}.loadBatch
}

func (l articlesLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var (
		n       = len(keys)
		results = make([]*dataloader.Result, n)
		wg      sync.WaitGroup
	)

	wg.Add(n)

	for i, key := range keys {
		go func(i int, key dataloader.Key) {
			defer wg.Done()

			data := inout.QueryArticlesIn{}
			if err := json.Unmarshal([]byte(key.String()), &data); err != nil {
				results[i] = &dataloader.Result{
					Error: errs.NewErr(
						errs.ServerInternalErrorCode,
						errors.Wrapf(err, "dataloader: [articlesLoader] json unmarshal failed"),
					)}
				return
			}

			if data.CategoryID != nil {
				id64, err := strconv.ParseInt(string(*data.CategoryID), 10, 64)
				if err != nil {
					results[i] = &dataloader.Result{
						Error: errs.NewErr(
							errs.RecordNotFoundErrorCode,
							errors.Wrapf(err, "dataloader: [articlesLoader] parse category id to int failed"),
						)}
					return
				}

				defer l.examiner.CheckArticles(ctx, data.CountryCode, data.Locale)

				// Get key-value from cache.
				value, exist := l.service.ArticlesCacheGet(ctx, key.String(), data.CountryCode, data.Locale)
				if exist {
					articlesOut := &inout.GetArticlesOut{}
					if err := json.Unmarshal([]byte(value), &articlesOut); err != nil {
						results[i] = &dataloader.Result{
							Error: errs.NewErr(
								errs.ServerInternalErrorCode,
								errors.Wrapf(err, "dataloader: [articlesLoader] json unmarshal failed"),
							)}
						return
					}
					results[i] = &dataloader.Result{Data: articlesOut}
				} else {
					articles, total, err := l.service.GetArticlesByCategoryID(ctx,
						&models.GetArticlesParams{
							Locale:      data.Locale,
							CountryCode: data.CountryCode,
							PerPage:     int(data.PerPage),
							Page:        int(data.Page),
							SortBy:      data.SortBy,
							SortOrder:   data.SortOrder,
							CategoryID:  int(id64),
						}, []string{})
					if err != nil {
						results[i] = &dataloader.Result{
							Error: errs.NewErr(
								errs.ServerInternalErrorCode,
								errors.Wrapf(err, "dataloader: [articlesLoader] service.GetArticlesByCategoryID failed"),
							)}
						return
					}

					articlesOut := &inout.GetArticlesOut{
						Articles: articles,
						BaseOut: &inout.BaseOut{
							Page:      int(math.Round(float64(data.Page)/float64(data.PerPage))) + 1,
							PerPage:   int(data.PerPage),
							PageCount: int(math.Ceil(float64(total) / float64(data.PerPage))),
							Count:     total,
						},
					}
					results[i] = &dataloader.Result{Data: articlesOut}

					// Set key-value to cache.
					if b, err := json.Marshal(articlesOut); err == nil {
						l.service.ArticlesCacheSet(ctx, key.String(), string(b), data.CountryCode, data.Locale)
					}
				}
			} else if data.SectionID != nil {
				id64, err := strconv.ParseInt(string(*data.SectionID), 10, 64)
				if err != nil {
					results[i] = &dataloader.Result{
						Error: errs.NewErr(
							errs.RecordNotFoundErrorCode,
							errors.Wrapf(err, "dataloader: [articlesLoader] parse section id to int failed"),
						)}
					return
				}

				defer l.examiner.CheckArticles(ctx, data.CountryCode, data.Locale)

				// Get key-value from cache.
				value, exist := l.service.ArticlesCacheGet(ctx, key.String(), data.CountryCode, data.Locale)
				if exist {
					articlesOut := &inout.GetArticlesOut{}
					if err := json.Unmarshal([]byte(value), &articlesOut); err != nil {
						results[i] = &dataloader.Result{
							Error: errs.NewErr(
								errs.ServerInternalErrorCode,
								errors.Wrapf(err, "dataloader: [articlesLoader] json unmarshal failed"),
							)}
						return
					}
					results[i] = &dataloader.Result{Data: articlesOut}
				} else {
					articles, total, err := l.service.GetArticlesBySectionID(ctx,
						&models.GetArticlesParams{
							Locale:      data.Locale,
							CountryCode: data.CountryCode,
							PerPage:     int(data.PerPage),
							Page:        int(data.Page),
							SortBy:      data.SortBy,
							SortOrder:   data.SortOrder,
							SectionID:   int(id64),
						})
					if err != nil {
						results[i] = &dataloader.Result{
							Error: errs.NewErr(
								errs.ServerInternalErrorCode,
								errors.Wrapf(err, "dataloader: [articlesLoader] service.GetArticlesBySectionID failed"),
							)}
						return
					}

					articlesOut := &inout.GetArticlesOut{
						Articles: articles,
						BaseOut: &inout.BaseOut{
							Page:      int(math.Round(float64(data.Page)/float64(data.PerPage))) + 1,
							PerPage:   int(data.PerPage),
							PageCount: int(math.Ceil(float64(total) / float64(data.PerPage))),
							Count:     total,
						},
					}
					results[i] = &dataloader.Result{Data: articlesOut}

					// Set key-value to cache.
					if b, err := json.Marshal(articlesOut); err == nil {
						l.service.ArticlesCacheSet(ctx, key.String(), string(b), data.CountryCode, data.Locale)
					}
				}
			} else {
				defer l.examiner.CheckArticles(ctx, data.CountryCode, data.Locale)

				// Get key-value from cache.
				value, exist := l.service.ArticlesCacheGet(ctx, key.String(), data.CountryCode, data.Locale)
				if exist {
					articlesOut := &inout.GetArticlesOut{}
					if err := json.Unmarshal([]byte(value), &articlesOut); err != nil {
						results[i] = &dataloader.Result{
							Error: errs.NewErr(
								errs.ServerInternalErrorCode,
								errors.Wrapf(err, "dataloader: [articlesLoader] json unmarshal failed"),
							)}
						return
					}
					results[i] = &dataloader.Result{Data: articlesOut}
				} else {
					articles, total, err := l.service.GetArticles(ctx,
						&models.GetArticlesParams{
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
								errors.Wrapf(err, "dataloader: [articlesLoader] service.GetArticles failed"),
							)}
						return
					}

					articlesOut := &inout.GetArticlesOut{
						Articles: articles,
						BaseOut: &inout.BaseOut{
							Page:      int(math.Round(float64(data.Page)/float64(data.PerPage))) + 1,
							PerPage:   int(data.PerPage),
							PageCount: int(math.Ceil(float64(total) / float64(data.PerPage))),
							Count:     total,
						},
					}
					results[i] = &dataloader.Result{Data: articlesOut}

					// Set key-value to cache.
					if b, err := json.Marshal(articlesOut); err == nil {
						l.service.ArticlesCacheSet(ctx, key.String(), string(b), data.CountryCode, data.Locale)
					}
				}
			}
		}(i, key)
	}

	wg.Wait()

	return results
}

// LoadTopArticles implements data loader.
func LoadTopArticles(ctx context.Context, params interface{}) ([]*models.Article, error) {
	b, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	ldr, err := extract(ctx, topArticlesLoaderKey)
	if err != nil {
		return nil, err
	}

	data, err := ldr.Load(ctx, dataloader.StringKey(b))()
	if err != nil {
		return nil, err
	}

	articles, ok := data.([]*models.Article)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", articles, data)
	}

	return articles, nil
}

type topArticlesLoader struct {
	service models.Service
}

func newTopArticlesLoader(service models.Service) dataloader.BatchFunc {
	return topArticlesLoader{service: service}.loadBatch
}

func (l topArticlesLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var (
		n       = len(keys)
		results = make([]*dataloader.Result, n)
		wg      sync.WaitGroup
	)

	wg.Add(n)

	for i, key := range keys {
		go func(i int, key dataloader.Key) {
			defer wg.Done()

			data := inout.QueryTopArticlesIn{}
			if err := json.Unmarshal([]byte(key.String()), &data); err != nil {
				results[i] = &dataloader.Result{
					Error: errs.NewErr(
						errs.ServerInternalErrorCode,
						errors.Wrapf(err, "dataloader: [topArticlesLoader] json unmarshal failed"),
					)}
				return
			}

			// Check input parameter top_n is a positive number.
			if data.TopN <= 0 {
				results[i] = &dataloader.Result{
					Error: errs.NewErr(
						errs.InvalidAttributeErrorCode,
						errors.New("dataloader: [topArticlesLoader] invalid top_n input failed"),
					)}
				return
			}

			// Get key-value from cache.
			value, exist := l.service.ArticlesCacheGet(ctx, key.String(), data.CountryCode, data.Locale)
			if exist {
				articlesOut := make([]*models.Article, 0)
				if err := json.Unmarshal([]byte(value), &articlesOut); err != nil {
					results[i] = &dataloader.Result{
						Error: errs.NewErr(
							errs.ServerInternalErrorCode,
							errors.Wrapf(err, "dataloader: [topArticlesLoader] json unmarshal failed"),
						)}
					return
				}
				results[i] = &dataloader.Result{Data: articlesOut}
			} else {
				articlesOut, err := l.service.GetTopNArticles(ctx, uint64(data.TopN), data.Locale, data.CountryCode)
				if err != nil {
					results[i] = &dataloader.Result{
						Error: errs.NewErr(
							errs.ServerInternalErrorCode,
							errors.Wrapf(err, "dataloader: [topArticlesLoader] service.GetTopNArticles failed"),
						)}
					return
				}
				results[i] = &dataloader.Result{Data: articlesOut}

				// Set key-value to cache.
				if b, err := json.Marshal(articlesOut); err == nil {
					l.service.ArticlesCacheSet(ctx, key.String(), string(b), data.CountryCode, data.Locale)
				}
			}
		}(i, key)
	}

	wg.Wait()

	return results
}

// LoadArticle implements data loader.
func LoadArticle(ctx context.Context, params interface{}) (*models.Article, error) {
	b, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	ldr, err := extract(ctx, articleLoaderKey)
	if err != nil {
		return nil, err
	}

	data, err := ldr.Load(ctx, dataloader.StringKey(b))()
	if err != nil {
		return nil, err
	}

	article, ok := data.(*models.Article)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", article, data)
	}

	return article, nil
}

type articleLoader struct {
	service  models.Service
	examiner *examiner.Examiner
}

func newArticleLoader(service models.Service, examiner *examiner.Examiner) dataloader.BatchFunc {
	return articleLoader{service: service, examiner: examiner}.loadBatch
}

func (l articleLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var (
		n       = len(keys)
		results = make([]*dataloader.Result, n)
		wg      sync.WaitGroup
	)

	wg.Add(n)

	for i, key := range keys {
		go func(i int, key dataloader.Key) {
			defer wg.Done()

			data := inout.QueryArticleIn{}
			if err := json.Unmarshal([]byte(key.String()), &data); err != nil {
				results[i] = &dataloader.Result{
					Error: errs.NewErr(
						errs.ServerInternalErrorCode,
						errors.Wrapf(err, "dataloader: [articleLoader] json unmarshal failed"),
					)}
				return
			}

			articleID64, err := strconv.ParseInt(string(data.ArticleID), 10, 64)
			if err != nil {
				results[i] = &dataloader.Result{
					Error: errs.NewErr(
						errs.RecordNotFoundErrorCode,
						errors.Wrapf(err, "dataloader: [articleLoader] parse article id to int failed"),
					)}
				return
			}

			defer l.examiner.CheckArticles(ctx, data.CountryCode, data.Locale)
			defer l.service.PlusOneArticleClickCounter(ctx, int(articleID64), data.Locale, data.CountryCode)

			// Get key-value from cache.
			value, exist := l.service.ArticlesCacheGet(ctx, key.String(), data.CountryCode, data.Locale)
			if exist {
				articlesOut := &models.Article{}
				if err := json.Unmarshal([]byte(value), &articlesOut); err != nil {
					results[i] = &dataloader.Result{
						Error: errs.NewErr(
							errs.ServerInternalErrorCode,
							errors.Wrapf(err, "dataloader: [articleLoader] json unmarshal failed"),
						)}
					return
				}
				results[i] = &dataloader.Result{Data: articlesOut}
			} else {
				articleOut, err := l.service.GetArticleByArticleID(ctx, int(articleID64), data.Locale, data.CountryCode)
				if err != nil {
					switch err {
					case models.ErrNotFound:
						results[i] = &dataloader.Result{
							Error: errs.NewErr(
								errs.RecordNotFoundErrorCode,
								errors.Wrapf(err, "dataloader: [articleLoader] service.GetArticleByArticleID not found"),
							)}
					default:
						results[i] = &dataloader.Result{
							Error: errs.NewErr(
								errs.ServerInternalErrorCode,
								errors.Wrapf(err, "dataloader: [articleLoader] service.GetArticleByArticleID failed"),
							)}
					}
					return
				}
				results[i] = &dataloader.Result{Data: articleOut}

				// Set key-value to cache.
				if b, err := json.Marshal(articleOut); err == nil {
					l.service.ArticlesCacheSet(ctx, key.String(), string(b), data.CountryCode, data.Locale)
				}
			}
		}(i, key)
	}

	wg.Wait()

	return results
}
