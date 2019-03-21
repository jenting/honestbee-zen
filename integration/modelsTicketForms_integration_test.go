// +build integration

package integration

import (
	"context"
	"testing"
	"time"

	"github.com/go-test/deep"

	"github.com/honestbee/Zen/models"
)

func TestModelsSyncWithTicketForms(t *testing.T) {
	service := newService()
	defer service.Close()

	testCases := []struct {
		description string
		inputForms  []*models.SyncTicketForm
		inputID     int
		expectForm  *models.TicketForm
		expectError bool
	}{
		{
			description: "testing sync with one mock form case",
			inputForms: []*models.SyncTicketForm{
				&models.SyncTicketForm{
					ID:                 3345678,
					URL:                "testing-form-url",
					Name:               "testing",
					RawName:            "testing",
					DisplayName:        "testing",
					RawDisplayName:     "testing",
					EndUserVisible:     true,
					Position:           100,
					Active:             true,
					InAllBrands:        true,
					RestrictedBrandIDs: nil,
					TicketFieldIDs:     nil,
					CreatedAt:          time.Date(1988, 10, 13, 3, 30, 59, 0, time.UTC),
					UpdatedAt:          time.Date(1988, 10, 13, 3, 30, 59, 0, time.UTC),
				},
			},
			inputID: 3345678,
			expectForm: &models.TicketForm{
				ID:                 3345678,
				URL:                "testing-form-url",
				Name:               "testing",
				RawName:            "testing",
				DisplayName:        "testing",
				RawDisplayName:     "testing",
				EndUserVisible:     true,
				Position:           100,
				Active:             true,
				InAllBrands:        true,
				RestrictedBrandIDs: []int64{},
				TicketFields:       []*models.TicketField{},
				CreatedAt:          time.Date(1988, 10, 13, 3, 30, 59, 0, time.UTC),
				UpdatedAt:          time.Date(1988, 10, 13, 3, 30, 59, 0, time.UTC),
			},
			expectError: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			err := service.SyncWithTicketForms(context.Background(), tt.inputForms)
			defer resetDB()

			if tt.expectError && err == nil {
				t.Errorf("[%s] expect an error, actual none", tt.description)
			} else if !tt.expectError && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			} else if err == nil {
				actualForm, err := service.GetTicketForm(context.Background(), tt.inputID, "en-us")
				if err != nil {
					t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
				}

				actualForm.CreatedAt = actualForm.CreatedAt.In(time.UTC)
				actualForm.UpdatedAt = actualForm.UpdatedAt.In(time.UTC)

				if diff := deep.Equal(tt.expectForm, actualForm); diff != nil {
					t.Errorf("[%s] %v", tt.description, diff)
				}
			}
		})
	}
}
