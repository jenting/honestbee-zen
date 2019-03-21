package resolvers

import (
	gographql "github.com/graph-gophers/graphql-go"
	"github.com/rs/zerolog"
	graphqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/graph-gophers/graphql-go"

	"github.com/honestbee/Zen/config"
	"github.com/honestbee/Zen/dataloader"
	"github.com/honestbee/Zen/examiner"
	"github.com/honestbee/Zen/models"
	"github.com/honestbee/Zen/schema"
	"github.com/honestbee/Zen/zendesk"
)

// GraphQL struct handles GraphQL API requests over HTTP.
type GraphQL struct {
	Schema *gographql.Schema
	Loader dataloader.Collection
}

// New return the new GraphQL.
func New(conf *config.Config,
	logger *zerolog.Logger,
	service models.Service,
	examiner *examiner.Examiner,
	zendesk *zendesk.ZenDesk) (*GraphQL, error) {

	return &GraphQL{
		Schema: gographql.MustParseSchema(
			schema.String(),
			&Resolver{
				conf:     conf,
				logger:   logger,
				service:  service,
				examiner: examiner,
				zendesk:  zendesk,
			},
			gographql.Tracer(graphqltrace.NewTracer(graphqltrace.WithServiceName("helpcenter-zendesk-graphql"))),
			gographql.MaxDepth(conf.GraphQL.MaxDepth),
			gographql.MaxParallelism(conf.GraphQL.MaxParallelism),
		),
		Loader: dataloader.Initialize(service, examiner, zendesk),
	}, nil
}
