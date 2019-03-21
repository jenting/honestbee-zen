package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/go-test/deep"
	"github.com/h2non/gock"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/honestbee/Zen/config"
	"github.com/honestbee/Zen/errs"
	"github.com/honestbee/Zen/examiner"
	"github.com/honestbee/Zen/inout"
	"github.com/honestbee/Zen/models"
	"github.com/honestbee/Zen/zendesk"
)

var e *Env

func init() {
	logger := zerolog.New(ioutil.Discard)
	conf := &config.Config{
		HTTP: &config.HTTP{
			BasicAuthUser: "admin",
			BasicAuthPwd:  "33456783345678",
		},
	}
	ms := &models.MockModels{}
	zend, _ := zendesk.NewZenDesk(&config.Config{
		ZenDesk: &config.ZenDesk{
			RequestTimeoutSec: 20,
			AuthToken:         "ZWRtdW5kLmthb0Bob25lc3RiZWUuY29tL3Rva2VuOmZXdmVMYXVvN0lzQVExQURrbE54ZFVySkIwMWN1aFltTnhVRmVIbE8=",
			HKBaseURL:         "https://honestbeehelp-hk.zendesk.com",
			IDBaseURL:         "https://honestbee-idn.zendesk.com",
			JPBaseURL:         "https://honestbeehelp-jp.zendesk.com",
			MYBaseURL:         "https://honestbee-my.zendesk.com",
			PHBaseURL:         "https://honestbee-ph.zendesk.com",
			SGBaseURL:         "https://honestbeehelp-sg.zendesk.com",
			THBaseURL:         "https://honestbee-th.zendesk.com",
			TWBaseURL:         "https://honestbeehelp-tw.zendesk.com",
		},
	})
	exam, _ := examiner.NewExaminer(&config.Config{
		Examiner: &config.Examiner{
			MaxWorkerSize:          1,
			MaxPoolSize:            2,
			CategoriesRefreshLimit: 1000,
			SectionsRefreshLimit:   1000,
			ArticlesRefreshLimit:   1000,
		},
	}, &logger, ms, zend)
	e = &Env{
		Config:   conf,
		Logger:   &logger,
		Service:  ms,
		Examiner: exam,
		ZenDesk:  zend,
	}
}
func TestMiddleware(t *testing.T) {
	testCases := [...]struct {
		description   string
		dec           decompressor
		fn            handler
		expectStatus  int
		expectContent string
		expectBody    map[string]interface{}
	}{
		{
			description: "testing normal work flow",
			dec: func(httprouter.Params, *http.Request) (interface{}, error) {
				return nil, nil
			},
			fn: func(ctx context.Context, e *Env, in interface{}) (interface{}, error) {
				return map[string]interface{}{
					"name": "tester",
				}, nil
			},
			expectStatus:  http.StatusOK,
			expectContent: "application/json",
			expectBody: map[string]interface{}{
				"name": "tester",
			},
		},
		{
			description: "testing decompressor failed goes to 500 response",
			dec: func(httprouter.Params, *http.Request) (interface{}, error) {
				return nil, errors.New("decompressor failed")
			},
			fn: func(ctx context.Context, e *Env, in interface{}) (interface{}, error) {
				return map[string]interface{}{
					"name": "tester",
				}, nil
			},
			expectStatus:  http.StatusInternalServerError,
			expectContent: "application/json",
			expectBody: map[string]interface{}{
				"error": errs.ServerInternalErrorMsg,
			},
		},
		{
			description: "testing handler failed goes to 500 response",
			dec: func(httprouter.Params, *http.Request) (interface{}, error) {
				return nil, nil
			},
			fn: func(ctx context.Context, e *Env, in interface{}) (interface{}, error) {
				return nil, errors.New("handler failed")
			},
			expectStatus:  http.StatusInternalServerError,
			expectContent: "application/json",
			expectBody: map[string]interface{}{
				"error": errs.ServerInternalErrorMsg,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			m := Middleware(e, tt.dec, tt.fn)
			req := httptest.NewRequest(http.MethodGet, "http://fake.url.com", nil)
			w := httptest.NewRecorder()
			m(w, req, nil)

			resp := w.Result()
			actualBody := make(map[string]interface{})
			json.NewDecoder(resp.Body).Decode(&actualBody)

			actualStatus := resp.StatusCode
			if tt.expectStatus != actualStatus {
				t.Errorf("[%s] expectStatus:%v, actual:%v", tt.description, tt.expectStatus, actualStatus)
			}
			actualContent := resp.Header.Get("Content-Type")
			if tt.expectContent != actualContent {
				t.Errorf("[%s] expectContent:%v, actual:%v", tt.description, tt.expectContent, actualContent)
			}
			if diff := deep.Equal(tt.expectBody, actualBody); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}

func TestCreateRequestRequest(t *testing.T) {
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
		description  string
		dec          decompressor
		fn           handler
		inputParams  map[string]interface{}
		expectStatus int
		expectBody   interface{}
	}{
		{
			description: "testing normal case",
			dec:         CreateRequestDecompressor,
			fn:          CreateRequestHandler,
			inputParams: map[string]interface{}{
				"country_code": "tw",
				"data": map[string]interface{}{
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
			expectStatus: http.StatusCreated,
			expectBody: map[string]interface{}{
				"error": "",
			},
		},
		{
			description: "testing country code empty case",
			dec:         CreateRequestDecompressor,
			fn:          CreateRequestHandler,
			inputParams: map[string]interface{}{
				"country_code": "",
			},
			expectStatus: http.StatusBadRequest,
			expectBody: map[string]interface{}{
				"error": errs.InvalidAttributeErrorMsg,
			},
		},
		{
			description: "testing requesting failed case",
			dec:         CreateRequestDecompressor,
			fn:          CreateRequestHandler,
			inputParams: map[string]interface{}{
				"country_code": "tw",
				"data": map[string]interface{}{
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
			expectStatus: http.StatusBadRequest,
			expectBody: map[string]interface{}{
				"error": errs.InvalidAttributeErrorMsg,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			m := Middleware(e, tt.dec, tt.fn)
			pb, err := json.Marshal(tt.inputParams)
			if err != nil {
				t.Fatalf("json marshal failed:%v", err)
			}

			r := httptest.NewRequest(http.MethodPost, "http://fake.url.com", bytes.NewReader(pb))
			w := httptest.NewRecorder()
			m(w, r, nil)

			resp := w.Result()
			actualStatus := resp.StatusCode
			if tt.expectStatus != actualStatus {
				t.Errorf("[%s] expectStatus:%v, actual:%v", tt.description, tt.expectStatus, actualStatus)
			}

			var actualBody interface{}
			if actualStatus == http.StatusOK {
				actualBody = &inout.GetCategoriesOut{BaseOut: &inout.BaseOut{}}
			} else {
				actualBody = make(map[string]interface{})
			}

			json.NewDecoder(resp.Body).Decode(&actualBody)
			if diff := deep.Equal(tt.expectBody, actualBody); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}

func TestGetCategoriesRequest(t *testing.T) {
	testCases := [...]struct {
		description  string
		dec          decompressor
		fn           handler
		inputParams  url.Values
		expectStatus int
		expectBody   interface{}
	}{
		{
			description: "testing normal case",
			dec:         GetCategoriesDecompressor,
			fn:          GetCategoriesHandler,
			inputParams: url.Values{
				"locale": []string{"en-us"},
			},
			expectStatus: http.StatusOK,
			expectBody: &inout.GetCategoriesOut{
				Categories: []*models.Category{
					&models.Category{
						ID:           3345678,
						Position:     0,
						CreatedAt:    models.FixCreatedAt1,
						UpdatedAt:    models.FixUpdatedAt1,
						SourceLocale: "en-us",
						Outdated:     false,
						CountryCode:  "tw",
						URL:          "www.honestbee.com",
						HTMLURL:      "www.honestbee.com",
						Name:         "testing category 1",
						Description:  "",
						Locale:       "en-us",
						KeyName:      "food",
					},
				},
				BaseOut: &inout.BaseOut{
					Page:      1,
					PerPage:   30,
					PageCount: 1,
					Count:     1,
				},
			},
		},
		{
			description: "testing 400 bad request",
			dec:         GetCategoriesDecompressor,
			fn:          GetCategoriesHandler,
			inputParams: url.Values{
				"locale": []string{"no-exist-locale"},
			},
			expectStatus: http.StatusBadRequest,
			expectBody: map[string]interface{}{
				"error": errs.InvalidAttributeErrorMsg,
			},
		},
		{
			description: "testing 500 bad request",
			dec:         GetCategoriesDecompressor,
			fn:          GetCategoriesHandler,
			inputParams: url.Values{
				"country_code": []string{models.ModelsReturnErrorCountryCode},
			},
			expectStatus: http.StatusInternalServerError,
			expectBody: map[string]interface{}{
				"error": errs.ServerInternalErrorMsg,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			m := Middleware(e, tt.dec, tt.fn)
			r := httptest.NewRequest(http.MethodGet, "http://fake.url.com", nil)
			r.Form = tt.inputParams
			w := httptest.NewRecorder()
			m(w, r, nil)

			resp := w.Result()
			actualStatus := resp.StatusCode
			if tt.expectStatus != actualStatus {
				t.Errorf("[%s] expectStatus:%v, actual:%v", tt.description, tt.expectStatus, actualStatus)
			}

			var actualBody interface{}
			if actualStatus == http.StatusOK {
				actualBody = &inout.GetCategoriesOut{BaseOut: &inout.BaseOut{}}
			} else {
				actualBody = make(map[string]interface{})
			}

			json.NewDecoder(resp.Body).Decode(&actualBody)
			if diff := deep.Equal(tt.expectBody, actualBody); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}

func TestGetSectionsRequest(t *testing.T) {
	testCases := [...]struct {
		description  string
		dec          decompressor
		fn           handler
		inputParams  url.Values
		inputPS      httprouter.Params
		expectStatus int
		expectBody   interface{}
	}{
		{
			description: "testing normal case",
			dec:         GetSectionsDecompressor,
			fn:          GetSectionsHandler,
			inputParams: url.Values{
				"locale": []string{"en-us"},
			},
			inputPS: httprouter.Params{
				httprouter.Param{
					Key:   "category_id",
					Value: "3345678",
				},
			},
			expectStatus: http.StatusOK,
			expectBody: &inout.GetSectionsOut{
				Sections: []*models.Section{
					&models.Section{
						ID:           3345679,
						Position:     0,
						CreatedAt:    models.FixCreatedAt1,
						UpdatedAt:    models.FixUpdatedAt1,
						SourceLocale: "en-us",
						Outdated:     false,
						CountryCode:  "tw",
						URL:          "www.honestbee.com",
						HTMLURL:      "www.honestbee.com",
						Name:         "testing section 1",
						Description:  "",
						Locale:       "en-us",
						CategoryID:   3345678,
					},
				},
				BaseOut: &inout.BaseOut{
					Page:      1,
					PerPage:   30,
					PageCount: 1,
					Count:     1,
				},
			},
		},
		{
			description: "testing 400 bad request",
			dec:         GetSectionsDecompressor,
			fn:          GetSectionsHandler,
			inputParams: url.Values{
				"locale": []string{"no-exist-locale"},
			},
			expectStatus: http.StatusBadRequest,
			expectBody: map[string]interface{}{
				"error": errs.InvalidAttributeErrorMsg,
			},
		},
		{
			description: "testing 500 bad request",
			dec:         GetSectionsDecompressor,
			fn:          GetSectionsHandler,
			inputParams: url.Values{
				"country_code": []string{models.ModelsReturnErrorCountryCode},
			},
			inputPS: httprouter.Params{
				httprouter.Param{
					Key:   "category_id",
					Value: "3345678",
				},
			},
			expectStatus: http.StatusInternalServerError,
			expectBody: map[string]interface{}{
				"error": errs.ServerInternalErrorMsg,
			},
		},
		{
			description: "testing 404 not found",
			dec:         GetSectionsDecompressor,
			fn:          GetSectionsHandler,
			inputParams: url.Values{
				"country_code": []string{models.ModelsReturnErrorCountryCode},
			},
			expectStatus: http.StatusNotFound,
			expectBody: map[string]interface{}{
				"error": errs.RecordNotFoundErrorMsg,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			m := Middleware(e, tt.dec, tt.fn)
			r := httptest.NewRequest(http.MethodGet, "http://fake.url.com", nil)
			r.Form = tt.inputParams
			w := httptest.NewRecorder()
			m(w, r, tt.inputPS)

			resp := w.Result()
			actualStatus := resp.StatusCode
			if tt.expectStatus != actualStatus {
				t.Errorf("[%s] expectStatus:%v, actual:%v", tt.description, tt.expectStatus, actualStatus)
			}

			var actualBody interface{}
			if actualStatus == http.StatusOK {
				actualBody = &inout.GetSectionsOut{BaseOut: &inout.BaseOut{}}
			} else {
				actualBody = make(map[string]interface{})
			}

			json.NewDecoder(resp.Body).Decode(&actualBody)
			if diff := deep.Equal(tt.expectBody, actualBody); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}

func TestGetArticlesRequest(t *testing.T) {
	testCases := [...]struct {
		description  string
		dec          decompressor
		fn           handler
		inputParams  url.Values
		inputPS      httprouter.Params
		expectStatus int
		expectBody   interface{}
	}{
		{
			description: "testing normal case",
			dec:         GetArticlesDecompressor,
			fn:          GetArticlesHandler,
			inputParams: url.Values{
				"locale": []string{"en-us"},
			},
			inputPS: httprouter.Params{
				httprouter.Param{
					Key:   "section_id",
					Value: "3345679",
				},
			},
			expectStatus: http.StatusOK,
			expectBody: &inout.GetArticlesOut{
				Articles: []*models.Article{
					&models.Article{
						ID:              33456710,
						AuthorID:        1234567,
						CommentsDisable: false,
						Draft:           false,
						Promoted:        false,
						Position:        0,
						VoteSum:         0,
						VoteCount:       0,
						CreatedAt:       models.FixCreatedAt1,
						UpdatedAt:       models.FixUpdatedAt1,
						SourceLocale:    "en-us",
						Outdated:        false,
						OutdatedLocales: []string{},
						EditedAt:        models.FixEditedAt1,
						LabelNames:      []string{},
						CountryCode:     "tw",
						URL:             "www.honestbee.com",
						HTMLURL:         "www.honestbee.com",
						Name:            "testing article 1",
						Title:           "testing article 1",
						Body:            "this is testing article 1",
						Locale:          "en-us",
						SectionID:       33456789,
					},
				},
				BaseOut: &inout.BaseOut{
					Page:      1,
					PerPage:   30,
					PageCount: 1,
					Count:     1,
				},
			},
		},
		{
			description: "testing 400 bad request",
			dec:         GetArticlesDecompressor,
			fn:          GetArticlesHandler,
			inputParams: url.Values{
				"locale": []string{"no-exist-locale"},
			},
			expectStatus: http.StatusBadRequest,
			expectBody: map[string]interface{}{
				"error": errs.InvalidAttributeErrorMsg,
			},
		},
		{
			description: "testing 500 bad request",
			dec:         GetArticlesDecompressor,
			fn:          GetArticlesHandler,
			inputParams: url.Values{
				"country_code": []string{models.ModelsReturnErrorCountryCode},
			},
			inputPS: httprouter.Params{
				httprouter.Param{
					Key:   "section_id",
					Value: "3345679",
				},
			},
			expectStatus: http.StatusInternalServerError,
			expectBody: map[string]interface{}{
				"error": errs.ServerInternalErrorMsg,
			},
		},
		{
			description: "testing 404 not found",
			dec:         GetArticlesDecompressor,
			fn:          GetArticlesHandler,
			inputParams: url.Values{
				"locale": []string{"en-us"},
			},
			expectStatus: http.StatusNotFound,
			expectBody: map[string]interface{}{
				"error": errs.RecordNotFoundErrorMsg,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			m := Middleware(e, tt.dec, tt.fn)
			r := httptest.NewRequest(http.MethodGet, "http://fake.url.com", nil)
			r.Form = tt.inputParams
			w := httptest.NewRecorder()
			m(w, r, tt.inputPS)

			resp := w.Result()
			actualStatus := resp.StatusCode
			if tt.expectStatus != actualStatus {
				t.Errorf("[%s] expectStatus:%v, actual:%v", tt.description, tt.expectStatus, actualStatus)
			}

			var actualBody interface{}
			if actualStatus == http.StatusOK {
				actualBody = &inout.GetArticlesOut{BaseOut: &inout.BaseOut{}}
			} else {
				actualBody = make(map[string]interface{})
			}

			json.NewDecoder(resp.Body).Decode(&actualBody)
			if diff := deep.Equal(tt.expectBody, actualBody); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}

func TestGetArticleRequest(t *testing.T) {
	testCases := [...]struct {
		description  string
		dec          decompressor
		fn           handler
		inputParams  url.Values
		inputPS      httprouter.Params
		expectStatus int
		expectBody   interface{}
	}{
		{
			description: "testing normal case",
			dec:         GetArticleDecompressor,
			fn:          GetArticleHandler,
			inputParams: url.Values{
				"locale": []string{"en-us"},
			},
			inputPS: httprouter.Params{
				httprouter.Param{
					Key:   "article_id",
					Value: "33456710",
				},
			},
			expectStatus: http.StatusOK,
			expectBody: &inout.GetArticleOut{
				Article: &models.Article{
					ID:              33456710,
					AuthorID:        1234567,
					CommentsDisable: false,
					Draft:           false,
					Promoted:        false,
					Position:        0,
					VoteSum:         0,
					VoteCount:       0,
					CreatedAt:       models.FixCreatedAt1,
					UpdatedAt:       models.FixUpdatedAt1,
					SourceLocale:    "en-us",
					Outdated:        false,
					OutdatedLocales: []string{},
					EditedAt:        models.FixEditedAt1,
					LabelNames:      []string{},
					CountryCode:     "tw",
					URL:             "www.honestbee.com",
					HTMLURL:         "www.honestbee.com",
					Name:            "testing article 1",
					Title:           "testing article 1",
					Body:            "this is testing article 1",
					Locale:          "en-us",
					SectionID:       33456789,
				},
			},
		},
		{
			description: "testing 400 bad request",
			dec:         GetArticleDecompressor,
			fn:          GetArticleHandler,
			inputParams: url.Values{
				"locale": []string{"no-exist-locale"},
			},
			expectStatus: http.StatusBadRequest,
			expectBody: map[string]interface{}{
				"error": errs.InvalidAttributeErrorMsg,
			},
		},
		{
			description: "testing 500 bad request",
			dec:         GetArticleDecompressor,
			fn:          GetArticleHandler,
			inputParams: url.Values{
				"country_code": []string{models.ModelsReturnErrorCountryCode},
			},
			inputPS: httprouter.Params{
				httprouter.Param{
					Key:   "article_id",
					Value: "33456710",
				},
			},
			expectStatus: http.StatusInternalServerError,
			expectBody: map[string]interface{}{
				"error": errs.ServerInternalErrorMsg,
			},
		},
		{
			description: "testing 404 not found case 1 invalid article id",
			dec:         GetArticleDecompressor,
			fn:          GetArticleHandler,
			inputParams: url.Values{
				"locale": []string{"en-us"},
			},
			expectStatus: http.StatusNotFound,
			expectBody: map[string]interface{}{
				"error": errs.RecordNotFoundErrorMsg,
			},
		},
		{
			description: "testing 404 not found case 2 database not found",
			dec:         GetArticleDecompressor,
			fn:          GetArticleHandler,
			inputParams: url.Values{
				"country_code": []string{models.ModelsReturnNotFoundCountryCode},
			},
			inputPS: httprouter.Params{
				httprouter.Param{
					Key:   "article_id",
					Value: "33456710",
				},
			},
			expectStatus: http.StatusNotFound,
			expectBody: map[string]interface{}{
				"error": errs.RecordNotFoundErrorMsg,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			m := Middleware(e, tt.dec, tt.fn)
			r := httptest.NewRequest(http.MethodGet, "http://fake.url.com", nil)
			r.Form = tt.inputParams
			w := httptest.NewRecorder()
			m(w, r, tt.inputPS)

			resp := w.Result()
			actualStatus := resp.StatusCode
			if tt.expectStatus != actualStatus {
				t.Errorf("[%s] expectStatus:%v, actual:%v", tt.description, tt.expectStatus, actualStatus)
			}

			var actualBody interface{}
			if actualStatus == http.StatusOK {
				actualBody = new(inout.GetArticleOut)
			} else {
				actualBody = make(map[string]interface{})
			}

			json.NewDecoder(resp.Body).Decode(&actualBody)
			if diff := deep.Equal(tt.expectBody, actualBody); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}
