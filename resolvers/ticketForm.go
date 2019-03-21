package resolvers

import (
	"context"
	"strconv"

	gographql "github.com/graph-gophers/graphql-go"

	"github.com/honestbee/Zen/dataloader"
	"github.com/honestbee/Zen/inout"
	"github.com/honestbee/Zen/models"
)

// TicketFormResolver defines resolver models.
type TicketFormResolver struct {
	m *models.SyncTicketForm
}

// ID is the TicketForm's field id.
func (r *TicketFormResolver) ID(ctx context.Context) gographql.ID {
	return gographql.ID(strconv.Itoa(r.m.ID))
}

// URL is the TicketForm's field url.
func (r *TicketFormResolver) URL(ctx context.Context) string {
	return r.m.URL
}

// Name is the TicketForm's field name.
func (r *TicketFormResolver) Name(ctx context.Context) string {
	return r.m.Name
}

// RawName is the TicketForm's field raw_name.
func (r *TicketFormResolver) RawName(ctx context.Context) string {
	return r.m.RawName
}

// DisplayName is the TicketForm's field display_name.
func (r *TicketFormResolver) DisplayName(ctx context.Context) string {
	return r.m.DisplayName
}

// RawDisplayName is the TicketForm's field raw_display_name.
func (r *TicketFormResolver) RawDisplayName(ctx context.Context) string {
	return r.m.RawDisplayName
}

// EndUserVisible is the TicketForm's field end_user_visible.
func (r *TicketFormResolver) EndUserVisible(ctx context.Context) bool {
	return r.m.EndUserVisible
}

// Position is the TicketForm's field end_user_visible.
func (r *TicketFormResolver) Position(ctx context.Context) int32 {
	return int32(r.m.Position)
}

// Active is the TicketForm's field active.
func (r *TicketFormResolver) Active(ctx context.Context) bool {
	return r.m.Active
}

// InAllBrands is the TicketForm's field in_all_brands.
func (r *TicketFormResolver) InAllBrands(ctx context.Context) bool {
	return r.m.InAllBrands
}

// RestrictedBrandIDs is the TicketForm's field restricted_brand_ids.
func (r *TicketFormResolver) RestrictedBrandIDs(ctx context.Context) []int32 {
	ret := make([]int32, 0)
	for _, id := range r.m.RestrictedBrandIDs {
		ret = append(ret, int32(id))
	}
	return ret
}

// TicketFieldsConnection is the TicketForm's field ticket_fields_connection.
func (r *TicketFormResolver) TicketFieldsConnection(ctx context.Context, data inout.QueryTicketFieldsIn) (*[]*TicketFieldResolver, error) {
	// Process input params.
	if err := data.ProcessInputParams(); err != nil {
		return nil, err
	}

	if data.FormID == nil {
		id := r.ID(ctx)
		data.FormID = &id
	}

	// Load ticket_fields.
	loads, err := dataloader.LoadTicketFields(ctx, data)
	if err != nil {
		return nil, err
	}

	// Translate results.
	ret := make([]*TicketFieldResolver, 0)
	for _, l := range loads {
		ret = append(ret, &TicketFieldResolver{m: l})
	}

	return &ret, nil
}

// CreatedAt is the TicketForm's field created_at.
func (r *TicketFormResolver) CreatedAt(ctx context.Context) gographql.Time {
	return gographql.Time{Time: r.m.CreatedAt}
}

// UpdatedAt is the TicketForm's field updated_at.
func (r *TicketFormResolver) UpdatedAt(ctx context.Context) gographql.Time {
	return gographql.Time{Time: r.m.UpdatedAt}
}
