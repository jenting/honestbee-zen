// +build integration

package integration

import (
	"context"
	"testing"
	"time"

	"github.com/honestbee/Zen/models"
)

func TestModelsSyncWithDynamicContentItems(t *testing.T) {
	service := newService()
	defer service.Close()

	testCases := []struct {
		description      string
		inputForms       []*models.SyncDynamicContentItem
		inputPlaceholder string
		expectError      bool
	}{
		{
			description: "testing sync with one mock form case",
			inputForms: []*models.SyncDynamicContentItem{
				&models.SyncDynamicContentItem{
					ID:              3345678,
					URL:             "testing-dc-items-url",
					Name:            "testing",
					Placeholder:     "{{dc.testing}}",
					DefaultLocaleID: 1,
					Outdated:        false,
					CreatedAt:       time.Date(2017, 12, 19, 6, 23, 48, 0, time.UTC),
					UpdatedAt:       time.Date(2017, 12, 19, 6, 23, 48, 0, time.UTC),
					Variants:        []byte("[]"),
				},
			},
			inputPlaceholder: "{{dc.testing}}",
			expectError:      false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			err := service.SyncWithDynamicContentItems(context.Background(), tt.inputForms)
			defer resetDB()

			if tt.expectError && err == nil {
				t.Errorf("[%s] expect an error, actual none", tt.description)
			} else if !tt.expectError && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			} else if err == nil {
				_, err := service.GetDynamicContentItem(context.Background(), tt.inputPlaceholder, "en-us")
				if err == nil {
					t.Errorf("[%s] expect an error, actual nil", tt.description)
				}
			}
		})
	}
}
