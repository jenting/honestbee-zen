package handlers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"

	gographql "github.com/graph-gophers/graphql-go"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"

	"github.com/honestbee/Zen/errs"
	"github.com/honestbee/Zen/inout"
)

// CreateGraphQLDecompressor combines params from URL or FORM
// and returns params in a structure that CreateGraphQLHandler needs.
func CreateGraphQLDecompressor(ps httprouter.Params, r *http.Request) (interface{}, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errs.NewErr(
			errs.InvalidAttributeErrorCode,
			errors.Errorf("handlers: [CreateGraphQLDecompressor] read http request body failed"),
		)
	}
	defer r.Body.Close()

	if len(body) == 0 {
		return nil, errs.NewErr(
			errs.InvalidAttributeErrorCode,
			errors.Errorf("handlers: [CreateGraphQLDecompressor] no queries to execute"),
		)
	}

	var queries []inout.GraphQLQuery
	var isBatch bool
	// Inspect the first character to inform how the body is parsed.
	switch body[0] {
	case '{':
		query := inout.GraphQLQuery{}
		if err := json.Unmarshal(body, &query); err != nil {
			return nil, errs.NewErr(
				errs.ServerInternalErrorCode,
				errors.Errorf("handlers: [CreateGraphQLDecompressor] json unmarshal failed"),
			)
		}
		queries = append(queries, query)
	case '[':
		isBatch = true
		if err := json.Unmarshal(body, &queries); err != nil {
			return nil, errs.NewErr(
				errs.ServerInternalErrorCode,
				errors.Errorf("handlers: [CreateGraphQLDecompressor] json unmarshal failed"),
			)
		}
	default:
		return nil, errs.NewErr(
			errs.InvalidAttributeErrorCode,
			errors.Errorf("handlers: [CreateGraphQLDecompressor] invalid query body"),
		)
	}

	return &inout.GraphQLIn{
		Ctx:     r.Context(),
		Queries: queries,
		IsBatch: isBatch,
	}, nil
}

// CreateGraphQLHandler handles graphql operation.
func CreateGraphQLHandler(ctx context.Context, e *Env, in interface{}) (interface{}, error) {
	data, ok := in.(*inout.GraphQLIn)
	if !ok {
		return nil, errs.NewErr(
			errs.ServerInternalErrorCode,
			errors.Errorf("handlers: [CreateGraphQLHandler] cast %v into *GraphQLIn failed", in),
		)
	}

	lens := len(data.Queries)
	ctx = e.GraphQL.Loader.Attach(ctx) // Attach dataloaders onto the request context.
	var (
		responses      = make([]*gographql.Response, lens) // Allocate a slice large enough for all responses.
		wg             sync.WaitGroup                      // Use the WaitGroup to wait for all executions to finish.
		resolverErrors error                               // Concatenate resolver errors.
	)

	wg.Add(lens)

	for i, q := range data.Queries {
		// Loop throught the parsed queries from the request.
		// These queries are executed in separate goroutines so they
		// process in parallel.
		go func(i int, q inout.GraphQLQuery) {
			res := e.GraphQL.Schema.Exec(ctx, q.Query, q.OpName, q.Variables)

			// We have to do some work here to expand errors when it is possible for a resolver to return
			// more than one error (for example, a list resolver).
			res.Errors, resolverErrors = Expand(res.Errors)
			responses[i] = res

			wg.Done()
		}(i, q)
	}

	wg.Wait()

	if data.IsBatch {
		return responses, resolverErrors
	}
	return responses[0], resolverErrors
}
