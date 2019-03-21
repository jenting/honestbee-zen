package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"

	"github.com/honestbee/Zen/config"
	"github.com/honestbee/Zen/errs"
	"github.com/honestbee/Zen/examiner"
	"github.com/honestbee/Zen/models"
	"github.com/honestbee/Zen/resolvers"
	"github.com/honestbee/Zen/zendesk"
)

// Env is the application-wid configuration.
type Env struct {
	Config   *config.Config
	Logger   *zerolog.Logger
	Service  models.HelpDeskService
	Examiner *examiner.Examiner
	ZenDesk  *zendesk.ZenDesk
	GraphQL  *resolvers.GraphQL
}

type decompressor func(httprouter.Params, *http.Request) (interface{}, error)
type handler func(ctx context.Context, e *Env, in interface{}) (interface{}, error)

type processor struct {
	err     error
	e       *Env
	source1 httprouter.Params
	source2 *http.Request
	product interface{}
}

func (p *processor) preparation(f decompressor) {
	p.product, p.err = f(p.source1, p.source2)
}

func (p *processor) handling(f handler) {
	if p.err != nil {
		return
	}
	p.product, p.err = f(p.source2.Context(), p.e, p.product)
}

func (p *processor) production(f func(v interface{}) error) {
	if p.err != nil {
		return
	}
	p.err = f(p.product)
}

// Middleware pre-handle every incoming request and generate output to the client based on every handler returned error code.
func Middleware(e *Env, dec decompressor, fn handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(w)

		proc := &processor{
			e:       e,
			source1: p,
			source2: r,
		}

		e.Logger.Info().Fields(map[string]interface{}{
			"from":   r.RemoteAddr,
			"path":   r.URL.Path,
			"method": r.Method,
			"agent":  r.UserAgent(),
		}).Msgf("receiving data")

		proc.preparation(dec)
		proc.handling(fn)
		proc.production(encoder.Encode)

		if proc.err != nil {
			er, ok := proc.err.(*errs.Error)
			if !ok {
				// Any error types we don't specifically look out for default
				// to serving a HTTP 500.
				er = errs.NewErr(errs.ServerInternalErrorCode, proc.err)
			}
			if er.InternalErr != nil {
				e.Logger.Error().Fields(map[string]interface{}{
					"from":   r.RemoteAddr,
					"path":   r.URL.Path,
					"method": r.Method,
					"agent":  r.UserAgent(),
					"error":  er.Error(),
				}).Msgf("middleware error occurred")
			}
			w.WriteHeader(er.Status)
			encoder.Encode(er)
		}
	}
}

// GraphQLMiddleware pre-handle every incoming request and generate output to the client based on every handler returned error code.
func GraphQLMiddleware(e *Env, dec decompressor, fn handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(w)

		proc := &processor{
			e:       e,
			source1: p,
			source2: r,
		}

		e.Logger.Info().Fields(map[string]interface{}{
			"from":   r.RemoteAddr,
			"path":   r.URL.Path,
			"method": r.Method,
		}).Msgf("receiving data")

		proc.preparation(dec)
		proc.handling(fn)
		encoder.Encode(proc.product)

		er, ok := proc.err.(*errs.Error)
		if !ok {
			// Any error types we don't specifically look out for default
			// to serving a HTTP 500.
			er = errs.NewErr(errs.ServerInternalErrorCode, proc.err)
		}
		if er.InternalErr != nil {
			e.Logger.Error().Fields(map[string]interface{}{
				"from":   r.RemoteAddr,
				"path":   r.URL.Path,
				"method": r.Method,
				"agent":  r.UserAgent(),
				"error":  er.Error(),
			}).Msgf("middleware error occurred")
		}
	}
}
