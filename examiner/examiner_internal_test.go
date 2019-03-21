package examiner

import (
	"context"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/go-test/deep"
	"github.com/rs/zerolog"

	"github.com/honestbee/Zen/config"
	"github.com/honestbee/Zen/models"
	"github.com/honestbee/Zen/zendesk"
)

var (
	logger  = zerolog.New(ioutil.Discard)
	zend, _ = zendesk.NewZenDesk(&config.Config{
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
	})
)

func TestCheckTicketForms(t *testing.T) {
	mockServ := models.NewMockService()
	exam, _ := NewExaminer(&config.Config{
		Examiner: &config.Examiner{
			MaxPoolSize:             10,
			MaxWorkerSize:           5,
			CategoriesRefreshLimit:  1,
			SectionsRefreshLimit:    1,
			ArticlesRefreshLimit:    1,
			TicketFormsRefreshLimit: 1,
		},
	}, &logger, mockServ, zend)
	defer exam.Close()

	testCases := []struct {
		description    string
		expectSequence map[string]bool
	}{
		{
			description: "testing normal case",
			expectSequence: map[string]bool{
				"PlusOneTicketFormsCounter":                   true,
				"LockTicketFormsCounter":                      true,
				"SyncWithTicketForms":                         true,
				"TicketFormCacheInvalidate":                   true,
				"SyncWithTicketFields":                        true,
				"TicketFieldCacheInvalidate":                  true,
				"TicketFieldCustomFieldOptionCacheInvalidate": true,
				"TicketFieldSystemFieldOptionCacheInvalidate": true,
				"SyncWithDynamicContentItems":                 true,
				"ResetTicketFormsCounter":                     true,
				"UnlockTicketFormsCounter":                    true,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			exam.ticketFormsWork(context.Background())
			actualSequence := mockServ.Sequence

			if diff := deep.Equal(tt.expectSequence, actualSequence); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
			mockServ.ResetSequence()
		})
	}
}

func TestCheckCategories(t *testing.T) {
	mockServ := models.NewMockService()
	exam, _ := NewExaminer(&config.Config{
		Examiner: &config.Examiner{
			MaxPoolSize:            10,
			MaxWorkerSize:          5,
			CategoriesRefreshLimit: 1,
			SectionsRefreshLimit:   1,
			ArticlesRefreshLimit:   1,
		},
	}, &logger, mockServ, zend)
	defer exam.Close()

	testCases := []struct {
		description    string
		countryCode    string
		locale         string
		expectSequence map[string]bool
	}{
		{
			description: "testing normal case",
			countryCode: "tw",
			locale:      "en-us",
			expectSequence: map[string]bool{
				"PlusOneCategoriesCounter":  true,
				"LockCategoriesCounter":     true,
				"SyncWithCategories":        true,
				"CategoriesCacheInvalidate": true,
				"ResetCategoriesCounter":    true,
				"UnlockCategoriesCounter":   true,
			},
		},
		{
			description: "testing PlusOneCategoriesCounter failed case",
			countryCode: models.PlusCounterReturnErrorCountryCode,
			locale:      "en-us",
			expectSequence: map[string]bool{
				"PlusOneCategoriesCounter": true,
			},
		},
		{
			description: "testing count < e.categoriesRefreshLimit case",
			countryCode: models.PlusCounterReturnSmallerCountryCode,
			locale:      "en-us",
			expectSequence: map[string]bool{
				"PlusOneCategoriesCounter": true,
			},
		},
		{
			description: "testing LockCategoriesCounter error failed case",
			countryCode: models.LockCounterReturnErrorCountryCode,
			locale:      "en-us",
			expectSequence: map[string]bool{
				"PlusOneCategoriesCounter": true,
				"LockCategoriesCounter":    true,
			},
		},
		{
			description: "testing LockCategoriesCounter lock failed case",
			countryCode: models.LockCounterFailedCountryCode,
			locale:      "en-us",
			expectSequence: map[string]bool{
				"PlusOneCategoriesCounter": true,
				"LockCategoriesCounter":    true,
			},
		},
		{
			description: "testing SyncWithCategories failed case",
			countryCode: "jp",
			locale:      models.SyncDBFailedLocale,
			expectSequence: map[string]bool{
				"PlusOneCategoriesCounter": true,
				"LockCategoriesCounter":    true,
				"SyncWithCategories":       true,
			},
		},
		{
			description: "testing ResetCategoriesCounter failed case",
			countryCode: models.ResetCounterFailedCountryCode,
			locale:      "en-us",
			expectSequence: map[string]bool{
				"PlusOneCategoriesCounter":  true,
				"LockCategoriesCounter":     true,
				"SyncWithCategories":        true,
				"CategoriesCacheInvalidate": true,
				"ResetCategoriesCounter":    true,
			},
		},
		{
			description: "testing UnlockCategoriesCounter failed case",
			countryCode: models.UnlockCounterFailedCountryCode,
			locale:      "en-us",
			expectSequence: map[string]bool{
				"PlusOneCategoriesCounter":  true,
				"LockCategoriesCounter":     true,
				"SyncWithCategories":        true,
				"CategoriesCacheInvalidate": true,
				"ResetCategoriesCounter":    true,
				"UnlockCategoriesCounter":   true,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			exam.categoriesWork(context.Background(), tt.countryCode, tt.locale)
			actualSequence := mockServ.Sequence

			if diff := deep.Equal(tt.expectSequence, actualSequence); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
			mockServ.ResetSequence()
		})
	}
}

func TestCheckSections(t *testing.T) {
	mockServ := models.NewMockService()
	exam, _ := NewExaminer(&config.Config{
		Examiner: &config.Examiner{
			MaxPoolSize:            10,
			MaxWorkerSize:          5,
			CategoriesRefreshLimit: 1,
			SectionsRefreshLimit:   1,
			ArticlesRefreshLimit:   1,
		},
	}, &logger, mockServ, zend)
	defer exam.Close()

	testCases := []struct {
		description    string
		countryCode    string
		locale         string
		expectSequence map[string]bool
	}{
		{
			description: "testing normal case",
			countryCode: "tw",
			locale:      "en-us",
			expectSequence: map[string]bool{
				"PlusOneSectionsCounter":  true,
				"LockSectionsCounter":     true,
				"SyncWithSections":        true,
				"SectionsCacheInvalidate": true,
				"ResetSectionsCounter":    true,
				"UnlockSectionsCounter":   true,
			},
		},
		{
			description: "testing PlusOneSectionsCounter failed case",
			countryCode: models.PlusCounterReturnErrorCountryCode,
			locale:      "en-us",
			expectSequence: map[string]bool{
				"PlusOneSectionsCounter": true,
			},
		},
		{
			description: "testing count < e.categoriesRefreshLimit case",
			countryCode: models.PlusCounterReturnSmallerCountryCode,
			locale:      "en-us",
			expectSequence: map[string]bool{
				"PlusOneSectionsCounter": true,
			},
		},
		{
			description: "testing LockSectionsCounter error failed case",
			countryCode: models.LockCounterReturnErrorCountryCode,
			locale:      "en-us",
			expectSequence: map[string]bool{
				"PlusOneSectionsCounter": true,
				"LockSectionsCounter":    true,
			},
		},
		{
			description: "testing LockSectionsCounter lock failed case",
			countryCode: models.LockCounterFailedCountryCode,
			locale:      "en-us",
			expectSequence: map[string]bool{
				"PlusOneSectionsCounter": true,
				"LockSectionsCounter":    true,
			},
		},
		{
			description: "testing SyncWithSections failed case",
			countryCode: "jp",
			locale:      models.SyncDBFailedLocale,
			expectSequence: map[string]bool{
				"PlusOneSectionsCounter": true,
				"LockSectionsCounter":    true,
				"SyncWithSections":       true,
			},
		},
		{
			description: "testing ResetSectionsCounter failed case",
			countryCode: models.ResetCounterFailedCountryCode,
			locale:      "en-us",
			expectSequence: map[string]bool{
				"PlusOneSectionsCounter":  true,
				"LockSectionsCounter":     true,
				"SyncWithSections":        true,
				"SectionsCacheInvalidate": true,
				"ResetSectionsCounter":    true,
			},
		},
		{
			description: "testing UnlockSectionsCounter failed case",
			countryCode: models.UnlockCounterFailedCountryCode,
			locale:      "en-us",
			expectSequence: map[string]bool{
				"PlusOneSectionsCounter":  true,
				"LockSectionsCounter":     true,
				"SyncWithSections":        true,
				"SectionsCacheInvalidate": true,
				"ResetSectionsCounter":    true,
				"UnlockSectionsCounter":   true,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			exam.sectionsWork(context.Background(), tt.countryCode, tt.locale)
			actualSequence := mockServ.Sequence

			if diff := deep.Equal(tt.expectSequence, actualSequence); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
			mockServ.ResetSequence()
		})
	}
}

func TestCheckArticles(t *testing.T) {
	mockServ := models.NewMockService()
	exam, _ := NewExaminer(&config.Config{
		Examiner: &config.Examiner{
			MaxPoolSize:            10,
			MaxWorkerSize:          5,
			CategoriesRefreshLimit: 1,
			SectionsRefreshLimit:   1,
			ArticlesRefreshLimit:   1,
		},
	}, &logger, mockServ, zend)
	defer exam.Close()

	testCases := []struct {
		description    string
		countryCode    string
		locale         string
		expectSequence map[string]bool
	}{
		{
			description: "testing normal case",
			countryCode: "tw",
			locale:      "en-us",
			expectSequence: map[string]bool{
				"PlusOneArticlesCounter":  true,
				"LockArticlesCounter":     true,
				"SyncWithArticles":        true,
				"ArticlesCacheInvalidate": true,
				"ResetArticlesCounter":    true,
				"UnlockArticlesCounter":   true,
			},
		},
		{
			description: "testing PlusOneArticlesCounter failed case",
			countryCode: models.PlusCounterReturnErrorCountryCode,
			locale:      "en-us",
			expectSequence: map[string]bool{
				"PlusOneArticlesCounter": true,
			},
		},
		{
			description: "testing count < e.categoriesRefreshLimit case",
			countryCode: models.PlusCounterReturnSmallerCountryCode,
			locale:      "en-us",
			expectSequence: map[string]bool{
				"PlusOneArticlesCounter": true,
			},
		},
		{
			description: "testing LockArticlesCounter error failed case",
			countryCode: models.LockCounterReturnErrorCountryCode,
			locale:      "en-us",
			expectSequence: map[string]bool{
				"PlusOneArticlesCounter": true,
				"LockArticlesCounter":    true,
			},
		},
		{
			description: "testing LockArticlesCounter lock failed case",
			countryCode: models.LockCounterFailedCountryCode,
			locale:      "en-us",
			expectSequence: map[string]bool{
				"PlusOneArticlesCounter": true,
				"LockArticlesCounter":    true,
			},
		},
		{
			description: "testing SyncWithArticles failed case",
			countryCode: "jp",
			locale:      models.SyncDBFailedLocale,
			expectSequence: map[string]bool{
				"PlusOneArticlesCounter": true,
				"LockArticlesCounter":    true,
				"SyncWithArticles":       true,
			},
		},
		{
			description: "testing ResetArticlesCounter failed case",
			countryCode: models.ResetCounterFailedCountryCode,
			locale:      "en-us",
			expectSequence: map[string]bool{
				"PlusOneArticlesCounter":  true,
				"LockArticlesCounter":     true,
				"SyncWithArticles":        true,
				"ArticlesCacheInvalidate": true,
				"ResetArticlesCounter":    true,
			},
		},
		{
			description: "testing UnlockArticlesCounter failed case",
			countryCode: models.UnlockCounterFailedCountryCode,
			locale:      "en-us",
			expectSequence: map[string]bool{
				"PlusOneArticlesCounter":  true,
				"LockArticlesCounter":     true,
				"SyncWithArticles":        true,
				"ArticlesCacheInvalidate": true,
				"ResetArticlesCounter":    true,
				"UnlockArticlesCounter":   true,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			exam.articlesWork(context.Background(), tt.countryCode, tt.locale)
			actualSequence := mockServ.Sequence

			if diff := deep.Equal(tt.expectSequence, actualSequence); diff != nil {
				t.Errorf("[%s] %v", tt.description, diff)
			}
			mockServ.ResetSequence()
		})
	}
}

func TestAlwaysNotSync(t *testing.T) {
	mockServ := models.NewMockService()
	exam, _ := NewExaminer(&config.Config{
		Examiner: &config.Examiner{
			MaxPoolSize:             10,
			MaxWorkerSize:           5,
			CategoriesRefreshLimit:  0,
			SectionsRefreshLimit:    0,
			ArticlesRefreshLimit:    0,
			TicketFormsRefreshLimit: 0,
		},
	}, &logger, mockServ, zend)
	defer exam.Close()

	testCases := []struct {
		description    string
		countryCode    string
		locale         string
		expectSequence map[string]bool
	}{
		{
			description: "testing normal case",
			countryCode: "tw",
			locale:      "en-us",
			expectSequence: map[string]bool{
				"PlusOneCategoriesCounter":  true,
				"PlusOneSectionsCounter":    true,
				"PlusOneArticlesCounter":    true,
				"PlusOneTicketFormsCounter": true,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			exam.categoriesWork(context.Background(), tt.countryCode, tt.locale)
			exam.sectionsWork(context.Background(), tt.countryCode, tt.locale)
			exam.articlesWork(context.Background(), tt.countryCode, tt.locale)
			exam.ticketFormsWork(context.Background())
			actualSequence := mockServ.Sequence

			if !reflect.DeepEqual(tt.expectSequence, actualSequence) {
				t.Errorf("[%s] expect:%v != actual:%v", tt.description, tt.expectSequence, actualSequence)
			}
			mockServ.ResetSequence()
		})
	}
}
