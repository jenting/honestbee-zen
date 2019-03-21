package zendesk

import (
	"time"
)

// BaseOut is the zendesk api basic params return form.
type BaseOut struct {
	Page         int     `json:"page,omitempty"`
	PerPage      int     `json:"per_page,omitempty"`
	PageCount    int     `json:"page_count,omitempty"`
	Count        int     `json:"count,omitempty"`
	NextPage     *string `json:"next_page,omitempty"`
	PreviousPage *string `json:"previous_page,omitempty"`
}

// Category is the zendesk category form.
type Category struct {
	ID           int       `json:"id,omitempty"`
	Position     int       `json:"position,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	SourceLocale string    `json:"source_locale,omitempty"`
	Outdated     bool      `json:"outdated,omitempty"`
	URL          string    `json:"url,omitempty"`
	HTMLURL      string    `json:"html_url,omitempty"`
	Name         string    `json:"name,omitempty"`
	Description  string    `json:"description,omitempty"`
	Locale       string    `json:"locale,omitempty"`
}

// ListCategories is the zendesk api:
// GET /api/v2/help_center/{locale}/categories.json
// return form.
type ListCategories struct {
	Categories []*Category `json:"categories,omitempty"`
	*BaseOut
}

// Section is the zendesk section form.
type Section struct {
	CategoryID   int       `json:"category_id,omitempty"`
	ID           int       `json:"id,omitempty"`
	Position     int       `json:"position,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	SourceLocale string    `json:"source_locale,omitempty"`
	Outdated     bool      `json:"outdated,omitempty"`
	URL          string    `json:"url,omitempty"`
	HTMLURL      string    `json:"html_url,omitempty"`
	Name         string    `json:"name,omitempty"`
	Description  string    `json:"description,omitempty"`
	Locale       string    `json:"locale,omitempty"`
}

// ListSections is the zendesk api:
// GET /api/v2/help_center/{locale}/sections.json
// return form.
type ListSections struct {
	Sections []*Section `json:"sections,omitempty"`
	*BaseOut
}

// Article is the zendesk article form.
type Article struct {
	SectionID       int       `json:"section_id,omitempty"`
	ID              int       `json:"id,omitempty"`
	AuthorID        int       `json:"author_id,omitempty"`
	CommentsDisable bool      `json:"comments_disable,omitempty"`
	Draft           bool      `json:"draft,omitempty"`
	Promoted        bool      `json:"promoted,omitempty"`
	Position        int       `json:"position,omitempty"`
	VoteSum         int       `json:"vote_sum,omitempty"`
	VoteCount       int       `json:"vote_count,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	SourceLocale    string    `json:"source_locale,omitempty"`
	Outdated        bool      `json:"outdated,omitempty"`
	OutdatedLocales []string  `json:"outdated_locales,omitempty"`
	EditedAt        time.Time `json:"edited_at,omitempty"`
	LabelNames      []string  `json:"label_names,omitempty"`
	URL             string    `json:"url,omitempty"`
	HTMLURL         string    `json:"html_url,omitempty"`
	Name            string    `json:"name,omitempty"`
	Title           string    `json:"title,omitempty"`
	Body            string    `json:"body,omitempty"`
	Locale          string    `json:"locale,omitempty"`
}

// ListArticles is the zendesk api:
// /api/v2/help_center/{locale}/articles.json
// return form.
type ListArticles struct {
	Articles []*Article `json:"articles,omitempty"`
	*BaseOut
}

// ShowArticle is the zendesk api:
// /api/v2/help_center/{locale}/articles/{id}.json
// return form.
type ShowArticle struct {
	Article *Article `json:"article,omitempty"`
}

// Vote is the zendesk POST
// hc/{locale}/articles/{id}/vote
// return form.
type Vote struct {
	ID          int64  `json:"id"`
	VoteSum     int    `json:"vote_sum"`
	VoteCount   int    `json:"vote_count"`
	UpvoteCount int    `json:"upvote_count"`
	Label       string `json:"label"`
	Value       string `json:"value"`
}

// TicketForm is the zendesk ticket form format.
type TicketForm struct {
	ID                 int       `json:"id,omitempty"`
	URL                string    `json:"url,omitempty"`
	Name               string    `json:"name,omitempty"`
	RawName            string    `json:"raw_name,omitempty"`
	DisplayName        string    `json:"display_name,omitempty"`
	RawDisplayName     string    `json:"raw_display_name,omitempty"`
	EndUserVisible     bool      `json:"end_user_visible,omitempty"`
	Position           int       `json:"position,omitempty"`
	Active             bool      `json:"active,omitempty"`
	InAllBrands        bool      `json:"in_all_brands,omitempty"`
	RestrictedBrandIDs []int64   `json:"restricted_brand_ids,omitempty"`
	TicketFieldIDs     []int64   `json:"ticket_field_ids,omitempty"`
	CreatedAt          time.Time `json:"created_at,omitempty"`
	UpdatedAt          time.Time `json:"updated_at,omitempty"`
}

// ListTicketForms is the zendesk api:
// /api/v2/ticket_forms.json
// return from.
type ListTicketForms struct {
	TicketForms []*TicketForm `json:"ticket_forms,omitempty"`
	*BaseOut
}

// TicketField is the zendesk ticket field format.
type TicketField struct {
	ID                  int                  `json:"id,omitempty"`
	URL                 string               `json:"url,omitempty"`
	Type                string               `json:"type,omitempty"`
	Title               string               `json:"title,omitempty"`
	RawTitle            string               `json:"raw_title,omitempty"`
	Description         string               `json:"description,omitempty"`
	RawDescription      string               `json:"raw_description,omitempty"`
	Position            int                  `json:"position,omitempty"`
	Active              bool                 `json:"active,omitempty"`
	Required            bool                 `json:"required,omitempty"`
	CollapsedForAgents  bool                 `json:"collapsed_for_agents,omitempty"`
	RegexpForValidation string               `json:"regexp_for_validation,omitempty"`
	TitleInPortal       string               `json:"title_in_portal,omitempty"`
	RawTitleInPortal    string               `json:"raw_title_in_portal,omitempty"`
	VisibleInPortal     bool                 `json:"visible_in_portal,omitempty"`
	EditableInPortal    bool                 `json:"editable_in_portal,omitempty"`
	RequiredInPortal    bool                 `json:"required_in_portal,omitempty"`
	Tag                 string               `json:"tag,omitempty"`
	CreatedAt           time.Time            `json:"created_at,omitempty"`
	UpdatedAt           time.Time            `json:"updated_at,omitempty"`
	Removable           bool                 `json:"removable,omitempty"`
	CustomFieldOptions  []*CustomFieldOption `json:"custom_field_options,omitempty"`
	SystemFieldOptions  []*SystemFieldOption `json:"system_field_options,omitempty"`
}

// CustomFieldOption is the TicketField struct CustomFieldOptions slice unit.
type CustomFieldOption struct {
	ID      int    `db:"id" json:"id,omitempty"`
	Name    string `db:"name" json:"name,omitempty"`
	RawName string `db:"raw_name" json:"raw_name,omitempty"`
	Value   string `db:"value" json:"value,omitempty"`
}

// SystemFieldOption is the TicketField struct SystemFieldOptions slice unit.
type SystemFieldOption struct {
	Name  string `db:"name" json:"name,omitempty"`
	Value string `db:"value" json:"value,omitempty"`
}

// ListTicketFields is the zendesk api:
// /api/v2/ticket_fields.json
// return form.
type ListTicketFields struct {
	TicketFields []*TicketField `json:"ticket_fields,omitempty"`
	*BaseOut
}

// DynamicContentItem is the zendesk dynamic content item form
type DynamicContentItem struct {
	ID              int        `json:"id,omitempty"`
	URL             string     `json:"url,omitempty"`
	Name            string     `json:"name,omitempty"`
	Placeholder     string     `json:"placeholder,omitempty"`
	DefaultLocaleID int        `json:"default_locale_id,omitempty"`
	Outdated        bool       `json:"outdated,omitempty"`
	CreatedAt       time.Time  `json:"created_at,omitempty"`
	UpdatedAt       time.Time  `json:"updated_at,omitempty"`
	Variants        []*Variant `json:"variants,omitempty"`
}

// Variant is the DynamicContentItem struct Variants slice unit.
type Variant struct {
	ID        int       `json:"id,omitempty"`
	URL       string    `json:"url,omitempty"`
	Content   string    `json:"content,omitempty"`
	LocaleID  int       `json:"locale_id,omitempty"`
	Outdated  bool      `json:"outdated,omitempty"`
	Active    bool      `json:"active,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// ListDynamicContentItems is the zendesk api:
// /api/v2/dynamic_content/items.json
// return form.
type ListDynamicContentItems struct {
	Items []*DynamicContentItem `json:"items,omitempty"`
	*BaseOut
}

// InstantSearch is the zendesk internal hc api:
// /hc/api/internal/instant_search.json
// return form.
type InstantSearch struct {
	Results []*InstantSearchResult `json:"results"`
}

// InstantSearchResult is the zendesk InstantSearch result form
type InstantSearchResult struct {
	Title         string `json:"title"`
	CategoryTitle string `json:"category_title"`
	URL           string `json:"url"`
}

// SearchArticle is the zendesk article which contains two additional properties for search.
type SearchArticle struct {
	*Article
	Snippet    string `json:"snippet"`
	ResultType string `json:"result_type"`
}

// Search is the zendesk api:
// /api/v2/help_center/articles/search.json
// return form.
type Search struct {
	Articles []*SearchArticle `json:"results"`
	*BaseOut
}
