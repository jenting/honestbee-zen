package dataloader

import (
	"context"
	"testing"

	"github.com/go-test/deep"
	gographql "github.com/graph-gophers/graphql-go"

	"github.com/honestbee/Zen/inout"
	"github.com/honestbee/Zen/models"
)

func TestLoadArticles(t *testing.T) {
	sid := gographql.ID("3345679")
	cid := gographql.ID("3345678")

	testCases := [...]struct {
		description  string
		inputContext context.Context
		inputParams  interface{}
		expectErr    bool
		expect       *inout.GetArticlesOut
	}{
		{
			description:  "testing normal w/o category id and section id case",
			inputContext: ctx,
			inputParams: inout.QueryArticlesIn{
				CountryCode: "tw",
				Locale:      "en-us",
				PerPage:     3,
				Page:        1,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectErr: false,
			expect: &inout.GetArticlesOut{
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
				},
				BaseOut: &inout.BaseOut{
					Page:      1,
					PerPage:   3,
					PageCount: 1,
					Count:     1,
				},
			},
		},
		{
			description:  "testing json marshal parameter failed case",
			inputContext: ctx,
			inputParams:  make(chan int),
			expectErr:    true,
			expect:       nil,
		},
		{
			description:  "testing json unmarshal parameter failed case",
			inputContext: ctx,
			inputParams: &struct {
				CountryCode int32
				Locale      int32
			}{
				CountryCode: 123,
				Locale:      456,
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description:  "testing extract dataloader failed case",
			inputContext: context.TODO(),
			inputParams: inout.QueryArticlesIn{
				CountryCode: "tw",
				Locale:      "en-us",
				PerPage:     3,
				Page:        1,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description:  "testing invalid country code case",
			inputContext: ctx,
			inputParams: inout.QueryArticlesIn{
				CountryCode: models.ModelsReturnErrorCountryCode,
				Locale:      "en-us",
				PerPage:     3,
				Page:        1,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description:  "testing normal w/ section id case",
			inputContext: ctx,
			inputParams: inout.QueryArticlesIn{
				SectionID:   &sid,
				CountryCode: "tw",
				Locale:      "en-us",
				PerPage:     3,
				Page:        1,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectErr: false,
			expect: &inout.GetArticlesOut{
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
				},
				BaseOut: &inout.BaseOut{
					Page:      1,
					PerPage:   3,
					PageCount: 1,
					Count:     1,
				},
			},
		},
		{
			description:  "testing invalid section id case",
			inputContext: ctx,
			inputParams: inout.QueryArticlesIn{
				SectionID:   new(gographql.ID),
				CountryCode: "tw",
				Locale:      "en-us",
				PerPage:     3,
				Page:        1,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description:  "testing invalid country code case",
			inputContext: ctx,
			inputParams: inout.QueryArticlesIn{
				SectionID:   &sid,
				CountryCode: models.ModelsReturnErrorCountryCode,
				Locale:      "en-us",
				PerPage:     3,
				Page:        1,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description:  "testing normal w/ category id case",
			inputContext: ctx,
			inputParams: inout.QueryArticlesIn{
				CategoryID:  &cid,
				CountryCode: "tw",
				Locale:      "en-us",
				PerPage:     3,
				Page:        1,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectErr: false,
			expect: &inout.GetArticlesOut{
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
						LabelNames:      []string{"confirmed"},
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
						LabelNames:      []string{"preparing"},
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
						LabelNames:      []string{"ontheway"},
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
						LabelNames:      []string{"delivered"},
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
				BaseOut: &inout.BaseOut{
					Page:      1,
					PerPage:   3,
					PageCount: 2,
					Count:     4,
				},
			},
		},
		{
			description:  "testing invalid section id case",
			inputContext: ctx,
			inputParams: inout.QueryArticlesIn{
				CategoryID:  new(gographql.ID),
				CountryCode: "tw",
				Locale:      "en-us",
				PerPage:     3,
				Page:        1,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description:  "testing invalid country code case",
			inputContext: ctx,
			inputParams: inout.QueryArticlesIn{
				CategoryID:  &cid,
				CountryCode: models.ModelsReturnErrorCountryCode,
				Locale:      "en-us",
				PerPage:     3,
				Page:        1,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := LoadArticles(tt.inputContext, tt.inputParams)

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

func TestLoadTopArticles(t *testing.T) {
	testCases := [...]struct {
		description  string
		inputContext context.Context
		inputParams  interface{}
		expectErr    bool
		expect       []*models.Article
	}{
		{
			description:  "testing normal case",
			inputContext: ctx,
			inputParams: inout.QueryTopArticlesIn{
				TopN:        4,
				CountryCode: "tw",
				Locale:      "en-us",
			},
			expectErr: false,
			expect: []*models.Article{
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
		{
			description:  "testing json marshal parameter failed case",
			inputContext: ctx,
			inputParams:  make(chan int),
			expectErr:    true,
			expect:       nil,
		},
		{
			description:  "testing json unmarshal parameter failed case",
			inputContext: ctx,
			inputParams: &struct {
				CountryCode int32
				Locale      int32
			}{
				CountryCode: 123,
				Locale:      456,
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description:  "testing extract dataloader failed case",
			inputContext: context.TODO(),
			inputParams: inout.QueryTopArticlesIn{
				TopN:        4,
				CountryCode: "tw",
				Locale:      "en-us",
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description:  "testing invalid country code case",
			inputContext: ctx,
			inputParams: inout.QueryTopArticlesIn{
				TopN:        4,
				CountryCode: models.ModelsReturnErrorCountryCode,
				Locale:      "en-us",
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description:  "testing invalid article id case",
			inputContext: ctx,
			inputParams: inout.QueryTopArticlesIn{
				TopN:        -1,
				CountryCode: "tw",
				Locale:      "en-us",
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description:  "testing article id not found case",
			inputContext: ctx,
			inputParams: inout.QueryTopArticlesIn{
				TopN:        4,
				CountryCode: models.ModelsReturnNotFoundCountryCode,
				Locale:      "en-us",
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := LoadTopArticles(tt.inputContext, tt.inputParams)

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

func TestLoadArticle(t *testing.T) {
	testCases := [...]struct {
		description  string
		inputContext context.Context
		inputParams  interface{}
		expectErr    bool
		expect       *models.Article
	}{
		{
			description:  "testing normal case",
			inputContext: ctx,
			inputParams: inout.QueryArticleIn{
				ArticleID:   gographql.ID("33456710"),
				CountryCode: "tw",
				Locale:      "en-us",
			},
			expectErr: false,
			expect: &models.Article{
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
		{
			description:  "testing json marshal parameter failed case",
			inputContext: ctx,
			inputParams:  make(chan int),
			expectErr:    true,
			expect:       nil,
		},
		{
			description:  "testing json unmarshal parameter failed case",
			inputContext: ctx,
			inputParams: &struct {
				CountryCode int32
				Locale      int32
			}{
				CountryCode: 123,
				Locale:      456,
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description:  "testing extract dataloader failed case",
			inputContext: context.TODO(),
			inputParams: inout.QueryArticleIn{
				ArticleID:   gographql.ID("33456710"),
				CountryCode: "tw",
				Locale:      "en-us",
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description:  "testing invalid country code case",
			inputContext: ctx,
			inputParams: inout.QueryArticleIn{
				ArticleID:   gographql.ID("3345679"),
				CountryCode: models.ModelsReturnErrorCountryCode,
				Locale:      "en-us",
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description:  "testing invalid article id case",
			inputContext: ctx,
			inputParams: inout.QueryArticleIn{
				ArticleID:   "",
				CountryCode: "tw",
				Locale:      "en-us",
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description:  "testing article id not found case",
			inputContext: ctx,
			inputParams: inout.QueryArticleIn{
				ArticleID:   gographql.ID("3345679"),
				CountryCode: models.ModelsReturnNotFoundCountryCode,
				Locale:      "en-us",
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := LoadArticle(tt.inputContext, tt.inputParams)

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
