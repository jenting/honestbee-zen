#!/usr/bin/make -f

all: dep schema
	go build

clean:
	go clean

unit-test: schema db
	go test -race -v -cover -count=1 ./...

integration-test: schema db
	go test -race -v -cover -count=1 -tags=integration ./integration -config_path=`pwd`/env.yml

dep:
	GO111MODULE=on go mod vendor

db:
	GO111MODULE=off go get bitbucket.org/liamstask/goose/cmd/goose
	goose up

schema:
	GO111MODULE=off go get github.com/jteeuwen/go-bindata/...
	go generate ./...

.PHONY: all clean test unit-test integration-test dep db schema
