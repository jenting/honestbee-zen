// +build integration

package integration

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/go-test/deep"
)

func TestHandlersGraphQLMutationCreateRequest(t *testing.T) {
	ts := newTserver()
	defer ts.closeAll()
	testCases := []struct {
		description string
		body        map[string]interface{}
		query       string
		expectBody  map[string]interface{}
	}{
		{
			description: "testing normal SG case",
			body: map[string]interface{}{
				"query": "mutation helpCenterSubmitRequest($countryCode: CountryCode, $data: RequestData!) {createRequest(countryCode: $countryCode, data: $data)}",
				"variables": map[string]interface{}{
					"countryCode": "SG",
					"data": map[string]interface{}{
						"request": map[string]interface{}{
							"comment": map[string]interface{}{
								"body": "testing, please ignore!!!",
							},
							"requester": map[string]interface{}{
								"name":  "zen project tester",
								"email": "zen.project.tester@honestbee.com",
							},
							"subject": "testing, please ignore",
						},
					},
				},
				"operationName": "helpCenterSubmitRequest",
			},
			expectBody: map[string]interface{}{
				"data": map[string]interface{}{
					"createRequest": "Created",
				},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			// Send requests.
			b, err := json.Marshal(tt.body)
			if err != nil {
				t.Fatalf("[%s] json marshal failed:%v", tt.description, err)
			}
			resp, err := ts.Client().Post(ts.URL+"/graphql", "application/json", ioutil.NopCloser(bytes.NewReader(b)))
			if err != nil {
				t.Fatalf("[%s] http client get failed:%v", tt.description, err)
			}
			defer resp.Body.Close()

			// Compare HTTP status code.
			if http.StatusOK != resp.StatusCode {
				t.Errorf("[%s] http status expect:%v != actual:%v", tt.description, http.StatusOK, resp.StatusCode)
			}

			// Compare HTTP body.
			actual := make(map[string]interface{})
			if err = json.NewDecoder(resp.Body).Decode(&actual); err != nil {
				t.Fatalf("[%s] json decoding failed:%v", tt.description, err)
			}
			// Converts integer to the same type.
			expectData, err := json.Marshal(tt.expectBody)
			if err != nil {
				t.Fatalf("[%s] json marshal failed:%v", tt.description, err)
			}
			expect := make(map[string]interface{})
			if err = json.Unmarshal(expectData, &expect); err != nil {
				t.Fatalf("[%s] json unmarshal failed:%v", tt.description, err)
			}
			// Compares and prints difference.
			if diff := deep.Equal(expect, actual); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
		})
	}
}
