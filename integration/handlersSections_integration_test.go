// +build integration

package integration

import (
	"encoding/json"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/go-test/deep"

	"github.com/honestbee/Zen/errs"
)

func TestHandlersSections(t *testing.T) {
	ts := newTserver()
	defer ts.closeAll()
	testCases := []struct {
		description  string
		args         url.Values
		addr         string
		expect       map[string]interface{}
		expectStatus int
	}{
		{
			description: "testing normal en-us locale case",
			addr:        "/api/categories/115002432448/sections",
			args: url.Values{
				"locale":       {"en-us"},
				"country_code": {"tw"},
			},
			expect: map[string]interface{}{
				"sections": []interface{}{
					map[string]interface{}{
						"category_id":   115002432448,
						"id":            115004118448,
						"position":      0,
						"created_at":    time.Date(2017, 12, 19, 6, 23, 48, 0, time.UTC),
						"updated_at":    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
						"source_locale": "en-us",
						"outdated":      false,
						"country_code":  "tw",
						"url":           "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/sections/115004118448-I-need-help-with-my-account.json",
						"html_url":      "https://help.honestbee.tw/hc/en-us/sections/115004118448-I-need-help-with-my-account",
						"name":          "I need help with my account",
						"description":   "",
						"locale":        "en-us",
					},
				},
				"page":       1,
				"per_page":   30,
				"page_count": 1,
				"count":      1,
			},
			expectStatus: http.StatusOK,
		},
		{
			description: "testing normal zh-tw locale case",
			addr:        "/api/categories/115002432448/sections",
			args: url.Values{
				"locale":       {"zh-tw"},
				"country_code": {"tw"},
			},
			expect: map[string]interface{}{
				"sections": []interface{}{
					map[string]interface{}{
						"category_id":   115002432448,
						"id":            115004118448,
						"position":      0,
						"created_at":    time.Date(2017, 12, 19, 6, 23, 48, 0, time.UTC),
						"updated_at":    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
						"source_locale": "en-us",
						"outdated":      false,
						"country_code":  "tw",
						"url":           "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/zh-tw/sections/115004118448-%E6%88%91%E9%9C%80%E8%A6%81%E5%B8%B3%E8%99%9F%E7%9B%B8%E9%97%9C%E7%9A%84%E5%8D%94%E5%8A%A9.json",
						"html_url":      "https://help.honestbee.tw/hc/zh-tw/sections/115004118448-%E6%88%91%E9%9C%80%E8%A6%81%E5%B8%B3%E8%99%9F%E7%9B%B8%E9%97%9C%E7%9A%84%E5%8D%94%E5%8A%A9",
						"name":          "我需要帳號相關的協助",
						"description":   "",
						"locale":        "zh-tw",
					},
				},
				"page":       1,
				"per_page":   30,
				"page_count": 1,
				"count":      1,
			},
			expectStatus: http.StatusOK,
		},
		{
			description: "testing not exist locale case",
			addr:        "/api/categories/115002432448/sections",
			args: url.Values{
				"locale":       {"not-exist-locale"},
				"country_code": {"tw"},
			},
			expectStatus: http.StatusBadRequest,
			expect: map[string]interface{}{
				"error": errs.InvalidAttributeErrorMsg,
			},
		},
		{
			description: "testing not exist country code case",
			addr:        "/api/categories/115002432448/sections",
			args: url.Values{
				"locale":       {"zh-tw"},
				"country_code": {"not-exist-country-code"},
			},
			expectStatus: http.StatusBadRequest,
			expect: map[string]interface{}{
				"error": errs.InvalidAttributeErrorMsg,
			},
		},
		{
			description: "testing not exist country code and locale case",
			addr:        "/api/categories/115002432448/sections",
			args: url.Values{
				"locale":       {"not-exist-locale"},
				"country_code": {"not-exist-country-code"},
			},
			expectStatus: http.StatusBadRequest,
			expect: map[string]interface{}{
				"error": errs.InvalidAttributeErrorMsg,
			},
		},
		{
			description: "testing not 0 per page case",
			addr:        "/api/categories/115002432448/sections",
			args: url.Values{
				"locale":       {"en-us"},
				"country_code": {"tw"},
				"per_page":     {"0"},
			},
			expectStatus: http.StatusBadRequest,
			expect: map[string]interface{}{
				"error": errs.InvalidAttributeErrorMsg,
			},
		},
		{
			description: "testing not 0 page case",
			addr:        "/api/categories/115002432448/sections",
			args: url.Values{
				"locale":       {"en-us"},
				"country_code": {"tw"},
				"page":         {"0"},
			},
			expectStatus: http.StatusBadRequest,
			expect: map[string]interface{}{
				"error": errs.InvalidAttributeErrorMsg,
			},
		},
		{
			description: "testing not 0 page and 0 per page case",
			addr:        "/api/categories/115002432448/sections",
			args: url.Values{
				"locale":       {"en-us"},
				"country_code": {"tw"},
				"page":         {"0"},
				"per_page":     {"0"},
			},
			expectStatus: http.StatusBadRequest,
			expect: map[string]interface{}{
				"error": errs.InvalidAttributeErrorMsg,
			},
		},
		{
			description: "testing sort by created at case",
			addr:        "/api/categories/115002432448/sections",
			args: url.Values{
				"locale":       {"en-us"},
				"country_code": {"tw"},
				"sort_by":      {"created_at"},
			},
			expectStatus: http.StatusOK,
			expect: map[string]interface{}{
				"sections": []interface{}{
					map[string]interface{}{
						"category_id":   115002432448,
						"id":            115004118448,
						"position":      0,
						"created_at":    time.Date(2017, 12, 19, 6, 23, 48, 0, time.UTC),
						"updated_at":    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
						"source_locale": "en-us",
						"outdated":      false,
						"country_code":  "tw",
						"url":           "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/sections/115004118448-I-need-help-with-my-account.json",
						"html_url":      "https://help.honestbee.tw/hc/en-us/sections/115004118448-I-need-help-with-my-account",
						"name":          "I need help with my account",
						"description":   "",
						"locale":        "en-us",
					},
				},
				"page":       1,
				"per_page":   30,
				"page_count": 1,
				"count":      1,
			},
		},
		{
			description: "testing sort by updated at case",
			addr:        "/api/categories/115002432448/sections",
			args: url.Values{
				"locale":       {"en-us"},
				"country_code": {"tw"},
				"sort_by":      {"updated_at"},
			},
			expectStatus: http.StatusOK,
			expect: map[string]interface{}{
				"sections": []interface{}{
					map[string]interface{}{
						"category_id":   115002432448,
						"id":            115004118448,
						"position":      0,
						"created_at":    time.Date(2017, 12, 19, 6, 23, 48, 0, time.UTC),
						"updated_at":    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
						"source_locale": "en-us",
						"outdated":      false,
						"country_code":  "tw",
						"url":           "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/sections/115004118448-I-need-help-with-my-account.json",
						"html_url":      "https://help.honestbee.tw/hc/en-us/sections/115004118448-I-need-help-with-my-account",
						"name":          "I need help with my account",
						"description":   "",
						"locale":        "en-us",
					},
				},
				"page":       1,
				"per_page":   30,
				"page_count": 1,
				"count":      1,
			},
		},
		{
			description: "testing sort by not correct case",
			addr:        "/api/categories/115002432448/sections",
			args: url.Values{
				"locale":       {"en-us"},
				"country_code": {"tw"},
				"sort_by":      {"unknown-sort-by"},
			},
			expectStatus: http.StatusBadRequest,
			expect: map[string]interface{}{
				"error": errs.InvalidAttributeErrorMsg,
			},
		},
		{
			description: "testing sort order desc case",
			addr:        "/api/categories/115002432448/sections",
			args: url.Values{
				"locale":       {"en-us"},
				"country_code": {"tw"},
				"sort_order":   {"desc"},
			},
			expectStatus: http.StatusOK,
			expect: map[string]interface{}{
				"sections": []interface{}{
					map[string]interface{}{
						"category_id":   115002432448,
						"id":            115004118448,
						"position":      0,
						"created_at":    time.Date(2017, 12, 19, 6, 23, 48, 0, time.UTC),
						"updated_at":    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
						"source_locale": "en-us",
						"outdated":      false,
						"country_code":  "tw",
						"url":           "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/sections/115004118448-I-need-help-with-my-account.json",
						"html_url":      "https://help.honestbee.tw/hc/en-us/sections/115004118448-I-need-help-with-my-account",
						"name":          "I need help with my account",
						"description":   "",
						"locale":        "en-us",
					},
				},
				"page":       1,
				"per_page":   30,
				"page_count": 1,
				"count":      1,
			},
		},
		{
			description: "testing sort order not correct case",
			addr:        "/api/categories/115002432448/sections",
			args: url.Values{
				"locale":       {"en-us"},
				"country_code": {"tw"},
				"sort_order":   {"unknown-sort-order"},
			},
			expectStatus: http.StatusBadRequest,
			expect: map[string]interface{}{
				"error": errs.InvalidAttributeErrorMsg,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			resp, err := ts.Client().Get(ts.URL + tt.addr + "?" + tt.args.Encode())
			if err != nil {
				t.Fatalf("[%s] http client do failed:%v", tt.description, err)
			}
			defer resp.Body.Close()

			if tt.expectStatus != resp.StatusCode {
				t.Errorf("[%s] http status expect:%v != actual:%v", tt.description, tt.expectStatus, resp.Status)
			}

			actual := make(map[string]interface{})
			if err = json.NewDecoder(resp.Body).Decode(&actual); err != nil {
				t.Fatalf("[%s] json decoding failed:%v", tt.description, err)
			}
			expectData, err := json.Marshal(tt.expect)
			if err != nil {
				t.Fatalf("[%s] json marshal failed:%v", tt.description, err)
			}
			expect := make(map[string]interface{})
			if err = json.Unmarshal(expectData, &expect); err != nil {
				t.Fatalf("[%s] json unmarshal failed:%v", tt.description, err)
			}
			if diff := deep.Equal(expect, actual); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}
