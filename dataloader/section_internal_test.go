package dataloader

import (
	"context"
	"testing"

	"github.com/go-test/deep"
	gographql "github.com/graph-gophers/graphql-go"

	"github.com/honestbee/Zen/inout"
	"github.com/honestbee/Zen/models"
)

func TestLoadSections(t *testing.T) {
	cid := gographql.ID("3345678")

	testCases := [...]struct {
		description  string
		inputContext context.Context
		inputParams  interface{}
		expectErr    bool
		expect       *inout.GetSectionsOut
	}{
		{
			description:  "testing normal w/o category id case",
			inputContext: ctx,
			inputParams: inout.QuerySectionsIn{
				CountryCode: "tw",
				Locale:      "en-us",
				PerPage:     3,
				Page:        1,
				SortBy:      "position",
				SortOrder:   "asc",
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
			inputParams: inout.QuerySectionsIn{
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
			inputParams: inout.QuerySectionsIn{
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
		{
			description:  "testing normal w/ category id case",
			inputContext: ctx,
			inputParams: inout.QuerySectionsIn{
				CategoryID:  &cid,
				CountryCode: "tw",
				Locale:      "en-us",
				PerPage:     3,
				Page:        1,
				SortBy:      "position",
				SortOrder:   "asc",
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
			description:  "testing invalid category id case",
			inputContext: ctx,
			inputParams: inout.QuerySectionsIn{
				CategoryID:  new(gographql.ID),
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
			inputParams: inout.QuerySectionsIn{
				CategoryID:  &cid,
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
			actual, err := LoadSections(tt.inputContext, tt.inputParams)

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

func TestLoadSection(t *testing.T) {
	testCases := [...]struct {
		description  string
		inputContext context.Context
		inputParams  interface{}
		expectErr    bool
		expect       *models.Section
	}{
		{
			description:  "testing normal case",
			inputContext: ctx,
			inputParams: inout.QuerySectionIn{
				SectionID:   gographql.ID("3345679"),
				CountryCode: "tw",
				Locale:      "en-us",
			},
			expectErr: false,
			expect: &models.Section{
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
			inputParams: inout.QuerySectionIn{
				SectionID:   gographql.ID("3345679"),
				CountryCode: "tw",
				Locale:      "en-us",
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description:  "testing invalid country code case",
			inputContext: ctx,
			inputParams: inout.QuerySectionIn{
				SectionID:   gographql.ID("3345679"),
				CountryCode: models.ModelsReturnErrorCountryCode,
				Locale:      "en-us",
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description:  "testing invalid section id case",
			inputContext: ctx,
			inputParams: inout.QuerySectionIn{
				SectionID:   "",
				CountryCode: "tw",
				Locale:      "en-us",
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description:  "testing section id not found case",
			inputContext: ctx,
			inputParams: inout.QuerySectionIn{
				SectionID:   gographql.ID("3345679"),
				CountryCode: models.ModelsReturnNotFoundCountryCode,
				Locale:      "en-us",
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := LoadSection(tt.inputContext, tt.inputParams)

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
