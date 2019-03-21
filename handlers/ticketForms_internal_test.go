package handlers

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/go-test/deep"
	"github.com/julienschmidt/httprouter"

	"github.com/honestbee/Zen/inout"
)

func TestGetTicketFormDecompressor(t *testing.T) {
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
					Key:   "form_id",
					Value: "3345678",
				},
			},
			input2: &http.Request{
				Form: url.Values{
					"locale":       []string{"en-us"},
					"country_code": []string{"tw"},
				},
			},
			expectErr: false,
			expect: &inout.GetTicketFormIn{
				CountryCode: "tw",
				Locale:      "en-us",
				FormID:      3345678,
			},
		},
		{
			description: "testing country code not exist case",
			input1: httprouter.Params{
				httprouter.Param{
					Key:   "form_id",
					Value: "3345678",
				},
			},
			input2: &http.Request{
				Form: url.Values{
					"locale":       []string{"en-us"},
					"country_code": []string{"not-exist-country-code"},
				},
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing locale not exist case",
			input1: httprouter.Params{
				httprouter.Param{
					Key:   "form_id",
					Value: "3345678",
				},
			},
			input2: &http.Request{
				Form: url.Values{
					"locale":       []string{"not-exist-locale"},
					"country_code": []string{"tw"},
				},
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing not valid form id case",
			input1: httprouter.Params{
				httprouter.Param{
					Key:   "form_id",
					Value: "abcdefghijk",
				},
			},
			input2: &http.Request{
				Form: url.Values{
					"locale":       []string{"zh-tw"},
					"country_code": []string{"tw"},
				},
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := GetTicketFormDecompressor(tt.input1, tt.input2)
			if tt.expectErr && err == nil {
				t.Errorf("[%s] expect an error, actual == nil", tt.description)
			} else if diff := deep.Equal(tt.expect, actual); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}
