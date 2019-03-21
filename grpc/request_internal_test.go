package grpc

import (
	"context"
	"net/http"
	"testing"

	"github.com/go-test/deep"
	"github.com/h2non/gock"

	"github.com/honestbee/Zen/protobuf"
)

func TestCreateRequest(t *testing.T) {
	gock.New("https://honestbeehelp-tw.zendesk.com").
		Post("/api/v2/requests.json").
		BodyString(
			`{
				"request": {
					"comment": {
					  "body": "testing, please ignore!!!"
					},
					"custom_fields": [
						{
							"id": "81469808",
							"value": "test"
						}
					],
					"requester": {
					  "name": "zen project tester",
					  "email": "zen.project.tester@honestbee.com"
					},
					"subject": "testing, please ignore",
					"ticket_form_id": "956988"
				}
			}`,
		).
		Reply(http.StatusCreated)

	gock.New("https://honestbeehelp-tw.zendesk.com").
		Post("/api/v2/requests.json").
		BodyString(
			`{
				"request": {
					"comment": {
					  "body": "testing, please ignore!!!"
					},
					"requester": {
					  "name": "zen project tester",
					  "email": "zen.project.tester@honestbee.com"
					},
					"subject": "testing, please ignore"
				}
			}`,
		).
		Reply(http.StatusCreated)

	gock.New("https://honestbeehelp-tw.zendesk.com").
		Post("/api/v2/requests.json").
		BodyString(
			`{
				"request": {
					"comment": {
						"body": ""
					},
					"requester": {
					  "name": "zen project tester",
					  "email": "zen.project.tester@honestbee.com"
					},
					"subject": "testing, please ignore"
				}
			}`,
		).
		Reply(http.StatusUnprocessableEntity)

	s := initServer()

	testCases := [...]struct {
		description string
		input       *protobuf.SetCreateRequestRequest
		expectErr   bool
		expect      *protobuf.SetCreateRequestResponse
	}{
		{
			description: "testing normal case w/ ticket form and custom fields",
			input: &protobuf.SetCreateRequestRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Data: &protobuf.SetCreateRequestRequest_Data{
					Request: &protobuf.SetCreateRequestRequest_Data_Request{
						Comment: &protobuf.SetCreateRequestRequest_Data_Request_Comment{
							Body: "testing, please ignore!!!",
						},
						Requester: &protobuf.SetCreateRequestRequest_Data_Request_Requester{
							Name:  "zen project tester",
							Email: "zen.project.tester@honestbee.com",
						},
						Subject:      "testing, please ignore",
						TicketFormId: "956988",
						CustomFields: []*protobuf.SetCreateRequestRequest_Data_Request_CustomField{
							{
								Id:    "81469808",
								Value: "test",
							},
						},
					},
				},
			},
			expectErr: false,
			expect:    &protobuf.SetCreateRequestResponse{Status: "Created"},
		},
		{
			description: "testing normal case w/o ticket form and custom fields",
			input: &protobuf.SetCreateRequestRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Data: &protobuf.SetCreateRequestRequest_Data{
					Request: &protobuf.SetCreateRequestRequest_Data_Request{
						Comment: &protobuf.SetCreateRequestRequest_Data_Request_Comment{
							Body: "testing, please ignore!!!",
						},
						Requester: &protobuf.SetCreateRequestRequest_Data_Request_Requester{
							Name:  "zen project tester",
							Email: "zen.project.tester@honestbee.com",
						},
						Subject: "testing, please ignore",
					},
				},
			},
			expectErr: false,
			expect:    &protobuf.SetCreateRequestResponse{Status: "Created"},
		},
		{
			description: "testing zendesk create request failed case",
			input: &protobuf.SetCreateRequestRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Data: &protobuf.SetCreateRequestRequest_Data{
					Request: &protobuf.SetCreateRequestRequest_Data_Request{
						Requester: &protobuf.SetCreateRequestRequest_Data_Request_Requester{
							Name:  "zen project tester",
							Email: "zen.project.tester@honestbee.com",
						},
						Subject: "testing, please ignore",
					},
				},
			},
			expectErr: true,
			expect:    &protobuf.SetCreateRequestResponse{Status: "Bad Request"},
		},
		{
			description: "testing empty data case",
			input: &protobuf.SetCreateRequestRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing empty request case",
			input: &protobuf.SetCreateRequestRequest{
				CountryCode: protobuf.CountryCode_COUNTRY_CODE_TW,
				Data:        &protobuf.SetCreateRequestRequest_Data{},
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			resp, err := s.SetCreateRequest(context.Background(), tt.input)

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
