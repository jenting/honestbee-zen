package handlers

import (
	"context"
	"encoding/base64"
	"net/http"
	"testing"

	"github.com/go-test/deep"
	"github.com/julienschmidt/httprouter"

	"github.com/honestbee/Zen/inout"
)

func TestGetForceSyncDecompressor(t *testing.T) {
	testCases := [...]struct {
		description string
		input1      httprouter.Params
		input2      *http.Request
		expect      interface{}
		expectErr   bool
	}{
		{
			description: "testing normal case",
			input1:      nil,
			input2: &http.Request{
				Header: http.Header{
					"Authorization": {"Basic " + base64.StdEncoding.EncodeToString([]byte("admin:1234"))},
				},
			},
			expectErr: false,
			expect: &inout.GetBasicAuthIn{
				User: "admin",
				Pwd:  "1234",
			},
		},
		{
			description: "testing fetchBasicAuth failed",
			input1:      nil,
			input2:      &http.Request{},
			expectErr:   true,
			expect:      nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := CreateForceSyncDecompressor(tt.input1, tt.input2)
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

func TestGetForceSyncHandler(t *testing.T) {
	testCases := [...]struct {
		description string
		input       interface{}
		expect      interface{}
		expectErr   bool
	}{
		/*
			{
				description: "testing normal case",
				input: &inout.GetBasicAuthIn{
					User: "admin",
					Pwd:  "33456783345678",
				},
				expectErr: false,
				expect:    "success trigger force sync job",
			},
		*/
		{
			description: "testing input casting failed case",
			input: &struct {
				name string
				age  int
			}{
				name: "honestbee",
				age:  99,
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing basic auth failed case",
			input: &inout.GetBasicAuthIn{
				User: "admin",
				Pwd:  "",
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := CreateForceSyncHandler(context.Background(), e, tt.input)
			if tt.expectErr && err == nil {
				t.Errorf("[%s] expect an error, actual nil", tt.description)
			} else if !tt.expectErr && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			} else if diff := deep.Equal(tt.expect, actual); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}
