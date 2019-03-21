package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"

	"github.com/honestbee/Zen/errs"
	"github.com/honestbee/Zen/inout"
)

// CreateRequestDecompressor combines params from URL or FORM
// and returns params in a structure that CreateRequestHandler needs.
func CreateRequestDecompressor(ps httprouter.Params, r *http.Request) (interface{}, error) {
	ret := new(inout.CreateRequestIn)
	if err := json.NewDecoder(r.Body).Decode(ret); err != nil {
		return nil, errs.NewErr(
			errs.InvalidAttributeErrorCode,
			errors.Wrapf(err, "handlers: [CreateRequestDecompressor] json decode failed"),
		)
	}

	if ret.CountryCode == "" {
		return nil, errs.NewErr(
			errs.InvalidAttributeErrorCode,
			errors.Errorf("handlers: [CreateRequestDecompressor] country code is empty"),
		)
	}

	return ret, nil
}

// CreateRequestHandler handles "create request" request.
func CreateRequestHandler(ctx context.Context, e *Env, in interface{}) (interface{}, error) {
	request, ok := in.(*inout.CreateRequestIn)
	if !ok {
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Errorf("handlers: [CreateRequestHandler] cast %v into *CreateRequestIn failed", in),
		)
	}

	if err := e.ZenDesk.CreateRequest(ctx, request.CountryCode, request.Data); err != nil {
		return nil, errs.NewErr(
			errs.InvalidAttributeErrorCode,
			errors.Wrapf(err, "handlers: [CreateRequestHandler] zendesk create request failed"),
		)
	}

	return nil, errs.NewErr(errs.SuccessCreatedCode, nil)
}
