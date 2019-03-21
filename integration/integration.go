package integration

import (
	"context"
	"io/ioutil"
	"log"
	"net/http/httptest"
	"os"
	"os/exec"

	"github.com/rs/zerolog"

	"github.com/honestbee/Zen/config"
	"github.com/honestbee/Zen/examiner"
	"github.com/honestbee/Zen/models"
	"github.com/honestbee/Zen/resolvers"
	"github.com/honestbee/Zen/router"
	"github.com/honestbee/Zen/zendesk"
)

var (
	conf *config.Config
)

const (
	timeFormat              = "2006-01-02 15:04:05"
	categoriesRefreshLimit  = 10
	sectionsRefreshLimit    = 10
	articlesRefreshLimit    = 10
	ticketFormsRefreshLimit = 10
)

func init() {
	var err error
	conf, err = config.New()
	if err != nil {
		log.Fatalf("new config failed:%v", err)
	}
	conf.Examiner.CategoriesRefreshLimit = categoriesRefreshLimit
	conf.Examiner.SectionsRefreshLimit = sectionsRefreshLimit
	conf.Examiner.ArticlesRefreshLimit = articlesRefreshLimit
	conf.Examiner.TicketFormsRefreshLimit = ticketFormsRefreshLimit

	if err = os.Setenv("PGPASSWORD", conf.Database.Password); err != nil {
		log.Fatalf("set up postgres password failed:%v", err)
	}
	if err = resetDB(); err != nil {
		log.Fatalf("reset db command run failed:%v", err)
	}
}

func resetDB() error {
	return exec.Command(
		"psql",
		"-U", conf.Database.User,
		"-h", conf.Database.Host,
		"-d", conf.Database.DBName,
		"-f", "fakeData.sql",
	).Run()
}

func newService() models.Service {
	service, err := models.New(conf)
	if err != nil {
		log.Fatalf("new models failed:%v", err)
	}
	return service
}

type tserver struct {
	*httptest.Server
	exam    *examiner.Examiner
	service models.Service
	zend    *zendesk.ZenDesk
}

func newTserver() *tserver {
	logger := zerolog.New(ioutil.Discard)
	zend, err := zendesk.NewZenDesk(conf)
	if err != nil {
		log.Fatalf("new zendesk failed:%v", err)
	}
	service, err := models.New(conf)
	if err != nil {
		log.Fatalf("new models failed:%v", err)
	}
	exam, err := examiner.NewExaminer(conf, &logger, service, zend)
	if err != nil {
		log.Fatalf("new examiner failed:%v", err)
	}
	resolver, err := resolvers.New(conf, &logger, service, exam, zend)
	if err != nil {
		log.Fatalf("new graphql resolver failed")
	}
	h, err := router.New(conf, &logger, service, exam, zend, resolver)
	if err != nil {
		log.Fatalf("new router failed:%v", err)
	}
	server := httptest.NewServer(h)
	return &tserver{
		Server:  server,
		exam:    exam,
		service: service,
		zend:    zend,
	}
}

func (t *tserver) resetAllCounter() {
	t.service.ResetArticlesCounter(context.Background(), "tw", "en-us")
	t.service.ResetArticlesCounter(context.Background(), "tw", "zh-tw")
	t.service.ResetArticlesCounter(context.Background(), "sg", "en-us")
	t.service.ResetArticlesCounter(context.Background(), "sg", "zh-cn")
	t.service.ResetSectionsCounter(context.Background(), "tw", "en-us")
	t.service.ResetSectionsCounter(context.Background(), "tw", "zh-tw")
	t.service.ResetSectionsCounter(context.Background(), "sg", "en-us")
	t.service.ResetSectionsCounter(context.Background(), "sg", "zh-cn")
	t.service.ResetCategoriesCounter(context.Background(), "tw", "en-us")
	t.service.ResetCategoriesCounter(context.Background(), "tw", "zh-tw")
	t.service.ResetCategoriesCounter(context.Background(), "sg", "en-us")
	t.service.ResetCategoriesCounter(context.Background(), "sg", "zh-cn")
	t.service.ResetTicketFormsCounter(context.Background())
}

func (t *tserver) dataloaderCacheInvalidateAll() {
	t.service.CategoriesCacheInvalidate(context.Background(), "sg", "en-us")
	t.service.CategoriesCacheInvalidate(context.Background(), "sg", "zh-cn")
	t.service.CategoriesCacheInvalidate(context.Background(), "tw", "en-us")
	t.service.CategoriesCacheInvalidate(context.Background(), "tw", "zh-tw")
	t.service.SectionsCacheInvalidate(context.Background(), "tw", "en-us")
	t.service.SectionsCacheInvalidate(context.Background(), "tw", "zh-tw")
	t.service.SectionsCacheInvalidate(context.Background(), "sg", "en-us")
	t.service.SectionsCacheInvalidate(context.Background(), "sg", "zh-cn")
	t.service.ArticlesCacheInvalidate(context.Background(), "tw", "en-us")
	t.service.ArticlesCacheInvalidate(context.Background(), "tw", "zh-tw")
	t.service.ArticlesCacheInvalidate(context.Background(), "sg", "en-us")
	t.service.ArticlesCacheInvalidate(context.Background(), "sg", "zh-cn")
	t.service.TicketFormCacheInvalidate(context.Background())
	t.service.TicketFieldCacheInvalidate(context.Background())
	t.service.TicketFieldCustomFieldOptionCacheInvalidate(context.Background())
	t.service.TicketFieldSystemFieldOptionCacheInvalidate(context.Background())
}

func (t *tserver) closeAll() {
	t.dataloaderCacheInvalidateAll()
	t.resetAllCounter()
	t.Server.Close()
	t.exam.Close()
	t.service.Close()
}
