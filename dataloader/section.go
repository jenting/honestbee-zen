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

// LoadSections implements data loader.
func LoadSections(ctx context.Context, params interface{}) (*inout.GetSectionsOut, error) {
	b, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	ldr, err := extract(ctx, sectionsLoaderKey)
	if err != nil {
		return nil, err
	}

	data, err := ldr.Load(ctx, dataloader.StringKey(b))()
	if err != nil {
		return nil, err
	}

	sections, ok := data.(*inout.GetSectionsOut)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", sections, data)
	}

	return sections, nil
}

type sectionsLoader struct {
	service  models.Service
	examiner *examiner.Examiner
}

func newSectionsLoader(service models.Service, examiner *examiner.Examiner) dataloader.BatchFunc {
	return sectionsLoader{service: service, examiner: examiner}.loadBatch
}

func (l sectionsLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var (
		n       = len(keys)
		results = make([]*dataloader.Result, n)
		wg      sync.WaitGroup
	)

	wg.Add(n)

	for i, key := range keys {
		go func(i int, key dataloader.Key) {
			defer wg.Done()

			data := inout.QuerySectionsIn{}
			if err := json.Unmarshal([]byte(key.String()), &data); err != nil {
				results[i] = &dataloader.Result{
					Error: errs.NewErr(
						errs.ServerInternalErrorCode,
						errors.Wrapf(err, "dataloader: [sectionsLoader] json unmarshal failed"),
					)}
				return
			}

			if data.CategoryID != nil {
				id64, err := strconv.ParseInt(string(*data.CategoryID), 10, 64)
				if err != nil {
					results[i] = &dataloader.Result{
						Error: errs.NewErr(
							errs.RecordNotFoundErrorCode,
							errors.Wrapf(err, "dataloader: [sectionsLoader] parse category id to int failed"),
						)}
					return
				}

				defer l.examiner.CheckSections(ctx, data.CountryCode, data.Locale)

				// Get key-value from cache.
				value, exist := l.service.SectionsCacheGet(ctx, key.String(), data.CountryCode, data.Locale)
				if exist {
					sectionsOut := &inout.GetSectionsOut{}
					if err := json.Unmarshal([]byte(value), &sectionsOut); err != nil {
						results[i] = &dataloader.Result{
							Error: errs.NewErr(
								errs.ServerInternalErrorCode,
								errors.Wrapf(err, "dataloader: [sectionsLoader] json unmarshal failed"),
							)}
						return
					}
					results[i] = &dataloader.Result{Data: sectionsOut}
				} else {
					sections, total, err := l.service.GetSectionsByCategoryID(ctx,
						&models.GetSectionsParams{
							Locale:      data.Locale,
							CountryCode: data.CountryCode,
							PerPage:     int(data.PerPage),
							Page:        int(data.Page),
							SortBy:      data.SortBy,
							SortOrder:   data.SortOrder,
							CategoryID:  int(id64),
						})
					if err != nil {
						results[i] = &dataloader.Result{
							Error: errs.NewErr(
								errs.ServerInternalErrorCode,
								errors.Wrapf(err, "dataloader: [sectionsLoader] service.GetSectionsByCategoryID failed"),
							)}
						return
					}

					sectionsOut := &inout.GetSectionsOut{
						Sections: sections,
						BaseOut: &inout.BaseOut{
							Page:      int(math.Round(float64(data.Page)/float64(data.PerPage))) + 1,
							PerPage:   int(data.PerPage),
							PageCount: int(math.Ceil(float64(total) / float64(data.PerPage))),
							Count:     total,
						},
					}
					results[i] = &dataloader.Result{Data: sectionsOut}

					// Set key-value to cache.
					if b, err := json.Marshal(sectionsOut); err == nil {
						l.service.SectionsCacheSet(ctx, key.String(), string(b), data.CountryCode, data.Locale)
					}
				}
			} else {
				defer l.examiner.CheckSections(ctx, data.CountryCode, data.Locale)

				// Get key-value from cache.
				value, exist := l.service.SectionsCacheGet(ctx, key.String(), data.CountryCode, data.Locale)
				if exist {
					sectionsOut := &inout.GetSectionsOut{}
					if err := json.Unmarshal([]byte(value), &sectionsOut); err != nil {
						results[i] = &dataloader.Result{
							Error: errs.NewErr(
								errs.ServerInternalErrorCode,
								errors.Wrapf(err, "dataloader: [sectionsLoader] json unmarshal failed"),
							)}
						return
					}
					results[i] = &dataloader.Result{Data: sectionsOut}
				} else {
					sections, total, err := l.service.GetSections(ctx,
						&models.GetSectionsParams{
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
								errors.Wrapf(err, "dataloader: [sectionsLoader] service.GetSections failed"),
							)}
						return
					}

					sectionsOut := &inout.GetSectionsOut{
						Sections: sections,
						BaseOut: &inout.BaseOut{
							Page:      int(math.Round(float64(data.Page)/float64(data.PerPage))) + 1,
							PerPage:   int(data.PerPage),
							PageCount: int(math.Ceil(float64(total) / float64(data.PerPage))),
							Count:     total,
						},
					}
					results[i] = &dataloader.Result{Data: sectionsOut}

					// Set key-value to cache.
					if b, err := json.Marshal(sectionsOut); err == nil {
						l.service.SectionsCacheSet(ctx, key.String(), string(b), data.CountryCode, data.Locale)
					}
				}
			}
		}(i, key)
	}

	wg.Wait()

	return results
}

// LoadSection implements data loader.
func LoadSection(ctx context.Context, params interface{}) (*models.Section, error) {
	b, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	ldr, err := extract(ctx, sectionLoaderKey)
	if err != nil {
		return nil, err
	}

	data, err := ldr.Load(ctx, dataloader.StringKey(b))()
	if err != nil {
		return nil, err
	}

	section, ok := data.(*models.Section)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", section, data)
	}

	return section, nil
}

type sectionLoader struct {
	service  models.Service
	examiner *examiner.Examiner
}

func newSectionLoader(service models.Service, examiner *examiner.Examiner) dataloader.BatchFunc {
	return sectionLoader{service: service, examiner: examiner}.loadBatch
}

func (l sectionLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var (
		n       = len(keys)
		results = make([]*dataloader.Result, n)
		wg      sync.WaitGroup
	)

	wg.Add(n)

	for i, key := range keys {
		go func(i int, key dataloader.Key) {
			defer wg.Done()

			data := inout.QuerySectionIn{}
			if err := json.Unmarshal([]byte(key.String()), &data); err != nil {
				results[i] = &dataloader.Result{
					Error: errs.NewErr(
						errs.ServerInternalErrorCode,
						errors.Wrapf(err, "dataloader: [sectionLoader] json unmarshal failed"),
					)}
				return
			}

			sectionID64, err := strconv.ParseInt(string(data.SectionID), 10, 64)
			if err != nil {
				results[i] = &dataloader.Result{
					Error: errs.NewErr(
						errs.RecordNotFoundErrorCode,
						errors.Wrapf(err, "dataloader: [sectionLoader] parse section id to int failed"),
					)}
				return
			}

			defer l.examiner.CheckSections(ctx, data.CountryCode, data.Locale)

			// Get key-value from cache.
			value, exist := l.service.SectionsCacheGet(ctx, key.String(), data.CountryCode, data.Locale)
			if exist {
				sectionOut := &models.Section{}
				if err := json.Unmarshal([]byte(value), &sectionOut); err != nil {
					results[i] = &dataloader.Result{
						Error: errs.NewErr(
							errs.ServerInternalErrorCode,
							errors.Wrapf(err, "dataloader: [sectionLoader] json unmarshal failed"),
						)}
					return
				}
				results[i] = &dataloader.Result{Data: sectionOut}
			} else {
				sectionOut, err := l.service.GetSectionBySectionID(ctx, int(sectionID64), data.Locale, data.CountryCode)
				if err != nil {
					switch err {
					case models.ErrNotFound:
						results[i] = &dataloader.Result{
							Error: errs.NewErr(
								errs.RecordNotFoundErrorCode,
								errors.Wrapf(err, "dataloader: [sectionLoader] service.GetSectionBySectionID not found"),
							)}
					default:
						results[i] = &dataloader.Result{
							Error: errs.NewErr(
								errs.ServerInternalErrorCode,
								errors.Wrapf(err, "dataloader: [sectionLoader] service.GetSectionBySectionID failed"),
							)}
					}
					return
				}
				results[i] = &dataloader.Result{Data: sectionOut}

				// Set key-value to cache.
				if b, err := json.Marshal(sectionOut); err == nil {
					l.service.SectionsCacheSet(ctx, key.String(), string(b), data.CountryCode, data.Locale)
				}
			}
		}(i, key)
	}

	wg.Wait()

	return results
}
