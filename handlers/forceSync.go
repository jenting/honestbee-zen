package handlers

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"

	"github.com/honestbee/Zen/errs"
	"github.com/honestbee/Zen/inout"
)

// CreateForceSyncDecompressor combines params from authorization header
// and returns params in a structure that CreateForceSyncHandler needs.
func CreateForceSyncDecompressor(ps httprouter.Params, r *http.Request) (interface{}, error) {
	gotUser, gotPwd, ok := r.BasicAuth()
	if !ok {
		return nil, errs.NewErr(
			errs.InvalidAttributeErrorCode,
			errors.Errorf("handlers: [CreateForceSyncDecompressor] fetchBasicAuth failed"),
		)
	}

	return &inout.GetBasicAuthIn{
		User: gotUser,
		Pwd:  gotPwd,
	}, nil
}

// CreateForceSyncHandler handles force sync request.
func CreateForceSyncHandler(ctx context.Context, e *Env, in interface{}) (interface{}, error) {
	data, ok := in.(*inout.GetBasicAuthIn)
	if !ok {
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Errorf("handlers: [CreateForceSyncHandler] cast %v into *GetBasicAuthIn failed", in),
		)
	}

	if e.Config.HTTP.BasicAuthUser != data.User || e.Config.HTTP.BasicAuthPwd != data.Pwd {
		return nil, errs.NewErr(
			errs.UnauthorizedErrCode,
			errors.Errorf("handlers: [CreateForceSyncHandler] user or password not match"),
		)
	}

	go func() {
		// Loop by countryCode.
		for countryCode, locales := range inout.SupportCountryLocaleMap {
			for _, locale := range locales {
				e.Logger.Info().Msgf("force sync categories, country code %s locales %s", countryCode, locale)
				if err := e.Examiner.ForceSyncCategories(ctx, countryCode, locale); err != nil {
					e.Logger.Error().Err(err).Msgf("sync categories, country code %s locale %s failed", countryCode, locale)
				}

				e.Logger.Info().Msgf("force sync sections, country code %s locales %s", countryCode, locale)
				if err := e.Examiner.ForceSyncSections(ctx, countryCode, locale); err != nil {
					e.Logger.Error().Err(err).Msgf("sync sections, country code %s locale %s failed", countryCode, locale)
				}

				e.Logger.Info().Msgf("force sync articles, country code %s locales %s", countryCode, locale)
				if err := e.Examiner.ForceSyncArticles(ctx, countryCode, locale); err != nil {
					e.Logger.Error().Err(err).Msgf("sync articles, country code %s locale %s failed", countryCode, locale)
				}
			}
		}

		e.Logger.Info().Msgf("force sync ticket forms")
		if err := e.Examiner.ForceSyncTicketForms(ctx); err != nil {
			e.Logger.Error().Err(err).Msgf("sync ticket forms failed")
		}
	}()

	return inout.SuccessForceSync, nil
}
