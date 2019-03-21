package inout

import (
	"testing"
)

func TestQueryCategoriesIn(t *testing.T) {
	testCases := [...]struct {
		description       string
		countryCode       string
		locale            string
		perPage           int32
		page              int32
		sortBy            string
		sortOrder         string
		expectCountryCode string
		expectLocale      string
		expectPerPage     int32
		expectPage        int32
		expectSortBy      string
		expectSortOrder   string
		expectErr         bool
	}{
		{
			description:       "normal case",
			countryCode:       "SG",
			locale:            "EN_US",
			perPage:           30,
			page:              1,
			sortBy:            "POSITION",
			sortOrder:         "ASC",
			expectCountryCode: "sg",
			expectLocale:      "en-us",
			expectPerPage:     30,
			expectPage:        0,
			expectSortBy:      "position",
			expectSortOrder:   "asc",
		},
		{
			description:     "no country code case",
			locale:          "EN_US",
			perPage:         30,
			page:            1,
			sortBy:          "POSITION",
			sortOrder:       "ASC",
			expectLocale:    "en-us",
			expectPerPage:   30,
			expectPage:      0,
			expectSortBy:    "position",
			expectSortOrder: "asc",
			expectErr:       true,
		},
		{
			description:       "no locale case",
			countryCode:       "SG",
			perPage:           30,
			page:              1,
			sortBy:            "POSITION",
			sortOrder:         "ASC",
			expectCountryCode: "sg",
			expectPerPage:     30,
			expectPage:        0,
			expectSortBy:      "position",
			expectSortOrder:   "asc",
			expectErr:         true,
		},
		{
			description:       "no sort by case",
			countryCode:       "SG",
			locale:            "EN_US",
			perPage:           30,
			page:              1,
			sortOrder:         "ASC",
			expectCountryCode: "sg",
			expectLocale:      "en-us",
			expectPerPage:     30,
			expectPage:        0,
			expectSortOrder:   "asc",
			expectErr:         true,
		},
		{
			description:       "no sort order case",
			countryCode:       "SG",
			locale:            "EN_US",
			perPage:           30,
			page:              1,
			sortBy:            "POSITION",
			expectCountryCode: "sg",
			expectLocale:      "en-us",
			expectPerPage:     30,
			expectPage:        0,
			expectSortBy:      "position",
			expectErr:         true,
		},
		{
			description:       "per page size > 100 case",
			countryCode:       "SG",
			locale:            "EN_US",
			perPage:           101,
			page:              1,
			sortBy:            "POSITION",
			sortOrder:         "ASC",
			expectCountryCode: "sg",
			expectLocale:      "en-us",
			expectPerPage:     100,
			expectPage:        0,
			expectSortBy:      "position",
			expectSortOrder:   "asc",
		},
		{
			description:       "per page size < 1 case",
			countryCode:       "SG",
			locale:            "EN_US",
			perPage:           0,
			page:              1,
			sortBy:            "POSITION",
			sortOrder:         "ASC",
			expectCountryCode: "sg",
			expectLocale:      "en-us",
			expectPerPage:     1,
			expectPage:        0,
			expectSortBy:      "position",
			expectSortOrder:   "asc",
		},
		{
			description:       "page size < 1 case",
			countryCode:       "SG",
			locale:            "EN_US",
			perPage:           30,
			page:              0,
			sortBy:            "POSITION",
			sortOrder:         "ASC",
			expectCountryCode: "sg",
			expectLocale:      "en-us",
			expectPerPage:     30,
			expectPage:        0,
			expectSortBy:      "position",
			expectSortOrder:   "asc",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			in := &QueryCategoriesIn{
				CountryCode: tt.countryCode,
				Locale:      tt.locale,
				PerPage:     tt.perPage,
				Page:        tt.page,
				SortBy:      tt.sortBy,
				SortOrder:   tt.sortOrder,
			}

			err := in.ProcessInputParams()
			if tt.expectErr {
				if err == nil {
					t.Errorf("[%s] expect an error, actual == nil", tt.description)
				}
			} else {
				if tt.expectCountryCode != in.CountryCode {
					t.Errorf("[%s] expect country code %s, actual %s", tt.description, tt.expectCountryCode, in.CountryCode)
				} else if tt.expectLocale != in.Locale {
					t.Errorf("[%s] expect locale %s, actual %s", tt.description, tt.expectLocale, in.Locale)
				} else if tt.expectPerPage != in.PerPage {
					t.Errorf("[%s] expect per page %d, actual %d", tt.description, tt.expectPerPage, in.PerPage)
				} else if tt.expectPage != in.Page {
					t.Errorf("[%s] expect page %d, actual %d", tt.description, tt.expectPage, in.Page)
				} else if tt.expectSortBy != in.SortBy {
					t.Errorf("[%s] expect sort by %s, actual %s", tt.description, tt.expectSortBy, in.SortBy)
				} else if tt.expectSortOrder != in.SortOrder {
					t.Errorf("[%s] expect sort order %s, actual %s", tt.description, tt.expectSortOrder, in.SortOrder)
				}
			}
		})
	}
}

func TestQueryCategoryIn(t *testing.T) {
	testCases := [...]struct {
		description       string
		countryCode       string
		locale            string
		expectCountryCode string
		expectLocale      string
		expectErr         bool
	}{
		{
			description:       "normal case",
			countryCode:       "SG",
			locale:            "EN_US",
			expectCountryCode: "sg",
			expectLocale:      "en-us",
		},
		{
			description:  "no country code case",
			locale:       "EN_US",
			expectLocale: "en-us",
			expectErr:    true,
		},
		{
			description:       "no locale case",
			countryCode:       "SG",
			expectCountryCode: "sg",
			expectErr:         true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			in := &QueryCategoryIn{
				CountryCode: tt.countryCode,
				Locale:      tt.locale,
			}

			err := in.ProcessInputParams()
			if tt.expectErr {
				if err == nil {
					t.Errorf("[%s] expect an error, actual == nil", tt.description)
				}
			} else {
				if tt.expectCountryCode != in.CountryCode {
					t.Errorf("[%s] expect country code %s, actual %s", tt.description, tt.expectCountryCode, in.CountryCode)
				} else if tt.expectLocale != in.Locale {
					t.Errorf("[%s] expect locale %s, actual %s", tt.description, tt.expectLocale, in.Locale)
				}
			}
		})
	}
}

func TestQuerySectionsIn(t *testing.T) {
	testCases := [...]struct {
		description       string
		countryCode       string
		locale            string
		perPage           int32
		page              int32
		sortBy            string
		sortOrder         string
		expectCountryCode string
		expectLocale      string
		expectPerPage     int32
		expectPage        int32
		expectSortBy      string
		expectSortOrder   string
		expectErr         bool
	}{
		{
			description:       "normal case",
			countryCode:       "TW",
			locale:            "ZH_TW",
			perPage:           30,
			page:              1,
			sortBy:            "CREATED_AT",
			sortOrder:         "DESC",
			expectCountryCode: "tw",
			expectLocale:      "zh-tw",
			expectPerPage:     30,
			expectPage:        0,
			expectSortBy:      "created_at",
			expectSortOrder:   "desc",
		},
		{
			description:     "no country code case",
			locale:          "ZH_TW",
			perPage:         30,
			page:            1,
			sortBy:          "CREATED_AT",
			sortOrder:       "DESC",
			expectLocale:    "zh-tw",
			expectPerPage:   30,
			expectPage:      0,
			expectSortBy:    "created_at",
			expectSortOrder: "desc",
			expectErr:       true,
		},
		{
			description:       "no locale case",
			countryCode:       "TW",
			perPage:           30,
			page:              1,
			sortBy:            "CREATED_AT",
			sortOrder:         "DESC",
			expectCountryCode: "tw",
			expectPerPage:     30,
			expectPage:        0,
			expectSortBy:      "created_at",
			expectSortOrder:   "desc",
			expectErr:         true,
		},
		{
			description:       "no sort by case",
			countryCode:       "TW",
			locale:            "ZH_TW",
			perPage:           30,
			page:              1,
			sortOrder:         "DESC",
			expectCountryCode: "tw",
			expectLocale:      "zh-tw",
			expectPerPage:     30,
			expectPage:        0,
			expectSortOrder:   "desc",
			expectErr:         true,
		},
		{
			description:       "no sort order case",
			countryCode:       "TW",
			locale:            "ZH_TW",
			perPage:           30,
			page:              1,
			sortBy:            "CREATED_AT",
			expectCountryCode: "tw",
			expectLocale:      "zh-tw",
			expectPerPage:     30,
			expectPage:        0,
			expectSortBy:      "created_at",
			expectErr:         true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			in := &QuerySectionsIn{
				CountryCode: tt.countryCode,
				Locale:      tt.locale,
				PerPage:     tt.perPage,
				Page:        tt.page,
				SortBy:      tt.sortBy,
				SortOrder:   tt.sortOrder,
			}

			err := in.ProcessInputParams()
			if tt.expectErr {
				if err == nil {
					t.Errorf("[%s] expect an error, actual == nil", tt.description)
				}
			} else {
				if tt.expectCountryCode != in.CountryCode {
					t.Errorf("[%s] expect country code %s, actual %s", tt.description, tt.expectCountryCode, in.CountryCode)
				} else if tt.expectLocale != in.Locale {
					t.Errorf("[%s] expect locale %s, actual %s", tt.description, tt.expectLocale, in.Locale)
				} else if tt.expectPerPage != in.PerPage {
					t.Errorf("[%s] expect per page %d, actual %d", tt.description, tt.expectPerPage, in.PerPage)
				} else if tt.expectPage != in.Page {
					t.Errorf("[%s] expect page %d, actual %d", tt.description, tt.expectPage, in.Page)
				} else if tt.expectSortBy != in.SortBy {
					t.Errorf("[%s] expect sort by %s, actual %s", tt.description, tt.expectSortBy, in.SortBy)
				} else if tt.expectSortOrder != in.SortOrder {
					t.Errorf("[%s] expect sort order %s, actual %s", tt.description, tt.expectSortOrder, in.SortOrder)
				}
			}
		})
	}
}

func TestQuerySectionIn(t *testing.T) {
	testCases := [...]struct {
		description       string
		countryCode       string
		locale            string
		expectCountryCode string
		expectLocale      string
		expectErr         bool
	}{
		{
			description:       "normal case",
			countryCode:       "ID",
			locale:            "ID",
			expectCountryCode: "id",
			expectLocale:      "id",
		},
		{
			description:       "no country code case",
			locale:            "ID",
			expectCountryCode: "id",
			expectLocale:      "id",
			expectErr:         true,
		},
		{
			description:       "no locale case",
			countryCode:       "ID",
			expectCountryCode: "id",
			expectLocale:      "id",
			expectErr:         true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			in := &QuerySectionIn{
				CountryCode: tt.countryCode,
				Locale:      tt.locale,
			}

			err := in.ProcessInputParams()
			if tt.expectErr {
				if err == nil {
					t.Errorf("[%s] expect an error, actual == nil", tt.description)
				}
			} else {
				if tt.expectCountryCode != in.CountryCode {
					t.Errorf("[%s] expect country code %s, actual %s", tt.description, tt.expectCountryCode, in.CountryCode)
				} else if tt.expectLocale != in.Locale {
					t.Errorf("[%s] expect locale %s, actual %s", tt.description, tt.expectLocale, in.Locale)
				}
			}
		})
	}
}

func TestQueryArticlesIn(t *testing.T) {
	testCases := [...]struct {
		description       string
		countryCode       string
		locale            string
		perPage           int32
		page              int32
		sortBy            string
		sortOrder         string
		expectCountryCode string
		expectLocale      string
		expectPerPage     int32
		expectPage        int32
		expectSortBy      string
		expectSortOrder   string
		expectErr         bool
	}{
		{
			description:       "normal case",
			countryCode:       "JP",
			locale:            "JA",
			perPage:           30,
			page:              1,
			sortBy:            "UPDATED_AT",
			sortOrder:         "DESC",
			expectCountryCode: "jp",
			expectLocale:      "ja",
			expectPerPage:     30,
			expectPage:        0,
			expectSortBy:      "updated_at",
			expectSortOrder:   "desc",
		},
		{
			description:     "no country code case",
			locale:          "JA",
			perPage:         30,
			page:            1,
			sortBy:          "UPDATED_AT",
			sortOrder:       "DESC",
			expectLocale:    "ja",
			expectPerPage:   30,
			expectPage:      0,
			expectSortBy:    "updated_at",
			expectSortOrder: "desc",
			expectErr:       true,
		},
		{
			description:       "no locale case",
			countryCode:       "JP",
			perPage:           30,
			page:              1,
			sortBy:            "UPDATED_AT",
			sortOrder:         "DESC",
			expectCountryCode: "jp",
			expectPerPage:     30,
			expectPage:        0,
			expectSortBy:      "updated_at",
			expectSortOrder:   "desc",
			expectErr:         true,
		},
		{
			description:       "no sort by case",
			countryCode:       "JP",
			locale:            "JA",
			perPage:           30,
			page:              1,
			sortOrder:         "DESC",
			expectCountryCode: "jp",
			expectLocale:      "ja",
			expectPerPage:     30,
			expectPage:        0,
			expectSortOrder:   "desc",
			expectErr:         true,
		},
		{
			description:       "no sort order case",
			countryCode:       "JP",
			locale:            "JA",
			perPage:           30,
			page:              1,
			sortBy:            "UPDATED_AT",
			expectCountryCode: "jp",
			expectLocale:      "ja",
			expectPerPage:     30,
			expectPage:        0,
			expectSortBy:      "updated_at",
			expectErr:         true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			in := &QueryArticlesIn{
				CountryCode: tt.countryCode,
				Locale:      tt.locale,
				PerPage:     tt.perPage,
				Page:        tt.page,
				SortBy:      tt.sortBy,
				SortOrder:   tt.sortOrder,
			}

			err := in.ProcessInputParams()
			if tt.expectErr {
				if err == nil {
					t.Errorf("[%s] expect an error, actual == nil", tt.description)
				}
			} else {
				if tt.expectCountryCode != in.CountryCode {
					t.Errorf("[%s] expect country code %s, actual %s", tt.description, tt.expectCountryCode, in.CountryCode)
				} else if tt.expectLocale != in.Locale {
					t.Errorf("[%s] expect locale %s, actual %s", tt.description, tt.expectLocale, in.Locale)
				} else if tt.expectPerPage != in.PerPage {
					t.Errorf("[%s] expect per page %d, actual %d", tt.description, tt.expectPerPage, in.PerPage)
				} else if tt.expectPage != in.Page {
					t.Errorf("[%s] expect page %d, actual %d", tt.description, tt.expectPage, in.Page)
				} else if tt.expectSortBy != in.SortBy {
					t.Errorf("[%s] expect sort by %s, actual %s", tt.description, tt.expectSortBy, in.SortBy)
				} else if tt.expectSortOrder != in.SortOrder {
					t.Errorf("[%s] expect sort order %s, actual %s", tt.description, tt.expectSortOrder, in.SortOrder)
				}
			}
		})
	}
}

func TestQueryTopArticlesIn(t *testing.T) {
	testCases := [...]struct {
		description       string
		countryCode       string
		locale            string
		expectCountryCode string
		expectLocale      string
		expectErr         bool
	}{
		{
			description:       "normal case",
			countryCode:       "TH",
			locale:            "TH",
			expectCountryCode: "th",
			expectLocale:      "th",
		},
		{
			description:  "no country code case",
			locale:       "TH",
			expectLocale: "th",
			expectErr:    true,
		},
		{
			description:       "no locale case",
			countryCode:       "TH",
			expectCountryCode: "th",
			expectErr:         true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			in := &QueryTopArticlesIn{
				CountryCode: tt.countryCode,
				Locale:      tt.locale,
			}

			err := in.ProcessInputParams()
			if tt.expectErr {
				if err == nil {
					t.Errorf("[%s] expect an error, actual == nil", tt.description)
				}
			} else {
				if tt.expectCountryCode != in.CountryCode {
					t.Errorf("[%s] expect country code %s, actual %s", tt.description, tt.expectCountryCode, in.CountryCode)
				} else if tt.expectLocale != in.Locale {
					t.Errorf("[%s] expect locale %s, actual %s", tt.description, tt.expectLocale, in.Locale)
				}
			}
		})
	}
}

func TestQueryArticleIn(t *testing.T) {
	testCases := [...]struct {
		description       string
		countryCode       string
		locale            string
		expectCountryCode string
		expectLocale      string
		expectErr         bool
	}{
		{
			description:       "normal case",
			countryCode:       "PH",
			locale:            "EN_US",
			expectCountryCode: "ph",
			expectLocale:      "en-us",
		},
		{
			description:  "no country code case",
			locale:       "EN_US",
			expectLocale: "en-us",
			expectErr:    true,
		},
		{
			description:       "no locale case",
			countryCode:       "PH",
			expectCountryCode: "ph",
			expectErr:         true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			in := &QueryArticleIn{
				CountryCode: tt.countryCode,
				Locale:      tt.locale,
			}

			err := in.ProcessInputParams()
			if tt.expectErr {
				if err == nil {
					t.Errorf("[%s] expect an error, actual == nil", tt.description)
				}
			} else {
				if tt.expectCountryCode != in.CountryCode {
					t.Errorf("[%s] expect country code %s, actual %s", tt.description, tt.expectCountryCode, in.CountryCode)
				} else if tt.expectLocale != in.Locale {
					t.Errorf("[%s] expect locale %s, actual %s", tt.description, tt.expectLocale, in.Locale)
				}
			}
		})
	}
}

func TestQueryTicketFieldsIn(t *testing.T) {
	testCases := [...]struct {
		description  string
		locale       string
		expectLocale string
		expectErr    bool
	}{
		{
			description:  "normal case",
			locale:       "EN_US",
			expectLocale: "en-us",
		},
		{
			description: "no locale case",
			expectErr:   true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			in := &QueryTicketFieldsIn{
				Locale: tt.locale,
			}

			err := in.ProcessInputParams()
			if tt.expectErr {
				if err == nil {
					t.Errorf("[%s] expect an error, actual == nil", tt.description)
				}
			} else {
				if tt.expectLocale != in.Locale {
					t.Errorf("[%s] expect locale %s, actual %s", tt.description, tt.expectLocale, in.Locale)
				}
			}
		})
	}
}

func TestQuerySearchTitleArticlesIn(t *testing.T) {
	testCases := [...]struct {
		description       string
		countryCode       string
		locale            string
		expectCountryCode string
		expectLocale      string
		expectErr         bool
	}{
		{
			description:       "normal case",
			countryCode:       "MY",
			locale:            "EN_US",
			expectCountryCode: "my",
			expectLocale:      "en-us",
		},
		{
			description:  "no country code case",
			locale:       "EN_US",
			expectLocale: "en-us",
			expectErr:    true,
		},
		{
			description:       "no locale case",
			countryCode:       "MY",
			expectCountryCode: "my",
			expectErr:         true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			in := &QuerySearchTitleArticlesIn{
				CountryCode: tt.countryCode,
				Locale:      tt.locale,
			}

			err := in.ProcessInputParams()
			if tt.expectErr {
				if err == nil {
					t.Errorf("[%s] expect an error, actual == nil", tt.description)
				}
			} else {
				if tt.expectCountryCode != in.CountryCode {
					t.Errorf("[%s] expect country code %s, actual %s", tt.description, tt.expectCountryCode, in.CountryCode)
				} else if tt.expectLocale != in.Locale {
					t.Errorf("[%s] expect locale %s, actual %s", tt.description, tt.expectLocale, in.Locale)
				}
			}
		})
	}
}

func TestQuerySearchBodyArticlesIn(t *testing.T) {
	testCases := [...]struct {
		description       string
		countryCode       string
		locale            string
		perPage           int32
		page              int32
		sortOrder         string
		expectCountryCode string
		expectLocale      string
		expectPerPage     int32
		expectPage        int32
		expectSortOrder   string
		expectErr         bool
	}{
		{
			description:       "normal case",
			countryCode:       "SG",
			locale:            "EN_US",
			perPage:           30,
			page:              1,
			sortOrder:         "ASC",
			expectCountryCode: "sg",
			expectLocale:      "en-us",
			expectPerPage:     30,
			expectPage:        0,
			expectSortOrder:   "asc",
		},
		{
			description:     "no country code case",
			locale:          "EN_US",
			perPage:         30,
			page:            1,
			sortOrder:       "ASC",
			expectLocale:    "en-us",
			expectPerPage:   30,
			expectPage:      0,
			expectSortOrder: "asc",
			expectErr:       true,
		},
		{
			description:       "no locale case",
			countryCode:       "SG",
			perPage:           30,
			page:              1,
			sortOrder:         "ASC",
			expectCountryCode: "sg",
			expectPerPage:     30,
			expectPage:        0,
			expectSortOrder:   "asc",
			expectErr:         true,
		},
		{
			description:       "no sort order case",
			countryCode:       "SG",
			locale:            "EN_US",
			perPage:           30,
			page:              1,
			expectCountryCode: "sg",
			expectLocale:      "en-us",
			expectPerPage:     30,
			expectPage:        0,
			expectErr:         true,
		},
		{
			description:       "per page size > 100 case",
			countryCode:       "SG",
			locale:            "EN_US",
			perPage:           101,
			page:              1,
			sortOrder:         "ASC",
			expectCountryCode: "sg",
			expectLocale:      "en-us",
			expectPerPage:     100,
			expectPage:        0,
			expectSortOrder:   "asc",
		},
		{
			description:       "per page size < 1 case",
			countryCode:       "SG",
			locale:            "EN_US",
			perPage:           0,
			page:              1,
			sortOrder:         "ASC",
			expectCountryCode: "sg",
			expectLocale:      "en-us",
			expectPerPage:     1,
			expectPage:        0,
			expectSortOrder:   "asc",
		},
		{
			description:       "page size < 1 case",
			countryCode:       "SG",
			locale:            "EN_US",
			perPage:           30,
			page:              0,
			sortOrder:         "ASC",
			expectCountryCode: "sg",
			expectLocale:      "en-us",
			expectPerPage:     30,
			expectPage:        0,
			expectSortOrder:   "asc",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			in := &QuerySearchBodyArticlesIn{
				CountryCode: tt.countryCode,
				Locale:      tt.locale,
				PerPage:     tt.perPage,
				Page:        tt.page,
				SortOrder:   tt.sortOrder,
			}

			err := in.ProcessInputParams()
			if tt.expectErr {
				if err == nil {
					t.Errorf("[%s] expect an error, actual == nil", tt.description)
				}
			} else {
				if tt.expectCountryCode != in.CountryCode {
					t.Errorf("[%s] expect country code %s, actual %s", tt.description, tt.expectCountryCode, in.CountryCode)
				} else if tt.expectLocale != in.Locale {
					t.Errorf("[%s] expect locale %s, actual %s", tt.description, tt.expectLocale, in.Locale)
				} else if tt.expectPerPage != in.PerPage {
					t.Errorf("[%s] expect per page %d, actual %d", tt.description, tt.expectPerPage, in.PerPage)
				} else if tt.expectPage != in.Page {
					t.Errorf("[%s] expect page %d, actual %d", tt.description, tt.expectPage, in.Page)
				} else if tt.expectSortOrder != in.SortOrder {
					t.Errorf("[%s] expect sort order %s, actual %s", tt.description, tt.expectSortOrder, in.SortOrder)
				}
			}
		})
	}
}

func TestMutationRequestsIn(t *testing.T) {
	testCases := [...]struct {
		description       string
		countryCode       string
		expectCountryCode string
		expectErr         bool
	}{
		{
			description:       "normal case",
			countryCode:       "MY",
			expectCountryCode: "my",
		},
		{
			description: "no country case",
			expectErr:   true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			in := &MutationRequestsIn{
				CountryCode: tt.countryCode,
			}

			err := in.ProcessInputParams()
			if tt.expectErr {
				if err == nil {
					t.Errorf("[%s] expect an error, actual == nil", tt.description)
				}
			} else {
				if tt.expectCountryCode != in.CountryCode {
					t.Errorf("[%s] expect country code %s, actual %s", tt.description, tt.expectCountryCode, in.CountryCode)
				}
			}
		})
	}
}

func TestMutationVoteArticleIn(t *testing.T) {
	testCases := [...]struct {
		description       string
		vote              string
		countryCode       string
		locale            string
		expectVote        string
		expectCountryCode string
		expectLocale      string
		expectErr         bool
	}{
		{
			description:       "normal vote up case",
			vote:              "UP",
			countryCode:       "ID",
			locale:            "ID",
			expectVote:        "up",
			expectCountryCode: "id",
			expectLocale:      "id",
		},
		{
			description:       "normal vote down case",
			vote:              "DOWN",
			countryCode:       "ID",
			locale:            "ID",
			expectVote:        "down",
			expectCountryCode: "id",
			expectLocale:      "id",
		},
		{
			description:       "no vote case",
			countryCode:       "ID",
			locale:            "ID",
			expectCountryCode: "id",
			expectLocale:      "id",
		},
		{
			description:  "no country code case",
			vote:         "UP",
			locale:       "ID",
			expectVote:   "up",
			expectLocale: "id",
			expectErr:    true,
		},
		{
			description:       "no locale case",
			vote:              "DOWN",
			countryCode:       "ID",
			expectVote:        "down",
			expectCountryCode: "id",

			expectErr: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			in := &MutationVoteArticleIn{
				Vote:        tt.vote,
				CountryCode: tt.countryCode,
				Locale:      tt.locale,
			}

			err := in.ProcessInputParams()
			if tt.expectErr {
				if err == nil {
					t.Errorf("[%s] expect an error, actual == nil", tt.description)
				}
			} else {
				if tt.expectVote != in.Vote {
					t.Errorf("[%s] expect vote %s, actual %s", tt.description, tt.expectVote, in.Vote)
				} else if tt.expectCountryCode != in.CountryCode {
					t.Errorf("[%s] expect country code %s, actual %s", tt.description, tt.expectCountryCode, in.CountryCode)
				} else if tt.expectLocale != in.Locale {
					t.Errorf("[%s] expect locale %s, actual %s", tt.description, tt.expectLocale, in.Locale)
				}
			}
		})
	}
}
