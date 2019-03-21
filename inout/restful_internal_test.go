package inout

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/go-test/deep"
)

func TestFetchBaseInParams(t *testing.T) {
	testCases := [...]struct {
		description string
		expect      *BaseIn
		expectErr   bool
		input       *http.Request
	}{
		{
			description: "testing normal case",
			expectErr:   false,
			expect: &BaseIn{
				Locale:      "en-us",
				CountryCode: "tw",
				PerPage:     3,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			input: &http.Request{
				Form: url.Values{
					"locale":       []string{"en-us"},
					"country_code": []string{"tw"},
					"per_page":     []string{"3"},
					"page":         []string{"1"},
					"sort_by":      []string{"position"},
					"sort_order":   []string{"asc"},
				},
			},
		},
		{
			description: "testing default case",
			expectErr:   false,
			expect: &BaseIn{
				Locale:      "en-us",
				CountryCode: "sg",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			input: &http.Request{
				Form: url.Values{},
			},
		},
		{
			description: "testing per page bigger than max case",
			expectErr:   false,
			expect: &BaseIn{
				Locale:      "en-us",
				CountryCode: "sg",
				PerPage:     100,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			input: &http.Request{
				Form: url.Values{
					"per_page": []string{"100000000"},
					"page":     []string{"1"},
				},
			},
		},
		{
			description: "testing per page error case",
			expectErr:   true,
			expect:      nil,
			input: &http.Request{
				Form: url.Values{
					"per_page": []string{"0"},
					"page":     []string{"1"},
				},
			},
		},
		{
			description: "testing page error case",
			expectErr:   true,
			expect:      nil,
			input: &http.Request{
				Form: url.Values{
					"per_page": []string{"10"},
					"page":     []string{"0"},
				},
			},
		},
		{
			description: "testing country code SG case",
			expectErr:   false,
			expect: &BaseIn{
				Locale:      "en-us",
				CountryCode: "sg",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			input: &http.Request{
				Form: url.Values{
					"country_code": []string{"sg"},
				},
			},
		},
		{
			description: "testing country code HK case",
			expectErr:   false,
			expect: &BaseIn{
				Locale:      "en-us",
				CountryCode: "hk",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			input: &http.Request{
				Form: url.Values{
					"country_code": []string{"hk"},
				},
			},
		},
		{
			description: "testing country code TW case",
			expectErr:   false,
			expect: &BaseIn{
				Locale:      "en-us",
				CountryCode: "tw",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			input: &http.Request{
				Form: url.Values{
					"country_code": []string{"tw"},
				},
			},
		},
		{
			description: "testing country code JP case",
			expectErr:   false,
			expect: &BaseIn{
				Locale:      "en-us",
				CountryCode: "jp",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			input: &http.Request{
				Form: url.Values{
					"country_code": []string{"jp"},
				},
			},
		},
		{
			description: "testing country code TH case",
			expectErr:   false,
			expect: &BaseIn{
				Locale:      "en-us",
				CountryCode: "th",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			input: &http.Request{
				Form: url.Values{
					"country_code": []string{"th"},
				},
			},
		},
		{
			description: "testing country code MY case",
			expectErr:   false,
			expect: &BaseIn{
				Locale:      "en-us",
				CountryCode: "my",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			input: &http.Request{
				Form: url.Values{
					"country_code": []string{"my"},
				},
			},
		},
		{
			description: "testing country code ID case",
			expectErr:   false,
			expect: &BaseIn{
				Locale:      "en-us",
				CountryCode: "id",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			input: &http.Request{
				Form: url.Values{
					"country_code": []string{"id"},
				},
			},
		},
		{
			description: "testing country code PH case",
			expectErr:   false,
			expect: &BaseIn{
				Locale:      "en-us",
				CountryCode: "ph",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			input: &http.Request{
				Form: url.Values{
					"country_code": []string{"ph"},
				},
			},
		},
		{
			description: "testing country code error case",
			expectErr:   true,
			expect:      nil,
			input: &http.Request{
				Form: url.Values{
					"country_code": []string{"no-this-country"},
				},
			},
		},
		{
			description: "testing locale ZH-TW case",
			expectErr:   false,
			expect: &BaseIn{
				Locale:      "zh-tw",
				CountryCode: "sg",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			input: &http.Request{
				Form: url.Values{
					"locale": []string{"zh-tw"},
				},
			},
		},
		{
			description: "testing locale ZH-CN case",
			expectErr:   false,
			expect: &BaseIn{
				Locale:      "zh-cn",
				CountryCode: "sg",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			input: &http.Request{
				Form: url.Values{
					"locale": []string{"zh-cn"},
				},
			},
		},
		{
			description: "testing locale EN-US case",
			expectErr:   false,
			expect: &BaseIn{
				Locale:      "en-us",
				CountryCode: "sg",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			input: &http.Request{
				Form: url.Values{
					"locale": []string{"en-us"},
				},
			},
		},
		{
			description: "testing locale JA case",
			expectErr:   false,
			expect: &BaseIn{
				Locale:      "ja",
				CountryCode: "sg",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			input: &http.Request{
				Form: url.Values{
					"locale": []string{"ja"},
				},
			},
		},
		{
			description: "testing locale TH case",
			expectErr:   false,
			expect: &BaseIn{
				Locale:      "th",
				CountryCode: "sg",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			input: &http.Request{
				Form: url.Values{
					"locale": []string{"th"},
				},
			},
		},
		{
			description: "testing locale ID case",
			expectErr:   false,
			expect: &BaseIn{
				Locale:      "id",
				CountryCode: "sg",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			input: &http.Request{
				Form: url.Values{
					"locale": []string{"id"},
				},
			},
		},
		{
			description: "testing locale error case",
			expectErr:   true,
			expect:      nil,
			input: &http.Request{
				Form: url.Values{
					"locale": []string{"no-this-locale"},
				},
			},
		},
		{
			description: "testing sort by position case",
			expectErr:   false,
			expect: &BaseIn{
				Locale:      "en-us",
				CountryCode: "sg",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			input: &http.Request{
				Form: url.Values{
					"sort_by": []string{"position"},
				},
			},
		},
		{
			description: "testing sort by created_at case",
			expectErr:   false,
			expect: &BaseIn{
				Locale:      "en-us",
				CountryCode: "sg",
				PerPage:     30,
				Page:        0,
				SortBy:      "created_at",
				SortOrder:   "asc",
			},
			input: &http.Request{
				Form: url.Values{
					"sort_by": []string{"created_at"},
				},
			},
		},
		{
			description: "testing sort by updated_at case",
			expectErr:   false,
			expect: &BaseIn{
				Locale:      "en-us",
				CountryCode: "sg",
				PerPage:     30,
				Page:        0,
				SortBy:      "updated_at",
				SortOrder:   "asc",
			},
			input: &http.Request{
				Form: url.Values{
					"sort_by": []string{"updated_at"},
				},
			},
		},
		{
			description: "testing sort by error case",
			expectErr:   true,
			expect:      nil,
			input: &http.Request{
				Form: url.Values{
					"sort_by": []string{"no-this-sort_by"},
				},
			},
		},
		{
			description: "testing sort order asc case",
			expectErr:   false,
			expect: &BaseIn{
				Locale:      "en-us",
				CountryCode: "sg",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			input: &http.Request{
				Form: url.Values{
					"sort_order": []string{"asc"},
				},
			},
		},
		{
			description: "testing sort order desc case",
			expectErr:   false,
			expect: &BaseIn{
				Locale:      "en-us",
				CountryCode: "sg",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "desc",
			},
			input: &http.Request{
				Form: url.Values{
					"sort_order": []string{"desc"},
				},
			},
		},
		{
			description: "testing sort order error case",
			expectErr:   true,
			expect:      nil,
			input: &http.Request{
				Form: url.Values{
					"sort_order": []string{"no-this-sort_order"},
				},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := FetchBaseParams(tt.input)
			if tt.expectErr && err == nil {
				t.Errorf("[%s] expect an error, actual == nil", tt.description)
			} else if diff := deep.Equal(tt.expect, actual); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}
