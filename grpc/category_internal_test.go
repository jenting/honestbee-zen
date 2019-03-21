package grpc

import (
	"context"
	"testing"

	"github.com/go-test/deep"

	"github.com/honestbee/Zen/models"
	"github.com/honestbee/Zen/protobuf"
)

func TestGetCategories(t *testing.T) {
	s := initServer()

	testCases := [...]struct {
		description string
		input       *protobuf.GetCategoriesRequest
		expectErr   bool
		expect      *protobuf.GetCategoriesResponse
	}{
		{
			description: "testing normal case",
			input: &protobuf.GetCategoriesRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				PerPage:     3,
				Page:        0,
				SortBy:      protobuf.SortBy_SORT_BY_POSITION,
				SortOrder:   protobuf.SortOrder_SORT_ORDER_ASC,
			},
			expectErr: false,
			expect: &protobuf.GetCategoriesResponse{
				PageInfo: &protobuf.PageInfo{
					PerPage:   3,
					Page:      1,
					PageCount: 1,
					Count:     1,
				},
				Categories: []*protobuf.Category{
					{
						Id:           "3345678",
						Position:     0,
						CreatedAt:    models.FixCreatedAtProto1,
						UpdatedAt:    models.FixUpdatedAtProto1,
						SourceLocale: "en-us",
						Outdated:     false,
						CountryCode:  "tw",
						Url:          "www.honestbee.com",
						HtmlUrl:      "www.honestbee.com",
						Name:         "testing category 1",
						Description:  "",
						Locale:       "en-us",
						KeyName:      "food",
					},
				},
			},
		},
		{
			description: "testing models return not found case",
			input: &protobuf.GetCategoriesRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_ID,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				PerPage:     3,
				Page:        0,
				SortBy:      protobuf.SortBy_SORT_BY_POSITION,
				SortOrder:   protobuf.SortOrder_SORT_ORDER_ASC,
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing models return error case",
			input: &protobuf.GetCategoriesRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TH,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				PerPage:     3,
				Page:        0,
				SortBy:      protobuf.SortBy_SORT_BY_POSITION,
				SortOrder:   protobuf.SortOrder_SORT_ORDER_ASC,
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			resp, err := s.GetCategories(context.Background(), tt.input)

			if tt.expectErr && err == nil {
				t.Errorf("[%s] expect an error, actual == nil", tt.description)
			} else if !tt.expectErr && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			} else if diff := deep.Equal(tt.expect, resp); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}

func TestGetCategoryByCategoryIdOrKeyname(t *testing.T) {
	s := initServer()

	testCases := [...]struct {
		description string
		input       *protobuf.GetCategoryRequest
		expectErr   bool
		expect      *protobuf.GetCategoryResponse
	}{
		{
			description: "testing normal case",
			input: &protobuf.GetCategoryRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Locale:      protobuf.Locale_LOCALE_ZH_TW,
				Id: &protobuf.GetCategoryRequest_CategoryIdOrKeyname{
					CategoryIdOrKeyname: "groceries",
				},
			},
			expectErr: false,
			expect: &protobuf.GetCategoryResponse{
				Category: &protobuf.Category{
					Id:           "3345678",
					Position:     0,
					CreatedAt:    models.FixCreatedAtProto1,
					UpdatedAt:    models.FixUpdatedAtProto1,
					SourceLocale: "en-us",
					Outdated:     false,
					CountryCode:  "tw",
					KeyName:      "food",
					Url:          "www.honestbee.com",
					HtmlUrl:      "www.honestbee.com",
					Name:         "testing category 1",
					Description:  "",
					Locale:       "en-us",
				},
			},
		},
		{
			description: "testing models return not found case",
			input: &protobuf.GetCategoryRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_ID,
				Locale:      protobuf.Locale_LOCALE_ZH_TW,
				Id: &protobuf.GetCategoryRequest_CategoryIdOrKeyname{
					CategoryIdOrKeyname: "groceries",
				},
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing models return error case",
			input: &protobuf.GetCategoryRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TH,
				Locale:      protobuf.Locale_LOCALE_ZH_TW,
				Id: &protobuf.GetCategoryRequest_CategoryIdOrKeyname{
					CategoryIdOrKeyname: "groceries",
				},
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			resp, err := s.GetCategory(context.Background(), tt.input)

			if tt.expectErr && err == nil {
				t.Errorf("[%s] expect an error, actual == nil", tt.description)
			} else if !tt.expectErr && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			} else if diff := deep.Equal(tt.expect, resp); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}

func TestGetCategoryBySectionId(t *testing.T) {
	s := initServer()

	testCases := [...]struct {
		description string
		input       *protobuf.GetCategoryRequest
		expectErr   bool
		expect      *protobuf.GetCategoryResponse
	}{
		{
			description: "testing normal case",
			input: &protobuf.GetCategoryRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Locale:      protobuf.Locale_LOCALE_ZH_TW,
				Id: &protobuf.GetCategoryRequest_SectionId{
					SectionId: "3345679",
				},
			},
			expectErr: false,
			expect: &protobuf.GetCategoryResponse{
				Category: &protobuf.Category{
					Id:           "3345678",
					Position:     0,
					CreatedAt:    models.FixCreatedAtProto1,
					UpdatedAt:    models.FixUpdatedAtProto1,
					SourceLocale: "en-us",
					Outdated:     false,
					CountryCode:  "tw",
					KeyName:      "",
					Url:          "www.honestbee.com",
					HtmlUrl:      "www.honestbee.com",
					Name:         "testing category 1",
					Description:  "",
					Locale:       "en-us",
				},
			},
		},
		{
			description: "testing models return not found case",
			input: &protobuf.GetCategoryRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Locale:      protobuf.Locale_LOCALE_ID,
				Id: &protobuf.GetCategoryRequest_SectionId{
					SectionId: "3345679",
				},
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing models return error case",
			input: &protobuf.GetCategoryRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Locale:      protobuf.Locale_LOCALE_ZH_CN,
				Id: &protobuf.GetCategoryRequest_SectionId{
					SectionId: "3345679",
				},
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing input invalid section id case",
			input: &protobuf.GetCategoryRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Locale:      protobuf.Locale_LOCALE_ZH_TW,
				Id:          &protobuf.GetCategoryRequest_SectionId{},
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			resp, err := s.GetCategory(context.Background(), tt.input)

			if tt.expectErr && err == nil {
				t.Errorf("[%s] expect an error, actual == nil", tt.description)
			} else if !tt.expectErr && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			} else if diff := deep.Equal(tt.expect, resp); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}

func TestGetCategoryByArticleId(t *testing.T) {
	s := initServer()

	testCases := [...]struct {
		description string
		input       *protobuf.GetCategoryRequest
		expectErr   bool
		expect      *protobuf.GetCategoryResponse
	}{
		{
			description: "testing normal case",
			input: &protobuf.GetCategoryRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Locale:      protobuf.Locale_LOCALE_ZH_TW,
				Id: &protobuf.GetCategoryRequest_ArticleId{
					ArticleId: "3345680",
				},
			},
			expectErr: false,
			expect: &protobuf.GetCategoryResponse{
				Category: &protobuf.Category{
					Id:           "3345678",
					Position:     0,
					CreatedAt:    models.FixCreatedAtProto1,
					UpdatedAt:    models.FixUpdatedAtProto1,
					SourceLocale: "en-us",
					Outdated:     false,
					CountryCode:  "tw",
					KeyName:      "",
					Url:          "www.honestbee.com",
					HtmlUrl:      "www.honestbee.com",
					Name:         "testing category 1",
					Description:  "",
					Locale:       "en-us",
				},
			},
		},
		{
			description: "testing models return not found case",
			input: &protobuf.GetCategoryRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Locale:      protobuf.Locale_LOCALE_ID,
				Id: &protobuf.GetCategoryRequest_ArticleId{
					ArticleId: "3345680",
				},
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing models return error case",
			input: &protobuf.GetCategoryRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Locale:      protobuf.Locale_LOCALE_ZH_CN,
				Id: &protobuf.GetCategoryRequest_ArticleId{
					ArticleId: "3345680",
				},
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing input invalid article id case",
			input: &protobuf.GetCategoryRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Locale:      protobuf.Locale_LOCALE_ZH_TW,
				Id:          &protobuf.GetCategoryRequest_ArticleId{},
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			resp, err := s.GetCategory(context.Background(), tt.input)

			if tt.expectErr && err == nil {
				t.Errorf("[%s] expect an error, actual == nil", tt.description)
			} else if !tt.expectErr && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			} else if diff := deep.Equal(tt.expect, resp); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}
