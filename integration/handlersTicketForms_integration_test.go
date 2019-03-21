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

func TestHandlersTicketForms(t *testing.T) {
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
			addr:        "/api/ticket_forms/951408",
			args: url.Values{
				"locale":       {"en-us"},
				"country_code": {"tw"},
			},
			expect: map[string]interface{}{
				"ticket_form": map[string]interface{}{
					"id":               951408,
					"name":             "951408 - My grocery bill is incorrect",
					"raw_name":         "951408 - My grocery bill is incorrect",
					"display_name":     "My grocery bill is incorrect",
					"raw_display_name": "My grocery bill is incorrect",
					"position":         1,
					"created_at":       time.Date(2017, 11, 30, 15, 45, 31, 0, time.UTC),
					"updated_at":       time.Date(2018, 1, 11, 5, 15, 20, 0, time.UTC),
					"ticket_fields": []map[string]interface{}{
						map[string]interface{}{
							"id":                  24681488,
							"type":                "subject",
							"title":               "Subject",
							"raw_title":           "Subject",
							"title_in_portal":     "Ticket form",
							"raw_title_in_portal": "Ticket form",
							"created_at":          time.Date(2015, 3, 2, 2, 47, 32, 0, time.UTC),
							"updated_at":          time.Date(2017, 12, 3, 10, 20, 4, 0, time.UTC),
						},
						map[string]interface{}{
							"id":                  24681498,
							"type":                "description",
							"title":               "Description",
							"raw_title":           "Description",
							"descript":            "Please elaborate on your request here.",
							"raw_descript":        "Please elaborate on your request here.",
							"title_in_portal":     "Description",
							"raw_title_in_portal": "Description",
							"created_at":          time.Date(2015, 3, 2, 2, 47, 32, 0, time.UTC),
							"updated_at":          time.Date(2017, 12, 1, 10, 49, 32, 0, time.UTC),
						},
						map[string]interface{}{
							"id":                  81469808,
							"type":                "text",
							"title":               "Order Number",
							"raw_title":           "Order Number",
							"position":            13,
							"title_in_portal":     "Order Number",
							"raw_title_in_portal": "Order Number",
							"created_at":          time.Date(2017, 11, 22, 13, 10, 41, 0, time.UTC),
							"updated_at":          time.Date(2018, 1, 4, 8, 41, 29, 0, time.UTC),
						},
					},
				},
			},
			expectStatus: http.StatusOK,
		},
		{
			description: "testing normal zh-tw locale case",
			addr:        "/api/ticket_forms/951408",
			args: url.Values{
				"locale":       {"zh-tw"},
				"country_code": {"tw"},
			},
			expect: map[string]interface{}{
				"ticket_form": map[string]interface{}{
					"id":               951408,
					"name":             "951408 - My grocery bill is incorrect",
					"raw_name":         "951408 - My grocery bill is incorrect",
					"display_name":     "My grocery bill is incorrect",
					"raw_display_name": "My grocery bill is incorrect",
					"position":         1,
					"created_at":       time.Date(2017, 11, 30, 15, 45, 31, 0, time.UTC),
					"updated_at":       time.Date(2018, 1, 11, 5, 15, 20, 0, time.UTC),
					"ticket_fields": []map[string]interface{}{
						map[string]interface{}{
							"id":                  24681488,
							"type":                "subject",
							"title":               "Subject",
							"raw_title":           "Subject",
							"title_in_portal":     "Ticket form",
							"raw_title_in_portal": "Ticket form",
							"created_at":          time.Date(2015, 3, 2, 2, 47, 32, 0, time.UTC),
							"updated_at":          time.Date(2017, 12, 3, 10, 20, 4, 0, time.UTC),
						},
						map[string]interface{}{
							"id":                  24681498,
							"type":                "description",
							"title":               "Description",
							"raw_title":           "Description",
							"descript":            "Please elaborate on your request here.",
							"raw_descript":        "Please elaborate on your request here.",
							"title_in_portal":     "Description",
							"raw_title_in_portal": "Description",
							"created_at":          time.Date(2015, 3, 2, 2, 47, 32, 0, time.UTC),
							"updated_at":          time.Date(2017, 12, 1, 10, 49, 32, 0, time.UTC),
						},
						map[string]interface{}{
							"id":                  81469808,
							"type":                "text",
							"title":               "Order Number",
							"raw_title":           "Order Number",
							"position":            13,
							"title_in_portal":     "Order Number",
							"raw_title_in_portal": "訂單號碼",
							"created_at":          time.Date(2017, 11, 22, 13, 10, 41, 0, time.UTC),
							"updated_at":          time.Date(2018, 1, 4, 8, 41, 29, 0, time.UTC),
						},
					},
				},
			},
			expectStatus: http.StatusOK,
		},
		{
			description: "testing not support locale goes default case",
			addr:        "/api/ticket_forms/951408",
			args: url.Values{
				"locale":       {"id"},
				"country_code": {"tw"},
			},
			expect: map[string]interface{}{
				"ticket_form": map[string]interface{}{
					"id":               951408,
					"name":             "951408 - My grocery bill is incorrect",
					"raw_name":         "951408 - My grocery bill is incorrect",
					"display_name":     "My grocery bill is incorrect",
					"raw_display_name": "My grocery bill is incorrect",
					"position":         1,
					"created_at":       time.Date(2017, 11, 30, 15, 45, 31, 0, time.UTC),
					"updated_at":       time.Date(2018, 1, 11, 5, 15, 20, 0, time.UTC),
					"ticket_fields": []map[string]interface{}{
						map[string]interface{}{
							"id":                  24681488,
							"type":                "subject",
							"title":               "Subject",
							"raw_title":           "Subject",
							"title_in_portal":     "Ticket form",
							"raw_title_in_portal": "Ticket form",
							"created_at":          time.Date(2015, 3, 2, 2, 47, 32, 0, time.UTC),
							"updated_at":          time.Date(2017, 12, 3, 10, 20, 4, 0, time.UTC),
						},
						map[string]interface{}{
							"id":                  24681498,
							"type":                "description",
							"title":               "Description",
							"raw_title":           "Description",
							"descript":            "Please elaborate on your request here.",
							"raw_descript":        "Please elaborate on your request here.",
							"title_in_portal":     "Description",
							"raw_title_in_portal": "Description",
							"created_at":          time.Date(2015, 3, 2, 2, 47, 32, 0, time.UTC),
							"updated_at":          time.Date(2017, 12, 1, 10, 49, 32, 0, time.UTC),
						},
						map[string]interface{}{
							"id":                  81469808,
							"type":                "text",
							"title":               "Order Number",
							"raw_title":           "Order Number",
							"position":            13,
							"title_in_portal":     "Order Number",
							"raw_title_in_portal": "Order Number",
							"created_at":          time.Date(2017, 11, 22, 13, 10, 41, 0, time.UTC),
							"updated_at":          time.Date(2018, 1, 4, 8, 41, 29, 0, time.UTC),
						},
					},
				},
			},
			expectStatus: http.StatusOK,
		},
		{
			description: "testing not exist form id case",
			addr:        "/api/ticket_forms/91293129847",
			args: url.Values{
				"locale":       {"id"},
				"country_code": {"tw"},
			},
			expect: map[string]interface{}{
				"error": errs.RecordNotFoundErrorMsg,
			},
			expectStatus: http.StatusNotFound,
		},
		{
			description: "testing wrong type form id case",
			addr:        "/api/ticket_forms/not-exist-form",
			args: url.Values{
				"locale":       {"id"},
				"country_code": {"tw"},
			},
			expect: map[string]interface{}{
				"error": errs.RecordNotFoundErrorMsg,
			},
			expectStatus: http.StatusNotFound,
		},
		{
			description: "testing not exist locale case",
			addr:        "/api/ticket_forms/951408",
			args: url.Values{
				"locale":       {"not-exist-locale"},
				"country_code": {"tw"},
			},
			expect: map[string]interface{}{
				"error": errs.InvalidAttributeErrorMsg,
			},
			expectStatus: http.StatusBadRequest,
		},
		{
			description: "testing not exist country code case",
			addr:        "/api/ticket_forms/951408",
			args: url.Values{
				"locale":       {"en-us"},
				"country_code": {"not-exist-country-code"},
			},
			expect: map[string]interface{}{
				"error": errs.InvalidAttributeErrorMsg,
			},
			expectStatus: http.StatusBadRequest,
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
