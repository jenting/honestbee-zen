package dataloader

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"

	"github.com/graph-gophers/dataloader"
	"github.com/pkg/errors"

	"github.com/honestbee/Zen/errs"
	"github.com/honestbee/Zen/inout"
	"github.com/honestbee/Zen/models"
)

// LoadTicketFields implements data loader.
func LoadTicketFields(ctx context.Context, params interface{}) ([]*models.TicketField, error) {
	b, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	ldr, err := extract(ctx, ticketFieldsLoaderKey)
	if err != nil {
		return nil, err
	}

	data, err := ldr.Load(ctx, dataloader.StringKey(b))()
	if err != nil {
		return nil, err
	}

	ticketFields, ok := data.([]*models.TicketField)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", ticketFields, data)
	}

	return ticketFields, nil
}

type ticketFieldsLoader struct {
	service models.Service
}

func newTicketFieldsLoader(service models.Service) dataloader.BatchFunc {
	return ticketFieldsLoader{service: service}.loadBatch
}

func (l ticketFieldsLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var (
		n       = len(keys)
		results = make([]*dataloader.Result, n)
		wg      sync.WaitGroup
	)

	wg.Add(n)

	for i, key := range keys {
		go func(i int, key dataloader.Key) {
			defer wg.Done()

			data := inout.QueryTicketFieldsIn{}
			if err := json.Unmarshal([]byte(key.String()), &data); err != nil {
				results[i] = &dataloader.Result{
					Error: errs.NewErr(
						errs.ServerInternalErrorCode,
						errors.Wrapf(err, "dataloader: [ticketFieldsLoader] json unmarshal failed"),
					)}
				return
			}

			if data.FormID != nil {
				formID64, err := strconv.ParseInt(string(*data.FormID), 10, 64)
				if err != nil {
					results[i] = &dataloader.Result{
						Error: errs.NewErr(
							errs.RecordNotFoundErrorCode,
							errors.Wrapf(err, "dataloader: [ticketFieldsLoader] parse form id to int failed"),
						)}
					return
				}

				// Get key-value from cache.
				value, exist := l.service.TicketFieldCacheGet(ctx, key.String())
				if exist {
					ticketFiledsOut := make([]*models.TicketField, 0)
					if err := json.Unmarshal([]byte(value), &ticketFiledsOut); err != nil {
						results[i] = &dataloader.Result{
							Error: errs.NewErr(
								errs.ServerInternalErrorCode,
								errors.Wrapf(err, "dataloader: [ticketFieldsLoader] json unmarshal failed"),
							)}
						return
					}
					results[i] = &dataloader.Result{Data: ticketFiledsOut}
				} else {
					ticketFiledsOut, err := l.service.GetTicketFieldByFormID(ctx, int(formID64), data.Locale)
					if err != nil {
						switch err {
						case models.ErrNotFound:
							results[i] = &dataloader.Result{
								Error: errs.NewErr(
									errs.RecordNotFoundErrorCode,
									errors.Wrapf(err, "dataloader: [ticketFieldsLoader] service.GetTicketFieldByFormID not found"),
								)}
						default:
							results[i] = &dataloader.Result{
								Error: errs.NewErr(
									errs.ServerInternalErrorCode,
									errors.Wrapf(err, "dataloader: [ticketFieldsLoader] service.GetTicketFieldByFormID failed"),
								)}
						}
						return
					}
					results[i] = &dataloader.Result{Data: ticketFiledsOut}

					// Set key-value to cache.
					if b, err := json.Marshal(ticketFiledsOut); err == nil {
						l.service.TicketFieldCacheSet(ctx, key.String(), string(b))
					}
				}
			}
		}(i, key)
	}

	wg.Wait()

	return results
}

// LoadTicketFieldCustomFieldOptions implements data loader.
func LoadTicketFieldCustomFieldOptions(ctx context.Context, params interface{}) ([]*models.CustomFieldOption, error) {
	b, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	ldr, err := extract(ctx, ticketFieldCustomFieldOption)
	if err != nil {
		return nil, err
	}

	data, err := ldr.Load(ctx, dataloader.StringKey(b))()
	if err != nil {
		return nil, err
	}

	customFieldOption, ok := data.([]*models.CustomFieldOption)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", customFieldOption, data)
	}

	return customFieldOption, nil
}

type ticketFieldCustomFieldOptions struct {
	service models.Service
}

func newTicketFieldCustomFieldOptionsLoader(service models.Service) dataloader.BatchFunc {
	return ticketFieldCustomFieldOptions{service: service}.loadBatch
}

func (l ticketFieldCustomFieldOptions) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var (
		n       = len(keys)
		results = make([]*dataloader.Result, n)
		wg      sync.WaitGroup
	)

	wg.Add(n)

	for i, key := range keys {
		go func(i int, key dataloader.Key) {
			defer wg.Done()

			data := inout.QueryCustomFieldOptionsIn{}
			if err := json.Unmarshal([]byte(key.String()), &data); err != nil {
				results[i] = &dataloader.Result{
					Error: errs.NewErr(
						errs.ServerInternalErrorCode,
						errors.Wrapf(err, "dataloader: [ticketFieldCustomFieldOptions] json unmarshal failed"),
					)}
				return
			}

			fieldID64, err := strconv.ParseInt(string(data.FieldID), 10, 64)
			if err != nil {
				results[i] = &dataloader.Result{
					Error: errs.NewErr(
						errs.RecordNotFoundErrorCode,
						errors.Wrapf(err, "dataloader: [ticketFieldCustomFieldOptions] parse field id to int failed"),
					)}
				return
			}

			// Get key-value from cache.
			value, exist := l.service.TicketFieldCustomFieldOptionCacheGet(ctx, key.String())
			if exist {
				customFieldsOut := make([]*models.CustomFieldOption, 0)
				if err := json.Unmarshal([]byte(value), &customFieldsOut); err != nil {
					results[i] = &dataloader.Result{
						Error: errs.NewErr(
							errs.ServerInternalErrorCode,
							errors.Wrapf(err, "dataloader: [ticketFieldCustomFieldOptions] json unmarshal failed"),
						)}
					return
				}
				results[i] = &dataloader.Result{Data: customFieldsOut}
			} else {
				customFieldsOut, err := l.service.GetTicketFieldCustomFieldOption(ctx, int(fieldID64))
				if err != nil {
					switch err {
					case models.ErrNotFound:
						results[i] = &dataloader.Result{
							Error: errs.NewErr(
								errs.RecordNotFoundErrorCode,
								errors.Wrapf(err, "dataloader: [ticketFieldCustomFieldOptions] service.GetTicketFieldCustomFieldOption not found"),
							)}
					default:
						results[i] = &dataloader.Result{
							Error: errs.NewErr(
								errs.ServerInternalErrorCode,
								errors.Wrapf(err, "dataloader: [ticketFieldCustomFieldOptions] service.GetTicketFieldCustomFieldOption failed"),
							)}
					}
					return
				}
				results[i] = &dataloader.Result{Data: customFieldsOut}

				// Set key-value to cache.
				if b, err := json.Marshal(customFieldsOut); err == nil {
					l.service.TicketFieldCustomFieldOptionCacheSet(ctx, key.String(), string(b))
				}
			}
		}(i, key)
	}

	wg.Wait()

	return results
}

// LoadTicketFieldSystemFieldOptions implements data loader.
func LoadTicketFieldSystemFieldOptions(ctx context.Context, params interface{}) ([]*models.SystemFieldOption, error) {
	b, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	ldr, err := extract(ctx, ticketFieldSystemFieldOption)
	if err != nil {
		return nil, err
	}

	data, err := ldr.Load(ctx, dataloader.StringKey(b))()
	if err != nil {
		return nil, err
	}

	systemFieldOption, ok := data.([]*models.SystemFieldOption)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", systemFieldOption, data)
	}

	return systemFieldOption, nil
}

type ticketFieldSystemFieldOptions struct {
	service models.Service
}

func newTicketFieldSystemFieldOptionsLoader(service models.Service) dataloader.BatchFunc {
	return ticketFieldSystemFieldOptions{service: service}.loadBatch
}

func (l ticketFieldSystemFieldOptions) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var (
		n       = len(keys)
		results = make([]*dataloader.Result, n)
		wg      sync.WaitGroup
	)

	wg.Add(n)

	for i, key := range keys {
		go func(i int, key dataloader.Key) {
			defer wg.Done()

			data := inout.QuerySystemFieldOptionsIn{}
			if err := json.Unmarshal([]byte(key.String()), &data); err != nil {
				results[i] = &dataloader.Result{
					Error: errs.NewErr(
						errs.ServerInternalErrorCode,
						errors.Wrapf(err, "dataloader: [ticketFieldSystemFieldOptions] json unmarshal failed"),
					)}
				return
			}

			fieldID64, err := strconv.ParseInt(string(data.FieldID), 10, 64)
			if err != nil {
				results[i] = &dataloader.Result{
					Error: errs.NewErr(
						errs.RecordNotFoundErrorCode,
						errors.Wrapf(err, "dataloader: [ticketFieldSystemFieldOptions] parse field id to int failed"),
					)}
				return
			}

			// Get key-value from cache.
			value, exist := l.service.TicketFieldSystemFieldOptionCacheGet(ctx, key.String())
			if exist {
				systemFieldsOut := make([]*models.SystemFieldOption, 0)
				if err := json.Unmarshal([]byte(value), &systemFieldsOut); err != nil {
					results[i] = &dataloader.Result{
						Error: errs.NewErr(
							errs.ServerInternalErrorCode,
							errors.Wrapf(err, "dataloader: [ticketFieldSystemFieldOptions] json unmarshal failed"),
						)}
					return
				}
				results[i] = &dataloader.Result{Data: systemFieldsOut}
			} else {
				systemFieldsOut, err := l.service.GetTicketFieldSystemFieldOption(ctx, int(fieldID64))
				if err != nil {
					switch err {
					case models.ErrNotFound:
						results[i] = &dataloader.Result{
							Error: errs.NewErr(
								errs.RecordNotFoundErrorCode,
								errors.Wrapf(err, "dataloader: [ticketFieldSystemFieldOptions] service.GetTicketFieldSystemFieldOption not found"),
							)}
					default:
						results[i] = &dataloader.Result{
							Error: errs.NewErr(
								errs.ServerInternalErrorCode,
								errors.Wrapf(err, "dataloader: [ticketFieldSystemFieldOptions] service.GetTicketFieldSystemFieldOption failed"),
							)}
					}
					return
				}
				results[i] = &dataloader.Result{Data: systemFieldsOut}

				// Set key-value to cache.
				if b, err := json.Marshal(systemFieldsOut); err == nil {
					l.service.TicketFieldSystemFieldOptionCacheSet(ctx, key.String(), string(b))
				}
			}
		}(i, key)
	}

	wg.Wait()

	return results
}
