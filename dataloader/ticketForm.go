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
	"github.com/honestbee/Zen/examiner"
	"github.com/honestbee/Zen/inout"
	"github.com/honestbee/Zen/models"
)

// LoadTicketForm implements data loader.
func LoadTicketForm(ctx context.Context, params interface{}) (*models.SyncTicketForm, error) {
	b, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	ldr, err := extract(ctx, ticketFormLoaderKey)
	if err != nil {
		return nil, err
	}

	data, err := ldr.Load(ctx, dataloader.StringKey(b))()
	if err != nil {
		return nil, err
	}

	ticketForm, ok := data.(*models.SyncTicketForm)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", ticketForm, data)
	}

	return ticketForm, nil
}

type ticketFormLoader struct {
	service  models.Service
	examiner *examiner.Examiner
}

func newTicketFormLoader(service models.Service, examiner *examiner.Examiner) dataloader.BatchFunc {
	return ticketFormLoader{service: service, examiner: examiner}.loadBatch
}

func (l ticketFormLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var (
		n       = len(keys)
		results = make([]*dataloader.Result, n)
		wg      sync.WaitGroup
	)

	wg.Add(n)

	for i, key := range keys {
		go func(i int, key dataloader.Key) {
			defer wg.Done()

			data := inout.QueryTicketFormIn{}
			if err := json.Unmarshal([]byte(key.String()), &data); err != nil {
				results[i] = &dataloader.Result{
					Error: errs.NewErr(
						errs.ServerInternalErrorCode,
						errors.Wrapf(err, "dataloader: [ticketFormLoader] json unmarshal failed"),
					)}
				return
			}

			formID64, err := strconv.ParseInt(string(data.FormID), 10, 64)
			if err != nil {
				results[i] = &dataloader.Result{
					Error: errs.NewErr(
						errs.RecordNotFoundErrorCode,
						errors.Wrapf(err, "dataloader: [ticketFormLoader] parse form id to int failed"),
					)}
				return
			}

			defer l.examiner.CheckTicketForms(ctx)

			// Get key-value from cache.
			value, exist := l.service.TicketFormCacheGet(ctx, key.String())
			if exist {
				ticketFormOut := &models.SyncTicketForm{}
				if err := json.Unmarshal([]byte(value), &ticketFormOut); err != nil {
					results[i] = &dataloader.Result{
						Error: errs.NewErr(
							errs.ServerInternalErrorCode,
							errors.Wrapf(err, "dataloader: [ticketFormLoader] json unmarshal failed"),
						)}
					return
				}
				results[i] = &dataloader.Result{Data: ticketFormOut}
			} else {
				ticketFormOut, err := l.service.GetTicketFormGraphQL(ctx, int(formID64))
				if err != nil {
					switch err {
					case models.ErrNotFound:
						results[i] = &dataloader.Result{
							Error: errs.NewErr(
								errs.RecordNotFoundErrorCode,
								errors.Wrapf(err, "dataloader: [ticketFormLoader] service.GetTicketFormByFormID not found"),
							)}
					default:
						results[i] = &dataloader.Result{
							Error: errs.NewErr(
								errs.ServerInternalErrorCode,
								errors.Wrapf(err, "dataloader: [ticketFormLoader] service.GetTicketFormByFormID failed"),
							)}
					}
					return
				}
				results[i] = &dataloader.Result{Data: ticketFormOut}

				// Set key-value to cache.
				if b, err := json.Marshal(ticketFormOut); err == nil {
					l.service.TicketFormCacheSet(ctx, key.String(), string(b))
				}
			}
		}(i, key)
	}

	wg.Wait()

	return results
}
