package grpc

import (
	"context"
	"testing"

	"github.com/go-test/deep"

	"github.com/honestbee/Zen/models"
	"github.com/honestbee/Zen/protobuf"
)

func TestGetTicketForm(t *testing.T) {
	s := initServer()

	testCases := [...]struct {
		description string
		input       *protobuf.GetTicketFormRequest
		expectErr   bool
		expect      *protobuf.GetTicketFormResponse
	}{
		{
			description: "testing normal case",
			input: &protobuf.GetTicketFormRequest{
				FormId: "191908",
			},
			expectErr: false,
			expect: &protobuf.GetTicketFormResponse{
				Id:                 "191908",
				Url:                "https://honestbeehelp-tw.zendesk.com/api/v2/ticket_forms/191908.json",
				Name:               "191908 - Default Ticket Form",
				RawName:            "191908 - Default Ticket Form",
				DisplayName:        "Default Ticket Form",
				RawDisplayName:     "Default Ticket Form",
				EndUserVisible:     true,
				Position:           0,
				Active:             true,
				InAllBrands:        true,
				RestrictedBrandIds: []int32{123, 456, 789},
				CreatedAt:          models.FixCreatedAtProto1,
				UpdatedAt:          models.FixUpdatedAtProto1,
			},
		},
		{
			description: "testing models return not found case",
			input: &protobuf.GetTicketFormRequest{
				FormId: "1919",
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing models return error case",
			input: &protobuf.GetTicketFormRequest{
				FormId: "-1",
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing invalid input case",
			input: &protobuf.GetTicketFormRequest{
				FormId: "",
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			resp, err := s.GetTicketForm(context.Background(), tt.input)

			if tt.expectErr && err == nil {
				t.Errorf("[%s] expect an error, actual == nil", tt.description)
			} else if !tt.expectErr && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			} else if diff := deep.Equal(tt.expect, resp); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}

func TestGetTicketFields(t *testing.T) {
	s := initServer()

	testCases := [...]struct {
		description string
		input       *protobuf.GetTicketFieldsRequest
		expectErr   bool
		expect      *protobuf.GetTicketFieldsResponse
	}{
		{
			description: "testing normal case",
			input: &protobuf.GetTicketFieldsRequest{
				FormId: "191908",
				Locale: protobuf.Locale_LOCALE_EN_US,
			},
			expectErr: false,
			expect: &protobuf.GetTicketFieldsResponse{
				TicketFields: []*protobuf.TicketField{
					{
						Id:                  "81469808",
						Url:                 "https://honestbee-ph.zendesk.com/api/v2/ticket_fields/81469808.json",
						Type:                "text",
						Title:               "Order Number",
						RawTitle:            "Order Number",
						Description:         "",
						RawDescription:      "",
						Position:            13,
						Active:              true,
						Required:            false,
						CollapsedForAgents:  false,
						RegexpForValidation: "",
						TitleInPortal:       "Order Number",
						RawTitleInPortal:    "Order Number",
						VisibleInPortal:     true,
						EditableInPortal:    true,
						RequiredInPortal:    true,
						Tag:                 "",
						CreatedAt:           models.FixCreatedAtProto1,
						UpdatedAt:           models.FixUpdatedAtProto1,
						Removable:           true,
						CustomFieldOptions: []*protobuf.CustomFieldOption{
							{
								Id:      "83592448",
								Name:    "Grocery",
								RawName: "Grocery",
								Value:   "grocery_form",
							},
							{
								Id:      "83592468",
								Name:    "Food",
								RawName: "Food",
								Value:   "food_form",
							},
							{
								Id:      "83592488",
								Name:    "Laundry",
								RawName: "Laundry",
								Value:   "laundry_form",
							},
							{
								Id:      "83592508",
								Name:    "Ticketing",
								RawName: "Ticketing",
								Value:   "ticketing_form",
							},
							{
								Id:      "83592528",
								Name:    "Rewards",
								RawName: "Rewards",
								Value:   "rewards_form",
							},
						},
						SystemFieldOptions: []*protobuf.SystemFieldOption{{}},
					},
				},
			},
		},
		{
			description: "testing invalid input case",
			input: &protobuf.GetTicketFieldsRequest{
				FormId: "",
				Locale: protobuf.Locale_LOCALE_EN_US,
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing models return not found case",
			input: &protobuf.GetTicketFieldsRequest{
				FormId: "191908",
				Locale: protobuf.Locale_LOCALE_ID,
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing models return error case",
			input: &protobuf.GetTicketFieldsRequest{
				FormId: "191908",
				Locale: protobuf.Locale_LOCALE_ZH_CN,
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing custom field options return not found case",
			input: &protobuf.GetTicketFieldsRequest{
				FormId: "123456",
				Locale: protobuf.Locale_LOCALE_EN_US,
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing custom field options return error case",
			input: &protobuf.GetTicketFieldsRequest{
				FormId: "123456",
				Locale: protobuf.Locale_LOCALE_ZH_TW,
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing system field options return not found case",
			input: &protobuf.GetTicketFieldsRequest{
				FormId: "654321",
				Locale: protobuf.Locale_LOCALE_EN_US,
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing system field options return error case",
			input: &protobuf.GetTicketFieldsRequest{
				FormId: "654321",
				Locale: protobuf.Locale_LOCALE_ZH_TW,
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			resp, err := s.GetTicketFields(context.Background(), tt.input)

			if tt.expectErr && err == nil {
				t.Errorf("[%s] expect an error, actual == nil", tt.description)
			} else if !tt.expectErr && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			} else if diff := deep.Equal(tt.expect, resp); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}
