package dataloader

import (
	"context"
	"net/http"
	"testing"

	"github.com/go-test/deep"
	"gopkg.in/h2non/gock.v1"

	"github.com/honestbee/Zen/inout"
	"github.com/honestbee/Zen/zendesk"
)

func TestLoadSearchTitleArticles(t *testing.T) {
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

	testCases := [...]struct {
		description  string
		inputContext context.Context
		inputParams  interface{}
		expectErr    bool
		expect       []*zendesk.InstantSearchResult
	}{
		{
			description:  "testing normal sg + en-us case",
			inputContext: ctx,
			inputParams: inout.QuerySearchTitleArticlesIn{
				Query:       "order",
				CountryCode: "sg",
				Locale:      "en-us",
			},
			expectErr: false,
			expect: []*zendesk.InstantSearchResult{
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
		},
		{
			description:  "testing normal tw + zh-tw case",
			inputContext: ctx,
			inputParams: inout.QuerySearchTitleArticlesIn{
				Query:       "訂單",
				CountryCode: "tw",
				Locale:      "zh-tw",
			},
			expectErr: false,
			expect: []*zendesk.InstantSearchResult{
				{
					Title:         "我可以重<em>訂</em>我的<em>訂</em><em>單</em>嗎？",
					CategoryTitle: "熟食外送",
					URL:           "/hc/search/instant_click?data=1",
				},
			},
		},
		{
			description:  "testing no result case",
			inputContext: ctx,
			inputParams: inout.QuerySearchTitleArticlesIn{
				Query:       "order",
				CountryCode: "tw",
				Locale:      "zh-tw",
			},
			expectErr: false,
			expect:    []*zendesk.InstantSearchResult{},
		},
		{
			description:  "testing too many request err case",
			inputContext: ctx,
			inputParams: inout.QuerySearchTitleArticlesIn{
				Query:       "too-many-request",
				CountryCode: "tw",
				Locale:      "en-us",
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
			inputParams: inout.QuerySearchTitleArticlesIn{
				Query:       "order",
				CountryCode: "tw",
				Locale:      "en-us",
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := LoadSearchTitleArticles(tt.inputContext, tt.inputParams)

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
