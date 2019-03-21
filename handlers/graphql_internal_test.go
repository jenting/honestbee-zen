package handlers

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/go-test/deep"
	"github.com/julienschmidt/httprouter"

	"github.com/honestbee/Zen/inout"
)

func TestCreateGraphQLDecompressor(t *testing.T) {
	testCases := [...]struct {
		description string
		input1      httprouter.Params
		input2      *http.Request
		expect      interface{}
		expectErr   bool
	}{
		{
			description: "testing normal request case",
			input1:      nil,
			input2: &http.Request{
				Body: ioutil.NopCloser(bytes.NewBufferString("{\"query\":\"query IntrospectionQuery { __schema { queryType { name } mutationType { name } subscriptionType { name } types { ...FullType } directives { name description locations args { ...InputValue } } }}fragment FullType on __Type { kind name description fields(includeDeprecated: true) { name description args { ...InputValue } type { ...TypeRef } isDeprecated deprecationReason } inputFields { ...InputValue } interfaces { ...TypeRef } enumValues(includeDeprecated: true) { name description isDeprecated deprecationReason } possibleTypes { ...TypeRef }}fragment InputValue on __InputValue { name description type { ...TypeRef } defaultValue}fragment TypeRef on __Type { kind name ofType { kind name ofType { kind name ofType { kind name ofType { kind name ofType { kind name ofType { kind name ofType { kind name } } } } } } }}\"}")),
			},
			expectErr: false,
			expect: &inout.GraphQLIn{
				Ctx: context.Background(),
				Queries: []inout.GraphQLQuery{
					inout.GraphQLQuery{
						Query: "query IntrospectionQuery { __schema { queryType { name } mutationType { name } subscriptionType { name } types { ...FullType } directives { name description locations args { ...InputValue } } }}fragment FullType on __Type { kind name description fields(includeDeprecated: true) { name description args { ...InputValue } type { ...TypeRef } isDeprecated deprecationReason } inputFields { ...InputValue } interfaces { ...TypeRef } enumValues(includeDeprecated: true) { name description isDeprecated deprecationReason } possibleTypes { ...TypeRef }}fragment InputValue on __InputValue { name description type { ...TypeRef } defaultValue}fragment TypeRef on __Type { kind name ofType { kind name ofType { kind name ofType { kind name ofType { kind name ofType { kind name ofType { kind name ofType { kind name } } } } } } }}",
					},
				},
				IsBatch: false,
			},
		},
		{
			description: "testing batch request case",
			input1:      nil,
			input2: &http.Request{
				Body: ioutil.NopCloser(bytes.NewBufferString("[{\"query\":\"query1\"},{\"query\":\"query2\"}]")),
			},
			expectErr: false,
			expect: &inout.GraphQLIn{
				Ctx: context.Background(),
				Queries: []inout.GraphQLQuery{
					inout.GraphQLQuery{
						Query: "query1",
					},
					inout.GraphQLQuery{
						Query: "query2",
					},
				},
				IsBatch: true,
			},
		},
		{
			description: "testing invalid first request case",
			input1:      nil,
			input2: &http.Request{
				Body: ioutil.NopCloser(bytes.NewBufferString("(\"\":\"\")")),
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing normal request json unmarshal failed case",
			input1:      nil,
			input2: &http.Request{
				Body: ioutil.NopCloser(bytes.NewBufferString("{\"query\":\"\"")),
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing batch request json unmarshal failed case",
			input1:      nil,
			input2: &http.Request{
				Body: ioutil.NopCloser(bytes.NewBufferString("[{\"query\":\"query1\",{\"query\":\"query2\"]")),
			},
			expectErr: true,
			expect:    nil,
		},
		{
			description: "testing read HTTP request body failed",
			input1:      nil,
			input2: &http.Request{
				Body: ioutil.NopCloser(bytes.NewBufferString("")),
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := CreateGraphQLDecompressor(tt.input1, tt.input2)
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
