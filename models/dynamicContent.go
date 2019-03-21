package models

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx/types"
	"github.com/pkg/errors"

	"github.com/honestbee/Zen/internal/db"
)

var (
	// based on honestbee https://{domain.name}/api/v2/locales.json
	localeMapper = map[string]int{
		"en-us": 1,
		"id":    77,
		"ja":    67,
		"zh-cn": 10,
		"zh-tw": 9,
		"th":    81,
	}
)

type dynamicContentService interface {
	GetDynamicContentItem(ctx context.Context, placeholder, locale string) (*DynamicContentItem, error)
	SyncWithDynamicContentItems(ctx context.Context, zendeskDCItems []*SyncDynamicContentItem) error
}

// DynamicContentItem is dynamic content item field model.
type DynamicContentItem struct {
	ID                int       `json:"id,omitempty"`
	Name              string    `json:"name,omitempty"`
	Placeholder       string    `json:"placeholder,omitempty"`
	DefaultLocaleID   int       `json:"default_locale_id,omitempty"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
	UpdatedAt         time.Time `json:"updated_at,omitempty"`
	VariantsID        int       `json:"variants_id,omitempty"`
	VariantsContent   string    `json:"variants_content,omitempty"`
	VariantsLocaleID  int       `json:"variants_locale_id,omitempty"`
	VariantsCreatedAt time.Time `json:"variants_created_at,omitempty"`
	VariantsUpdatedAt time.Time `json:"variants_updated_at,omitempty"`
}

// SyncDynamicContentItem is the struct for SyncWithDynamicContentItems to sync up database data.
type SyncDynamicContentItem struct {
	ID              int
	URL             string
	Name            string
	Placeholder     string
	DefaultLocaleID int
	Outdated        bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Variants        json.RawMessage
}

type dynamicContentOps struct {
	db db.Database
}

const (
	updateDCItemsQuery = `
		UPDATE dynamic_content_items SET
			url = :url,
			name = :name,
			placeholder = :placeholder,
			default_locale_id = :default_locale_id,
			outdated = :outdated,
			created_at = :created_at,
			updated_at = :updated_at,
			variants = :variants
		WHERE id = :id`

	insertDCItemsQuery = `
		INSERT INTO dynamic_content_items (
			id,
			url,
			name,
			placeholder,
			default_locale_id,
			outdated,
			created_at,
			updated_at,
			variants
		)
		VALUES (
			:id,
			:url,
			:name,
			:placeholder,
			:default_locale_id,
			:outdated,
			:created_at,
			:updated_at,
			:variants
		)`

	deleteDCItemsQuery = `DELETE FROM dynamic_content_items WHERE id = :id`
)

func (d *dynamicContentOps) SyncWithDynamicContentItems(ctx context.Context, zendeskDCItems []*SyncDynamicContentItem) error {
	tx, err := d.db.Begin()
	if err != nil {
		return errors.Wrapf(err, "models: [SyncWithDynamicContentItems] db.Begin failed")
	}

	ids := make([]int, 0)
	tx.Select(&ids, "SELECT id FROM dynamic_content_items")

	dbIDs := make(map[int]struct{})
	for _, id := range ids {
		dbIDs[id] = struct{}{}
	}

	for _, zendeskDCItem := range zendeskDCItems {
		dbDCItem := &db.DynamicContentItems{
			ID:              zendeskDCItem.ID,
			URL:             zendeskDCItem.URL,
			Name:            zendeskDCItem.Name,
			Placeholder:     zendeskDCItem.Placeholder,
			DefaultLocaleID: zendeskDCItem.DefaultLocaleID,
			Outdated:        zendeskDCItem.Outdated,
			CreatedAt:       zendeskDCItem.CreatedAt,
			UpdatedAt:       zendeskDCItem.UpdatedAt,
			Variants:        types.JSONText(zendeskDCItem.Variants),
		}

		if _, exist := dbIDs[zendeskDCItem.ID]; exist {
			tx.NamedExec(updateDCItemsQuery, dbDCItem)
			delete(dbIDs, zendeskDCItem.ID)
		} else {
			tx.NamedExec(insertDCItemsQuery, dbDCItem)
		}
	}

	for id := range dbIDs {
		tx.NamedExec(deleteDCItemsQuery, map[string]interface{}{"id": id})
	}
	tx.Commit()

	return errors.Wrapf(tx.Err(), "models: [SyncWithDynamicContentItems] db transaction failed")
}

func (d *dynamicContentOps) GetDynamicContentItem(ctx context.Context, placeholder, locale string) (*DynamicContentItem, error) {
	item := new(db.DynamicContentItems)
	query := fmt.Sprintf(
		`SELECT id,name,placeholder,default_locale_id,created_at,updated_at,variants 
		FROM dynamic_content_items WHERE placeholder = '%s'`,
		placeholder,
	)
	if err := d.db.Get(ctx, item, query); err != nil {
		switch err {
		case db.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, errors.Wrapf(
				err,
				"models: [GetDynamicContentItem] db get dc item by placeholder:%s, locale:%s failed",
				placeholder, locale,
			)
		}
	}

	ret := &DynamicContentItem{
		ID:              item.ID,
		Name:            item.Name,
		Placeholder:     item.Placeholder,
		DefaultLocaleID: item.DefaultLocaleID,
		CreatedAt:       item.CreatedAt,
		UpdatedAt:       item.UpdatedAt,
	}

	localeID := localeMapper[locale]
	variants := make([]*db.Variant, 0)
	if err := item.Variants.Unmarshal(&variants); err != nil {
		return nil, errors.Wrapf(err, "models: [GetDynamicContentItem] variants unmarshal failed")
	}

	vmaper := make(map[int]*db.Variant)
	// max loop depends on how many locale we have
	for _, variant := range variants {
		vmaper[variant.LocaleID] = variant
	}

	if v := vmaper[localeID]; v != nil && v.ID != 0 {
		ret.VariantsContent = vmaper[localeID].Content
		ret.VariantsCreatedAt = vmaper[localeID].CreatedAt
		ret.VariantsID = vmaper[localeID].ID
		ret.VariantsLocaleID = vmaper[localeID].LocaleID
		ret.VariantsUpdatedAt = vmaper[localeID].UpdatedAt

	} else if v := vmaper[ret.DefaultLocaleID]; v != nil && v.ID != 0 {
		// go default locale id
		ret.VariantsContent = vmaper[ret.DefaultLocaleID].Content
		ret.VariantsCreatedAt = vmaper[ret.DefaultLocaleID].CreatedAt
		ret.VariantsID = vmaper[ret.DefaultLocaleID].ID
		ret.VariantsLocaleID = vmaper[ret.DefaultLocaleID].LocaleID
		ret.VariantsUpdatedAt = vmaper[ret.DefaultLocaleID].UpdatedAt

	} else {
		return nil, errors.Errorf(
			"modes: [GetDynamicContentItem] placeholder:%s, locale:%s, defaultLocaleID:%d not found",
			placeholder, locale, ret.DefaultLocaleID,
		)
	}

	return ret, nil
}
