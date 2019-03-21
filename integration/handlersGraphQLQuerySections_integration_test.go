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

func TestHandlersGraphQLQuerySections(t *testing.T) {
	ts := newTserver()
	defer ts.closeAll()
	testCases := []struct {
		description string
		body        map[string]interface{}
		query       string
		expectBody  map[string]interface{}
	}{
		{
			description: "testing normal SG + EN_US case",
			body: map[string]interface{}{
				"query": `
				{
					allSections {
						page
						perPage
						pageCount
						count
						sections {
							id
							position
							createdAt
							updatedAt
							sourceLocale
							outdated
							countryCode
							url
							htmlUrl
							name
							description
							locale
							categoryConnection {
								id
								name
								keyName
							}
							articlesConnection {
								page
								perPage
								pageCount
								count
							}
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"data": map[string]interface{}{
					"allSections": map[string]interface{}{
						"page":      1,
						"perPage":   30,
						"pageCount": 0,
						"count":     0,
						"sections":  []interface{}{},
					},
				},
			},
		},
		{
			description: "testing normal TW + EN_US case",
			body: map[string]interface{}{
				"query": `
				{
					allSections(countryCode: TW, locale: EN_US) {
						page
						perPage
						pageCount
						count
						sections {
							id
							position
							createdAt
							updatedAt
							sourceLocale
							outdated
							countryCode
							url
							htmlUrl
							name
							description
							locale
							categoryConnection {
								id
								name
								keyName
							}
							articlesConnection {
								page
								perPage
								pageCount
								count
							}
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"data": map[string]interface{}{
					"allSections": map[string]interface{}{
						"page":      1,
						"perPage":   30,
						"pageCount": 1,
						"count":     1,
						"sections": []interface{}{
							map[string]interface{}{
								"id":           "115004118448",
								"position":     0,
								"createdAt":    time.Date(2017, 12, 19, 6, 23, 48, 0, time.UTC),
								"updatedAt":    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
								"sourceLocale": "en-us",
								"outdated":     false,
								"countryCode":  "tw",
								"url":          "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/sections/115004118448-I-need-help-with-my-account.json",
								"htmlUrl":      "https://help.honestbee.tw/hc/en-us/sections/115004118448-I-need-help-with-my-account",
								"name":         "I need help with my account",
								"description":  "",
								"locale":       "en-us",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"articlesConnection": map[string]interface{}{
									"page":      1,
									"perPage":   30,
									"pageCount": 1,
									"count":     5,
								},
							},
						},
					},
				},
			},
		},
		{
			description: "testing normal TW + ZH_TW case",
			body: map[string]interface{}{
				"query": `
				{
					allSections(countryCode: TW, locale: ZH_TW) {
						page
						perPage
						pageCount
						count
						sections {
							id
							position
							createdAt
							updatedAt
							sourceLocale
							outdated
							countryCode
							url
							htmlUrl
							name
							description
							locale
							categoryConnection {
								id
								name
								keyName
							}
							articlesConnection {
								page
								perPage
								pageCount
								count
							}
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"data": map[string]interface{}{
					"allSections": map[string]interface{}{
						"page":      1,
						"perPage":   30,
						"pageCount": 1,
						"count":     1,
						"sections": []interface{}{
							map[string]interface{}{
								"id":           "115004118448",
								"position":     0,
								"createdAt":    time.Date(2017, 12, 19, 6, 23, 48, 0, time.UTC),
								"updatedAt":    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
								"sourceLocale": "en-us",
								"outdated":     false,
								"countryCode":  "tw",
								"url":          "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/zh-tw/sections/115004118448-%E6%88%91%E9%9C%80%E8%A6%81%E5%B8%B3%E8%99%9F%E7%9B%B8%E9%97%9C%E7%9A%84%E5%8D%94%E5%8A%A9.json",
								"htmlUrl":      "https://help.honestbee.tw/hc/zh-tw/sections/115004118448-%E6%88%91%E9%9C%80%E8%A6%81%E5%B8%B3%E8%99%9F%E7%9B%B8%E9%97%9C%E7%9A%84%E5%8D%94%E5%8A%A9",
								"name":         "我需要帳號相關的協助",
								"description":  "",
								"locale":       "zh-tw",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"articlesConnection": map[string]interface{}{
									"page":      1,
									"perPage":   30,
									"pageCount": 1,
									"count":     5,
								},
							},
						},
					},
				},
			},
		},
		{
			description: "testing sort by created at case",
			body: map[string]interface{}{
				"query": `
				{
					allSections(countryCode: TW, locale: EN_US, sortBy: CREATED_AT) {
						page
						perPage
						pageCount
						count
						sections {
							id
							position
							createdAt
							updatedAt
							sourceLocale
							outdated
							countryCode
							url
							htmlUrl
							name
							description
							locale
							categoryConnection {
								id
								name
								keyName
							}
							articlesConnection {
								page
								perPage
								pageCount
								count
							}
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"data": map[string]interface{}{
					"allSections": map[string]interface{}{
						"page":      1,
						"perPage":   30,
						"pageCount": 1,
						"count":     1,
						"sections": []interface{}{
							map[string]interface{}{
								"id":           "115004118448",
								"position":     0,
								"createdAt":    time.Date(2017, 12, 19, 6, 23, 48, 0, time.UTC),
								"updatedAt":    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
								"sourceLocale": "en-us",
								"outdated":     false,
								"countryCode":  "tw",
								"url":          "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/sections/115004118448-I-need-help-with-my-account.json",
								"htmlUrl":      "https://help.honestbee.tw/hc/en-us/sections/115004118448-I-need-help-with-my-account",
								"name":         "I need help with my account",
								"description":  "",
								"locale":       "en-us",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"articlesConnection": map[string]interface{}{
									"page":      1,
									"perPage":   30,
									"pageCount": 1,
									"count":     5,
								},
							},
						},
					},
				},
			},
		},
		{
			description: "testing sort by updated at case",
			body: map[string]interface{}{
				"query": `
				{
					allSections(countryCode: TW, locale: EN_US, sortBy: UPDATED_AT) {
						page
						perPage
						pageCount
						count
						sections {
							id
							position
							createdAt
							updatedAt
							sourceLocale
							outdated
							countryCode
							url
							htmlUrl
							name
							description
							locale
							categoryConnection {
								id
								name
								keyName
							}
							articlesConnection {
								page
								perPage
								pageCount
								count
							}
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"data": map[string]interface{}{
					"allSections": map[string]interface{}{
						"page":      1,
						"perPage":   30,
						"pageCount": 1,
						"count":     1,
						"sections": []interface{}{
							map[string]interface{}{
								"id":           "115004118448",
								"position":     0,
								"createdAt":    time.Date(2017, 12, 19, 6, 23, 48, 0, time.UTC),
								"updatedAt":    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
								"sourceLocale": "en-us",
								"outdated":     false,
								"countryCode":  "tw",
								"url":          "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/sections/115004118448-I-need-help-with-my-account.json",
								"htmlUrl":      "https://help.honestbee.tw/hc/en-us/sections/115004118448-I-need-help-with-my-account",
								"name":         "I need help with my account",
								"description":  "",
								"locale":       "en-us",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"articlesConnection": map[string]interface{}{
									"page":      1,
									"perPage":   30,
									"pageCount": 1,
									"count":     5,
								},
							},
						},
					},
				},
			},
		},
		{
			description: "testing sort order desc case",
			body: map[string]interface{}{
				"query": `
				{
					allSections(countryCode: TW, locale: EN_US, sortOrder: DESC) {
						page
						perPage
						pageCount
						count
						sections {
							id
							position
							createdAt
							updatedAt
							sourceLocale
							outdated
							countryCode
							url
							htmlUrl
							name
							description
							locale
							categoryConnection {
								id
								name
								keyName
							}
							articlesConnection {
  								page
								perPage
								pageCount
								count
							}
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"data": map[string]interface{}{
					"allSections": map[string]interface{}{
						"page":      1,
						"perPage":   30,
						"pageCount": 1,
						"count":     1,
						"sections": []interface{}{
							map[string]interface{}{
								"id":           "115004118448",
								"position":     0,
								"createdAt":    time.Date(2017, 12, 19, 6, 23, 48, 0, time.UTC),
								"updatedAt":    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
								"sourceLocale": "en-us",
								"outdated":     false,
								"countryCode":  "tw",
								"url":          "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/sections/115004118448-I-need-help-with-my-account.json",
								"htmlUrl":      "https://help.honestbee.tw/hc/en-us/sections/115004118448-I-need-help-with-my-account",
								"name":         "I need help with my account",
								"description":  "",
								"locale":       "en-us",
								"categoryConnection": map[string]interface{}{
									"id":      "115002432448",
									"name":    "My Account",
									"keyName": "myAccount",
								},
								"articlesConnection": map[string]interface{}{
									"page":      1,
									"perPage":   30,
									"pageCount": 1,
									"count":     5,
								},
							},
						},
					},
				},
			},
		},
		{
			description: "testing not exist country code case",
			body: map[string]interface{}{
				"query": `
				{
				  allSections(countryCode: not_exist_country_code) {
						page
						perPage
						pageCount
						count
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"errors": []interface{}{
					map[string]interface{}{
						"message": "Argument \"countryCode\" has invalid value not_exist_country_code.\nExpected type \"CountryCode\", found not_exist_country_code.",
						"locations": []interface{}{
							map[string]interface{}{
								"line":   3,
								"column": 32,
							},
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
					allSections(locale: not_exist_locale) {
						page
						perPage
						pageCount
						count
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"errors": []interface{}{
					map[string]interface{}{
						"message": "Argument \"locale\" has invalid value not_exist_locale.\nExpected type \"Locale\", found not_exist_locale.",
						"locations": []interface{}{
							map[string]interface{}{
								"line":   3,
								"column": 26,
							},
						},
					},
				},
			},
		},
		{
			description: "testing sort by not correct case",
			body: map[string]interface{}{
				"query": `
				{
					allSections(sortBy: unknown_sort_by) {
						page
						perPage
						pageCount
						count
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"errors": []interface{}{
					map[string]interface{}{
						"message": "Argument \"sortBy\" has invalid value unknown_sort_by.\nExpected type \"SortBy\", found unknown_sort_by.",
						"locations": []interface{}{
							map[string]interface{}{
								"line":   3,
								"column": 26,
							},
						},
					},
				},
			},
		},
		{
			description: "testing sort order not correct case",
			body: map[string]interface{}{
				"query": `
				{
					allSections(sortOrder: unknown_sort_order) {
						page
						perPage
						pageCount
						count
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"errors": []interface{}{
					map[string]interface{}{
						"message": "Argument \"sortOrder\" has invalid value unknown_sort_order.\nExpected type \"SortOrder\", found unknown_sort_order.",
						"locations": []interface{}{
							map[string]interface{}{
								"line":   3,
								"column": 29,
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

func TestHandlersGraphQLQuerySection(t *testing.T) {
	ts := newTserver()
	defer ts.closeAll()
	testCases := []struct {
		description string
		body        map[string]interface{}
		query       string
		expectBody  map[string]interface{}
	}{
		{
			description: "testing normal query section id and TW + EN_US case",
			body: map[string]interface{}{
				"query": `
				{
					oneSection(sectionId: "115004118448", countryCode: TW) {
						id
						position
						createdAt
						updatedAt
						sourceLocale
						outdated
						countryCode
						url
						htmlUrl
						name
						description
						locale
						categoryConnection {
							id
							name
							keyName
						}
						articlesConnection {
							page
							perPage
							pageCount
							count
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"data": map[string]interface{}{
					"oneSection": map[string]interface{}{
						"id":           "115004118448",
						"position":     0,
						"createdAt":    time.Date(2017, 12, 19, 6, 23, 48, 0, time.UTC),
						"updatedAt":    time.Date(2018, 3, 6, 12, 39, 30, 0, time.UTC),
						"sourceLocale": "en-us",
						"outdated":     false,
						"countryCode":  "tw",
						"url":          "https://honestbeehelp-tw.zendesk.com/api/v2/help_center/en-us/sections/115004118448-I-need-help-with-my-account.json",
						"htmlUrl":      "https://help.honestbee.tw/hc/en-us/sections/115004118448-I-need-help-with-my-account",
						"name":         "I need help with my account",
						"description":  "",
						"locale":       "en-us",
						"categoryConnection": map[string]interface{}{
							"id":      "115002432448",
							"name":    "My Account",
							"keyName": "myAccount",
						},
						"articlesConnection": map[string]interface{}{
							"page":      1,
							"perPage":   30,
							"pageCount": 1,
							"count":     5,
						},
					},
				},
			},
		},
		{
			description: "testing normal query section id and SG + EN_US not found case",
			body: map[string]interface{}{
				"query": `
				{
					oneSection(sectionId: "115004118448") {
						id
						position
						createdAt
						updatedAt
						sourceLocale
						outdated
						countryCode
						url
						htmlUrl
						name
						description
						locale
						categoryConnection {
							id
							name
							keyName
						}
						articlesConnection {
							page
							perPage
							pageCount
							count
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"data": map[string]interface{}{
					"oneSection": nil,
				},
				"errors": []interface{}{
					map[string]interface{}{
						"message": "Record Not Found",
						"path": []interface{}{
							"oneSection",
						},
					},
				},
			},
		},
		{
			description: "testing not exist country code case",
			body: map[string]interface{}{
				"query": `
				{
					oneSection(sectionId: "sectionId", countryCode: not_exist_country_code) {
						id
						position
						createdAt
						updatedAt
						sourceLocale
						outdated
						countryCode
						url
						htmlUrl
						name
						description
						locale
						categoryConnection {
							id
							name
							keyName
						}
						articlesConnection {
							page
							perPage
							pageCount
							count
						}
					}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"errors": []interface{}{
					map[string]interface{}{
						"message": "Argument \"countryCode\" has invalid value not_exist_country_code.\nExpected type \"CountryCode\", found not_exist_country_code.",
						"locations": []interface{}{
							map[string]interface{}{
								"line":   3,
								"column": 54,
							},
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
					oneSection(sectionId: "sectionId", locale: not_exist_locale) {
						id
						position
						createdAt
						updatedAt
						sourceLocale
						outdated
						countryCode
						url
						htmlUrl
						name
						description
						locale
						categoryConnection {
							id
							name
							keyName
						}
						articlesConnection {
							page
							perPage
							pageCount
							count
						}
				 	}
				}
				`,
			},
			expectBody: map[string]interface{}{
				"errors": []interface{}{
					map[string]interface{}{
						"message": "Argument \"locale\" has invalid value not_exist_locale.\nExpected type \"Locale\", found not_exist_locale.",
						"locations": []interface{}{
							map[string]interface{}{
								"line":   3,
								"column": 49,
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
