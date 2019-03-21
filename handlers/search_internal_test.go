package handlers

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/go-test/deep"
	"gopkg.in/h2non/gock.v1"

	"github.com/honestbee/Zen/inout"
	"github.com/honestbee/Zen/models"
	"github.com/honestbee/Zen/zendesk"
)

func TestInstantSearchDecompressor(t *testing.T) {
	testCases := [...]struct {
		description string
		input       *http.Request
		expect      interface{}
		expectErr   bool
	}{
		{
			description: "testing normal case",
			input: &http.Request{
				Form: url.Values{
					"query":        []string{"order"},
					"locale":       []string{"en-us"},
					"country_code": []string{"sg"},
				},
			},
			expectErr: false,
			expect: &inout.GetInstantSearchIn{
				Query:       "order",
				Locale:      "en-us",
				CountryCode: "sg",
			},
		},
		{
			description: "testing fetchBaseIn failed",
			input: &http.Request{
				Form: url.Values{
					"locale":       []string{"en-us"},
					"country_code": []string{"gg"},
				},
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing parse query value failed",
			input: &http.Request{
				Form: url.Values{
					"locale":       []string{"en-us"},
					"country_code": []string{"sg"},
				},
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := GetInstantSearchDecompressor(nil, tt.input)
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

func TestInstantSearchHandler(t *testing.T) {
	// Race condition happens when gock change back HTTP DefaultTransport after
	// TestCreateVoteHandler function, but defer function have not finish.
	// We remove gock.Off() to pass testing with race detector enabled.
	//defer gock.Off()

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

	testCases := [...]struct {
		description string
		input       interface{}
		expect      interface{}
		expectErr   bool
	}{
		{
			description: "testing normal case",
			input: &inout.GetInstantSearchIn{
				Query:       "order",
				Locale:      "en-us",
				CountryCode: "sg",
			},
			expectErr: false,
			expect: &inout.GetInstantSearchOut{
				Results: []*inout.InstantSearchResult{
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
		},
		{
			description: "testing tw normal case",
			input: &inout.GetInstantSearchIn{
				Query:       "訂單",
				Locale:      "zh-tw",
				CountryCode: "tw",
			},
			expectErr: false,
			expect: &inout.GetInstantSearchOut{
				Results: []*inout.InstantSearchResult{
					{
						Title:         "我可以重<em>訂</em>我的<em>訂</em><em>單</em>嗎？",
						CategoryTitle: "熟食外送",
						URL:           "/hc/search/instant_click?data=1",
					},
				},
			},
		},
		{
			description: "testing no result case",
			input: &inout.GetInstantSearchIn{
				Query:       "order",
				Locale:      "zh-tw",
				CountryCode: "tw",
			},
			expectErr: false,
			expect: &inout.GetInstantSearchOut{
				Results: []*inout.InstantSearchResult{},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := GetInstantSearchHandler(context.Background(), e, tt.input)
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

func TestSearchDecompressor(t *testing.T) {
	testCases := [...]struct {
		description string
		input       *http.Request
		expect      interface{}
		expectErr   bool
	}{
		{
			description: "testing normal case",
			input: &http.Request{
				Form: url.Values{
					"query":        []string{"order"},
					"locale":       []string{"en-us"},
					"country_code": []string{"sg"},
					"per_page":     []string{"3"},
					"page":         []string{"1"},
					"sort_by":      []string{"position"},
					"sort_order":   []string{"asc"},
				},
			},
			expectErr: false,
			expect: &inout.GetSearchIn{
				Query: "order",
				BaseIn: &inout.BaseIn{
					Locale:      "en-us",
					CountryCode: "sg",
					PerPage:     3,
					Page:        0,
					SortBy:      "position",
					SortOrder:   "asc"},
			},
		},
		{
			description: "testing fetchBaseIn failed",
			input: &http.Request{
				Form: url.Values{
					"locale":       []string{"en-us"},
					"country_code": []string{"gg"},
				},
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing parse query value failed",
			input: &http.Request{
				Form: url.Values{
					"locale":       []string{"en-us"},
					"country_code": []string{"sg"},
				},
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := GetSearchDecompressor(nil, tt.input)
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

func TestGetSearchHandler(t *testing.T) {
	// Race condition happens when gock change back HTTP DefaultTransport after
	// TestCreateVoteHandler function, but defer function have not finish.
	// We remove gock.Off() to pass testing with race detector enabled.
	//defer gock.Off()

	gock.New("https://honestbeehelp-sg.zendesk.com").
		Get("/api/v2/help_center/articles/search.json").
		Filter(func(req *http.Request) bool {
			return req.URL.Query().Get("query") == "order" && req.URL.Query().Get("locale") == "en-us"
		}).
		Reply(http.StatusOK).
		JSON(&zendesk.Search{
			Articles: []*zendesk.SearchArticle{
				{
					Article: &zendesk.Article{ID: 33456710,
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

	testCases := [...]struct {
		description string
		input       interface{}
		expect      interface{}
		expectErr   bool
	}{
		{
			description: "testing normal case",
			input: &inout.GetSearchIn{
				Query: "order",
				BaseIn: &inout.BaseIn{
					Locale:      "en-us",
					CountryCode: "sg",
				},
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
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := GetSearchHandler(context.Background(), e, tt.input)
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
