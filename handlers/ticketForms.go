package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"

	"github.com/honestbee/Zen/errs"
	"github.com/honestbee/Zen/inout"
	"github.com/honestbee/Zen/models"
)

// GetTicketFormDecompressor combines params from URL or FORM
// and returns params in a structure that GetTicketFormHandler needs.
func GetTicketFormDecompressor(ps httprouter.Params, r *http.Request) (interface{}, error) {
	baseParams, err := inout.FetchBaseParams(r)
	if err != nil {
		return nil, errs.NewErr(
			errs.InvalidAttributeErrorCode,
			errors.Wrapf(err, "handlers: [GetTicketFormDecompressor] inout.FetchBaseParams failed"),
		)
	}

	formID, err := strconv.ParseInt(ps.ByName("form_id"), 10, 64)
	if err != nil {
		return nil, errs.NewErr(
			errs.RecordNotFoundErrorCode,
			errors.Wrapf(err, "handlers: [GetTicketFormDecompressor] parse form id to int failed"),
		)
	}

	return &inout.GetTicketFormIn{
		CountryCode: baseParams.CountryCode,
		Locale:      baseParams.Locale,
		FormID:      int(formID),
	}, nil
}

// GetTicketFormHandler handles get ticket form request.
func GetTicketFormHandler(ctx context.Context, e *Env, in interface{}) (interface{}, error) {
	data, ok := in.(*inout.GetTicketFormIn)
	if !ok {
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Errorf("handlers: [GetTicketFormHandler] cast %v into *GetTicketFormIn failed", in),
		)
	}

	defer e.Examiner.CheckTicketForms(ctx)

	form, err := e.Service.GetTicketForm(ctx, data.FormID, data.Locale)
	if err != nil {
		switch err {
		case models.ErrNotFound:
			return nil, errs.NewErr(
				errs.RecordNotFoundErrorCode,
				errors.Wrapf(err, "handlers: [GetTicketFormHandler] Service.GetTicketForm formID:%d not found", data.FormID),
			)
		default:
			return nil, errs.NewErr(
				errs.ServerInternalErrorCode,
				errors.Wrapf(err, "handlers: [GetTicketFormHandler] Service.GetTicketForm formID:%d failed", data.FormID),
			)
		}
	}

	return &inout.GetTicketFormOut{
		TicketForm: form,
	}, nil
}
