package examiner

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/honestbee/Zen/config"
	"github.com/honestbee/Zen/models"
	"github.com/honestbee/Zen/zendesk"
)

const (
	categoriesItem  = "categories"
	sectionsItem    = "sections"
	articlesItem    = "articles"
	ticketFormsItem = "ticket_forms"
)

var (
	// ErrAcquireCounterLockFailed means trying to lock cache counter failed,
	// only have two scanior this error may occured:
	// 1. another goroutine is syncing the data
	// 2. another goroutine failed to sync the data and the lock is waiting to be expired
	ErrAcquireCounterLockFailed = errors.New("acquire counter lock failed")
)

type task struct {
	ctx         context.Context
	item        string
	locale      string
	countryCode string
}

type articleTask struct {
	ctx         context.Context
	locale      string
	countryCode string
	articleID   int
}

// Examiner is the structure checking the counter number,
// if the counter number of each cache subject reach the limit,
// it will refresh the database data by reaching zendesk api.
type Examiner struct {
	tasks                   chan interface{}
	wg                      *sync.WaitGroup
	logger                  *zerolog.Logger
	categoriesRefreshLimit  int
	sectionsRefreshLimit    int
	articlesRefreshLimit    int
	ticketFormsRefreshLimit int
	service                 models.Service
	zendesk                 *zendesk.ZenDesk
}

// NewExaminer returns a Examiner instance and runs workers to work.
func NewExaminer(conf *config.Config,
	logger *zerolog.Logger,
	service models.Service,
	zendesk *zendesk.ZenDesk) (*Examiner, error) {

	e := &Examiner{
		tasks:                   make(chan interface{}, conf.Examiner.MaxPoolSize),
		wg:                      new(sync.WaitGroup),
		logger:                  logger,
		service:                 service,
		zendesk:                 zendesk,
		categoriesRefreshLimit:  conf.Examiner.CategoriesRefreshLimit,
		sectionsRefreshLimit:    conf.Examiner.SectionsRefreshLimit,
		articlesRefreshLimit:    conf.Examiner.ArticlesRefreshLimit,
		ticketFormsRefreshLimit: conf.Examiner.TicketFormsRefreshLimit,
	}

	for i := 0; i < conf.Examiner.MaxWorkerSize; i++ {
		e.wg.Add(1)
		go e.worker(i)
	}

	return e, nil
}

func (e *Examiner) worker(workerID int) {
	defer e.wg.Done()
	defer e.logger.Info().Msgf("examiner: [%d]worker return", workerID)

	for eachTask := range e.tasks {
		switch vTask := eachTask.(type) {
		case *task:
			if err := e.work(vTask); err != nil {
				if err != ErrAcquireCounterLockFailed {
					e.logger.Error().Err(err).Fields(map[string]interface{}{
						"item":        vTask.item,
						"countryCode": vTask.countryCode,
						"locale":      vTask.locale,
					}).Msgf("examiner: [%d]worker work failed ", workerID)
				}
			}
		case *articleTask:
			if err := e.articleWork(vTask); err != nil {
				if err != ErrAcquireCounterLockFailed {
					e.logger.Error().Err(err).Fields(map[string]interface{}{
						"countryCode": vTask.countryCode,
						"locale":      vTask.locale,
						"articleID":   vTask.articleID,
					}).Msgf("examiner: [%d]article worker work failed ", workerID)
				}
			}
		}
	}
}

func (e *Examiner) work(task *task) (err error) {
	switch task.item {
	case categoriesItem:
		err = e.categoriesWork(task.ctx, task.countryCode, task.locale)
	case sectionsItem:
		err = e.sectionsWork(task.ctx, task.countryCode, task.locale)
	case articlesItem:
		err = e.articlesWork(task.ctx, task.countryCode, task.locale)
	case ticketFormsItem:
		err = e.ticketFormsWork(task.ctx)
	default:
		err = errors.Errorf("examiner: [work] receive unknown item:%s", task.item)
	}

	return err
}

func (e *Examiner) articleWork(task *articleTask) (err error) {
	return e.articleSync(task.ctx, task.articleID, task.countryCode, task.locale)
}

func (e *Examiner) categoriesWork(ctx context.Context, countryCode, locale string) error {
	count, err := e.service.PlusOneCategoriesCounter(ctx, countryCode, locale)
	if err != nil {
		return errors.Wrapf(err, "examiner: [categoriesWork] service.PlusOneCategoriesCounter failed")
	}

	// refresh limit <= 0: always not sync.
	// refresh limit == 1: always sync.
	// refresh limit > 1: sync when count >= refresh limit.
	if e.categoriesRefreshLimit <= 0 {
		return nil
	}
	if count < e.categoriesRefreshLimit {
		return nil
	}

	if err := e.categoriesSync(ctx, countryCode, locale); err != nil {
		switch err {
		case ErrAcquireCounterLockFailed:
			return ErrAcquireCounterLockFailed
		default:
			return errors.Wrapf(err, "examiner: [categoriesWork] categoriesSync failed")
		}
	}
	return nil
}

func (e *Examiner) categoriesSync(ctx context.Context, countryCode, locale string) error {
	isLock, err := e.service.LockCategoriesCounter(ctx, countryCode, locale)
	if err != nil {
		return errors.Wrapf(err, "examiner: [categoriesSync] service.LockCategoriesCounter failed")
	}
	if !isLock {
		return ErrAcquireCounterLockFailed
	}

	zendeskCategories, err := e.zendesk.ListCategories(ctx, countryCode, locale)
	if err != nil {
		return errors.Wrapf(err, "examiner: [categoriesSync] zendesk.ListCategories failed")
	}

	e.logger.Info().Msgf("examiner: [categoriesSync] pulling from zendesk categories length:%d", len(zendeskCategories))

	if len(zendeskCategories) == 0 {
		return errors.Wrapf(err, "examiner: [categoriesSync] pulling from zendesk categories length is 0")
	}

	categories := make([]*models.Category, len(zendeskCategories))
	for i, zendeskCategory := range zendeskCategories {
		categories[i] = new(models.Category)
		categories[i].CountryCode = countryCode
		categories[i].CreatedAt = zendeskCategory.CreatedAt
		categories[i].Description = zendeskCategory.Description
		categories[i].HTMLURL = zendeskCategory.HTMLURL
		categories[i].ID = zendeskCategory.ID
		categories[i].Locale = zendeskCategory.Locale
		categories[i].Name = zendeskCategory.Name
		categories[i].Outdated = zendeskCategory.Outdated
		categories[i].Position = zendeskCategory.Position
		categories[i].SourceLocale = zendeskCategory.SourceLocale
		categories[i].UpdatedAt = zendeskCategory.UpdatedAt
		categories[i].URL = zendeskCategory.URL
	}

	if err = e.service.SyncWithCategories(ctx, categories, countryCode, locale); err != nil {
		return errors.Wrapf(err, "examiner: [categoriesSync] service.SyncWithCategories failed")
	}
	if err = e.service.CategoriesCacheInvalidate(ctx, countryCode, locale); err != nil {
		return errors.Wrapf(err, "examiner: [categoriesSync] service.CategoriesCacheInvalidate failed")
	}
	if err = e.service.ResetCategoriesCounter(ctx, countryCode, locale); err != nil {
		return errors.Wrapf(err, "examiner: [categoriesSync] service.ResetCategoriesCounter failed")
	}
	if err = e.service.UnlockCategoriesCounter(ctx, countryCode, locale); err != nil {
		return errors.Wrapf(err, "examiner: [categoriesSync] service.UnlockCategoriesCounter failed")
	}

	return nil
}

func (e *Examiner) sectionsWork(ctx context.Context, countryCode, locale string) error {
	count, err := e.service.PlusOneSectionsCounter(ctx, countryCode, locale)
	if err != nil {
		return errors.Wrapf(err, "examiner: [sectionsWork] service.PlusOneSectionsCounter failed")
	}

	// refresh limit <= 0: always not sync.
	// refresh limit == 1: always sync.
	// refresh limit > 1: sync when count >= refresh limit.
	if e.sectionsRefreshLimit <= 0 {
		return nil
	}
	if count < e.sectionsRefreshLimit {
		return nil
	}

	if err = e.sectionsSync(ctx, countryCode, locale); err != nil {
		switch err {
		case ErrAcquireCounterLockFailed:
			return ErrAcquireCounterLockFailed
		default:
			return errors.Wrapf(err, "examiner: [sectionsWork] sectionsSync failed")
		}
	}
	return nil
}

func (e *Examiner) sectionsSync(ctx context.Context, countryCode, locale string) error {
	isLock, err := e.service.LockSectionsCounter(ctx, countryCode, locale)
	if err != nil {
		return errors.Wrapf(err, "examiner: [sectionsSync] service.LockSectionsCounter failed")
	}
	if !isLock {
		return ErrAcquireCounterLockFailed
	}

	zendeskSections, err := e.zendesk.ListSections(ctx, countryCode, locale)
	if err != nil {
		return errors.Wrapf(err, "examiner: [sectionsSync] zendesk.ListSections failed")
	}

	e.logger.Info().Msgf("examiner: [sectionsSync] pulling from zendesk sections length:%d", len(zendeskSections))

	if len(zendeskSections) == 0 {
		return errors.Wrapf(err, "examiner: [sectionsSync] pulling from zendesk sections length is 0")
	}

	sections := make([]*models.Section, len(zendeskSections))
	for i, zendeskSection := range zendeskSections {
		sections[i] = new(models.Section)
		sections[i].CategoryID = zendeskSection.CategoryID
		sections[i].CountryCode = countryCode
		sections[i].CreatedAt = zendeskSection.CreatedAt
		sections[i].Description = zendeskSection.Description
		sections[i].HTMLURL = zendeskSection.HTMLURL
		sections[i].ID = zendeskSection.ID
		sections[i].Locale = zendeskSection.Locale
		sections[i].Name = zendeskSection.Name
		sections[i].Outdated = zendeskSection.Outdated
		sections[i].Position = zendeskSection.Position
		sections[i].SourceLocale = zendeskSection.SourceLocale
		sections[i].UpdatedAt = zendeskSection.UpdatedAt
		sections[i].URL = zendeskSection.URL
	}

	if err = e.service.SyncWithSections(ctx, sections, countryCode, locale); err != nil {
		return errors.Wrapf(err, "examiner: [sectionsSync] service.SyncWithSections failed")
	}
	if err = e.service.SectionsCacheInvalidate(ctx, countryCode, locale); err != nil {
		return errors.Wrapf(err, "examiner: [sectionsSync] service.SectionsCacheInvalidate failed")
	}
	if err = e.service.ResetSectionsCounter(ctx, countryCode, locale); err != nil {
		return errors.Wrapf(err, "examiner: [sectionsSync] service.ResetSectionsCounter failed")
	}
	if err = e.service.UnlockSectionsCounter(ctx, countryCode, locale); err != nil {
		return errors.Wrapf(err, "examiner: [sectionsSync] service.UnlockSectionsCounter failed")
	}

	return nil
}

func (e *Examiner) articlesWork(ctx context.Context, countryCode, locale string) error {
	count, err := e.service.PlusOneArticlesCounter(ctx, countryCode, locale)
	if err != nil {
		return errors.Wrapf(err, "examiner: [articlesWork] service.PlusOneArticlesCounter failed")
	}

	// refresh limit <= 0: always not sync.
	// refresh limit == 1: always sync.
	// refresh limit > 1: sync when count >= refresh limit.
	if e.articlesRefreshLimit <= 0 {
		return nil
	}
	if count < e.articlesRefreshLimit {
		return nil
	}

	if err := e.articlesSync(ctx, countryCode, locale); err != nil {
		switch err {
		case ErrAcquireCounterLockFailed:
			return ErrAcquireCounterLockFailed
		default:
			return errors.Wrapf(err, "examiner: [articlesWork] articlesSync failed")
		}
	}
	return nil
}

func (e *Examiner) articlesSync(ctx context.Context, countryCode, locale string) error {
	isLock, err := e.service.LockArticlesCounter(ctx, countryCode, locale)
	if err != nil {
		return errors.Wrapf(err, "examiner: [articlesSync] service.LockArticlesCounter failed")
	}
	if !isLock {
		return ErrAcquireCounterLockFailed
	}

	zendeskArticles, err := e.zendesk.ListArticles(ctx, countryCode, locale)
	if err != nil {
		return errors.Wrapf(err, "examiner: [articlesSync] zendesk.ListArticles failed")
	}

	e.logger.Info().Msgf("examiner: [articlesSync] pulling from zendesk articles length:%d", len(zendeskArticles))

	if len(zendeskArticles) == 0 {
		return errors.Wrapf(err, "examiner: [articlesSync] pulling from zendesk articles length is 0")
	}

	articles := make([]*models.Article, len(zendeskArticles))
	for i, zendeskArticle := range zendeskArticles {
		articles[i] = new(models.Article)
		articles[i].SectionID = zendeskArticle.SectionID
		articles[i].AuthorID = zendeskArticle.AuthorID
		articles[i].Body = zendeskArticle.Body
		articles[i].CommentsDisable = zendeskArticle.CommentsDisable
		articles[i].CountryCode = countryCode
		articles[i].CreatedAt = zendeskArticle.CreatedAt
		articles[i].Draft = zendeskArticle.Draft
		articles[i].EditedAt = zendeskArticle.EditedAt
		articles[i].HTMLURL = zendeskArticle.HTMLURL
		articles[i].ID = zendeskArticle.ID
		articles[i].LabelNames = zendeskArticle.LabelNames
		articles[i].Locale = zendeskArticle.Locale
		articles[i].Name = zendeskArticle.Name
		articles[i].Outdated = zendeskArticle.Outdated
		articles[i].OutdatedLocales = zendeskArticle.OutdatedLocales
		articles[i].Position = zendeskArticle.Position
		articles[i].Promoted = zendeskArticle.Promoted
		articles[i].SourceLocale = zendeskArticle.SourceLocale
		articles[i].Title = zendeskArticle.Title
		articles[i].UpdatedAt = zendeskArticle.UpdatedAt
		articles[i].URL = zendeskArticle.URL
		articles[i].VoteCount = zendeskArticle.VoteCount
		articles[i].VoteSum = zendeskArticle.VoteSum
	}

	if err = e.service.SyncWithArticles(ctx, articles, countryCode, locale); err != nil {
		return errors.Wrapf(err, "examiner: [articlesSync] service.SyncWithArticles failed")
	}
	if err = e.service.ArticlesCacheInvalidate(ctx, countryCode, locale); err != nil {
		return errors.Wrapf(err, "examiner: [articlesSync] service.ArticlesCacheInvalidate failed")
	}
	if err = e.service.ResetArticlesCounter(ctx, countryCode, locale); err != nil {
		return errors.Wrapf(err, "examiner: [articlesSync] service.ResetArticlesCounter failed")
	}
	if err = e.service.UnlockArticlesCounter(ctx, countryCode, locale); err != nil {
		return errors.Wrapf(err, "examiner: [articlesSync] service.UnlockArticlesCounter failed")
	}

	return nil
}

func (e *Examiner) articleSync(ctx context.Context, articleID int, countryCode, locale string) error {
	zendeskArticle, err := e.zendesk.ShowArticle(ctx, articleID, countryCode, locale)

	if err != nil {
		return errors.Wrapf(err, "examiner: [articleSync] zendesk.ShowArticle failed")
	}

	article := &models.Article{
		SectionID:       zendeskArticle.SectionID,
		ID:              zendeskArticle.ID,
		AuthorID:        zendeskArticle.AuthorID,
		CommentsDisable: zendeskArticle.CommentsDisable,
		Draft:           zendeskArticle.Draft,
		Promoted:        zendeskArticle.Promoted,
		Position:        zendeskArticle.Position,
		VoteSum:         zendeskArticle.VoteSum,
		VoteCount:       zendeskArticle.VoteCount,
		CreatedAt:       zendeskArticle.CreatedAt,
		UpdatedAt:       zendeskArticle.UpdatedAt,
		SourceLocale:    zendeskArticle.SourceLocale,
		Outdated:        zendeskArticle.Outdated,
		OutdatedLocales: zendeskArticle.OutdatedLocales,
		EditedAt:        zendeskArticle.EditedAt,
		LabelNames:      zendeskArticle.LabelNames,
		CountryCode:     countryCode,
		URL:             zendeskArticle.URL,
		HTMLURL:         zendeskArticle.HTMLURL,
		Name:            zendeskArticle.Name,
		Title:           zendeskArticle.Title,
		Body:            zendeskArticle.Body,
		Locale:          zendeskArticle.Locale,
	}

	if err = e.service.SyncWithArticle(ctx, articleID, article, countryCode, locale); err != nil {
		return errors.Wrapf(err, "examiner: [articleSync] service.SyncWithArticle failed")
	}

	return nil
}

func (e *Examiner) ticketFormsWork(ctx context.Context) error {
	count, err := e.service.PlusOneTicketFormsCounter(ctx)
	if err != nil {
		return errors.Wrapf(err, "examiner: [ticketFormsWork] service.PlusOneTicketFormsCounter failed")
	}

	// refresh limit <= 0: always not sync.
	// refresh limit == 1: always sync.
	// refresh limit > 1: sync when count >= refresh limit.
	if e.ticketFormsRefreshLimit <= 0 {
		return nil
	}
	if count < e.ticketFormsRefreshLimit {
		return nil
	}

	if err := e.ticketFormsSync(ctx); err != nil {
		switch err {
		case ErrAcquireCounterLockFailed:
			return ErrAcquireCounterLockFailed
		default:
			return errors.Wrapf(err, "examiner: [ticketFormsWork] ticketFormsSync failed")
		}
	}
	return nil
}

func (e *Examiner) ticketFormsSync(ctx context.Context) error {
	isLock, err := e.service.LockTicketFormsCounter(ctx)
	if err != nil {
		return errors.Wrapf(err, "examiner: [ticketFormsSync] service.LockTicketFormsCounter failed")
	}
	if !isLock {
		return ErrAcquireCounterLockFailed
	}

	zendeskForms, err := e.zendesk.ListTicketForms(ctx)
	if err != nil {
		return errors.Wrapf(err, "examiner: [ticketFormsSync] zendesk.ListTicketForms failed")
	}

	e.logger.Info().Msgf("examiner: [ticketFormsSync] pulling from zendesk ticket forms length:%d", len(zendeskForms))

	if len(zendeskForms) == 0 {
		return errors.Wrapf(err, "examiner: [ticketFormsSync] pulling from zendesk ticket forms length is 0")
	}

	forms := make([]*models.SyncTicketForm, len(zendeskForms))
	for i, zendeskForm := range zendeskForms {
		forms[i] = &models.SyncTicketForm{
			ID:                 zendeskForm.ID,
			URL:                zendeskForm.URL,
			Name:               zendeskForm.Name,
			RawName:            zendeskForm.RawName,
			DisplayName:        zendeskForm.DisplayName,
			RawDisplayName:     zendeskForm.RawDisplayName,
			EndUserVisible:     zendeskForm.EndUserVisible,
			Position:           zendeskForm.Position,
			Active:             zendeskForm.Active,
			InAllBrands:        zendeskForm.InAllBrands,
			RestrictedBrandIDs: zendeskForm.RestrictedBrandIDs,
			TicketFieldIDs:     zendeskForm.TicketFieldIDs,
			CreatedAt:          zendeskForm.CreatedAt,
			UpdatedAt:          zendeskForm.UpdatedAt,
		}
	}

	if err = e.service.SyncWithTicketForms(ctx, forms); err != nil {
		return errors.Wrapf(err, "examiner: [ticketFormsSync] service.SyncWithTicketForms failed")
	}
	if err = e.service.TicketFormCacheInvalidate(ctx); err != nil {
		return errors.Wrapf(err, "examiner: [ticketFormsSync] service.TicketFormCacheInvalidate failed")
	}

	zendeskFields, err := e.zendesk.ListTicketFields(ctx)
	if err != nil {
		return errors.Wrapf(err, "examiner: [ticketFormsSync] zendesk.ListTicketFields failed")
	}

	e.logger.Info().Msgf("examiner: [ticketFormsSync] pulling from zendesk ticket fields length:%d", len(zendeskFields))

	if len(zendeskFields) == 0 {
		return errors.Wrapf(err, "examiner: [ticketFormsSync] pulling from zendesk ticket fields length is 0")
	}

	fields := make([]*models.SyncTicketField, len(zendeskFields))
	for i, zendeskField := range zendeskFields {
		fields[i] = &models.SyncTicketField{
			ID:                  zendeskField.ID,
			URL:                 zendeskField.URL,
			Type:                zendeskField.Type,
			Title:               zendeskField.Title,
			RawTitle:            zendeskField.RawTitle,
			Description:         zendeskField.Description,
			RawDescription:      zendeskField.RawDescription,
			Position:            zendeskField.Position,
			Active:              zendeskField.Active,
			Required:            zendeskField.Required,
			CollapsedForAgents:  zendeskField.CollapsedForAgents,
			RegexpForValidation: zendeskField.RegexpForValidation,
			TitleInPortal:       zendeskField.TitleInPortal,
			RawTitleInPortal:    zendeskField.RawTitleInPortal,
			VisibleInPortal:     zendeskField.VisibleInPortal,
			EditableInPortal:    zendeskField.EditableInPortal,
			RequiredInPortal:    zendeskField.RequiredInPortal,
			Tag:                 zendeskField.Tag,
			CreatedAt:           zendeskField.CreatedAt,
			UpdatedAt:           zendeskField.UpdatedAt,
			Removable:           zendeskField.Removable,
		}
		if zendeskField.CustomFieldOptions == nil {
			fields[i].CustomFieldOptions = []byte("[]")
		} else {
			customOps, err := json.Marshal(zendeskField.CustomFieldOptions)
			if err != nil {
				return errors.Wrapf(err, "examiner: [ticketFormsSync] marshal CustomFieldOptions failed")
			}
			fields[i].CustomFieldOptions = customOps
		}

		if zendeskField.SystemFieldOptions == nil {
			fields[i].SystemFieldOptions = []byte("[]")
		} else {
			systemOps, err := json.Marshal(zendeskField.SystemFieldOptions)
			if err != nil {
				return errors.Wrapf(err, "examiner: [ticketFormsSync] marshal SystemFieldOptions failed")
			}
			fields[i].SystemFieldOptions = systemOps
		}
	}

	if err = e.service.SyncWithTicketFields(ctx, fields); err != nil {
		return errors.Wrapf(err, "examiner: [ticketFormsSync] service.SyncWithTicketFields failed")
	}
	if err = e.service.TicketFieldCacheInvalidate(ctx); err != nil {
		return errors.Wrapf(err, "examiner: [ticketFormsSync] service.TicketFieldCacheInvalidate failed")
	}
	if err = e.service.TicketFieldCustomFieldOptionCacheInvalidate(ctx); err != nil {
		return errors.Wrapf(err, "examiner: [ticketFormsSync] service.TicketFieldCustomFieldOptionCacheInvalidate failed")
	}
	if err = e.service.TicketFieldSystemFieldOptionCacheInvalidate(ctx); err != nil {
		return errors.Wrapf(err, "examiner: [ticketFormsSync] service.TicketFieldSystemFieldOptionCacheInvalidate failed")
	}

	zendeskDCItems, err := e.zendesk.ListDynamicContentItems(ctx)
	if err != nil {
		return errors.Wrapf(err, "examiner: [ticketFormsSync] zendesk.ListDynamicContentItems failed")
	}

	e.logger.Info().Msgf("examiner: [ticketFormsSync] pulling from zendesk dynamic content items length:%d", len(zendeskDCItems))

	if len(zendeskDCItems) == 0 {
		return errors.Wrapf(err, "examiner: [ticketFormsSync] pulling from zendesk dynamic content items length is 0")
	}

	dcItems := make([]*models.SyncDynamicContentItem, len(zendeskDCItems))
	for i, zendeskDCItem := range zendeskDCItems {
		dcItems[i] = &models.SyncDynamicContentItem{
			ID:              zendeskDCItem.ID,
			URL:             zendeskDCItem.URL,
			Name:            zendeskDCItem.Name,
			Placeholder:     zendeskDCItem.Placeholder,
			DefaultLocaleID: zendeskDCItem.DefaultLocaleID,
			Outdated:        zendeskDCItem.Outdated,
			CreatedAt:       zendeskDCItem.CreatedAt,
			UpdatedAt:       zendeskDCItem.UpdatedAt,
		}
		if zendeskDCItem.Variants == nil {
			dcItems[i].Variants = []byte("[]")
		} else {
			variants, err := json.Marshal(zendeskDCItem.Variants)
			if err != nil {
				return errors.Wrapf(err, "examiner: [ticketFormsSync] marshal Variants failed")
			}
			dcItems[i].Variants = variants
		}
	}

	if err = e.service.SyncWithDynamicContentItems(ctx, dcItems); err != nil {
		return errors.Wrapf(err, "examiner: [ticketFormsSync] service.SyncWithDynamicContentItems failed")
	}
	if err = e.service.ResetTicketFormsCounter(ctx); err != nil {
		return errors.Wrapf(err, "examiner: [ticketFormsSync] service.ResetTicketFormsCounter failed")
	}
	if err = e.service.UnlockTicketFormsCounter(ctx); err != nil {
		return errors.Wrapf(err, "examiner: [ticketFormsSync] service.UnlockTicketFormsCounter failed")
	}

	return nil
}

func (e *Examiner) check(ctx context.Context, item, countryCode, locale string) {
	e.tasks <- &task{
		ctx:         ctx,
		item:        item,
		countryCode: countryCode,
		locale:      locale,
	}
}

func (e *Examiner) syncArticle(ctx context.Context, articleID int, countryCode, locale string) {
	e.tasks <- &articleTask{
		ctx:         ctx,
		articleID:   articleID,
		countryCode: countryCode,
		locale:      locale,
	}
}

func (e *Examiner) force(ctx context.Context, item, countryCode, locale string) (err error) {
	switch item {
	case categoriesItem:
		err = e.categoriesSync(ctx, countryCode, locale)
	case sectionsItem:
		err = e.sectionsSync(ctx, countryCode, locale)
	case articlesItem:
		err = e.articlesSync(ctx, countryCode, locale)
	case ticketFormsItem:
		err = e.ticketFormsSync(ctx)
	default:
		err = errors.Errorf("examiner: [force] receive unknown item:%s", item)
	}
	return err
}

// CheckCategories puts the categories check task into worker pool.
func (e *Examiner) CheckCategories(ctx context.Context, countryCode, locale string) {
	e.check(ctx, categoriesItem, countryCode, locale)
}

// CheckSections puts the sections check task into worker pool.
func (e *Examiner) CheckSections(ctx context.Context, countryCode, locale string) {
	e.check(ctx, sectionsItem, countryCode, locale)
}

// CheckArticles puts the articles check task into worker pool.
func (e *Examiner) CheckArticles(ctx context.Context, countryCode, locale string) {
	e.check(ctx, articlesItem, countryCode, locale)
}

// SyncArticle puts the article task into worker pool.
func (e *Examiner) SyncArticle(ctx context.Context, articleID int, countryCode, locale string) {
	e.syncArticle(ctx, articleID, countryCode, locale)
}

// CheckTicketForms puts the ticket forms check task into worker pool.
// it sync:
// 1. ticket forms
// 2. ticket fields
// 3. dynamic content items
func (e *Examiner) CheckTicketForms(ctx context.Context) {
	e.check(ctx, ticketFormsItem, "", "")
}

// ForceSyncCategories force to sync with zendesk categories data.
func (e *Examiner) ForceSyncCategories(ctx context.Context, countryCode, locale string) error {
	return errors.Wrapf(
		e.force(ctx, categoriesItem, countryCode, locale),
		"examiner: [ForceSyncCategories] force to sync categories failed",
	)
}

// ForceSyncSections force to sync with zendesk sections data.
func (e *Examiner) ForceSyncSections(ctx context.Context, countryCode, locale string) error {
	return errors.Wrapf(
		e.force(ctx, sectionsItem, countryCode, locale),
		"examiner: [ForceSyncSections] force to sync sections failed",
	)
}

// ForceSyncArticles force to sync with zendesk articles data.
func (e *Examiner) ForceSyncArticles(ctx context.Context, countryCode, locale string) error {
	return errors.Wrapf(
		e.force(ctx, articlesItem, countryCode, locale),
		"examiner: [ForceSyncArticles] force to sync articles failed",
	)
}

// ForceSyncTicketForms force to sync with zendesk ticket forms data.
// it sync:
// 1. ticket forms
// 2. ticket fields
// 3. dynamic content items
func (e *Examiner) ForceSyncTicketForms(ctx context.Context) error {
	return errors.Wrapf(
		e.force(ctx, ticketFormsItem, "", ""),
		"examiner: [ForceSyncTicketForms] force to sync ticket forms failed",
	)
}

// Close let the gone out goroutine to stop it self.
func (e *Examiner) Close() error {
	close(e.tasks)
	e.wg.Wait()
	return nil
}
