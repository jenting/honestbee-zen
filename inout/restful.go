package inout

import (
	"context"
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	"github.com/honestbee/Zen/models"
)

// BaseIn is the basic input parameters.
type BaseIn struct {
	Locale      string `json:"locale,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	PerPage     int    `json:"per_page,omitempty"`
	Page        int    `json:"page,omitempty"`
	SortBy      string `json:"sort_by,omitempty"`
	SortOrder   string `json:"sort_order,omitempty"`
}

// BaseOut is the basic output parameters.
type BaseOut struct {
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
	PageCount int `json:"page_count"`
	Count     int `json:"count"`
}

// GetCategoriesIn is the input parameters of GET categories.
type GetCategoriesIn struct {
	*BaseIn
}

// GetCategoriesOut is the output parameters of GET categories.
type GetCategoriesOut struct {
	Categories []*models.Category `json:"categories"`
	*BaseOut
}

// GetCategoryKeyNameToIDIn is the input parameters of GET category_key_name_to_id.
type GetCategoryKeyNameToIDIn struct {
	CategoryKeyName string `json:"category_key_name,omitempty"`
	*BaseIn
}

// GetCategoryKeyNameToIDOut is the output parameters of GET category_key_name_to_id.
type GetCategoryKeyNameToIDOut struct {
	CategoryID int `json:"category_id,omitempty"`
}

// GetSectionsIn is the input parameters of GET categories/sections.
type GetSectionsIn struct {
	CategoryID int `json:"category_id,omitempty"`
	*BaseIn
}

// GetSectionsOut is the output parameters of GET sections.
type GetSectionsOut struct {
	Sections []*models.Section `json:"sections"`
	*BaseOut
}

// GetCategoriesArticlesIn is the input parameters of GET categories/articles.
type GetCategoriesArticlesIn struct {
	CategoryID int    `json:"category_id,omitempty"`
	LabelNames string `json:"label_names"`
	*BaseIn
}

// GetSectionIn is the input parameters of GET sections.
type GetSectionIn struct {
	SectionID int `json:"section_id,omitempty"`
	*BaseIn
}

// GetSectionOut is the output parameters of GET sections.
type GetSectionOut struct {
	Section *models.Section `json:"section"`
}

// GetArticlesIn is the input parameters of GET sections/articles.
type GetArticlesIn struct {
	SectionID int `json:"section_id,omitempty"`
	*BaseIn
}

// GetArticlesOut is the output parameters of GET sections/articles.
type GetArticlesOut struct {
	Articles []*models.Article `json:"articles"`
	*BaseOut
}

// GetArticleIn is the input parameters of GET article.
type GetArticleIn struct {
	ArticleID   int    `json:"article_id,omitempty"`
	Locale      string `json:"locale,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
}

// GetArticleOut is the output parameters of GET article.
type GetArticleOut struct {
	Article *models.Article `json:"article"`
}

// GetTopNArticlesIn is the input parameters of GET top_n articles.
type GetTopNArticlesIn struct {
	TopN        uint64 `json:"top_n,omitempty"`
	Locale      string `json:"locale,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
}

// GetTopNArticlesOut is the output parameters of GET top_n articles.
type GetTopNArticlesOut struct {
	Articles []*models.Article `json:"articles"`
}

// CreateRequestIn is the input parameters of POST request.
type CreateRequestIn struct {
	CountryCode string                 `json:"country_code,omitempty"`
	Data        map[string]interface{} `json:"data,omitempty"`
}

// GetTicketFormIn is the input parameters of GET ticket_form.
type GetTicketFormIn struct {
	CountryCode string `json:"country_code,omitempty"`
	Locale      string `json:"locale,omitempty"`
	FormID      int    `json:"form_id,omitempty"`
}

// GetTicketFormOut is the output parameters of GET ticket_form.
type GetTicketFormOut struct {
	TicketForm *models.TicketForm `json:"ticket_form"`
}

// CreateVoteIn is the input parameters of POST vote.
type CreateVoteIn struct {
	ArticleID   int    `json:"article_id,omitempty"`
	Value       string `json:"value,omitempty"`
	Locale      string `json:"locale,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
}

// CreateVoteOut is the output parameters of POST vote.
type CreateVoteOut struct {
	VoteSum   int `json:"vote_sum"`
	VoteCount int `json:"vote_count"`
}

// GetInstantSearchIn is the input parameters of GET instant_search.
type GetInstantSearchIn struct {
	Query       string `json:"query,omitempty"`
	Locale      string `json:"locale,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
}

// GetInstantSearchOut is the output parameters of GET instant_search.
type GetInstantSearchOut struct {
	Results []*InstantSearchResult `json:"results"`
}

// InstantSearchResult is the output parameters of GET instant_search.
type InstantSearchResult struct {
	Title         string `json:"title"`
	CategoryTitle string `json:"category_title"`
	URL           string `json:"url"`
}

// GetSearchIn is the input parameters of GET search.
type GetSearchIn struct {
	Query string `json:"query,omitempty"`
	*BaseIn
}

// GetSearchOut is the output parameters of GET search.
type GetSearchOut struct {
	Articles []*models.SearchArticle `json:"results"`
	*BaseOut
}

// GetBasicAuthIn is the input parameters of basic auth.
type GetBasicAuthIn struct {
	User string
	Pwd  string
}

// GraphQLIn is the input parameters of GraphQL query.
type GraphQLIn struct {
	Ctx     context.Context
	Queries []GraphQLQuery
	IsBatch bool
}

// GraphQLQuery is the input parameters of GraphQL query.
type GraphQLQuery struct {
	Query     string                 `json:"query"`
	OpName    string                 `json:"operationName"`
	Variables map[string]interface{} `json:"variables"`
}

// FetchBaseParams fetches RESTful base parameters.
func FetchBaseParams(r *http.Request) (*BaseIn, error) {
	locale := r.FormValue("locale")
	countryCode := r.FormValue("country_code")
	perPageStr := r.FormValue("per_page")
	pageStr := r.FormValue("page")
	sortBy := r.FormValue("sort_by")
	sortOrder := r.FormValue("sort_order")

	switch countryCode {
	case countryCodeSG:
	case countryCodeHK:
	case countryCodeTW:
	case countryCodeJP:
	case countryCodeTH:
	case countryCodeMY:
	case countryCodeID:
	case countryCodePH:
	case "":
		countryCode = defaultCountryCode
	default:
		return nil, errors.Errorf("inout: [fetchBaseIn] countryCode:%v is not in the list", countryCode)
	}

	switch locale {
	case localeZHTW:
	case localeZHCN:
	case localeENUS:
	case localeJA:
	case localeTH:
	case localeID:
	case "":
		locale = defaultLocale
	default:
		return nil, errors.Errorf("inout: [fetchBaseIn] locale:%v is not in the list", locale)
	}

	switch sortBy {
	case sortByPosition:
	case sortByCreatedAt:
	case sortByUpdatedAt:
	case "":
		sortBy = defaultSortBy
	default:
		return nil, errors.Errorf("inout: [fetchBaseIn] sort by:%v is not in the list", sortBy)
	}

	switch sortOrder {
	case sortOrderAsc:
	case sortOrderDesc:
	case "":
		sortOrder = defaultSortOrder
	default:
		return nil, errors.Errorf("inout: [fetchBaseIn] sort order:%v is not in the list", sortOrder)
	}

	perPage, err := strconv.ParseInt(perPageStr, 10, 64)
	if err != nil {
		perPage = defaultPerPage
	} else if perPage > maxPerPage {
		perPage = maxPerPage
	} else if perPage < minPerPage {
		return nil, errors.Errorf("inout: [fetchBaseIn] perPage:%v < %v", perPage, minPerPage)
	}

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	} else if page < minPage {
		return nil, errors.Errorf("inout: [fetchBaseIn] page:%v < %v", page, minPage)
	}

	// convert to offset
	page = (page - 1) * perPage

	return &BaseIn{
		Locale:      locale,
		CountryCode: countryCode,
		PerPage:     int(perPage),
		Page:        int(page),
		SortBy:      sortBy,
		SortOrder:   sortOrder,
	}, nil
}
