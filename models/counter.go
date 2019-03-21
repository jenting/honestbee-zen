package models

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/honestbee/Zen/internal/cache"
)

const (
	categoriesCounterForm     = "zen_categories_counter_%s_%s"
	sectionsCounterForm       = "zen_sections_counter_%s_%s"
	articlesCounterForm       = "zen_articles_counter_%s_%s"
	ticketFormsCounterForm    = "zen_ticket_forms_counter"
	categoriesCounterLockForm = "zen_cateogires_counter_lock_%s_%s"
	sectionsCounterLockForm   = "zen_sections_counter_lock_%s_%s"
	articlesCounterLockForm   = "zen_articles_counter_lock_%s_%s"
	ticketFormsLockForm       = "zen_ticket_forms_lock"
)

const (
	maxLockTimeSec = 60
)

type counterService interface {
	PlusOneCategoriesCounter(ctx context.Context, countryCode, locale string) (int, error)
	PlusOneSectionsCounter(ctx context.Context, countryCode, locale string) (int, error)
	PlusOneArticlesCounter(ctx context.Context, countryCode, locale string) (int, error)
	PlusOneTicketFormsCounter(ctx context.Context) (int, error)
	ResetCategoriesCounter(ctx context.Context, countryCode, locale string) error
	ResetSectionsCounter(ctx context.Context, countryCode, locale string) error
	ResetArticlesCounter(ctx context.Context, countryCode, locale string) error
	ResetTicketFormsCounter(ctx context.Context) error
	LockCategoriesCounter(ctx context.Context, countryCode, locale string) (bool, error)
	LockSectionsCounter(ctx context.Context, countryCode, locale string) (bool, error)
	LockArticlesCounter(ctx context.Context, countryCode, locale string) (bool, error)
	LockTicketFormsCounter(ctx context.Context) (bool, error)
	UnlockCategoriesCounter(ctx context.Context, countryCode, locale string) error
	UnlockSectionsCounter(ctx context.Context, countryCode, locale string) error
	UnlockArticlesCounter(ctx context.Context, countryCode, locale string) error
	UnlockTicketFormsCounter(ctx context.Context) error
}

type counterOps struct {
	cache cache.Cache
}

func (c *counterOps) LockCategoriesCounter(ctx context.Context, countryCode, locale string) (bool, error) {
	key := fmt.Sprintf(categoriesCounterLockForm, countryCode, locale)
	reply, err := c.cache.StringDo("SET", key, true, "EX", maxLockTimeSec, "NX", ctx)
	return reply == "OK", errors.Wrapf(err, "models: [LockCategoriesCounter] cache StringDo failed")
}

func (c *counterOps) LockSectionsCounter(ctx context.Context, countryCode, locale string) (bool, error) {
	key := fmt.Sprintf(sectionsCounterLockForm, countryCode, locale)
	reply, err := c.cache.StringDo("SET", key, true, "EX", maxLockTimeSec, "NX", ctx)
	return reply == "OK", errors.Wrapf(err, "models: [LockSectionsCounter] cache StringDo failed")
}

func (c *counterOps) LockArticlesCounter(ctx context.Context, countryCode, locale string) (bool, error) {
	key := fmt.Sprintf(articlesCounterLockForm, countryCode, locale)
	reply, err := c.cache.StringDo("SET", key, true, "EX", maxLockTimeSec, "NX", ctx)
	return reply == "OK", errors.Wrapf(err, "models: [LockArticlesCounter] cache StringDo failed")
}

func (c *counterOps) LockTicketFormsCounter(ctx context.Context) (bool, error) {
	reply, err := c.cache.StringDo("SET", ticketFormsLockForm, true, "EX", maxLockTimeSec, "NX", ctx)
	return reply == "OK", errors.Wrapf(err, "models: [LockTicketFormsCounter] cache StringDo failed")
}

func (c *counterOps) UnlockCategoriesCounter(ctx context.Context, countryCode, locale string) error {
	_, err := c.cache.BoolDo("DEL", fmt.Sprintf(categoriesCounterLockForm, countryCode, locale), ctx)
	return errors.Wrapf(err, "models: [UnlockCategoriesCounter] cache BoolDo failed")
}

func (c *counterOps) UnlockSectionsCounter(ctx context.Context, countryCode, locale string) error {
	_, err := c.cache.BoolDo("DEL", fmt.Sprintf(sectionsCounterLockForm, countryCode, locale), ctx)
	return errors.Wrapf(err, "models: [UnlockSectionsCounter] cache BoolDo failed")
}

func (c *counterOps) UnlockArticlesCounter(ctx context.Context, countryCode, locale string) error {
	_, err := c.cache.BoolDo("DEL", fmt.Sprintf(articlesCounterLockForm, countryCode, locale), ctx)
	return errors.Wrapf(err, "models: [UnlockArticlesCounter] cache BoolDo failed")
}

func (c *counterOps) UnlockTicketFormsCounter(ctx context.Context) error {
	_, err := c.cache.BoolDo("DEL", ticketFormsLockForm, ctx)
	return errors.Wrapf(err, "models: [UnlockTicketFormsCounter] cache BoolDo failed")
}

func (c *counterOps) PlusOneCategoriesCounter(ctx context.Context, countryCode, locale string) (int, error) {
	reply, err := c.cache.IntDo("INCR", fmt.Sprintf(categoriesCounterForm, countryCode, locale), ctx)
	return reply, errors.Wrapf(err, "models: [PlusOneCategoriesCounter] cache IntDo failed")
}

func (c *counterOps) PlusOneSectionsCounter(ctx context.Context, countryCode, locale string) (int, error) {
	reply, err := c.cache.IntDo("INCR", fmt.Sprintf(sectionsCounterForm, countryCode, locale), ctx)
	return reply, errors.Wrapf(err, "models: [PlusOneSectionsCounter] cache IntDo failed")
}

func (c *counterOps) PlusOneArticlesCounter(ctx context.Context, countryCode, locale string) (int, error) {
	reply, err := c.cache.IntDo("INCR", fmt.Sprintf(articlesCounterForm, countryCode, locale), ctx)
	return reply, errors.Wrapf(err, "models: [PlusOneArticlesCounter] cache IntDo failed")
}

func (c *counterOps) PlusOneTicketFormsCounter(ctx context.Context) (int, error) {
	reply, err := c.cache.IntDo("INCR", ticketFormsCounterForm, ctx)
	return reply, errors.Wrapf(err, "models: [PlusOneTicketFormsCounter] cache IntDo failed")
}

func (c *counterOps) ResetCategoriesCounter(ctx context.Context, countryCode, locale string) error {
	_, err := c.cache.StringDo("SET", fmt.Sprintf(categoriesCounterForm, countryCode, locale), 0, ctx)
	return errors.Wrapf(err, "models: [ResetCategoriesCounter] cache StringDo failed")
}

func (c *counterOps) ResetSectionsCounter(ctx context.Context, countryCode, locale string) error {
	_, err := c.cache.StringDo("SET", fmt.Sprintf(sectionsCounterForm, countryCode, locale), 0, ctx)
	return errors.Wrapf(err, "models: [ResetSectionsCounter] cache StringDo failed")
}

func (c *counterOps) ResetArticlesCounter(ctx context.Context, countryCode, locale string) error {
	_, err := c.cache.StringDo("SET", fmt.Sprintf(articlesCounterForm, countryCode, locale), 0, ctx)
	return errors.Wrapf(err, "models: [ResetArticlesCounter] cache StringDo failed")
}

func (c *counterOps) ResetTicketFormsCounter(ctx context.Context) error {
	_, err := c.cache.StringDo("SET", ticketFormsCounterForm, 0, ctx)
	return errors.Wrapf(err, "models: [ResetTicketFormsCounter] cache StringDo failed")
}
