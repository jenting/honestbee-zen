package resolvers

import (
	"context"
	"strconv"

	gographql "github.com/graph-gophers/graphql-go"

	"github.com/honestbee/Zen/dataloader"
	"github.com/honestbee/Zen/inout"
	"github.com/honestbee/Zen/models"
)

// TicketFieldResolver defines resolver models.
type TicketFieldResolver struct {
	m *models.TicketField
}

// ID is the TicketForm's field updated_at.
func (r *TicketFieldResolver) ID(ctx context.Context) gographql.ID {
	return gographql.ID(strconv.Itoa(r.m.ID))
}

// URL is the TicketForm's field url.
func (r *TicketFieldResolver) URL(ctx context.Context) string {
	return r.m.URL
}

// Type is the TicketForm's field type.
func (r *TicketFieldResolver) Type(ctx context.Context) string {
	return r.m.Type
}

// Title is the TicketForm's field title.
func (r *TicketFieldResolver) Title(ctx context.Context) string {
	return r.m.Title
}

// RawTitle is the TicketForm's field raw_title.
func (r *TicketFieldResolver) RawTitle(ctx context.Context) string {
	return r.m.RawTitle
}

// Description is the TicketForm's field description.
func (r *TicketFieldResolver) Description(ctx context.Context) string {
	return r.m.Description
}

// RawDescription is the TicketForm's field raw_description.
func (r *TicketFieldResolver) RawDescription(ctx context.Context) string {
	return r.m.RawDescription
}

// Position is the TicketForm's field position.
func (r *TicketFieldResolver) Position(ctx context.Context) int32 {
	return int32(r.m.Position)
}

// Active is the TicketForm's field active.
func (r *TicketFieldResolver) Active(ctx context.Context) bool {
	return r.m.Active
}

// Required is the TicketForm's field required.
func (r *TicketFieldResolver) Required(ctx context.Context) bool {
	return r.m.Required
}

// CollapsedForAgents is the TicketForm's field collapsed_for_agents.
func (r *TicketFieldResolver) CollapsedForAgents(ctx context.Context) bool {
	return r.m.CollapsedForAgents
}

// RegexpForValidation is the TicketForm's field regexp_for_validation.
func (r *TicketFieldResolver) RegexpForValidation(ctx context.Context) string {
	return r.m.RegexpForValidation
}

// TitleInPortal is the TicketForm's field title_in_portal.
func (r *TicketFieldResolver) TitleInPortal(ctx context.Context) string {
	return r.m.TitleInPortal
}

// RawTitleInPortal is the TicketForm's field raw_title_in_portal.
func (r *TicketFieldResolver) RawTitleInPortal(ctx context.Context) string {
	return r.m.RawTitleInPortal
}

// VisibleInPortal is the TicketForm's field visible_in_portal.
func (r *TicketFieldResolver) VisibleInPortal(ctx context.Context) bool {
	return r.m.VisibleInPortal
}

// EditableInPortal is the TicketForm's field editable_in_portal.
func (r *TicketFieldResolver) EditableInPortal(ctx context.Context) bool {
	return r.m.EditableInPortal
}

// RequiredInPortal is the TicketForm's field required_in_portal.
func (r *TicketFieldResolver) RequiredInPortal(ctx context.Context) bool {
	return r.m.RequiredInPortal
}

// Tag is the TicketForm's field tag.
func (r *TicketFieldResolver) Tag(ctx context.Context) string {
	return r.m.Tag
}

// CreatedAt is the TicketForm's field created_at.
func (r *TicketFieldResolver) CreatedAt(ctx context.Context) gographql.Time {
	return gographql.Time{Time: r.m.CreatedAt}
}

// UpdatedAt is the TicketForm's field updated_at.
func (r *TicketFieldResolver) UpdatedAt(ctx context.Context) gographql.Time {
	return gographql.Time{Time: r.m.UpdatedAt}
}

// Removable is the TicketForm's field removable.
func (r *TicketFieldResolver) Removable(ctx context.Context) bool {
	return r.m.Removable
}

// CustomFieldOptions is the TicketForm's field custom_field_options.
func (r *TicketFieldResolver) CustomFieldOptions(ctx context.Context) (*[]*CustomFieldOptionResolver, error) {
	id := gographql.ID(strconv.Itoa(r.m.ID))
	data := inout.QueryCustomFieldOptionsIn{
		FieldID: id,
	}

	loads, err := dataloader.LoadTicketFieldCustomFieldOptions(ctx, data)
	if err != nil {
		return nil, err
	}

	// Translate results.
	ret := make([]*CustomFieldOptionResolver, 0)
	for _, l := range loads {
		ret = append(ret, &CustomFieldOptionResolver{m: l})
	}

	return &ret, nil
}

// SystemFieldOptions is the TicketForm's field system_field_options.
func (r *TicketFieldResolver) SystemFieldOptions(ctx context.Context) (*[]*SystemFieldOptionResolver, error) {
	id := gographql.ID(strconv.Itoa(r.m.ID))
	data := inout.QuerySystemFieldOptionsIn{
		FieldID: id,
	}

	loads, err := dataloader.LoadTicketFieldSystemFieldOptions(ctx, data)
	if err != nil {
		return nil, err
	}

	// Translate result.
	ret := make([]*SystemFieldOptionResolver, 0)
	for _, l := range loads {
		ret = append(ret, &SystemFieldOptionResolver{m: l})
	}

	return &ret, nil
}

// CustomFieldOptionResolver defines resolver models.
type CustomFieldOptionResolver struct {
	m *models.CustomFieldOption
}

// ID is the CustomFieldOptionResolver's field id.
func (r *CustomFieldOptionResolver) ID(ctx context.Context) gographql.ID {
	return gographql.ID(strconv.Itoa(r.m.ID))
}

// Name is the CustomFieldOptionResolver's field name.
func (r *CustomFieldOptionResolver) Name(ctx context.Context) string {
	return r.m.Name
}

// RawName is the CustomFieldOptionResolver's field raw_name.
func (r *CustomFieldOptionResolver) RawName(ctx context.Context) string {
	return r.m.RawName
}

// Value is the CustomFieldOptionResolver's field value.
func (r *CustomFieldOptionResolver) Value(ctx context.Context) string {
	return r.m.Value
}

// SystemFieldOptionResolver defines resolver models.
type SystemFieldOptionResolver struct {
	m *models.SystemFieldOption
}

// Name is the SystemFieldOptionResolver's field name.
func (r *SystemFieldOptionResolver) Name(ctx context.Context) string {
	return r.m.Name
}

// Value is the SystemFieldOptionResolver's field value.
func (r *SystemFieldOptionResolver) Value(ctx context.Context) string {
	return r.m.Value
}
