package resolvers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	"github.com/honestbee/Zen/errs"
	"github.com/honestbee/Zen/inout"
)

// CreateRequest create a new createRequest resolver.
func (r *Resolver) CreateRequest(ctx context.Context, data inout.MutationRequestsIn) (*string, error) {
	// Process input params.
	if err := data.ProcessInputParams(); err != nil {
		return nil, err
	}

	if err := r.zendesk.CreateRequest(ctx, data.CountryCode, data.Data); err != nil {
		return nil, errs.NewErr(
			errs.InvalidAttributeErrorCode,
			errors.Wrapf(err, "resolver: [CreateRequest] zendesk.CreateRequest failed"),
		)
	}

	ret := http.StatusText(http.StatusCreated)
	return &ret, nil
}

// VoteArticle create a new voteArticle resolver.
func (r *Resolver) VoteArticle(ctx context.Context, data inout.MutationVoteArticleIn) (*ArticleResolver, error) {
	// Process input params.
	if err := data.ProcessInputParams(); err != nil {
		return nil, errs.NewErr(
			errs.InvalidAttributeErrorCode,
			errors.Wrapf(err, "resolver: [VoteArticle] invalid input params"),
		)
	}

	articleID64, err := strconv.ParseInt(string(data.ArticleID), 10, 64)
	if err != nil {
		return nil, errs.NewErr(
			errs.InvalidAttributeErrorCode,
			errors.Wrapf(err, "resolver: [VoteArticle] parse article id to int failed"),
		)
	}

	voteResult, err := r.zendesk.CreateVote(ctx, int(articleID64), data.Vote, data.CountryCode, data.Locale)
	if err != nil {
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Wrapf(err, "resolver: [VoteArticle] zendesk.CreateVote failed"),
		)
	}

	defer r.examiner.SyncArticle(ctx, int(articleID64), data.CountryCode, data.Locale)

	articleOut, err := r.service.GetArticleByArticleID(ctx, int(articleID64), data.Locale, data.CountryCode)
	if err != nil {
		return nil, errs.NewErr(
			errs.RecordNotFoundErrorCode,
			errors.Wrapf(err, "resolver: [VoteArticle] service.GetArticleByArticleID failed"),
		)
	}
	articleOut.VoteSum = voteResult.VoteSum
	articleOut.VoteCount = voteResult.VoteCount

	return &ArticleResolver{m: articleOut}, nil
}

// ForceSync create a new forceSync resolver.
func (r *Resolver) ForceSync(ctx context.Context, data inout.MutationForceSyncIn) (*string, error) {
	if r.conf.HTTP.BasicAuthUser != data.Username || r.conf.HTTP.BasicAuthPwd != data.Password {
		return nil, errs.NewErr(
			errs.UnauthorizedErrCode,
			errors.Errorf("resolver: [ForceSync] user or password not match"),
		)
	}

	go func() {
		// Loop by countryCode.
		for countryCode, locales := range inout.SupportCountryLocaleMap {
			for _, locale := range locales {
				r.logger.Info().Msgf("force sync categories, country code %s locales %s", countryCode, locale)
				if err := r.examiner.ForceSyncCategories(ctx, countryCode, locale); err != nil {
					r.logger.Error().Err(err).Msgf("sync categories, country code %s locale %s failed", countryCode, locale)
				}

				r.logger.Info().Msgf("force sync sections, country code %s locales %s", countryCode, locale)
				if err := r.examiner.ForceSyncSections(ctx, countryCode, locale); err != nil {
					r.logger.Error().Err(err).Msgf("sync sections, country code %s locale %s failed", countryCode, locale)
				}

				r.logger.Info().Msgf("force sync articles, country code %s locales %s", countryCode, locale)
				if err := r.examiner.ForceSyncArticles(ctx, countryCode, locale); err != nil {
					r.logger.Error().Err(err).Msgf("sync articles, country code %s locale %s failed", countryCode, locale)
				}
			}
		}

		r.logger.Info().Msgf("force sync ticket forms")
		if err := r.examiner.ForceSyncTicketForms(ctx); err != nil {
			r.logger.Error().Err(err).Msgf("sync ticket forms failed")
		}
	}()

	ret := inout.SuccessForceSync
	return &ret, nil
}
