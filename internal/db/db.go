package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx/types"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

var (
	// ErrNoRows means the query got sql.ErrNoRows error
	ErrNoRows = errors.New("no such rows")
)

// Database is the interface of defining all normal operations.
type Database interface {
	Close() error
	Select(ctx context.Context, dest interface{}, query string) error
	Get(ctx context.Context, dest interface{}, query string) error
	NamedExec(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	Begin() (DatabaseTransaction, error)
}

// DatabaseTransaction is the interface of defining all transaction operations.
// All operations will doing rollback if there has an error occurred.
type DatabaseTransaction interface {
	Select(dest interface{}, query string)
	Get(dest interface{}, query string)
	NamedExec(query string, arg interface{}) sql.Result
	Err() error
	Commit()
}

// Categories is the categories table columns.
type Categories struct {
	SN           int       `db:"sn"`
	ID           int       `db:"id"`
	Position     int       `db:"position"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	SourceLocale string    `db:"source_locale"`
	Outdated     bool      `db:"outdated"`
	CountryCode  string    `db:"country_code"`
}

// Category is the categories join categoryTranslates table columns.
type Category struct {
	ID           int       `db:"id"`
	Position     int       `db:"position"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	SourceLocale string    `db:"source_locale"`
	Outdated     bool      `db:"outdated"`
	CountryCode  string    `db:"country_code"`

	CategoryTranslates
}

// CategoryTranslates is the category_translates table columns.
type CategoryTranslates struct {
	SN          int    `db:"sn"`
	CategoryID  int    `db:"category_id"`
	URL         string `db:"url"`
	HTMLURL     string `db:"html_url"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Locale      string `db:"locale"`
}

// CategoryKey is the category_key table columns.
type CategoryKey struct {
	SN          int       `db:"sn"`
	CategoryID  int       `db:"category_id"`
	KeyName     string    `db:"key_name"`
	CountryCode string    `db:"country_code"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

// Sections is the sections table columns.
type Sections struct {
	SN           int       `db:"sn"`
	CategoryID   int       `db:"category_id"`
	ID           int       `db:"id"`
	Position     int       `db:"position"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	SourceLocale string    `db:"source_locale"`
	Outdated     bool      `db:"outdated"`
	CountryCode  string    `db:"country_code"`
}

// SectionTranslates is the section_translates table columns.
type SectionTranslates struct {
	SN          int    `db:"sn"`
	SectionID   int    `db:"section_id"`
	URL         string `db:"url"`
	HTMLURL     string `db:"html_url"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Locale      string `db:"locale"`
}

// Articles is the articles table columns.
type Articles struct {
	SN              int            `db:"sn"`
	SectionID       int            `db:"section_id"`
	ID              int            `db:"id"`
	AuthorID        int            `db:"author_id"`
	CommentsDisable bool           `db:"comments_disable"`
	Draft           bool           `db:"draft"`
	Promoted        bool           `db:"promoted"`
	Position        int            `db:"position"`
	VoteSum         int            `db:"vote_sum"`
	VoteCount       int            `db:"vote_count"`
	CreatedAt       time.Time      `db:"created_at"`
	UpdatedAt       time.Time      `db:"updated_at"`
	SourceLocale    string         `db:"source_locale"`
	Outdated        bool           `db:"outdated"`
	OutdatedLocales pq.StringArray `db:"outdated_locales"`
	EditedAt        time.Time      `db:"edited_at"`
	LabelNames      pq.StringArray `db:"label_names"`
	CountryCode     string         `db:"country_code"`
}

// ArticleTranslates is the article_translates table columns.
type ArticleTranslates struct {
	SN        int    `db:"sn"`
	ArticleID int    `db:"article_id"`
	URL       string `db:"url"`
	HTMLURL   string `db:"html_url"`
	Name      string `db:"name"`
	Title     string `db:"title"`
	Body      string `db:"body"`
	Locale    string `db:"locale"`
}

// TicketForms is the ticket_forms table columns.
type TicketForms struct {
	SN                 int           `db:"sn"`
	ID                 int           `db:"id"`
	URL                string        `db:"url"`
	Name               string        `db:"name"`
	RawName            string        `db:"raw_name"`
	DisplayName        string        `db:"display_name"`
	RawDisplayName     string        `db:"raw_display_name"`
	EndUserVisible     bool          `db:"end_user_visible"`
	Position           int           `db:"position"`
	Active             bool          `db:"active"`
	InAllBrands        bool          `db:"in_all_brands"`
	RestrictedBrandIDs pq.Int64Array `db:"restricted_brand_ids"`
	TicketFieldIDs     pq.Int64Array `db:"ticket_field_ids"`
	CreatedAt          time.Time     `db:"created_at"`
	UpdatedAt          time.Time     `db:"updated_at"`
}

// TicketFields is the ticket fields table columns.
type TicketFields struct {
	SN                  int            `db:"sn"`
	ID                  int            `db:"id"`
	URL                 string         `db:"url"`
	Type                string         `db:"type"`
	Title               string         `db:"title"`
	RawTitle            string         `db:"raw_title"`
	Description         string         `db:"description"`
	RawDescription      string         `db:"raw_description"`
	Position            int            `db:"position"`
	Active              bool           `db:"active"`
	Required            bool           `db:"required"`
	CollapsedForAgents  bool           `db:"collapsed_for_agents"`
	RegexpForValidation string         `db:"regexp_for_validation"`
	TitleInPortal       string         `db:"title_in_portal"`
	RawTitleInPortal    string         `db:"raw_title_in_portal"`
	VisibleInPortal     bool           `db:"visible_in_portal"`
	EditableInPortal    bool           `db:"editable_in_portal"`
	RequiredInPortal    bool           `db:"required_in_portal"`
	Tag                 string         `db:"tag"`
	CreatedAt           time.Time      `db:"created_at"`
	UpdatedAt           time.Time      `db:"updated_at"`
	Removable           bool           `db:"removable"`
	CustomFieldOptions  types.JSONText `db:"custom_field_options"`
	SystemFieldOptions  types.JSONText `db:"system_field_options"`
}

// CustomFieldOption is the ticket fields table CustomFieldOptions column json form.
type CustomFieldOption struct {
	ID      int    `db:"id" json:"id,omitempty"`
	Name    string `db:"name" json:"name,omitempty"`
	RawName string `db:"raw_name" json:"raw_name,omitempty"`
	Value   string `db:"value" json:"value,omitempty"`
}

// SystemFieldOption is the ticket fields table SystemFieldOptions column json form.
type SystemFieldOption struct {
	Name  string `db:"name" json:"name,omitempty"`
	Value string `db:"value" json:"value,omitempty"`
}

// DynamicContentItems is the dynamic_content_items table columns.
type DynamicContentItems struct {
	SN              int            `db:"sn"`
	ID              int            `db:"id"`
	URL             string         `db:"url"`
	Name            string         `db:"name"`
	Placeholder     string         `db:"placeholder"`
	DefaultLocaleID int            `db:"default_locale_id"`
	Outdated        bool           `db:"outdated"`
	CreatedAt       time.Time      `db:"created_at"`
	UpdatedAt       time.Time      `db:"updated_at"`
	Variants        types.JSONText `db:"variants"`
}

// Variant is the dynamic_content_items table variants column json form.
type Variant struct {
	ID        int       `db:"id" json:"id,omitempty"`
	URL       string    `db:"url" json:"url,omitempty"`
	Content   string    `db:"content" json:"content,omitempty"`
	LocaleID  int       `db:"locale_id" json:"locale_id,omitempty"`
	Outdated  bool      `db:"outdated" json:"outdated,omitempty"`
	Active    bool      `db:"active" json:"active,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at,omitempty"`
}
