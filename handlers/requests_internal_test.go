package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/go-test/deep"
	"github.com/h2non/gock"

	"github.com/honestbee/Zen/errs"
	"github.com/honestbee/Zen/inout"
)

func TestCreateRequestDecompressor(t *testing.T) {
	normal := &inout.CreateRequestIn{
		CountryCode: "tw",
		Data: map[string]interface{}{
			"request": map[string]interface{}{
				"requester": map[string]interface{}{
					"name":  "zen project tester",
					"email": "zen.project.tester@honestbee.com",
				},
				"subject": "testing, please ignore",
				"comment": map[string]interface{}{
					"body": "testing, please ignore!!!",
				},
			},
		},
	}

	normalBytes, err := json.Marshal(normal)
	if err != nil {
		t.Fatalf("json marshal normal failed:%v", err)
	}

	testCases := [...]struct {
		description string
		input       *http.Request
		expect      interface{}
		expectErr   bool
	}{
		{
			description: "testing normal case",
			input: &http.Request{
				Body: ioutil.NopCloser(bytes.NewReader(normalBytes)),
			},
			expect:    normal,
			expectErr: false,
		},
		{
			description: "testing json decode failed case",
			input: &http.Request{
				Body: ioutil.NopCloser(bytes.NewBufferString("")),
			},
			expect:    nil,
			expectErr: true,
		},
		{
			description: "testing country code empty case",
			input: &http.Request{
				Body: ioutil.NopCloser(bytes.NewBufferString("{\"code\":1,\"msg\":\"\"}")),
			},
			expect:    nil,
			expectErr: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := CreateRequestDecompressor(nil, tt.input)
			if tt.expectErr && err == nil {
				t.Errorf("[%s] expect an error, actual == nil", tt.description)
			} else if diff := deep.Equal(tt.expect, actual); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}

func TestCreateRequestHandler(t *testing.T) {
	gock.New("https://honestbeehelp-tw.zendesk.com").
		Post("/api/v2/requests.json").
		BodyString(
			`{
				"request": {
					"comment": {
					  "body": "testing, please ignore!!!"
					},
					"requester": {
					  "name": "zen project tester",
					  "email": "zen.project.tester@honestbee.com"
					},
					"subject": "testing, please ignore"
				}
			}`,
		).
		Reply(http.StatusCreated)

	gock.New("https://honestbeehelp-tw.zendesk.com").
		Post("/api/v2/requests.json").
		BodyString(
			`{
				"request": {
					"requester": {
					  "name": "zen project tester",
					  "email": "zen.project.tester@honestbee.com"
					},
					"subject": "testing, please ignore"
				}
			}`,
		).
		Reply(http.StatusUnprocessableEntity)

	testCases := [...]struct {
		description   string
		input         interface{}
		expectErrCode int
	}{
		{
			description: "testing normal case",
			input: &inout.CreateRequestIn{
				CountryCode: "tw",
				Data: map[string]interface{}{
					"request": map[string]interface{}{
						"requester": map[string]interface{}{
							"name":  "zen project tester",
							"email": "zen.project.tester@honestbee.com",
						},
						"subject": "testing, please ignore",
						"comment": map[string]interface{}{
							"body": "testing, please ignore!!!",
						},
					},
				},
			},
			expectErrCode: http.StatusCreated,
		},
		{
			description:   "testing cast failed case",
			input:         map[string]interface{}{"cast": "failed"},
			expectErrCode: http.StatusInternalServerError,
		},
		{
			description: "testing zendesk create request failed case",
			input: &inout.CreateRequestIn{
				CountryCode: "tw",
				Data: map[string]interface{}{
					"requester": map[string]interface{}{
						"name":  "zen project tester",
						"email": "zen.project.tester@honestbee.com",
					},
					"subject": "testing, please ignore",
					"comment": map[string]interface{}{
						"body": "testing, please ignore!!!",
					},
				},
			},
			expectErrCode: http.StatusBadRequest,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			_, err := CreateRequestHandler(context.Background(), e, tt.input)
			if tt.expectErrCode != err.(*errs.Error).Status {
				t.Errorf("[%s] error code expect:%v, actual:%v", tt.description, tt.expectErrCode, err)
			}
		})
	}
}
