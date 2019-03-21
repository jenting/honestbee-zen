package grpc

import (
	"context"
	"testing"

	"github.com/go-test/deep"

	"github.com/honestbee/Zen/protobuf"
)

func TestSetForceSync(t *testing.T) {
	s := initServer()

	testCases := [...]struct {
		description string
		input       *protobuf.SetForceSyncRequest
		expectErr   bool
		expect      *protobuf.SetForceSyncResponse
	}{
		/*
			{
				description: "testing normal case",
				input: &protobuf.SetForceSyncRequest{
					Username: "admin",
					Password: "33456783345678",
				},
				expectErr: false,
				expect: &protobuf.SetForceSyncResponse{
					Status: "success trigger force sync job",
				},
			},
		*/
		{
			description: "testing basic auth failed case",
			input: &protobuf.SetForceSyncRequest{
				Username: "admin",
				Password: "",
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			resp, err := s.SetForceSync(context.Background(), tt.input)

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
