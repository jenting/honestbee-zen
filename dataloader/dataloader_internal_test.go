package dataloader

import (
	"context"
	"io/ioutil"

	"github.com/rs/zerolog"

	"github.com/honestbee/Zen/config"
	"github.com/honestbee/Zen/examiner"
	"github.com/honestbee/Zen/models"
	"github.com/honestbee/Zen/zendesk"
)

var ctx context.Context

func init() {
	logger := zerolog.New(ioutil.Discard)
	ms := &models.MockModels{}
	zend, _ := zendesk.NewZenDesk(&config.Config{
		ZenDesk: &config.ZenDesk{
			RequestTimeoutSec: 20,
			AuthToken:         "ZWRtdW5kLmthb0Bob25lc3RiZWUuY29tL3Rva2VuOmZXdmVMYXVvN0lzQVExQURrbE54ZFVySkIwMWN1aFltTnhVRmVIbE8=",
			HKBaseURL:         "https://honestbeehelp-hk.zendesk.com",
			IDBaseURL:         "https://honestbee-idn.zendesk.com",
			JPBaseURL:         "https://honestbeehelp-jp.zendesk.com",
			MYBaseURL:         "https://honestbee-my.zendesk.com",
			PHBaseURL:         "https://honestbee-ph.zendesk.com",
			SGBaseURL:         "https://honestbeehelp-sg.zendesk.com",
			THBaseURL:         "https://honestbee-th.zendesk.com",
			TWBaseURL:         "https://honestbeehelp-tw.zendesk.com",
		},
	})
	exam, _ := examiner.NewExaminer(&config.Config{
		Examiner: &config.Examiner{
			MaxWorkerSize:          1,
			MaxPoolSize:            2,
			CategoriesRefreshLimit: 1000,
			SectionsRefreshLimit:   1000,
			ArticlesRefreshLimit:   1000,
		},
	}, &logger, ms, zend)

	ctx = Initialize(ms, exam, zend).Attach(context.Background())
}
