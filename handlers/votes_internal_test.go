package handlers

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/go-test/deep"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/h2non/gock.v1"

	"github.com/honestbee/Zen/inout"
	"github.com/honestbee/Zen/zendesk"
)

func TestCreateVoteDecompressor(t *testing.T) {
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
					Key:   "article_id",
					Value: "33456711",
				},
				httprouter.Param{
					Key:   "value",
					Value: "up",
				},
			},
			input2: &http.Request{
				Form: url.Values{
					"locale":       []string{"en-us"},
					"country_code": []string{"tw"},
				},
			},
			expectErr: false,
			expect: &inout.CreateVoteIn{
				ArticleID:   33456711,
				Value:       "up",
				Locale:      "en-us",
				CountryCode: "tw",
			},
		},
		{
			description: "testing fetchBaseIn failed",
			input1:      nil,
			input2: &http.Request{
				Form: url.Values{
					"locale":       []string{"en-us"},
					"country_code": []string{"gg"},
				},
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing parse article id failed",
			input1: httprouter.Params{
				httprouter.Param{
					Key:   "article_id",
					Value: "fake",
				},
			},
			input2: &http.Request{
				Form: url.Values{
					"locale":       []string{"en-us"},
					"country_code": []string{"tw"},
				},
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing parse value key failed",
			input1: httprouter.Params{
				httprouter.Param{
					Key:   "article_id",
					Value: "33456711",
				},
				httprouter.Param{
					Key:   "value",
					Value: "fake",
				},
			},
			input2: &http.Request{
				Form: url.Values{
					"locale":       []string{"en-us"},
					"country_code": []string{"tw"},
				},
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := CreateVoteDecompressor(tt.input1, tt.input2)
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

func TestCreateVoteHandler(t *testing.T) {
	// Race condition happens when gock change back HTTP DefaultTransport after
	// TestCreateVoteHandler function, but defer e.Examiner.SyncArticle have not
	// finish. We remove gock.Off() to pass testing with race detector enabled.
	//defer gock.Off()

	gock.New("https://honestbeehelp-tw.zendesk.com/").
		Post("/hc/en-us/articles/3345679/vote").
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
		Post("/hc/en-us/articles/3345678/vote").
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

	testCases := [...]struct {
		description string
		input       interface{}
		expect      interface{}
		expectErr   bool
	}{
		{
			description: "testing normal case, vote up",
			input: &inout.CreateVoteIn{
				Locale:      "en-us",
				CountryCode: "tw",
				ArticleID:   3345679,
				Value:       "up",
			},
			expectErr: false,
			expect: &inout.CreateVoteOut{
				VoteSum:   5,
				VoteCount: 7,
			},
		},
		{
			description: "testing normal case, vote down",
			input: &inout.CreateVoteIn{
				Locale:      "en-us",
				CountryCode: "tw",
				ArticleID:   3345678,
				Value:       "down",
			},
			expectErr: false,
			expect: &inout.CreateVoteOut{
				VoteSum:   -3,
				VoteCount: 7,
			},
		},
		{
			description: "testing article not found error case",
			input: &inout.CreateVoteIn{
				Locale:      "en-us",
				CountryCode: "tw",
				ArticleID:   111,
				Value:       "down",
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := CreateVoteHandler(context.Background(), e, tt.input)
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
