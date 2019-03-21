// +build integration

package zendesk_test

import (
	"context"
	"log"
	"testing"

	"github.com/honestbee/Zen/config"
	"github.com/honestbee/Zen/zendesk"
)

var (
	zend *zendesk.ZenDesk
)

func init() {
	var err error
	conf := &config.Config{
		ZenDesk: &config.ZenDesk{
			RequestTimeoutSec: 60,
			AuthToken:         "ZWRtdW5kLmthb0Bob25lc3RiZWUuY29tL3Rva2VuOmZXdmVMYXVvN0lzQVExQURrbE54ZFVySkIwMWN1aFltTnhVRmVIbE8=",
			HKBaseURL:         "https://honestbeehelp-hk.zendesk.com",
			IDBaseURL:         "https://honestbee-idn.zendesk.com",
			JPBaseURL:         "https://honestbeehelp-jp.zendesk.com",
			MYBaseURL:         "https://honestbee-my.zendesk.com",
			PHBaseURL:         "https://honestbee-ph.zendesk.com",
			SGBaseURL:         "https://honestbeehelp-sg.zendesk.com",
			THBaseURL:         "https://honestbee-th.zendesk.com",
			TWBaseURL:         "https://honestbeehelp-tw.zendesk.com",
		},
	}
	zend, err = zendesk.NewZenDesk(conf)
	if err != nil {
		log.Fatalf("create testing zendesk failed:%v", err)
	}
}

func TestCreateRequest(t *testing.T) {
	testCases := []struct {
		description string
		countryCode string
		data        interface{}
		expectErr   bool
	}{
		{
			description: "testing tw normal case",
			countryCode: "tw",
			data: map[string]interface{}{
				"request": map[string]interface{}{
					"requester": map[string]interface{}{
						"name":  "zen project tester",
						"email": "zen.project.tester@honestbee.com",
					},
					"subject": "testing, please ignore",
					"comment": map[string]interface{}{
						"body": "testing, please ignore!!!",
					},
				},
			},
			expectErr: false,
		},
		{
			description: "testing no country code case",
			countryCode: "not-exist-country-code",
			data: map[string]interface{}{
				"request": map[string]interface{}{
					"requester": map[string]interface{}{
						"name":  "zen project tester",
						"email": "zen.project.tester@honestbee.com",
					},
					"subject": "testing, please ignore",
					"comment": map[string]interface{}{
						"body": "testing, please ignore!!!",
					},
				},
			},
			expectErr: true,
		},
		{
			description: "testing empty data case",
			countryCode: "tw",
			data:        map[string]interface{}{},
			expectErr:   true,
		},
		{
			description: "testing nil data case",
			countryCode: "tw",
			data:        nil,
			expectErr:   true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			err := zend.CreateRequest(context.Background(), tt.countryCode, tt.data)
			if tt.expectErr && err == nil {
				t.Errorf("[%s] expect an error, actual nil", tt.description)
			} else if !tt.expectErr && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			}
		})
	}
}

func TestListTicketForms(t *testing.T) {
	testCases := []struct {
		description string
		expectErr   bool
	}{
		{
			description: "testing normal case",
			expectErr:   false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			_, err := zend.ListTicketForms(context.Background())
			if tt.expectErr && err == nil {
				t.Errorf("[%s] expect an error, actual nil", tt.description)
			} else if !tt.expectErr && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			}
		})
	}
}

func TestListTicketFields(t *testing.T) {
	testCases := []struct {
		description string
		expectErr   bool
	}{
		{
			description: "testing normal case",
			expectErr:   false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			_, err := zend.ListTicketFields(context.Background())
			if tt.expectErr && err == nil {
				t.Errorf("[%s] expect an error, actual nil", tt.description)
			} else if !tt.expectErr && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			}
		})
	}
}

func TestListDynamicContentItems(t *testing.T) {
	testCases := []struct {
		description string
		expectErr   bool
	}{
		{
			description: "testing normal case",
			expectErr:   false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			_, err := zend.ListDynamicContentItems(context.Background())
			if tt.expectErr && err == nil {
				t.Errorf("[%s] expect an error, actual nil", tt.description)
			} else if !tt.expectErr && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			}
		})
	}
}

func TestListCategories(t *testing.T) {
	testCases := []struct {
		description string
		countryCode string
		locale      string
		expectErr   bool
	}{
		{
			description: "testing sg en-us case",
			countryCode: "sg",
			locale:      "en-us",
			expectErr:   false,
		},
		{
			description: "testing tw en-us case",
			countryCode: "tw",
			locale:      "en-us",
			expectErr:   false,
		},
		{
			description: "testing hk en-us case",
			countryCode: "hk",
			locale:      "en-us",
			expectErr:   false,
		},
		{
			description: "testing jp en-us case",
			countryCode: "jp",
			locale:      "en-us",
			expectErr:   false,
		},
		{
			description: "testing th en-us case",
			countryCode: "th",
			locale:      "en-us",
			expectErr:   false,
		},
		{
			description: "testing my en-us case",
			countryCode: "my",
			locale:      "en-us",
			expectErr:   false,
		},
		{
			description: "testing id en-us case",
			countryCode: "id",
			locale:      "en-us",
			expectErr:   false,
		},
		{
			description: "testing ph en-us case",
			countryCode: "ph",
			locale:      "en-us",
			expectErr:   false,
		},
		{
			description: "testing tw zh-tw case",
			countryCode: "tw",
			locale:      "zh-tw",
			expectErr:   false,
		},
		{
			description: "testing hk zh-tw case",
			countryCode: "hk",
			locale:      "zh-tw",
			expectErr:   false,
		},
		{
			description: "testing jp ja case",
			countryCode: "jp",
			locale:      "ja",
			expectErr:   false,
		},
		{
			description: "testing th th case",
			countryCode: "th",
			locale:      "th",
			expectErr:   false,
		},
		{
			description: "testing id id case",
			countryCode: "id",
			locale:      "id",
			expectErr:   false,
		},
		{
			description: "testing sg wrong-locale case",
			countryCode: "sg",
			locale:      "wrong-locale",
			expectErr:   true,
		},
		{
			description: "testing tw wrong-locale case",
			countryCode: "tw",
			locale:      "wrong-locale",
			expectErr:   true,
		},
		{
			description: "testing hk wrong-locale case",
			countryCode: "hk",
			locale:      "wrong-locale",
			expectErr:   true,
		},
		{
			description: "testing jp wrong-locale case",
			countryCode: "jp",
			locale:      "wrong-locale",
			expectErr:   true,
		},
		{
			description: "testing th wrong-locale case",
			countryCode: "th",
			locale:      "wrong-locale",
			expectErr:   true,
		},
		{
			description: "testing my wrong-locale case",
			countryCode: "my",
			locale:      "wrong-locale",
			expectErr:   true,
		},
		{
			description: "testing id wrong-locale case",
			countryCode: "id",
			locale:      "wrong-locale",
			expectErr:   true,
		},
		{
			description: "testing ph wrong-locale case",
			countryCode: "ph",
			locale:      "wrong-locale",
			expectErr:   true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			_, err := zend.ListCategories(context.Background(), tt.countryCode, tt.locale)
			if tt.expectErr && err == nil {
				t.Errorf("[%s] expect an error, actual nil", tt.description)
			} else if !tt.expectErr && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			}
		})
	}
}

func TestListSections(t *testing.T) {
	testCases := []struct {
		description string
		countryCode string
		locale      string
		expectErr   bool
	}{
		{
			description: "testing sg en-us case",
			countryCode: "sg",
			locale:      "en-us",
			expectErr:   false,
		},
		{
			description: "testing tw en-us case",
			countryCode: "tw",
			locale:      "en-us",
			expectErr:   false,
		},
		{
			description: "testing hk en-us case",
			countryCode: "hk",
			locale:      "en-us",
			expectErr:   false,
		},
		{
			description: "testing jp en-us case",
			countryCode: "jp",
			locale:      "en-us",
			expectErr:   false,
		},
		{
			description: "testing th en-us case",
			countryCode: "th",
			locale:      "en-us",
			expectErr:   false,
		},
		{
			description: "testing my en-us case",
			countryCode: "my",
			locale:      "en-us",
			expectErr:   false,
		},
		{
			description: "testing id en-us case",
			countryCode: "id",
			locale:      "en-us",
			expectErr:   false,
		},
		{
			description: "testing ph en-us case",
			countryCode: "ph",
			locale:      "en-us",
			expectErr:   false,
		},
		{
			description: "testing tw zh-tw case",
			countryCode: "tw",
			locale:      "zh-tw",
			expectErr:   false,
		},
		{
			description: "testing hk zh-tw case",
			countryCode: "hk",
			locale:      "zh-tw",
			expectErr:   false,
		},
		{
			description: "testing jp ja case",
			countryCode: "jp",
			locale:      "ja",
			expectErr:   false,
		},
		{
			description: "testing th th case",
			countryCode: "th",
			locale:      "th",
			expectErr:   false,
		},
		{
			description: "testing id id case",
			countryCode: "id",
			locale:      "id",
			expectErr:   false,
		},
		{
			description: "testing sg wrong-locale case",
			countryCode: "sg",
			locale:      "wrong-locale",
			expectErr:   true,
		},
		{
			description: "testing tw wrong-locale case",
			countryCode: "tw",
			locale:      "wrong-locale",
			expectErr:   true,
		},
		{
			description: "testing hk wrong-locale case",
			countryCode: "hk",
			locale:      "wrong-locale",
			expectErr:   true,
		},
		{
			description: "testing jp wrong-locale case",
			countryCode: "jp",
			locale:      "wrong-locale",
			expectErr:   true,
		},
		{
			description: "testing th wrong-locale case",
			countryCode: "th",
			locale:      "wrong-locale",
			expectErr:   true,
		},
		{
			description: "testing my wrong-locale case",
			countryCode: "my",
			locale:      "wrong-locale",
			expectErr:   true,
		},
		{
			description: "testing id wrong-locale case",
			countryCode: "id",
			locale:      "wrong-locale",
			expectErr:   true,
		},
		{
			description: "testing ph wrong-locale case",
			countryCode: "ph",
			locale:      "wrong-locale",
			expectErr:   true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			_, err := zend.ListSections(context.Background(), tt.countryCode, tt.locale)
			if tt.expectErr && err == nil {
				t.Errorf("[%s] expect an error, actual nil", tt.description)
			} else if !tt.expectErr && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			}
		})
	}
}

func TestListArticles(t *testing.T) {
	testCases := []struct {
		description string
		countryCode string
		locale      string
		expectErr   bool
	}{
		{
			description: "testing sg en-us case",
			countryCode: "sg",
			locale:      "en-us",
			expectErr:   false,
		},
		{
			description: "testing tw en-us case",
			countryCode: "tw",
			locale:      "en-us",
			expectErr:   false,
		},
		{
			description: "testing hk en-us case",
			countryCode: "hk",
			locale:      "en-us",
			expectErr:   false,
		},
		{
			description: "testing jp en-us case",
			countryCode: "jp",
			locale:      "en-us",
			expectErr:   false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			_, err := zend.ListArticles(context.Background(), tt.countryCode, tt.locale)
			if tt.expectErr && err == nil {
				t.Errorf("[%s] expect an error, actual nil", tt.description)
			} else if !tt.expectErr && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			}
		})
	}
}

func TestShowArticle(t *testing.T) {
	testCases := []struct {
		description string
		id          int
		countryCode string
		locale      string
		expectErr   bool
	}{
		{
			description: "testing sg en-us case",
			countryCode: "sg",
			id:          360000766988,
			locale:      "en-us",
			expectErr:   false,
		},
		{
			description: "testing tw en-us case",
			countryCode: "tw",
			id:          115015959188,
			locale:      "en-us",
			expectErr:   false,
		},
		{
			description: "testing hk en-us case",
			countryCode: "hk",
			locale:      "en-us",
			id:          360000777288,
			expectErr:   false,
		},
		{
			description: "testing jp en-us case",
			countryCode: "jp",
			locale:      "en-us",
			id:          115015927467,
			expectErr:   false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			_, err := zend.ShowArticle(context.Background(), tt.id, tt.countryCode, tt.locale)
			if tt.expectErr && err == nil {
				t.Errorf("[%s] expect an error, actual nil", tt.description)
			} else if !tt.expectErr && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			}
		})
	}
}

func TestInstantSearch(t *testing.T) {
	testCases := []struct {
		description string
		query       string
		countryCode string
		locale      string
		expectErr   bool
	}{
		{
			description: "testing sg en-us case",
			countryCode: "sg",
			query:       "order",
			locale:      "en-us",
			expectErr:   false,
		},
		{
			description: "testing tw zh-tw case",
			countryCode: "tw",
			query:       "訂單",
			locale:      "zh-tw",
			expectErr:   false,
		},
		{
			description: "testing tw en-us case",
			countryCode: "tw",
			query:       "order",
			locale:      "zn-tw",
			expectErr:   false,
		},
		{
			description: "testing hk en-us case",
			countryCode: "hk",
			query:       "order",
			locale:      "en-us",
			expectErr:   false,
		},
		{
			description: "testing jp en-us case",
			countryCode: "jp",
			query:       "order",
			locale:      "en-us",
			expectErr:   false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			_, err := zend.InstantSearch(context.Background(), tt.query, tt.countryCode, tt.locale)
			if tt.expectErr && err == nil {
				t.Errorf("[%s] expect an error, actual nil", tt.description)
			} else if !tt.expectErr && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			}
		})
	}
}

func TestZenDeskSearch(t *testing.T) {
	testCases := []struct {
		description string
		query       string
		countryCode string
		locale      string
		perPage     int
		page        int
		SortOrder   string
		expectErr   bool
	}{
		{
			description: "testing sg en-us case",
			countryCode: "sg",
			query:       "order",
			locale:      "en-us",
			perPage:     30,
			page:        1,
			SortOrder:   "ASC",
			expectErr:   false,
		},
		{
			description: "testing tw zh-tw case",
			countryCode: "tw",
			query:       "訂單",
			locale:      "zh-tw",
			perPage:     30,
			page:        1,
			SortOrder:   "DESC",
			expectErr:   false,
		},
		{
			description: "testing tw en-us case",
			countryCode: "tw",
			query:       "order",
			locale:      "zn-tw",
			perPage:     30,
			page:        1,
			SortOrder:   "ASC",
			expectErr:   false,
		},
		{
			description: "testing hk en-us case",
			countryCode: "hk",
			query:       "order",
			locale:      "en-us",
			perPage:     30,
			page:        1,
			SortOrder:   "DESC",
			expectErr:   false,
		},
		{
			description: "testing jp en-us case",
			countryCode: "jp",
			query:       "order",
			locale:      "en-us",
			perPage:     30,
			page:        1,
			SortOrder:   "ASC",
			expectErr:   false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			_, err := zend.Search(context.Background(), nil, tt.query, tt.countryCode, tt.locale, &zendesk.Pagination{
				PerPage:   tt.perPage,
				Page:      tt.page,
				SortOrder: tt.SortOrder,
			})
			if tt.expectErr && err == nil {
				t.Errorf("[%s] expect an error, actual nil", tt.description)
			} else if !tt.expectErr && err != nil {
				t.Errorf("[%s] expect no error, actual:%v", tt.description, err)
			}
		})
	}
}
