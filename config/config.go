package config

import (
	"flag"
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

var (
	// Version is the app version number at build time
	Version = "No Version Provided"
)

// HTTP is the configuration of http server.
type HTTP struct {
	ReadTimeoutSec  int    `yaml:"read_timeout_sec"`
	WriteTimeoutSec int    `yaml:"write_timeout_sec"`
	IdleTimeoutSec  int    `yaml:"idle_timeout_sec"`
	ListenAddr      string `yaml:"listen_addr"`
	BasicAuthUser   string `yaml:"basic_auth_user"`
	BasicAuthPwd    string `yaml:"basic_auth_pwd"`
}

// Database is the database configuration.
type Database struct {
	MaxIdle                  int    `yaml:"max_idle"`
	MaxActive                int    `yaml:"max_active"`
	ConnectTimeoutSec        int    `yaml:"connect_timeout_sec"`
	ReadTimeoutSec           int    `yaml:"read_timeout_sec"`
	WriteTimeoutSec          int    `yaml:"write_timeout_sec"`
	TransactionMaxTimeoutSec int    `yaml:"transaction_max_timeout_sec"`
	Host                     string `yaml:"host"`
	Port                     string `yaml:"port"`
	User                     string `yaml:"user"`
	Password                 string `yaml:"password"`
	DBName                   string `yaml:"db_name"`
}

// ZenDesk is the configurations for zendesk package.
type ZenDesk struct {
	RequestTimeoutSec int    `yaml:"request_timeout_sec"`
	AuthToken         string `yaml:"auth_token"`
	SGBaseURL         string `yaml:"sg_base_url"`
	TWBaseURL         string `yaml:"tw_base_url"`
	HKBaseURL         string `yaml:"hk_base_url"`
	THBaseURL         string `yaml:"th_base_url"`
	MYBaseURL         string `yaml:"my_base_url"`
	JPBaseURL         string `yaml:"jp_base_url"`
	PHBaseURL         string `yaml:"ph_base_url"`
	IDBaseURL         string `yaml:"id_base_url"`
}

// Cache is the cache configuration.
type Cache struct {
	MaxIdle           int    `yaml:"max_idle"`
	MaxActive         int    `yaml:"max_active"`
	IdleTimeoutSec    int    `yaml:"idle_timeout_sec"`
	Wait              bool   `yaml:"wait"`
	ConnectTimeoutSec int    `yaml:"connect_timeout_sec"`
	ReadTimeoutSec    int    `yaml:"read_timeout_sec"`
	WriteTimeoutSec   int    `yaml:"write_timeout_sec"`
	Host              string `yaml:"host"`
	Port              string `yaml:"port"`
	Password          string `yaml:"password"`
}

// Examiner is the examiner package configurations.
type Examiner struct {
	MaxWorkerSize           int `yaml:"max_worker_size"`
	MaxPoolSize             int `yaml:"max_pool_size"`
	CategoriesRefreshLimit  int `yaml:"categories_refresh_limit"`
	SectionsRefreshLimit    int `yaml:"sections_refresh_limit"`
	ArticlesRefreshLimit    int `yaml:"articles_refresh_limit"`
	TicketFormsRefreshLimit int `yaml:"ticket_forms_refresh_limit"`
}

// GraphQL is the GraphQL package configurations.
type GraphQL struct {
	MaxDepth       int `yaml:"max_depth"`
	MaxParallelism int `yaml:"max_parallelism"`
}

// Datadog is the Datadog package configurations.
type Datadog struct {
	Enable bool   `yaml:"enable"`
	Debug  bool   `yaml:"debug"`
	Env    string `yaml:"env"`
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
}

// GRPC is the gRPC package configurations.
type GRPC struct {
	ListenAddr string `yaml:"listen_addr"`
}

// Config is the main configuration for Zen server.
type Config struct {
	HTTP     *HTTP     `yaml:"http"`
	Database *Database `yaml:"database"`
	ZenDesk  *ZenDesk  `yaml:"zendesk"`
	Cache    *Cache    `yaml:"cache"`
	Examiner *Examiner `yaml:"examiner"`
	GraphQL  *GraphQL  `yaml:"graphql"`
	Datadog  *Datadog  `yaml:"datadog"`
	GRPC     *GRPC     `yaml:"grpc"`
}

// New returns a Config instance.
func New() (*Config, error) {
	c := &Config{
		HTTP:     &HTTP{},
		Database: &Database{},
		Cache:    &Cache{},
		Examiner: &Examiner{},
		ZenDesk:  &ZenDesk{},
		GraphQL:  &GraphQL{},
		Datadog:  &Datadog{},
		GRPC:     &GRPC{},
	}

	path := flag.String("config_path", "env.yml", "config file path, if provided will replace flag setting values")
	flag.StringVar(&c.HTTP.ListenAddr, "http_listen_addr", ":8080", "http server listening address")
	flag.IntVar(&c.HTTP.IdleTimeoutSec, "http_idle_timeout_sec", 1200, "http idle timeout second")
	flag.IntVar(&c.HTTP.ReadTimeoutSec, "http_read_timeout_sec", 30, "http read timeout second")
	flag.IntVar(&c.HTTP.WriteTimeoutSec, "http_write_timeout_sec", 60, "http write timeout second")
	flag.StringVar(&c.HTTP.BasicAuthUser, "http_basic_auth_user", "admin", "basic auth user")
	flag.StringVar(&c.HTTP.BasicAuthPwd, "http_basic_auth_pwd", "33456783345678", "basic auth password")
	flag.IntVar(&c.Database.MaxIdle, "db_max_idle", 500, "database max idle")
	flag.IntVar(&c.Database.MaxActive, "db_max_active", 1000, "database max active")
	flag.IntVar(&c.Database.ConnectTimeoutSec, "db_connect_timeout_sec", 5, "database connect timeout second")
	flag.IntVar(&c.Database.ReadTimeoutSec, "db_read_timeout_sec", 10, "database read timeout second")
	flag.IntVar(&c.Database.WriteTimeoutSec, "db_write_timeout_sec", 15, "database write timeout second")
	flag.IntVar(&c.Database.TransactionMaxTimeoutSec, "db_transaction_max_timeout_sec", 60, "database transaction max timeout second")
	flag.StringVar(&c.Database.Host, "db_host", "localhost", "database host")
	flag.StringVar(&c.Database.Port, "db_port", "5432", "database port")
	flag.StringVar(&c.Database.User, "db_user", "root", "database user")
	flag.StringVar(&c.Database.Password, "db_password", "", "database password")
	flag.StringVar(&c.Database.DBName, "db_dbname", "", "database db name")
	flag.IntVar(&c.ZenDesk.RequestTimeoutSec, "zendesk_request_timeout_sec", 10, "zendesk api http request timeout")
	flag.StringVar(&c.ZenDesk.AuthToken, "zendesk_auth_token", "", "zendesk api authorization token")
	flag.StringVar(&c.ZenDesk.HKBaseURL, "zendesk_hk_base_url", "https://honestbeehelp-hk.zendesk.com", "zendesk hk base url")
	flag.StringVar(&c.ZenDesk.IDBaseURL, "zendesk_id_base_url", "https://honestbee-idn.zendesk.com", "zendesk id base url")
	flag.StringVar(&c.ZenDesk.JPBaseURL, "zendesk_jp_base_url", "https://honestbeehelp-jp.zendesk.com", "zendesk jp base url")
	flag.StringVar(&c.ZenDesk.MYBaseURL, "zendesk_my_base_url", "https://honestbee-my.zendesk.com", "zendesk my base url")
	flag.StringVar(&c.ZenDesk.PHBaseURL, "zendesk_ph_base_url", "https://honestbee-ph.zendesk.com", "zendesk ph base url")
	flag.StringVar(&c.ZenDesk.SGBaseURL, "zendesk_sg_base_url", "https://honestbeehelp-sg.zendesk.com", "zendesk sg base url")
	flag.StringVar(&c.ZenDesk.THBaseURL, "zendesk_th_base_url", "https://honestbee-th.zendesk.com", "zendesk th base url")
	flag.StringVar(&c.ZenDesk.TWBaseURL, "zendesk_tw_base_url", "https://honestbeehelp-tw.zendesk.com", "zendesk tw base url")
	flag.IntVar(&c.Cache.MaxIdle, "cache_max_idle", 500, "cache max idle")
	flag.IntVar(&c.Cache.MaxActive, "cache_max_active", 1000, "cache max active")
	flag.IntVar(&c.Cache.IdleTimeoutSec, "cache_idle_timeout_sec", 1200, "close connections after remaining idle for this duration")
	flag.BoolVar(&c.Cache.Wait, "cache_wait", false, "if true and the pool is at the MaxActive limit then Get() waits for a connection to be returned to the pool before returning")
	flag.IntVar(&c.Cache.ConnectTimeoutSec, "cache_connect_timeout_sec", 5, "cache connect timeout second")
	flag.IntVar(&c.Cache.ReadTimeoutSec, "cache_read_timeout_sec", 10, "cache read timeout second")
	flag.IntVar(&c.Cache.WriteTimeoutSec, "cache_write_timeout_sec", 15, "cache write timeout second")
	flag.StringVar(&c.Cache.Host, "cache_host", "127.0.0.1", "cache host")
	flag.StringVar(&c.Cache.Port, "cache_port", "6379", "cache port")
	flag.StringVar(&c.Cache.Password, "cache_password", "", "cache password")
	flag.IntVar(&c.Examiner.MaxWorkerSize, "examiner_max_worker_size", 100, "examiner max worker size")
	flag.IntVar(&c.Examiner.MaxPoolSize, "examiner_max_pool_size", 200, "examiner max pool size")
	flag.IntVar(&c.Examiner.CategoriesRefreshLimit, "examiner_categories_refresh_limit", 0, "examiner categories refresh limit")
	flag.IntVar(&c.Examiner.SectionsRefreshLimit, "examiner_sections_refresh_limit", 0, "examiner sections refresh limit")
	flag.IntVar(&c.Examiner.ArticlesRefreshLimit, "examiner_articles_refresh_limit", 0, "examiner articles refresh limit")
	flag.IntVar(&c.Examiner.TicketFormsRefreshLimit, "examiner_ticket_forms_refresh_limit", 0, "examiner ticket forms refresh limit")
	flag.IntVar(&c.GraphQL.MaxDepth, "graphql_max_depth", 13, "max field nesting depth in a query")
	flag.IntVar(&c.GraphQL.MaxParallelism, "graphql_max_parallelism", 10, "max number of resolvers per request allowed to run in parallel")
	flag.BoolVar(&c.Datadog.Enable, "datadog_enable", true, "datadog enable")
	flag.BoolVar(&c.Datadog.Debug, "datadog_debug", false, "datadog debug")
	flag.StringVar(&c.Datadog.Env, "datadog_env", "development", "datadog environment (development/staging/production)")
	flag.StringVar(&c.Datadog.Host, "datadog_host", "localhost", "datadog host")
	flag.StringVar(&c.Datadog.Port, "datadog_port", "8126", "datadog port")
	flag.StringVar(&c.GRPC.ListenAddr, "grpc_listen_addr", ":50051", "grpc server listening address")

	flag.Parse()

	if *path != "" {
		yFile, err := ioutil.ReadFile(*path)
		if err != nil {
			return nil, errors.Wrapf(err, "config: [New] ioutil read path:%q failed", *path)
		}

		if err = yaml.Unmarshal(yFile, c); err != nil {
			return nil, errors.Wrapf(err, "config: [New] yaml unmarshal failed")
		}
	}

	return c, nil
}
