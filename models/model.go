package models

import (
	"github.com/pkg/errors"

	"github.com/honestbee/Zen/config"
	"github.com/honestbee/Zen/internal/cache"
	"github.com/honestbee/Zen/internal/db"
)

var (
	// ErrNotFound means sql query got sql.ErrNoRows error.
	ErrNotFound = errors.New("not found")
)

// Service is the interface of all model service.
type Service interface {
	categoriesService
	articlesService
	sectionsService
	ticketFormsService
	ticketFieldsService
	dynamicContentService
	counterService
	dataloaderService
	Close() error
}

// HelpDeskService is the interface of help desk service.
// This interface is defined to avoid handlers package can access counter service.
type HelpDeskService interface {
	categoriesService
	articlesService
	sectionsService
	ticketFormsService
	ticketFieldsService
	dynamicContentService
}

type service struct {
	*categoriesOps
	*sectionsOps
	*articlesOps
	*ticketFormsOps
	*ticketFieldsOps
	*dynamicContentOps
	*counterOps
	*dataloaderOps
	close func() error
}

const (
	counterServiceRedisIndex    = iota // 0
	dataloaderServiceRedisIndex        // 1
)

// New returns a Service instance for operating all model service.
func New(conf *config.Config) (Service, error) {
	d, err := db.NewPostgres(conf)
	if err != nil {
		return nil, errors.Wrapf(err, "model: [New] new postgress failed")
	}

	cc, err := cache.NewRedis(conf, counterServiceRedisIndex)
	if err != nil {
		return nil, errors.Wrapf(err, "model: [New] new redis failed")
	}
	dlc, err := cache.NewRedis(conf, dataloaderServiceRedisIndex)
	if err != nil {
		return nil, errors.Wrapf(err, "model: [New] new redis failed")
	}

	dcOps := &dynamicContentOps{d}
	fieldsOps := &ticketFieldsOps{db: d, dcOps: dcOps}

	return &service{
		categoriesOps:     &categoriesOps{d},
		sectionsOps:       &sectionsOps{d},
		articlesOps:       &articlesOps{d},
		counterOps:        &counterOps{cc},
		dataloaderOps:     &dataloaderOps{dlc},
		ticketFormsOps:    &ticketFormsOps{db: d, fieldsOps: fieldsOps, dcOps: dcOps},
		ticketFieldsOps:   fieldsOps,
		dynamicContentOps: dcOps,
		close: func() error {
			derr := errors.Wrapf(d.Close(), "db close failed")
			ccerr := errors.Wrapf(cc.Close(), "counter cache close failed")
			dlcerr := errors.Wrapf(dlc.Close(), "dataloader cache close failed")
			if derr != nil {
				return derr
			} else if ccerr != nil {
				return ccerr
			} else if dlcerr != nil {
				return dlcerr
			}
			return nil
		},
	}, nil
}

func (s *service) Close() error {
	return errors.Wrapf(s.close(), "model: [Close] close failed")
}
