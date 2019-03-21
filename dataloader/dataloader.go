package dataloader

import (
	"context"
	"fmt"

	"github.com/graph-gophers/dataloader"

	"github.com/honestbee/Zen/examiner"
	"github.com/honestbee/Zen/models"
	"github.com/honestbee/Zen/zendesk"
)

// DataLoader cache only serves the purpose of not repeatedly loading the
// same data in the context of a single request to your Application.

// Avoid multiple requests from different users using the DataLoader instance,
// which could result in cached data incorrectly appearing in each request.
// Typically, DataLoader instances are created when a Request begins,
// and are not used once the Request ends.
const (
	categoriesLoaderKey          dataloader.StringKey = "categories"
	categoryLoaderKey            dataloader.StringKey = "category"
	sectionsLoaderKey            dataloader.StringKey = "sections"
	sectionLoaderKey             dataloader.StringKey = "section"
	articlesLoaderKey            dataloader.StringKey = "articles"
	topArticlesLoaderKey         dataloader.StringKey = "toparticles"
	articleLoaderKey             dataloader.StringKey = "article"
	ticketFormLoaderKey          dataloader.StringKey = "ticket_form"
	ticketFieldsLoaderKey        dataloader.StringKey = "ticket_fields"
	ticketFieldCustomFieldOption dataloader.StringKey = "ticket_field_custom_field_option"
	ticketFieldSystemFieldOption dataloader.StringKey = "ticket_field_system_field_option"
	searchTitleArticlesLoaderKey dataloader.StringKey = "search_title_articles"
	searchBodyArticlesLoaderKey  dataloader.StringKey = "search_body_articles"
)

// Collection holds an internal lookup of initialized loader data load functions.
type Collection struct {
	lookup map[dataloader.StringKey]dataloader.BatchFunc
}

// Initialize a lookup map of context keys to loader functions.
func Initialize(
	service models.Service,
	examiner *examiner.Examiner,
	zend *zendesk.ZenDesk) Collection {
	return Collection{
		lookup: map[dataloader.StringKey]dataloader.BatchFunc{
			categoriesLoaderKey:          newCategoriesLoader(service, examiner),
			categoryLoaderKey:            newCategoryLoader(service, examiner),
			sectionsLoaderKey:            newSectionsLoader(service, examiner),
			sectionLoaderKey:             newSectionLoader(service, examiner),
			articlesLoaderKey:            newArticlesLoader(service, examiner),
			topArticlesLoaderKey:         newTopArticlesLoader(service),
			articleLoaderKey:             newArticleLoader(service, examiner),
			ticketFormLoaderKey:          newTicketFormLoader(service, examiner),
			ticketFieldsLoaderKey:        newTicketFieldsLoader(service),
			ticketFieldCustomFieldOption: newTicketFieldCustomFieldOptionsLoader(service),
			ticketFieldSystemFieldOption: newTicketFieldSystemFieldOptionsLoader(service),
			searchTitleArticlesLoaderKey: newSearchTitleArticlesLoader(zend),
			searchBodyArticlesLoaderKey:  newSearchBodyArticlesLoader(service, zend),
		},
	}
}

// Attach creates new instances of dataloader.Loader and attaches the instances on the request context.
// We do this because the dataloader has an in-memory cache that is scoped to the request.
func (c Collection) Attach(ctx context.Context) context.Context {
	for k, v := range c.lookup {
		ctx = context.WithValue(ctx, k, dataloader.NewBatchedLoader(v))
	}

	return ctx
}

// extract is a helper function to make common get-value, assert-type, return-error-or-value
// operations easier.
func extract(ctx context.Context, k dataloader.StringKey) (*dataloader.Loader, error) {
	ldr, ok := ctx.Value(k).(*dataloader.Loader)
	if !ok {
		return nil, fmt.Errorf("unable to find %s loader on the request context", k)
	}

	return ldr, nil
}
