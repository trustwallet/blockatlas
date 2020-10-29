#! /usr/bin/make -f

# Project variables.
PACKAGE := github.com/trustwallet/blockatlas
VERSION := $(shell git describe --tags 2>/dev/null || git describe --all)
BUILD := $(shell git rev-parse --short HEAD)
DATETIME := $(shell date +"%Y.%m.%d-%H:%M:%S")
PROJECT_NAME := $(shell basename "$(PWD)")

API := api
NOTIFIER := notifier
PARSER := parser
SUBSCRIBER := subscriber
SEARCHER := searcher
INDEXER := indexer

DOCKER_LOCAL_DB_IMAGE_NAME := test_db
DOCKER_LOCAL_MQ_IMAGE_NAME := mq
DOCKER_LOCAL_DB_USER :=user
DOCKER_LOCAL_DB_PASS :=pass
DOCKER_LOCAL_DB := blockatlas

# Go related variables.
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin

# A valid GOPATH is required to use the `go get` command.
# If $GOPATH is not specified, $HOME/go will be used by default
GOPATH := $(if $(GOPATH),$(GOPATH),~/go)

# Environment variables
CONFIG_FILE=$(GOBASE)/config.yml

# Go files
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

# Use linker flags to provide version/build settings
LDFLAGS=-ldflags "-X=$(PACKAGE)/build.Version=$(VERSION) -X=$(PACKAGE)/build.Build=$(BUILD) -X=$(PACKAGE)/build.Date=$(DATETIME)"

# Redirect error output to a file, so we can show it in development mode.
STDERR := /tmp/.$(PROJECT_NAME)-stderr.txt

# PID file will keep the process id of the server
PID_API := /tmp/.$(PROJECT_NAME).$(API).pid
PID_SEARCHER := /tmp/.$(PROJECT_NAME).$(SEARCHER).pid
PID_INDEXER := /tmp/.$(PROJECT_NAME).$(INDEXER).pid
PID_NOTIFIER := /tmp/.$(PROJECT_NAME).$(NOTIFIER).pid
PID_PARSER := /tmp/.$(PROJECT_NAME).$(PARSER).pid
PID_SUBSCRIBER := /tmp/.$(PROJECT_NAME).$(SUBSCRIBER).pid
PID_MOCKSERVER := /tmp/.$(PROJECT_NAME).mockserver.pid
# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

## install: Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar
install: go-get

## compile: Compile the project.
compile:
	@-touch $(STDERR)
	@-rm $(STDERR)
	@-$(MAKE) -s go-compile 2> $(STDERR)
	@cat $(STDERR) | sed -e '1s/.*/\nError:\n/'  | sed 's/make\[.*/ /' | sed "/^/s/^/     /" 1>&2

## test: Run all unit tests.
test: go-test

## integration: Run all integration tests.
integration: go-integration

## golint: Run golint.
lint: go-lint-install go-lint

install-swag:
ifeq (,$(wildcard test -f $(GOPATH)/bin/swag))
	@echo "  >  Installing swagger"
	@-bash -c "go get github.com/swaggo/swag/cmd/swag"
endif

swag: install-swag
	@bash -c "$(GOPATH)/bin/swag init --parseDependency -g ./cmd/api/main.go -o ./docs"

go-compile: go-get go-build

go-build: go-build-api go-build-notifier go-build-parser go-build-subscriber go-build-searcher go-build-indexer

docker-shutdown:
	@echo "  >  Shutdown docker containers..."
	@-bash -c "docker rm -f $(DOCKER_LOCAL_DB_IMAGE_NAME) 2> /dev/null"
	@-bash -c "docker rm -f $(DOCKER_LOCAL_MQ_IMAGE_NAME) 2> /dev/null"

start-docker-services: docker-shutdown
	@echo "  >  Starting docker containers"
	docker run -d -p 5432:5432 --name $(DOCKER_LOCAL_DB_IMAGE_NAME) -e POSTGRES_USER=$(DOCKER_LOCAL_DB_USER) -e POSTGRES_PASSWORD=$(DOCKER_LOCAL_DB_PASS) -e POSTGRES_DB=$(DOCKER_LOCAL_DB) postgres
	docker run -d -p 5672:5672 --name $(DOCKER_LOCAL_MQ_IMAGE_NAME) rabbitmq

go-build-api:
	@echo "  >  Building api binary..."
	GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(API)/api ./cmd/$(API)

go-build-notifier:
	@echo "  >  Building notifier binary..."
	GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(NOTIFIER)/notifier ./cmd/$(NOTIFIER)

go-build-parser:
	@echo "  >  Building parser binary..."
	GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(PARSER)/parser ./cmd/$(PARSER)

go-build-subscriber:
	@echo "  >  Building subscriber binary..."
	GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(SUBSCRIBER)/subscriber ./cmd/$(SUBSCRIBER)

go-build-searcher:
	@echo "  >  Building searcher binary..."
	GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(SEARCHER)/searcher ./cmd/$(SEARCHER)

go-build-indexer:
	@echo "  >  Building indexer binary..."
	GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(INDEXER)/indexer ./cmd/$(INDEXER)

go-generate:
	@echo "  >  Generating dependency files..."
	GOBIN=$(GOBIN) go generate $(generate)

go-get:
	@echo "  >  Checking if there is any missing dependencies..."
	GOBIN=$(GOBIN) go get cmd/... $(get)

go-test:
	@echo "  >  Running unit tests"
	GOBIN=$(GOBIN) go test -cover -race -coverprofile=coverage.txt -covermode=atomic -v ./...

go-integration:
	@echo "  >  Running integration tests"
	GOBIN=$(GOBIN) TEST_CONFIG=$(CONFIG_FILE) go test -race -tags=integration -v ./tests/integration/...

go-fmt:
	@echo "  >  Format all go files"
	GOBIN=$(GOBIN) gofmt -w ${GOFMT_FILES}

go-gen-coins:
	@echo "  >  Generating coin file"
	COIN_FILE=$(COIN_FILE) COIN_GO_FILE=$(COIN_GO_FILE) GOBIN=$(GOBIN) go run -tags=coins $(GEN_COIN_FILE)

go-lint-install:
	@echo "  >  Installing golint"
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s

go-lint:
	@echo "  >  Running golint"
	bin/golangci-lint run --timeout=2m
