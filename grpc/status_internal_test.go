package grpc

import (
	"context"
	"testing"

	"github.com/honestbee/Zen/protobuf"
)

func TestGetStatus(t *testing.T) {
	s := initServer()

	testCases := [...]struct {
		description string
		input       *protobuf.GetStatusRequest
		expectErr   bool
	}{
		{
			description: "testing normal case",
			input:       &protobuf.GetStatusRequest{},
			expectErr:   false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			resp, err := s.GetStatus(context.Background(), tt.input)

			if tt.expectErr && err == nil {
				t.Errorf("[%s] expect an error, actual == nil", tt.description)
			} else if !tt.expectErr && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			} else {
				if resp.GoVersion == "" {
					t.Errorf("[%s] should not be empty", tt.description)
				}
				if resp.AppVersion == "" {
					t.Errorf("[%s] should not be empty", tt.description)
				}
			}
		})
	}
}
