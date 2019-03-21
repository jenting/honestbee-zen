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

func TestSetVoteArticle(t *testing.T) {
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

	gock.New("https://honestbee-th.zendesk.com/").
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

	gock.New("https://honestbee-idn.zendesk.com").
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

	s := initServer()

	testCases := [...]struct {
		description string
		input       *protobuf.SetVoteArticleRequest
		expectErr   bool
		expect      *protobuf.SetVoteArticleResponse
	}{
		{
			description: "testing normal vote up case",
			input: &protobuf.SetVoteArticleRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				ArticleId:   "3345679",
				Vote:        protobuf.Vote_VOTE_UP,
			},
			expectErr: false,
			expect: &protobuf.SetVoteArticleResponse{
				Article: &protobuf.Article{
					Id:              "33456710",
					AuthorId:        "1234567",
					CommentsDisable: false,
					Draft:           false,
					Promoted:        false,
					Position:        0,
					VoteSum:         5,
					VoteCount:       7,
					CreatedAt:       models.FixCreatedAtProto1,
					UpdatedAt:       models.FixUpdatedAtProto1,
					SourceLocale:    "en-us",
					Outdated:        false,
					OutdatedLocales: []string{},
					EditedAt:        models.FixEditedAtProto1,
					LabelNames:      []string{},
					CountryCode:     "tw",
					Url:             "www.honestbee.com",
					HtmlUrl:         "www.honestbee.com",
					Name:            "testing article 1",
					Title:           "testing article 1",
					Body:            "this is testing article 1",
					Locale:          "en-us",
					SectionId:       "33456789",
				},
			},
		},
		{
			description: "testing normal vote down case",
			input: &protobuf.SetVoteArticleRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				ArticleId:   "3345678",
				Vote:        protobuf.Vote_VOTE_DOWN,
			},
			expectErr: false,
			expect: &protobuf.SetVoteArticleResponse{
				Article: &protobuf.Article{
					Id:              "33456710",
					AuthorId:        "1234567",
					CommentsDisable: false,
					Draft:           false,
					Promoted:        false,
					Position:        0,
					VoteSum:         -3,
					VoteCount:       7,
					CreatedAt:       models.FixCreatedAtProto1,
					UpdatedAt:       models.FixUpdatedAtProto1,
					SourceLocale:    "en-us",
					Outdated:        false,
					OutdatedLocales: []string{},
					EditedAt:        models.FixEditedAtProto1,
					LabelNames:      []string{},
					CountryCode:     "tw",
					Url:             "www.honestbee.com",
					HtmlUrl:         "www.honestbee.com",
					Name:            "testing article 1",
					Title:           "testing article 1",
					Body:            "this is testing article 1",
					Locale:          "en-us",
					SectionId:       "33456789",
				},
			},
		},
		{
			description: "testing invalid input case",
			input: &protobuf.SetVoteArticleRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				ArticleId:   "",
				Vote:        protobuf.Vote_VOTE_DOWN,
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing models return not found case",
			input: &protobuf.SetVoteArticleRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_ID,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				ArticleId:   "3345678",
				Vote:        protobuf.Vote_VOTE_DOWN,
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing models return error case",
			input: &protobuf.SetVoteArticleRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TH,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				ArticleId:   "3345679",
				Vote:        protobuf.Vote_VOTE_UP,
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing create vote error response case",
			input: &protobuf.SetVoteArticleRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Locale:      protobuf.Locale_LOCALE_EN_US,
				ArticleId:   "3345680",
				Vote:        protobuf.Vote_VOTE_UP,
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			resp, err := s.SetVoteArticle(context.Background(), tt.input)

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
