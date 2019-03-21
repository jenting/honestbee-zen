FROM golang:1.11-alpine AS build-stage

ENV GO111MODULE=on

ARG DEPLOY_ENV
ENV DEPLOY_ENV=${DEPLOY_ENV}
ARG GIT_COMMIT_HASH
ENV GIT_COMMIT_HASH=${GIT_COMMIT_HASH}

RUN apk add --update alpine-sdk
RUN mkdir -p /go/src/github.com/honestbee/Zen
WORKDIR /go/src/github.com/honestbee/Zen

COPY go.mod .
COPY go.sum .
# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

COPY . .
RUN go build -ldflags "-X github.com/honestbee/Zen/config.Version=${DEPLOY_ENV}-${GIT_COMMIT_HASH:0:8}" -tags=${DEPLOY_ENV} -o bin/zendesk

FROM alpine
RUN apk add --update ca-certificates
WORKDIR /go/bin
COPY --from=build-stage /go/src/github.com/honestbee/Zen/bin/zendesk /go/bin
COPY --from=build-stage /go/src/github.com/honestbee/Zen/env.yml /go/bin

ENTRYPOINT ["/go/bin/zendesk"]
