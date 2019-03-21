package dataloader

import (
	"context"
	"testing"

	"github.com/go-test/deep"
	gographql "github.com/graph-gophers/graphql-go"

	"github.com/honestbee/Zen/inout"
	"github.com/honestbee/Zen/models"
)

func TestLoadCategories(t *testing.T) {
	testCases := [...]struct {
		description  string
		inputContext context.Context
		inputParams  interface{}
		expectErr    bool
		expect       *inout.GetCategoriesOut
	}{
		{
			description:  "testing normal case",
			inputContext: ctx,
			inputParams: inout.QueryCategoriesIn{
				CountryCode: "tw",
				Locale:      "en-us",
				PerPage:     3,
				Page:        1,
				SortBy:      "position",
				SortOrder:   "asc",
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
			description:  "testing json marshal parameter failed case",
			inputContext: ctx,
			inputParams:  make(chan int),
			expectErr:    true,
			expect:       nil,
		},
		{
			description:  "testing json unmarshal parameter failed case",
			inputContext: ctx,
			inputParams: &struct {
				CountryCode int32
				Locale      int32
			}{
				CountryCode: 123,
				Locale:      456,
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description:  "testing extract dataloader failed case",
			inputContext: context.TODO(),
			inputParams: inout.QueryCategoriesIn{
				CountryCode: "tw",
				Locale:      "en-us",
				PerPage:     3,
				Page:        1,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description:  "testing invalid country code case",
			inputContext: ctx,
			inputParams: inout.QueryCategoriesIn{
				CountryCode: models.ModelsReturnErrorCountryCode,
				Locale:      "en-us",
				PerPage:     3,
				Page:        1,
				SortBy:      "position",
				SortOrder:   "asc",
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := LoadCategories(tt.inputContext, tt.inputParams)

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

func TestLoadCategory(t *testing.T) {
	testCases := [...]struct {
		description  string
		inputContext context.Context
		inputParams  interface{}
		expectErr    bool
		expect       *models.Category
	}{
		{
			description:  "testing normal case",
			inputContext: ctx,
			inputParams: inout.QueryCategoryIn{
				CategoryIDOrKeyName: gographql.ID("3345678"),
				CountryCode:         "tw",
				Locale:              "en-us",
			},
			expectErr: false,
			expect: &models.Category{
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
		{
			description:  "testing json marshal parameter failed case",
			inputContext: ctx,
			inputParams:  make(chan int),
			expectErr:    true,
			expect:       nil,
		},
		{
			description:  "testing json unmarshal parameter failed case",
			inputContext: ctx,
			inputParams: &struct {
				CountryCode int32
				Locale      int32
			}{
				CountryCode: 123,
				Locale:      456,
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description:  "testing extract dataloader failed case",
			inputContext: context.TODO(),
			inputParams: inout.QueryCategoryIn{
				CategoryIDOrKeyName: gographql.ID("3345678"),
				CountryCode:         "tw",
				Locale:              "en-us",
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description:  "testing invalid country code case",
			inputContext: ctx,
			inputParams: inout.QueryCategoryIn{
				CategoryIDOrKeyName: gographql.ID("3345678"),
				CountryCode:         models.ModelsReturnErrorCountryCode,
				Locale:              "en-us",
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description:  "testing category id or key name not found case",
			inputContext: ctx,
			inputParams: inout.QueryCategoryIn{
				CategoryIDOrKeyName: gographql.ID("3345678"),
				CountryCode:         models.ModelsReturnNotFoundCountryCode,
				Locale:              "en-us",
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := LoadCategory(tt.inputContext, tt.inputParams)

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
