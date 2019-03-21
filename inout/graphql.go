package inout

import (
	gographql "github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
)

const (
	graphqlEnumCountryCodeSG = "SG"
	graphqlEnumCountryCodeHK = "HK"
	graphqlEnumCountryCodeTW = "TW"
	graphqlEnumCountryCodeJP = "JP"
	graphqlEnumCountryCodeTH = "TH"
	graphqlEnumCountryCodeMY = "MY"
	graphqlEnumCountryCodeID = "ID"
	graphqlEnumCountryCodePH = "PH"
)

const (
	graphqlEnumLocaleZHTW = "ZH_TW"
	graphqlEnumLocaleZHCN = "ZH_CN"
	graphqlEnumLocaleENUS = "EN_US"
	graphqlEnumLocaleJA   = "JA"
	graphqlEnumLocaleTH   = "TH"
	graphqlEnumLocaleID   = "ID"
)

const (
	graphqlEnumSortByPostion   = "POSITION"
	qraphqlEnumSortByCreatedAt = "CREATED_AT"
	graphqlEnumSortByUpdatedAt = "UPDATED_AT"
)

const (
	graphqlEnumSortOrderAsc  = "ASC"
	graphqlEnumSortOrderDesc = "DESC"
)

const (
	graphqlEnumVoteUp   = "UP"
	graphqlEnumVoteDown = "DOWN"
)

var graphqlCountryCodeMap = map[string]string{
	graphqlEnumCountryCodeSG: countryCodeSG,
	graphqlEnumCountryCodeHK: countryCodeHK,
	graphqlEnumCountryCodeTW: countryCodeTW,
	graphqlEnumCountryCodeJP: countryCodeJP,
	graphqlEnumCountryCodeTH: countryCodeTH,
	graphqlEnumCountryCodeMY: countryCodeMY,
	graphqlEnumCountryCodeID: countryCodeID,
	graphqlEnumCountryCodePH: countryCodePH,
}

var graphqlLocaleMap = map[string]string{
	graphqlEnumLocaleZHTW: localeZHTW,
	graphqlEnumLocaleZHCN: localeZHCN,
	graphqlEnumLocaleENUS: localeENUS,
	graphqlEnumLocaleJA:   localeJA,
	graphqlEnumLocaleTH:   localeTH,
	graphqlEnumLocaleID:   localeID,
}

var graphqlSortByMap = map[string]string{
	graphqlEnumSortByPostion:   sortByPosition,
	qraphqlEnumSortByCreatedAt: sortByCreatedAt,
	graphqlEnumSortByUpdatedAt: sortByUpdatedAt,
}

var graphqlSortOrderMap = map[string]string{
	graphqlEnumSortOrderAsc:  sortOrderAsc,
	graphqlEnumSortOrderDesc: sortOrderDesc,
}

var graphqlVoteMap = map[string]string{
	graphqlEnumVoteUp:   voteUp,
	graphqlEnumVoteDown: voteDown,
}

func processGraphQLCountryCode(countryCode string) (string, error) {
	val, ok := graphqlCountryCodeMap[countryCode]
	if !ok {
		return "", errors.Errorf("inout: [processGraphQL] countryCode:%v is not in the list", countryCode)
	}
	return val, nil
}

func processGraphQLLocale(locale string) (string, error) {
	val, ok := graphqlLocaleMap[locale]
	if !ok {
		return "", errors.Errorf("inout: [processGraphQL] locale:%v is not in the list", locale)
	}
	return val, nil
}

func processGraphQLSortBy(sortBy string) (string, error) {
	val, ok := graphqlSortByMap[sortBy]
	if !ok {
		return "", errors.Errorf("inout: [processGraphQL] sort by:%v is not in the list", sortBy)
	}
	return val, nil
}

func processGraphQLSortOrder(sortOrder string) (string, error) {
	val, ok := graphqlSortOrderMap[sortOrder]
	if !ok {
		return "", errors.Errorf("inout: [processGraphQL] sort order:%v is not in the list", sortOrder)
	}
	return val, nil
}

func processGraphQLVote(vote string) (string, error) {
	val, ok := graphqlVoteMap[vote]
	if !ok {
		return "", errors.Errorf("inout: [processGraphQL] vote:%v is not in the list", vote)
	}
	return val, nil
}

// ProcessPage process input params perPage and page.
func ProcessPage(perPage, page int32) (int32, int32) {
	if perPage > maxPerPage {
		perPage = maxPerPage
	} else if perPage < minPerPage {
		perPage = minPerPage
	}

	if page < minPage {
		page = minPage
	}
	page = (page - 1) * (perPage)

	return perPage, page
}

// QueryCategoriesIn are the arguments for the "allCategories" query.
type QueryCategoriesIn struct {
	CountryCode string
	Locale      string
	PerPage     int32
	Page        int32
	SortBy      string
	SortOrder   string
}

// ProcessInputParams process QueryCategoriesIn input parameters.
func (in *QueryCategoriesIn) ProcessInputParams() error {
	var err error

	in.CountryCode, err = processGraphQLCountryCode(in.CountryCode)
	if err != nil {
		return err
	}

	in.Locale, err = processGraphQLLocale(in.Locale)
	if err != nil {
		return err
	}

	in.SortBy, err = processGraphQLSortBy(in.SortBy)
	if err != nil {
		return err
	}

	in.SortOrder, err = processGraphQLSortOrder(in.SortOrder)
	if err != nil {
		return err
	}

	in.PerPage, in.Page = ProcessPage(in.PerPage, in.Page)

	return nil
}

// QueryCategoryIn are the arguments for the "oneCategory" query.
type QueryCategoryIn struct {
	CategoryIDOrKeyName gographql.ID
	CountryCode         string
	Locale              string
}

// ProcessInputParams process QueryCategoryIn input parameters.
func (in *QueryCategoryIn) ProcessInputParams() error {
	var err error

	in.CountryCode, err = processGraphQLCountryCode(in.CountryCode)
	if err != nil {
		return err
	}

	in.Locale, err = processGraphQLLocale(in.Locale)
	if err != nil {
		return err
	}

	return nil
}

// QuerySectionsIn are the arguments for the "allSections" query.
type QuerySectionsIn struct {
	CategoryID  *gographql.ID
	CountryCode string
	Locale      string
	PerPage     int32
	Page        int32
	SortBy      string
	SortOrder   string
}

// ProcessInputParams process QuerySectionsIn input parameters.
func (in *QuerySectionsIn) ProcessInputParams() error {
	var err error

	in.CountryCode, err = processGraphQLCountryCode(in.CountryCode)
	if err != nil {
		return err
	}

	in.Locale, err = processGraphQLLocale(in.Locale)
	if err != nil {
		return err
	}

	in.SortBy, err = processGraphQLSortBy(in.SortBy)
	if err != nil {
		return err
	}

	in.SortOrder, err = processGraphQLSortOrder(in.SortOrder)
	if err != nil {
		return err
	}

	in.PerPage, in.Page = ProcessPage(in.PerPage, in.Page)

	return nil
}

// QuerySectionIn are the arguments for the "oneSection" query.
type QuerySectionIn struct {
	SectionID   gographql.ID
	CountryCode string
	Locale      string
}

// ProcessInputParams process QuerySectionIn input parameters.
func (in *QuerySectionIn) ProcessInputParams() error {
	var err error

	in.CountryCode, err = processGraphQLCountryCode(in.CountryCode)
	if err != nil {
		return err
	}

	in.Locale, err = processGraphQLLocale(in.Locale)
	if err != nil {
		return err
	}

	return nil
}

// QueryArticlesIn are the arguments for the "allArticles" query.
type QueryArticlesIn struct {
	CategoryID  *gographql.ID
	SectionID   *gographql.ID
	CountryCode string
	Locale      string
	PerPage     int32
	Page        int32
	SortBy      string
	SortOrder   string
}

// ProcessInputParams process QueryArticlesIn input parameters.
func (in *QueryArticlesIn) ProcessInputParams() error {
	var err error

	in.CountryCode, err = processGraphQLCountryCode(in.CountryCode)
	if err != nil {
		return err
	}

	in.Locale, err = processGraphQLLocale(in.Locale)
	if err != nil {
		return err
	}

	in.SortBy, err = processGraphQLSortBy(in.SortBy)
	if err != nil {
		return err
	}

	in.SortOrder, err = processGraphQLSortOrder(in.SortOrder)
	if err != nil {
		return err
	}

	in.PerPage, in.Page = ProcessPage(in.PerPage, in.Page)

	return nil
}

// QueryTopArticlesIn are the arguments for the "topArticles" query.
type QueryTopArticlesIn struct {
	TopN        int32
	CountryCode string
	Locale      string
}

// ProcessInputParams process QueryTopArticlesIn input parameters.
func (in *QueryTopArticlesIn) ProcessInputParams() error {
	var err error

	in.CountryCode, err = processGraphQLCountryCode(in.CountryCode)
	if err != nil {
		return err
	}

	in.Locale, err = processGraphQLLocale(in.Locale)
	if err != nil {
		return err
	}

	return nil
}

// QueryArticleIn are the arguments for the "oneArticle" query.
type QueryArticleIn struct {
	ArticleID   gographql.ID
	CountryCode string
	Locale      string
}

// ProcessInputParams process QueryArticleIn input parameters.
func (in *QueryArticleIn) ProcessInputParams() error {
	var err error

	in.CountryCode, err = processGraphQLCountryCode(in.CountryCode)
	if err != nil {
		return err
	}

	in.Locale, err = processGraphQLLocale(in.Locale)
	if err != nil {
		return err
	}

	return nil
}

// QueryTicketFormIn are the arguments for the "ticketForm" query.
type QueryTicketFormIn struct {
	FormID gographql.ID
}

// QueryTicketFieldsIn are the arguments for the "ticketField" query.
type QueryTicketFieldsIn struct {
	FormID *gographql.ID
	Locale string
}

// ProcessInputParams process QueryTicketFieldsIn input parameters.
func (in *QueryTicketFieldsIn) ProcessInputParams() error {
	var err error

	in.Locale, err = processGraphQLLocale(in.Locale)
	if err != nil {
		return err
	}

	return nil
}

// QueryCustomFieldOptionsIn are the arguments for the "customFieldOptions" query.
type QueryCustomFieldOptionsIn struct {
	FieldID gographql.ID
}

// QuerySystemFieldOptionsIn are the arguments for the "systemFieldOptions" query.
type QuerySystemFieldOptionsIn struct {
	FieldID gographql.ID
}

// QuerySearchTitleArticlesIn are the arguments for the "searchTitleArticles" query.
type QuerySearchTitleArticlesIn struct {
	Query       string
	CountryCode string
	Locale      string
}

// ProcessInputParams process QuerySearchTitleArticlesIn input parameters.
func (in *QuerySearchTitleArticlesIn) ProcessInputParams() error {
	var err error

	in.CountryCode, err = processGraphQLCountryCode(in.CountryCode)
	if err != nil {
		return err
	}

	in.Locale, err = processGraphQLLocale(in.Locale)
	if err != nil {
		return err
	}

	return err
}

// QuerySearchBodyArticlesIn are the arguments for the "searchBodyArticles" query.
type QuerySearchBodyArticlesIn struct {
	Query       string
	CountryCode string
	Locale      string
	PerPage     int32
	Page        int32
	SortOrder   string
}

// ProcessInputParams process QuerySearchBodyArticlesIn input parameters.
func (in *QuerySearchBodyArticlesIn) ProcessInputParams() error {
	var err error

	in.CountryCode, err = processGraphQLCountryCode(in.CountryCode)
	if err != nil {
		return err
	}

	in.Locale, err = processGraphQLLocale(in.Locale)
	if err != nil {
		return err
	}

	in.SortOrder, err = processGraphQLSortOrder(in.SortOrder)
	if err != nil {
		return err
	}

	in.PerPage, in.Page = ProcessPage(in.PerPage, in.Page)

	return nil
}

// MutationRequestsIn are the arguments for the "requests" mutation.
type MutationRequestsIn struct {
	CountryCode string            `json:"country_code"`
	Data        CreateRequestData `json:"data"`
}

// CreateRequestData are the definition of create request data field.
type CreateRequestData struct {
	Request CreateRequestDataRequest `json:"request"`
}

// CreateRequestDataRequest are the definition of create request data request field.
type CreateRequestDataRequest struct {
	Comment      CreateRequestDataRequestComment        `json:"comment"`
	Requester    CreateRequestDataRequestRequester      `json:"requester"`
	Subject      string                                 `json:"subject"`
	TicketFormID *string                                `json:"ticket_form_id,omitempty"`
	CustomFields *[]CreateRequestDataRequestCustomField `json:"custom_fields,omitempty"`
}

// CreateRequestDataRequestComment are the definition of create request comment field.
type CreateRequestDataRequestComment struct {
	Body string `json:"body"`
}

// CreateRequestDataRequestCustomField are the definition of create request custom field field.
type CreateRequestDataRequestCustomField struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

// CreateRequestDataRequestRequester are the definition of create request requester field.
type CreateRequestDataRequestRequester struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ProcessInputParams process MutationRequestsIn input parameters.
func (in *MutationRequestsIn) ProcessInputParams() error {
	var err error

	in.CountryCode, err = processGraphQLCountryCode(in.CountryCode)
	if err != nil {
		return err
	}

	return nil
}

// MutationVoteArticleIn are the arguments for the "voteArticle" mutation.
type MutationVoteArticleIn struct {
	ArticleID   gographql.ID
	Vote        string
	CountryCode string
	Locale      string
}

// ProcessInputParams process MutationVoteArticleIn input parameters.
func (in *MutationVoteArticleIn) ProcessInputParams() error {
	var err error

	in.CountryCode, err = processGraphQLCountryCode(in.CountryCode)
	if err != nil {
		return err
	}

	in.Locale, err = processGraphQLLocale(in.Locale)
	if err != nil {
		return err
	}

	in.Vote, err = processGraphQLVote(in.Vote)
	if err != nil {
		return err
	}

	return nil
}

// MutationForceSyncIn are the arguments for the "forceSync" mutation.
type MutationForceSyncIn struct {
	Username string
	Password string
}
