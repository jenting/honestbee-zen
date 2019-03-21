# Zen
zendesk api proxy service to avoid rate limit

[doc](https://honestbee.atlassian.net/wiki/spaces/EN/pages/302613609/2018+Q1+Zendesk+FAQ+micro-service)

- [Zen](#zen)
    - [Preparations](#preparations)
        - [Install Golang](#install-golang)
        - [Checking Golang Version](#checking-golang-version)
        - [Install Dependencies](#install-dependencies)
        - [Setup Config Variables](#setup-config-variables)
        - [Install Cache](#install-cache)
        - [Database](#database)
            - [Setup Database Env](#setup-database-env)
            - [Install Database](#install-database)
            - [Database Migration](#database-migration)
            - [Install Goose](#install-goose)
            - [Goose version (Print the current version of the database)](#goose-version-print-the-current-version-of-the-database)
            - [Goose up (Apply all available migrations)](#goose-up-apply-all-available-migrations)
            - [Goose down (Roll back a single migration from the current version)](#goose-down-roll-back-a-single-migration-from-the-current-version)
            - [Goose status (Dump the migration status for the current DB)](#goose-status-dump-the-migration-status-for-the-current-db)
        - [Datadog Agent](#Datadog-agent)
        - [Testing](#testing)
        - [Integration Testing](#integration-testing)
    - [DevOps](#devops)
        - [Build Time Setup App Version](#build-time-setup-app-version)
        - [Check Server Status](#check-server-status)
        - [Interact With cc-test-reporter](#interact-with-cc-test-reporter)

## Preparations

### Install Golang
please follow the instruction here for [mac](https://golang.org/doc/install)
and install the latest version.

### Checking Golang Version
```bash
go version >= go1.11
1. go test multi package converage issue, fixed on go1.10
2. go version >= go1.11
```

### Install Dependencies
using go built-in module, after installation, just type
```bash
go mod vendor
go mod verify
```
then all done.


### Setup Config Variables
the logic of config is 
```bash
if config_path flag has been given
then using the given config file as config variables

else if config_path flag is empty
then using input flags as config variables

else 
then using default flags as config variables

```

the flags

| variable name                       | default value                               | description                                                                                                                  |
| ----------------------------------- | ------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| config_path                         | ""                                          | leave it empty will using flags as config variables                                                                          |
| http_listen_addr                    | ":8080"                                     | server listening address                                                                                                     |
| http_idle_timeout_sec               | 1200                                        | server http idle timeout in seconds                                                                                          |
| http_read_timeout_sec               | 30                                          | server http read timeout in seconds                                                                                          |
| http_write_timeout_sec              | 60                                          | server http write timeout in seconds                                                                                         |
| db_max_idle                         | 500                                         | database max idle                                                                                                            |
| db_max_active                       | 1000                                        | database max active                                                                                                          |
| db_connect_timeout_sec              | 5                                           | database connect timeout in seconds                                                                                          |
| db_read_timeout_sec                 | 10                                          | database read timeout in seconds                                                                                             |
| db_write_timeout_sec                | 15                                          | database write timeout in seconds                                                                                            |
| db_transaction_max_timeout_sec      | 60                                          | database transaction max timeout in seconds                                                                                  |
| db_host                             | "localhost"                                 | database host                                                                                                                |
| db_port                             | "3306"                                      | database password                                                                                                            |
| db_user                             | "root"                                      | database user                                                                                                                |
| db_password                         | ""                                          | database password                                                                                                            |
| db_dbname                           | ""                                          | database db name                                                                                                             |
| cache_max_idle                      | 500                                         | cache max idle                                                                                                               |
| cache_max_active                    | 1000                                        | cache idle                                                                                                                   |
| cache_idle_timeout_sec              | 1200                                        | close connections after remaining idle for this duration                                                                     |
| cache_wait                          | false                                       | if true and the pool is at the MaxActive limit then Get() waits for a connection to be returned to the pool before returning |
| cache_connect_timeout_sec           | 5                                           | cache connect timeout second                                                                                                 |
| cache_read_timeout_sec              | 10                                          | cache read timeout second                                                                                                    |
| cache_write_timeout_sec             | 15                                          | cache write timeout second                                                                                                   |
| cache_host                          | "127.0.0.1"                                 | cache host                                                                                                                   |
| cache_port                          | "6379"                                      | cache port                                                                                                                   |
| cache_password                      | ""                                          | cache password                                                                                                               |
| cache_db_index                      | 1                                           | cache db index                                                                                                               |
| zendesk_request_timeout             | 10                                          | zendesk api http request timeout second                                                                                      |
| zendesk_auth_token                  | ""                                          | zendesk api authorization token                                                                                              |
| zendesk_hk_base_url                 | https://honestbeehelp-hk.zendesk.com | zendesk hk base url                                                                                                          |
| zendesk_id_base_url                 | https://honestbee-idn.zendesk.com    | zendesk id base url                                                                                                          |
| zendesk_jp_base_url                 | https://honestbeehelp-jp.zendesk.com | zendesk jp base url                                                                                                          |
| zendesk_my_base_url                 | https://honestbee-my.zendesk.com     | zendesk my base url                                                                                                          |
| zendesk_ph_base_url                 | https://honestbee-ph.zendesk.com     | zendesk ph base url                                                                                                          |
| zendesk_sg_base_url                 | https://honestbeehelp-sg.zendesk.com | zendesk sg base url                                                                                                          |
| zendesk_th_base_url                 | https://honestbee-th.zendesk.com     | zendesk th base url                                                                                                          |
| zendesk_tw_base_url                 | https://honestbeehelp-tw.zendesk.com | zendesk tw base url                                                                                                          |
| examiner_max_worker_size            | 100                                         | examiner max worker size                                                                                                     |
| examiner_max_pool_size              | 200                                         | examiner max pool size                                                                                                       |
| examiner_categories_refresh_limit   | 0                                        | examiner categories refresh limit                                                                                            |
| examiner_sections_refresh_limit     | 0                                        | examiner sections refresh limit                                                                                              |
| examiner_articles_refresh_limit     | 0                                        | examiner articles refresh limit                                                                                              |
| examiner_ticket_forms_refresh_limit | 0                                        | examiner ticket forms refresh limit                                                                                          |
| graphql_max_depth                   | 13                                          | graphql max field nesting depth in a query                                                                                   |
| graphql_max_parallelism             | 10                                          | graphql max number of resolvers per request allowed to run in parallel                                                       |
| datadog_enable                       | true                                       | datadog enable |
| datadog_debug                       | false                                       | datadog debug |
| datadog_env                       | development                                       | datadog environment (development/staging/production) |
| datadog_host                       | localhost                                       | datadog host |
| datadog_port                       | 8126                                       | datadog port |
| grpc_listen_addr                       | :50051                                       | gRPC server address  |


### Install Cache
using redis cache, please follow [the instruction](https://redis.io/download)
or using docker image
```bash
docker run --name redis -p 6379:6379 -d redis
```

### Database

#### Setup Database Env
setup ENV for goose to connect to the database

develop usage:

using database name **zen**

```bash
export DB_USER={the user name}
export DB_PASSWORD={the user password}
export ZEN_DATABASE_URI=localhost
export ZEN_DATABASE_NAME=zen
```

ci usage:

```bash
export DB_USER=
export DB_PASSWORD=
export ZEN_DATABASE_URI=
export ZEN_DATABASE_NAME=
```

### Install Database
using Postgres database, please follow [the instruction](https://www.postgresql.org/download/)
or using docker image
```bash
docker run --name postgres -e POSTGRES_USER=${DB_USER} -e POSTGRES_PASSWORD=${DB_PASSWORD} -e POSTGRES_DB=${ZEN_DATABASE_NAME} -p 5432:5432 -d postgres
```

### Database Migration
using [goose](https://bitbucket.org/liamstask/goose)

#### Install Goose
```bash
go get bitbucket.org/liamstask/goose/cmd/goose
```

#### Goose version (Print the current version of the database)
```bash
goose dbversion

goose: dbversion 20180301103424
```

#### Goose up (Apply all available migrations)
```bash
goose -env=ci up

goose: migrating db environment 'development', current version: 0, target: 20180301103754
OK    20180301100347_addCategories.sql
OK    20180301103424_addSections.sql
OK    20180301103754_addArticles.sql
```

#### Goose down (Roll back a single migration from the current version)
```bash
goose down

goose: migrating db environment 'development', current version: 20180301103754, target: 20180301103424
OK    20180301103754_addArticles.sql
...
2018/03/02 10:43:08 no previous version found
```

#### Goose status (Dump the migration status for the current DB)
```bash
goose status

goose: status for environment 'development'
    Applied At                  Migration
    =======================================
    Fri Mar  2 02:44:31 2018 -- 20180301100347_addCategories.sql
    Fri Mar  2 02:44:31 2018 -- 20180301103424_addSections.sql
    Fri Mar  2 02:44:31 2018 -- 20180301103754_addArticles.sql
```

### Datadog Agent
using datadog agent docker image
```
docker run -d -v /var/run/docker.sock:/var/run/docker.sock:ro -p 8126:8126 \
              -v /proc/:/host/proc/:ro \
              -v /sys/fs/cgroup/:/host/sys/fs/cgroup:ro \
              -e DD_API_KEY=xxx \
              -e DD_APM_ENABLED=true --name datadog-agent \
              datadog/agent:latest
```

### Testing
using TDD testing, so simply type in Zen root folder
```bash
go test -race -v -cover -count=1 ./...
```

### Integration Testing
in Zen root folder, setting up the env.yml then type:
```bash
go test -race -v -cover -count=1 -tags=integration ./integration -config_path=`pwd`/env.yml
```


## DevOps

### Build Time Setup App Version
```bash
go build -ldflags "-X github.com/honestbee/Zen/config.Version={$APP_VERSION}"
```

### Check Server Status
```bash
curl localhost:8080/api/status
{"go-version":"go1.11","app-version":"1.0.0","server-time":"2018-03-03 05:23:50.469746859 +0000 UTC"}
```

### Interact With cc-test-reporter
since the alpine image will failed on race test, [see issue](https://github.com/golang/go/issues/14481)

we will not using race test in the CI.

setting up like this:
```bash
go test ./... -v -coverprofile=unit.out
go test -v -coverprofile=integration.out -tags=integration ./integration
./cc-test-reporter format-coverage -t gocov unit.out integration.out
```
see .drone.yml for full example
