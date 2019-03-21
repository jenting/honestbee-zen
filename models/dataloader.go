package models

import (
	"context"
	"fmt"

	"github.com/garyburd/redigo/redis"
	"github.com/pkg/errors"

	"github.com/honestbee/Zen/internal/cache"
)

const (
	categoriesDataloaderForm         = "zen_categories_dataloader_%s_%s_%s"
	sectionsDataloaderForm           = "zen_sections_dataloader_%s_%s_%s"
	articlesDataloaderForm           = "zen_articles_dataloader_%s_%s_%s"
	ticketFormDataloaderForm         = "zen_ticket_form_dataloader_%s"
	ticketFieldsDataloaderForm       = "zen_ticket_fields_dataloader_%s"
	ticketFieldCustomFieldOptionForm = "zen_ticket_field_custom_field_option_%s"
	ticketFieldSystemFieldOptionForm = "zen_ticket_field_system_field_option_%s"
)

const (
	dataloaderTTLSec = 3600
)

type dataloaderService interface {
	CategoriesCacheGet(ctx context.Context, key, countryCode, locale string) (string, bool)
	CategoriesCacheSet(ctx context.Context, key, value, countryCode, locale string) (bool, error)
	CategoriesCacheInvalidate(ctx context.Context, countryCode, locale string) error
	SectionsCacheGet(ctx context.Context, key, countryCode, locale string) (string, bool)
	SectionsCacheSet(ctx context.Context, key, value, countryCode, locale string) (bool, error)
	SectionsCacheInvalidate(ctx context.Context, countryCode, locale string) error
	ArticlesCacheGet(ctx context.Context, key, countryCode, locale string) (string, bool)
	ArticlesCacheSet(ctx context.Context, key, value, countryCode, locale string) (bool, error)
	ArticlesCacheInvalidate(ctx context.Context, countryCode, locale string) error
	TicketFormCacheGet(ctx context.Context, key string) (string, bool)
	TicketFormCacheSet(ctx context.Context, key, value string) (bool, error)
	TicketFormCacheInvalidate(ctx context.Context) error
	TicketFieldCacheGet(ctx context.Context, key string) (string, bool)
	TicketFieldCacheSet(ctx context.Context, key, value string) (bool, error)
	TicketFieldCacheInvalidate(ctx context.Context) error
	TicketFieldCustomFieldOptionCacheGet(ctx context.Context, key string) (string, bool)
	TicketFieldCustomFieldOptionCacheSet(ctx context.Context, key, value string) (bool, error)
	TicketFieldCustomFieldOptionCacheInvalidate(ctx context.Context) error
	TicketFieldSystemFieldOptionCacheGet(ctx context.Context, key string) (string, bool)
	TicketFieldSystemFieldOptionCacheSet(ctx context.Context, key, value string) (bool, error)
	TicketFieldSystemFieldOptionCacheInvalidate(ctx context.Context) error
}

type dataloaderOps struct {
	cache cache.Cache
}

func (d *dataloaderOps) CategoriesCacheGet(ctx context.Context, key, countryCode, locale string) (string, bool) {
	key = fmt.Sprintf(categoriesDataloaderForm, countryCode, locale, key)
	reply, err := d.cache.StringDo("GET", key, ctx)
	if err == redis.ErrNil {
		return "", false
	}

	// Update TTL.
	d.cache.StringDo("SETEX", key, dataloaderTTLSec, reply, ctx)
	return reply, true
}

func (d *dataloaderOps) CategoriesCacheSet(ctx context.Context, key, value, countryCode, locale string) (bool, error) {
	reply, err := d.cache.StringDo("SET", fmt.Sprintf(categoriesDataloaderForm, countryCode, locale, key), value, "EX", dataloaderTTLSec, "NX", ctx)
	return reply == "OK", errors.Wrapf(err, "models: [CategoriesCacheSet] cache StringDo failed")
}

func (d *dataloaderOps) CategoriesCacheInvalidate(ctx context.Context, countryCode, locale string) error {
	replys, err := d.cache.StringsDo("KEYS", fmt.Sprintf(categoriesDataloaderForm, countryCode, locale, "*"), ctx)
	if err != nil {
		return err
	}

	for _, reply := range replys {
		d.cache.StringDo("DEL", reply, ctx)
	}
	return nil
}

func (d *dataloaderOps) SectionsCacheGet(ctx context.Context, key, countryCode, locale string) (string, bool) {
	key = fmt.Sprintf(sectionsDataloaderForm, countryCode, locale, key)
	reply, err := d.cache.StringDo("GET", key, ctx)
	if err == redis.ErrNil {
		return "", false
	}

	// Update TTL.
	d.cache.StringDo("SETEX", key, dataloaderTTLSec, reply, ctx)
	return reply, true
}

func (d *dataloaderOps) SectionsCacheSet(ctx context.Context, key, value, countryCode, locale string) (bool, error) {
	reply, err := d.cache.StringDo("SET", fmt.Sprintf(sectionsDataloaderForm, countryCode, locale, key), value, "EX", dataloaderTTLSec, "NX", ctx)
	return reply == "OK", errors.Wrapf(err, "models: [SectionsCacheSet] cache StringDo failed")
}

func (d *dataloaderOps) SectionsCacheInvalidate(ctx context.Context, countryCode, locale string) error {
	replys, err := d.cache.StringsDo("KEYS", fmt.Sprintf(sectionsDataloaderForm, countryCode, locale, "*"), ctx)
	if err != nil {
		return err
	}

	for _, reply := range replys {
		d.cache.StringDo("DEL", reply, ctx)
	}
	return nil
}

func (d *dataloaderOps) ArticlesCacheGet(ctx context.Context, key, countryCode, locale string) (string, bool) {
	key = fmt.Sprintf(articlesDataloaderForm, countryCode, locale, key)
	reply, err := d.cache.StringDo("GET", key, ctx)
	if err == redis.ErrNil {
		return "", false
	}

	// Update TTL.
	d.cache.StringDo("SETEX", key, dataloaderTTLSec, reply, ctx)
	return reply, true
}

func (d *dataloaderOps) ArticlesCacheSet(ctx context.Context, key, value, countryCode, locale string) (bool, error) {
	reply, err := d.cache.StringDo("SET", fmt.Sprintf(articlesDataloaderForm, countryCode, locale, key), value, "EX", dataloaderTTLSec, "NX", ctx)
	return reply == "OK", errors.Wrapf(err, "models: [ArticlesCacheSet] cache StringDo failed")
}

func (d *dataloaderOps) ArticlesCacheInvalidate(ctx context.Context, countryCode, locale string) error {
	replys, err := d.cache.StringsDo("KEYS", fmt.Sprintf(articlesDataloaderForm, countryCode, locale, "*"), ctx)
	if err != nil {
		return err
	}

	for _, reply := range replys {
		d.cache.StringDo("DEL", reply, ctx)
	}
	return nil
}

func (d *dataloaderOps) TicketFormCacheGet(ctx context.Context, key string) (string, bool) {
	key = fmt.Sprintf(ticketFormDataloaderForm, key)
	reply, err := d.cache.StringDo("GET", key, ctx)
	if err == redis.ErrNil {
		return "", false
	}

	// Update TTL.
	d.cache.StringDo("SETEX", key, dataloaderTTLSec, reply, ctx)
	return reply, true
}

func (d *dataloaderOps) TicketFormCacheSet(ctx context.Context, key, value string) (bool, error) {
	reply, err := d.cache.StringDo("SET", fmt.Sprintf(ticketFormDataloaderForm, key), value, "EX", dataloaderTTLSec, "NX", ctx)
	return reply == "OK", errors.Wrapf(err, "models: [TicketFormCacheSet] cache StringDo failed")
}

func (d *dataloaderOps) TicketFormCacheInvalidate(ctx context.Context) error {
	replys, err := d.cache.StringsDo("KEYS", fmt.Sprintf(ticketFormDataloaderForm, "*"), ctx)
	if err != nil {
		return err
	}

	for _, reply := range replys {
		d.cache.StringDo("DEL", reply, ctx)
	}
	return nil
}

func (d *dataloaderOps) TicketFieldCacheGet(ctx context.Context, key string) (string, bool) {
	key = fmt.Sprintf(ticketFieldsDataloaderForm, key)
	reply, err := d.cache.StringDo("GET", key, ctx)
	if err == redis.ErrNil {
		return "", false
	}

	// Update TTL.
	d.cache.StringDo("SETEX", key, dataloaderTTLSec, reply, ctx)
	return reply, true
}

func (d *dataloaderOps) TicketFieldCacheSet(ctx context.Context, key, value string) (bool, error) {
	reply, err := d.cache.StringDo("SET", fmt.Sprintf(ticketFieldsDataloaderForm, key), value, "EX", dataloaderTTLSec, "NX", ctx)
	return reply == "OK", errors.Wrapf(err, "models: [TicketFieldCacheSet] cache StringDo failed")
}

func (d *dataloaderOps) TicketFieldCacheInvalidate(ctx context.Context) error {
	replys, err := d.cache.StringsDo("KEYS", fmt.Sprintf(ticketFieldsDataloaderForm, "*"), ctx)
	if err != nil {
		return err
	}

	for _, reply := range replys {
		d.cache.StringDo("DEL", reply, ctx)
	}
	return nil
}

func (d *dataloaderOps) TicketFieldCustomFieldOptionCacheGet(ctx context.Context, key string) (string, bool) {
	key = fmt.Sprintf(ticketFieldCustomFieldOptionForm, key)
	reply, err := d.cache.StringDo("GET", key, ctx)
	if err == redis.ErrNil {
		return "", false
	}

	// Update TTL.
	d.cache.StringDo("SETEX", key, dataloaderTTLSec, reply, ctx)
	return reply, true
}

func (d *dataloaderOps) TicketFieldCustomFieldOptionCacheSet(ctx context.Context, key, value string) (bool, error) {
	reply, err := d.cache.StringDo("SET", fmt.Sprintf(ticketFieldCustomFieldOptionForm, key), value, "EX", dataloaderTTLSec, "NX", ctx)
	return reply == "OK", errors.Wrapf(err, "models: [TicketFieldCustomFieldOptionCacheSet] cache StringDo failed")
}

func (d *dataloaderOps) TicketFieldCustomFieldOptionCacheInvalidate(ctx context.Context) error {
	replys, err := d.cache.StringsDo("KEYS", fmt.Sprintf(ticketFieldCustomFieldOptionForm, "*"), ctx)
	if err != nil {
		return err
	}

	for _, reply := range replys {
		d.cache.StringDo("DEL", reply, ctx)
	}
	return nil
}

func (d *dataloaderOps) TicketFieldSystemFieldOptionCacheGet(ctx context.Context, key string) (string, bool) {
	key = fmt.Sprintf(ticketFieldSystemFieldOptionForm, key)
	reply, err := d.cache.StringDo("GET", key, ctx)
	if err == redis.ErrNil {
		return "", false
	}

	// Update TTL.
	d.cache.StringDo("SETEX", key, dataloaderTTLSec, reply, ctx)
	return reply, true
}

func (d *dataloaderOps) TicketFieldSystemFieldOptionCacheSet(ctx context.Context, key, value string) (bool, error) {
	reply, err := d.cache.StringDo("SET", fmt.Sprintf(ticketFieldSystemFieldOptionForm, key), value, "EX", dataloaderTTLSec, "NX", ctx)
	return reply == "OK", errors.Wrapf(err, "models: [TicketFieldSystemFieldOptionCacheSet] cache StringDo failed")
}

func (d *dataloaderOps) TicketFieldSystemFieldOptionCacheInvalidate(ctx context.Context) error {
	replys, err := d.cache.StringsDo("KEYS", fmt.Sprintf(ticketFieldSystemFieldOptionForm, "*"), ctx)
	if err != nil {
		return err
	}

	for _, reply := range replys {
		d.cache.StringDo("DEL", reply, ctx)
	}
	return nil
}
