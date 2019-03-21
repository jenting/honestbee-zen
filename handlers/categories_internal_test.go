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

func TestGetSectionsDecompressor(t *testing.T) {
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
					Key:   "category_id",
					Value: "3345678",
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
			expect: &inout.GetSectionsIn{
				CategoryID: 3345678,
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
			description: "testing parse category id failed",
			input1: httprouter.Params{
				httprouter.Param{
					Key:   "category_id",
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
			actual, err := GetSectionsDecompressor(tt.input1, tt.input2)
			if tt.expectErr && err == nil {
				t.Errorf("[%s] expect an error, actual == nil", tt.description)
			} else if diff := deep.Equal(tt.expect, actual); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}

func TestGetSectionsHandler(t *testing.T) {
	testCases := [...]struct {
		description string
		input       interface{}
		expect      interface{}
		expectErr   bool
	}{
		{
			description: "testing normal case",
			input: &inout.GetSectionsIn{
				BaseIn: &inout.BaseIn{
					Locale:      "en-us",
					CountryCode: "tw",
					PerPage:     3,
					Page:        0,
					SortBy:      "position",
					SortOrder:   "asc",
				},
				CategoryID: 3345678,
			},
			expectErr: false,
			expect: &inout.GetSectionsOut{
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
			input: &inout.GetSectionsIn{
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
			actual, err := GetSectionsHandler(context.Background(), e, tt.input)
			if tt.expectErr && err == nil {
				t.Errorf("[%s] expect an error, actual nil", tt.description)
			} else if diff := deep.Equal(tt.expect, actual); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}

func TestGetCategoryKeyNameToIDDecompressor(t *testing.T) {
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
					Key:   "category_key_name",
					Value: "groceries",
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
			expect: &inout.GetCategoryKeyNameToIDIn{
				CategoryKeyName: "groceries",
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
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := GetCategoryKeyNameToIDDecompressor(tt.input1, tt.input2)
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

func TestGetCategoryKeyNameToIDHandler(t *testing.T) {
	testCases := [...]struct {
		description string
		input       interface{}
		expect      interface{}
		expectErr   bool
	}{
		{
			description: "testing normal case",
			input: &inout.GetCategoryKeyNameToIDIn{
				BaseIn: &inout.BaseIn{
					Locale:      "en-us",
					CountryCode: "tw",
					PerPage:     3,
					Page:        0,
					SortBy:      "position",
					SortOrder:   "asc",
				},
				CategoryKeyName: "groceries",
			},
			expectErr: false,
			expect: &inout.GetCategoryKeyNameToIDOut{
				CategoryID: 3345678,
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
			input: &inout.GetCategoryKeyNameToIDIn{
				BaseIn: &inout.BaseIn{
					Locale:      "en-us",
					CountryCode: "tw",
					PerPage:     3,
					Page:        0,
					SortBy:      "position",
					SortOrder:   "asc",
				},
				CategoryKeyName: "hahaha",
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing models return error case",
			input: &inout.GetCategoryKeyNameToIDIn{
				BaseIn: &inout.BaseIn{
					Locale:      "en-us",
					CountryCode: models.ModelsReturnErrorCountryCode,
					PerPage:     3,
					Page:        0,
					SortBy:      "position",
					SortOrder:   "asc",
				},
				CategoryKeyName: "fekjfkelwjf23m",
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := GetCategoryKeyNameToIDHandler(context.Background(), e, tt.input)
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

func TestGetCategoriesDecompressor(t *testing.T) {
	testCases := [...]struct {
		description string
		input1      httprouter.Params
		input2      *http.Request
		expect      interface{}
		expectErr   bool
	}{
		{
			description: "testing normal case",
			input1:      nil,
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
			expect: &inout.GetCategoriesIn{
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
			description: "testing error case 1",
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
			description: "testing error case 2",
			input1:      nil,
			input2: &http.Request{
				Form: url.Values{
					"locale":       []string{"en-us"},
					"country_code": []string{"tw"},
					"per_page":     []string{"1"},
					"page":         []string{"-1"},
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
			actual, err := GetCategoriesDecompressor(tt.input1, tt.input2)
			if tt.expectErr && err == nil {
				t.Errorf("[%s] expect an error, actual == nil", tt.description)
			} else if diff := deep.Equal(tt.expect, actual); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}

func TestGetCategoriesHandler(t *testing.T) {
	testCases := [...]struct {
		description string
		input       interface{}
		expect      interface{}
		expectErr   bool
	}{
		{
			description: "testing normal case",
			input: &inout.GetCategoriesIn{
				BaseIn: &inout.BaseIn{
					Locale:      "en-us",
					CountryCode: "tw",
					PerPage:     3,
					Page:        0,
					SortBy:      "position",
					SortOrder:   "asc",
				},
			},
			expectErr: false,
			expect: &inout.GetCategoriesOut{
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
			input: &inout.GetCategoriesIn{
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
			actual, err := GetCategoriesHandler(context.Background(), e, tt.input)
			if tt.expectErr && err == nil {
				t.Errorf("[%s] expect an error, actual nil", tt.description)
			} else if diff := deep.Equal(tt.expect, actual); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}

func TestGetCategoriesArticlesDecompressor(t *testing.T) {
	testCases := [...]struct {
		description string
		input1      httprouter.Params
		input2      *http.Request
		expect      interface{}
		expectErr   bool
	}{
		{
			description: "testing normal case 1",
			input1: httprouter.Params{
				httprouter.Param{
					Key:   "category_id",
					Value: "1234",
				},
			},
			input2: &http.Request{
				Form: url.Values{
					"locale":       []string{"en-us"},
					"country_code": []string{"sg"},
					"per_page":     []string{"3"},
					"page":         []string{"1"},
					"sort_by":      []string{"position"},
					"sort_order":   []string{"asc"},
					"label_names":  []string{""},
				},
			},
			expectErr: false,
			expect: &inout.GetCategoriesArticlesIn{
				CategoryID: 1234,
				LabelNames: "",
				BaseIn: &inout.BaseIn{
					Locale:      "en-us",
					CountryCode: "sg",
					PerPage:     3,
					Page:        0,
					SortBy:      "position",
					SortOrder:   "asc",
				},
			},
		},
		{
			description: "testing normal case 2",
			input1: httprouter.Params{
				httprouter.Param{
					Key:   "category_id",
					Value: "5678",
				},
			},
			input2: &http.Request{
				Form: url.Values{
					"locale":       []string{"en-us"},
					"country_code": []string{"sg"},
					"per_page":     []string{"3"},
					"page":         []string{"1"},
					"sort_by":      []string{"position"},
					"sort_order":   []string{"asc"},
					"label_names":  []string{"confirmed,preparing,ontheway,delivered"},
				},
			},
			expectErr: false,
			expect: &inout.GetCategoriesArticlesIn{
				CategoryID: 5678,
				LabelNames: "confirmed,preparing,ontheway,delivered",
				BaseIn: &inout.BaseIn{
					Locale:      "en-us",
					CountryCode: "sg",
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
			description: "testing parse category id failed",
			input1: httprouter.Params{
				httprouter.Param{
					Key:   "category_id",
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
			actual, err := GetCategoriesArticlesDecompressor(tt.input1, tt.input2)
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

func TestGetCategoriesArticlesHandler(t *testing.T) {
	testCases := [...]struct {
		description string
		input       interface{}
		expect      interface{}
		expectErr   bool
	}{
		{
			description: "testing normal case without labels",
			input: &inout.GetCategoriesArticlesIn{
				BaseIn: &inout.BaseIn{
					Locale:      "en-us",
					CountryCode: "tw",
					PerPage:     5,
					Page:        0,
					SortBy:      "position",
					SortOrder:   "asc",
				},
				CategoryID: 3345678,
				LabelNames: "",
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
						LabelNames:      []string{"confirmed"},
						CountryCode:     "tw",
						URL:             "www.honestbee.com",
						HTMLURL:         "www.honestbee.com",
						Name:            "testing article 1",
						Title:           "testing article 1",
						Body:            "this is testing article 1",
						Locale:          "en-us",
						SectionID:       33456789,
					},
					&models.Article{
						ID:              33456711,
						AuthorID:        1234568,
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
						LabelNames:      []string{"preparing"},
						CountryCode:     "tw",
						URL:             "www.honestbee.com",
						HTMLURL:         "www.honestbee.com",
						Name:            "testing article 2",
						Title:           "testing article 2",
						Body:            "this is testing article 2",
						Locale:          "en-us",
						SectionID:       33456789,
					},
					&models.Article{
						ID:              33456712,
						AuthorID:        1234569,
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
						LabelNames:      []string{"ontheway"},
						CountryCode:     "tw",
						URL:             "www.honestbee.com",
						HTMLURL:         "www.honestbee.com",
						Name:            "testing article 3",
						Title:           "testing article 3",
						Body:            "this is testing article 3",
						Locale:          "en-us",
						SectionID:       33456789,
					},
					&models.Article{
						ID:              33456713,
						AuthorID:        1234570,
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
						LabelNames:      []string{"delivered"},
						CountryCode:     "tw",
						URL:             "www.honestbee.com",
						HTMLURL:         "www.honestbee.com",
						Name:            "testing article 4",
						Title:           "testing article 4",
						Body:            "this is testing article 4",
						Locale:          "en-us",
						SectionID:       33456789,
					},
				},
				BaseOut: &inout.BaseOut{
					Page:      1,
					PerPage:   5,
					PageCount: 1,
					Count:     4,
				},
			},
		},
		{
			description: "testing normal case with labels",
			input: &inout.GetCategoriesArticlesIn{
				BaseIn: &inout.BaseIn{
					Locale:      "en-us",
					CountryCode: "tw",
					PerPage:     5,
					Page:        0,
					SortBy:      "position",
					SortOrder:   "asc",
				},
				CategoryID: 3345678,
				LabelNames: "confirmed",
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
						LabelNames:      []string{"confirmed"},
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
					PerPage:   5,
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
			input: &inout.GetCategoriesArticlesIn{
				BaseIn: &inout.BaseIn{
					Locale:      "en-us",
					CountryCode: models.ModelsReturnErrorCountryCode,
					PerPage:     3,
					Page:        0,
					SortBy:      "position",
					SortOrder:   "asc",
				},
				CategoryID: 3345678,
				LabelNames: "",
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := GetCategoriesArticlesHandler(context.Background(), e, tt.input)
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
