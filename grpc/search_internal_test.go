package grpc

import (
	"context"
	"net/http"
	"testing"

	"github.com/go-test/deep"
	"github.com/h2non/gock"

	"github.com/honestbee/Zen/models"
	"github.com/honestbee/Zen/protobuf"
	"github.com/honestbee/Zen/zendesk"
)

func TestGetSearchTitleArticles(t *testing.T) {
	gock.New("https://honestbeehelp-sg.zendesk.com").
		Get("hc/api/internal/instant_search.json").
		Filter(func(req *http.Request) bool {
			return req.URL.Query().Get("query") == "order" && req.URL.Query().Get("locale") == "en-us"
		}).
		Reply(http.StatusOK).
		JSON(&zendesk.InstantSearch{
			Results: []*zendesk.InstantSearchResult{
				{
					Title:         "How is my <em>order</em> confirmed?",
					CategoryTitle: "Goodship",
					URL:           "/hc/search/instant_click?data=1",
				},
				{
					Title:         "Where does my honestbee <em>order</em> come from?",
					CategoryTitle: "Grocery",
					URL:           "/hc/search/instant_click?data=2",
				},
				{
					Title:         "Can i reschedule my <em>order</em>?",
					CategoryTitle: "Laundry",
					URL:           "/hc/search/instant_click?data=3",
				},
			},
		})

	gock.New("https://honestbeehelp-sg.zendesk.com").
		Get("hc/api/internal/instant_search.json").
		Filter(func(req *http.Request) bool {
			return req.URL.Query().Get("query") == "order" && req.URL.Query().Get("locale") == "en-us"
		}).
		Reply(http.StatusOK).
		JSON(&zendesk.InstantSearch{
			Results: []*zendesk.InstantSearchResult{
				{
					Title:         "How is my <em>order</em> confirmed?",
					CategoryTitle: "Goodship",
					URL:           "/hc/search/instant_click?data=1",
				},
				{
					Title:         "Where does my honestbee <em>order</em> come from?",
					CategoryTitle: "Grocery",
					URL:           "/hc/search/instant_click?data=2",
				},
				{
					Title:         "Can i reschedule my <em>order</em>?",
					CategoryTitle: "Laundry",
					URL:           "/hc/search/instant_click?data=3",
				},
			},
		})

	gock.New("https://honestbeehelp-tw.zendesk.com").
		Get("hc/api/internal/instant_search.json").
		Filter(func(req *http.Request) bool {
			return req.URL.Query().Get("query") == "訂單" && req.URL.Query().Get("locale") == "zh-tw"
		}).
		Reply(http.StatusOK).
		JSON(&zendesk.InstantSearch{
			Results: []*zendesk.InstantSearchResult{
				{
					Title:         "我可以重<em>訂</em>我的<em>訂</em><em>單</em>嗎？",
					CategoryTitle: "熟食外送",
					URL:           "/hc/search/instant_click?data=1",
				},
			},
		})

	gock.New("https://honestbeehelp-tw.zendesk.com").
		Get("hc/api/internal/instant_search.json").
		Filter(func(req *http.Request) bool {
			return req.URL.Query().Get("query") == "order" && req.URL.Query().Get("locale") == "zh-tw"
		}).
		Reply(http.StatusOK).
		JSON(&zendesk.InstantSearch{
			Results: []*zendesk.InstantSearchResult{},
		})

	gock.New("https://honestbeehelp-tw.zendesk.com").
		Get("hc/api/internal/instant_search.json").
		Filter(func(req *http.Request) bool {
			return req.URL.Query().Get("query") == "too-many-request"
		}).
		Reply(http.StatusTooManyRequests).
		JSON(&zendesk.InstantSearch{})

	s := initServer()

	testCases := [...]struct {
		description string
		input       *protobuf.GetSearchTitleArticlesRequest
		expectErr   bool
		expect      *protobuf.GetSearchTitleArticlesResponse
	}{
		{
			description: "testing normal sg + en-us case",
			input: &protobuf.GetSearchTitleArticlesRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_SG,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				Query:       "order",
			},
			expectErr: false,
			expect: &protobuf.GetSearchTitleArticlesResponse{
				Articles: []*protobuf.SearchTitleArticle{
					{
						Title:         "How is my <em>order</em> confirmed?",
						CategoryTitle: "Goodship",
						Url:           "/hc/search/instant_click?data=1",
					},
					{
						Title:         "Where does my honestbee <em>order</em> come from?",
						CategoryTitle: "Grocery",
						Url:           "/hc/search/instant_click?data=2",
					},
					{
						Title:         "Can i reschedule my <em>order</em>?",
						CategoryTitle: "Laundry",
						Url:           "/hc/search/instant_click?data=3",
					},
				},
			},
		},
		{
			description: "testing normal tw + zh-tw case",
			input: &protobuf.GetSearchTitleArticlesRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Locale:      protobuf.Locale_LOCALE_ZH_TW,
				Query:       "訂單",
			},
			expectErr: false,
			expect: &protobuf.GetSearchTitleArticlesResponse{
				Articles: []*protobuf.SearchTitleArticle{
					{
						Title:         "我可以重<em>訂</em>我的<em>訂</em><em>單</em>嗎？",
						CategoryTitle: "熟食外送",
						Url:           "/hc/search/instant_click?data=1",
					},
				},
			},
		},
		{
			description: "testing no result case",
			input: &protobuf.GetSearchTitleArticlesRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Locale:      protobuf.Locale_LOCALE_ZH_TW,
				Query:       "order",
			},
			expectErr: false,
			expect: &protobuf.GetSearchTitleArticlesResponse{
				Articles: []*protobuf.SearchTitleArticle{},
			},
		},
		{
			description: "testing too many request err case",
			input: &protobuf.GetSearchTitleArticlesRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				Query:       "too-many-request",
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			resp, err := s.GetSearchTitleArticles(context.Background(), tt.input)

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

func TestGetSearchBodyArticles(t *testing.T) {
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
			return req.URL.Query().Get("query") == "too-many-request" && req.URL.Query().Get("locale") == "en-us"
		}).
		Reply(http.StatusTooManyRequests)

	s := initServer()

	testCases := [...]struct {
		description string
		input       *protobuf.GetSearchBodyArticlesRequest
		expectErr   bool
		expect      *protobuf.GetSearchBodyArticlesResponse
	}{
		{
			description: "testing normal case",
			input: &protobuf.GetSearchBodyArticlesRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_SG,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				SortOrder:   protobuf.SortOrder_SORT_ORDER_ASC,
				PerPage:     3,
				Page:        0,
				Query:       "order",
			},
			expectErr: false,
			expect: &protobuf.GetSearchBodyArticlesResponse{
				PageInfo: &protobuf.PageInfo{
					PerPage:   10,
					Page:      1,
					PageCount: 2,
					Count:     12,
				},
				Articles: []*protobuf.SearchBodyArticle{
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
						SourceLocale:    "",
						Outdated:        false,
						EditedAt:        models.FixEditedAtProto1,
						CountryCode:     "sg",
						Url:             "www.honestbee.com",
						HtmlUrl:         "www.honestbee.com",
						Name:            "testing article 1",
						Title:           "testing article 1",
						Body:            "this is testing article 1",
						Locale:          "en-us",
						Snippet:         "We’ll send you an <em>order</em> confirmation email with: Your <em>order</em> number Items ordered Delivery time",
						SectionId:       "7654321",
						CategoryId:      "3345678",
						CategoryName:    "testing category 1",
					},
				},
			},
		},
		{
			description: "testing models return not found case",
			input: &protobuf.GetSearchBodyArticlesRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_ID,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				SortOrder:   protobuf.SortOrder_SORT_ORDER_ASC,
				PerPage:     3,
				Page:        0,
				Query:       "order",
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing models return error case",
			input: &protobuf.GetSearchBodyArticlesRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TH,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				SortOrder:   protobuf.SortOrder_SORT_ORDER_ASC,
				PerPage:     3,
				Page:        0,
				Query:       "order",
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing zendesk search return too-many-request case",
			input: &protobuf.GetSearchBodyArticlesRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_SG,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				SortOrder:   protobuf.SortOrder_SORT_ORDER_ASC,
				PerPage:     3,
				Page:        0,
				Query:       "too-many-request",
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			resp, err := s.GetSearchBodyArticles(context.Background(), tt.input)

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
