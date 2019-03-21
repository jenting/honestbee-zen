package grpc

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"runtime"
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	grpctrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/google.golang.org/grpc"

	"github.com/honestbee/Zen/config"
	"github.com/honestbee/Zen/errs"
	"github.com/honestbee/Zen/examiner"
	"github.com/honestbee/Zen/inout"
	"github.com/honestbee/Zen/models"
	"github.com/honestbee/Zen/protobuf"
	"github.com/honestbee/Zen/zendesk"
)

type server struct {
	conf     *config.Config
	logger   *zerolog.Logger
	service  models.Service
	examiner *examiner.Examiner
	zend     *zendesk.ZenDesk
}

// New register ZendeskServer instance to gRPC server and returns it.
func New(conf *config.Config,
	logger *zerolog.Logger,
	service models.Service,
	examiner *examiner.Examiner,
	zend *zendesk.ZenDesk) (*grpc.Server, error) {
	// Initialize the grpc server as normal, using the tracing and logging interceptor.
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpcmiddleware.ChainUnaryServer(
			logUnaryInterceptor(logger),
			grpctrace.UnaryServerInterceptor(grpctrace.WithServiceName("helpcenter-zendesk-grpc")),
		)),
	)

	// Register service.
	protobuf.RegisterZendeskServer(s, &server{
		conf:     conf,
		logger:   logger,
		service:  service,
		examiner: examiner,
		zend:     zend,
	})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	return s, nil
}

func (s *server) GetCategories(ctx context.Context, in *protobuf.GetCategoriesRequest) (*protobuf.GetCategoriesResponse, error) {
	perPage, page := inout.ProcessPage(in.PerPage, in.Page)

	defer s.examiner.CheckCategories(ctx, inout.GRPCCountryCodeMap[in.CountryCode], inout.GRPCLocaleMap[in.Locale])

	categories, total, err := s.service.GetCategories(ctx,
		&models.GetCategoriesParams{
			Locale:      inout.GRPCLocaleMap[in.Locale],
			CountryCode: inout.GRPCCountryCodeMap[in.CountryCode],
			PerPage:     int(perPage),
			Page:        int(page),
			SortBy:      inout.GRPCSortByMap[in.SortBy],
			SortOrder:   inout.GRPCSortOrderMap[in.SortOrder],
		})
	if err != nil {
		if err == models.ErrNotFound {
			return nil, errs.NewErr(
				errs.RecordNotFoundErrorCode,
				errors.Wrapf(err, "grpc: [GetCategories] not found"),
			)
		}
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Wrapf(err, "grpc: [GetCategories] failed"),
		)
	}

	categoriesOut := &protobuf.GetCategoriesResponse{
		PageInfo: &protobuf.PageInfo{
			Page:      int32(math.Round(float64(page)/float64(perPage))) + 1,
			PerPage:   int32(perPage),
			PageCount: int32(math.Ceil(float64(total) / float64(perPage))),
			Count:     int32(total),
		},
		Categories: make([]*protobuf.Category, 0),
	}
	for _, category := range categories {
		outCategory := &protobuf.Category{
			Id:           strconv.Itoa(category.ID),
			Position:     int32(category.Position),
			SourceLocale: category.SourceLocale,
			Outdated:     category.Outdated,
			CountryCode:  category.CountryCode,
			KeyName:      category.KeyName,
			Url:          category.URL,
			HtmlUrl:      category.HTMLURL,
			Name:         category.Name,
			Description:  category.Description,
			Locale:       category.Locale,
		}
		// Due to timestamp format needs to be convert,
		// we cannot use json Marshal + Unmarshal to converts format.
		outCategory.CreatedAt, _ = ptypes.TimestampProto(category.CreatedAt)
		outCategory.UpdatedAt, _ = ptypes.TimestampProto(category.UpdatedAt)

		categoriesOut.Categories = append(categoriesOut.Categories, outCategory)
	}

	return categoriesOut, nil
}

func (s *server) GetCategory(ctx context.Context, in *protobuf.GetCategoryRequest) (*protobuf.GetCategoryResponse, error) {
	switch in.Id.(type) {
	case *protobuf.GetCategoryRequest_CategoryIdOrKeyname:
		defer s.examiner.CheckCategories(ctx, inout.GRPCCountryCodeMap[in.CountryCode], inout.GRPCLocaleMap[in.Locale])

		category, err := s.service.GetCategoryByCategoryIDOrKeyName(ctx,
			in.GetCategoryIdOrKeyname(),
			inout.GRPCLocaleMap[in.Locale],
			inout.GRPCCountryCodeMap[in.CountryCode],
		)
		if err != nil {
			if err == models.ErrNotFound {
				return nil, errs.NewErr(
					errs.RecordNotFoundErrorCode,
					errors.Wrapf(err, "grpc: [GetCategoryByCategoryIDOrKeyName] not found"),
				)
			}
			return nil, errs.NewErr(
				errs.ServerInternalErrorCode,
				errors.Wrapf(err, "grpc: [GetCategoryByCategoryIDOrKeyName] failed"),
			)
		}

		out := &protobuf.GetCategoryResponse{
			Category: &protobuf.Category{
				Id:           strconv.Itoa(category.ID),
				Position:     int32(category.Position),
				SourceLocale: category.SourceLocale,
				Outdated:     category.Outdated,
				CountryCode:  category.CountryCode,
				KeyName:      category.KeyName,
				Url:          category.URL,
				HtmlUrl:      category.HTMLURL,
				Name:         category.Name,
				Description:  category.Description,
				Locale:       category.Locale,
			},
		}
		// Due to timestamp format needs to be convert,
		// we cannot use json Marshal + Unmarshal to converts format.
		out.Category.CreatedAt, _ = ptypes.TimestampProto(category.CreatedAt)
		out.Category.UpdatedAt, _ = ptypes.TimestampProto(category.UpdatedAt)

		return out, nil
	case *protobuf.GetCategoryRequest_SectionId:
		sectionID, err := strconv.Atoi(in.GetSectionId())
		if err != nil {
			return nil, errs.NewErr(
				errs.InvalidAttributeErrorCode,
				errors.Wrapf(err, "grpc: [GetCategoryBySectionID] failed"),
			)
		}

		defer s.examiner.CheckCategories(ctx, inout.GRPCCountryCodeMap[in.CountryCode], inout.GRPCLocaleMap[in.Locale])

		category, err := s.service.GetCategoryBySectionID(ctx, sectionID, inout.GRPCLocaleMap[in.Locale])
		if err != nil {
			if err == models.ErrNotFound {
				return nil, errs.NewErr(
					errs.RecordNotFoundErrorCode,
					errors.Wrapf(err, "grpc: [GetCategoryBySectionID] not found"),
				)
			}
			return nil, errs.NewErr(
				errs.ServerInternalErrorCode,
				errors.Wrapf(err, "grpc: [GetCategoryBySectionID] failed"),
			)
		}

		out := &protobuf.GetCategoryResponse{
			Category: &protobuf.Category{
				Id:           strconv.Itoa(category.ID),
				Position:     int32(category.Position),
				SourceLocale: category.SourceLocale,
				Outdated:     category.Outdated,
				CountryCode:  category.CountryCode,
				KeyName:      category.KeyName,
				Url:          category.URL,
				HtmlUrl:      category.HTMLURL,
				Name:         category.Name,
				Description:  category.Description,
				Locale:       category.Locale,
			},
		}
		// Due to timestamp format needs to be convert,
		// we cannot use json Marshal + Unmarshal to converts format.
		out.Category.CreatedAt, _ = ptypes.TimestampProto(category.CreatedAt)
		out.Category.UpdatedAt, _ = ptypes.TimestampProto(category.UpdatedAt)

		return out, nil
	case *protobuf.GetCategoryRequest_ArticleId:
		articleID, err := strconv.Atoi(in.GetArticleId())
		if err != nil {
			return nil, errs.NewErr(
				errs.InvalidAttributeErrorCode,
				errors.Wrapf(err, "grpc: [GetCategoryByArticleID] failed"),
			)
		}

		defer s.examiner.CheckCategories(ctx, inout.GRPCCountryCodeMap[in.CountryCode], inout.GRPCLocaleMap[in.Locale])

		category, err := s.service.GetCategoryByArticleID(ctx, articleID, inout.GRPCLocaleMap[in.Locale])
		if err != nil {
			if err == models.ErrNotFound {
				return nil, errs.NewErr(
					errs.RecordNotFoundErrorCode,
					errors.Wrapf(err, "grpc: [GetCategoryByArticleID] not found"),
				)
			}
			return nil, errs.NewErr(
				errs.ServerInternalErrorCode,
				errors.Wrapf(err, "grpc: [GetCategoryByArticleID] failed"),
			)
		}

		out := &protobuf.GetCategoryResponse{
			Category: &protobuf.Category{
				Id:           strconv.Itoa(category.ID),
				Position:     int32(category.Position),
				SourceLocale: category.SourceLocale,
				Outdated:     category.Outdated,
				CountryCode:  category.CountryCode,
				KeyName:      category.KeyName,
				Url:          category.URL,
				HtmlUrl:      category.HTMLURL,
				Name:         category.Name,
				Description:  category.Description,
				Locale:       category.Locale,
			},
		}
		// Due to timestamp format needs to be convert,
		// we cannot use json Marshal + Unmarshal to converts format.
		out.Category.CreatedAt, _ = ptypes.TimestampProto(category.CreatedAt)
		out.Category.UpdatedAt, _ = ptypes.TimestampProto(category.UpdatedAt)

		return out, nil
	}

	return nil, fmt.Errorf("[GetCategory] invalid type: %v", in.Id)
}

func (s *server) GetSections(ctx context.Context, in *protobuf.GetSectionsRequest) (*protobuf.GetSectionsResponse, error) {
	switch in.Id.(type) {
	case *protobuf.GetSectionsRequest_All:
		perPage, page := inout.ProcessPage(in.PerPage, in.Page)

		defer s.examiner.CheckSections(ctx, inout.GRPCCountryCodeMap[in.CountryCode], inout.GRPCLocaleMap[in.Locale])

		sections, total, err := s.service.GetSections(ctx,
			&models.GetSectionsParams{
				Locale:      inout.GRPCLocaleMap[in.Locale],
				CountryCode: inout.GRPCCountryCodeMap[in.CountryCode],
				PerPage:     int(perPage),
				Page:        int(page),
				SortBy:      inout.GRPCSortByMap[in.SortBy],
				SortOrder:   inout.GRPCSortOrderMap[in.SortOrder],
			})
		if err != nil {
			if err == models.ErrNotFound {
				return nil, errs.NewErr(
					errs.RecordNotFoundErrorCode,
					errors.Wrapf(err, "grpc: [GetSections] not found"),
				)
			}
			return nil, errs.NewErr(
				errs.ServerInternalErrorCode,
				errors.Wrapf(err, "grpc: [GetSections] failed"),
			)
		}

		sectionsOut := &protobuf.GetSectionsResponse{
			PageInfo: &protobuf.PageInfo{
				Page:      int32(math.Round(float64(page)/float64(perPage))) + 1,
				PerPage:   int32(perPage),
				PageCount: int32(math.Ceil(float64(total) / float64(perPage))),
				Count:     int32(total),
			},
			Sections: make([]*protobuf.Section, 0),
		}
		for _, section := range sections {
			outSection := &protobuf.Section{
				Id:           strconv.Itoa(section.ID),
				Position:     int32(section.Position),
				SourceLocale: section.SourceLocale,
				Outdated:     section.Outdated,
				CountryCode:  section.CountryCode,
				Url:          section.URL,
				HtmlUrl:      section.HTMLURL,
				Name:         section.Name,
				Description:  section.Description,
				Locale:       section.Locale,
				CategoryId:   strconv.Itoa(section.CategoryID),
			}
			// Due to timestamp format needs to be convert,
			// we cannot use json Marshal + Unmarshal to converts format.
			outSection.CreatedAt, _ = ptypes.TimestampProto(section.CreatedAt)
			outSection.UpdatedAt, _ = ptypes.TimestampProto(section.UpdatedAt)

			sectionsOut.Sections = append(sectionsOut.Sections, outSection)
		}

		return sectionsOut, nil
	case *protobuf.GetSectionsRequest_CategoryId:
		perPage, page := inout.ProcessPage(in.PerPage, in.Page)
		categoryID, err := strconv.Atoi(in.GetCategoryId())
		if err != nil {
			return nil, errs.NewErr(
				errs.InvalidAttributeErrorCode,
				errors.Wrapf(err, "grpc: [GetSectionsByCategoryID] failed"),
			)
		}

		defer s.examiner.CheckSections(ctx, inout.GRPCCountryCodeMap[in.CountryCode], inout.GRPCLocaleMap[in.Locale])

		sections, total, err := s.service.GetSectionsByCategoryID(ctx,
			&models.GetSectionsParams{
				CategoryID:  categoryID,
				Locale:      inout.GRPCLocaleMap[in.Locale],
				CountryCode: inout.GRPCCountryCodeMap[in.CountryCode],
				PerPage:     int(perPage),
				Page:        int(page),
				SortBy:      inout.GRPCSortByMap[in.SortBy],
				SortOrder:   inout.GRPCSortOrderMap[in.SortOrder],
			})
		if err != nil {
			if err == models.ErrNotFound {
				return nil, errs.NewErr(
					errs.RecordNotFoundErrorCode,
					errors.Wrapf(err, "grpc: [GetSectionsByCategoryID] not found"),
				)
			}
			return nil, errs.NewErr(
				errs.ServerInternalErrorCode,
				errors.Wrapf(err, "grpc: [GetSectionsByCategoryID] failed"),
			)
		}

		sectionsOut := &protobuf.GetSectionsResponse{
			PageInfo: &protobuf.PageInfo{
				Page:      int32(math.Round(float64(page)/float64(perPage))) + 1,
				PerPage:   int32(perPage),
				PageCount: int32(math.Ceil(float64(total) / float64(perPage))),
				Count:     int32(total),
			},
			Sections: make([]*protobuf.Section, 0),
		}
		for _, section := range sections {
			outSection := &protobuf.Section{
				Id:           strconv.Itoa(section.ID),
				Position:     int32(section.Position),
				SourceLocale: section.SourceLocale,
				Outdated:     section.Outdated,
				CountryCode:  section.CountryCode,
				Url:          section.URL,
				HtmlUrl:      section.HTMLURL,
				Name:         section.Name,
				Description:  section.Description,
				Locale:       section.Locale,
				CategoryId:   strconv.Itoa(section.CategoryID),
			}
			// Due to timestamp format needs to be convert,
			// we cannot use json Marshal + Unmarshal to converts format.
			outSection.CreatedAt, _ = ptypes.TimestampProto(section.CreatedAt)
			outSection.UpdatedAt, _ = ptypes.TimestampProto(section.UpdatedAt)

			sectionsOut.Sections = append(sectionsOut.Sections, outSection)
		}

		return sectionsOut, nil
	}

	return nil, fmt.Errorf("[GetSections] invalid type: %v", in.Id)
}

func (s *server) GetSection(ctx context.Context, in *protobuf.GetSectionRequest) (*protobuf.GetSectionResponse, error) {
	switch in.Id.(type) {
	case *protobuf.GetSectionRequest_SectionId:
		sectionID, err := strconv.Atoi(in.GetSectionId())
		if err != nil {
			return nil, errs.NewErr(
				errs.InvalidAttributeErrorCode,
				errors.Wrapf(err, "grpc: [GetSectionBySectionID] failed"),
			)
		}

		defer s.examiner.CheckSections(ctx, inout.GRPCCountryCodeMap[in.CountryCode], inout.GRPCLocaleMap[in.Locale])

		section, err := s.service.GetSectionBySectionID(ctx,
			sectionID,
			inout.GRPCLocaleMap[in.Locale],
			inout.GRPCCountryCodeMap[in.CountryCode],
		)
		if err != nil {
			if err == models.ErrNotFound {
				return nil, errs.NewErr(
					errs.RecordNotFoundErrorCode,
					errors.Wrapf(err, "grpc: [GetSectionBySectionID] not found"),
				)
			}
			return nil, errs.NewErr(
				errs.ServerInternalErrorCode,
				errors.Wrapf(err, "grpc: [GetSectionBySectionID] failed"),
			)
		}

		out := &protobuf.GetSectionResponse{
			Section: &protobuf.Section{
				Id:           strconv.Itoa(section.ID),
				Position:     int32(section.Position),
				SourceLocale: section.SourceLocale,
				Outdated:     section.Outdated,
				CountryCode:  section.CountryCode,
				Url:          section.URL,
				HtmlUrl:      section.HTMLURL,
				Name:         section.Name,
				Description:  section.Description,
				Locale:       section.Locale,
				CategoryId:   strconv.Itoa(section.CategoryID),
			},
		}
		// Due to timestamp format needs to be convert,
		// we cannot use json Marshal + Unmarshal to converts format.
		out.Section.CreatedAt, _ = ptypes.TimestampProto(section.CreatedAt)
		out.Section.UpdatedAt, _ = ptypes.TimestampProto(section.UpdatedAt)

		return out, nil
	case *protobuf.GetSectionRequest_ArticleId:
		articleID, err := strconv.Atoi(in.GetArticleId())
		if err != nil {
			return nil, errs.NewErr(
				errs.InvalidAttributeErrorCode,
				errors.Wrapf(err, "grpc: [GetSectionByArticleID] failed"),
			)
		}

		defer s.examiner.CheckSections(ctx, inout.GRPCCountryCodeMap[in.CountryCode], inout.GRPCLocaleMap[in.Locale])

		section, err := s.service.GetSectionByArticleID(ctx,
			articleID,
			inout.GRPCLocaleMap[in.Locale],
			inout.GRPCCountryCodeMap[in.CountryCode],
		)
		if err != nil {
			if err == models.ErrNotFound {
				return nil, errs.NewErr(
					errs.RecordNotFoundErrorCode,
					errors.Wrapf(err, "grpc: [GetSectionByArticleID] not found"),
				)
			}
			return nil, errs.NewErr(
				errs.ServerInternalErrorCode,
				errors.Wrapf(err, "grpc: [GetSectionByArticleID] failed"),
			)
		}

		out := &protobuf.GetSectionResponse{
			Section: &protobuf.Section{
				Id:           strconv.Itoa(section.ID),
				Position:     int32(section.Position),
				SourceLocale: section.SourceLocale,
				Outdated:     section.Outdated,
				CountryCode:  section.CountryCode,
				Url:          section.URL,
				HtmlUrl:      section.HTMLURL,
				Name:         section.Name,
				Description:  section.Description,
				Locale:       section.Locale,
				CategoryId:   strconv.Itoa(section.CategoryID),
			},
		}
		// Due to timestamp format needs to be convert,
		// we cannot use json Marshal + Unmarshal to converts format.
		out.Section.CreatedAt, _ = ptypes.TimestampProto(section.CreatedAt)
		out.Section.UpdatedAt, _ = ptypes.TimestampProto(section.UpdatedAt)

		return out, nil
	}

	return nil, fmt.Errorf("[GetSection] invalid type: %v", in.Id)
}

func (s *server) GetArticles(ctx context.Context, in *protobuf.GetArticlesRequest) (*protobuf.GetArticlesResponse, error) {
	switch in.Id.(type) {
	case *protobuf.GetArticlesRequest_All:
		perPage, page := inout.ProcessPage(in.PerPage, in.Page)

		defer s.examiner.CheckArticles(ctx, inout.GRPCCountryCodeMap[in.CountryCode], inout.GRPCLocaleMap[in.Locale])

		articles, total, err := s.service.GetArticles(ctx,
			&models.GetArticlesParams{
				Locale:      inout.GRPCLocaleMap[in.Locale],
				CountryCode: inout.GRPCCountryCodeMap[in.CountryCode],
				PerPage:     int(perPage),
				Page:        int(page),
				SortBy:      inout.GRPCSortByMap[in.SortBy],
				SortOrder:   inout.GRPCSortOrderMap[in.SortOrder],
			})
		if err != nil {
			if err == models.ErrNotFound {
				return nil, errs.NewErr(
					errs.RecordNotFoundErrorCode,
					errors.Wrapf(err, "grpc: [GetArticles] not found"),
				)
			}
			return nil, errs.NewErr(
				errs.ServerInternalErrorCode,
				errors.Wrapf(err, "grpc: [GetArticles] failed"),
			)
		}

		out := &protobuf.GetArticlesResponse{
			PageInfo: &protobuf.PageInfo{
				Page:      int32(math.Round(float64(page)/float64(perPage))) + 1,
				PerPage:   int32(perPage),
				PageCount: int32(math.Ceil(float64(total) / float64(perPage))),
				Count:     int32(total),
			},
			Articles: make([]*protobuf.Article, 0),
		}
		for _, article := range articles {
			outArticle := &protobuf.Article{
				Id:              strconv.Itoa(article.ID),
				AuthorId:        strconv.Itoa(article.AuthorID),
				CommentsDisable: article.CommentsDisable,
				Draft:           article.Draft,
				Promoted:        article.Promoted,
				Position:        int32(article.Position),
				VoteSum:         int32(article.VoteSum),
				VoteCount:       int32(article.VoteCount),
				SourceLocale:    article.SourceLocale,
				Outdated:        article.Outdated,
				OutdatedLocales: article.OutdatedLocales,
				LabelNames:      article.LabelNames,
				CountryCode:     article.CountryCode,
				Url:             article.URL,
				HtmlUrl:         article.HTMLURL,
				Name:            article.Name,
				Title:           article.Title,
				Body:            article.Body,
				Locale:          article.Locale,
				SectionId:       strconv.Itoa(article.SectionID),
			}
			// Due to timestamp format needs to be convert,
			// we cannot use json Marshal + Unmarshal to converts format.
			outArticle.CreatedAt, _ = ptypes.TimestampProto(article.CreatedAt)
			outArticle.UpdatedAt, _ = ptypes.TimestampProto(article.UpdatedAt)
			outArticle.EditedAt, _ = ptypes.TimestampProto(article.EditedAt)

			out.Articles = append(out.Articles, outArticle)
		}

		return out, nil
	case *protobuf.GetArticlesRequest_CategoryId:
		perPage, page := inout.ProcessPage(in.PerPage, in.Page)
		categoryID, err := strconv.Atoi(in.GetCategoryId())
		if err != nil {
			return nil, errs.NewErr(
				errs.InvalidAttributeErrorCode,
				errors.Wrapf(err, "grpc: [GetArticlesByCategoryID] failed"),
			)
		}

		defer s.examiner.CheckArticles(ctx, inout.GRPCCountryCodeMap[in.CountryCode], inout.GRPCLocaleMap[in.Locale])

		articles, total, err := s.service.GetArticlesByCategoryID(ctx,
			&models.GetArticlesParams{
				CategoryID:  categoryID,
				Locale:      inout.GRPCLocaleMap[in.Locale],
				CountryCode: inout.GRPCCountryCodeMap[in.CountryCode],
				PerPage:     int(perPage),
				Page:        int(page),
				SortBy:      inout.GRPCSortByMap[in.SortBy],
				SortOrder:   inout.GRPCSortOrderMap[in.SortOrder],
			}, in.LabelNames)
		if err != nil {
			if err == models.ErrNotFound {
				return nil, errs.NewErr(
					errs.RecordNotFoundErrorCode,
					errors.Wrapf(err, "grpc: [GetArticlesByCategoryID] not found"),
				)
			}
			return nil, errs.NewErr(
				errs.ServerInternalErrorCode,
				errors.Wrapf(err, "grpc: [GetArticlesByCategoryID] failed"),
			)
		}

		out := &protobuf.GetArticlesResponse{
			PageInfo: &protobuf.PageInfo{
				Page:      int32(math.Round(float64(page)/float64(perPage))) + 1,
				PerPage:   int32(perPage),
				PageCount: int32(math.Ceil(float64(total) / float64(perPage))),
				Count:     int32(total),
			},
			Articles: make([]*protobuf.Article, 0),
		}
		for _, article := range articles {
			outArticle := &protobuf.Article{
				Id:              strconv.Itoa(article.ID),
				AuthorId:        strconv.Itoa(article.AuthorID),
				CommentsDisable: article.CommentsDisable,
				Draft:           article.Draft,
				Promoted:        article.Promoted,
				Position:        int32(article.Position),
				VoteSum:         int32(article.VoteSum),
				VoteCount:       int32(article.VoteCount),
				SourceLocale:    article.SourceLocale,
				Outdated:        article.Outdated,
				OutdatedLocales: article.OutdatedLocales,
				LabelNames:      article.LabelNames,
				CountryCode:     article.CountryCode,
				Url:             article.URL,
				HtmlUrl:         article.HTMLURL,
				Name:            article.Name,
				Title:           article.Title,
				Body:            article.Body,
				Locale:          article.Locale,
				SectionId:       strconv.Itoa(article.SectionID),
			}
			// Due to timestamp format needs to be convert,
			// we cannot use json Marshal + Unmarshal to converts format.
			outArticle.CreatedAt, _ = ptypes.TimestampProto(article.CreatedAt)
			outArticle.UpdatedAt, _ = ptypes.TimestampProto(article.UpdatedAt)
			outArticle.EditedAt, _ = ptypes.TimestampProto(article.EditedAt)

			out.Articles = append(out.Articles, outArticle)
		}

		return out, nil
	case *protobuf.GetArticlesRequest_SectionId:
		perPage, page := inout.ProcessPage(in.PerPage, in.Page)
		sectionID, err := strconv.Atoi(in.GetSectionId())
		if err != nil {
			return nil, errs.NewErr(
				errs.InvalidAttributeErrorCode,
				errors.Wrapf(err, "grpc: [GetArticlesBySectionID] failed"),
			)
		}

		defer s.examiner.CheckArticles(ctx, inout.GRPCCountryCodeMap[in.CountryCode], inout.GRPCLocaleMap[in.Locale])

		articles, total, err := s.service.GetArticlesBySectionID(ctx,
			&models.GetArticlesParams{
				SectionID:   sectionID,
				Locale:      inout.GRPCLocaleMap[in.Locale],
				CountryCode: inout.GRPCCountryCodeMap[in.CountryCode],
				PerPage:     int(perPage),
				Page:        int(page),
				SortBy:      inout.GRPCSortByMap[in.SortBy],
				SortOrder:   inout.GRPCSortOrderMap[in.SortOrder],
			})
		if err != nil {
			if err == models.ErrNotFound {
				return nil, errs.NewErr(
					errs.RecordNotFoundErrorCode,
					errors.Wrapf(err, "grpc: [GetArticlesBySectionID] not found"),
				)
			}
			return nil, errs.NewErr(
				errs.ServerInternalErrorCode,
				errors.Wrapf(err, "grpc: [GetArticlesBySectionID] failed"),
			)
		}

		out := &protobuf.GetArticlesResponse{
			PageInfo: &protobuf.PageInfo{
				Page:      int32(math.Round(float64(page)/float64(perPage))) + 1,
				PerPage:   int32(perPage),
				PageCount: int32(math.Ceil(float64(total) / float64(perPage))),
				Count:     int32(total),
			},
			Articles: make([]*protobuf.Article, 0),
		}
		for _, article := range articles {
			outArticle := &protobuf.Article{
				Id:              strconv.Itoa(article.ID),
				AuthorId:        strconv.Itoa(article.AuthorID),
				CommentsDisable: article.CommentsDisable,
				Draft:           article.Draft,
				Promoted:        article.Promoted,
				Position:        int32(article.Position),
				VoteSum:         int32(article.VoteSum),
				VoteCount:       int32(article.VoteCount),
				SourceLocale:    article.SourceLocale,
				Outdated:        article.Outdated,
				OutdatedLocales: article.OutdatedLocales,
				LabelNames:      article.LabelNames,
				CountryCode:     article.CountryCode,
				Url:             article.URL,
				HtmlUrl:         article.HTMLURL,
				Name:            article.Name,
				Title:           article.Title,
				Body:            article.Body,
				Locale:          article.Locale,
				SectionId:       strconv.Itoa(article.SectionID),
			}
			// Due to timestamp format needs to be convert,
			// we cannot use json Marshal + Unmarshal to converts format.
			outArticle.CreatedAt, _ = ptypes.TimestampProto(article.CreatedAt)
			outArticle.UpdatedAt, _ = ptypes.TimestampProto(article.UpdatedAt)
			outArticle.EditedAt, _ = ptypes.TimestampProto(article.EditedAt)

			out.Articles = append(out.Articles, outArticle)
		}

		return out, nil
	}

	return nil, fmt.Errorf("[GetArticles] invalid type: %v", in.Id)
}

func (s *server) GetTopArticles(ctx context.Context, in *protobuf.GetTopArticlesRequest) (*protobuf.GetTopArticlesResponse, error) {
	articles, err := s.service.GetTopNArticles(ctx,
		uint64(in.TopN),
		inout.GRPCLocaleMap[in.Locale],
		inout.GRPCCountryCodeMap[in.CountryCode],
	)
	if err != nil {
		if err == models.ErrNotFound {
			return nil, errs.NewErr(
				errs.RecordNotFoundErrorCode,
				errors.Wrapf(err, "grpc: [GetTopNArticles] not found"),
			)
		}
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Wrapf(err, "grpc: [GetTopNArticles] failed"),
		)
	}

	out := &protobuf.GetTopArticlesResponse{
		Articles: make([]*protobuf.Article, 0),
	}
	for _, article := range articles {
		outArticle := &protobuf.Article{
			Id:              strconv.Itoa(article.ID),
			AuthorId:        strconv.Itoa(article.AuthorID),
			CommentsDisable: article.CommentsDisable,
			Draft:           article.Draft,
			Promoted:        article.Promoted,
			Position:        int32(article.Position),
			VoteSum:         int32(article.VoteSum),
			VoteCount:       int32(article.VoteCount),
			SourceLocale:    article.SourceLocale,
			Outdated:        article.Outdated,
			OutdatedLocales: article.OutdatedLocales,
			LabelNames:      article.LabelNames,
			CountryCode:     article.CountryCode,
			Url:             article.URL,
			HtmlUrl:         article.HTMLURL,
			Name:            article.Name,
			Title:           article.Title,
			Body:            article.Body,
			Locale:          article.Locale,
			SectionId:       strconv.Itoa(article.SectionID),
		}
		// Due to timestamp format needs to be convert,
		// we cannot use json Marshal + Unmarshal to converts format.
		outArticle.CreatedAt, _ = ptypes.TimestampProto(article.CreatedAt)
		outArticle.UpdatedAt, _ = ptypes.TimestampProto(article.UpdatedAt)
		outArticle.EditedAt, _ = ptypes.TimestampProto(article.EditedAt)

		out.Articles = append(out.Articles, outArticle)
	}

	return out, nil
}

func (s *server) GetArticle(ctx context.Context, in *protobuf.GetArticleRequest) (*protobuf.GetArticleResponse, error) {
	articleID, err := strconv.Atoi(in.ArticleId)
	if err != nil {
		return nil, errs.NewErr(
			errs.InvalidAttributeErrorCode,
			errors.Wrapf(err, "grpc: [GetArticleByArticleID] failed"),
		)
	}

	defer s.examiner.CheckArticles(ctx, inout.GRPCCountryCodeMap[in.CountryCode], inout.GRPCLocaleMap[in.Locale])
	defer s.service.PlusOneArticleClickCounter(ctx, articleID, inout.GRPCLocaleMap[in.Locale], inout.GRPCCountryCodeMap[in.CountryCode])

	article, err := s.service.GetArticleByArticleID(ctx,
		articleID,
		inout.GRPCLocaleMap[in.Locale],
		inout.GRPCCountryCodeMap[in.CountryCode],
	)
	if err != nil {
		if err == models.ErrNotFound {
			return nil, errs.NewErr(
				errs.RecordNotFoundErrorCode,
				errors.Wrapf(err, "grpc: [GetArticleByArticleID] not found"),
			)
		}
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Wrapf(err, "grpc: [GetArticleByArticleID] failed"),
		)
	}

	out := &protobuf.GetArticleResponse{
		Article: &protobuf.Article{
			Id:              strconv.Itoa(article.ID),
			AuthorId:        strconv.Itoa(article.AuthorID),
			CommentsDisable: article.CommentsDisable,
			Draft:           article.Draft,
			Promoted:        article.Promoted,
			Position:        int32(article.Position),
			VoteSum:         int32(article.VoteSum),
			VoteCount:       int32(article.VoteCount),
			SourceLocale:    article.SourceLocale,
			Outdated:        article.Outdated,
			OutdatedLocales: article.OutdatedLocales,
			LabelNames:      article.LabelNames,
			CountryCode:     article.CountryCode,
			Url:             article.URL,
			HtmlUrl:         article.HTMLURL,
			Name:            article.Name,
			Title:           article.Title,
			Body:            article.Body,
			Locale:          article.Locale,
			SectionId:       strconv.Itoa(article.SectionID),
		},
	}
	// Due to timestamp format needs to be convert,
	// we cannot use json Marshal + Unmarshal to converts format.
	out.Article.CreatedAt, _ = ptypes.TimestampProto(article.CreatedAt)
	out.Article.UpdatedAt, _ = ptypes.TimestampProto(article.UpdatedAt)
	out.Article.EditedAt, _ = ptypes.TimestampProto(article.EditedAt)

	return out, nil
}

func (s *server) GetTicketForm(ctx context.Context, in *protobuf.GetTicketFormRequest) (*protobuf.GetTicketFormResponse, error) {
	formID, err := strconv.Atoi(in.FormId)
	if err != nil {
		return nil, errs.NewErr(
			errs.InvalidAttributeErrorCode,
			errors.Wrapf(err, "grpc: [GetTicketFormGraphQL] failed"),
		)
	}

	ticketForm, err := s.service.GetTicketFormGraphQL(ctx, formID)
	if err != nil {
		if err == models.ErrNotFound {
			return nil, errs.NewErr(
				errs.RecordNotFoundErrorCode,
				errors.Wrapf(err, "grpc: [GetTicketFormGraphQL] not found"),
			)
		}
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Wrapf(err, "grpc: [GetTicketFormGraphQL] failed"),
		)
	}

	out := &protobuf.GetTicketFormResponse{
		Id:                 strconv.Itoa(ticketForm.ID),
		Url:                ticketForm.URL,
		Name:               ticketForm.Name,
		RawName:            ticketForm.RawName,
		DisplayName:        ticketForm.DisplayName,
		RawDisplayName:     ticketForm.RawDisplayName,
		EndUserVisible:     ticketForm.EndUserVisible,
		Position:           int32(ticketForm.Position),
		Active:             ticketForm.Active,
		InAllBrands:        ticketForm.InAllBrands,
		RestrictedBrandIds: make([]int32, 0),
	}
	for _, id := range ticketForm.RestrictedBrandIDs {
		out.RestrictedBrandIds = append(out.RestrictedBrandIds, int32(id))
	}
	// Due to timestamp format needs to be convert,
	// we cannot use json Marshal + Unmarshal to converts format.
	out.CreatedAt, _ = ptypes.TimestampProto(ticketForm.CreatedAt)
	out.UpdatedAt, _ = ptypes.TimestampProto(ticketForm.UpdatedAt)

	return out, nil
}

func (s *server) GetTicketFields(ctx context.Context, in *protobuf.GetTicketFieldsRequest) (*protobuf.GetTicketFieldsResponse, error) {
	formID, err := strconv.Atoi(in.FormId)
	if err != nil {
		return nil, errs.NewErr(
			errs.InvalidAttributeErrorCode,
			errors.Wrapf(err, "grpc: [GetTicketFieldByFormID] failed"),
		)
	}

	ticketFields, err := s.service.GetTicketFieldByFormID(ctx, formID, inout.GRPCLocaleMap[in.Locale])
	if err != nil {
		if err == models.ErrNotFound {
			return nil, errs.NewErr(
				errs.RecordNotFoundErrorCode,
				errors.Wrapf(err, "grpc: [GetTicketFieldByFormID] not found"),
			)
		}
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Wrapf(err, "grpc: [GetTicketFieldByFormID] failed"),
		)
	}

	out := &protobuf.GetTicketFieldsResponse{
		TicketFields: make([]*protobuf.TicketField, 0),
	}
	for _, ticket := range ticketFields {
		outTicketField := &protobuf.TicketField{
			Id:                  strconv.Itoa(ticket.ID),
			Url:                 ticket.URL,
			Type:                ticket.Type,
			Title:               ticket.Title,
			RawTitle:            ticket.RawTitle,
			Description:         ticket.Description,
			RawDescription:      ticket.RawDescription,
			Position:            int32(ticket.Position),
			Active:              ticket.Active,
			Required:            ticket.Required,
			CollapsedForAgents:  ticket.CollapsedForAgents,
			RegexpForValidation: ticket.RegexpForValidation,
			TitleInPortal:       ticket.TitleInPortal,
			RawTitleInPortal:    ticket.RawTitleInPortal,
			VisibleInPortal:     ticket.VisibleInPortal,
			EditableInPortal:    ticket.EditableInPortal,
			RequiredInPortal:    ticket.RequiredInPortal,
			Tag:                 ticket.Tag,
			Removable:           ticket.Removable,
			CustomFieldOptions:  make([]*protobuf.CustomFieldOption, 0),
			SystemFieldOptions:  make([]*protobuf.SystemFieldOption, 0),
		}
		// Due to timestamp format needs to be convert,
		// we cannot use json Marshal + Unmarshal to converts format.
		outTicketField.CreatedAt, _ = ptypes.TimestampProto(ticket.CreatedAt)
		outTicketField.UpdatedAt, _ = ptypes.TimestampProto(ticket.UpdatedAt)

		customOption, err := s.service.GetTicketFieldCustomFieldOption(ctx, ticket.ID)
		if err != nil {
			if err == models.ErrNotFound {
				return nil, errs.NewErr(
					errs.RecordNotFoundErrorCode,
					errors.Wrapf(err, "grpc: [GetTicketFieldCustomFieldOption] not found"),
				)
			}
			return nil, errs.NewErr(
				errs.ServerInternalErrorCode,
				errors.Wrapf(err, "grpc: [GetTicketFieldCustomFieldOption] failed"),
			)
		}
		for _, option := range customOption {
			outTicketField.CustomFieldOptions = append(outTicketField.CustomFieldOptions, &protobuf.CustomFieldOption{
				Id:      strconv.Itoa(option.ID),
				Name:    option.Name,
				RawName: option.RawName,
				Value:   option.Value,
			})
		}

		systemOption, err := s.service.GetTicketFieldSystemFieldOption(ctx, ticket.ID)
		if err != nil {
			if err == models.ErrNotFound {
				return nil, errs.NewErr(
					errs.RecordNotFoundErrorCode,
					errors.Wrapf(err, "grpc: [GetTicketFieldSystemFieldOption] not found"),
				)
			}
			return nil, errs.NewErr(
				errs.ServerInternalErrorCode,
				errors.Wrapf(err, "grpc: [GetTicketFieldSystemFieldOption] failed"),
			)
		}
		for _, option := range systemOption {
			outTicketField.SystemFieldOptions = append(outTicketField.SystemFieldOptions, &protobuf.SystemFieldOption{
				Name:  option.Name,
				Value: option.Value,
			})
		}

		out.TicketFields = append(out.TicketFields, outTicketField)
	}

	return out, nil
}

func (s *server) GetSearchTitleArticles(ctx context.Context, in *protobuf.GetSearchTitleArticlesRequest) (*protobuf.GetSearchTitleArticlesResponse, error) {
	zendeskInstantSearch, err := s.zend.InstantSearch(ctx,
		in.Query,
		inout.GRPCCountryCodeMap[in.CountryCode],
		inout.GRPCLocaleMap[in.Locale],
	)
	if err != nil {
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Wrapf(err, "grpc: [InstantSearch] failed"),
		)
	}

	out := &protobuf.GetSearchTitleArticlesResponse{
		Articles: make([]*protobuf.SearchTitleArticle, 0),
	}
	for _, result := range zendeskInstantSearch.Results {
		out.Articles = append(out.Articles, &protobuf.SearchTitleArticle{
			Title:         result.Title,
			CategoryTitle: result.CategoryTitle,
			Url:           result.URL,
		})
	}

	return out, nil
}

func (s *server) GetSearchBodyArticles(ctx context.Context, in *protobuf.GetSearchBodyArticlesRequest) (*protobuf.GetSearchBodyArticlesResponse, error) {
	perPage, page := inout.ProcessPage(in.PerPage, in.Page)
	categoryIDs, err := s.service.GetCategoriesID(ctx, inout.GRPCCountryCodeMap[in.CountryCode])
	if err != nil {
		if err == models.ErrNotFound {
			return nil, errs.NewErr(
				errs.RecordNotFoundErrorCode,
				errors.Wrapf(err, "grpc: [GetCategoriesID] not found"),
			)
		}
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Wrapf(err, "grpc: [GetCategoriesID] failed"),
		)
	}

	zendeskSearch, err := s.zend.Search(ctx,
		categoryIDs, in.Query,
		inout.GRPCCountryCodeMap[in.CountryCode], inout.GRPCLocaleMap[in.Locale],
		&zendesk.Pagination{
			PerPage:   int(perPage),
			Page:      int(math.Round(float64(page)/float64(perPage))) + 1,
			SortOrder: inout.GRPCSortOrderMap[in.SortOrder],
		})
	if err != nil {
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Wrapf(err, "grpc: [Search] failed"),
		)
	}

	out := &protobuf.GetSearchBodyArticlesResponse{
		PageInfo: &protobuf.PageInfo{
			Page:      int32(zendeskSearch.Page),
			PerPage:   int32(zendeskSearch.PerPage),
			PageCount: int32(zendeskSearch.PageCount),
			Count:     int32(zendeskSearch.Count),
		},
		Articles: make([]*protobuf.SearchBodyArticle, 0),
	}

	for _, article := range zendeskSearch.Articles {
		category, err := s.service.GetCategoryByArticleID(ctx, article.ID, article.Locale)
		if err != nil {
			if err == models.ErrNotFound {
				// if article not found in locale db, ignore it.
				continue
			}
			return nil, errs.NewErr(
				errs.RecordNotFoundErrorCode,
				errors.Wrapf(err, "grpc: [GetCategoryByArticleID] failed"),
			)
		}

		outArticle := &protobuf.SearchBodyArticle{
			Id:              strconv.Itoa(article.ID),
			AuthorId:        strconv.Itoa(article.AuthorID),
			CommentsDisable: article.CommentsDisable,
			Draft:           article.Draft,
			Promoted:        article.Promoted,
			Position:        int32(article.Position),
			VoteSum:         int32(article.VoteSum),
			VoteCount:       int32(article.VoteCount),
			SourceLocale:    article.SourceLocale,
			Outdated:        article.Outdated,
			OutdatedLocales: article.OutdatedLocales,
			LabelNames:      article.LabelNames,
			CountryCode:     inout.GRPCCountryCodeMap[in.CountryCode],
			Url:             article.URL,
			HtmlUrl:         article.HTMLURL,
			Name:            article.Name,
			Title:           article.Title,
			Body:            article.Body,
			Locale:          article.Locale,
			Snippet:         article.Snippet,
			SectionId:       strconv.Itoa(article.SectionID),
			CategoryId:      strconv.Itoa(category.ID),
			CategoryName:    category.Name,
		}
		// Due to timestamp format needs to be convert,
		// we cannot use json Marshal + Unmarshal to converts format.
		outArticle.CreatedAt, _ = ptypes.TimestampProto(article.CreatedAt)
		outArticle.UpdatedAt, _ = ptypes.TimestampProto(article.UpdatedAt)
		outArticle.EditedAt, _ = ptypes.TimestampProto(article.EditedAt)

		out.Articles = append(out.Articles, outArticle)
	}

	return out, nil
}

func (s *server) GetStatus(ctx context.Context, in *protobuf.GetStatusRequest) (*protobuf.GetStatusResponse, error) {
	serverTime, _ := ptypes.TimestampProto(time.Now().UTC())

	return &protobuf.GetStatusResponse{
		GoVersion:  runtime.Version(),
		AppVersion: config.Version,
		ServerTime: serverTime,
	}, nil
}

func (s *server) SetCreateRequest(ctx context.Context, in *protobuf.SetCreateRequestRequest) (*protobuf.SetCreateRequestResponse, error) {
	if in.Data == nil {
		return nil, errs.NewErr(
			errs.InvalidAttributeErrorCode,
			errors.New("grpc: [SetCreateRequest] failed"),
		)
	}
	if in.Data.Request == nil {
		return nil, errs.NewErr(
			errs.InvalidAttributeErrorCode,
			errors.New("grpc: [SetCreateRequest] failed"),
		)
	}

	request := &inout.MutationRequestsIn{
		CountryCode: inout.GRPCCountryCodeMap[in.CountryCode],
		Data: inout.CreateRequestData{
			Request: inout.CreateRequestDataRequest{
				Subject: in.Data.Request.Subject,
			},
		},
	}
	if in.Data.Request.Comment != nil {
		request.Data.Request.Comment.Body = in.Data.Request.Comment.Body
	}
	if in.Data.Request.Requester != nil {
		request.Data.Request.Requester.Name = in.Data.Request.Requester.Name
		request.Data.Request.Requester.Email = in.Data.Request.Requester.Email
	}
	if in.Data.Request.TicketFormId != "" {
		request.Data.Request.TicketFormID = &in.Data.Request.TicketFormId
	}
	if len(in.Data.Request.CustomFields) > 0 {
		customFields := make([]inout.CreateRequestDataRequestCustomField, 0)
		for _, field := range in.Data.Request.CustomFields {
			customFields = append(customFields,
				inout.CreateRequestDataRequestCustomField{
					ID:    field.Id,
					Value: field.Value,
				},
			)
		}
		request.Data.Request.CustomFields = &customFields
	}

	if err := s.zend.CreateRequest(ctx, request.CountryCode, request.Data); err != nil {
		return &protobuf.SetCreateRequestResponse{
				Status: http.StatusText(http.StatusBadRequest),
			}, errs.NewErr(
				errs.InvalidAttributeErrorCode,
				errors.Wrapf(err, "grpc: [CreateRequest] failed"),
			)
	}

	return &protobuf.SetCreateRequestResponse{
		Status: http.StatusText(http.StatusCreated),
	}, nil
}

func (s *server) SetVoteArticle(ctx context.Context, in *protobuf.SetVoteArticleRequest) (*protobuf.SetVoteArticleResponse, error) {
	articleID, err := strconv.Atoi(in.ArticleId)
	if err != nil {
		return nil, errs.NewErr(
			errs.InvalidAttributeErrorCode,
			errors.Wrapf(err, "grpc: [CreateVote] failed"),
		)
	}

	voteResult, err := s.zend.CreateVote(ctx,
		articleID,
		inout.GRPCVoteMap[in.Vote],
		inout.GRPCCountryCodeMap[in.CountryCode],
		inout.GRPCLocaleMap[in.Locale],
	)
	if err != nil {
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Wrapf(err, "grpc: [CreateVote] failed"),
		)
	}

	defer s.examiner.SyncArticle(ctx, articleID, inout.GRPCCountryCodeMap[in.CountryCode], inout.GRPCLocaleMap[in.Locale])

	article, err := s.service.GetArticleByArticleID(ctx,
		articleID,
		inout.GRPCLocaleMap[in.Locale],
		inout.GRPCCountryCodeMap[in.CountryCode],
	)
	if err != nil {
		if err == models.ErrNotFound {
			return nil, errs.NewErr(
				errs.RecordNotFoundErrorCode,
				errors.Wrapf(err, "grpc: [GetArticleByArticleID] not found"),
			)
		}
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Wrapf(err, "grpc: [GetArticleByArticleID] failed"),
		)
	}

	outArticle := &protobuf.Article{
		Id:              strconv.Itoa(article.ID),
		AuthorId:        strconv.Itoa(article.AuthorID),
		CommentsDisable: article.CommentsDisable,
		Draft:           article.Draft,
		Promoted:        article.Promoted,
		Position:        int32(article.Position),
		VoteSum:         int32(voteResult.VoteSum),
		VoteCount:       int32(voteResult.VoteCount),
		SourceLocale:    article.SourceLocale,
		Outdated:        article.Outdated,
		OutdatedLocales: article.OutdatedLocales,
		LabelNames:      article.LabelNames,
		CountryCode:     article.CountryCode,
		Url:             article.URL,
		HtmlUrl:         article.HTMLURL,
		Name:            article.Name,
		Title:           article.Title,
		Body:            article.Body,
		Locale:          article.Locale,
		SectionId:       strconv.Itoa(article.SectionID),
	}
	// Due to timestamp format needs to be convert,
	// we cannot use json Marshal + Unmarshal to converts format.
	outArticle.CreatedAt, _ = ptypes.TimestampProto(article.CreatedAt)
	outArticle.UpdatedAt, _ = ptypes.TimestampProto(article.UpdatedAt)
	outArticle.EditedAt, _ = ptypes.TimestampProto(article.EditedAt)

	return &protobuf.SetVoteArticleResponse{Article: outArticle}, nil
}

func (s *server) SetForceSync(ctx context.Context, in *protobuf.SetForceSyncRequest) (*protobuf.SetForceSyncResponse, error) {
	if s.conf.HTTP.BasicAuthUser != in.Username || s.conf.HTTP.BasicAuthPwd != in.Password {
		return nil, errs.NewErr(
			errs.UnauthorizedErrCode,
			errors.New("grpc: [SetForceSync] failed"),
		)
	}

	go func() {
		// Loop by countryCode.
		for countryCode, locales := range inout.SupportCountryLocaleMap {
			for _, locale := range locales {
				s.logger.Info().Msgf("force sync categories, country code %s locales %s", countryCode, locale)
				if err := s.examiner.ForceSyncCategories(ctx, countryCode, locale); err != nil {
					s.logger.Error().Err(err).Msgf("sync categories, country code %s locale %s failed", countryCode, locale)
				}

				s.logger.Info().Msgf("force sync sections, country code %s locales %s", countryCode, locale)
				if err := s.examiner.ForceSyncSections(ctx, countryCode, locale); err != nil {
					s.logger.Error().Err(err).Msgf("sync sections, country code %s locale %s failed", countryCode, locale)
				}

				s.logger.Info().Msgf("force sync articles, country code %s locales %s", countryCode, locale)
				if err := s.examiner.ForceSyncArticles(ctx, countryCode, locale); err != nil {
					s.logger.Error().Err(err).Msgf("sync articles, country code %s locale %s failed", countryCode, locale)
				}
			}
		}

		s.logger.Info().Msgf("force sync ticket forms")
		if err := s.examiner.ForceSyncTicketForms(ctx); err != nil {
			s.logger.Error().Err(err).Msgf("sync ticket forms failed")
		}
	}()

	return &protobuf.SetForceSyncResponse{Status: inout.SuccessForceSync}, nil
}
