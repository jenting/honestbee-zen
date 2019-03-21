// +build integration

package integration

import (
	"context"
	"testing"
	"time"

	"github.com/go-test/deep"

	"github.com/honestbee/Zen/models"
)

func TestModelsSyncWithTicketFields(t *testing.T) {
	service := newService()
	defer service.Close()

	testCases := []struct {
		description string
		inputForms  []*models.SyncTicketField
		inputID     int
		expectField *models.TicketField
		expectError bool
	}{
		{
			description: "testing sync with one mock form case",
			inputForms: []*models.SyncTicketField{
				&models.SyncTicketField{
					ID:                  3345678,
					URL:                 "testing-field-url",
					Type:                "testing",
					Title:               "testing",
					RawTitle:            "testing",
					Description:         "testing",
					RawDescription:      "testing",
					Position:            99,
					Active:              true,
					Required:            true,
					CollapsedForAgents:  true,
					RegexpForValidation: "{[*]}",
					TitleInPortal:       "testing",
					RawTitleInPortal:    "testing",
					VisibleInPortal:     true,
					EditableInPortal:    true,
					RequiredInPortal:    true,
					Tag:                 "",
					CreatedAt:           time.Date(1988, 10, 13, 3, 30, 59, 0, time.UTC),
					UpdatedAt:           time.Date(1988, 10, 13, 3, 30, 59, 0, time.UTC),
					Removable:           false,
					CustomFieldOptions:  []byte("[]"),
					SystemFieldOptions:  []byte("[]"),
				},
			},
			inputID: 3345678,
			expectField: &models.TicketField{
				ID:                  3345678,
				Type:                "testing",
				Title:               "testing",
				RawTitle:            "testing",
				Description:         "testing",
				RawDescription:      "testing",
				Position:            99,
				RegexpForValidation: "{[*]}",
				TitleInPortal:       "testing",
				RawTitleInPortal:    "testing",
				CreatedAt:           time.Date(1988, 10, 13, 3, 30, 59, 0, time.UTC),
				UpdatedAt:           time.Date(1988, 10, 13, 3, 30, 59, 0, time.UTC),
				CustomFieldOptions:  []*models.CustomFieldOption{},
				SystemFieldOptions:  []*models.SystemFieldOption{},
			},
			expectError: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			err := service.SyncWithTicketFields(context.Background(), tt.inputForms)
			defer resetDB()

			if tt.expectError && err == nil {
				t.Errorf("[%s] expect an error, actual none", tt.description)
			} else if !tt.expectError && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			} else if err == nil {
				actualField, err := service.GetTicketFieldByFieldID(context.Background(), tt.inputID, "en-us")
				if err != nil {
					t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
				}

				actualField.CreatedAt = actualField.CreatedAt.In(time.UTC)
				actualField.UpdatedAt = actualField.UpdatedAt.In(time.UTC)

				if diff := deep.Equal(tt.expectField, actualField); diff != nil {
					t.Errorf("[%s] %v", tt.description, diff)
				}
			}
		})
	}
}

func TestModelsGetTicketFieldByFormID(t *testing.T) {
	service := newService()
	defer service.Close()

	testCases := []struct {
		description  string
		inputFormID  int
		inputLocale  string
		expectFields []*models.TicketField
		expectError  bool
	}{
		{
			description: "testing normal case",
			inputFormID: 951408,
			inputLocale: "en-us",
			expectFields: []*models.TicketField{
				&models.TicketField{
					ID:                  24681488,
					URL:                 "https://honestbee-ph.zendesk.com/api/v2/ticket_fields/24681488.json",
					Type:                "subject",
					Title:               "Subject",
					RawTitle:            "Subject",
					Description:         "",
					RawDescription:      "",
					Position:            0,
					Active:              true,
					Required:            false,
					CollapsedForAgents:  false,
					RegexpForValidation: "",
					TitleInPortal:       "Ticket form",
					RawTitleInPortal:    "Ticket form",
					VisibleInPortal:     true,
					EditableInPortal:    true,
					RequiredInPortal:    false,
					Tag:                 "",
					CreatedAt:           time.Date(2015, 3, 2, 2, 47, 32, 0, time.UTC),
					UpdatedAt:           time.Date(2017, 12, 3, 10, 20, 4, 0, time.UTC),
					Removable:           false,
				},
				&models.TicketField{
					ID:                  24681498,
					URL:                 "https://honestbee-ph.zendesk.com/api/v2/ticket_fields/24681498.json",
					Type:                "description",
					Title:               "Description",
					RawTitle:            "Description",
					Description:         "Please elaborate on your request here.",
					RawDescription:      "Please elaborate on your request here.",
					Position:            0,
					Active:              true,
					Required:            false,
					CollapsedForAgents:  false,
					RegexpForValidation: "",
					TitleInPortal:       "Description",
					RawTitleInPortal:    "Description",
					VisibleInPortal:     true,
					EditableInPortal:    true,
					RequiredInPortal:    true,
					Tag:                 "",
					CreatedAt:           time.Date(2015, 3, 2, 2, 47, 32, 0, time.UTC),
					UpdatedAt:           time.Date(2017, 12, 1, 10, 49, 32, 0, time.UTC),
					Removable:           false,
				},
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
					CreatedAt:           time.Date(2017, 11, 22, 13, 10, 41, 0, time.UTC),
					UpdatedAt:           time.Date(2018, 1, 4, 8, 41, 29, 0, time.UTC),
					Removable:           true,
				},
			},
		},
		{
			description:  "testing non-exist form id case",
			inputFormID:  123456789,
			inputLocale:  "en-us",
			expectError:  true,
			expectFields: nil,
		},
		{
			description: "testing locale back to en-us case",
			inputFormID: 951408,
			inputLocale: "id",
			expectFields: []*models.TicketField{
				&models.TicketField{
					ID:                  24681488,
					URL:                 "https://honestbee-ph.zendesk.com/api/v2/ticket_fields/24681488.json",
					Type:                "subject",
					Title:               "Subject",
					RawTitle:            "Subject",
					Description:         "",
					RawDescription:      "",
					Position:            0,
					Active:              true,
					Required:            false,
					CollapsedForAgents:  false,
					RegexpForValidation: "",
					TitleInPortal:       "Ticket form",
					RawTitleInPortal:    "Ticket form",
					VisibleInPortal:     true,
					EditableInPortal:    true,
					RequiredInPortal:    false,
					Tag:                 "",
					CreatedAt:           time.Date(2015, 3, 2, 2, 47, 32, 0, time.UTC),
					UpdatedAt:           time.Date(2017, 12, 3, 10, 20, 4, 0, time.UTC),
					Removable:           false,
				},
				&models.TicketField{
					ID:                  24681498,
					URL:                 "https://honestbee-ph.zendesk.com/api/v2/ticket_fields/24681498.json",
					Type:                "description",
					Title:               "Description",
					RawTitle:            "Description",
					Description:         "Please elaborate on your request here.",
					RawDescription:      "Please elaborate on your request here.",
					Position:            0,
					Active:              true,
					Required:            false,
					CollapsedForAgents:  false,
					RegexpForValidation: "",
					TitleInPortal:       "Description",
					RawTitleInPortal:    "Description",
					VisibleInPortal:     true,
					EditableInPortal:    true,
					RequiredInPortal:    true,
					Tag:                 "",
					CreatedAt:           time.Date(2015, 3, 2, 2, 47, 32, 0, time.UTC),
					UpdatedAt:           time.Date(2017, 12, 1, 10, 49, 32, 0, time.UTC),
					Removable:           false,
				},
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
					CreatedAt:           time.Date(2017, 11, 22, 13, 10, 41, 0, time.UTC),
					UpdatedAt:           time.Date(2018, 1, 4, 8, 41, 29, 0, time.UTC),
					Removable:           true,
				},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actualFields, err := service.GetTicketFieldByFormID(context.Background(), tt.inputFormID, tt.inputLocale)

			if tt.expectError && err == nil {
				t.Errorf("[%s] expect an error, actual none", tt.description)
			} else if !tt.expectError && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			} else if err == nil {
				if diff := deep.Equal(tt.expectFields, actualFields); diff != nil {
					t.Errorf("[%s] %v", tt.description, diff)
				}
			}
		})
	}
}

func TestModelsGetTicketFieldCustomFieldOption(t *testing.T) {
	service := newService()
	defer service.Close()

	testCases := []struct {
		description   string
		inputFieldID  int
		expectOptions []*models.CustomFieldOption
		expectError   bool
	}{
		{
			description:  "testing normal case",
			inputFieldID: 28932938,
			expectOptions: []*models.CustomFieldOption{
				&models.CustomFieldOption{
					ID:      50431968,
					Name:    "Pre-sales",
					RawName: "Pre-sales",
					Value:   "pre_sales",
				},
				&models.CustomFieldOption{
					ID:      50431978,
					Name:    "Post-sales",
					RawName: "Post-sales",
					Value:   "post_sales",
				},
			},
		},
		{
			description:   "testing non-exist field id case",
			inputFieldID:  123456789,
			expectError:   true,
			expectOptions: nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actualOptions, err := service.GetTicketFieldCustomFieldOption(context.Background(), tt.inputFieldID)

			if tt.expectError && err == nil {
				t.Errorf("[%s] expect an error, actual none", tt.description)
			} else if !tt.expectError && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			} else if err == nil {
				if diff := deep.Equal(tt.expectOptions, actualOptions); diff != nil {
					t.Errorf("[%s] %v", tt.description, diff)
				}
			}
		})
	}
}

func TestModelsGetTicketFieldSystemFieldOption(t *testing.T) {
	service := newService()
	defer service.Close()

	testCases := []struct {
		description   string
		inputFieldID  int
		expectOptions []*models.SystemFieldOption
		expectError   bool
	}{
		{
			description:  "testing normal case",
			inputFieldID: 24681508,
			expectOptions: []*models.SystemFieldOption{
				&models.SystemFieldOption{
					Name:  "Open",
					Value: "open",
				},
				&models.SystemFieldOption{
					Name:  "Pending",
					Value: "pending",
				},
				&models.SystemFieldOption{
					Name:  "Solved",
					Value: "solved",
				},
			},
		},
		{
			description:   "testing non-exist field id case",
			inputFieldID:  123456789,
			expectError:   true,
			expectOptions: nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actualOptions, err := service.GetTicketFieldSystemFieldOption(context.Background(), tt.inputFieldID)
			defer resetDB()

			if tt.expectError && err == nil {
				t.Errorf("[%s] expect an error, actual none", tt.description)
			} else if !tt.expectError && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			} else if err == nil {
				if diff := deep.Equal(tt.expectOptions, actualOptions); diff != nil {
					t.Errorf("[%s] %v", tt.description, diff)
				}
			}
		})
	}
}
