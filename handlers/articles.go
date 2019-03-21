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

// GetArticleDecompressor combines params from URL or FORM
// and returns params in a structure that GetArticleHandler needs.
func GetArticleDecompressor(ps httprouter.Params, r *http.Request) (interface{}, error) {
	baseParams, err := inout.FetchBaseParams(r)
	if err != nil {
		return nil, errs.NewErr(
			errs.InvalidAttributeErrorCode,
			errors.Wrapf(err, "handlers: [GetArticleDecompressor] inout.FetchBaseParams failed"),
		)
	}

	articleID, err := strconv.ParseInt(ps.ByName("article_id"), 10, 64)
	if err != nil {
		return nil, errs.NewErr(
			errs.RecordNotFoundErrorCode,
			errors.Wrapf(err, "handlers: [GetArticleDecompressor] parse article id to int failed"),
		)
	}

	return &inout.GetArticleIn{
		Locale:      baseParams.Locale,
		CountryCode: baseParams.CountryCode,
		ArticleID:   int(articleID),
	}, nil
}

// GetArticleHandler handles get article request.
func GetArticleHandler(ctx context.Context, e *Env, in interface{}) (interface{}, error) {
	data, ok := in.(*inout.GetArticleIn)
	if !ok {
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Errorf("handlers: [GetArticleHandler] cast %v into *GetArticleIn failed", in),
		)
	}

	defer e.Examiner.CheckArticles(ctx, data.CountryCode, data.Locale)
	defer e.Service.PlusOneArticleClickCounter(ctx, data.ArticleID, data.Locale, data.CountryCode)

	article, err := e.Service.GetArticleByArticleID(ctx, data.ArticleID, data.Locale, data.CountryCode)
	if err != nil {
		if err == models.ErrNotFound {
			return nil, errs.NewErr(
				errs.RecordNotFoundErrorCode,
				errors.Wrapf(err, "handlers: [GetArticleHandler] Service.GetArticle not found"),
			)
		}
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Wrapf(err, "handlers: [GetArticleHandler] Service.GetArticle failed"),
		)
	}

	return &inout.GetArticleOut{
		Article: article,
	}, nil
}

// GetTopNArticlesDecompressor combines params from URL or FORM
// and returns params in a structure that GetTopNArticlesHandler needs.
func GetTopNArticlesDecompressor(ps httprouter.Params, r *http.Request) (interface{}, error) {
	baseParams, err := inout.FetchBaseParams(r)
	if err != nil {
		return nil, errs.NewErr(
			errs.InvalidAttributeErrorCode,
			errors.Wrapf(err, "handlers: [GetTopNArticlesDecompressor] inout.FetchBaseParams failed"),
		)
	}

	topN, err := strconv.ParseUint(ps.ByName("top_n"), 10, 64)
	if err != nil {
		return nil, errs.NewErr(
			errs.InvalidAttributeErrorCode,
			errors.Wrapf(err, "handlers: [GetTopNArticlesDecompressor] parse top_n to uint64 failed"),
		)
	}

	return &inout.GetTopNArticlesIn{
		TopN:        topN,
		Locale:      baseParams.Locale,
		CountryCode: baseParams.CountryCode,
	}, nil
}

// GetTopNArticlesHandler handles get topN articles request.
func GetTopNArticlesHandler(ctx context.Context, e *Env, in interface{}) (interface{}, error) {
	data, ok := in.(*inout.GetTopNArticlesIn)
	if !ok {
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Errorf("handlers: [GetTopNArticlesHandler] cast %v into *GetTopNArticlesIn failed", in),
		)
	}

	articles, err := e.Service.GetTopNArticles(ctx, data.TopN, data.Locale, data.CountryCode)
	if err != nil {
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Wrapf(err, "handlers: [GetTopNArticlesHandler] Service.GetTopNArticles failed"),
		)
	}

	return &inout.GetTopNArticlesOut{
		Articles: articles,
	}, nil
}
