package dataloader

import (
	"context"
	"testing"

	"github.com/go-test/deep"
	gographql "github.com/graph-gophers/graphql-go"

	"github.com/honestbee/Zen/inout"
	"github.com/honestbee/Zen/models"
)

func TestLoadTicketForm(t *testing.T) {
	testCases := [...]struct {
		description  string
		inputContext context.Context
		inputParams  interface{}
		expectErr    bool
		expect       *models.SyncTicketForm
	}{
		{
			description:  "testing normal case",
			inputContext: ctx,
			inputParams: inout.QueryTicketFormIn{
				FormID: gographql.ID("191908"),
			},
			expectErr: false,
			expect: &models.SyncTicketForm{
				ID:                 191908,
				URL:                "https://honestbeehelp-tw.zendesk.com/api/v2/ticket_forms/191908.json",
				Name:               "191908 - Default Ticket Form",
				RawName:            "191908 - Default Ticket Form",
				DisplayName:        "Default Ticket Form",
				RawDisplayName:     "Default Ticket Form",
				EndUserVisible:     true,
				Position:           0,
				Active:             true,
				InAllBrands:        true,
				RestrictedBrandIDs: []int64{123, 456, 789},
				TicketFieldIDs:     []int64{24681488, 24681498},
				CreatedAt:          models.FixCreatedAt1,
				UpdatedAt:          models.FixUpdatedAt1,
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
			inputParams: inout.QueryTicketFormIn{
				FormID: gographql.ID("191908"),
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description:  "testing invalid article id case",
			inputContext: ctx,
			inputParams: inout.QueryTicketFormIn{
				FormID: gographql.ID(""),
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := LoadTicketForm(tt.inputContext, tt.inputParams)

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
