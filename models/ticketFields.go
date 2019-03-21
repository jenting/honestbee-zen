package models

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx/types"
	"github.com/pkg/errors"

	"github.com/honestbee/Zen/internal/db"
)

type ticketFieldsService interface {
	SyncWithTicketFields(ctx context.Context, zendeskTicketFields []*SyncTicketField) error
	GetTicketFieldByFieldID(ctx context.Context, fieldID int, locale string) (*TicketField, error)
	GetTicketFieldByFormID(ctx context.Context, formID int, locale string) ([]*TicketField, error)
	GetTicketFieldCustomFieldOption(ctx context.Context, fieldID int) ([]*CustomFieldOption, error)
	GetTicketFieldSystemFieldOption(ctx context.Context, fieldID int) ([]*SystemFieldOption, error)
}

// TicketField is the ticket field model.
type TicketField struct {
	ID                  int                  `json:"id,omitempty"`
	URL                 string               `json:"url,omitempty"`
	Type                string               `json:"type,omitempty"`
	Title               string               `json:"title,omitempty"`
	RawTitle            string               `json:"raw_title,omitempty"`
	Description         string               `json:"descript,omitempty"`
	RawDescription      string               `json:"raw_descript,omitempty"`
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

// CustomFieldOption is the sub-struct of TicketField.
type CustomFieldOption struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	RawName string `json:"raw_name,omitempty"`
	Value   string `json:"value,omitempty"`
}

// SystemFieldOption is the sub-struct of TicketField.
type SystemFieldOption struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// SyncTicketField is the struct for SyncWithTicketFields to sync up database data.
type SyncTicketField struct {
	ID                  int
	URL                 string
	Type                string
	Title               string
	RawTitle            string
	Description         string
	RawDescription      string
	Position            int
	Active              bool
	Required            bool
	CollapsedForAgents  bool
	RegexpForValidation string
	TitleInPortal       string
	RawTitleInPortal    string
	VisibleInPortal     bool
	EditableInPortal    bool
	RequiredInPortal    bool
	Tag                 string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	Removable           bool
	CustomFieldOptions  json.RawMessage
	SystemFieldOptions  json.RawMessage
}

type ticketFieldsOps struct {
	db    db.Database
	dcOps *dynamicContentOps
}

const (
	updateTicketFieldsQuery = `
		UPDATE ticket_fields SET
			url = :url,
			type = :type,
			title = :title,
			raw_title = :raw_title,
			description = :description,
			raw_description = :raw_description,
			position = :position,
			active = :active,
			required = :required,
			collapsed_for_agents = :collapsed_for_agents,
			regexp_for_validation = :regexp_for_validation,
			title_in_portal = :title_in_portal,
			raw_title_in_portal = :raw_title_in_portal,
			visible_in_portal = :visible_in_portal,
			editable_in_portal = :editable_in_portal,
			required_in_portal = :required_in_portal,
			tag = :tag,
			created_at = :created_at,
			updated_at = :updated_at,
			removable = :removable,
			custom_field_options = :custom_field_options,
			system_field_options = :system_field_options
		WHERE id = :id`

	insertTicketFieldsQuery = `
		INSERT INTO ticket_fields (
			id,
			url,
			type,
			title,
			raw_title,
			description,
			raw_description,
			position,
			active,
			required,
			collapsed_for_agents,
			regexp_for_validation,
			title_in_portal,
			raw_title_in_portal,
			visible_in_portal,
			editable_in_portal,
			required_in_portal,
			tag,
			created_at,
			updated_at,
			removable,
			custom_field_options,
			system_field_options
		)
		VALUES (
			:id,
			:url,
			:type,
			:title,
			:raw_title,
			:description,
			:raw_description,
			:position,
			:active,
			:required,
			:collapsed_for_agents,
			:regexp_for_validation,
			:title_in_portal,
			:raw_title_in_portal,
			:visible_in_portal,
			:editable_in_portal,
			:required_in_portal,
			:tag,
			:created_at,
			:updated_at,
			:removable,
			:custom_field_options,
			:system_field_options
		)`

	deleteTicketFieldsQuery = `DELETE FROM ticket_fields WHERE id = :id`
)

func (t *ticketFieldsOps) SyncWithTicketFields(ctx context.Context, zendeskTicketFields []*SyncTicketField) error {
	tx, err := t.db.Begin()
	if err != nil {
		return errors.Wrapf(err, "models: [SyncWithTicketFields] db.Begin failed")
	}

	ids := make([]int, 0)
	tx.Select(&ids, "SELECT id FROM ticket_fields")

	dbIDs := make(map[int]struct{})
	for _, id := range ids {
		dbIDs[id] = struct{}{}
	}

	for _, zendeskTicketField := range zendeskTicketFields {
		dbTicketField := &db.TicketFields{
			ID:                  zendeskTicketField.ID,
			URL:                 zendeskTicketField.URL,
			Type:                zendeskTicketField.Type,
			Title:               zendeskTicketField.Title,
			RawTitle:            zendeskTicketField.RawTitle,
			Description:         zendeskTicketField.Description,
			RawDescription:      zendeskTicketField.RawDescription,
			Position:            zendeskTicketField.Position,
			Active:              zendeskTicketField.Active,
			Required:            zendeskTicketField.Required,
			CollapsedForAgents:  zendeskTicketField.CollapsedForAgents,
			RegexpForValidation: zendeskTicketField.RegexpForValidation,
			TitleInPortal:       zendeskTicketField.TitleInPortal,
			RawTitleInPortal:    zendeskTicketField.RawTitleInPortal,
			VisibleInPortal:     zendeskTicketField.VisibleInPortal,
			EditableInPortal:    zendeskTicketField.EditableInPortal,
			RequiredInPortal:    zendeskTicketField.RequiredInPortal,
			Tag:                 zendeskTicketField.Tag,
			CreatedAt:           zendeskTicketField.CreatedAt,
			UpdatedAt:           zendeskTicketField.UpdatedAt,
			Removable:           zendeskTicketField.Removable,
			CustomFieldOptions:  types.JSONText(zendeskTicketField.CustomFieldOptions),
			SystemFieldOptions:  types.JSONText(zendeskTicketField.SystemFieldOptions),
		}

		if _, exist := dbIDs[zendeskTicketField.ID]; exist {
			tx.NamedExec(updateTicketFieldsQuery, dbTicketField)
			delete(dbIDs, zendeskTicketField.ID)
		} else {
			tx.NamedExec(insertTicketFieldsQuery, dbTicketField)
		}
	}

	for id := range dbIDs {
		tx.NamedExec(deleteTicketFieldsQuery, map[string]interface{}{"id": id})
	}
	tx.Commit()

	return errors.Wrapf(tx.Err(), "models: [SyncWithTicketFields] db transaction failed")
}

func (t *ticketFieldsOps) GetTicketFieldByFieldID(ctx context.Context, fieldID int, locale string) (*TicketField, error) {
	field := new(db.TicketFields)
	query := fmt.Sprintf(
		`SELECT id,type,title,raw_title,description,raw_description,position,
		regexp_for_validation,title_in_portal,raw_title_in_portal,visible_in_portal,
		editable_in_portal,created_at,updated_at,custom_field_options,system_field_options 
		FROM ticket_fields WHERE id = '%d'`,
		fieldID,
	)
	if err := t.db.Get(ctx, field, query); err != nil {
		switch err {
		case db.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, errors.Wrapf(
				err,
				"models: [GetTicketField] db get field by id:%d, locale:%s failed",
				fieldID, locale,
			)
		}
	}

	// fontend only needs VisibleInPortal == true && EditableInPortal == true
	if !field.VisibleInPortal || !field.EditableInPortal {
		return nil, ErrNotFound
	}

	ret := &TicketField{
		ID:                  field.ID,
		Type:                field.Type,
		Title:               field.Title,
		RawTitle:            field.RawTitle,
		Description:         field.Description,
		RawDescription:      field.RawDescription,
		TitleInPortal:       field.TitleInPortal,
		RawTitleInPortal:    field.RawTitleInPortal,
		Position:            field.Position,
		RegexpForValidation: field.RegexpForValidation,
		CreatedAt:           field.CreatedAt,
		UpdatedAt:           field.UpdatedAt,
	}

	customFieldOptions := make([]*CustomFieldOption, 0)
	if err := field.CustomFieldOptions.Unmarshal(&customFieldOptions); err != nil {
		return nil, errors.Wrapf(err, "models: [GetTicketField] unmarshal custom field options failed")
	}
	ret.CustomFieldOptions = customFieldOptions

	systemFieldOptions := make([]*SystemFieldOption, 0)
	if err := field.SystemFieldOptions.Unmarshal(&systemFieldOptions); err != nil {
		return nil, errors.Wrapf(err, "models: [GetTicketField] unmarshal system field options failed")
	}
	ret.SystemFieldOptions = systemFieldOptions

	if strings.HasPrefix(ret.RawTitleInPortal, "{{") && strings.HasSuffix(ret.RawTitleInPortal, "}}") {
		dc, err := t.dcOps.GetDynamicContentItem(ctx, ret.RawTitleInPortal, locale)
		if err != nil {
			return nil, errors.Wrapf(
				err,
				"models: [GetTicketField] placeholder:%s, locale:%s on ret.RawTitleInPortal failed",
				ret.RawTitleInPortal,
				locale,
			)
		}
		ret.RawTitleInPortal = dc.VariantsContent
	}

	return ret, nil
}

func (t *ticketFieldsOps) GetTicketFieldByFormID(ctx context.Context, formID int, locale string) ([]*TicketField, error) {
	form := new(db.TicketForms)
	query := fmt.Sprintf(`SELECT ticket_field_ids FROM ticket_forms WHERE id = '%d'`, formID)
	if err := t.db.Get(ctx, form, query); err != nil {
		switch err {
		case db.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, errors.Wrapf(err, "models: [GetTicketFieldByFormID] db get form by id:%d locale:%s failed", formID, locale)
		}
	}

	ticketFieldRet := make([]*TicketField, 0)
	for _, fieldID := range form.TicketFieldIDs {
		field := new(db.TicketFields)
		query := fmt.Sprintf(
			`SELECT id,url,type,title,raw_title,description,raw_description,position,
			active,required,collapsed_for_agents,regexp_for_validation,title_in_portal,
			raw_title_in_portal,visible_in_portal,editable_in_portal,required_in_portal,
			tag,created_at,updated_at,removable,custom_field_options,system_field_options 
			FROM ticket_fields WHERE id = '%d'`,
			fieldID,
		)

		if err := t.db.Get(ctx, field, query); err != nil {
			switch err {
			case db.ErrNoRows:
				continue
			default:
				return nil, errors.Wrapf(err, "models: [GetTicketFieldByFormID] db get field by id:%d, locale:%s failed", fieldID, locale)
			}
		}

		// fontend only needs VisibleInPortal == true && EditableInPortal == true
		if !field.VisibleInPortal || !field.EditableInPortal {
			continue
		}

		if strings.HasPrefix(field.RawTitleInPortal, "{{") && strings.HasSuffix(field.RawTitleInPortal, "}}") {
			dc, err := t.dcOps.GetDynamicContentItem(ctx, field.RawTitleInPortal, locale)
			if err != nil {
				return nil, errors.Wrapf(
					err,
					"models: [GetTicketFieldByFormID] placeholder:%s, locale:%s on field.RawTitleInPortal failed",
					field.RawTitleInPortal,
					locale,
				)
			}
			field.RawTitleInPortal = dc.VariantsContent
		}

		ticketFieldRet = append(ticketFieldRet, &TicketField{
			ID:                  field.ID,
			URL:                 field.URL,
			Type:                field.Type,
			Title:               field.Title,
			RawTitle:            field.RawTitle,
			Description:         field.Description,
			RawDescription:      field.RawDescription,
			Position:            field.Position,
			Active:              field.Active,
			Required:            field.Required,
			CollapsedForAgents:  field.CollapsedForAgents,
			RegexpForValidation: field.RegexpForValidation,
			TitleInPortal:       field.TitleInPortal,
			RawTitleInPortal:    field.RawTitleInPortal,
			VisibleInPortal:     field.VisibleInPortal,
			EditableInPortal:    field.EditableInPortal,
			RequiredInPortal:    field.RequiredInPortal,
			Tag:                 field.Tag,
			CreatedAt:           field.CreatedAt,
			UpdatedAt:           field.UpdatedAt,
			Removable:           field.Removable,
		})
	}

	return ticketFieldRet, nil
}

func (t *ticketFieldsOps) GetTicketFieldCustomFieldOption(ctx context.Context, fieldID int) ([]*CustomFieldOption, error) {
	field := new(db.TicketFields)
	query := fmt.Sprintf(`SELECT custom_field_options FROM ticket_fields WHERE id = '%d'`, fieldID)
	if err := t.db.Get(ctx, field, query); err != nil {
		switch err {
		case db.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, errors.Wrapf(err, "models: [GetTicketFieldCustomFieldOption] db get field by id:%d failed", fieldID)
		}
	}

	customFieldOptions := make([]*CustomFieldOption, 0)
	if err := field.CustomFieldOptions.Unmarshal(&customFieldOptions); err != nil {
		return nil, errors.Wrapf(err, "models: [GetTicketFieldCustomFieldOption] unmarshal custom field options failed")
	}

	return customFieldOptions, nil
}

func (t *ticketFieldsOps) GetTicketFieldSystemFieldOption(ctx context.Context, fieldID int) ([]*SystemFieldOption, error) {
	field := new(db.TicketFields)
	query := fmt.Sprintf(`SELECT system_field_options FROM ticket_fields WHERE id = '%d'`, fieldID)
	if err := t.db.Get(ctx, field, query); err != nil {
		switch err {
		case db.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, errors.Wrapf(err, "models: [GetTicketFieldSystemFieldOption] db get field by id:%d failed", fieldID)
		}
	}

	systemFieldOptions := make([]*SystemFieldOption, 0)
	if err := field.SystemFieldOptions.Unmarshal(&systemFieldOptions); err != nil {
		return nil, errors.Wrapf(err, "models: [GetTicketFieldSystemFieldOption] unmarshal system field options failed")
	}

	return systemFieldOptions, nil
}
