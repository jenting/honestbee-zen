package zendesk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	"github.com/honestbee/Zen/config"
)

// ZenDesk is the instance to conmunicate with zendesk API.
type ZenDesk struct {
	token    string
	client   *http.Client
	urlTable map[string]string
}

// Pagination is the instance to present pagination.
type Pagination struct {
	PerPage   int    `json:"per_page,omitempty"`
	Page      int    `json:"page,omitempty"`
	SortOrder string `json:"sort_order,omitempty"`
}

// NewZenDesk returns a ZenDesk instance.
func NewZenDesk(conf *config.Config) (*ZenDesk, error) {
	return &ZenDesk{
		token: conf.ZenDesk.AuthToken,
		client: &http.Client{
			Timeout: time.Duration(conf.ZenDesk.RequestTimeoutSec) * time.Second,
		},
		urlTable: map[string]string{
			"hk": conf.ZenDesk.HKBaseURL,
			"id": conf.ZenDesk.IDBaseURL,
			"jp": conf.ZenDesk.JPBaseURL,
			"my": conf.ZenDesk.MYBaseURL,
			"ph": conf.ZenDesk.PHBaseURL,
			"sg": conf.ZenDesk.SGBaseURL,
			"th": conf.ZenDesk.THBaseURL,
			"tw": conf.ZenDesk.TWBaseURL,
		},
	}, nil
}

func (z *ZenDesk) connect(ctx context.Context, dest interface{}, expectStatus int, req *http.Request) error {
	span, _ := tracer.StartSpanFromContext(ctx, req.URL.Path)
	defer span.Finish()

	req.Header.Set("Cache-Control", "no-cache")
	resp, err := z.client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "zendesk: [connect] url[%s] http client do failed", req.URL.String())
	}
	defer resp.Body.Close()

	if resp.StatusCode != expectStatus {
		return errors.Errorf("zendesk: [connect] url[%s] status expect[%v], actual[%v]",
			req.RequestURI,
			expectStatus,
			resp.Status,
		)
	}

	if dest != nil {
		if err = json.NewDecoder(resp.Body).Decode(dest); err != nil {
			return errors.Wrapf(err, "zendesk: [connect] url[%s] json decode failed", req.URL.String())
		}
	}

	return nil
}

func (z *ZenDesk) authConnectPOST(ctx context.Context, dest interface{}, url string, expectStatus int, params io.Reader) error {
	req, err := http.NewRequest(http.MethodPost, url, params)
	if err != nil {
		return errors.Wrapf(err, "zendesk: [authConnectPOST] url[%s] http NewRequest failed", url)
	}
	req.Header.Set("Authorization", "Basic "+z.token)
	req.Header.Set("Content-Type", "application/json")

	return errors.Wrapf(
		z.connect(ctx, dest, expectStatus, req),
		"zendesk: [authConnectPOST] url[%s] connect failed",
		url,
	)
}

func (z *ZenDesk) connectPOST(ctx context.Context, dest interface{}, url string, expectStatus int, params io.Reader) error {
	req, err := http.NewRequest(http.MethodPost, url, params)
	if err != nil {
		return errors.Wrapf(err, "zendesk: [connectPOST] url[%s] http NewRequest failed", url)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return errors.Wrapf(
		z.connect(ctx, dest, expectStatus, req),
		"zendesk: [connectPOST] url[%s] connect failed",
		url,
	)
}

func (z *ZenDesk) authConnectGET(ctx context.Context, dest interface{}, url string, expectStatus int, params io.Reader) error {
	req, err := http.NewRequest(http.MethodGet, url, params)
	if err != nil {
		return errors.Wrapf(err, "zendesk: [authConnectGET] url[%s] http NewRequest failed", url)
	}
	req.Header.Set("Authorization", "Basic "+z.token)

	return errors.Wrapf(
		z.connect(ctx, dest, expectStatus, req),
		"zendesk: [authConnectGET] url[%s] connect failed",
		url,
	)
}

func (z *ZenDesk) connectGET(ctx context.Context, dest interface{}, url string, expectStatus int, params io.Reader) error {
	req, err := http.NewRequest(http.MethodGet, url, params)
	if err != nil {
		return errors.Wrapf(err, "zendesk: [connectGET] url[%s] http NewRequest failed", url)
	}

	return errors.Wrapf(
		z.connect(ctx, dest, expectStatus, req),
		"zendesk: [connectGET] url[%s] connect failed",
		url,
	)
}

// CreateRequest do a POST request to zendesk API to create a request.
func (z *ZenDesk) CreateRequest(ctx context.Context, countryCode string, data interface{}) error {
	url := z.identifyCountryCode(countryCode) + "/api/v2/requests.json"
	binaryData, err := json.Marshal(data)
	if err != nil {
		return errors.Wrapf(err, "zendesk: [CreateRequest] json marshal failed")
	}

	return errors.Wrapf(
		z.authConnectPOST(ctx, nil, url, http.StatusCreated, bytes.NewReader(binaryData)),
		"zendesk: [CreateRequest] connect failed",
	)
}

// ListTicketForms returns the ticket forms from zendesk API.
// *NOTE*: zendesk ticket_forms.json api does not care country code all returns the same ouput.
func (z *ZenDesk) ListTicketForms(ctx context.Context) ([]*TicketForm, error) {
	url := z.identifyCountryCode("tw") + "/api/v2/ticket_forms.json"

	ret := make([]*TicketForm, 0)
	for {
		forms := new(ListTicketForms)
		if err := z.authConnectGET(ctx, forms, url, http.StatusOK, nil); err != nil {
			return nil, errors.Wrapf(err, "zendesk: [ListTicketForms] connect failed")
		}

		ret = append(ret, forms.TicketForms...)
		if forms.NextPage == nil || *forms.NextPage == "" {
			break
		}
		url = *forms.NextPage
	}

	return ret, nil
}

// ListTicketFields returns the ticket fiels from zendesk API.
// *NOTE*: zendesk ticket_fields.json api does not care country code all returns the same ouput.
func (z *ZenDesk) ListTicketFields(ctx context.Context) ([]*TicketField, error) {
	url := z.identifyCountryCode("tw") + "/api/v2/ticket_fields.json"

	ret := make([]*TicketField, 0)
	for {
		fields := new(ListTicketFields)
		if err := z.authConnectGET(ctx, fields, url, http.StatusOK, nil); err != nil {
			return nil, errors.Wrapf(err, "zendesk: [ListTicketFields] connect failed")
		}

		ret = append(ret, fields.TicketFields...)
		if fields.NextPage == nil || *fields.NextPage == "" {
			break
		}
		url = *fields.NextPage
	}

	return ret, nil
}

// ListDynamicContentItems returns the dynamic content (DC) items from zendesk API.
// *NOTE*: zendesk dynamic_content/items.json api does not care country code all returns the same ouput.
func (z *ZenDesk) ListDynamicContentItems(ctx context.Context) ([]*DynamicContentItem, error) {
	url := z.identifyCountryCode("tw") + "/api/v2/dynamic_content/items.json"

	ret := make([]*DynamicContentItem, 0)
	for {
		items := new(ListDynamicContentItems)
		if err := z.authConnectGET(ctx, items, url, http.StatusOK, nil); err != nil {
			return nil, errors.Wrapf(err, "zendesk: [ListDynamicContentItems] connect failed")
		}

		ret = append(ret, items.Items...)
		if items.NextPage == nil || *items.NextPage == "" {
			break
		}
		url = *items.NextPage
	}

	return ret, nil
}

// ListCategories returns the categories depends on country code and locale.
func (z *ZenDesk) ListCategories(ctx context.Context, countryCode, locale string) ([]*Category, error) {
	url := fmt.Sprintf("%s/api/v2/help_center/%s/categories.json?page=1&per_page=100",
		z.identifyCountryCode(countryCode),
		locale,
	)

	ret := make([]*Category, 0)
	for {
		categories := new(ListCategories)
		if err := z.connectGET(ctx, categories, url, http.StatusOK, nil); err != nil {
			return nil, errors.Wrapf(err, "zendesk: [ListCategories] connect failed")
		}

		ret = append(ret, categories.Categories...)

		if categories.NextPage == nil || *categories.NextPage == "" {
			break
		}
		url = *categories.NextPage
		if !strings.Contains(url, "per_page") {
			url += "&per_page=100"
		}
	}

	return ret, nil
}

// ListSections returns the sections depends on country code and locale.
func (z *ZenDesk) ListSections(ctx context.Context, countryCode, locale string) ([]*Section, error) {
	url := fmt.Sprintf("%s/api/v2/help_center/%s/sections.json?page=1&per_page=100",
		z.identifyCountryCode(countryCode),
		locale,
	)

	ret := make([]*Section, 0)
	for {
		sections := new(ListSections)
		if err := z.connectGET(ctx, sections, url, http.StatusOK, nil); err != nil {
			return nil, errors.Wrapf(err, "zendesk: [ListSections] connect failed")
		}

		ret = append(ret, sections.Sections...)

		if sections.NextPage == nil || *sections.NextPage == "" {
			break
		}
		url = *sections.NextPage
		if !strings.Contains(url, "per_page") {
			url += "&per_page=100"
		}
	}

	return ret, nil
}

// ListArticles returns the articles depends on country code and locale.
func (z *ZenDesk) ListArticles(ctx context.Context, countryCode, locale string) ([]*Article, error) {
	url := fmt.Sprintf("%s/api/v2/help_center/%s/articles.json?page=1&per_page=100",
		z.identifyCountryCode(countryCode),
		locale,
	)

	ret := make([]*Article, 0)
	for {
		articles := new(ListArticles)
		if err := z.connectGET(ctx, articles, url, http.StatusOK, nil); err != nil {
			return nil, errors.Wrapf(err, "zendesk: [ListArticles] connect failed")
		}

		ret = append(ret, articles.Articles...)

		if articles.NextPage == nil || *articles.NextPage == "" {
			break
		}
		url = *articles.NextPage
		if !strings.Contains(url, "per_page") {
			url += "&per_page=100"
		}
	}

	return ret, nil
}

// ShowArticle returns the article depends on id, country code and locale.
func (z *ZenDesk) ShowArticle(ctx context.Context, id int, countryCode, locale string) (*Article, error) {
	url := fmt.Sprintf("%s/api/v2/help_center/%s/articles/%d.json",
		z.identifyCountryCode(countryCode),
		locale,
		id,
	)

	showArticle := &ShowArticle{}
	if err := z.connectGET(ctx, showArticle, url, http.StatusOK, nil); err != nil {
		return nil, errors.Wrapf(err, "zendesk: [ShowArticle] connect failed")
	}

	return showArticle.Article, nil
}

// CreateVote returns the article depends on id,expectVote, country code and locale.
func (z *ZenDesk) CreateVote(ctx context.Context, id int, expectVote, countryCode, locale string) (*Vote, error) {
	url := fmt.Sprintf("%s/hc/%s/articles/%d/vote",
		z.identifyCountryCode(countryCode),
		locale,
		id,
	)

	params := fmt.Sprintf("value=%s",
		expectVote,
	)

	vote := &Vote{}
	if err := z.connectPOST(ctx, vote, url, http.StatusOK, strings.NewReader(params)); err != nil {
		return nil, errors.Wrapf(err, "zendesk: [CreateVote] connect failed")
	}

	return vote, nil
}

// InstantSearch returns search result depends on query text, country code and locale.
func (z *ZenDesk) InstantSearch(ctx context.Context, queryText, countryCode, locale string) (*InstantSearch, error) {
	url := fmt.Sprintf("%s/hc/api/internal/instant_search.json?locale=%s&query=%s",
		z.identifyCountryCode(countryCode),
		locale,
		url.QueryEscape(queryText),
	)

	instantSearch := new(InstantSearch)
	if err := z.connectGET(ctx, instantSearch, url, http.StatusOK, nil); err != nil {
		return nil, errors.Wrapf(err, "zendesk: [InstantSearch] connect failed")
	}

	return instantSearch, nil
}

// Search returns search result depends on per page, page, query text, country code and locale.
func (z *ZenDesk) Search(ctx context.Context, categoryIDs []int, queryText, countryCode, locale string, pagination *Pagination) (*Search, error) {
	var filter string
	if len(categoryIDs) > 0 {
		filter += "&category="
		for _, id := range categoryIDs {
			filter += strconv.Itoa(id) + ","
		}
		filter = filter[:len(filter)-1]
	}

	url := fmt.Sprintf("%s/api/v2/help_center/articles/search.json?per_page=%d&page=%d&sort_order=%s&locale=%s&query=%s"+filter,
		z.identifyCountryCode(countryCode),
		pagination.PerPage,
		pagination.Page,
		pagination.SortOrder,
		locale,
		url.QueryEscape(queryText),
	)

	search := new(Search)
	if err := z.connectGET(ctx, search, url, http.StatusOK, nil); err != nil {
		return nil, errors.Wrapf(err, "zendesk: [Search] connect failed")
	}

	return search, nil
}

func (z *ZenDesk) identifyCountryCode(countryCode string) string {
	return z.urlTable[countryCode]
}
