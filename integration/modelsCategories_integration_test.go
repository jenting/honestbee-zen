// +build integration

package integration

import (
	"context"
	"testing"
	"time"

	"github.com/go-test/deep"

	"github.com/honestbee/Zen/models"
)

func TestModelsSyncWithCategories(t *testing.T) {
	service := newService()
	defer service.Close()
	testCases := []struct {
		description              string
		inputCategories          []*models.Category
		inputGetCategoriesParams *models.GetCategoriesParams
		inputCountryCode         string
		inputLocale              string
		expectCategories         []*models.Category
		expectCount              int
		expectError              bool
	}{
		{
			description: "testing sync with en-us locale one mock category case",
			inputCategories: []*models.Category{
				&models.Category{
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
			inputGetCategoriesParams: &models.GetCategoriesParams{
				Locale:      "en-us",
				CountryCode: "tw",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			inputLocale:      "en-us",
			inputCountryCode: "tw",
			expectCategories: []*models.Category{
				&models.Category{
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
			description: "testing sync with zh-tw locale one mock category case",
			inputCategories: []*models.Category{
				&models.Category{
					ID:           3345678,
					Position:     0,
					CreatedAt:    time.Date(1988, 10, 13, 3, 30, 59, 0, time.UTC),
					UpdatedAt:    time.Date(1988, 10, 13, 3, 30, 59, 0, time.UTC),
					SourceLocale: "en-us",
					Outdated:     false,
					URL:          "https://sync.test.1",
					HTMLURL:      "https://sync.test.1",
					Name:         "同步測試 1",
					Description:  "",
					Locale:       "zh-tw",
					CountryCode:  "tw",
				},
			},
			inputGetCategoriesParams: &models.GetCategoriesParams{
				Locale:      "zh-tw",
				CountryCode: "tw",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			inputLocale:      "zh-tw",
			inputCountryCode: "tw",
			expectCategories: []*models.Category{
				&models.Category{
					ID:           3345678,
					Position:     0,
					CreatedAt:    time.Date(1988, 10, 13, 3, 30, 59, 0, time.UTC),
					UpdatedAt:    time.Date(1988, 10, 13, 3, 30, 59, 0, time.UTC),
					SourceLocale: "en-us",
					Outdated:     false,
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
			description: "testing sync with en-us locale two mock categories case",
			inputCategories: []*models.Category{
				&models.Category{
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
				&models.Category{
					ID:           3345679,
					Position:     0,
					CreatedAt:    time.Date(1989, 10, 13, 3, 30, 59, 0, time.UTC),
					UpdatedAt:    time.Date(1989, 10, 13, 3, 30, 59, 0, time.UTC),
					SourceLocale: "en-us",
					Outdated:     false,
					URL:          "https://sync.test.2",
					HTMLURL:      "https://sync.test.2",
					Name:         "sync test 2",
					Description:  "",
					Locale:       "en-us",
					CountryCode:  "tw",
				},
			},
			inputGetCategoriesParams: &models.GetCategoriesParams{
				Locale:      "en-us",
				CountryCode: "tw",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			inputLocale:      "en-us",
			inputCountryCode: "tw",
			expectCategories: []*models.Category{
				&models.Category{
					ID:           3345679,
					Position:     0,
					CreatedAt:    time.Date(1989, 10, 13, 3, 30, 59, 0, time.UTC),
					UpdatedAt:    time.Date(1989, 10, 13, 3, 30, 59, 0, time.UTC),
					SourceLocale: "en-us",
					Outdated:     false,
					URL:          "https://sync.test.2",
					HTMLURL:      "https://sync.test.2",
					Name:         "sync test 2",
					Description:  "",
					Locale:       "en-us",
					CountryCode:  "tw",
				},
				&models.Category{
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
			expectCount: 2,
		},
		{
			description: "testing sync with zh-tw locale two mock categories case",
			inputCategories: []*models.Category{
				&models.Category{
					ID:           3345678,
					Position:     0,
					CreatedAt:    time.Date(1988, 10, 13, 3, 30, 59, 0, time.UTC),
					UpdatedAt:    time.Date(1988, 10, 13, 3, 30, 59, 0, time.UTC),
					SourceLocale: "en-us",
					Outdated:     false,
					URL:          "https://sync.test.1",
					HTMLURL:      "https://sync.test.1",
					Name:         "同步測試 1",
					Description:  "",
					Locale:       "zh-tw",
					CountryCode:  "tw",
				},
				&models.Category{
					ID:           3345679,
					Position:     0,
					CreatedAt:    time.Date(1989, 10, 13, 3, 30, 59, 0, time.UTC),
					UpdatedAt:    time.Date(1989, 10, 13, 3, 30, 59, 0, time.UTC),
					SourceLocale: "en-us",
					Outdated:     false,
					URL:          "https://sync.test.2",
					HTMLURL:      "https://sync.test.2",
					Name:         "同步測試 2",
					Description:  "",
					Locale:       "zh-tw",
					CountryCode:  "tw",
				},
			},
			inputGetCategoriesParams: &models.GetCategoriesParams{
				Locale:      "zh-tw",
				CountryCode: "tw",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			inputLocale:      "zh-tw",
			inputCountryCode: "tw",
			expectCategories: []*models.Category{
				&models.Category{
					ID:           3345679,
					Position:     0,
					CreatedAt:    time.Date(1989, 10, 13, 3, 30, 59, 0, time.UTC),
					UpdatedAt:    time.Date(1989, 10, 13, 3, 30, 59, 0, time.UTC),
					SourceLocale: "en-us",
					Outdated:     false,
					URL:          "https://sync.test.2",
					HTMLURL:      "https://sync.test.2",
					Name:         "同步測試 2",
					Description:  "",
					Locale:       "zh-tw",
					CountryCode:  "tw",
				},
				&models.Category{
					ID:           3345678,
					Position:     0,
					CreatedAt:    time.Date(1988, 10, 13, 3, 30, 59, 0, time.UTC),
					UpdatedAt:    time.Date(1988, 10, 13, 3, 30, 59, 0, time.UTC),
					SourceLocale: "en-us",
					Outdated:     false,
					URL:          "https://sync.test.1",
					HTMLURL:      "https://sync.test.1",
					Name:         "同步測試 1",
					Description:  "",
					Locale:       "zh-tw",
					CountryCode:  "tw",
				},
			},
			expectCount: 2,
		},
		{
			description: "testing sync back fake data en-us locale case",
			inputCategories: []*models.Category{
				&models.Category{
					ID:           115002432448,
					Position:     2,
					CreatedAt:    time.Date(2017, 12, 19, 6, 21, 45, 0, time.UTC),
					UpdatedAt:    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
					SourceLocale: "en-us",
					Outdated:     false,
					CountryCode:  "tw",
					URL:          "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/categories/115002432448-My-Account.json",
					HTMLURL:      "https://help.honestbee.tw/hc/en-us/categories/115002432448-My-Account",
					Name:         "My Account",
					Description:  "",
					Locale:       "en-us",
				},
			},
			inputGetCategoriesParams: &models.GetCategoriesParams{
				Locale:      "en-us",
				CountryCode: "tw",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			inputLocale:      "en-us",
			inputCountryCode: "tw",
			expectCategories: []*models.Category{
				&models.Category{
					ID:           115002432448,
					Position:     2,
					CreatedAt:    time.Date(2017, 12, 19, 6, 21, 45, 0, time.UTC),
					UpdatedAt:    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
					SourceLocale: "en-us",
					Outdated:     false,
					CountryCode:  "tw",
					URL:          "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/categories/115002432448-My-Account.json",
					HTMLURL:      "https://help.honestbee.tw/hc/en-us/categories/115002432448-My-Account",
					Name:         "My Account",
					Description:  "",
					Locale:       "en-us",
					KeyName:      "myAccount",
				},
			},
			expectCount: 1,
		},
		{
			description: "testing sync back fake data zh-tw locale case",
			inputCategories: []*models.Category{
				&models.Category{
					ID:           115002432448,
					Position:     2,
					CreatedAt:    time.Date(2017, 12, 19, 6, 21, 45, 0, time.UTC),
					UpdatedAt:    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
					SourceLocale: "en-us",
					Outdated:     false,
					CountryCode:  "tw",
					URL:          "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/zh-tw/categories/115002432448-%E6%88%91%E7%9A%84%E5%B8%B3%E8%99%9F.json",
					HTMLURL:      "https://help.honestbee.tw/hc/zh-tw/categories/115002432448-%E6%88%91%E7%9A%84%E5%B8%B3%E8%99%9F",
					Name:         "我的帳號",
					Description:  "",
					Locale:       "zh-tw",
				},
			},
			inputGetCategoriesParams: &models.GetCategoriesParams{
				Locale:      "zh-tw",
				CountryCode: "tw",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			inputLocale:      "zh-tw",
			inputCountryCode: "tw",
			expectCategories: []*models.Category{
				&models.Category{
					ID:           115002432448,
					Position:     2,
					CreatedAt:    time.Date(2017, 12, 19, 6, 21, 45, 0, time.UTC),
					UpdatedAt:    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
					SourceLocale: "en-us",
					Outdated:     false,
					CountryCode:  "tw",
					URL:          "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/zh-tw/categories/115002432448-%E6%88%91%E7%9A%84%E5%B8%B3%E8%99%9F.json",
					HTMLURL:      "https://help.honestbee.tw/hc/zh-tw/categories/115002432448-%E6%88%91%E7%9A%84%E5%B8%B3%E8%99%9F",
					Name:         "我的帳號",
					Description:  "",
					Locale:       "zh-tw",
					KeyName:      "myAccount",
				},
			},
			expectCount: 1,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			err := service.SyncWithCategories(context.Background(), tt.inputCategories, tt.inputCountryCode, tt.inputLocale)
			defer resetDB()

			if tt.expectError && err == nil {
				t.Errorf("[%s] expect an error, actual none", tt.description)
			} else if !tt.expectError && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			} else if err == nil {
				actualCategories, actualCount, err := service.GetCategories(context.Background(), tt.inputGetCategoriesParams)
				if err != nil {
					t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
				}
				for _, category := range actualCategories {
					category.CreatedAt = category.CreatedAt.In(time.UTC)
					category.UpdatedAt = category.UpdatedAt.In(time.UTC)
				}

				if tt.expectCount != actualCount {
					t.Errorf("[%s] count expect:%v, actual:%v", tt.description, tt.expectCount, actualCount)
				}
				if diff := deep.Equal(tt.expectCategories, actualCategories); diff != nil {
					t.Errorf("[%s] %v", tt.description, diff)
				}
			}
		})
	}
}

func TestModelsForceSyncCategories(t *testing.T) {
	service := newService()
	defer service.Close()
	testCases := []struct {
		description               string
		inputCategories1          []*models.Category
		inputCountryCode1         string
		inputLocale1              string
		inputCategories2          []*models.Category
		inputCountryCode2         string
		inputLocale2              string
		inputGetCategoriesParams1 *models.GetCategoriesParams
		expectCategories1         []*models.Category
		expectCount1              int
		expectError1              bool
		inputGetCategoriesParams2 *models.GetCategoriesParams
		expectCategories2         []*models.Category
		expectCount2              int
		expectError2              bool
	}{
		{
			description: "testing sync with sg + en-us/zh-cn one mock category case",
			inputCategories1: []*models.Category{
				&models.Category{
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
			inputCategories2:  []*models.Category{},
			inputCountryCode2: "sg",
			inputLocale2:      "zh-cn",
			inputGetCategoriesParams1: &models.GetCategoriesParams{
				Locale:      "en-us",
				CountryCode: "sg",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectCount1: 1,
			expectCategories1: []*models.Category{
				&models.Category{
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
			inputGetCategoriesParams2: &models.GetCategoriesParams{
				Locale:      "zh-cn",
				CountryCode: "sg",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectCount2:      0,
			expectCategories2: []*models.Category{},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			service.SyncWithCategories(context.Background(), tt.inputCategories1, tt.inputCountryCode1, tt.inputLocale1)
			service.SyncWithCategories(context.Background(), tt.inputCategories2, tt.inputCountryCode2, tt.inputLocale2)
			defer resetDB()

			actualCategories1, actualCount1, err := service.GetCategories(context.Background(), tt.inputGetCategoriesParams1)
			if err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			}
			for _, category := range actualCategories1 {
				category.CreatedAt = category.CreatedAt.In(time.UTC)
				category.UpdatedAt = category.UpdatedAt.In(time.UTC)
			}

			if tt.expectCount1 != actualCount1 {
				t.Errorf("[%s] count expect:%v, actual:%v", tt.description, tt.expectCount1, actualCount1)
			}
			if diff := deep.Equal(tt.expectCategories1, actualCategories1); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}

			actualCategories2, actualCount2, err := service.GetCategories(context.Background(), tt.inputGetCategoriesParams2)
			if err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			}
			for _, category := range actualCategories2 {
				category.CreatedAt = category.CreatedAt.In(time.UTC)
				category.UpdatedAt = category.UpdatedAt.In(time.UTC)
			}

			if tt.expectCount2 != actualCount2 {
				t.Errorf("[%s] count expect:%v, actual:%v", tt.description, tt.expectCount2, actualCount2)
			}
			if diff := deep.Equal(tt.expectCategories2, actualCategories2); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}

func TestModelsGetCategories(t *testing.T) {
	service := newService()
	defer service.Close()
	testCases := []struct {
		description      string
		input            *models.GetCategoriesParams
		expectCategories []*models.Category
		expectCount      int
		expectError      bool
	}{
		{
			description: "testing normal zh-tw locale case",
			input: &models.GetCategoriesParams{
				Locale:      "zh-tw",
				CountryCode: "tw",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectCount: 1,
			expectCategories: []*models.Category{
				&models.Category{
					ID:           115002432448,
					Position:     2,
					CreatedAt:    time.Date(2017, 12, 19, 6, 21, 45, 0, time.UTC),
					UpdatedAt:    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
					SourceLocale: "en-us",
					Outdated:     false,
					CountryCode:  "tw",
					URL:          "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/zh-tw/categories/115002432448-%E6%88%91%E7%9A%84%E5%B8%B3%E8%99%9F.json",
					HTMLURL:      "https://help.honestbee.tw/hc/zh-tw/categories/115002432448-%E6%88%91%E7%9A%84%E5%B8%B3%E8%99%9F",
					Name:         "我的帳號",
					Description:  "",
					Locale:       "zh-tw",
					KeyName:      "myAccount",
				},
			},
		},
		{
			description: "testing normal en-us locale case",
			input: &models.GetCategoriesParams{
				Locale:      "en-us",
				CountryCode: "tw",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectCount: 1,
			expectCategories: []*models.Category{
				&models.Category{
					ID:           115002432448,
					Position:     2,
					CreatedAt:    time.Date(2017, 12, 19, 6, 21, 45, 0, time.UTC),
					UpdatedAt:    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
					SourceLocale: "en-us",
					Outdated:     false,
					CountryCode:  "tw",
					URL:          "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/categories/115002432448-My-Account.json",
					HTMLURL:      "https://help.honestbee.tw/hc/en-us/categories/115002432448-My-Account",
					Name:         "My Account",
					Description:  "",
					Locale:       "en-us",
					KeyName:      "myAccount",
				},
			},
		},
		{
			description: "testing not exist locale case",
			input: &models.GetCategoriesParams{
				Locale:      "not-exist-locale",
				CountryCode: "tw",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectCount:      0,
			expectCategories: []*models.Category{},
		},
		{
			description: "testing not exist country code case",
			input: &models.GetCategoriesParams{
				Locale:      "en-us",
				CountryCode: "not-exist-country-code",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectCount:      0,
			expectCategories: []*models.Category{},
		},
		{
			description: "testing per page 1 case",
			input: &models.GetCategoriesParams{
				Locale:      "en-us",
				CountryCode: "tw",
				PerPage:     1,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectCount: 1,
			expectCategories: []*models.Category{
				&models.Category{
					ID:           115002432448,
					Position:     2,
					CreatedAt:    time.Date(2017, 12, 19, 6, 21, 45, 0, time.UTC),
					UpdatedAt:    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
					SourceLocale: "en-us",
					Outdated:     false,
					CountryCode:  "tw",
					URL:          "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/categories/115002432448-My-Account.json",
					HTMLURL:      "https://help.honestbee.tw/hc/en-us/categories/115002432448-My-Account",
					Name:         "My Account",
					Description:  "",
					Locale:       "en-us",
					KeyName:      "myAccount",
				},
			},
		},
		{
			description: "testing per page 1 page 1 case",
			input: &models.GetCategoriesParams{
				Locale:      "en-us",
				CountryCode: "tw",
				PerPage:     1,
				Page:        1,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectCount:      1,
			expectCategories: []*models.Category{},
		},
		{
			description: "testing page overflow case",
			input: &models.GetCategoriesParams{
				Locale:      "en-us",
				CountryCode: "tw",
				PerPage:     30,
				Page:        100000,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectCount:      1,
			expectCategories: []*models.Category{},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actualCategories, actualCount, err := service.GetCategories(context.Background(), tt.input)
			if tt.expectError && err == nil {
				t.Errorf("[%s] expect an error, actual none", tt.description)
			} else if !tt.expectError && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			} else if err == nil {
				// doing in this convertions is because: https://groups.google.com/forum/#!topic/golang-nuts/RPMLsEkvY3E
				for _, category := range actualCategories {
					category.CreatedAt = category.CreatedAt.In(time.UTC)
					category.UpdatedAt = category.UpdatedAt.In(time.UTC)
				}
				if tt.expectCount != actualCount {
					t.Errorf("[%s] count expect:%v, actual:%v", tt.description, tt.expectCount, actualCount)
				}
				if diff := deep.Equal(tt.expectCategories, actualCategories); diff != nil {
					t.Errorf("[%s] %v", tt.description, diff)
				}
			}
		})
	}
}

func TestModelsGetCategoriesID(t *testing.T) {
	service := newService()
	defer service.Close()
	testCases := []struct {
		description        string
		inputCountryCode   string
		expectCategoriesID []int
		expectError        bool
	}{
		{
			description:        "testing normal tw country code case",
			inputCountryCode:   "tw",
			expectCategoriesID: []int{115002432448},
		},
		{
			description:        "testing not exist country code case",
			inputCountryCode:   "not-exist-country-code",
			expectCategoriesID: []int{},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actualCategoriesID, err := service.GetCategoriesID(context.Background(), tt.inputCountryCode)
			if tt.expectError && err == nil {
				t.Errorf("[%s] expect an error, actual none", tt.description)
			} else if !tt.expectError && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			} else if err == nil {
				if diff := deep.Equal(tt.expectCategoriesID, actualCategoriesID); diff != nil {
					t.Errorf("[%s] %v", tt.description, diff)
				}
			}
		})
	}
}

func TestModelsGetCategoryKeyNameToID(t *testing.T) {
	service := newService()
	defer service.Close()
	testCases := []struct {
		description      string
		inputKeyName     string
		inputCountryCode string
		expectCategoryID int
		expectError      bool
	}{
		{
			description:      "testing membership sg case",
			inputKeyName:     "memberShips",
			inputCountryCode: "sg",
			expectCategoryID: 360000148587,
		},
		{
			description:      "testing membership tw case",
			inputKeyName:     "memberShips",
			inputCountryCode: "tw",
			expectCategoryID: 360000148927,
		},
		{
			description:      "testing supports all upper case",
			inputKeyName:     "MEMBERSHIPS",
			inputCountryCode: "sg",
			expectCategoryID: 360000148587,
		},
		{
			description:      "testing supports all lower case",
			inputKeyName:     "memberships",
			inputCountryCode: "sg",
			expectCategoryID: 360000148587,
		},
		{
			description:      "testing not exist country code case",
			inputKeyName:     "memberShips",
			inputCountryCode: "not-exist-country-code",
			expectError:      true,
		},
		{
			description:      "testing not exist key name case",
			inputKeyName:     "non-exist-key-name",
			inputCountryCode: "sg",
			expectError:      true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actualCategoryID, err := service.GetCategoryKeyNameToID(context.Background(), tt.inputKeyName, tt.inputCountryCode)
			if tt.expectError && err == nil {
				t.Errorf("[%s] expect an error, actual none", tt.description)
			} else if !tt.expectError && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			} else if err == nil {
				if diff := deep.Equal(tt.expectCategoryID, actualCategoryID); diff != nil {
					t.Errorf("[%s] %v", tt.description, diff)
				}
			}
		})
	}
}

func TestModelsGetCategoryByArticleID(t *testing.T) {
	service := newService()
	defer service.Close()
	testCases := []struct {
		description    string
		inputArticleID int
		inputLocale    string
		expectCategory *models.Category
		expectError    bool
	}{
		{
			description:    "testing normal zh-tw locale case",
			inputArticleID: 115015959148,
			inputLocale:    "zh-tw",
			expectCategory: &models.Category{
				ID:           115002432448,
				Position:     2,
				CreatedAt:    time.Date(2017, 12, 19, 6, 21, 45, 0, time.UTC),
				UpdatedAt:    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
				SourceLocale: "en-us",
				Outdated:     false,
				CountryCode:  "tw",
				URL:          "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/zh-tw/categories/115002432448-%E6%88%91%E7%9A%84%E5%B8%B3%E8%99%9F.json",
				HTMLURL:      "https://help.honestbee.tw/hc/zh-tw/categories/115002432448-%E6%88%91%E7%9A%84%E5%B8%B3%E8%99%9F",
				Name:         "我的帳號",
				Description:  "",
				Locale:       "zh-tw",
			},
		},
		{
			description:    "testing normal en-us locale case",
			inputArticleID: 115015959148,
			inputLocale:    "en-us",
			expectCategory: &models.Category{
				ID:           115002432448,
				Position:     2,
				CreatedAt:    time.Date(2017, 12, 19, 6, 21, 45, 0, time.UTC),
				UpdatedAt:    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
				SourceLocale: "en-us",
				Outdated:     false,
				CountryCode:  "tw",
				URL:          "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/categories/115002432448-My-Account.json",
				HTMLURL:      "https://help.honestbee.tw/hc/en-us/categories/115002432448-My-Account",
				Name:         "My Account",
				Description:  "",
				Locale:       "en-us",
			},
		},
		{
			description:    "testing not exist locale case",
			inputArticleID: 115015959148,
			inputLocale:    "not-exist-locale",
			expectError:    true,
			expectCategory: nil,
		},
		{
			description:    "testing not exist section id case",
			inputArticleID: 123456789,
			inputLocale:    "en-us",
			expectError:    true,
			expectCategory: nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actualCategory, err := service.GetCategoryByArticleID(context.Background(), tt.inputArticleID, tt.inputLocale)
			if tt.expectError && err == nil {
				t.Errorf("[%s] expect an error, actual none", tt.description)
			} else if !tt.expectError && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			} else if err == nil {
				if diff := deep.Equal(tt.expectCategory, actualCategory); diff != nil {
					t.Errorf("[%s] %v", tt.description, diff)
				}
			}
		})
	}
}

func TestModelsGetCategoryByCategoryIDOrKeyName(t *testing.T) {
	service := newService()
	defer service.Close()
	testCases := []struct {
		description      string
		inputIDOrKeyName string
		inputLocale      string
		inputCountryCode string
		expectCategory   *models.Category
		expectError      bool
	}{
		{
			description:      "testing normal category id zh-tw locale case",
			inputIDOrKeyName: "115002432448",
			inputLocale:      "zh-tw",
			inputCountryCode: "tw",
			expectCategory: &models.Category{
				ID:           115002432448,
				Position:     2,
				CreatedAt:    time.Date(2017, 12, 19, 6, 21, 45, 0, time.UTC),
				UpdatedAt:    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
				SourceLocale: "en-us",
				Outdated:     false,
				CountryCode:  "tw",
				URL:          "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/zh-tw/categories/115002432448-%E6%88%91%E7%9A%84%E5%B8%B3%E8%99%9F.json",
				HTMLURL:      "https://help.honestbee.tw/hc/zh-tw/categories/115002432448-%E6%88%91%E7%9A%84%E5%B8%B3%E8%99%9F",
				Name:         "我的帳號",
				Description:  "",
				Locale:       "zh-tw",
				KeyName:      "myAccount",
			},
		},
		{
			description:      "testing normal category id en-us locale case",
			inputIDOrKeyName: "115002432448",
			inputLocale:      "en-us",
			inputCountryCode: "tw",
			expectCategory: &models.Category{
				ID:           115002432448,
				Position:     2,
				CreatedAt:    time.Date(2017, 12, 19, 6, 21, 45, 0, time.UTC),
				UpdatedAt:    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
				SourceLocale: "en-us",
				Outdated:     false,
				CountryCode:  "tw",
				URL:          "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/categories/115002432448-My-Account.json",
				HTMLURL:      "https://help.honestbee.tw/hc/en-us/categories/115002432448-My-Account",
				Name:         "My Account",
				Description:  "",
				Locale:       "en-us",
				KeyName:      "myAccount",
			},
		},
		{
			description:      "testing normal category keyname zh-tw locale case",
			inputIDOrKeyName: "myAccount",
			inputLocale:      "zh-tw",
			inputCountryCode: "tw",
			expectCategory: &models.Category{
				ID:           115002432448,
				Position:     2,
				CreatedAt:    time.Date(2017, 12, 19, 6, 21, 45, 0, time.UTC),
				UpdatedAt:    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
				SourceLocale: "en-us",
				Outdated:     false,
				CountryCode:  "tw",
				URL:          "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/zh-tw/categories/115002432448-%E6%88%91%E7%9A%84%E5%B8%B3%E8%99%9F.json",
				HTMLURL:      "https://help.honestbee.tw/hc/zh-tw/categories/115002432448-%E6%88%91%E7%9A%84%E5%B8%B3%E8%99%9F",
				Name:         "我的帳號",
				Description:  "",
				Locale:       "zh-tw",
				KeyName:      "myAccount",
			},
		},
		{
			description:      "testing normal category keyname en-us locale case",
			inputIDOrKeyName: "myAccount",
			inputLocale:      "en-us",
			inputCountryCode: "tw",
			expectCategory: &models.Category{
				ID:           115002432448,
				Position:     2,
				CreatedAt:    time.Date(2017, 12, 19, 6, 21, 45, 0, time.UTC),
				UpdatedAt:    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
				SourceLocale: "en-us",
				Outdated:     false,
				CountryCode:  "tw",
				URL:          "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/categories/115002432448-My-Account.json",
				HTMLURL:      "https://help.honestbee.tw/hc/en-us/categories/115002432448-My-Account",
				Name:         "My Account",
				Description:  "",
				Locale:       "en-us",
				KeyName:      "myAccount",
			},
		},
		{
			description:      "testing not exist locale case",
			inputIDOrKeyName: "myAccount",
			inputLocale:      "not-exist-locale",
			inputCountryCode: "tw",
			expectError:      true,
			expectCategory:   nil,
		},
		{
			description:      "testing not exist country code case",
			inputIDOrKeyName: "myAccount",
			inputLocale:      "not-exist-country-code",
			inputCountryCode: "tw",
			expectError:      true,
			expectCategory:   nil,
		},
		{
			description:      "testing not exist category id case",
			inputIDOrKeyName: "123456789",
			inputLocale:      "en-us",
			inputCountryCode: "tw",
			expectError:      true,
			expectCategory:   nil,
		},
		{
			description:      "testing not exist category keyname case",
			inputIDOrKeyName: "non-exist-key-name",
			inputLocale:      "en-us",
			inputCountryCode: "tw",
			expectError:      true,
			expectCategory:   nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actualCategory, err := service.GetCategoryByCategoryIDOrKeyName(context.Background(), tt.inputIDOrKeyName, tt.inputLocale, tt.inputCountryCode)
			if tt.expectError && err == nil {
				t.Errorf("[%s] expect an error, actual none", tt.description)
			} else if !tt.expectError && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			} else if err == nil {
				if diff := deep.Equal(tt.expectCategory, actualCategory); diff != nil {
					t.Errorf("[%s] %v", tt.description, diff)
				}
			}
		})
	}
}
