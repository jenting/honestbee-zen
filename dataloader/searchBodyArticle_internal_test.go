package dataloader

import (
	"context"
	"net/http"
	"testing"

	"github.com/go-test/deep"
	"gopkg.in/h2non/gock.v1"

	"github.com/honestbee/Zen/inout"
	"github.com/honestbee/Zen/models"
	"github.com/honestbee/Zen/zendesk"
)

func TestLoadSearchBodyArticles(t *testing.T) {
	gock.New("https://honestbeehelp-sg.zendesk.com").
		Get("/api/v2/help_center/articles/search.json").
		Filter(func(req *http.Request) bool {
			return req.URL.Query().Get("query") == "order" && req.URL.Query().Get("locale") == "en-us"
		}).
		Reply(http.StatusOK).
		JSON(&zendesk.Search{
			Articles: []*zendesk.SearchArticle{
				{
					Article: &zendesk.Article{
						ID:              33456710,
						SectionID:       7654321,
						AuthorID:        1234567,
						CommentsDisable: false,
						Draft:           false,
						Promoted:        false,
						Position:        0,
						VoteSum:         0,
						VoteCount:       0,
						CreatedAt:       models.FixCreatedAt1,
						UpdatedAt:       models.FixUpdatedAt1,
						EditedAt:        models.FixEditedAt1,
						URL:             "www.honestbee.com",
						HTMLURL:         "www.honestbee.com",
						Name:            "testing article 1",
						Title:           "testing article 1",
						Body:            "this is testing article 1",
						Locale:          "en-us"},
					Snippet:    "We’ll send you an <em>order</em> confirmation email with: Your <em>order</em> number Items ordered Delivery time",
					ResultType: "article",
				},
			},
			BaseOut: &zendesk.BaseOut{
				PerPage:   10,
				Page:      1,
				PageCount: 2,
				Count:     12,
			},
		})

	gock.New("https://honestbeehelp-sg.zendesk.com").
		Get("/api/v2/help_center/articles/search.json").
		Filter(func(req *http.Request) bool {
			return req.URL.Query().Get("query") == "too-many-request"
		}).
		Reply(http.StatusTooManyRequests).
		JSON(&zendesk.Search{})

	testCases := [...]struct {
		description  string
		inputContext context.Context
		inputParams  interface{}
		expectErr    bool
		expect       *inout.GetSearchOut
	}{
		{
			description:  "testing sg + en-us normal case",
			inputContext: ctx,
			inputParams: inout.QuerySearchBodyArticlesIn{
				Query:       "order",
				CountryCode: "sg",
				Locale:      "en-us",
				PerPage:     30,
				Page:        1,
				SortOrder:   "asc",
			},
			expectErr: false,
			expect: &inout.GetSearchOut{
				Articles: []*models.SearchArticle{
					{
						Article: &models.Article{
							ID:              33456710,
							SectionID:       7654321,
							AuthorID:        1234567,
							CommentsDisable: false,
							Draft:           false,
							Promoted:        false,
							Position:        0,
							VoteSum:         0,
							VoteCount:       0,
							CreatedAt:       models.FixCreatedAt1,
							UpdatedAt:       models.FixUpdatedAt1,
							EditedAt:        models.FixEditedAt1,
							URL:             "www.honestbee.com",
							HTMLURL:         "www.honestbee.com",
							Name:            "testing article 1",
							Title:           "testing article 1",
							Body:            "this is testing article 1",
							Locale:          "en-us",
							CountryCode:     "sg",
						},
						Snippet:      "We’ll send you an <em>order</em> confirmation email with: Your <em>order</em> number Items ordered Delivery time",
						CategoryName: "testing category 1",
						CategoryID:   3345678,
					},
				},
				BaseOut: &inout.BaseOut{
					PerPage:   10,
					Page:      1,
					PageCount: 2,
					Count:     12,
				},
			},
		},
		{
			description:  "testing too many request err case",
			inputContext: ctx,
			inputParams: inout.QuerySearchBodyArticlesIn{
				Query:       "too-many-request",
				CountryCode: "sg",
				Locale:      "en-us",
				PerPage:     30,
				Page:        1,
				SortOrder:   "asc",
			},
			expectErr: true,
			expect:    nil,
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
			inputParams: inout.QuerySearchBodyArticlesIn{
				Query:       "order",
				CountryCode: "tw",
				Locale:      "en-us",
				PerPage:     3,
				Page:        1,
				SortOrder:   "asc",
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := LoadSearchBodyArticles(tt.inputContext, tt.inputParams)

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
