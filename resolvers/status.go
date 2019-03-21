package resolvers

import (
	"context"
	"runtime"
	"time"

	gographql "github.com/graph-gophers/graphql-go"

	"github.com/honestbee/Zen/config"
)

// StatusResolver defines resolver models.
type StatusResolver struct {
}

// GoVersion is the Status's field go_version.
func (r *StatusResolver) GoVersion(ctx context.Context) string {
	return runtime.Version()
}

// AppVersion is the Status's field app_version.
func (r *StatusResolver) AppVersion(ctx context.Context) string {
	return config.Version
}

// ServerTime is the Status's field server_time.
func (r *StatusResolver) ServerTime(ctx context.Context) gographql.Time {
	return gographql.Time{Time: time.Now().UTC()}
}
