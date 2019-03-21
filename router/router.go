package router

import (
	"net/http"

	"github.com/rs/zerolog"
	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/julienschmidt/httprouter"

	"github.com/honestbee/Zen/config"
	"github.com/honestbee/Zen/examiner"
	"github.com/honestbee/Zen/handlers"
	"github.com/honestbee/Zen/models"
	"github.com/honestbee/Zen/resolvers"
	"github.com/honestbee/Zen/zendesk"
)

// New returns a http.Server capable router with all routing handlers.
func New(conf *config.Config,
	logger *zerolog.Logger,
	service models.Service,
	examiner *examiner.Examiner,
	zend *zendesk.ZenDesk,
	graphql *resolvers.GraphQL) (*httptrace.Router, error) {

	e := &handlers.Env{
		Config:   conf,
		Logger:   logger,
		Service:  service,
		Examiner: examiner,
		ZenDesk:  zend,
		GraphQL:  graphql,
	}

	mux := httptrace.New(httptrace.WithServiceName("helpcenter-zendesk-http"))
	mux.PanicHandler = func(w http.ResponseWriter, r *http.Request, v interface{}) {
		e.Logger.Error().Fields(map[string]interface{}{
			"from":   r.RemoteAddr,
			"path":   r.RequestURI,
			"method": r.Method,
		}).Msgf("panic:%v", v)
	}

	// RESTful handlers.
	mux.GET("/api/categories", handlers.Middleware(e, handlers.GetCategoriesDecompressor, handlers.GetCategoriesHandler))
	mux.GET("/api/categories/:category_id/sections", handlers.Middleware(e, handlers.GetSectionsDecompressor, handlers.GetSectionsHandler))
	mux.GET("/api/categories/:category_id/articles", handlers.Middleware(e, handlers.GetCategoriesArticlesDecompressor, handlers.GetCategoriesArticlesHandler))
	mux.GET("/api/category/:category_key_name", handlers.Middleware(e, handlers.GetCategoryKeyNameToIDDecompressor, handlers.GetCategoryKeyNameToIDHandler))
	mux.GET("/api/sections/:section_id/articles", handlers.Middleware(e, handlers.GetArticlesDecompressor, handlers.GetArticlesHandler))
	mux.GET("/api/sections/:section_id", handlers.Middleware(e, handlers.GetSectionDecompressor, handlers.GetSectionHandler))
	mux.GET("/api/articles/:article_id", handlers.Middleware(e, handlers.GetArticleDecompressor, handlers.GetArticleHandler))
	mux.GET("/api/toparticles/:top_n", handlers.Middleware(e, handlers.GetTopNArticlesDecompressor, handlers.GetTopNArticlesHandler))
	mux.GET("/api/ticket_forms/:form_id", handlers.Middleware(e, handlers.GetTicketFormDecompressor, handlers.GetTicketFormHandler))
	mux.GET("/api/instant_search", handlers.Middleware(e, handlers.GetInstantSearchDecompressor, handlers.GetInstantSearchHandler))
	mux.GET("/api/search", handlers.Middleware(e, handlers.GetSearchDecompressor, handlers.GetSearchHandler))
	mux.GET("/api/status", handlers.StatusHandler)
	mux.POST("/api/requests", handlers.Middleware(e, handlers.CreateRequestDecompressor, handlers.CreateRequestHandler))
	mux.POST("/api/vote/:article_id/:value", handlers.Middleware(e, handlers.CreateVoteDecompressor, handlers.CreateVoteHandler))
	mux.POST("/api/forcesync", handlers.Middleware(e, handlers.CreateForceSyncDecompressor, handlers.CreateForceSyncHandler))

	// GraphQL handlers.
	mux.POST("/graphql", handlers.GraphQLMiddleware(e, handlers.CreateGraphQLDecompressor, handlers.CreateGraphQLHandler))
	mux.Handler("GET", "/graphiql", handlers.GraphiQL{})

	return mux, nil
}
