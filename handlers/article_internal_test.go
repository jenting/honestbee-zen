package handlers

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/go-test/deep"
	"github.com/julienschmidt/httprouter"

	"github.com/honestbee/Zen/inout"
	"github.com/honestbee/Zen/models"
)

func TestGetArticleDecompressor(t *testing.T) {
	testCases := [...]struct {
		description string
		input1      httprouter.Params
		input2      *http.Request
		expect      interface{}
		expectErr   bool
	}{
		{
			description: "testing normal case",
			input1: httprouter.Params{
				httprouter.Param{
					Key:   "article_id",
					Value: "33456711",
				},
			},
			input2: &http.Request{
				Form: url.Values{
					"locale":       []string{"en-us"},
					"country_code": []string{"tw"},
				},
			},
			expectErr: false,
			expect: &inout.GetArticleIn{
				ArticleID:   33456711,
				Locale:      "en-us",
				CountryCode: "tw",
			},
		},
		{
			description: "testing fetchBaseIn failed",
			input1:      nil,
			input2: &http.Request{
				Form: url.Values{
					"locale":       []string{"en-us"},
					"country_code": []string{"gg"},
				},
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing parse article id failed",
			input1: httprouter.Params{
				httprouter.Param{
					Key:   "article_id",
					Value: "abcdefg",
				},
			},
			input2: &http.Request{
				Form: url.Values{
					"locale":       []string{"en-us"},
					"country_code": []string{"tw"},
				},
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := GetArticleDecompressor(tt.input1, tt.input2)
			if tt.expectErr && err == nil {
				t.Errorf("[%s] expect an error, actual == nil", tt.description)
			} else if !tt.expectErr && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			} else if diff := deep.Equal(tt.expect, actual); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}

func TestGetArticleHandler(t *testing.T) {
	testCases := [...]struct {
		description string
		input       interface{}
		expect      interface{}
		expectErr   bool
	}{
		{
			description: "testing normal case",
			input: &inout.GetArticleIn{
				Locale:      "en-us",
				CountryCode: "tw",
				ArticleID:   3345679,
			},
			expectErr: false,
			expect: &inout.GetArticleOut{
				Article: &models.Article{
					ID:              33456710,
					AuthorID:        1234567,
					CommentsDisable: false,
					Draft:           false,
					Promoted:        false,
					Position:        0,
					VoteSum:         0,
					VoteCount:       0,
					CreatedAt:       models.FixCreatedAt1,
					UpdatedAt:       models.FixUpdatedAt1,
					SourceLocale:    "en-us",
					Outdated:        false,
					OutdatedLocales: []string{},
					EditedAt:        models.FixEditedAt1,
					LabelNames:      []string{},
					CountryCode:     "tw",
					URL:             "www.honestbee.com",
					HTMLURL:         "www.honestbee.com",
					Name:            "testing article 1",
					Title:           "testing article 1",
					Body:            "this is testing article 1",
					Locale:          "en-us",
					SectionID:       33456789,
				},
			},
		},
		{
			description: "testing input casting failed case",
			input: &struct {
				name string
				age  int
			}{
				name: "honestbee",
				age:  99,
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing models return error case",
			input: &inout.GetArticleIn{
				Locale:      "en-us",
				CountryCode: models.ModelsReturnErrorCountryCode,
				ArticleID:   3345679,
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing models return not found error case",
			input: &inout.GetArticleIn{
				Locale:      "en-us",
				CountryCode: models.ModelsReturnNotFoundCountryCode,
				ArticleID:   3345679,
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := GetArticleHandler(context.Background(), e, tt.input)
			if tt.expectErr && err == nil {
				t.Errorf("[%s] expect an error, actual nil", tt.description)
			} else if !tt.expectErr && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			} else if diff := deep.Equal(tt.expect, actual); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}

func TestGetTopNArticlesDecompressor(t *testing.T) {
	testCases := [...]struct {
		description string
		input1      httprouter.Params
		input2      *http.Request
		expect      interface{}
		expectErr   bool
	}{
		{
			description: "testing normal case",
			input1: httprouter.Params{
				httprouter.Param{
					Key:   "top_n",
					Value: "5",
				},
			},
			input2: &http.Request{
				Form: url.Values{
					"locale":       []string{"en-us"},
					"country_code": []string{"tw"},
				},
			},
			expectErr: false,
			expect: &inout.GetTopNArticlesIn{
				TopN:        5,
				Locale:      "en-us",
				CountryCode: "tw",
			},
		},
		{
			description: "testing fetchBaseIn failed",
			input1:      nil,
			input2: &http.Request{
				Form: url.Values{
					"locale":       []string{"en-us"},
					"country_code": []string{"gg"},
				},
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing parse top n failed",
			input1: httprouter.Params{
				httprouter.Param{
					Key:   "top_n",
					Value: "abcdefg",
				},
			},
			input2: &http.Request{
				Form: url.Values{
					"locale":       []string{"en-us"},
					"country_code": []string{"tw"},
				},
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := GetTopNArticlesDecompressor(tt.input1, tt.input2)
			if tt.expectErr && err == nil {
				t.Errorf("[%s] expect an error, actual == nil", tt.description)
			} else if !tt.expectErr && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			} else if diff := deep.Equal(tt.expect, actual); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}

func TestGetTopNArticlesHandler(t *testing.T) {
	testCases := [...]struct {
		description string
		input       interface{}
		expect      interface{}
		expectErr   bool
	}{
		{
			description: "testing normal case",
			input: &inout.GetTopNArticlesIn{
				TopN:        5,
				Locale:      "en-us",
				CountryCode: "tw",
			},
			expectErr: false,
			expect: &inout.GetTopNArticlesOut{
				Articles: []*models.Article{
					&models.Article{
						ID:              33456710,
						AuthorID:        1234567,
						CommentsDisable: false,
						Draft:           false,
						Promoted:        false,
						Position:        0,
						VoteSum:         0,
						VoteCount:       0,
						CreatedAt:       models.FixCreatedAt1,
						UpdatedAt:       models.FixUpdatedAt1,
						SourceLocale:    "en-us",
						Outdated:        false,
						OutdatedLocales: []string{},
						EditedAt:        models.FixEditedAt1,
						LabelNames:      []string{},
						CountryCode:     "tw",
						URL:             "www.honestbee.com",
						HTMLURL:         "www.honestbee.com",
						Name:            "testing article 1",
						Title:           "testing article 1",
						Body:            "this is testing article 1",
						Locale:          "en-us",
						SectionID:       33456789,
					},
					&models.Article{
						ID:              33456711,
						AuthorID:        1234568,
						CommentsDisable: false,
						Draft:           false,
						Promoted:        false,
						Position:        0,
						VoteSum:         0,
						VoteCount:       0,
						CreatedAt:       models.FixCreatedAt1,
						UpdatedAt:       models.FixUpdatedAt1,
						SourceLocale:    "en-us",
						Outdated:        false,
						OutdatedLocales: []string{},
						EditedAt:        models.FixEditedAt1,
						LabelNames:      []string{},
						CountryCode:     "tw",
						URL:             "www.honestbee.com",
						HTMLURL:         "www.honestbee.com",
						Name:            "testing article 2",
						Title:           "testing article 2",
						Body:            "this is testing article 2",
						Locale:          "en-us",
						SectionID:       33456789,
					},
					&models.Article{
						ID:              33456712,
						AuthorID:        1234569,
						CommentsDisable: false,
						Draft:           false,
						Promoted:        false,
						Position:        0,
						VoteSum:         0,
						VoteCount:       0,
						CreatedAt:       models.FixCreatedAt1,
						UpdatedAt:       models.FixUpdatedAt1,
						SourceLocale:    "en-us",
						Outdated:        false,
						OutdatedLocales: []string{},
						EditedAt:        models.FixEditedAt1,
						LabelNames:      []string{},
						CountryCode:     "tw",
						URL:             "www.honestbee.com",
						HTMLURL:         "www.honestbee.com",
						Name:            "testing article 3",
						Title:           "testing article 3",
						Body:            "this is testing article 3",
						Locale:          "en-us",
						SectionID:       33456789,
					},
					&models.Article{
						ID:              33456713,
						AuthorID:        1234570,
						CommentsDisable: false,
						Draft:           false,
						Promoted:        false,
						Position:        0,
						VoteSum:         0,
						VoteCount:       0,
						CreatedAt:       models.FixCreatedAt1,
						UpdatedAt:       models.FixUpdatedAt1,
						SourceLocale:    "en-us",
						Outdated:        false,
						OutdatedLocales: []string{},
						EditedAt:        models.FixEditedAt1,
						LabelNames:      []string{},
						CountryCode:     "tw",
						URL:             "www.honestbee.com",
						HTMLURL:         "www.honestbee.com",
						Name:            "testing article 4",
						Title:           "testing article 4",
						Body:            "this is testing article 4",
						Locale:          "en-us",
						SectionID:       33456789,
					},
					&models.Article{
						ID:              33456714,
						AuthorID:        1234571,
						CommentsDisable: false,
						Draft:           false,
						Promoted:        false,
						Position:        0,
						VoteSum:         0,
						VoteCount:       0,
						CreatedAt:       models.FixCreatedAt1,
						UpdatedAt:       models.FixUpdatedAt1,
						SourceLocale:    "en-us",
						Outdated:        false,
						OutdatedLocales: []string{},
						EditedAt:        models.FixEditedAt1,
						LabelNames:      []string{},
						CountryCode:     "tw",
						URL:             "www.honestbee.com",
						HTMLURL:         "www.honestbee.com",
						Name:            "testing article 5",
						Title:           "testing article 5",
						Body:            "this is testing article 5",
						Locale:          "en-us",
						SectionID:       33456789,
					},
				},
			},
		},
		{
			description: "testing different topN case",
			input: &inout.GetTopNArticlesIn{
				TopN:        4,
				Locale:      "en-us",
				CountryCode: "tw",
			},
			expectErr: false,
			expect: &inout.GetTopNArticlesOut{
				Articles: []*models.Article{
					&models.Article{
						ID:              33456710,
						AuthorID:        1234567,
						CommentsDisable: false,
						Draft:           false,
						Promoted:        false,
						Position:        0,
						VoteSum:         0,
						VoteCount:       0,
						CreatedAt:       models.FixCreatedAt1,
						UpdatedAt:       models.FixUpdatedAt1,
						SourceLocale:    "en-us",
						Outdated:        false,
						OutdatedLocales: []string{},
						EditedAt:        models.FixEditedAt1,
						LabelNames:      []string{},
						CountryCode:     "tw",
						URL:             "www.honestbee.com",
						HTMLURL:         "www.honestbee.com",
						Name:            "testing article 1",
						Title:           "testing article 1",
						Body:            "this is testing article 1",
						Locale:          "en-us",
						SectionID:       33456789,
					},
					&models.Article{
						ID:              33456711,
						AuthorID:        1234568,
						CommentsDisable: false,
						Draft:           false,
						Promoted:        false,
						Position:        0,
						VoteSum:         0,
						VoteCount:       0,
						CreatedAt:       models.FixCreatedAt1,
						UpdatedAt:       models.FixUpdatedAt1,
						SourceLocale:    "en-us",
						Outdated:        false,
						OutdatedLocales: []string{},
						EditedAt:        models.FixEditedAt1,
						LabelNames:      []string{},
						CountryCode:     "tw",
						URL:             "www.honestbee.com",
						HTMLURL:         "www.honestbee.com",
						Name:            "testing article 2",
						Title:           "testing article 2",
						Body:            "this is testing article 2",
						Locale:          "en-us",
						SectionID:       33456789,
					},
					&models.Article{
						ID:              33456712,
						AuthorID:        1234569,
						CommentsDisable: false,
						Draft:           false,
						Promoted:        false,
						Position:        0,
						VoteSum:         0,
						VoteCount:       0,
						CreatedAt:       models.FixCreatedAt1,
						UpdatedAt:       models.FixUpdatedAt1,
						SourceLocale:    "en-us",
						Outdated:        false,
						OutdatedLocales: []string{},
						EditedAt:        models.FixEditedAt1,
						LabelNames:      []string{},
						CountryCode:     "tw",
						URL:             "www.honestbee.com",
						HTMLURL:         "www.honestbee.com",
						Name:            "testing article 3",
						Title:           "testing article 3",
						Body:            "this is testing article 3",
						Locale:          "en-us",
						SectionID:       33456789,
					},
					&models.Article{
						ID:              33456713,
						AuthorID:        1234570,
						CommentsDisable: false,
						Draft:           false,
						Promoted:        false,
						Position:        0,
						VoteSum:         0,
						VoteCount:       0,
						CreatedAt:       models.FixCreatedAt1,
						UpdatedAt:       models.FixUpdatedAt1,
						SourceLocale:    "en-us",
						Outdated:        false,
						OutdatedLocales: []string{},
						EditedAt:        models.FixEditedAt1,
						LabelNames:      []string{},
						CountryCode:     "tw",
						URL:             "www.honestbee.com",
						HTMLURL:         "www.honestbee.com",
						Name:            "testing article 4",
						Title:           "testing article 4",
						Body:            "this is testing article 4",
						Locale:          "en-us",
						SectionID:       33456789,
					},
				},
			},
		},
		{
			description: "testing input casting failed case",
			input: &struct {
				topn int
			}{
				topn: 1,
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing models return error case",
			input: &inout.GetTopNArticlesIn{
				TopN:        5,
				Locale:      "en-us",
				CountryCode: models.ModelsReturnErrorCountryCode,
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing models return not found error case",
			input: &inout.GetTopNArticlesIn{
				TopN:        5,
				Locale:      "en-us",
				CountryCode: models.ModelsReturnNotFoundCountryCode,
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := GetTopNArticlesHandler(context.Background(), e, tt.input)
			if tt.expectErr && err == nil {
				t.Errorf("[%s] expect an error, actual nil", tt.description)
			} else if !tt.expectErr && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			} else if diff := deep.Equal(tt.expect, actual); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}
