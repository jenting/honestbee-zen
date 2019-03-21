package grpc

import (
	"context"
	"testing"

	"github.com/go-test/deep"

	"github.com/honestbee/Zen/models"
	"github.com/honestbee/Zen/protobuf"
)

func TestGetArticles(t *testing.T) {
	s := initServer()

	testCases := [...]struct {
		description string
		input       *protobuf.GetArticlesRequest
		expectErr   bool
		expect      *protobuf.GetArticlesResponse
	}{
		{
			description: "testing normal case",
			input: &protobuf.GetArticlesRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				PerPage:     3,
				Page:        0,
				SortBy:      protobuf.SortBy_SORT_BY_POSITION,
				SortOrder:   protobuf.SortOrder_SORT_ORDER_ASC,
				Id: &protobuf.GetArticlesRequest_All{
					All: true,
				},
			},
			expectErr: false,
			expect: &protobuf.GetArticlesResponse{
				PageInfo: &protobuf.PageInfo{
					PerPage:   3,
					Page:      1,
					PageCount: 1,
					Count:     1,
				},
				Articles: []*protobuf.Article{
					{
						Id:              "33456710",
						AuthorId:        "1234567",
						CommentsDisable: false,
						Draft:           false,
						Promoted:        false,
						Position:        0,
						VoteSum:         0,
						VoteCount:       0,
						CreatedAt:       models.FixCreatedAtProto1,
						UpdatedAt:       models.FixUpdatedAtProto1,
						SourceLocale:    "en-us",
						Outdated:        false,
						OutdatedLocales: []string{},
						EditedAt:        models.FixEditedAtProto1,
						LabelNames:      []string{},
						CountryCode:     "tw",
						Url:             "www.honestbee.com",
						HtmlUrl:         "www.honestbee.com",
						Name:            "testing article 1",
						Title:           "testing article 1",
						Body:            "this is testing article 1",
						Locale:          "en-us",
						SectionId:       "33456789",
					},
				},
			},
		},
		{
			description: "testing models return not found case",
			input: &protobuf.GetArticlesRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_ID,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				PerPage:     3,
				Page:        0,
				SortBy:      protobuf.SortBy_SORT_BY_POSITION,
				SortOrder:   protobuf.SortOrder_SORT_ORDER_ASC,
				Id: &protobuf.GetArticlesRequest_All{
					All: true,
				},
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing models return error case",
			input: &protobuf.GetArticlesRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TH,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				PerPage:     3,
				Page:        0,
				SortBy:      protobuf.SortBy_SORT_BY_POSITION,
				SortOrder:   protobuf.SortOrder_SORT_ORDER_ASC,
				Id: &protobuf.GetArticlesRequest_All{
					All: true,
				},
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			resp, err := s.GetArticles(context.Background(), tt.input)

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

func TestGetArticlesByCategoryId(t *testing.T) {
	s := initServer()

	testCases := [...]struct {
		description string
		input       *protobuf.GetArticlesRequest
		expectErr   bool
		expect      *protobuf.GetArticlesResponse
	}{
		{
			description: "testing normal case w/o labels",
			input: &protobuf.GetArticlesRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				PerPage:     3,
				Page:        0,
				SortBy:      protobuf.SortBy_SORT_BY_POSITION,
				SortOrder:   protobuf.SortOrder_SORT_ORDER_ASC,
				Id: &protobuf.GetArticlesRequest_CategoryId{
					CategoryId: "3345678",
				},
				LabelNames: []string{},
			},
			expectErr: false,
			expect: &protobuf.GetArticlesResponse{
				PageInfo: &protobuf.PageInfo{
					PerPage:   3,
					Page:      1,
					PageCount: 2,
					Count:     4,
				},
				Articles: []*protobuf.Article{
					{
						Id:              "33456710",
						AuthorId:        "1234567",
						CommentsDisable: false,
						Draft:           false,
						Promoted:        false,
						Position:        0,
						VoteSum:         0,
						VoteCount:       0,
						CreatedAt:       models.FixCreatedAtProto1,
						UpdatedAt:       models.FixUpdatedAtProto1,
						SourceLocale:    "en-us",
						Outdated:        false,
						OutdatedLocales: []string{},
						EditedAt:        models.FixEditedAtProto1,
						LabelNames:      []string{"confirmed"},
						CountryCode:     "tw",
						Url:             "www.honestbee.com",
						HtmlUrl:         "www.honestbee.com",
						Name:            "testing article 1",
						Title:           "testing article 1",
						Body:            "this is testing article 1",
						Locale:          "en-us",
						SectionId:       "33456789",
					},
					{
						Id:              "33456711",
						AuthorId:        "1234568",
						CommentsDisable: false,
						Draft:           false,
						Promoted:        false,
						Position:        0,
						VoteSum:         0,
						VoteCount:       0,
						CreatedAt:       models.FixCreatedAtProto1,
						UpdatedAt:       models.FixUpdatedAtProto1,
						SourceLocale:    "en-us",
						Outdated:        false,
						OutdatedLocales: []string{},
						EditedAt:        models.FixEditedAtProto1,
						LabelNames:      []string{"preparing"},
						CountryCode:     "tw",
						Url:             "www.honestbee.com",
						HtmlUrl:         "www.honestbee.com",
						Name:            "testing article 2",
						Title:           "testing article 2",
						Body:            "this is testing article 2",
						Locale:          "en-us",
						SectionId:       "33456789",
					},
					{
						Id:              "33456712",
						AuthorId:        "1234569",
						CommentsDisable: false,
						Draft:           false,
						Promoted:        false,
						Position:        0,
						VoteSum:         0,
						VoteCount:       0,
						CreatedAt:       models.FixCreatedAtProto1,
						UpdatedAt:       models.FixUpdatedAtProto1,
						SourceLocale:    "en-us",
						Outdated:        false,
						OutdatedLocales: []string{},
						EditedAt:        models.FixEditedAtProto1,
						LabelNames:      []string{"ontheway"},
						CountryCode:     "tw",
						Url:             "www.honestbee.com",
						HtmlUrl:         "www.honestbee.com",
						Name:            "testing article 3",
						Title:           "testing article 3",
						Body:            "this is testing article 3",
						Locale:          "en-us",
						SectionId:       "33456789",
					},
					{
						Id:              "33456713",
						AuthorId:        "1234570",
						CommentsDisable: false,
						Draft:           false,
						Promoted:        false,
						Position:        0,
						VoteSum:         0,
						VoteCount:       0,
						CreatedAt:       models.FixCreatedAtProto1,
						UpdatedAt:       models.FixUpdatedAtProto1,
						SourceLocale:    "en-us",
						Outdated:        false,
						OutdatedLocales: []string{},
						EditedAt:        models.FixEditedAtProto1,
						LabelNames:      []string{"delivered"},
						CountryCode:     "tw",
						Url:             "www.honestbee.com",
						HtmlUrl:         "www.honestbee.com",
						Name:            "testing article 4",
						Title:           "testing article 4",
						Body:            "this is testing article 4",
						Locale:          "en-us",
						SectionId:       "33456789",
					},
				},
			},
		},
		{
			description: "testing normal case w/ labels confirmed",
			input: &protobuf.GetArticlesRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				PerPage:     3,
				Page:        0,
				SortBy:      protobuf.SortBy_SORT_BY_POSITION,
				SortOrder:   protobuf.SortOrder_SORT_ORDER_ASC,
				Id: &protobuf.GetArticlesRequest_CategoryId{
					CategoryId: "3345678",
				},
				LabelNames: []string{"confirmed"},
			},
			expectErr: false,
			expect: &protobuf.GetArticlesResponse{
				PageInfo: &protobuf.PageInfo{
					PerPage:   3,
					Page:      1,
					PageCount: 1,
					Count:     1,
				},
				Articles: []*protobuf.Article{
					{
						Id:              "33456710",
						AuthorId:        "1234567",
						CommentsDisable: false,
						Draft:           false,
						Promoted:        false,
						Position:        0,
						VoteSum:         0,
						VoteCount:       0,
						CreatedAt:       models.FixCreatedAtProto1,
						UpdatedAt:       models.FixUpdatedAtProto1,
						SourceLocale:    "en-us",
						Outdated:        false,
						OutdatedLocales: []string{},
						EditedAt:        models.FixEditedAtProto1,
						LabelNames:      []string{"confirmed"},
						CountryCode:     "tw",
						Url:             "www.honestbee.com",
						HtmlUrl:         "www.honestbee.com",
						Name:            "testing article 1",
						Title:           "testing article 1",
						Body:            "this is testing article 1",
						Locale:          "en-us",
						SectionId:       "33456789",
					},
				},
			},
		},
		{
			description: "testing models return not found case",
			input: &protobuf.GetArticlesRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_ID,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				PerPage:     3,
				Page:        0,
				SortBy:      protobuf.SortBy_SORT_BY_POSITION,
				SortOrder:   protobuf.SortOrder_SORT_ORDER_ASC,
				Id: &protobuf.GetArticlesRequest_CategoryId{
					CategoryId: "3345678",
				},
				LabelNames: []string{},
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing models return error case",
			input: &protobuf.GetArticlesRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TH,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				PerPage:     3,
				Page:        0,
				SortBy:      protobuf.SortBy_SORT_BY_POSITION,
				SortOrder:   protobuf.SortOrder_SORT_ORDER_ASC,
				Id: &protobuf.GetArticlesRequest_CategoryId{
					CategoryId: "3345678",
				},
				LabelNames: []string{},
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing invalid input case",
			input: &protobuf.GetArticlesRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TH,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				PerPage:     3,
				Page:        0,
				SortBy:      protobuf.SortBy_SORT_BY_POSITION,
				SortOrder:   protobuf.SortOrder_SORT_ORDER_ASC,
				Id: &protobuf.GetArticlesRequest_CategoryId{
					CategoryId: "",
				},
				LabelNames: []string{},
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			resp, err := s.GetArticles(context.Background(), tt.input)

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

func TestGetArticlesBySectionId(t *testing.T) {
	s := initServer()

	testCases := [...]struct {
		description string
		input       *protobuf.GetArticlesRequest
		expectErr   bool
		expect      *protobuf.GetArticlesResponse
	}{
		{
			description: "testing normal case",
			input: &protobuf.GetArticlesRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				PerPage:     3,
				Page:        0,
				SortBy:      protobuf.SortBy_SORT_BY_POSITION,
				SortOrder:   protobuf.SortOrder_SORT_ORDER_ASC,
				Id: &protobuf.GetArticlesRequest_SectionId{
					SectionId: "3345679",
				},
			},
			expectErr: false,
			expect: &protobuf.GetArticlesResponse{
				PageInfo: &protobuf.PageInfo{
					PerPage:   3,
					Page:      1,
					PageCount: 1,
					Count:     1,
				},
				Articles: []*protobuf.Article{
					{
						Id:              "33456710",
						AuthorId:        "1234567",
						CommentsDisable: false,
						Draft:           false,
						Promoted:        false,
						Position:        0,
						VoteSum:         0,
						VoteCount:       0,
						CreatedAt:       models.FixCreatedAtProto1,
						UpdatedAt:       models.FixUpdatedAtProto1,
						SourceLocale:    "en-us",
						Outdated:        false,
						OutdatedLocales: []string{},
						EditedAt:        models.FixEditedAtProto1,
						LabelNames:      []string{},
						CountryCode:     "tw",
						Url:             "www.honestbee.com",
						HtmlUrl:         "www.honestbee.com",
						Name:            "testing article 1",
						Title:           "testing article 1",
						Body:            "this is testing article 1",
						Locale:          "en-us",
						SectionId:       "33456789",
					},
				},
			},
		},
		{
			description: "testing models return not found case",
			input: &protobuf.GetArticlesRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_ID,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				PerPage:     3,
				Page:        0,
				SortBy:      protobuf.SortBy_SORT_BY_POSITION,
				SortOrder:   protobuf.SortOrder_SORT_ORDER_ASC,
				Id: &protobuf.GetArticlesRequest_SectionId{
					SectionId: "3345679",
				},
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing models return error case",
			input: &protobuf.GetArticlesRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TH,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				PerPage:     3,
				Page:        0,
				SortBy:      protobuf.SortBy_SORT_BY_POSITION,
				SortOrder:   protobuf.SortOrder_SORT_ORDER_ASC,
				Id: &protobuf.GetArticlesRequest_SectionId{
					SectionId: "3345679",
				},
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing invalid input case",
			input: &protobuf.GetArticlesRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TH,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				PerPage:     3,
				Page:        0,
				SortBy:      protobuf.SortBy_SORT_BY_POSITION,
				SortOrder:   protobuf.SortOrder_SORT_ORDER_ASC,
				Id: &protobuf.GetArticlesRequest_SectionId{
					SectionId: "",
				},
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			resp, err := s.GetArticles(context.Background(), tt.input)

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

func TestGetTopArticles(t *testing.T) {
	s := initServer()

	testCases := [...]struct {
		description string
		input       *protobuf.GetTopArticlesRequest
		expectErr   bool
		expect      *protobuf.GetTopArticlesResponse
	}{
		{
			description: "testing top5 case",
			input: &protobuf.GetTopArticlesRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				TopN:        5,
			},
			expectErr: false,
			expect: &protobuf.GetTopArticlesResponse{
				Articles: []*protobuf.Article{
					{
						Id:              "33456710",
						AuthorId:        "1234567",
						CommentsDisable: false,
						Draft:           false,
						Promoted:        false,
						Position:        0,
						VoteSum:         0,
						VoteCount:       0,
						CreatedAt:       models.FixCreatedAtProto1,
						UpdatedAt:       models.FixUpdatedAtProto1,
						SourceLocale:    "en-us",
						Outdated:        false,
						OutdatedLocales: []string{},
						EditedAt:        models.FixEditedAtProto1,
						LabelNames:      []string{},
						CountryCode:     "tw",
						Url:             "www.honestbee.com",
						HtmlUrl:         "www.honestbee.com",
						Name:            "testing article 1",
						Title:           "testing article 1",
						Body:            "this is testing article 1",
						Locale:          "en-us",
						SectionId:       "33456789",
					},
					{
						Id:              "33456711",
						AuthorId:        "1234568",
						CommentsDisable: false,
						Draft:           false,
						Promoted:        false,
						Position:        0,
						VoteSum:         0,
						VoteCount:       0,
						CreatedAt:       models.FixCreatedAtProto1,
						UpdatedAt:       models.FixUpdatedAtProto1,
						SourceLocale:    "en-us",
						Outdated:        false,
						OutdatedLocales: []string{},
						EditedAt:        models.FixEditedAtProto1,
						LabelNames:      []string{},
						CountryCode:     "tw",
						Url:             "www.honestbee.com",
						HtmlUrl:         "www.honestbee.com",
						Name:            "testing article 2",
						Title:           "testing article 2",
						Body:            "this is testing article 2",
						Locale:          "en-us",
						SectionId:       "33456789",
					},
					{
						Id:              "33456712",
						AuthorId:        "1234569",
						CommentsDisable: false,
						Draft:           false,
						Promoted:        false,
						Position:        0,
						VoteSum:         0,
						VoteCount:       0,
						CreatedAt:       models.FixCreatedAtProto1,
						UpdatedAt:       models.FixUpdatedAtProto1,
						SourceLocale:    "en-us",
						Outdated:        false,
						OutdatedLocales: []string{},
						EditedAt:        models.FixEditedAtProto1,
						LabelNames:      []string{},
						CountryCode:     "tw",
						Url:             "www.honestbee.com",
						HtmlUrl:         "www.honestbee.com",
						Name:            "testing article 3",
						Title:           "testing article 3",
						Body:            "this is testing article 3",
						Locale:          "en-us",
						SectionId:       "33456789",
					},
					{
						Id:              "33456713",
						AuthorId:        "1234570",
						CommentsDisable: false,
						Draft:           false,
						Promoted:        false,
						Position:        0,
						VoteSum:         0,
						VoteCount:       0,
						CreatedAt:       models.FixCreatedAtProto1,
						UpdatedAt:       models.FixUpdatedAtProto1,
						SourceLocale:    "en-us",
						Outdated:        false,
						OutdatedLocales: []string{},
						EditedAt:        models.FixEditedAtProto1,
						LabelNames:      []string{},
						CountryCode:     "tw",
						Url:             "www.honestbee.com",
						HtmlUrl:         "www.honestbee.com",
						Name:            "testing article 4",
						Title:           "testing article 4",
						Body:            "this is testing article 4",
						Locale:          "en-us",
						SectionId:       "33456789",
					},
					{
						Id:              "33456714",
						AuthorId:        "1234571",
						CommentsDisable: false,
						Draft:           false,
						Promoted:        false,
						Position:        0,
						VoteSum:         0,
						VoteCount:       0,
						CreatedAt:       models.FixCreatedAtProto1,
						UpdatedAt:       models.FixUpdatedAtProto1,
						SourceLocale:    "en-us",
						Outdated:        false,
						OutdatedLocales: []string{},
						EditedAt:        models.FixEditedAtProto1,
						LabelNames:      []string{},
						CountryCode:     "tw",
						Url:             "www.honestbee.com",
						HtmlUrl:         "www.honestbee.com",
						Name:            "testing article 5",
						Title:           "testing article 5",
						Body:            "this is testing article 5",
						Locale:          "en-us",
						SectionId:       "33456789",
					},
				},
			},
		},
		{
			description: "testing top4 case",
			input: &protobuf.GetTopArticlesRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				TopN:        4,
			},
			expectErr: false,
			expect: &protobuf.GetTopArticlesResponse{
				Articles: []*protobuf.Article{
					{
						Id:              "33456710",
						AuthorId:        "1234567",
						CommentsDisable: false,
						Draft:           false,
						Promoted:        false,
						Position:        0,
						VoteSum:         0,
						VoteCount:       0,
						CreatedAt:       models.FixCreatedAtProto1,
						UpdatedAt:       models.FixUpdatedAtProto1,
						SourceLocale:    "en-us",
						Outdated:        false,
						OutdatedLocales: []string{},
						EditedAt:        models.FixEditedAtProto1,
						LabelNames:      []string{},
						CountryCode:     "tw",
						Url:             "www.honestbee.com",
						HtmlUrl:         "www.honestbee.com",
						Name:            "testing article 1",
						Title:           "testing article 1",
						Body:            "this is testing article 1",
						Locale:          "en-us",
						SectionId:       "33456789",
					},
					{
						Id:              "33456711",
						AuthorId:        "1234568",
						CommentsDisable: false,
						Draft:           false,
						Promoted:        false,
						Position:        0,
						VoteSum:         0,
						VoteCount:       0,
						CreatedAt:       models.FixCreatedAtProto1,
						UpdatedAt:       models.FixUpdatedAtProto1,
						SourceLocale:    "en-us",
						Outdated:        false,
						OutdatedLocales: []string{},
						EditedAt:        models.FixEditedAtProto1,
						LabelNames:      []string{},
						CountryCode:     "tw",
						Url:             "www.honestbee.com",
						HtmlUrl:         "www.honestbee.com",
						Name:            "testing article 2",
						Title:           "testing article 2",
						Body:            "this is testing article 2",
						Locale:          "en-us",
						SectionId:       "33456789",
					},
					{
						Id:              "33456712",
						AuthorId:        "1234569",
						CommentsDisable: false,
						Draft:           false,
						Promoted:        false,
						Position:        0,
						VoteSum:         0,
						VoteCount:       0,
						CreatedAt:       models.FixCreatedAtProto1,
						UpdatedAt:       models.FixUpdatedAtProto1,
						SourceLocale:    "en-us",
						Outdated:        false,
						OutdatedLocales: []string{},
						EditedAt:        models.FixEditedAtProto1,
						LabelNames:      []string{},
						CountryCode:     "tw",
						Url:             "www.honestbee.com",
						HtmlUrl:         "www.honestbee.com",
						Name:            "testing article 3",
						Title:           "testing article 3",
						Body:            "this is testing article 3",
						Locale:          "en-us",
						SectionId:       "33456789",
					},
					{
						Id:              "33456713",
						AuthorId:        "1234570",
						CommentsDisable: false,
						Draft:           false,
						Promoted:        false,
						Position:        0,
						VoteSum:         0,
						VoteCount:       0,
						CreatedAt:       models.FixCreatedAtProto1,
						UpdatedAt:       models.FixUpdatedAtProto1,
						SourceLocale:    "en-us",
						Outdated:        false,
						OutdatedLocales: []string{},
						EditedAt:        models.FixEditedAtProto1,
						LabelNames:      []string{},
						CountryCode:     "tw",
						Url:             "www.honestbee.com",
						HtmlUrl:         "www.honestbee.com",
						Name:            "testing article 4",
						Title:           "testing article 4",
						Body:            "this is testing article 4",
						Locale:          "en-us",
						SectionId:       "33456789",
					},
				},
			},
		},
		{
			description: "testing models return not found case",
			input: &protobuf.GetTopArticlesRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_ID,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				TopN:        5,
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing models return error case",
			input: &protobuf.GetTopArticlesRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TH,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				TopN:        5,
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			resp, err := s.GetTopArticles(context.Background(), tt.input)

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

func TestGetArticle(t *testing.T) {
	s := initServer()

	testCases := [...]struct {
		description string
		input       *protobuf.GetArticleRequest
		expectErr   bool
		expect      *protobuf.GetArticleResponse
	}{
		{
			description: "testing normal case",
			input: &protobuf.GetArticleRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				ArticleId:   "33456710",
			},
			expectErr: false,
			expect: &protobuf.GetArticleResponse{
				Article: &protobuf.Article{
					Id:              "33456710",
					AuthorId:        "1234567",
					CommentsDisable: false,
					Draft:           false,
					Promoted:        false,
					Position:        0,
					VoteSum:         0,
					VoteCount:       0,
					CreatedAt:       models.FixCreatedAtProto1,
					UpdatedAt:       models.FixUpdatedAtProto1,
					SourceLocale:    "en-us",
					Outdated:        false,
					OutdatedLocales: []string{},
					EditedAt:        models.FixEditedAtProto1,
					LabelNames:      []string{},
					CountryCode:     "tw",
					Url:             "www.honestbee.com",
					HtmlUrl:         "www.honestbee.com",
					Name:            "testing article 1",
					Title:           "testing article 1",
					Body:            "this is testing article 1",
					Locale:          "en-us",
					SectionId:       "33456789",
				},
			},
		},
		{
			description: "testing invalid input case",
			input: &protobuf.GetArticleRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				ArticleId:   "",
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing models return not found case",
			input: &protobuf.GetArticleRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_ID,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				ArticleId:   "33456710",
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing models return error case",
			input: &protobuf.GetArticleRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TH,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				ArticleId:   "33456710",
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			resp, err := s.GetArticle(context.Background(), tt.input)

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
