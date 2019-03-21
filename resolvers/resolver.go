package resolvers

import (
	"github.com/rs/zerolog"

	"github.com/honestbee/Zen/config"
	"github.com/honestbee/Zen/examiner"
	"github.com/honestbee/Zen/models"
	"github.com/honestbee/Zen/zendesk"
)

// Resolver is the query/mutation resolver.
type Resolver struct {
	conf     *config.Config
	logger   *zerolog.Logger
	service  models.Service
	examiner *examiner.Examiner
	zendesk  *zendesk.ZenDesk
}
