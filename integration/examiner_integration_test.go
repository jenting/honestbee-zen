// +build integration

package integration

import (
	"context"
	"net/http"
	"testing"

	"github.com/pkg/errors"

	"github.com/honestbee/Zen/models"
)

const (
	examinerArticlesTestTimes    = 5
	examinerSectionsTestTimes    = 5
	examinerCategoriesTestTimes  = 5
	examinerTicketFormsTestTimes = 5
)

func TestExaminerTicketForms(t *testing.T) {
	cleaner := func(ts *tserver) error {
		if err := ts.service.SyncWithTicketForms(context.Background(), make([]*models.SyncTicketForm, 0)); err != nil {
			return err
		}
		if err := ts.service.SyncWithTicketFields(context.Background(), make([]*models.SyncTicketField, 0)); err != nil {
			return err
		}
		if err := ts.service.SyncWithDynamicContentItems(context.Background(), make([]*models.SyncDynamicContentItem, 0)); err != nil {
			return err
		}
		return nil
	}
	reacher := func(ts *tserver) error {
		addr := ts.URL + "/api/ticket_forms/951408"
		for i := 0; i <= ticketFormsRefreshLimit; i++ {
			resp, err := ts.Client().Get(addr)
			if err != nil {
				return err
			}
			resp.Body.Close()
			if resp.StatusCode == http.StatusNotFound {
				continue
			}
			if resp.StatusCode != http.StatusOK {
				return errors.Errorf("http resp status code expect ok, actual:%v", resp.Status)
			}
		}
		return nil
	}
	checker := func(ts *tserver) error {
		if _, err := ts.service.GetTicketForm(context.Background(), 951408, "en-us"); err != nil {
			return errors.Errorf("GetTicketForm error expect nil, actual:%v", err)
		}

		if _, err := ts.service.GetTicketFieldByFieldID(context.Background(), 24681498, "en-us"); err != nil {
			return errors.Errorf("GetTicketFields error expect nil, actual:%v", err)
		}

		if _, err := ts.service.GetDynamicContentItem(context.Background(), "{{dc.form_order_number_field}}", "en-us"); err != nil {
			return errors.Errorf("GetDynamicContentItem error expect nil, actual:%v", err)
		}
		return nil
	}

	for i := 0; i < examinerTicketFormsTestTimes; i++ {
		testExaminerHelper(t, cleaner, reacher, checker)
	}
}

func TestExaminerArticles(t *testing.T) {
	cleaner := func(ts *tserver) error {
		return ts.service.SyncWithArticles(context.Background(), make([]*models.Article, 0), "tw", "en-us")
	}
	reacher := func(ts *tserver) error {
		addr := ts.URL + "/api/sections/115004118448/articles?country_code=tw&locale=en-us"
		for i := 0; i <= articlesRefreshLimit; i++ {
			resp, err := ts.Client().Get(addr)
			if err != nil {
				return err
			}
			resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				return errors.Errorf("http resp status code expect ok, actual:%v", resp.Status)
			}
		}
		return nil
	}
	checker := func(ts *tserver) error {
		articles, count, err := ts.service.GetArticlesBySectionID(context.Background(),
			&models.GetArticlesParams{
				SectionID:   115004118448,
				Locale:      "en-us",
				CountryCode: "tw",
				PerPage:     100,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			})
		if err != nil {
			return err
		}

		if len(articles) == 0 || count == 0 {
			return errors.Errorf(
				"GetArticles expect len != 0, actual len(articles)=%v, count=%v",
				len(articles),
				count,
			)
		}
		return nil
	}

	for i := 0; i < examinerArticlesTestTimes; i++ {
		testExaminerHelper(t, cleaner, reacher, checker)
	}
}

func TestExaminerSections(t *testing.T) {
	cleaner := func(ts *tserver) error {
		return ts.service.SyncWithSections(context.Background(), make([]*models.Section, 0), "tw", "en-us")
	}
	reacher := func(ts *tserver) error {
		addr := ts.URL + "/api/categories/115002432448/sections?country_code=tw&locale=en-us"
		for i := 0; i <= sectionsRefreshLimit; i++ {
			resp, err := ts.Client().Get(addr)
			if err != nil {
				return err
			}
			resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				return errors.Errorf("http resp status code expect ok, actual:%v", resp.Status)
			}
		}
		return nil
	}
	checker := func(ts *tserver) error {
		sections, count, err := ts.service.GetSectionsByCategoryID(context.Background(),
			&models.GetSectionsParams{
				CategoryID:  115002432448,
				Locale:      "en-us",
				CountryCode: "tw",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			})
		if err != nil {
			return err
		}

		if len(sections) == 0 || count == 0 {
			return errors.Errorf(
				"GetSections expect len != 0, actual len(sections)=%v, count=%v",
				len(sections),
				count,
			)
		}
		return nil
	}

	for i := 0; i < examinerSectionsTestTimes; i++ {
		testExaminerHelper(t, cleaner, reacher, checker)
	}
}

func TestExaminerCategories(t *testing.T) {
	cleaner := func(ts *tserver) error {
		return ts.service.SyncWithCategories(context.Background(), make([]*models.Category, 0), "tw", "en-us")
	}
	reacher := func(ts *tserver) error {
		addr := ts.URL + "/api/categories?country_code=tw&locale=en-us"
		for i := 0; i <= categoriesRefreshLimit; i++ {
			resp, err := ts.Client().Get(addr)
			if err != nil {
				return err
			}
			resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				return errors.Errorf("http resp status code expect ok, actual:%v", resp.Status)
			}
		}
		return nil
	}
	checker := func(ts *tserver) error {
		categories, count, err := ts.service.GetCategories(context.Background(),
			&models.GetCategoriesParams{
				Locale:      "en-us",
				CountryCode: "tw",
				PerPage:     30,
				Page:        0,
				SortBy:      "position",
				SortOrder:   "asc",
			})
		if err != nil {
			return err
		}

		if len(categories) == 0 || count == 0 {
			return errors.Errorf(
				"GetCategories expect len != 0, actual len(categories)=%v, count=%v",
				len(categories),
				count,
			)
		}
		return nil
	}

	for i := 0; i < examinerCategoriesTestTimes; i++ {
		testExaminerHelper(t, cleaner, reacher, checker)
	}
}

func testExaminerHelper(t *testing.T, cleaner, reacher, checker func(ts *tserver) error) {
	ts := newTserver()
	defer ts.Server.Close()
	defer ts.service.Close()
	defer ts.resetAllCounter()

	// cleaning fake data from database
	if err := cleaner(ts); err != nil {
		t.Fatalf("cleaner failed:%v", err)
	}

	// reach the examiner update limit
	if err := reacher(ts); err != nil {
		t.Fatalf("reacher failed:%v", err)
	}

	// wait all examiner goroutine to finish
	ts.exam.Close()

	// checkout is the data sync with remote API
	if err := checker(ts); err != nil {
		t.Errorf("checker failed:%v", err)
	}

	// rest back to fake database
	if err := resetDB(); err != nil {
		t.Fatalf("reset db command run failed:%v", err)
	}
}
