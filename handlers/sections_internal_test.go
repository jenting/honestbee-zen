package handlers

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/go-test/deep"
	"github.com/julienschmidt/httprouter"

	"github.com/honestbee/Zen/inout"
	"github.com/honestbee/Zen/models"
)

func TestGetArticlesDecompressor(t *testing.T) {
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
					Key:   "section_id",
					Value: "3345679",
				},
			},
			input2: &http.Request{
				Form: url.Values{
					"locale":       []string{"en-us"},
					"country_code": []string{"tw"},
					"per_page":     []string{"3"},
					"page":         []string{"1"},
					"sort_by":      []string{"position"},
					"sort_order":   []string{"asc"},
				},
			},
			expectErr: false,
			expect: &inout.GetArticlesIn{
				SectionID: 3345679,
				BaseIn: &inout.BaseIn{
					Locale:      "en-us",
					CountryCode: "tw",
					PerPage:     3,
					Page:        0,
					SortBy:      "position",
					SortOrder:   "asc",
				},
			},
		},
		{
			description: "testing fetchBaseIn failed",
			input1:      nil,
			input2: &http.Request{
				Form: url.Values{
					"locale":       []string{"en-us"},
					"country_code": []string{"tw"},
					"per_page":     []string{"0"},
					"page":         []string{"1"},
					"sort_by":      []string{"position"},
					"sort_order":   []string{"asc"},
				},
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing parse section id failed",
			input1: httprouter.Params{
				httprouter.Param{
					Key:   "section_id",
					Value: "abcdefg",
				},
			},
			input2: &http.Request{
				Form: url.Values{
					"locale":       []string{"en-us"},
					"country_code": []string{"tw"},
					"per_page":     []string{"1"},
					"page":         []string{"1"},
					"sort_by":      []string{"position"},
					"sort_order":   []string{"asc"},
				},
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := GetArticlesDecompressor(tt.input1, tt.input2)
			if tt.expectErr && err == nil {
				t.Errorf("[%s] expect an error, actual == nil", tt.description)
			} else if diff := deep.Equal(tt.expect, actual); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}

func TestGetArticlesHandler(t *testing.T) {
	testCases := [...]struct {
		description string
		input       interface{}
		expect      interface{}
		expectErr   bool
	}{
		{
			description: "testing normal case",
			input: &inout.GetArticlesIn{
				BaseIn: &inout.BaseIn{
					Locale:      "en-us",
					CountryCode: "tw",
					PerPage:     3,
					Page:        0,
					SortBy:      "position",
					SortOrder:   "asc",
				},
				SectionID: 3345679,
			},
			expectErr: false,
			expect: &inout.GetArticlesOut{
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
					PerPage:   3,
					PageCount: 1,
					Count:     1,
				},
			},
		},
		{
			description: "testing input casting failed case",
			input: &struct {
				name string
				age  int
			}{
				name: "honestbee",
				age:  99,
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing models return error case",
			input: &inout.GetArticlesIn{
				BaseIn: &inout.BaseIn{
					Locale:      "en-us",
					CountryCode: models.ModelsReturnErrorCountryCode,
					PerPage:     3,
					Page:        0,
					SortBy:      "position",
					SortOrder:   "asc",
				},
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := GetArticlesHandler(context.Background(), e, tt.input)
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

func TestGetSectionDecompressor(t *testing.T) {
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
					Key:   "section_id",
					Value: "115003853607",
				},
			},
			input2: &http.Request{
				Form: url.Values{
					"locale":       []string{"en-us"},
					"country_code": []string{"tw"},
				},
			},
			expectErr: false,
			expect: &inout.GetSectionIn{
				SectionID: 115003853607,
				BaseIn: &inout.BaseIn{
					Locale:      "en-us",
					CountryCode: "tw",
					PerPage:     30,
					Page:        0,
					SortBy:      "position",
					SortOrder:   "asc",
				},
			},
		},
		{
			description: "testing fetchBaseIn failed",
			input1:      nil,
			input2: &http.Request{
				Form: url.Values{
					"locale":       []string{"en-us"},
					"country_code": []string{"tw"},
					"per_page":     []string{"0"},
					"page":         []string{"1"},
					"sort_by":      []string{"position"},
					"sort_order":   []string{"asc"},
				},
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing parse section id failed",
			input1: httprouter.Params{
				httprouter.Param{
					Key:   "section_id",
					Value: "abcdefg",
				},
			},
			input2: &http.Request{
				Form: url.Values{
					"locale":       []string{"en-us"},
					"country_code": []string{"tw"},
					"per_page":     []string{"1"},
					"page":         []string{"1"},
					"sort_by":      []string{"position"},
					"sort_order":   []string{"asc"},
				},
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := GetSectionDecompressor(tt.input1, tt.input2)
			if tt.expectErr && err == nil {
				t.Errorf("[%s] expect an error, actual == nil", tt.description)
			} else if diff := deep.Equal(tt.expect, actual); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}

func TestGetSectionHandler(t *testing.T) {
	testCases := [...]struct {
		description string
		input       interface{}
		expect      interface{}
		expectErr   bool
	}{
		{
			description: "testing normal case",
			input: &inout.GetSectionIn{
				BaseIn: &inout.BaseIn{
					Locale:      "en-us",
					CountryCode: "sg",
				},
				SectionID: 3345679,
			},
			expectErr: false,
			expect: &inout.GetSectionOut{
				Section: &models.Section{
					ID:           3345679,
					Position:     0,
					CreatedAt:    models.FixCreatedAt1,
					UpdatedAt:    models.FixUpdatedAt1,
					SourceLocale: "en-us",
					Outdated:     false,
					CountryCode:  "sg",
					URL:          "www.honestbee.com",
					HTMLURL:      "www.honestbee.com",
					Name:         "testing section 1",
					Description:  "",
					Locale:       "en-us",
					CategoryID:   3345678,
				},
			},
		},
		{
			description: "testing input casting failed case",
			input: &struct {
				name string
				age  int
			}{
				name: "honestbee",
				age:  99,
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing models return error case",
			input: &inout.GetArticlesIn{
				BaseIn: &inout.BaseIn{
					Locale:      "en-us",
					CountryCode: models.ModelsReturnErrorCountryCode,
					PerPage:     3,
					Page:        0,
					SortBy:      "position",
					SortOrder:   "asc",
				},
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := GetSectionHandler(context.Background(), e, tt.input)
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
