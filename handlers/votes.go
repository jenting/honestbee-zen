package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"

	"github.com/honestbee/Zen/errs"
	"github.com/honestbee/Zen/inout"
)

// CreateVoteDecompressor combines params from URL or FORM
// and returns params in a structure that GetVoteHandler needs.
func CreateVoteDecompressor(ps httprouter.Params, r *http.Request) (interface{}, error) {
	baseParams, err := inout.FetchBaseParams(r)
	if err != nil {
		return nil, errs.NewErr(
			errs.InvalidAttributeErrorCode,
			errors.Wrapf(err, "handlers: [CreateVoteDecompressor] inout.FetchBaseParams failed"),
		)
	}

	voteValue := ps.ByName("value")
	if !(voteValue == "up" || voteValue == "down") {
		return nil, errs.NewErr(
			errs.InvalidAttributeErrorCode,
			errors.Wrapf(err, "handlers: [CreateVoteDecompressor] parse vote value failed"),
		)
	}

	articleID, err := strconv.ParseInt(ps.ByName("article_id"), 10, 64)
	if err != nil {
		return nil, errs.NewErr(
			errs.RecordNotFoundErrorCode,
			errors.Wrapf(err, "handlers: [CreateVoteDecompressor] parse article id to int failed"),
		)
	}

	return &inout.CreateVoteIn{
		ArticleID:   int(articleID),
		Value:       voteValue,
		Locale:      baseParams.Locale,
		CountryCode: baseParams.CountryCode,
	}, nil
}

// CreateVoteHandler handles get article request.
func CreateVoteHandler(ctx context.Context, e *Env, in interface{}) (interface{}, error) {
	data, ok := in.(*inout.CreateVoteIn)
	if !ok {
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Errorf("handlers: [CreateVoteHandler] cast %v into *GetArticleIn failed", in),
		)
	}

	voteResult, err := e.ZenDesk.CreateVote(ctx, data.ArticleID, data.Value, data.CountryCode, data.Locale)
	if err != nil {
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Wrapf(err, "handlers: [CreateVoteHandler] ZenDesk.CreateVote failed"),
		)
	}

	defer e.Examiner.SyncArticle(ctx, data.ArticleID, data.CountryCode, data.Locale)

	return &inout.CreateVoteOut{
		VoteSum:   voteResult.VoteSum,
		VoteCount: voteResult.VoteCount,
	}, nil
}
