package handlers

import (
	"context"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"

	"github.com/honestbee/Zen/errs"
	"github.com/honestbee/Zen/inout"
	"github.com/honestbee/Zen/models"
)

// GetCategoriesDecompressor combines params from URL or FORM
// and returns params in a structure that GetCategoriesHandler needs.
func GetCategoriesDecompressor(ps httprouter.Params, r *http.Request) (interface{}, error) {
	baseParams, err := inout.FetchBaseParams(r)
	if err != nil {
		return nil, errs.NewErr(
			errs.InvalidAttributeErrorCode,
			errors.Wrapf(err, "handlers: [GetCategoriesDecompressor] inout.FetchBaseParams failed"),
		)
	}

	return &inout.GetCategoriesIn{
		BaseIn: baseParams,
	}, nil
}

// GetCategoriesHandler handles get categories request.
func GetCategoriesHandler(ctx context.Context, e *Env, in interface{}) (interface{}, error) {
	data, ok := in.(*inout.GetCategoriesIn)
	if !ok {
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Errorf("handlers: [GetCategoriesHandler] cast %v into *GetCategoriesIn failed", in),
		)
	}

	defer e.Examiner.CheckCategories(ctx, data.CountryCode, data.Locale)

	categories, total, err := e.Service.GetCategories(ctx, &models.GetCategoriesParams{
		Locale:      data.Locale,
		CountryCode: data.CountryCode,
		PerPage:     data.PerPage,
		Page:        data.Page,
		SortBy:      data.SortBy,
		SortOrder:   data.SortOrder,
	})
	if err != nil {
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Wrapf(err, "handlers: [GetCategoriesHandler] Service.GetCategories failed"),
		)
	}

	return &inout.GetCategoriesOut{
		Categories: categories,
		BaseOut: &inout.BaseOut{
			Page:      int(math.Round(float64(data.Page)/float64(data.PerPage))) + 1,
			PerPage:   data.PerPage,
			PageCount: int(math.Ceil(float64(total) / float64(data.PerPage))),
			Count:     total,
		},
	}, nil
}

// GetCategoryKeyNameToIDDecompressor combines params from URL or FORM
// and returns params in a structure that GetCategoryKeyNameToIDHandler needs.
func GetCategoryKeyNameToIDDecompressor(ps httprouter.Params, r *http.Request) (interface{}, error) {
	baseParams, err := inout.FetchBaseParams(r)
	if err != nil {
		return nil, errs.NewErr(
			errs.InvalidAttributeErrorCode,
			errors.Wrapf(err, "handlers: [GetCategoryKeyNameToIDDecompressor] inout.FetchBaseParams failed"),
		)
	}

	categoryKeyName := ps.ByName("category_key_name")

	return &inout.GetCategoryKeyNameToIDIn{
		BaseIn:          baseParams,
		CategoryKeyName: categoryKeyName,
	}, nil
}

// GetCategoryKeyNameToIDHandler handles get key ID request.
func GetCategoryKeyNameToIDHandler(ctx context.Context, e *Env, in interface{}) (interface{}, error) {
	data, ok := in.(*inout.GetCategoryKeyNameToIDIn)
	if !ok {
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Errorf("handlers: [GetCategoryKeyNameToIDHandler] cast %v into *GetCategoryKeyNameToIDIn failed", in),
		)
	}

	categoryID, err := e.Service.GetCategoryKeyNameToID(ctx, data.CategoryKeyName, data.CountryCode)
	if err != nil {
		if err == models.ErrNotFound {
			return nil, errs.NewErr(
				errs.RecordNotFoundErrorCode,
				errors.Wrapf(err, "handlers: [GetCategoryKeyNameToIDHandler] Service.GetCategoryKeyNameToID not found"),
			)
		}
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Wrapf(err, "handlers: [GetCategoryKeyNameToIDHandler] Service.GetCategoryKeyNameToID failed"),
		)
	}

	return &inout.GetCategoryKeyNameToIDOut{
		CategoryID: categoryID,
	}, nil
}

// GetSectionsDecompressor combines params from URL or FORM
// and returns params in a structure that GetSectionsHandler needs.
func GetSectionsDecompressor(ps httprouter.Params, r *http.Request) (interface{}, error) {
	baseParams, err := inout.FetchBaseParams(r)
	if err != nil {
		return nil, errs.NewErr(
			errs.InvalidAttributeErrorCode,
			errors.Wrapf(err, "handlers: [GetSectionsDecompressor] inout.FetchBaseParams failed"),
		)
	}

	categoryID, err := strconv.ParseInt(ps.ByName("category_id"), 10, 64)
	if err != nil {
		return nil, errs.NewErr(
			errs.RecordNotFoundErrorCode,
			errors.Wrapf(err, "handlers: [GetSectionsDecompressor] parse category id to int failed"),
		)
	}

	return &inout.GetSectionsIn{
		BaseIn:     baseParams,
		CategoryID: int(categoryID),
	}, nil
}

// GetSectionsHandler handles get sections request.
func GetSectionsHandler(ctx context.Context, e *Env, in interface{}) (interface{}, error) {
	data, ok := in.(*inout.GetSectionsIn)
	if !ok {
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Errorf("handlers: [GetSectionsHandler] cast %v into *GetSectionsIn failed", in),
		)
	}

	defer e.Examiner.CheckSections(ctx, data.CountryCode, data.Locale)

	sections, total, err := e.Service.GetSectionsByCategoryID(ctx,
		&models.GetSectionsParams{
			Locale:      data.Locale,
			CountryCode: data.CountryCode,
			PerPage:     data.PerPage,
			Page:        data.Page,
			SortBy:      data.SortBy,
			SortOrder:   data.SortOrder,
			CategoryID:  data.CategoryID,
		})
	if err != nil {
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Wrapf(err, "handlers: [GetSectionsHandler] Service.GetSectionsByCategoryID failed"),
		)
	}

	return &inout.GetSectionsOut{
		Sections: sections,
		BaseOut: &inout.BaseOut{
			Page:      int(math.Round(float64(data.Page)/float64(data.PerPage))) + 1,
			PerPage:   data.PerPage,
			PageCount: int(math.Ceil(float64(total) / float64(data.PerPage))),
			Count:     total,
		},
	}, nil
}

// GetCategoriesArticlesDecompressor combines params from URL or FORM
// and returns params in a structure that GetArticlesHandler needs.
func GetCategoriesArticlesDecompressor(ps httprouter.Params, r *http.Request) (interface{}, error) {
	baseParams, err := inout.FetchBaseParams(r)
	if err != nil {
		return nil, errs.NewErr(
			errs.InvalidAttributeErrorCode,
			errors.Wrapf(err, "handlers: [GetCategoriesArticlesDecompressor] inout.FetchBaseParams failed"),
		)
	}

	categoryID, err := strconv.ParseInt(ps.ByName("category_id"), 10, 64)
	if err != nil {
		return nil, errs.NewErr(
			errs.InvalidAttributeErrorCode,
			errors.Wrapf(err, "handlers: [GetCategoriesArticlesDecompressor] parse category id to int failed"),
		)
	}

	labelNames := r.FormValue("label_names")

	return &inout.GetCategoriesArticlesIn{
		CategoryID: int(categoryID),
		LabelNames: labelNames,
		BaseIn:     baseParams,
	}, nil
}

// GetCategoriesArticlesHandler handles get sections request.
func GetCategoriesArticlesHandler(ctx context.Context, e *Env, in interface{}) (interface{}, error) {
	data, ok := in.(*inout.GetCategoriesArticlesIn)
	if !ok {
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Errorf("handlers: [GetCategoriesArticlesHandler] cast %v into *GetCategoriesArticlesIn failed", in),
		)
	}

	var labels []string
	if data.LabelNames != "" {
		labels = strings.Split(data.LabelNames, ",")
	}

	articles, total, err := e.Service.GetArticlesByCategoryID(ctx,
		&models.GetArticlesParams{
			Locale:      data.Locale,
			CountryCode: data.CountryCode,
			PerPage:     data.PerPage,
			Page:        data.Page,
			SortBy:      data.SortBy,
			SortOrder:   data.SortOrder,
			CategoryID:  data.CategoryID,
		}, labels)
	if err != nil {
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Wrapf(err, "handlers: [GetSectionsHandler] Service.GetSections failed"),
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
