package grpc

import (
	"context"
	"testing"

	"github.com/go-test/deep"

	"github.com/honestbee/Zen/models"
	"github.com/honestbee/Zen/protobuf"
)

func TestGetSections(t *testing.T) {
	s := initServer()

	testCases := [...]struct {
		description string
		input       *protobuf.GetSectionsRequest
		expectErr   bool
		expect      *protobuf.GetSectionsResponse
	}{
		{
			description: "testing normal case",
			input: &protobuf.GetSectionsRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				PerPage:     3,
				Page:        0,
				SortBy:      protobuf.SortBy_SORT_BY_POSITION,
				SortOrder:   protobuf.SortOrder_SORT_ORDER_ASC,
				Id: &protobuf.GetSectionsRequest_All{
					All: true,
				},
			},
			expectErr: false,
			expect: &protobuf.GetSectionsResponse{
				PageInfo: &protobuf.PageInfo{
					PerPage:   3,
					Page:      1,
					PageCount: 1,
					Count:     1,
				},
				Sections: []*protobuf.Section{
					{
						Id:           "3345679",
						Position:     0,
						CreatedAt:    models.FixCreatedAtProto1,
						UpdatedAt:    models.FixUpdatedAtProto1,
						SourceLocale: "en-us",
						Outdated:     false,
						CountryCode:  "tw",
						Url:          "www.honestbee.com",
						HtmlUrl:      "www.honestbee.com",
						Name:         "testing section 1",
						Description:  "",
						Locale:       "en-us",
						CategoryId:   "3345678",
					},
				},
			},
		},
		{
			description: "testing models return not found case",
			input: &protobuf.GetSectionsRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_ID,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				PerPage:     3,
				Page:        0,
				SortBy:      protobuf.SortBy_SORT_BY_POSITION,
				SortOrder:   protobuf.SortOrder_SORT_ORDER_ASC,
				Id: &protobuf.GetSectionsRequest_All{
					All: true,
				},
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing models return error case",
			input: &protobuf.GetSectionsRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TH,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				PerPage:     3,
				Page:        0,
				SortBy:      protobuf.SortBy_SORT_BY_POSITION,
				SortOrder:   protobuf.SortOrder_SORT_ORDER_ASC,
				Id: &protobuf.GetSectionsRequest_All{
					All: true,
				},
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			resp, err := s.GetSections(context.Background(), tt.input)

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

func TestGetSectionsByCategoryId(t *testing.T) {
	s := initServer()

	testCases := [...]struct {
		description string
		input       *protobuf.GetSectionsRequest
		expectErr   bool
		expect      *protobuf.GetSectionsResponse
	}{
		{
			description: "testing normal case",
			input: &protobuf.GetSectionsRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				PerPage:     3,
				Page:        0,
				SortBy:      protobuf.SortBy_SORT_BY_POSITION,
				SortOrder:   protobuf.SortOrder_SORT_ORDER_ASC,
				Id: &protobuf.GetSectionsRequest_CategoryId{
					CategoryId: "3345678",
				},
			},
			expectErr: false,
			expect: &protobuf.GetSectionsResponse{
				PageInfo: &protobuf.PageInfo{
					PerPage:   3,
					Page:      1,
					PageCount: 1,
					Count:     1,
				},
				Sections: []*protobuf.Section{
					{
						Id:           "3345679",
						Position:     0,
						CreatedAt:    models.FixCreatedAtProto1,
						UpdatedAt:    models.FixUpdatedAtProto1,
						SourceLocale: "en-us",
						Outdated:     false,
						CountryCode:  "tw",
						Url:          "www.honestbee.com",
						HtmlUrl:      "www.honestbee.com",
						Name:         "testing section 1",
						Description:  "",
						Locale:       "en-us",
						CategoryId:   "3345678",
					},
				},
			},
		},
		{
			description: "testing models return not found case",
			input: &protobuf.GetSectionsRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_ID,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				PerPage:     3,
				Page:        0,
				SortBy:      protobuf.SortBy_SORT_BY_POSITION,
				SortOrder:   protobuf.SortOrder_SORT_ORDER_ASC,
				Id: &protobuf.GetSectionsRequest_CategoryId{
					CategoryId: "3345678",
				},
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing models return error case",
			input: &protobuf.GetSectionsRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TH,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				PerPage:     3,
				Page:        0,
				SortBy:      protobuf.SortBy_SORT_BY_POSITION,
				SortOrder:   protobuf.SortOrder_SORT_ORDER_ASC,
				Id: &protobuf.GetSectionsRequest_CategoryId{
					CategoryId: "3345678",
				},
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing invalid input case",
			input: &protobuf.GetSectionsRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				PerPage:     3,
				Page:        0,
				SortBy:      protobuf.SortBy_SORT_BY_POSITION,
				SortOrder:   protobuf.SortOrder_SORT_ORDER_ASC,
				Id: &protobuf.GetSectionsRequest_CategoryId{
					CategoryId: "",
				},
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			resp, err := s.GetSections(context.Background(), tt.input)

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

func TestGetSectionBySectionId(t *testing.T) {
	s := initServer()

	testCases := [...]struct {
		description string
		input       *protobuf.GetSectionRequest
		expectErr   bool
		expect      *protobuf.GetSectionResponse
	}{
		{
			description: "testing normal case",
			input: &protobuf.GetSectionRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				Id: &protobuf.GetSectionRequest_SectionId{
					SectionId: "3345679",
				},
			},
			expectErr: false,
			expect: &protobuf.GetSectionResponse{
				Section: &protobuf.Section{
					Id:           "3345679",
					Position:     0,
					CreatedAt:    models.FixCreatedAtProto1,
					UpdatedAt:    models.FixUpdatedAtProto1,
					SourceLocale: "en-us",
					Outdated:     false,
					CountryCode:  "sg",
					Url:          "www.honestbee.com",
					HtmlUrl:      "www.honestbee.com",
					Name:         "testing section 1",
					Description:  "",
					Locale:       "en-us",
					CategoryId:   "3345678",
				},
			},
		},
		{
			description: "testing models return not found case",
			input: &protobuf.GetSectionRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_ID,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				Id: &protobuf.GetSectionRequest_SectionId{
					SectionId: "3345679",
				},
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing models return error case",
			input: &protobuf.GetSectionRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TH,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				Id: &protobuf.GetSectionRequest_SectionId{
					SectionId: "3345679",
				},
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing invalid input case",
			input: &protobuf.GetSectionRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				Id: &protobuf.GetSectionRequest_SectionId{
					SectionId: "",
				},
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			resp, err := s.GetSection(context.Background(), tt.input)

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

func TestGetSectionByArticleId(t *testing.T) {
	s := initServer()

	testCases := [...]struct {
		description string
		input       *protobuf.GetSectionRequest
		expectErr   bool
		expect      *protobuf.GetSectionResponse
	}{
		{
			description: "testing normal case",
			input: &protobuf.GetSectionRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				Id: &protobuf.GetSectionRequest_ArticleId{
					ArticleId: "3345679",
				},
			},
			expectErr: false,
			expect: &protobuf.GetSectionResponse{
				Section: &protobuf.Section{
					Id:           "3345679",
					Position:     0,
					CreatedAt:    models.FixCreatedAtProto1,
					UpdatedAt:    models.FixUpdatedAtProto1,
					SourceLocale: "en-us",
					Outdated:     false,
					CountryCode:  "sg",
					Url:          "www.honestbee.com",
					HtmlUrl:      "www.honestbee.com",
					Name:         "testing section 1",
					Description:  "",
					Locale:       "en-us",
					CategoryId:   "3345678",
				},
			},
		},
		{
			description: "testing models return not found case",
			input: &protobuf.GetSectionRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_ID,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				Id: &protobuf.GetSectionRequest_ArticleId{
					ArticleId: "3345679",
				},
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing models return error case",
			input: &protobuf.GetSectionRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TH,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				Id: &protobuf.GetSectionRequest_ArticleId{
					ArticleId: "3345679",
				},
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing invalid input case",
			input: &protobuf.GetSectionRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				Id: &protobuf.GetSectionRequest_ArticleId{
					ArticleId: "",
				},
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			resp, err := s.GetSection(context.Background(), tt.input)

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
