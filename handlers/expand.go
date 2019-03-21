package handlers

import (
	gographqlerrors "github.com/graph-gophers/graphql-go/errors"
	"github.com/pkg/errors"

	"github.com/honestbee/Zen/errs"
)

type slicer interface {
	Slice() []error
}

type indexedCauser interface {
	Index() int
	Cause() error
}

// Expand expands graphql query error.
func Expand(queryErrors []*gographqlerrors.QueryError) ([]*gographqlerrors.QueryError, error) {
	expanded := make([]*gographqlerrors.QueryError, 0, len(queryErrors))
	var resolverErrors error

	for _, queryError := range queryErrors {
		switch t := queryError.ResolverError.(type) {
		case slicer:
			for _, e := range t.Slice() {
				qe := &gographqlerrors.QueryError{
					Message:   queryError.Message,
					Locations: queryError.Locations,
					Path:      queryError.Path,
				}

				if ic, ok := e.(indexedCauser); ok {
					qe.Path = append(qe.Path, ic.Index())
					qe.Message = ic.Cause().Error()
				}

				expanded = append(expanded, qe)

				// Concatenate resolver errors.
				if qe.ResolverError != nil {
					if resolverErrors == nil {
						resolverErrors = qe.ResolverError
					} else {
						resolverErrors = errors.Wrap(resolverErrors, qe.ResolverError.Error())
					}
				}
			}
		case *errs.Error: // Custom errors
			// Replace output message to output error.
			queryError.Message = t.OutputErr
			expanded = append(expanded, queryError)

			// Concatenate resolver errors.
			if queryError.ResolverError != nil {
				if resolverErrors == nil {
					resolverErrors = queryError.ResolverError
				} else {
					resolverErrors = errors.Wrap(resolverErrors, queryError.ResolverError.Error())
				}
			}
		default:
			expanded = append(expanded, queryError)
		}
	}

	return expanded, resolverErrors
}
