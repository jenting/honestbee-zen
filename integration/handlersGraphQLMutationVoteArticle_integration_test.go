// +build integration

package integration

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/go-test/deep"
	"gopkg.in/h2non/gock.v1"

	"github.com/honestbee/Zen/zendesk"
)

func TestHandlersGraphQLMutationVoteArticle(t *testing.T) {
	gock.New("https://honestbeehelp-tw.zendesk.com/").
		Post("/hc/en-us/articles/115015959188/vote").
		Filter(func(req *http.Request) bool { return req.PostFormValue("value") == "up" }).
		Reply(http.StatusOK).
		JSON(&zendesk.Vote{
			ID:          360002569612,
			VoteSum:     5,
			VoteCount:   7,
			UpvoteCount: 6,
			Label:       "6 out of 7 found this helpful",
			Value:       "up",
		})

	gock.New("https://honestbeehelp-tw.zendesk.com/").
		Post("/hc/en-us/articles/115015959188/vote").
		Filter(func(req *http.Request) bool { return req.PostFormValue("value") == "down" }).
		Reply(http.StatusOK).
		JSON(&zendesk.Vote{
			ID:          360002569612,
			VoteSum:     -3,
			VoteCount:   7,
			UpvoteCount: 2,
			Label:       "2 out of 7 found this helpful",
			Value:       "down",
		})

	ts := newTserver()
	defer ts.closeAll()
	testCases := []struct {
		description string
		body        map[string]interface{}
		query       string
		expectBody  map[string]interface{}
	}{
		{
			description: "testing normal tw + en-us + up case",
			body: map[string]interface{}{
				"query": `mutation
				{
					voteArticle(articleId: "115015959188", vote: UP, countryCode: TW, locale: EN_US) {
						id
						authorId
						commentsDisable
						draft
						promoted
						position
						voteSum
						voteCount
						createdAt
						updatedAt
						sourceLocale
						outdated
						outdatedLocales
						editedAt
						labelNames
						countryCode
						url
						htmlUrl
						name
						title
						body
						locale
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"data": map[string]interface{}{
					"voteArticle": map[string]interface{}{
						"id":              "115015959188",
						"authorId":        "24400224208",
						"commentsDisable": false,
						"draft":           false,
						"promoted":        false,
						"position":        0,
						"voteSum":         5,
						"voteCount":       7,
						"createdAt":       time.Date(2017, 12, 27, 3, 14, 16, 0, time.UTC),
						"updatedAt":       time.Date(2017, 12, 27, 3, 14, 48, 0, time.UTC),
						"sourceLocale":    "zh-tw",
						"outdated":        false,
						"outdatedLocales": []string{},
						"editedAt":        time.Date(2017, 12, 27, 3, 14, 16, 0, time.UTC),
						"labelNames":      []string{},
						"countryCode":     "tw",
						"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/articles/115015959188-What-can-I-do-when-my-cart-is-locked-.json",
						"htmlUrl":         "https://help.honestbee.tw/hc/en-us/articles/115015959188-What-can-I-do-when-my-cart-is-locked-",
						"name":            "What can I do when my cart is locked?",
						"title":           "What can I do when my cart is locked?",
						"body":            "<p>When there is an error completing your checkout,your cart will be temporarily locked to prevent further changes to your order. To unlock your cart,click ‘Yes,unlock my cart,’ when prompted.</p>",
						"locale":          "en-us",
					},
				},
			},
		},
		{
			description: "testing normal tw + en-us + down case",
			body: map[string]interface{}{
				"query": `mutation
				{
					voteArticle(articleId: "115015959188", vote: DOWN, countryCode: TW, locale: EN_US) {
						id
						authorId
						commentsDisable
						draft
						promoted
						position
						voteSum
						voteCount
						createdAt
						updatedAt
						sourceLocale
						outdated
						outdatedLocales
						editedAt
						labelNames
						countryCode
						url
						htmlUrl
						name
						title
						body
						locale
					}
				}
			`,
			},
			expectBody: map[string]interface{}{
				"data": map[string]interface{}{
					"voteArticle": map[string]interface{}{
						"id":              "115015959188",
						"authorId":        "24400224208",
						"commentsDisable": false,
						"draft":           false,
						"promoted":        false,
						"position":        0,
						"voteSum":         -3,
						"voteCount":       7,
						"createdAt":       time.Date(2017, 12, 27, 3, 14, 16, 0, time.UTC),
						"updatedAt":       time.Date(2017, 12, 27, 3, 14, 48, 0, time.UTC),
						"sourceLocale":    "zh-tw",
						"outdated":        false,
						"outdatedLocales": []string{},
						"editedAt":        time.Date(2017, 12, 27, 3, 14, 16, 0, time.UTC),
						"labelNames":      []string{},
						"countryCode":     "tw",
						"url":             "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/articles/115015959188-What-can-I-do-when-my-cart-is-locked-.json",
						"htmlUrl":         "https://help.honestbee.tw/hc/en-us/articles/115015959188-What-can-I-do-when-my-cart-is-locked-",
						"name":            "What can I do when my cart is locked?",
						"title":           "What can I do when my cart is locked?",
						"body":            "<p>When there is an error completing your checkout,your cart will be temporarily locked to prevent further changes to your order. To unlock your cart,click ‘Yes,unlock my cart,’ when prompted.</p>",
						"locale":          "en-us",
					},
				},
			},
		},
		{
			description: "testing not exist country code case",
			body: map[string]interface{}{
				"query": `mutation
				{
					voteArticle(articleId: "115015959188", vote: DOWN, countryCode: not_exist_country_code, locale: EN_US) {
						id
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"errors": []interface{}{
					map[string]interface{}{
						"message": "Argument \"countryCode\" has invalid value not_exist_country_code.\nExpected type \"CountryCode\", found not_exist_country_code.",
						"locations": []interface{}{
							map[string]interface{}{
								"line":   3,
								"column": 70,
							},
						},
					},
				},
			},
		},
		{
			description: "testing not exist locale case",
			body: map[string]interface{}{
				"query": `mutation
				{
					voteArticle(articleId: "115015959188", vote: DOWN, countryCode: SG, locale: not_exist_locale) {
						id
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"errors": []interface{}{
					map[string]interface{}{
						"message": "Argument \"locale\" has invalid value not_exist_locale.\nExpected type \"Locale\", found not_exist_locale.",
						"locations": []interface{}{
							map[string]interface{}{
								"line":   3,
								"column": 82,
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			// Send requests.
			b, err := json.Marshal(tt.body)
			if err != nil {
				t.Fatalf("[%s] json marshal failed:%v", tt.description, err)
			}
			resp, err := ts.Client().Post(ts.URL+"/graphql", "application/json", ioutil.NopCloser(bytes.NewReader(b)))
			if err != nil {
				t.Fatalf("[%s] http client get failed:%v", tt.description, err)
			}
			defer resp.Body.Close()

			// Compare HTTP status code.
			if http.StatusOK != resp.StatusCode {
				t.Errorf("[%s] http status expect:%v != actual:%v", tt.description, http.StatusOK, resp.StatusCode)
			}

			// Compare HTTP body.
			actual := make(map[string]interface{})
			if err = json.NewDecoder(resp.Body).Decode(&actual); err != nil {
				t.Fatalf("[%s] json decoding failed:%v", tt.description, err)
			}
			// Converts integer to the same type.
			expectData, err := json.Marshal(tt.expectBody)
			if err != nil {
				t.Fatalf("[%s] json marshal failed:%v", tt.description, err)
			}
			expect := make(map[string]interface{})
			if err = json.Unmarshal(expectData, &expect); err != nil {
				t.Fatalf("[%s] json unmarshal failed:%v", tt.description, err)
			}
			// Compares and prints difference.
			if diff := deep.Equal(expect, actual); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}
