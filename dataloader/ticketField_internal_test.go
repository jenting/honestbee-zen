package dataloader

import (
	"context"
	"testing"

	"github.com/go-test/deep"
	gographql "github.com/graph-gophers/graphql-go"

	"github.com/honestbee/Zen/inout"
	"github.com/honestbee/Zen/models"
)

func TestLoadTicketField(t *testing.T) {
	fid := gographql.ID("191908")
	efid := gographql.ID("")

	testCases := [...]struct {
		description  string
		inputContext context.Context
		inputParams  interface{}
		expectErr    bool
		expect       []*models.TicketField
	}{
		{
			description:  "testing normal en-us case",
			inputContext: ctx,
			inputParams: inout.QueryTicketFieldsIn{
				FormID: &fid,
				Locale: "en-us",
			},
			expectErr: false,
			expect: []*models.TicketField{
				&models.TicketField{
					ID:                  81469808,
					URL:                 "https://honestbee-ph.zendesk.com/api/v2/ticket_fields/81469808.json",
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
					CreatedAt:           models.FixCreatedAt1,
					UpdatedAt:           models.FixUpdatedAt1,
					Removable:           true,
					CustomFieldOptions:  make([]*models.CustomFieldOption, 0),
					SystemFieldOptions:  make([]*models.SystemFieldOption, 0),
				},
			},
		},
		{
			description:  "testing normal zh-tw case",
			inputContext: ctx,
			inputParams: inout.QueryTicketFieldsIn{
				FormID: &fid,
				Locale: "zh-tw",
			},
			expectErr: false,
			expect: []*models.TicketField{
				&models.TicketField{
					ID:                  81469808,
					URL:                 "https://honestbee-ph.zendesk.com/api/v2/ticket_fields/81469808.json",
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
					RawTitleInPortal:    "訂單號碼",
					VisibleInPortal:     true,
					EditableInPortal:    true,
					RequiredInPortal:    true,
					Tag:                 "",
					CreatedAt:           models.FixCreatedAt1,
					UpdatedAt:           models.FixUpdatedAt1,
					Removable:           true,
					CustomFieldOptions:  make([]*models.CustomFieldOption, 0),
					SystemFieldOptions:  make([]*models.SystemFieldOption, 0),
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
			inputParams: inout.QueryTicketFieldsIn{
				FormID: &fid,
				Locale: "en-us",
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description:  "testing invalid field id case",
			inputContext: ctx,
			inputParams: inout.QueryTicketFieldsIn{
				FormID: &efid,
				Locale: "en-us",
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description:  "testing locale error case",
			inputContext: ctx,
			inputParams: inout.QueryTicketFieldsIn{
				FormID: &fid,
				Locale: models.ModelsReturnErrorLocale,
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description:  "testing locale not found case",
			inputContext: ctx,
			inputParams: inout.QueryTicketFieldsIn{
				FormID: &fid,
				Locale: models.ModelsReturnNotFoundLocale,
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := LoadTicketFields(tt.inputContext, tt.inputParams)

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

func TestLoadTicketFieldCustomFieldOptions(t *testing.T) {
	testCases := [...]struct {
		description  string
		inputContext context.Context
		inputParams  interface{}
		expectErr    bool
		expect       []*models.CustomFieldOption
	}{
		{
			description:  "testing normal case",
			inputContext: ctx,
			inputParams: inout.QueryCustomFieldOptionsIn{
				FieldID: gographql.ID("81469808"),
			},
			expectErr: false,
			expect: []*models.CustomFieldOption{
				&models.CustomFieldOption{
					ID:      83592448,
					Name:    "Grocery",
					RawName: "Grocery",
					Value:   "grocery_form",
				},
				&models.CustomFieldOption{
					ID:      83592468,
					Name:    "Food",
					RawName: "Food",
					Value:   "food_form",
				},
				&models.CustomFieldOption{
					ID:      83592488,
					Name:    "Laundry",
					RawName: "Laundry",
					Value:   "laundry_form",
				},
				&models.CustomFieldOption{
					ID:      83592508,
					Name:    "Ticketing",
					RawName: "Ticketing",
					Value:   "ticketing_form",
				},
				&models.CustomFieldOption{
					ID:      83592528,
					Name:    "Rewards",
					RawName: "Rewards",
					Value:   "rewards_form",
				},
			},
		},
		{
			description:  "testing internal error case",
			inputContext: ctx,
			inputParams: inout.QueryCustomFieldOptionsIn{
				FieldID: gographql.ID("123456789"),
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description:  "testing field id not found case",
			inputContext: ctx,
			inputParams: inout.QueryCustomFieldOptionsIn{
				FieldID: gographql.ID("987654321"),
			},
			expectErr: true,
			expect:    nil,
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
				FieldID int32
			}{
				FieldID: 123,
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description:  "testing extract dataloader failed case",
			inputContext: context.TODO(),
			inputParams: inout.QueryCustomFieldOptionsIn{
				FieldID: gographql.ID("81469808"),
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description:  "testing invalid field id case",
			inputContext: ctx,
			inputParams: inout.QueryCustomFieldOptionsIn{
				FieldID: gographql.ID(""),
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := LoadTicketFieldCustomFieldOptions(tt.inputContext, tt.inputParams)

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

func TestLoadTicketFieldSystemFieldOptions(t *testing.T) {
	testCases := [...]struct {
		description  string
		inputContext context.Context
		inputParams  interface{}
		expectErr    bool
		expect       []*models.SystemFieldOption
	}{
		{
			description:  "testing normal case",
			inputContext: ctx,
			inputParams: inout.QuerySystemFieldOptionsIn{
				FieldID: gographql.ID("81469808"),
			},
			expectErr: false,
			expect: []*models.SystemFieldOption{
				&models.SystemFieldOption{},
			},
		},
		{
			description:  "testing internal error case",
			inputContext: ctx,
			inputParams: inout.QuerySystemFieldOptionsIn{
				FieldID: gographql.ID("987654"),
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description:  "testing field id not found case",
			inputContext: ctx,
			inputParams: inout.QuerySystemFieldOptionsIn{
				FieldID: gographql.ID("456789"),
			},
			expectErr: true,
			expect:    nil,
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
				FieldID int32
			}{
				FieldID: 123,
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description:  "testing extract dataloader failed case",
			inputContext: context.TODO(),
			inputParams: inout.QuerySystemFieldOptionsIn{
				FieldID: gographql.ID("81469808"),
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description:  "testing invalid field id case",
			inputContext: ctx,
			inputParams: inout.QuerySystemFieldOptionsIn{
				FieldID: gographql.ID(""),
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := LoadTicketFieldSystemFieldOptions(tt.inputContext, tt.inputParams)

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
