package models

import (
	"context"
	"reflect"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
)

const (
	// ModelsReturnErrorCountryCode is a mock country code for return error.
	ModelsReturnErrorCountryCode = "th"
	// ModelsReturnNotFoundCountryCode is a mock country code for return not found.
	ModelsReturnNotFoundCountryCode = "id"
	// ModelsReturnErrorLocale is a mock locale for return error.
	ModelsReturnErrorLocale = "zh-cn"
	// ModelsReturnNotFoundLocale is a mock locale for return not found.
	ModelsReturnNotFoundLocale = "id"
	// PlusCounterReturnErrorCountryCode is a mock country code for plus counter return return error.
	PlusCounterReturnErrorCountryCode = "not_exist_country_code_1"
	// PlusCounterReturnSmallerCountryCode is a mock country code for plus counter return smaller counter.
	PlusCounterReturnSmallerCountryCode = "not_exist_country_code_2"
	// LockCounterReturnErrorCountryCode is a mock country code for lock counter return error..
	LockCounterReturnErrorCountryCode = "not_exist_country_code_3"
	// LockCounterFailedCountryCode is a mock country code for lock counter return failed country code.
	LockCounterFailedCountryCode = "not_exist_country_code_4"
	// ResetCounterFailedCountryCode is a mock country code for reset counter return failed country code.
	ResetCounterFailedCountryCode = "sg"
	// UnlockCounterFailedCountryCode is a mock country code for unlink counter return failed country code.
	UnlockCounterFailedCountryCode = "hk"
	// SyncDBFailedLocale is a mock for locale for sync db return failed.
	SyncDBFailedLocale = "ja"
)

var (
	// FixCreatedAt1 is a mock created_at time.
	FixCreatedAt1 = time.Now().UTC()
	// FixUpdatedAt1 is a mock updated_at time.
	FixUpdatedAt1 = time.Now().UTC()
	// FixEditedAt1 is a mock edited_at time.
	FixEditedAt1 = time.Now().UTC()
)

var (
	// FixCreatedAtProto1 is a mock created_at protobuf time.
	FixCreatedAtProto1, _ = ptypes.TimestampProto(FixCreatedAt1)
	// FixUpdatedAtProto1 is a mock updated_at time.
	FixUpdatedAtProto1, _ = ptypes.TimestampProto(FixUpdatedAt1)
	// FixEditedAtProto1 is a mock edited_at time.
	FixEditedAtProto1, _ = ptypes.TimestampProto(FixEditedAt1)
)

// MockModels is a mock service.
type MockModels struct {
	Sequence map[string]bool
}

// NewMockService return a new mock service with sequece initialized.
func NewMockService() *MockModels {
	return &MockModels{
		Sequence: make(map[string]bool),
	}
}

// ResetSequence reset the sequence data.
func (m *MockModels) ResetSequence() {
	m.Sequence = make(map[string]bool)
}

// Close is the mock function of Close.
func (m *MockModels) Close() error { return nil }

// SyncWithDynamicContentItems is the mock function of SyncWithDynamicContentItems.
func (m *MockModels) SyncWithDynamicContentItems(ctx context.Context, zendeskDCItems []*SyncDynamicContentItem) error {
	if m.Sequence != nil {
		m.Sequence["SyncWithDynamicContentItems"] = true
	}
	return nil
}

// SyncWithTicketFields is the mock function of SyncWithTicketFields.
func (m *MockModels) SyncWithTicketFields(ctx context.Context, zendeskTicketFields []*SyncTicketField) error {
	if m.Sequence != nil {
		m.Sequence["SyncWithTicketFields"] = true
	}
	return nil
}

// SyncWithTicketForms is the mock function of SyncWithTicketForms.
func (m *MockModels) SyncWithTicketForms(ctx context.Context, zendeskTicketForms []*SyncTicketForm) error {
	if m.Sequence != nil {
		m.Sequence["SyncWithTicketForms"] = true
	}
	return nil
}

// PlusOneTicketFormsCounter is the mock function of PlusOneTicketFormsCounter.
func (m *MockModels) PlusOneTicketFormsCounter(ctx context.Context) (int, error) {
	if m.Sequence != nil {
		m.Sequence["PlusOneTicketFormsCounter"] = true
	}
	return 10, nil
}

// ResetTicketFormsCounter is the mock function of ResetTicketFormsCounter.
func (m *MockModels) ResetTicketFormsCounter(ctx context.Context) error {
	if m.Sequence != nil {
		m.Sequence["ResetTicketFormsCounter"] = true
	}
	return nil
}

// LockTicketFormsCounter is the mock function of LockTicketFormsCounter.
func (m *MockModels) LockTicketFormsCounter(ctx context.Context) (bool, error) {
	if m.Sequence != nil {
		m.Sequence["LockTicketFormsCounter"] = true
	}
	return true, nil
}

// UnlockTicketFormsCounter is the mock function of UnlockTicketFormsCounter.
func (m *MockModels) UnlockTicketFormsCounter(ctx context.Context) error {
	if m.Sequence != nil {
		m.Sequence["UnlockTicketFormsCounter"] = true
	}
	return nil
}

// SyncWithArticles is the mock function of SyncWithArticles.
func (m *MockModels) SyncWithArticles(ctx context.Context, zendeskArticles []*Article, countryCode, locale string) error {
	if m.Sequence != nil {
		m.Sequence["SyncWithArticles"] = true
	}
	if len(zendeskArticles) > 0 {
		if zendeskArticles[0].Locale == SyncDBFailedLocale {
			return errors.Errorf("return error")
		}
	}
	return nil
}

// SyncWithArticle is the mock function of SyncWithArticle.
func (m *MockModels) SyncWithArticle(ctx context.Context, articleID int, zendeskArticle *Article, countryCode, locale string) error {
	if m.Sequence != nil {
		m.Sequence["SyncWithArticle"] = true
	}
	return nil
}

// SyncWithSections is the mock function of SyncWithSections.
func (m *MockModels) SyncWithSections(ctx context.Context, zendeskSections []*Section, countryCode, locale string) error {
	if m.Sequence != nil {
		m.Sequence["SyncWithSections"] = true
	}
	if len(zendeskSections) > 0 {
		if zendeskSections[0].Locale == SyncDBFailedLocale {
			return errors.Errorf("return error")
		}
	}
	return nil
}

// SyncWithCategories is the mock function of SyncWithCategories.
func (m *MockModels) SyncWithCategories(ctx context.Context, zendeskCategories []*Category, countryCode, locale string) error {
	if m.Sequence != nil {
		m.Sequence["SyncWithCategories"] = true
	}
	if len(zendeskCategories) > 0 {
		if zendeskCategories[0].Locale == SyncDBFailedLocale {
			return errors.Errorf("return error")
		}
	}
	return nil
}

// PlusOneCategoriesCounter is the mock function of PlusOneCategoriesCounter.
func (m *MockModels) PlusOneCategoriesCounter(ctx context.Context, countryCode, locale string) (int, error) {
	if m.Sequence != nil {
		m.Sequence["PlusOneCategoriesCounter"] = true
	}

	switch countryCode {
	case PlusCounterReturnErrorCountryCode:
		return 0, errors.Errorf("return error")
	case PlusCounterReturnSmallerCountryCode:
		return 0, nil
	}
	return 10, nil
}

// PlusOneSectionsCounter is the mock function of PlusOneSectionsCounter.
func (m *MockModels) PlusOneSectionsCounter(ctx context.Context, countryCode, locale string) (int, error) {
	if m.Sequence != nil {
		m.Sequence["PlusOneSectionsCounter"] = true
	}

	switch countryCode {
	case PlusCounterReturnErrorCountryCode:
		return 0, errors.Errorf("return error")
	case PlusCounterReturnSmallerCountryCode:
		return 0, nil
	}
	return 10, nil
}

// PlusOneArticlesCounter is the mock function of PlusOneArticlesCounter.
func (m *MockModels) PlusOneArticlesCounter(ctx context.Context, countryCode, locale string) (int, error) {
	if m.Sequence != nil {
		m.Sequence["PlusOneArticlesCounter"] = true
	}

	switch countryCode {
	case PlusCounterReturnErrorCountryCode:
		return 0, errors.Errorf("return error")
	case PlusCounterReturnSmallerCountryCode:
		return 0, nil
	}
	return 10, nil
}

// ResetCategoriesCounter is the mock function of ResetCategoriesCounter.
func (m *MockModels) ResetCategoriesCounter(ctx context.Context, countryCode, locale string) error {
	if m.Sequence != nil {
		m.Sequence["ResetCategoriesCounter"] = true
	}

	switch countryCode {
	case ResetCounterFailedCountryCode:
		return errors.Errorf("return error")
	}
	return nil
}

// ResetSectionsCounter is the mock function of ResetSectionsCounter.
func (m *MockModels) ResetSectionsCounter(ctx context.Context, countryCode, locale string) error {
	if m.Sequence != nil {
		m.Sequence["ResetSectionsCounter"] = true
	}

	switch countryCode {
	case ResetCounterFailedCountryCode:
		return errors.Errorf("return error")
	}
	return nil
}

// ResetArticlesCounter is the mock function of ResetArticlesCounter.
func (m *MockModels) ResetArticlesCounter(ctx context.Context, countryCode, locale string) error {
	if m.Sequence != nil {
		m.Sequence["ResetArticlesCounter"] = true
	}

	switch countryCode {
	case ResetCounterFailedCountryCode:
		return errors.Errorf("return error")
	}
	return nil
}

// LockCategoriesCounter is the mock function of LockCategoriesCounter.
func (m *MockModels) LockCategoriesCounter(ctx context.Context, countryCode, locale string) (bool, error) {
	if m.Sequence != nil {
		m.Sequence["LockCategoriesCounter"] = true
	}

	switch countryCode {
	case LockCounterReturnErrorCountryCode:
		return false, errors.Errorf("return error")
	case LockCounterFailedCountryCode:
		return false, nil
	}
	return true, nil
}

// LockSectionsCounter is the mock function of LockSectionsCounter.
func (m *MockModels) LockSectionsCounter(ctx context.Context, countryCode, locale string) (bool, error) {
	if m.Sequence != nil {
		m.Sequence["LockSectionsCounter"] = true
	}

	switch countryCode {
	case LockCounterReturnErrorCountryCode:
		return false, errors.Errorf("return error")
	case LockCounterFailedCountryCode:
		return false, nil
	}
	return true, nil
}

// LockArticlesCounter is the mock function of LockArticlesCounter.
func (m *MockModels) LockArticlesCounter(ctx context.Context, countryCode, locale string) (bool, error) {
	if m.Sequence != nil {
		m.Sequence["LockArticlesCounter"] = true
	}

	switch countryCode {
	case LockCounterReturnErrorCountryCode:
		return false, errors.Errorf("return error")
	case LockCounterFailedCountryCode:
		return false, nil
	}
	return true, nil
}

// UnlockCategoriesCounter is the mock function of UnlockCategoriesCounter.
func (m *MockModels) UnlockCategoriesCounter(ctx context.Context, countryCode, locale string) error {
	if m.Sequence != nil {
		m.Sequence["UnlockCategoriesCounter"] = true
	}

	switch countryCode {
	case UnlockCounterFailedCountryCode:
		return errors.Errorf("return error")
	}
	return nil
}

// UnlockSectionsCounter is the mock function of UnlockSectionsCounter.
func (m *MockModels) UnlockSectionsCounter(ctx context.Context, countryCode, locale string) error {
	if m.Sequence != nil {
		m.Sequence["UnlockSectionsCounter"] = true
	}

	switch countryCode {
	case UnlockCounterFailedCountryCode:
		return errors.Errorf("return error")
	}
	return nil
}

// UnlockArticlesCounter is the mock function of UnlockArticlesCounter.
func (m *MockModels) UnlockArticlesCounter(ctx context.Context, countryCode, locale string) error {
	if len(m.Sequence) != 0 {
		m.Sequence["UnlockArticlesCounter"] = true
	}

	switch countryCode {
	case UnlockCounterFailedCountryCode:
		return errors.Errorf("return error")
	}
	return nil
}

// CategoriesCacheGet is the mock function of CategoriesCacheGet.
func (m *MockModels) CategoriesCacheGet(ctx context.Context, key, countryCode, locale string) (string, bool) {
	if m.Sequence != nil {
		m.Sequence["CategoriesCacheGet"] = true
	}
	return "", false
}

// CategoriesCacheSet is the mock function of CategoriesCacheSet.
func (m *MockModels) CategoriesCacheSet(ctx context.Context, key, value, countryCode, locale string) (bool, error) {
	if m.Sequence != nil {
		m.Sequence["CategoriesCacheSet"] = true
	}
	return true, nil
}

// CategoriesCacheInvalidate is the mock function of .
func (m *MockModels) CategoriesCacheInvalidate(ctx context.Context, countryCode, locale string) error {
	if m.Sequence != nil {
		m.Sequence["CategoriesCacheInvalidate"] = true
	}
	return nil
}

// SectionsCacheGet is the mock function of SectionsCacheGet.
func (m *MockModels) SectionsCacheGet(ctx context.Context, key, countryCode, locale string) (string, bool) {
	if m.Sequence != nil {
		m.Sequence["SectionsCacheGet"] = true
	}
	return "", false
}

// SectionsCacheSet is the mock function of SectionsCacheSet.
func (m *MockModels) SectionsCacheSet(ctx context.Context, key, value, countryCode, locale string) (bool, error) {
	if m.Sequence != nil {
		m.Sequence["SectionsCacheSet"] = true
	}
	return true, nil
}

// SectionsCacheInvalidate is the mock function of SectionsCacheInvalidate.
func (m *MockModels) SectionsCacheInvalidate(ctx context.Context, countryCode, locale string) error {
	if m.Sequence != nil {
		m.Sequence["SectionsCacheInvalidate"] = true
	}
	return nil
}

// ArticlesCacheGet is the mock function of ArticlesCacheGet.
func (m *MockModels) ArticlesCacheGet(ctx context.Context, key, countryCode, locale string) (string, bool) {
	if m.Sequence != nil {
		m.Sequence["ArticlesCacheGet"] = true
	}
	return "", false
}

// ArticlesCacheSet is the mock function of ArticlesCacheSet.
func (m *MockModels) ArticlesCacheSet(ctx context.Context, key, value, countryCode, locale string) (bool, error) {
	if m.Sequence != nil {
		m.Sequence["ArticlesCacheSet"] = true
	}
	return true, nil
}

// ArticlesCacheInvalidate is the mock function of ArticlesCacheInvalidate.
func (m *MockModels) ArticlesCacheInvalidate(ctx context.Context, countryCode, locale string) error {
	if m.Sequence != nil {
		m.Sequence["ArticlesCacheInvalidate"] = true
	}
	return nil
}

// TicketFormCacheGet is the mock function of TicketFormCacheGet.
func (m *MockModels) TicketFormCacheGet(ctx context.Context, key string) (string, bool) {
	if m.Sequence != nil {
		m.Sequence["TicketFormCacheGet"] = true
	}
	return "", false
}

// TicketFormCacheSet is the mock function of TicketFormCacheSet.
func (m *MockModels) TicketFormCacheSet(ctx context.Context, key, value string) (bool, error) {
	if m.Sequence != nil {
		m.Sequence["TicketFormCacheSet"] = true
	}
	return true, nil
}

// TicketFormCacheInvalidate is the mock function of TicketFormCacheInvalidate.
func (m *MockModels) TicketFormCacheInvalidate(ctx context.Context) error {
	if m.Sequence != nil {
		m.Sequence["TicketFormCacheInvalidate"] = true
	}
	return nil
}

// TicketFieldCacheGet is the mock function of TicketFieldCacheGet.
func (m *MockModels) TicketFieldCacheGet(ctx context.Context, key string) (string, bool) {
	if m.Sequence != nil {
		m.Sequence["TicketFieldCacheGet"] = true
	}
	return "", false
}

// TicketFieldCacheSet is the mock function of TicketFieldCacheSet.
func (m *MockModels) TicketFieldCacheSet(ctx context.Context, key, value string) (bool, error) {
	if m.Sequence != nil {
		m.Sequence["TicketFieldCacheSet"] = true
	}
	return true, nil
}

// TicketFieldCacheInvalidate is the mock function of TicketFieldCacheInvalidate.
func (m *MockModels) TicketFieldCacheInvalidate(ctx context.Context) error {
	if m.Sequence != nil {
		m.Sequence["TicketFieldCacheInvalidate"] = true
	}
	return nil
}

// TicketFieldCustomFieldOptionCacheGet is the mock function of TicketFieldCustomFieldOptionCacheGet.
func (m *MockModels) TicketFieldCustomFieldOptionCacheGet(ctx context.Context, key string) (string, bool) {
	if m.Sequence != nil {
		m.Sequence["TicketFieldCustomFieldOptionCacheGet"] = true
	}
	return "", false
}

// TicketFieldCustomFieldOptionCacheSet is the mock function of TicketFieldCustomFieldOptionCacheSet.
func (m *MockModels) TicketFieldCustomFieldOptionCacheSet(ctx context.Context, key, value string) (bool, error) {
	if m.Sequence != nil {
		m.Sequence["TicketFieldCustomFieldOptionCacheSet"] = true
	}
	return true, nil
}

// TicketFieldCustomFieldOptionCacheInvalidate is the mock function of TicketFieldCustomFieldOptionCacheInvalidate.
func (m *MockModels) TicketFieldCustomFieldOptionCacheInvalidate(ctx context.Context) error {
	if m.Sequence != nil {
		m.Sequence["TicketFieldCustomFieldOptionCacheInvalidate"] = true
	}
	return nil
}

// TicketFieldSystemFieldOptionCacheGet is the mock function of TicketFieldSystemFieldOptionCacheGet.
func (m *MockModels) TicketFieldSystemFieldOptionCacheGet(ctx context.Context, key string) (string, bool) {
	if m.Sequence != nil {
		m.Sequence["TicketFieldSystemFieldOptionCacheGet"] = true
	}
	return "", false
}

// TicketFieldSystemFieldOptionCacheSet is the mock function of TicketFieldSystemFieldOptionCacheSet.
func (m *MockModels) TicketFieldSystemFieldOptionCacheSet(ctx context.Context, key, value string) (bool, error) {
	if m.Sequence != nil {
		m.Sequence["TicketFieldSystemFieldOptionCacheSet"] = true
	}
	return true, nil
}

// TicketFieldSystemFieldOptionCacheInvalidate is the mock function of TicketFieldSystemFieldOptionCacheInvalidate.
func (m *MockModels) TicketFieldSystemFieldOptionCacheInvalidate(ctx context.Context) error {
	if m.Sequence != nil {
		m.Sequence["TicketFieldSystemFieldOptionCacheInvalidate"] = true
	}
	return nil
}

// GetCategories is the mock function of GetCategories.
func (m *MockModels) GetCategories(ctx context.Context, params *GetCategoriesParams) ([]*Category, int, error) {
	switch params.CountryCode {
	case ModelsReturnErrorCountryCode:
		return nil, 0, errors.New("MockModels GetCategories return error")
	case ModelsReturnNotFoundCountryCode:
		return nil, 0, ErrNotFound
	}

	return []*Category{
		&Category{
			ID:           3345678,
			Position:     0,
			CreatedAt:    FixCreatedAt1,
			UpdatedAt:    FixUpdatedAt1,
			SourceLocale: "en-us",
			Outdated:     false,
			CountryCode:  "tw",
			URL:          "www.honestbee.com",
			HTMLURL:      "www.honestbee.com",
			Name:         "testing category 1",
			Description:  "",
			Locale:       "en-us",
			KeyName:      "food",
		},
	}, 1, nil
}

// GetCategoriesID is the mock function of GetCategoriesID.
func (m *MockModels) GetCategoriesID(ctx context.Context, countryCode string) ([]int, error) {
	switch countryCode {
	case ModelsReturnErrorCountryCode:
		return nil, errors.New("MockModels GetCategories return error")
	case ModelsReturnNotFoundCountryCode:
		return nil, ErrNotFound
	}

	return []int{3345678}, nil
}

// GetCategoryBySectionID is the mock function of GetCategoryBySectionID.
func (m *MockModels) GetCategoryBySectionID(ctx context.Context, sectionID int, locale string) (*Category, error) {
	switch locale {
	case ModelsReturnErrorLocale:
		return nil, errors.New("MockModels GetCategoryBySectionID return error")
	case ModelsReturnNotFoundLocale:
		return nil, ErrNotFound
	}

	return &Category{
		ID:           3345678,
		Position:     0,
		CreatedAt:    FixCreatedAt1,
		UpdatedAt:    FixUpdatedAt1,
		SourceLocale: "en-us",
		Outdated:     false,
		CountryCode:  "tw",
		URL:          "www.honestbee.com",
		HTMLURL:      "www.honestbee.com",
		Name:         "testing category 1",
		Description:  "",
		Locale:       "en-us",
	}, nil
}

// GetCategoryByArticleID is the mock function of GetCategoryByArticleID.
func (m *MockModels) GetCategoryByArticleID(ctx context.Context, articleID int, locale string) (*Category, error) {
	switch locale {
	case ModelsReturnErrorLocale:
		return nil, errors.New("MockModels GetCategoryByArticleID return error")
	case ModelsReturnNotFoundLocale:
		return nil, ErrNotFound
	}

	return &Category{
		ID:           3345678,
		Position:     0,
		CreatedAt:    FixCreatedAt1,
		UpdatedAt:    FixUpdatedAt1,
		SourceLocale: "en-us",
		Outdated:     false,
		CountryCode:  "tw",
		URL:          "www.honestbee.com",
		HTMLURL:      "www.honestbee.com",
		Name:         "testing category 1",
		Description:  "",
		Locale:       "en-us",
	}, nil
}

// GetCategoryKeyNameToID is the mock function of GetCategoryKeyNameToID.
func (m *MockModels) GetCategoryKeyNameToID(ctx context.Context, keyName, locale string) (int, error) {
	switch keyName {
	case "groceries":
		return 3345678, nil
	case "hahaha":
		return 0, ErrNotFound
	default:
		return 0, errors.New("something wrong")
	}
}

// GetCategoryByCategoryIDOrKeyName is the mock function of GetCategoryByCategoryIDOrKeyName.
func (m *MockModels) GetCategoryByCategoryIDOrKeyName(ctx context.Context, idOrKeyName, locale, countryCode string) (*Category, error) {
	switch countryCode {
	case ModelsReturnErrorCountryCode:
		return nil, errors.New("MockModels GetCategoryByCategoryIDOrKeyName return error")
	case ModelsReturnNotFoundCountryCode:
		return nil, ErrNotFound
	}

	return &Category{
		ID:           3345678,
		Position:     0,
		CreatedAt:    FixCreatedAt1,
		UpdatedAt:    FixUpdatedAt1,
		SourceLocale: "en-us",
		Outdated:     false,
		CountryCode:  "tw",
		URL:          "www.honestbee.com",
		HTMLURL:      "www.honestbee.com",
		Name:         "testing category 1",
		Description:  "",
		Locale:       "en-us",
		KeyName:      "food",
	}, nil
}

// GetSections is the mock function of GetSections.
func (m *MockModels) GetSections(ctx context.Context, params *GetSectionsParams) ([]*Section, int, error) {
	switch params.CountryCode {
	case ModelsReturnErrorCountryCode:
		return nil, 0, errors.New("MockModels GetSections return error")
	case ModelsReturnNotFoundCountryCode:
		return nil, 0, ErrNotFound
	}

	return []*Section{
		&Section{
			ID:           3345679,
			Position:     0,
			CreatedAt:    FixCreatedAt1,
			UpdatedAt:    FixUpdatedAt1,
			SourceLocale: "en-us",
			Outdated:     false,
			CountryCode:  "tw",
			URL:          "www.honestbee.com",
			HTMLURL:      "www.honestbee.com",
			Name:         "testing section 1",
			Description:  "",
			Locale:       "en-us",
			CategoryID:   3345678,
		},
	}, 1, nil
}

// GetSectionsByCategoryID is the mock function of GetSectionsByCategoryID.
func (m *MockModels) GetSectionsByCategoryID(ctx context.Context, params *GetSectionsParams) ([]*Section, int, error) {
	switch params.CountryCode {
	case ModelsReturnErrorCountryCode:
		return nil, 0, errors.New("MockModels GetSectionsByCategoryID return error")
	case ModelsReturnNotFoundCountryCode:
		return nil, 0, ErrNotFound
	}

	return []*Section{
		&Section{
			ID:           3345679,
			Position:     0,
			CreatedAt:    FixCreatedAt1,
			UpdatedAt:    FixUpdatedAt1,
			SourceLocale: "en-us",
			Outdated:     false,
			CountryCode:  "tw",
			URL:          "www.honestbee.com",
			HTMLURL:      "www.honestbee.com",
			Name:         "testing section 1",
			Description:  "",
			Locale:       "en-us",
			CategoryID:   3345678,
		},
	}, 1, nil
}

// GetSectionBySectionID is the mock function of GetSectionBySectionID.
func (m *MockModels) GetSectionBySectionID(ctx context.Context, sectionID int, locale, countryCode string) (*Section, error) {
	switch countryCode {
	case ModelsReturnErrorCountryCode:
		return nil, errors.New("MockModels GetSectionBySectionID return error")
	case ModelsReturnNotFoundCountryCode:
		return nil, ErrNotFound
	}

	return &Section{
		ID:           3345679,
		Position:     0,
		CreatedAt:    FixCreatedAt1,
		UpdatedAt:    FixUpdatedAt1,
		SourceLocale: "en-us",
		Outdated:     false,
		CountryCode:  "sg",
		URL:          "www.honestbee.com",
		HTMLURL:      "www.honestbee.com",
		Name:         "testing section 1",
		Description:  "",
		Locale:       "en-us",
		CategoryID:   3345678,
	}, nil
}

// GetSectionByArticleID is the mock function of GetSectionByArticleID.
func (m *MockModels) GetSectionByArticleID(ctx context.Context, articleID int, locale, countryCode string) (*Section, error) {
	switch countryCode {
	case ModelsReturnErrorCountryCode:
		return nil, errors.New("MockModels GetSectionBySectionID return error")
	case ModelsReturnNotFoundCountryCode:
		return nil, ErrNotFound
	}

	return &Section{
		ID:           3345679,
		Position:     0,
		CreatedAt:    FixCreatedAt1,
		UpdatedAt:    FixUpdatedAt1,
		SourceLocale: "en-us",
		Outdated:     false,
		CountryCode:  "sg",
		URL:          "www.honestbee.com",
		HTMLURL:      "www.honestbee.com",
		Name:         "testing section 1",
		Description:  "",
		Locale:       "en-us",
		CategoryID:   3345678,
	}, nil
}

// GetArticles is the mock function of GetArticles.
func (m *MockModels) GetArticles(ctx context.Context, params *GetArticlesParams) ([]*Article, int, error) {
	switch params.CountryCode {
	case ModelsReturnErrorCountryCode:
		return nil, 0, errors.New("MockModels GetArticles return error")
	case ModelsReturnNotFoundCountryCode:
		return nil, 0, ErrNotFound
	}

	return []*Article{
		&Article{
			ID:              33456710,
			AuthorID:        1234567,
			CommentsDisable: false,
			Draft:           false,
			Promoted:        false,
			Position:        0,
			VoteSum:         0,
			VoteCount:       0,
			CreatedAt:       FixCreatedAt1,
			UpdatedAt:       FixUpdatedAt1,
			SourceLocale:    "en-us",
			Outdated:        false,
			OutdatedLocales: []string{},
			EditedAt:        FixEditedAt1,
			LabelNames:      []string{},
			CountryCode:     "tw",
			URL:             "www.honestbee.com",
			HTMLURL:         "www.honestbee.com",
			Name:            "testing article 1",
			Title:           "testing article 1",
			Body:            "this is testing article 1",
			Locale:          "en-us",
			SectionID:       33456789,
		},
	}, 1, nil
}

// GetArticlesByCategoryID is the mock function of GetArticlesByCategoryID.
func (m *MockModels) GetArticlesByCategoryID(ctx context.Context, params *GetArticlesParams, labels []string) ([]*Article, int, error) {
	switch params.CountryCode {
	case ModelsReturnErrorCountryCode:
		return nil, 0, errors.New("MockModels GetArticlesByCategoryID return error")
	case ModelsReturnNotFoundCountryCode:
		return nil, 0, ErrNotFound
	}

	if len(labels) == 0 {
		return []*Article{
			&Article{
				ID:              33456710,
				AuthorID:        1234567,
				CommentsDisable: false,
				Draft:           false,
				Promoted:        false,
				Position:        0,
				VoteSum:         0,
				VoteCount:       0,
				CreatedAt:       FixCreatedAt1,
				UpdatedAt:       FixUpdatedAt1,
				SourceLocale:    "en-us",
				Outdated:        false,
				OutdatedLocales: []string{},
				EditedAt:        FixEditedAt1,
				LabelNames:      []string{"confirmed"},
				CountryCode:     "tw",
				URL:             "www.honestbee.com",
				HTMLURL:         "www.honestbee.com",
				Name:            "testing article 1",
				Title:           "testing article 1",
				Body:            "this is testing article 1",
				Locale:          "en-us",
				SectionID:       33456789,
			},
			&Article{
				ID:              33456711,
				AuthorID:        1234568,
				CommentsDisable: false,
				Draft:           false,
				Promoted:        false,
				Position:        0,
				VoteSum:         0,
				VoteCount:       0,
				CreatedAt:       FixCreatedAt1,
				UpdatedAt:       FixUpdatedAt1,
				SourceLocale:    "en-us",
				Outdated:        false,
				OutdatedLocales: []string{},
				EditedAt:        FixEditedAt1,
				LabelNames:      []string{"preparing"},
				CountryCode:     "tw",
				URL:             "www.honestbee.com",
				HTMLURL:         "www.honestbee.com",
				Name:            "testing article 2",
				Title:           "testing article 2",
				Body:            "this is testing article 2",
				Locale:          "en-us",
				SectionID:       33456789,
			},
			&Article{
				ID:              33456712,
				AuthorID:        1234569,
				CommentsDisable: false,
				Draft:           false,
				Promoted:        false,
				Position:        0,
				VoteSum:         0,
				VoteCount:       0,
				CreatedAt:       FixCreatedAt1,
				UpdatedAt:       FixUpdatedAt1,
				SourceLocale:    "en-us",
				Outdated:        false,
				OutdatedLocales: []string{},
				EditedAt:        FixEditedAt1,
				LabelNames:      []string{"ontheway"},
				CountryCode:     "tw",
				URL:             "www.honestbee.com",
				HTMLURL:         "www.honestbee.com",
				Name:            "testing article 3",
				Title:           "testing article 3",
				Body:            "this is testing article 3",
				Locale:          "en-us",
				SectionID:       33456789,
			},
			&Article{
				ID:              33456713,
				AuthorID:        1234570,
				CommentsDisable: false,
				Draft:           false,
				Promoted:        false,
				Position:        0,
				VoteSum:         0,
				VoteCount:       0,
				CreatedAt:       FixCreatedAt1,
				UpdatedAt:       FixUpdatedAt1,
				SourceLocale:    "en-us",
				Outdated:        false,
				OutdatedLocales: []string{},
				EditedAt:        FixEditedAt1,
				LabelNames:      []string{"delivered"},
				CountryCode:     "tw",
				URL:             "www.honestbee.com",
				HTMLURL:         "www.honestbee.com",
				Name:            "testing article 4",
				Title:           "testing article 4",
				Body:            "this is testing article 4",
				Locale:          "en-us",
				SectionID:       33456789,
			},
		}, 4, nil
	}

	if reflect.DeepEqual(labels, []string{"confirmed"}) {
		return []*Article{
			&Article{
				ID:              33456710,
				AuthorID:        1234567,
				CommentsDisable: false,
				Draft:           false,
				Promoted:        false,
				Position:        0,
				VoteSum:         0,
				VoteCount:       0,
				CreatedAt:       FixCreatedAt1,
				UpdatedAt:       FixUpdatedAt1,
				SourceLocale:    "en-us",
				Outdated:        false,
				OutdatedLocales: []string{},
				EditedAt:        FixEditedAt1,
				LabelNames:      []string{"confirmed"},
				CountryCode:     "tw",
				URL:             "www.honestbee.com",
				HTMLURL:         "www.honestbee.com",
				Name:            "testing article 1",
				Title:           "testing article 1",
				Body:            "this is testing article 1",
				Locale:          "en-us",
				SectionID:       33456789,
			},
		}, 1, nil
	}

	if reflect.DeepEqual(labels, []string{"preparing"}) {
		return []*Article{
			&Article{
				ID:              33456711,
				AuthorID:        1234568,
				CommentsDisable: false,
				Draft:           false,
				Promoted:        false,
				Position:        0,
				VoteSum:         0,
				VoteCount:       0,
				CreatedAt:       FixCreatedAt1,
				UpdatedAt:       FixUpdatedAt1,
				SourceLocale:    "en-us",
				Outdated:        false,
				OutdatedLocales: []string{},
				EditedAt:        FixEditedAt1,
				LabelNames:      []string{"preparing"},
				CountryCode:     "tw",
				URL:             "www.honestbee.com",
				HTMLURL:         "www.honestbee.com",
				Name:            "testing article 2",
				Title:           "testing article 2",
				Body:            "this is testing article 2",
				Locale:          "en-us",
				SectionID:       33456789,
			},
		}, 1, nil
	}

	if reflect.DeepEqual(labels, []string{"ontheway"}) {
		return []*Article{
			&Article{
				ID:              33456712,
				AuthorID:        1234569,
				CommentsDisable: false,
				Draft:           false,
				Promoted:        false,
				Position:        0,
				VoteSum:         0,
				VoteCount:       0,
				CreatedAt:       FixCreatedAt1,
				UpdatedAt:       FixUpdatedAt1,
				SourceLocale:    "en-us",
				Outdated:        false,
				OutdatedLocales: []string{},
				EditedAt:        FixEditedAt1,
				LabelNames:      []string{"ontheway"},
				CountryCode:     "tw",
				URL:             "www.honestbee.com",
				HTMLURL:         "www.honestbee.com",
				Name:            "testing article 3",
				Title:           "testing article 3",
				Body:            "this is testing article 3",
				Locale:          "en-us",
				SectionID:       33456789,
			},
		}, 1, nil
	}

	if reflect.DeepEqual(labels, []string{"delivered"}) {
		return []*Article{
			&Article{
				ID:              33456712,
				AuthorID:        1234569,
				CommentsDisable: false,
				Draft:           false,
				Promoted:        false,
				Position:        0,
				VoteSum:         0,
				VoteCount:       0,
				CreatedAt:       FixCreatedAt1,
				UpdatedAt:       FixUpdatedAt1,
				SourceLocale:    "en-us",
				Outdated:        false,
				OutdatedLocales: []string{},
				EditedAt:        FixEditedAt1,
				LabelNames:      []string{"delivered"},
				CountryCode:     "tw",
				URL:             "www.honestbee.com",
				HTMLURL:         "www.honestbee.com",
				Name:            "testing article 4",
				Title:           "testing article 4",
				Body:            "this is testing article 4",
				Locale:          "en-us",
				SectionID:       33456789,
			},
		}, 1, nil
	}

	return nil, 0, nil
}

// GetArticlesBySectionID is the mock function of GetArticlesBySectionID.
func (m *MockModels) GetArticlesBySectionID(ctx context.Context, params *GetArticlesParams) ([]*Article, int, error) {
	switch params.CountryCode {
	case ModelsReturnErrorCountryCode:
		return nil, 0, errors.New("MockModels GetArticlesBySectionID return error")
	case ModelsReturnNotFoundCountryCode:
		return nil, 0, ErrNotFound
	}

	return []*Article{
		&Article{
			ID:              33456710,
			AuthorID:        1234567,
			CommentsDisable: false,
			Draft:           false,
			Promoted:        false,
			Position:        0,
			VoteSum:         0,
			VoteCount:       0,
			CreatedAt:       FixCreatedAt1,
			UpdatedAt:       FixUpdatedAt1,
			SourceLocale:    "en-us",
			Outdated:        false,
			OutdatedLocales: []string{},
			EditedAt:        FixEditedAt1,
			LabelNames:      []string{},
			CountryCode:     "tw",
			URL:             "www.honestbee.com",
			HTMLURL:         "www.honestbee.com",
			Name:            "testing article 1",
			Title:           "testing article 1",
			Body:            "this is testing article 1",
			Locale:          "en-us",
			SectionID:       33456789,
		},
	}, 1, nil
}

// GetArticleByArticleID is the mock function of GetArticleByArticleID.
func (m *MockModels) GetArticleByArticleID(ctx context.Context, articleID int, locale, countryCode string) (*Article, error) {
	switch countryCode {
	case ModelsReturnErrorCountryCode:
		return nil, errors.New("MockModels GetArticleByArticleID return error")
	case ModelsReturnNotFoundCountryCode:
		return nil, ErrNotFound
	}

	return &Article{
		ID:              33456710,
		AuthorID:        1234567,
		CommentsDisable: false,
		Draft:           false,
		Promoted:        false,
		Position:        0,
		VoteSum:         0,
		VoteCount:       0,
		CreatedAt:       FixCreatedAt1,
		UpdatedAt:       FixUpdatedAt1,
		SourceLocale:    "en-us",
		Outdated:        false,
		OutdatedLocales: []string{},
		EditedAt:        FixEditedAt1,
		LabelNames:      []string{},
		CountryCode:     "tw",
		URL:             "www.honestbee.com",
		HTMLURL:         "www.honestbee.com",
		Name:            "testing article 1",
		Title:           "testing article 1",
		Body:            "this is testing article 1",
		Locale:          "en-us",
		SectionID:       33456789,
	}, nil
}

// PlusOneArticleClickCounter is the mock function of PlusOneArticleClickCounter.
func (m *MockModels) PlusOneArticleClickCounter(ctx context.Context, articleID int, locale, countryCode string) error {
	switch countryCode {
	case ModelsReturnErrorCountryCode:
		return errors.New("MockModels PlusOneArticleClickCounter return error")
	case ModelsReturnNotFoundCountryCode:
		return ErrNotFound
	}

	return nil
}

// GetTopNArticles is the mock function of GetTopNArticles.
func (m *MockModels) GetTopNArticles(ctx context.Context, topN uint64, locale, countryCode string) ([]*Article, error) {
	switch countryCode {
	case ModelsReturnErrorCountryCode:
		return nil, errors.New("MockModels GetTopNArticles return error")
	case ModelsReturnNotFoundCountryCode:
		return nil, ErrNotFound
	}

	switch topN {
	case 4:
		return []*Article{
			&Article{
				ID:              33456710,
				AuthorID:        1234567,
				CommentsDisable: false,
				Draft:           false,
				Promoted:        false,
				Position:        0,
				VoteSum:         0,
				VoteCount:       0,
				CreatedAt:       FixCreatedAt1,
				UpdatedAt:       FixUpdatedAt1,
				SourceLocale:    "en-us",
				Outdated:        false,
				OutdatedLocales: []string{},
				EditedAt:        FixEditedAt1,
				LabelNames:      []string{},
				CountryCode:     "tw",
				URL:             "www.honestbee.com",
				HTMLURL:         "www.honestbee.com",
				Name:            "testing article 1",
				Title:           "testing article 1",
				Body:            "this is testing article 1",
				Locale:          "en-us",
				SectionID:       33456789,
			},
			&Article{
				ID:              33456711,
				AuthorID:        1234568,
				CommentsDisable: false,
				Draft:           false,
				Promoted:        false,
				Position:        0,
				VoteSum:         0,
				VoteCount:       0,
				CreatedAt:       FixCreatedAt1,
				UpdatedAt:       FixUpdatedAt1,
				SourceLocale:    "en-us",
				Outdated:        false,
				OutdatedLocales: []string{},
				EditedAt:        FixEditedAt1,
				LabelNames:      []string{},
				CountryCode:     "tw",
				URL:             "www.honestbee.com",
				HTMLURL:         "www.honestbee.com",
				Name:            "testing article 2",
				Title:           "testing article 2",
				Body:            "this is testing article 2",
				Locale:          "en-us",
				SectionID:       33456789,
			},
			&Article{
				ID:              33456712,
				AuthorID:        1234569,
				CommentsDisable: false,
				Draft:           false,
				Promoted:        false,
				Position:        0,
				VoteSum:         0,
				VoteCount:       0,
				CreatedAt:       FixCreatedAt1,
				UpdatedAt:       FixUpdatedAt1,
				SourceLocale:    "en-us",
				Outdated:        false,
				OutdatedLocales: []string{},
				EditedAt:        FixEditedAt1,
				LabelNames:      []string{},
				CountryCode:     "tw",
				URL:             "www.honestbee.com",
				HTMLURL:         "www.honestbee.com",
				Name:            "testing article 3",
				Title:           "testing article 3",
				Body:            "this is testing article 3",
				Locale:          "en-us",
				SectionID:       33456789,
			},
			&Article{
				ID:              33456713,
				AuthorID:        1234570,
				CommentsDisable: false,
				Draft:           false,
				Promoted:        false,
				Position:        0,
				VoteSum:         0,
				VoteCount:       0,
				CreatedAt:       FixCreatedAt1,
				UpdatedAt:       FixUpdatedAt1,
				SourceLocale:    "en-us",
				Outdated:        false,
				OutdatedLocales: []string{},
				EditedAt:        FixEditedAt1,
				LabelNames:      []string{},
				CountryCode:     "tw",
				URL:             "www.honestbee.com",
				HTMLURL:         "www.honestbee.com",
				Name:            "testing article 4",
				Title:           "testing article 4",
				Body:            "this is testing article 4",
				Locale:          "en-us",
				SectionID:       33456789,
			},
		}, nil
	case 5:
		return []*Article{
			&Article{
				ID:              33456710,
				AuthorID:        1234567,
				CommentsDisable: false,
				Draft:           false,
				Promoted:        false,
				Position:        0,
				VoteSum:         0,
				VoteCount:       0,
				CreatedAt:       FixCreatedAt1,
				UpdatedAt:       FixUpdatedAt1,
				SourceLocale:    "en-us",
				Outdated:        false,
				OutdatedLocales: []string{},
				EditedAt:        FixEditedAt1,
				LabelNames:      []string{},
				CountryCode:     "tw",
				URL:             "www.honestbee.com",
				HTMLURL:         "www.honestbee.com",
				Name:            "testing article 1",
				Title:           "testing article 1",
				Body:            "this is testing article 1",
				Locale:          "en-us",
				SectionID:       33456789,
			},
			&Article{
				ID:              33456711,
				AuthorID:        1234568,
				CommentsDisable: false,
				Draft:           false,
				Promoted:        false,
				Position:        0,
				VoteSum:         0,
				VoteCount:       0,
				CreatedAt:       FixCreatedAt1,
				UpdatedAt:       FixUpdatedAt1,
				SourceLocale:    "en-us",
				Outdated:        false,
				OutdatedLocales: []string{},
				EditedAt:        FixEditedAt1,
				LabelNames:      []string{},
				CountryCode:     "tw",
				URL:             "www.honestbee.com",
				HTMLURL:         "www.honestbee.com",
				Name:            "testing article 2",
				Title:           "testing article 2",
				Body:            "this is testing article 2",
				Locale:          "en-us",
				SectionID:       33456789,
			},
			&Article{
				ID:              33456712,
				AuthorID:        1234569,
				CommentsDisable: false,
				Draft:           false,
				Promoted:        false,
				Position:        0,
				VoteSum:         0,
				VoteCount:       0,
				CreatedAt:       FixCreatedAt1,
				UpdatedAt:       FixUpdatedAt1,
				SourceLocale:    "en-us",
				Outdated:        false,
				OutdatedLocales: []string{},
				EditedAt:        FixEditedAt1,
				LabelNames:      []string{},
				CountryCode:     "tw",
				URL:             "www.honestbee.com",
				HTMLURL:         "www.honestbee.com",
				Name:            "testing article 3",
				Title:           "testing article 3",
				Body:            "this is testing article 3",
				Locale:          "en-us",
				SectionID:       33456789,
			},
			&Article{
				ID:              33456713,
				AuthorID:        1234570,
				CommentsDisable: false,
				Draft:           false,
				Promoted:        false,
				Position:        0,
				VoteSum:         0,
				VoteCount:       0,
				CreatedAt:       FixCreatedAt1,
				UpdatedAt:       FixUpdatedAt1,
				SourceLocale:    "en-us",
				Outdated:        false,
				OutdatedLocales: []string{},
				EditedAt:        FixEditedAt1,
				LabelNames:      []string{},
				CountryCode:     "tw",
				URL:             "www.honestbee.com",
				HTMLURL:         "www.honestbee.com",
				Name:            "testing article 4",
				Title:           "testing article 4",
				Body:            "this is testing article 4",
				Locale:          "en-us",
				SectionID:       33456789,
			},
			&Article{
				ID:              33456714,
				AuthorID:        1234571,
				CommentsDisable: false,
				Draft:           false,
				Promoted:        false,
				Position:        0,
				VoteSum:         0,
				VoteCount:       0,
				CreatedAt:       FixCreatedAt1,
				UpdatedAt:       FixUpdatedAt1,
				SourceLocale:    "en-us",
				Outdated:        false,
				OutdatedLocales: []string{},
				EditedAt:        FixEditedAt1,
				LabelNames:      []string{},
				CountryCode:     "tw",
				URL:             "www.honestbee.com",
				HTMLURL:         "www.honestbee.com",
				Name:            "testing article 5",
				Title:           "testing article 5",
				Body:            "this is testing article 5",
				Locale:          "en-us",
				SectionID:       33456789,
			},
		}, nil
	}

	return []*Article{}, nil
}

// GetTicketForm is the mock function of GetTicketForm.
func (m *MockModels) GetTicketForm(ctx context.Context, formID int, locale string) (*TicketForm, error) {
	switch locale {
	case ModelsReturnErrorLocale:
		return nil, errors.New("MockModels GetTicketForm return error")
	case ModelsReturnNotFoundLocale:
		return nil, ErrNotFound
	}

	return &TicketForm{
		ID:             191908,
		Name:           "191908 - Default Ticket Form",
		RawName:        "191908 - Default Ticket Form",
		DisplayName:    "Default Ticket Form",
		RawDisplayName: "Default Ticket Form",
		Position:       0,
		TicketFields:   make([]*TicketField, 0),
		CreatedAt:      FixCreatedAt1,
		UpdatedAt:      FixUpdatedAt1,
	}, nil
}

// GetTicketFormGraphQL is the mock function of GetTicketFormGraphQL.
func (m *MockModels) GetTicketFormGraphQL(ctx context.Context, formID int) (*SyncTicketForm, error) {
	switch formID {
	case -1:
		return nil, errors.New("MockModels GetTicketFormGraphQL return error")
	case 191908:
		return &SyncTicketForm{
			ID:                 191908,
			URL:                "https://honestbeehelp-tw.zendesk.com/api/v2/ticket_forms/191908.json",
			Name:               "191908 - Default Ticket Form",
			RawName:            "191908 - Default Ticket Form",
			DisplayName:        "Default Ticket Form",
			RawDisplayName:     "Default Ticket Form",
			EndUserVisible:     true,
			Position:           0,
			Active:             true,
			InAllBrands:        true,
			RestrictedBrandIDs: []int64{123, 456, 789},
			TicketFieldIDs:     []int64{24681488, 24681498},
			CreatedAt:          FixCreatedAt1,
			UpdatedAt:          FixUpdatedAt1,
		}, nil
	default:
		return nil, ErrNotFound
	}
}

// GetTicketFieldByFieldID is the mock function of GetTicketFieldByFieldID.
func (m *MockModels) GetTicketFieldByFieldID(ctx context.Context, id int, locale string) (*TicketField, error) {
	return &TicketField{
		ID:                  81469808,
		Type:                "text",
		Title:               "Order Number",
		RawTitle:            "Order Number",
		Description:         "",
		RawDescription:      "",
		Position:            13,
		RegexpForValidation: "",
		TitleInPortal:       "Order Number",
		RawTitleInPortal:    "訂單號碼",
		CreatedAt:           FixCreatedAt1,
		UpdatedAt:           FixUpdatedAt1,
		CustomFieldOptions:  make([]*CustomFieldOption, 0),
	}, nil
}

// GetTicketFieldByFormID is the mock function of GetTicketFieldByFormID.
func (m *MockModels) GetTicketFieldByFormID(ctx context.Context, formID int, locale string) ([]*TicketField, error) {
	switch locale {
	case ModelsReturnErrorLocale:
		return nil, errors.New("MockModels GetTicketFieldByFormID return error")
	case ModelsReturnNotFoundLocale:
		return nil, ErrNotFound
	}

	if formID == 191908 {
		switch locale {
		case "en-us":
			return []*TicketField{
				&TicketField{
					ID:                  81469808,
					URL:                 "https://honestbee-ph.zendesk.com/api/v2/ticket_fields/81469808.json",
					Type:                "text",
					Title:               "Order Number",
					RawTitle:            "Order Number",
					Description:         "",
					RawDescription:      "",
					Position:            13,
					Active:              true,
					Required:            false,
					CollapsedForAgents:  false,
					RegexpForValidation: "",
					TitleInPortal:       "Order Number",
					RawTitleInPortal:    "Order Number",
					VisibleInPortal:     true,
					EditableInPortal:    true,
					RequiredInPortal:    true,
					Tag:                 "",
					CreatedAt:           FixCreatedAt1,
					UpdatedAt:           FixUpdatedAt1,
					Removable:           true,
					CustomFieldOptions:  make([]*CustomFieldOption, 0),
					SystemFieldOptions:  make([]*SystemFieldOption, 0),
				},
			}, nil
		case "zh-tw":
			return []*TicketField{
				&TicketField{
					ID:                  81469808,
					URL:                 "https://honestbee-ph.zendesk.com/api/v2/ticket_fields/81469808.json",
					Type:                "text",
					Title:               "Order Number",
					RawTitle:            "Order Number",
					Description:         "",
					RawDescription:      "",
					Position:            13,
					Active:              true,
					Required:            false,
					CollapsedForAgents:  false,
					RegexpForValidation: "",
					TitleInPortal:       "Order Number",
					RawTitleInPortal:    "訂單號碼",
					VisibleInPortal:     true,
					EditableInPortal:    true,
					RequiredInPortal:    true,
					Tag:                 "",
					CreatedAt:           FixCreatedAt1,
					UpdatedAt:           FixUpdatedAt1,
					Removable:           true,
					CustomFieldOptions:  make([]*CustomFieldOption, 0),
					SystemFieldOptions:  make([]*SystemFieldOption, 0),
				},
			}, nil
		}
	}
	if formID == 123456 {
		switch locale {
		case "en-us":
			return []*TicketField{
				&TicketField{
					ID: 123456789,
				},
			}, nil
		case "zh-tw":
			return []*TicketField{
				&TicketField{
					ID: 987654321,
				},
			}, nil
		}
	}
	if formID == 654321 {
		switch locale {
		case "en-us":
			return []*TicketField{
				&TicketField{
					ID: 987654,
				},
			}, nil
		case "zh-tw":
			return []*TicketField{
				&TicketField{
					ID: 456789,
				},
			}, nil
		}
	}

	return nil, nil
}

// GetTicketFieldCustomFieldOption is the mock function of GetTicketFieldCustomFieldOption.
func (m *MockModels) GetTicketFieldCustomFieldOption(ctx context.Context, fieldID int) ([]*CustomFieldOption, error) {
	switch fieldID {
	case 123456789:
		return nil, errors.New("MockModels GetTicketFieldCustomFieldOption return error")
	case 987654321:
		return nil, ErrNotFound
	case 81469808:
		return []*CustomFieldOption{
			&CustomFieldOption{
				ID:      83592448,
				Name:    "Grocery",
				RawName: "Grocery",
				Value:   "grocery_form",
			},
			&CustomFieldOption{
				ID:      83592468,
				Name:    "Food",
				RawName: "Food",
				Value:   "food_form",
			},
			&CustomFieldOption{
				ID:      83592488,
				Name:    "Laundry",
				RawName: "Laundry",
				Value:   "laundry_form",
			},
			&CustomFieldOption{
				ID:      83592508,
				Name:    "Ticketing",
				RawName: "Ticketing",
				Value:   "ticketing_form",
			},
			&CustomFieldOption{
				ID:      83592528,
				Name:    "Rewards",
				RawName: "Rewards",
				Value:   "rewards_form",
			},
		}, nil
	default:
		return nil, nil
	}
}

// GetTicketFieldSystemFieldOption is the mock function of GetTicketFieldSystemFieldOption.
func (m *MockModels) GetTicketFieldSystemFieldOption(ctx context.Context, fieldID int) ([]*SystemFieldOption, error) {
	switch fieldID {
	case 987654:
		return nil, errors.New("MockModels GetTicketFieldSystemFieldOption return error")
	case 456789:
		return nil, ErrNotFound
	case 81469808:
		return []*SystemFieldOption{
			&SystemFieldOption{},
		}, nil
	default:
		return nil, nil
	}
}

// GetDynamicContentItem is the mock function of GetDynamicContentItem.
func (m *MockModels) GetDynamicContentItem(ctx context.Context, placeholder, locale string) (*DynamicContentItem, error) {
	return &DynamicContentItem{
		ID:                754768,
		Name:              "Form: Order Number field",
		Placeholder:       "{{dc.form_order_number_field}}",
		DefaultLocaleID:   1,
		CreatedAt:         FixCreatedAt1,
		UpdatedAt:         FixUpdatedAt1,
		VariantsID:        3392648,
		VariantsContent:   "訂單號碼",
		VariantsLocaleID:  9,
		VariantsCreatedAt: FixCreatedAt1,
		VariantsUpdatedAt: FixUpdatedAt1,
	}, nil
}
