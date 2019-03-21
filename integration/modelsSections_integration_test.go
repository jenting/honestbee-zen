// +build integration

package integration

import (
	"context"
	"testing"
	"time"

	"github.com/go-test/deep"

	"github.com/honestbee/Zen/models"
)

func TestModelsSyncWithSections(t *testing.T) {
	service := newService()
	defer service.Close()
	testCases := []struct {
		description            string
		inputSections          []*models.Section
		inputGetSectionsParams *models.GetSectionsParams
		inputCountryCode       string
		inputLocale            string
		expectSections         []*models.Section
		expectCount            int
		expectError            bool
	}{
		{
			description: "testing sync with en-us locale one mock section case",
			inputSections: []*models.Section{
				&models.Section{
					CategoryID:   1234567,
					ID:           3345678,
					Position:     0,
					CreatedAt:    time.Date(1988, 10, 13, 3, 30, 59, 0, time.UTC),
					UpdatedAt:    time.Date(1988, 10, 13, 3, 30, 59, 0, time.UTC),
					SourceLocale: "en-us",
					Outdated:     false,
					URL:          "https://sync.test.1",
					HTMLURL:      "https://sync.test.1",
					Name:         "sync test 1",
					Description:  "",
					Locale:       "en-us",
					CountryCode:  "tw",
				},
			},
			inputGetSectionsParams: &models.GetSectionsParams{
				CategoryID:  1234567,
				Locale:      "en-us",
				CountryCode: "tw",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			inputLocale:      "en-us",
			inputCountryCode: "tw",
			expectSections: []*models.Section{
				&models.Section{
					CategoryID:   1234567,
					ID:           3345678,
					Position:     0,
					CreatedAt:    time.Date(1988, 10, 13, 3, 30, 59, 0, time.UTC),
					UpdatedAt:    time.Date(1988, 10, 13, 3, 30, 59, 0, time.UTC),
					SourceLocale: "en-us",
					Outdated:     false,
					URL:          "https://sync.test.1",
					HTMLURL:      "https://sync.test.1",
					Name:         "sync test 1",
					Description:  "",
					Locale:       "en-us",
					CountryCode:  "tw",
				},
			},
			expectCount: 1,
		},
		{
			description: "testing sync with all mock section case",
			inputSections: []*models.Section{
				&models.Section{
					CategoryID:   7654321,
					ID:           3345678,
					Position:     101,
					CreatedAt:    time.Date(1999, 10, 13, 3, 30, 59, 0, time.UTC),
					UpdatedAt:    time.Date(1999, 10, 13, 3, 30, 59, 0, time.UTC),
					SourceLocale: "zh-cn",
					Outdated:     true,
					URL:          "https://sync.test.1",
					HTMLURL:      "https://sync.test.1",
					Name:         "同步測試 1",
					Description:  "",
					Locale:       "zh-tw",
					CountryCode:  "tw",
				},
			},
			inputGetSectionsParams: &models.GetSectionsParams{
				CategoryID:  7654321,
				Locale:      "zh-tw",
				CountryCode: "tw",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			inputLocale:      "zh-tw",
			inputCountryCode: "tw",
			expectSections: []*models.Section{
				&models.Section{
					CategoryID:   7654321,
					ID:           3345678,
					Position:     101,
					CreatedAt:    time.Date(1999, 10, 13, 3, 30, 59, 0, time.UTC),
					UpdatedAt:    time.Date(1999, 10, 13, 3, 30, 59, 0, time.UTC),
					SourceLocale: "zh-cn",
					Outdated:     true,
					URL:          "https://sync.test.1",
					HTMLURL:      "https://sync.test.1",
					Name:         "同步測試 1",
					Description:  "",
					Locale:       "zh-tw",
					CountryCode:  "tw",
				},
			},
			expectCount: 1,
		},
		{
			description: "testing sync back fake data en-us locale case",
			inputSections: []*models.Section{
				&models.Section{
					CategoryID:   115002432448,
					ID:           115004118448,
					Position:     0,
					CreatedAt:    time.Date(2017, 12, 19, 6, 23, 48, 0, time.UTC),
					UpdatedAt:    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
					SourceLocale: "en-us",
					Outdated:     false,
					CountryCode:  "tw",
					URL:          "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/sections/115004118448-I-need-help-with-my-account.json",
					HTMLURL:      "https://help.honestbee.tw/hc/en-us/sections/115004118448-I-need-help-with-my-account",
					Name:         "I need help with my account",
					Description:  "",
					Locale:       "en-us",
				},
			},
			inputGetSectionsParams: &models.GetSectionsParams{
				CategoryID:  115002432448,
				Locale:      "en-us",
				CountryCode: "tw",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			inputLocale:      "en-us",
			inputCountryCode: "tw",
			expectSections: []*models.Section{
				&models.Section{
					CategoryID:   115002432448,
					ID:           115004118448,
					Position:     0,
					CreatedAt:    time.Date(2017, 12, 19, 6, 23, 48, 0, time.UTC),
					UpdatedAt:    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
					SourceLocale: "en-us",
					Outdated:     false,
					CountryCode:  "tw",
					URL:          "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/sections/115004118448-I-need-help-with-my-account.json",
					HTMLURL:      "https://help.honestbee.tw/hc/en-us/sections/115004118448-I-need-help-with-my-account",
					Name:         "I need help with my account",
					Description:  "",
					Locale:       "en-us",
				},
			},
			expectCount: 1,
		},
		{
			description: "testing sync back fake data zh-tw locale case",
			inputSections: []*models.Section{
				&models.Section{
					CategoryID:   115002432448,
					ID:           115004118448,
					Position:     0,
					CreatedAt:    time.Date(2017, 12, 19, 6, 23, 48, 0, time.UTC),
					UpdatedAt:    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
					SourceLocale: "en-us",
					Outdated:     false,
					CountryCode:  "tw",
					URL:          "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/zh-tw/sections/115004118448-%E6%88%91%E9%9C%80%E8%A6%81%E5%B8%B3%E8%99%9F%E7%9B%B8%E9%97%9C%E7%9A%84%E5%8D%94%E5%8A%A9.json",
					HTMLURL:      "https://help.honestbee.tw/hc/zh-tw/sections/115004118448-%E6%88%91%E9%9C%80%E8%A6%81%E5%B8%B3%E8%99%9F%E7%9B%B8%E9%97%9C%E7%9A%84%E5%8D%94%E5%8A%A9",
					Name:         "我需要帳號相關的協助",
					Description:  "",
					Locale:       "zh-tw",
				},
			},
			inputGetSectionsParams: &models.GetSectionsParams{
				CategoryID:  115002432448,
				Locale:      "zh-tw",
				CountryCode: "tw",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			inputLocale:      "zh-tw",
			inputCountryCode: "tw",
			expectSections: []*models.Section{
				&models.Section{
					CategoryID:   115002432448,
					ID:           115004118448,
					Position:     0,
					CreatedAt:    time.Date(2017, 12, 19, 6, 23, 48, 0, time.UTC),
					UpdatedAt:    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
					SourceLocale: "en-us",
					Outdated:     false,
					CountryCode:  "tw",
					URL:          "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/zh-tw/sections/115004118448-%E6%88%91%E9%9C%80%E8%A6%81%E5%B8%B3%E8%99%9F%E7%9B%B8%E9%97%9C%E7%9A%84%E5%8D%94%E5%8A%A9.json",
					HTMLURL:      "https://help.honestbee.tw/hc/zh-tw/sections/115004118448-%E6%88%91%E9%9C%80%E8%A6%81%E5%B8%B3%E8%99%9F%E7%9B%B8%E9%97%9C%E7%9A%84%E5%8D%94%E5%8A%A9",
					Name:         "我需要帳號相關的協助",
					Description:  "",
					Locale:       "zh-tw",
				},
			},
			expectCount: 1,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			err := service.SyncWithSections(context.Background(), tt.inputSections, tt.inputCountryCode, tt.inputLocale)
			defer resetDB()

			if tt.expectError && err == nil {
				t.Errorf("[%s] expect an error, actual none", tt.description)
			} else if !tt.expectError && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			} else if err == nil {
				actualSections, actualCount, err := service.GetSectionsByCategoryID(context.Background(), tt.inputGetSectionsParams)
				if err != nil {
					t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
				}
				for _, section := range actualSections {
					section.CreatedAt = section.CreatedAt.In(time.UTC)
					section.UpdatedAt = section.UpdatedAt.In(time.UTC)
				}

				if tt.expectCount != actualCount {
					t.Errorf("[%s] count expect:%v, actual:%v", tt.description, tt.expectCount, actualCount)
				}
				if diff := deep.Equal(tt.expectSections, actualSections); diff != nil {
					t.Errorf("[%s] %v", tt.description, diff)
				}
			}
		})
	}
}

func TestModelsForceSyncWithSections(t *testing.T) {
	service := newService()
	defer service.Close()
	testCases := []struct {
		description             string
		inputSections1          []*models.Section
		inputCountryCode1       string
		inputLocale1            string
		inputSections2          []*models.Section
		inputCountryCode2       string
		inputLocale2            string
		inputGetSectionsParams1 *models.GetSectionsParams
		expectSections1         []*models.Section
		expectCount1            int
		expectError1            bool
		inputGetSectionsParams2 *models.GetSectionsParams
		expectSections2         []*models.Section
		expectCount2            int
		expectError2            bool
	}{
		{
			description: "testing sync with sg + en-us/zh-cn one mock section case",
			inputSections1: []*models.Section{
				&models.Section{
					CategoryID:   987,
					ID:           978,
					Position:     888,
					CreatedAt:    time.Date(2000, 12, 12, 2, 22, 22, 0, time.UTC),
					UpdatedAt:    time.Date(2000, 12, 12, 2, 22, 22, 0, time.UTC),
					SourceLocale: "en-us",
					Outdated:     false,
					CountryCode:  "sg",
					URL:          "https://force.sync.test.1",
					HTMLURL:      "https://force.sync.test.1",
					Name:         "force sync test 1",
					Description:  "force sync test 1",
					Locale:       "en-us",
				},
			},
			inputCountryCode1: "sg",
			inputLocale1:      "en-us",
			inputSections2:    []*models.Section{},
			inputCountryCode2: "sg",
			inputLocale2:      "zh-cn",
			inputGetSectionsParams1: &models.GetSectionsParams{
				CategoryID:  987,
				Locale:      "en-us",
				CountryCode: "sg",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectCount1: 1,
			expectSections1: []*models.Section{
				&models.Section{
					CategoryID:   987,
					ID:           978,
					Position:     888,
					CreatedAt:    time.Date(2000, 12, 12, 2, 22, 22, 0, time.UTC),
					UpdatedAt:    time.Date(2000, 12, 12, 2, 22, 22, 0, time.UTC),
					SourceLocale: "en-us",
					Outdated:     false,
					CountryCode:  "sg",
					URL:          "https://force.sync.test.1",
					HTMLURL:      "https://force.sync.test.1",
					Name:         "force sync test 1",
					Description:  "force sync test 1",
					Locale:       "en-us",
				},
			},
			inputGetSectionsParams2: &models.GetSectionsParams{
				CategoryID:  987,
				Locale:      "zh-cn",
				CountryCode: "sg",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectCount2:    0,
			expectSections2: []*models.Section{},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			service.SyncWithSections(context.Background(), tt.inputSections1, tt.inputCountryCode1, tt.inputLocale1)
			service.SyncWithSections(context.Background(), tt.inputSections2, tt.inputCountryCode2, tt.inputLocale2)
			defer resetDB()

			actualSections1, actualCount1, err := service.GetSections(context.Background(), tt.inputGetSectionsParams1)
			if err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			}
			for _, section := range actualSections1 {
				section.CreatedAt = section.CreatedAt.In(time.UTC)
				section.UpdatedAt = section.UpdatedAt.In(time.UTC)
			}

			if tt.expectCount1 != actualCount1 {
				t.Errorf("[%s] count expect:%v, actual:%v", tt.description, tt.expectCount1, actualCount1)
			}
			if diff := deep.Equal(tt.expectSections1, actualSections1); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}

			actualSections2, actualCount2, err := service.GetSections(context.Background(), tt.inputGetSectionsParams2)
			if err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			}
			for _, section := range actualSections2 {
				section.CreatedAt = section.CreatedAt.In(time.UTC)
				section.UpdatedAt = section.UpdatedAt.In(time.UTC)
			}

			if tt.expectCount2 != actualCount2 {
				t.Errorf("[%s] count expect:%v, actual:%v", tt.description, tt.expectCount2, actualCount2)
			}
			if diff := deep.Equal(tt.expectSections2, actualSections2); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}

func TestModelsGetSections(t *testing.T) {
	service := newService()
	defer service.Close()
	testCases := []struct {
		description    string
		input          *models.GetSectionsParams
		expectSections []*models.Section
		expectCount    int
		expectError    bool
	}{
		{
			description: "testing normal zh-tw locale case",
			input: &models.GetSectionsParams{
				Locale:      "zh-tw",
				CountryCode: "tw",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectCount: 1,
			expectSections: []*models.Section{
				&models.Section{
					CategoryID:   115002432448,
					ID:           115004118448,
					Position:     0,
					CreatedAt:    time.Date(2017, 12, 19, 6, 23, 48, 0, time.UTC),
					UpdatedAt:    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
					SourceLocale: "en-us",
					Outdated:     false,
					CountryCode:  "tw",
					URL:          "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/zh-tw/sections/115004118448-%E6%88%91%E9%9C%80%E8%A6%81%E5%B8%B3%E8%99%9F%E7%9B%B8%E9%97%9C%E7%9A%84%E5%8D%94%E5%8A%A9.json",
					HTMLURL:      "https://help.honestbee.tw/hc/zh-tw/sections/115004118448-%E6%88%91%E9%9C%80%E8%A6%81%E5%B8%B3%E8%99%9F%E7%9B%B8%E9%97%9C%E7%9A%84%E5%8D%94%E5%8A%A9",
					Name:         "我需要帳號相關的協助",
					Description:  "",
					Locale:       "zh-tw",
				},
			},
		},
		{
			description: "testing normal en-us locale case",
			input: &models.GetSectionsParams{
				Locale:      "en-us",
				CountryCode: "tw",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectCount: 1,
			expectSections: []*models.Section{
				&models.Section{
					CategoryID:   115002432448,
					ID:           115004118448,
					Position:     0,
					CreatedAt:    time.Date(2017, 12, 19, 6, 23, 48, 0, time.UTC),
					UpdatedAt:    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
					SourceLocale: "en-us",
					Outdated:     false,
					CountryCode:  "tw",
					URL:          "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/sections/115004118448-I-need-help-with-my-account.json",
					HTMLURL:      "https://help.honestbee.tw/hc/en-us/sections/115004118448-I-need-help-with-my-account",
					Name:         "I need help with my account",
					Description:  "",
					Locale:       "en-us",
				},
			},
		},
		{
			description: "testing not exist locale case",
			input: &models.GetSectionsParams{
				Locale:      "not-exist-locale",
				CountryCode: "tw",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectCount:    0,
			expectSections: []*models.Section{},
		},
		{
			description: "testing not exist country code case",
			input: &models.GetSectionsParams{
				Locale:      "en-us",
				CountryCode: "not-exist-country-code",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectCount:    0,
			expectSections: []*models.Section{},
		},
		{
			description: "testing per page 1 case",
			input: &models.GetSectionsParams{
				Locale:      "en-us",
				CountryCode: "tw",
				PerPage:     1,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectCount: 1,
			expectSections: []*models.Section{
				&models.Section{
					CategoryID:   115002432448,
					ID:           115004118448,
					Position:     0,
					CreatedAt:    time.Date(2017, 12, 19, 6, 23, 48, 0, time.UTC),
					UpdatedAt:    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
					SourceLocale: "en-us",
					Outdated:     false,
					CountryCode:  "tw",
					URL:          "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/sections/115004118448-I-need-help-with-my-account.json",
					HTMLURL:      "https://help.honestbee.tw/hc/en-us/sections/115004118448-I-need-help-with-my-account",
					Name:         "I need help with my account",
					Description:  "",
					Locale:       "en-us",
				},
			},
		},
		{
			description: "testing per page 1 page 1 case",
			input: &models.GetSectionsParams{
				Locale:      "en-us",
				CountryCode: "tw",
				PerPage:     1,
				Page:        1,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectCount:    1,
			expectSections: []*models.Section{},
		},
		{
			description: "testing page overflow case",
			input: &models.GetSectionsParams{
				Locale:      "en-us",
				CountryCode: "tw",
				PerPage:     30,
				Page:        100000,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectCount:    1,
			expectSections: []*models.Section{},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actualSections, actualCount, err := service.GetSections(context.Background(), tt.input)
			if tt.expectError && err == nil {
				t.Errorf("[%s] expect an error, actual none", tt.description)
			} else if !tt.expectError && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			} else if err == nil {
				for _, section := range actualSections {
					section.CreatedAt = section.CreatedAt.In(time.UTC)
					section.UpdatedAt = section.UpdatedAt.In(time.UTC)
				}
				if tt.expectCount != actualCount {
					t.Errorf("[%s] count expect:%v, actual:%v", tt.description, tt.expectCount, actualCount)
				}
				if diff := deep.Equal(tt.expectSections, actualSections); diff != nil {
					t.Errorf("[%s] %v", tt.description, diff)
				}
			}
		})
	}
}

func TestModelsGetSectionsByCategoryID(t *testing.T) {
	service := newService()
	defer service.Close()
	testCases := []struct {
		description    string
		input          *models.GetSectionsParams
		expectSections []*models.Section
		expectCount    int
		expectError    bool
	}{
		{
			description: "testing normal zh-tw locale case",
			input: &models.GetSectionsParams{
				CategoryID:  115002432448,
				Locale:      "zh-tw",
				CountryCode: "tw",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectCount: 1,
			expectSections: []*models.Section{
				&models.Section{
					CategoryID:   115002432448,
					ID:           115004118448,
					Position:     0,
					CreatedAt:    time.Date(2017, 12, 19, 6, 23, 48, 0, time.UTC),
					UpdatedAt:    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
					SourceLocale: "en-us",
					Outdated:     false,
					CountryCode:  "tw",
					URL:          "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/zh-tw/sections/115004118448-%E6%88%91%E9%9C%80%E8%A6%81%E5%B8%B3%E8%99%9F%E7%9B%B8%E9%97%9C%E7%9A%84%E5%8D%94%E5%8A%A9.json",
					HTMLURL:      "https://help.honestbee.tw/hc/zh-tw/sections/115004118448-%E6%88%91%E9%9C%80%E8%A6%81%E5%B8%B3%E8%99%9F%E7%9B%B8%E9%97%9C%E7%9A%84%E5%8D%94%E5%8A%A9",
					Name:         "我需要帳號相關的協助",
					Description:  "",
					Locale:       "zh-tw",
				},
			},
		},
		{
			description: "testing normal en-us locale case",
			input: &models.GetSectionsParams{
				CategoryID:  115002432448,
				Locale:      "en-us",
				CountryCode: "tw",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectCount: 1,
			expectSections: []*models.Section{
				&models.Section{
					CategoryID:   115002432448,
					ID:           115004118448,
					Position:     0,
					CreatedAt:    time.Date(2017, 12, 19, 6, 23, 48, 0, time.UTC),
					UpdatedAt:    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
					SourceLocale: "en-us",
					Outdated:     false,
					CountryCode:  "tw",
					URL:          "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/sections/115004118448-I-need-help-with-my-account.json",
					HTMLURL:      "https://help.honestbee.tw/hc/en-us/sections/115004118448-I-need-help-with-my-account",
					Name:         "I need help with my account",
					Description:  "",
					Locale:       "en-us",
				},
			},
		},
		{
			description: "testing not exist locale case",
			input: &models.GetSectionsParams{
				CategoryID:  115002432448,
				Locale:      "not-exist-locale",
				CountryCode: "tw",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectCount:    0,
			expectSections: []*models.Section{},
		},
		{
			description: "testing not exist country code case",
			input: &models.GetSectionsParams{
				CategoryID:  115002432448,
				Locale:      "en-us",
				CountryCode: "not-exist-country-code",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectCount:    0,
			expectSections: []*models.Section{},
		},
		{
			description: "testing per page 1 case",
			input: &models.GetSectionsParams{
				CategoryID:  115002432448,
				Locale:      "en-us",
				CountryCode: "tw",
				PerPage:     1,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectCount: 1,
			expectSections: []*models.Section{
				&models.Section{
					CategoryID:   115002432448,
					ID:           115004118448,
					Position:     0,
					CreatedAt:    time.Date(2017, 12, 19, 6, 23, 48, 0, time.UTC),
					UpdatedAt:    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
					SourceLocale: "en-us",
					Outdated:     false,
					CountryCode:  "tw",
					URL:          "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/sections/115004118448-I-need-help-with-my-account.json",
					HTMLURL:      "https://help.honestbee.tw/hc/en-us/sections/115004118448-I-need-help-with-my-account",
					Name:         "I need help with my account",
					Description:  "",
					Locale:       "en-us",
				},
			},
		},
		{
			description: "testing per page 1 page 1 case",
			input: &models.GetSectionsParams{
				CategoryID:  115002432448,
				Locale:      "en-us",
				CountryCode: "tw",
				PerPage:     1,
				Page:        1,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectCount:    1,
			expectSections: []*models.Section{},
		},
		{
			description: "testing page overflow case",
			input: &models.GetSectionsParams{
				CategoryID:  115002432448,
				Locale:      "en-us",
				CountryCode: "tw",
				PerPage:     30,
				Page:        100000,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectCount:    1,
			expectSections: []*models.Section{},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actualSections, actualCount, err := service.GetSectionsByCategoryID(context.Background(), tt.input)
			if tt.expectError && err == nil {
				t.Errorf("[%s] expect an error, actual none", tt.description)
			} else if !tt.expectError && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			} else if err == nil {
				for _, section := range actualSections {
					section.CreatedAt = section.CreatedAt.In(time.UTC)
					section.UpdatedAt = section.UpdatedAt.In(time.UTC)
				}
				if tt.expectCount != actualCount {
					t.Errorf("[%s] count expect:%v, actual:%v", tt.description, tt.expectCount, actualCount)
				}
				if diff := deep.Equal(tt.expectSections, actualSections); diff != nil {
					t.Errorf("[%s] %v", tt.description, diff)
				}
			}
		})
	}
}

func TestModelsGetSectionBySectionID(t *testing.T) {
	service := newService()
	defer service.Close()
	testCases := []struct {
		description      string
		inputSectionID   int
		inputLocale      string
		inputCountryCode string
		expectSection    *models.Section
		expectError      bool
	}{
		{
			description:      "testing normal zh-tw locale case",
			inputSectionID:   115004118448,
			inputLocale:      "zh-tw",
			inputCountryCode: "tw",
			expectSection: &models.Section{
				CategoryID:   115002432448,
				ID:           115004118448,
				Position:     0,
				CreatedAt:    time.Date(2017, 12, 19, 6, 23, 48, 0, time.UTC),
				UpdatedAt:    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
				SourceLocale: "en-us",
				Outdated:     false,
				CountryCode:  "tw",
				URL:          "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/zh-tw/sections/115004118448-%E6%88%91%E9%9C%80%E8%A6%81%E5%B8%B3%E8%99%9F%E7%9B%B8%E9%97%9C%E7%9A%84%E5%8D%94%E5%8A%A9.json",
				HTMLURL:      "https://help.honestbee.tw/hc/zh-tw/sections/115004118448-%E6%88%91%E9%9C%80%E8%A6%81%E5%B8%B3%E8%99%9F%E7%9B%B8%E9%97%9C%E7%9A%84%E5%8D%94%E5%8A%A9",
				Name:         "我需要帳號相關的協助",
				Description:  "",
				Locale:       "zh-tw",
			},
		},
		{
			description:      "testing normal en-us locale case",
			inputSectionID:   115004118448,
			inputLocale:      "en-us",
			inputCountryCode: "tw",
			expectSection: &models.Section{
				CategoryID:   115002432448,
				ID:           115004118448,
				Position:     0,
				CreatedAt:    time.Date(2017, 12, 19, 6, 23, 48, 0, time.UTC),
				UpdatedAt:    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
				SourceLocale: "en-us",
				Outdated:     false,
				CountryCode:  "tw",
				URL:          "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/sections/115004118448-I-need-help-with-my-account.json",
				HTMLURL:      "https://help.honestbee.tw/hc/en-us/sections/115004118448-I-need-help-with-my-account",
				Name:         "I need help with my account",
				Description:  "",
				Locale:       "en-us",
			},
		},
		{
			description:      "testing not exist country code case",
			inputSectionID:   115004118448,
			inputLocale:      "zh-tw",
			inputCountryCode: "not-exist-country-code",
			expectError:      true,
			expectSection:    nil,
		},
		{
			description:      "testing not exist locale case",
			inputSectionID:   115004118448,
			inputLocale:      "not-exist-locale",
			inputCountryCode: "tw",
			expectError:      false,
			expectSection: &models.Section{
				CategoryID:   115002432448,
				ID:           115004118448,
				Position:     0,
				CreatedAt:    time.Date(2017, 12, 19, 6, 23, 48, 0, time.UTC),
				UpdatedAt:    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
				SourceLocale: "en-us",
				Outdated:     false,
				CountryCode:  "tw",
				URL:          "",
				HTMLURL:      "",
				Name:         "",
				Description:  "",
				Locale:       "",
			},
		},
		{
			description:      "testing not exist section id case",
			inputSectionID:   675849302348,
			inputLocale:      "zh-tw",
			inputCountryCode: "tw",
			expectError:      true,
			expectSection:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actualSection, err := service.GetSectionBySectionID(context.Background(), tt.inputSectionID, tt.inputLocale, tt.inputCountryCode)
			if tt.expectError && err == nil {
				t.Errorf("[%s] expect an error, actual none", tt.description)
			} else if !tt.expectError && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			} else if err == nil {
				actualSection.CreatedAt = actualSection.CreatedAt.In(time.UTC)
				actualSection.UpdatedAt = actualSection.UpdatedAt.In(time.UTC)
				if diff := deep.Equal(tt.expectSection, actualSection); diff != nil {
					t.Errorf("[%s] %v", tt.description, diff)
				}
			}
		})
	}
}
