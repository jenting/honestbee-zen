// +build integration

package integration

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/go-test/deep"
)

func TestHandlersGraphQLQueryTicketForm(t *testing.T) {
	ts := newTserver()
	defer ts.closeAll()
	testCases := []struct {
		description string
		body        map[string]interface{}
		query       string
		expectBody  map[string]interface{}
	}{
		{
			description: "testing normal ticket form tw + en-us case",
			body: map[string]interface{}{
				"query": `
				{
					oneTicketForm(formId: "825847") {
						id
						url
						name
						rawName
						displayName
						rawDisplayName
						endUserVisible
						position
						active
						inAllBrands
						restrictedBrandIds
						createdAt
						updatedAt
						ticketFieldsConnection {
							id
							url
							type
							title
							rawTitle
							description
							rawDescription
							position
							active
							required
							collapsedForAgents
							regexpForValidation
							titleInPortal
							rawTitleInPortal
							visibleInPortal
							editableInPortal
							requiredInPortal
							tag
							createdAt
							updatedAt
							removable
							customFieldOptions {
								id
								name
								rawName
								value
							}
							systemFieldOptions {
								name
								value
							}
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"data": map[string]interface{}{
					"oneTicketForm": map[string]interface{}{
						"id":                 "825847",
						"url":                "https://honestbeehelp-tw.zendesk.com/api/v2/ticket_forms/825847.json",
						"name":               "shin - wrong/defect item",
						"rawName":            "shin - wrong/defect item",
						"displayName":        "shin - wrong/defect item",
						"rawDisplayName":     "shin - wrong/defect item",
						"endUserVisible":     true,
						"position":           52,
						"active":             false,
						"inAllBrands":        true,
						"restrictedBrandIds": []int{},
						"createdAt":          time.Date(2017, 11, 9, 17, 49, 32, 0, time.UTC),
						"updatedAt":          time.Date(2018, 7, 13, 7, 40, 9, 0, time.UTC),
						"ticketFieldsConnection": []map[string]interface{}{
							map[string]interface{}{
								"id":                  "24681488",
								"url":                 "https://honestbee-ph.zendesk.com/api/v2/ticket_fields/24681488.json",
								"type":                "subject",
								"title":               "Subject",
								"rawTitle":            "Subject",
								"description":         "",
								"rawDescription":      "",
								"position":            0,
								"active":              true,
								"required":            false,
								"collapsedForAgents":  false,
								"regexpForValidation": "",
								"titleInPortal":       "Ticket form",
								"rawTitleInPortal":    "Ticket form",
								"visibleInPortal":     true,
								"editableInPortal":    true,
								"requiredInPortal":    false,
								"tag":                 "",
								"createdAt":           time.Date(2015, 3, 2, 2, 47, 32, 0, time.UTC),
								"updatedAt":           time.Date(2017, 12, 3, 10, 20, 4, 0, time.UTC),
								"removable":           false,
								"customFieldOptions":  []map[string]interface{}{},
								"systemFieldOptions":  []map[string]interface{}{},
							},
							map[string]interface{}{
								"id":                  "81469808",
								"url":                 "https://honestbee-ph.zendesk.com/api/v2/ticket_fields/81469808.json",
								"type":                "text",
								"title":               "Order Number",
								"rawTitle":            "Order Number",
								"description":         "",
								"rawDescription":      "",
								"position":            13,
								"active":              true,
								"required":            false,
								"collapsedForAgents":  false,
								"regexpForValidation": "",
								"titleInPortal":       "Order Number",
								"rawTitleInPortal":    "Order Number",
								"visibleInPortal":     true,
								"editableInPortal":    true,
								"requiredInPortal":    true,
								"tag":                 "",
								"createdAt":           time.Date(2017, 11, 22, 13, 10, 41, 0, time.UTC),
								"updatedAt":           time.Date(2018, 1, 4, 8, 41, 29, 0, time.UTC),
								"removable":           true,
								"customFieldOptions":  []map[string]interface{}{},
								"systemFieldOptions":  []map[string]interface{}{},
							},
							map[string]interface{}{
								"id":                  "81421968",
								"url":                 "https://honestbeehelp-tw.zendesk.com/api/v2/ticket_fields/81421968.json",
								"type":                "tagger",
								"title":               "Type of service",
								"rawTitle":            "Type of service",
								"description":         "",
								"rawDescription":      "",
								"position":            15,
								"active":              true,
								"required":            false,
								"collapsedForAgents":  false,
								"regexpForValidation": "",
								"titleInPortal":       "Type of service",
								"rawTitleInPortal":    "Type of service",
								"visibleInPortal":     true,
								"editableInPortal":    true,
								"requiredInPortal":    true,
								"tag":                 "",
								"createdAt":           time.Date(2017, 11, 9, 17, 41, 55, 0, time.UTC),
								"updatedAt":           time.Date(2018, 3, 28, 10, 13, 3, 0, time.UTC),
								"removable":           true,
								"customFieldOptions": []map[string]interface{}{
									map[string]interface{}{
										"id":      "83592448",
										"name":    "Grocery",
										"rawName": "Grocery",
										"value":   "grocery_form",
									},
									map[string]interface{}{
										"id":      "83592468",
										"name":    "Food",
										"rawName": "Food",
										"value":   "food_form",
									},
									map[string]interface{}{
										"id":      "83592488",
										"name":    "Laundry",
										"rawName": "Laundry",
										"value":   "laundry_form",
									},
									map[string]interface{}{
										"id":      "83592508",
										"name":    "Ticketing",
										"rawName": "Ticketing",
										"value":   "ticketing_form",
									},
									map[string]interface{}{
										"id":      "83592528",
										"name":    "Rewards",
										"rawName": "Rewards",
										"value":   "rewards_form",
									},
								},
								"systemFieldOptions": []map[string]interface{}{},
							},
						},
					},
				},
			},
		},
		{
			description: "testing normal ticket form tw + zh-tw case",
			body: map[string]interface{}{
				"query": `
				{
					oneTicketForm(formId: "825847") {
						id
						url
						name
						rawName
						displayName
						rawDisplayName
						endUserVisible
						position
						active
						inAllBrands
						restrictedBrandIds
						createdAt
						updatedAt
						ticketFieldsConnection(locale: ZH_TW) {
							id
							url
							type
							title
							rawTitle
							description
							rawDescription
							position
							active
							required
							collapsedForAgents
							regexpForValidation
							titleInPortal
							rawTitleInPortal
							visibleInPortal
							editableInPortal
							requiredInPortal
							tag
							createdAt
							updatedAt
							removable
							customFieldOptions {
								id
								name
								rawName
								value
							}
							systemFieldOptions {
								name
								value
							}
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"data": map[string]interface{}{
					"oneTicketForm": map[string]interface{}{
						"id":                 "825847",
						"url":                "https://honestbeehelp-tw.zendesk.com/api/v2/ticket_forms/825847.json",
						"name":               "shin - wrong/defect item",
						"rawName":            "shin - wrong/defect item",
						"displayName":        "shin - wrong/defect item",
						"rawDisplayName":     "shin - wrong/defect item",
						"endUserVisible":     true,
						"position":           52,
						"active":             false,
						"inAllBrands":        true,
						"restrictedBrandIds": []int{},
						"createdAt":          time.Date(2017, 11, 9, 17, 49, 32, 0, time.UTC),
						"updatedAt":          time.Date(2018, 7, 13, 7, 40, 9, 0, time.UTC),
						"ticketFieldsConnection": []map[string]interface{}{
							map[string]interface{}{
								"id":                  "24681488",
								"url":                 "https://honestbee-ph.zendesk.com/api/v2/ticket_fields/24681488.json",
								"type":                "subject",
								"title":               "Subject",
								"rawTitle":            "Subject",
								"description":         "",
								"rawDescription":      "",
								"position":            0,
								"active":              true,
								"required":            false,
								"collapsedForAgents":  false,
								"regexpForValidation": "",
								"titleInPortal":       "Ticket form",
								"rawTitleInPortal":    "Ticket form",
								"visibleInPortal":     true,
								"editableInPortal":    true,
								"requiredInPortal":    false,
								"tag":                 "",
								"createdAt":           time.Date(2015, 3, 2, 2, 47, 32, 0, time.UTC),
								"updatedAt":           time.Date(2017, 12, 3, 10, 20, 4, 0, time.UTC),
								"removable":           false,
								"customFieldOptions":  []map[string]interface{}{},
								"systemFieldOptions":  []map[string]interface{}{},
							},
							map[string]interface{}{
								"id":                  "81469808",
								"url":                 "https://honestbee-ph.zendesk.com/api/v2/ticket_fields/81469808.json",
								"type":                "text",
								"title":               "Order Number",
								"rawTitle":            "Order Number",
								"description":         "",
								"rawDescription":      "",
								"position":            13,
								"active":              true,
								"required":            false,
								"collapsedForAgents":  false,
								"regexpForValidation": "",
								"titleInPortal":       "Order Number",
								"rawTitleInPortal":    "訂單號碼",
								"visibleInPortal":     true,
								"editableInPortal":    true,
								"requiredInPortal":    true,
								"tag":                 "",
								"createdAt":           time.Date(2017, 11, 22, 13, 10, 41, 0, time.UTC),
								"updatedAt":           time.Date(2018, 1, 4, 8, 41, 29, 0, time.UTC),
								"removable":           true,
								"customFieldOptions":  []map[string]interface{}{},
								"systemFieldOptions":  []map[string]interface{}{},
							},
							map[string]interface{}{
								"id":                  "81421968",
								"url":                 "https://honestbeehelp-tw.zendesk.com/api/v2/ticket_fields/81421968.json",
								"type":                "tagger",
								"title":               "Type of service",
								"rawTitle":            "Type of service",
								"description":         "",
								"rawDescription":      "",
								"position":            15,
								"active":              true,
								"required":            false,
								"collapsedForAgents":  false,
								"regexpForValidation": "",
								"titleInPortal":       "Type of service",
								"rawTitleInPortal":    "服務種類",
								"visibleInPortal":     true,
								"editableInPortal":    true,
								"requiredInPortal":    true,
								"tag":                 "",
								"createdAt":           time.Date(2017, 11, 9, 17, 41, 55, 0, time.UTC),
								"updatedAt":           time.Date(2018, 3, 28, 10, 13, 3, 0, time.UTC),
								"removable":           true,
								"customFieldOptions": []map[string]interface{}{
									map[string]interface{}{
										"id":      "83592448",
										"name":    "Grocery",
										"rawName": "Grocery",
										"value":   "grocery_form",
									},
									map[string]interface{}{
										"id":      "83592468",
										"name":    "Food",
										"rawName": "Food",
										"value":   "food_form",
									},
									map[string]interface{}{
										"id":      "83592488",
										"name":    "Laundry",
										"rawName": "Laundry",
										"value":   "laundry_form",
									},
									map[string]interface{}{
										"id":      "83592508",
										"name":    "Ticketing",
										"rawName": "Ticketing",
										"value":   "ticketing_form",
									},
									map[string]interface{}{
										"id":      "83592528",
										"name":    "Rewards",
										"rawName": "Rewards",
										"value":   "rewards_form",
									},
								},
								"systemFieldOptions": []map[string]interface{}{},
							},
						},
					},
				},
			},
		},
		{
			description: "testing not support locale goes default case",
			body: map[string]interface{}{
				"query": `
				{
					oneTicketForm(formId: "825847") {
						id
						url
						name
						rawName
						displayName
						rawDisplayName
						endUserVisible
						position
						active
						inAllBrands
						restrictedBrandIds
						createdAt
						updatedAt
						ticketFieldsConnection(locale: ID) {
							id
							url
							type
							title
							rawTitle
							description
							rawDescription
							position
							active
							required
							collapsedForAgents
							regexpForValidation
							titleInPortal
							rawTitleInPortal
							visibleInPortal
							editableInPortal
							requiredInPortal
							tag
							createdAt
							updatedAt
							removable
							customFieldOptions {
								id
								name
								rawName
								value
							}
							systemFieldOptions {
								name
								value
							}
						}
					}
				}				  
				`,
			},
			expectBody: map[string]interface{}{
				"data": map[string]interface{}{
					"oneTicketForm": map[string]interface{}{
						"id":                 "825847",
						"url":                "https://honestbeehelp-tw.zendesk.com/api/v2/ticket_forms/825847.json",
						"name":               "shin - wrong/defect item",
						"rawName":            "shin - wrong/defect item",
						"displayName":        "shin - wrong/defect item",
						"rawDisplayName":     "shin - wrong/defect item",
						"endUserVisible":     true,
						"position":           52,
						"active":             false,
						"inAllBrands":        true,
						"restrictedBrandIds": []int{},
						"createdAt":          time.Date(2017, 11, 9, 17, 49, 32, 0, time.UTC),
						"updatedAt":          time.Date(2018, 7, 13, 7, 40, 9, 0, time.UTC),
						"ticketFieldsConnection": []map[string]interface{}{
							map[string]interface{}{
								"id":                  "24681488",
								"url":                 "https://honestbee-ph.zendesk.com/api/v2/ticket_fields/24681488.json",
								"type":                "subject",
								"title":               "Subject",
								"rawTitle":            "Subject",
								"description":         "",
								"rawDescription":      "",
								"position":            0,
								"active":              true,
								"required":            false,
								"collapsedForAgents":  false,
								"regexpForValidation": "",
								"titleInPortal":       "Ticket form",
								"rawTitleInPortal":    "Ticket form",
								"visibleInPortal":     true,
								"editableInPortal":    true,
								"requiredInPortal":    false,
								"tag":                 "",
								"createdAt":           time.Date(2015, 3, 2, 2, 47, 32, 0, time.UTC),
								"updatedAt":           time.Date(2017, 12, 3, 10, 20, 4, 0, time.UTC),
								"removable":           false,
								"customFieldOptions":  []map[string]interface{}{},
								"systemFieldOptions":  []map[string]interface{}{},
							},
							map[string]interface{}{
								"id":                  "81469808",
								"url":                 "https://honestbee-ph.zendesk.com/api/v2/ticket_fields/81469808.json",
								"type":                "text",
								"title":               "Order Number",
								"rawTitle":            "Order Number",
								"description":         "",
								"rawDescription":      "",
								"position":            13,
								"active":              true,
								"required":            false,
								"collapsedForAgents":  false,
								"regexpForValidation": "",
								"titleInPortal":       "Order Number",
								"rawTitleInPortal":    "Order Number",
								"visibleInPortal":     true,
								"editableInPortal":    true,
								"requiredInPortal":    true,
								"tag":                 "",
								"createdAt":           time.Date(2017, 11, 22, 13, 10, 41, 0, time.UTC),
								"updatedAt":           time.Date(2018, 1, 4, 8, 41, 29, 0, time.UTC),
								"removable":           true,
								"customFieldOptions":  []map[string]interface{}{},
								"systemFieldOptions":  []map[string]interface{}{},
							},
							map[string]interface{}{
								"id":                  "81421968",
								"url":                 "https://honestbeehelp-tw.zendesk.com/api/v2/ticket_fields/81421968.json",
								"type":                "tagger",
								"title":               "Type of service",
								"rawTitle":            "Type of service",
								"description":         "",
								"rawDescription":      "",
								"position":            15,
								"active":              true,
								"required":            false,
								"collapsedForAgents":  false,
								"regexpForValidation": "",
								"titleInPortal":       "Type of service",
								"rawTitleInPortal":    "Type of service",
								"visibleInPortal":     true,
								"editableInPortal":    true,
								"requiredInPortal":    true,
								"tag":                 "",
								"createdAt":           time.Date(2017, 11, 9, 17, 41, 55, 0, time.UTC),
								"updatedAt":           time.Date(2018, 3, 28, 10, 13, 3, 0, time.UTC),
								"removable":           true,
								"customFieldOptions": []map[string]interface{}{
									map[string]interface{}{
										"id":      "83592448",
										"name":    "Grocery",
										"rawName": "Grocery",
										"value":   "grocery_form",
									},
									map[string]interface{}{
										"id":      "83592468",
										"name":    "Food",
										"rawName": "Food",
										"value":   "food_form",
									},
									map[string]interface{}{
										"id":      "83592488",
										"name":    "Laundry",
										"rawName": "Laundry",
										"value":   "laundry_form",
									},
									map[string]interface{}{
										"id":      "83592508",
										"name":    "Ticketing",
										"rawName": "Ticketing",
										"value":   "ticketing_form",
									},
									map[string]interface{}{
										"id":      "83592528",
										"name":    "Rewards",
										"rawName": "Rewards",
										"value":   "rewards_form",
									},
								},
								"systemFieldOptions": []map[string]interface{}{},
							},
						},
					},
				},
			},
		},
		{
			description: "testing not exist form id case",
			body: map[string]interface{}{
				"query": `
				{
					oneTicketForm(formId: "91293129847") {
						id
						url
						name
						rawName
						displayName
						rawDisplayName
						endUserVisible
						position
						active
						inAllBrands
						restrictedBrandIds
						createdAt
						updatedAt
						ticketFieldsConnection(locale: ID) {
							id
							url
							type
							title
							rawTitle
							description
							rawDescription
							position
							active
							required
							collapsedForAgents
							regexpForValidation
							titleInPortal
							rawTitleInPortal
							visibleInPortal
							editableInPortal
							requiredInPortal
							tag
							createdAt
							updatedAt
							removable
							customFieldOptions {
								id
								name
								rawName
								value
							}
							systemFieldOptions {
								name
								value
							}
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"data": map[string]interface{}{
					"oneTicketForm": nil,
				},
				"errors": []interface{}{
					map[string]interface{}{
						"message": "Record Not Found",
						"path": []interface{}{
							"oneTicketForm",
						},
					},
				},
			},
		},
		{
			description: "testing not exist locale case",
			body: map[string]interface{}{
				"query": `
				{
					oneTicketForm(formId: "91293129847") {
						id
						url
						name
						rawName
						displayName
						rawDisplayName
						endUserVisible
						position
						active
						inAllBrands
						restrictedBrandIds
						createdAt
						updatedAt
						ticketFieldsConnection(locale: non_exist_locale) {
							id
							url
							type
							title
							rawTitle
							description
							rawDescription
							position
							active
							required
							collapsedForAgents
							regexpForValidation
							titleInPortal
							rawTitleInPortal
							visibleInPortal
							editableInPortal
							requiredInPortal
							tag
							createdAt
							updatedAt
							removable
							customFieldOptions {
								id
								name
								rawName
								value
							}
							systemFieldOptions {
								name
								value
							}
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"errors": []interface{}{
					map[string]interface{}{
						"message": "Argument \"locale\" has invalid value non_exist_locale.\nExpected type \"Locale\", found non_exist_locale.",
						"locations": []interface{}{
							map[string]interface{}{
								"line":   17,
								"column": 38,
							},
						},
					},
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
