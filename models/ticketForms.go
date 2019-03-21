package models

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"

	"github.com/honestbee/Zen/internal/db"
)

type ticketFormsService interface {
	SyncWithTicketForms(ctx context.Context, zendeskTicketForms []*SyncTicketForm) error
	GetTicketForm(ctx context.Context, formID int, locale string) (*TicketForm, error)
	GetTicketFormGraphQL(ctx context.Context, formID int) (*SyncTicketForm, error)
}

// TicketForm is the ticket form model.
type TicketForm struct {
	ID                 int            `json:"id,omitempty"`
	URL                string         `json:"-"`
	Name               string         `json:"name,omitempty"`
	RawName            string         `json:"raw_name,omitempty"`
	DisplayName        string         `json:"display_name,omitempty"`
	RawDisplayName     string         `json:"raw_display_name,omitempty"`
	EndUserVisible     bool           `json:"-"`
	Position           int            `json:"position,omitempty"`
	Active             bool           `json:"-"`
	InAllBrands        bool           `json:"-"`
	RestrictedBrandIDs []int64        `json:"-"`
	TicketFields       []*TicketField `json:"ticket_fields,omitempty"`
	CreatedAt          time.Time      `json:"created_at,omitempty"`
	UpdatedAt          time.Time      `json:"updated_at,omitempty"`
}

// SyncTicketForm is the struct for SyncWithTicketForms to sync up database data.
type SyncTicketForm struct {
	ID                 int
	URL                string
	Name               string
	RawName            string
	DisplayName        string
	RawDisplayName     string
	EndUserVisible     bool
	Position           int
	Active             bool
	InAllBrands        bool
	RestrictedBrandIDs []int64
	TicketFieldIDs     []int64
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type ticketFormsOps struct {
	db        db.Database
	fieldsOps *ticketFieldsOps
	dcOps     *dynamicContentOps
}

const (
	updateTicketFormsQuery = `
		UPDATE ticket_forms SET
			url = :url,
			name = :name,
			raw_name = :raw_name,
			display_name = :display_name,
			raw_display_name = :raw_display_name,
			end_user_visible = :end_user_visible,
			position = :position,
			active = :active,
			in_all_brands = :in_all_brands,
			restricted_brand_ids = :restricted_brand_ids,
			ticket_field_ids = :ticket_field_ids,
			created_at = :created_at,
			updated_at = :updated_at
		WHERE id = :id`

	insertTicketFormsQuery = `
		INSERT INTO ticket_forms (
			id,
			url,
			name,
			raw_name,
			display_name,
			raw_display_name,
			end_user_visible,
			position,
			active,
			in_all_brands,
			restricted_brand_ids,
			ticket_field_ids,
			created_at,
			updated_at
		) 
		VALUES (
			:id,
			:url,
			:name,
			:raw_name,
			:display_name,
			:raw_display_name,
			:end_user_visible,
			:position,
			:active,
			:in_all_brands,
			:restricted_brand_ids,
			:ticket_field_ids,
			:created_at,
			:updated_at
		)`

	deleteTicketFormsQuery = `DELETE FROM ticket_forms WHERE id = :id`
)

func (t *ticketFormsOps) SyncWithTicketForms(ctx context.Context, zendeskTicketForms []*SyncTicketForm) error {
	tx, err := t.db.Begin()
	if err != nil {
		return errors.Wrapf(err, "models: [SyncWithTicketForms] db.Begin failed")
	}

	ids := make([]int, 0)
	tx.Select(&ids, "SELECT id FROM ticket_forms")

	dbIDs := make(map[int]struct{})
	for _, id := range ids {
		dbIDs[id] = struct{}{}
	}

	for _, zendeskTicketForm := range zendeskTicketForms {
		dbTicketForm := &db.TicketForms{
			ID:                 zendeskTicketForm.ID,
			URL:                zendeskTicketForm.URL,
			Name:               zendeskTicketForm.Name,
			RawName:            zendeskTicketForm.RawName,
			DisplayName:        zendeskTicketForm.DisplayName,
			RawDisplayName:     zendeskTicketForm.RawDisplayName,
			EndUserVisible:     zendeskTicketForm.EndUserVisible,
			Position:           zendeskTicketForm.Position,
			Active:             zendeskTicketForm.Active,
			InAllBrands:        zendeskTicketForm.InAllBrands,
			RestrictedBrandIDs: zendeskTicketForm.RestrictedBrandIDs,
			TicketFieldIDs:     zendeskTicketForm.TicketFieldIDs,
			CreatedAt:          zendeskTicketForm.CreatedAt,
			UpdatedAt:          zendeskTicketForm.UpdatedAt,
		}
		if dbTicketForm.RestrictedBrandIDs == nil {
			dbTicketForm.RestrictedBrandIDs = make([]int64, 0)
		}
		if dbTicketForm.TicketFieldIDs == nil {
			dbTicketForm.TicketFieldIDs = make([]int64, 0)
		}

		if _, exist := dbIDs[zendeskTicketForm.ID]; exist {
			tx.NamedExec(updateTicketFormsQuery, dbTicketForm)
			delete(dbIDs, zendeskTicketForm.ID)
		} else {
			tx.NamedExec(insertTicketFormsQuery, dbTicketForm)
		}
	}

	for id := range dbIDs {
		tx.NamedExec(deleteTicketFormsQuery, map[string]interface{}{"id": id})
	}
	tx.Commit()

	return errors.Wrapf(tx.Err(), "models: [SyncWithTicketForms] db transaction failed")
}

func (t *ticketFormsOps) GetTicketForm(ctx context.Context, formID int, locale string) (*TicketForm, error) {
	form := new(db.TicketForms)
	query := fmt.Sprintf(
		`SELECT id,url,name,raw_name,display_name,raw_display_name,end_user_visible,
		position,active,in_all_brands,restricted_brand_ids,created_at,updated_at,ticket_field_ids 
		FROM ticket_forms WHERE id = '%d'`,
		formID,
	)
	if err := t.db.Get(ctx, form, query); err != nil {
		switch err {
		case db.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, errors.Wrapf(
				err,
				"models: [GetTicketForm] db get form by id:%d locale:%s failed",
				formID, locale,
			)
		}
	}

	ret := &TicketForm{
		ID:                 form.ID,
		URL:                form.URL,
		Name:               form.Name,
		RawName:            form.RawName,
		DisplayName:        form.DisplayName,
		RawDisplayName:     form.RawDisplayName,
		EndUserVisible:     form.EndUserVisible,
		Position:           form.Position,
		Active:             form.Active,
		InAllBrands:        form.InAllBrands,
		RestrictedBrandIDs: form.RestrictedBrandIDs,
		CreatedAt:          form.CreatedAt,
		UpdatedAt:          form.UpdatedAt,
		TicketFields:       make([]*TicketField, 0),
	}

	for _, fieldID := range form.TicketFieldIDs {
		field, err := t.fieldsOps.GetTicketFieldByFieldID(ctx, int(fieldID), locale)
		if err != nil {
			switch err {
			case ErrNotFound:
				continue
			default:
				return nil, errors.Wrapf(err, "models: [GetTicketForm] db get field failed")
			}
		}
		ret.TicketFields = append(ret.TicketFields, field)
	}

	return ret, nil
}

func (t *ticketFormsOps) GetTicketFormGraphQL(ctx context.Context, formID int) (*SyncTicketForm, error) {
	form := new(db.TicketForms)
	query := fmt.Sprintf(
		`SELECT id,url,name,raw_name,display_name,raw_display_name,end_user_visible,
		position,active,in_all_brands,restricted_brand_ids,created_at,updated_at,ticket_field_ids 
		FROM ticket_forms WHERE id = '%d'`,
		formID,
	)
	if err := t.db.Get(ctx, form, query); err != nil {
		switch err {
		case db.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, errors.Wrapf(
				err,
				"models: [GetTicketForm] db get form by id:%d failed",
				formID,
			)
		}
	}

	ret := &SyncTicketForm{
		ID:                 form.ID,
		URL:                form.URL,
		Name:               form.Name,
		RawName:            form.RawName,
		DisplayName:        form.DisplayName,
		RawDisplayName:     form.RawDisplayName,
		EndUserVisible:     form.EndUserVisible,
		Position:           form.Position,
		Active:             form.Active,
		InAllBrands:        form.InAllBrands,
		RestrictedBrandIDs: form.RestrictedBrandIDs,
		TicketFieldIDs:     form.TicketFieldIDs,
		CreatedAt:          form.CreatedAt,
		UpdatedAt:          form.UpdatedAt,
	}

	return ret, nil
}
