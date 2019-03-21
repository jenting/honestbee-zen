package handlers

import (
	"context"
	"math"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"

	"github.com/honestbee/Zen/errs"
	"github.com/honestbee/Zen/inout"
	"github.com/honestbee/Zen/models"
)

// GetArticlesDecompressor combines params from URL or FORM
// and returns params in a structure that GetArticlesHandler needs.
func GetArticlesDecompressor(ps httprouter.Params, r *http.Request) (interface{}, error) {
	baseParams, err := inout.FetchBaseParams(r)
	if err != nil {
		return nil, errs.NewErr(
			errs.InvalidAttributeErrorCode,
			errors.Wrapf(err, "handlers: [GetArticlesDecompressor] baseParams failed"),
		)
	}

	sectionID, err := strconv.ParseInt(ps.ByName("section_id"), 10, 64)
	if err != nil {
		return nil, errs.NewErr(
			errs.RecordNotFoundErrorCode,
			errors.Wrapf(err, "handlers: [GetArticlesDecompressor] parse section id to int failed"),
		)
	}

	return &inout.GetArticlesIn{
		BaseIn:    baseParams,
		SectionID: int(sectionID),
	}, nil
}

// GetArticlesHandler handles get articles request.
func GetArticlesHandler(ctx context.Context, e *Env, in interface{}) (interface{}, error) {
	data, ok := in.(*inout.GetArticlesIn)
	if !ok {
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Errorf("handlers: [GetArticlesHandler] cast %v into *GetArticlesIn failed", in),
		)
	}

	defer e.Examiner.CheckArticles(ctx, data.CountryCode, data.Locale)

	articles, total, err := e.Service.GetArticlesBySectionID(ctx,
		&models.GetArticlesParams{
			Locale:      data.Locale,
			CountryCode: data.CountryCode,
			PerPage:     data.PerPage,
			Page:        data.Page,
			SortBy:      data.SortBy,
			SortOrder:   data.SortOrder,
			SectionID:   data.SectionID,
		})
	if err != nil {
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Wrapf(err, "handlers: [GetArticlesHandler] Service.GetArticles failed"),
		)
	}

	return &inout.GetArticlesOut{
		Articles: articles,
		BaseOut: &inout.BaseOut{
			Page:      int(math.Round(float64(data.Page)/float64(data.PerPage))) + 1,
			PerPage:   data.PerPage,
			PageCount: int(math.Ceil(float64(total) / float64(data.PerPage))),
			Count:     total,
		},
	}, nil
}

// GetSectionDecompressor combines params from URL or FORM
// and returns params in a structure that GetSectionHandler needs.
func GetSectionDecompressor(ps httprouter.Params, r *http.Request) (interface{}, error) {
	baseParams, err := inout.FetchBaseParams(r)
	if err != nil {
		return nil, errs.NewErr(
			errs.InvalidAttributeErrorCode,
			errors.Wrapf(err, "handlers: [GetSectionDecompressor] baseParams failed"),
		)
	}

	sectionID, err := strconv.ParseInt(ps.ByName("section_id"), 10, 64)
	if err != nil {
		return nil, errs.NewErr(
			errs.RecordNotFoundErrorCode,
			errors.Wrapf(err, "handlers: [GetSectionDecompressor] parse section id to int failed"),
		)
	}

	return &inout.GetSectionIn{
		BaseIn:    baseParams,
		SectionID: int(sectionID),
	}, nil
}

// GetSectionHandler handles show section request.
func GetSectionHandler(ctx context.Context, e *Env, in interface{}) (interface{}, error) {
	data, ok := in.(*inout.GetSectionIn)
	if !ok {
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Errorf("handlers: [GetSectionHandler] cast %v into *GetSectionIn failed", in),
		)
	}

	defer e.Examiner.CheckSections(ctx, data.CountryCode, data.Locale)

	section, err := e.Service.GetSectionBySectionID(ctx, data.SectionID, data.Locale, data.CountryCode)
	if err != nil {
		if err == models.ErrNotFound {
			return nil, errs.NewErr(
				errs.RecordNotFoundErrorCode,
				errors.Wrapf(err, "handlers: [GetSectionHandler] Service.GetSection not found"),
			)
		}
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Wrapf(err, "handlers: [GetSectionHandler] Service.GetSection failed"),
		)
	}
	return &inout.GetSectionOut{
		Section: section,
	}, nil
}
