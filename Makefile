#! /usr/bin/make -f

# Project variables.
VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
PROJECT_NAME := $(shell basename "$(PWD)")
API_SERVICE := api
OBSERVER_SERVICE := observer
SYNC_SERVICE := syncmarkets
COIN_FILE := coin/coins.yml
COIN_GO_FILE := coin/coins.go
GEN_COIN_FILE := coin/gen.go

# Go related variables.
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
GOPKG := $(.)

# Environment variables
CONFIG_FILE=$(GOBASE)/config.yml

# Go files
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

# Use linker flags to provide version/build settings
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

# Redirect error output to a file, so we can show it in development mode.
STDERR := /tmp/.$(PROJECT_NAME)-stderr.txt

# PID file will keep the process id of the server
PID_API := /tmp/.$(PROJECT_NAME).$(API_SERVICE).pid
PID_OBSERVER := /tmp/.$(PROJECT_NAME).$(OBSERVER_SERVICE).pid
PID_SYNC := /tmp/.$(PROJECT_NAME).$(SYNC_SERVICE).pid

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

## install: Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar
install: go-get

## start: Start API, Observer and Sync in development mode.
start:
	@bash -c "$(MAKE) clean compile start-api start-observer start-syncmarkets"

## start-api: Start API in development mode.
start-api: stop
	@echo "  >  Starting $(PROJECT_NAME) API"
	@-$(GOBIN)/$(API_SERVICE)/api -c $(CONFIG_FILE) 2>&1 & echo $$! > $(PID_API)
	@cat $(PID_API) | sed "/^/s/^/  \>  API PID: /"
	@echo "  >  Error log: $(STDERR)"

## start-observer: Start Observer in development mode.
start-observer: stop
	@echo "  >  Starting $(PROJECT_NAME) Observer"
	@-$(GOBIN)/$(OBSERVER_SERVICE)/observer -c $(CONFIG_FILE) 2>&1 & echo $$! > $(PID_OBSERVER)
	@cat $(PID_OBSERVER) | sed "/^/s/^/  \>  Observer PID: /"
	@echo "  >  Error log: $(STDERR)"

## start-sync-markets: Start Sync markets in development mode.
start-syncmarkets: stop
	@echo "  >  Starting $(PROJECT_NAME) Sync"
	@-$(GOBIN)/$(SYNC_SERVICE)/syncmarkets -c $(CONFIG_FILE) 2>&1 & echo $$! > $(PID_SYNC)
	@cat $(PID_SYNC) | sed "/^/s/^/  \>  Sync PID: /"
	@echo "  >  Error log: $(STDERR)"

## stop: Stop development mode.
stop:
	@-touch $(PID_API) $(PID_OBSERVER)
	@-kill `cat $(PID_API)` 2> /dev/null || true
	@-kill `cat $(PID_OBSERVER)` 2> /dev/null || true
	@-kill `cat $(PID_SYNC)` 2> /dev/null || true
	@-rm $(PID_API) $(PID_OBSERVER)

## compile: Compile the project.
compile:
	@-touch $(STDERR)
	@-rm $(STDERR)
	@-$(MAKE) -s go-compile 2> $(STDERR)
	@cat $(STDERR) | sed -e '1s/.*/\nError:\n/'  | sed 's/make\[.*/ /' | sed "/^/s/^/     /" 1>&2

## exec: Run given command. e.g; make exec run="go test ./..."
exec:
	GOBIN=$(GOBIN) $(run)

## clean: Clean build files. Runs `go clean` internally.
clean:
	@-rm $(GOBIN)/$(PROJECT_NAME) 2> /dev/null
	@-$(MAKE) go-clean

## test: Run all unit tests.
test: go-test

## functional: Run all functional tests.
functional: go-functional

## integration: Run all functional tests.
integration: go-integration

## fmt: Run `go fmt` for all go files.
fmt: go-fmt

## gen-coins: Generate a new coin file.
gen-coins: remove-coin-file go-gen-coins

## remove-coin-file: Remove auto generated coin file.
remove-coin-file:
	@echo "  >  Removing "$(PROJECT_NAME)""
	@-rm $(GOBASE)/$(COIN_GO_FILE)

## goreleaser: Release the last tag version with GoReleaser.
goreleaser: go-goreleaser

## govet: Run go vet.
govet: go-vet

## golint: Run golint.
golint: go-lint

## docs: Generate swagger docs.
docs: go-gen-docs

## install-newman: Install Postman Newman for tests.
install-newman:
	@echo "  >  Running Postman Newman"
ifeq (,$(shell which newman))
	@echo "  >  Installing Postman Newman"
	@-brew install newman
endif

## run-newman: Run Postman Newman tests.
run-newman: install-newman newman-txs newman-tokens newman-staking newman-collections newman-domains newman-observer newman-healthcheck

## newman-txs: Run Newman tests for Transaction API's.
newman-txs:
	@echo " - Running Newman Test for Transaction API"
	@newman run pkg/tests/postman/Blockatlas.postman_collection.json --folder txs -d pkg/tests/postman/tx_data.json

## newman-tokens: Run Newman tests for tokens API's.
newman-tokens:
	@echo " - Running Newman Test for Tokens API"
	@newman run pkg/tests/postman/Blockatlas.postman_collection.json --folder tokens -d pkg/tests/postman/token_data.json

## newman-staking: Run Newman tests for staking API's.
newman-staking:
	@echo " - Running Newman Test for Staking API"
	@newman run pkg/tests/postman/Blockatlas.postman_collection.json --folder collection -d pkg/tests/postman/collection_data.json

## newman-collections: Run Newman tests for collections API's.
newman-collections:
	@echo " - Running Newman Test for Collection API"
	@newman run pkg/tests/postman/Blockatlas.postman_collection.json --folder stake -d pkg/tests/postman/staking_data.json

## newman-domains: Run Newman tests for domains API's.
newman-domains:
	@echo " - Running Newman Test for Domain API"
	@newman run pkg/tests/postman/Blockatlas.postman_collection.json --folder domain -d pkg/tests/postman/domain_data.json

## newman-observer: Run Newman tests for observer API's.
newman-observer:
	@echo " - Running Newman Test for Observer API"
	@newman run pkg/tests/postman/Blockatlas.postman_collection.json --folder observer -d pkg/tests/postman/observer_data.json

## newman-healthcheck: Run Newman tests for healthcheck API's.
newman-healthcheck:
	@echo " - Running Newman Test for Healthcheck API"
	@newman run pkg/tests/postman/Blockatlas.postman_collection.json --folder healthcheck

go-compile: go-get go-build

go-build:
	@echo "  >  Building api binary..."
	GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(API_SERVICE)/api ./cmd/$(API_SERVICE)
	@echo "  >  Building syncmarkets binary..."
	GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(SYNC_SERVICE)/syncmarkets ./cmd/$(SYNC_SERVICE)
	@echo "  >  Building observer binary..."
	GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(OBSERVER_SERVICE)/observer ./cmd/$(OBSERVER_SERVICE)

go-generate:
	@echo "  >  Generating dependency files..."
	GOBIN=$(GOBIN) go generate $(generate)

go-get:
	@echo "  >  Checking if there is any missing dependencies..."
	GOBIN=$(GOBIN) go get cmd/... $(get)

go-install:
	GOBIN=$(GOBIN) go install $(GOPKG)

go-clean:
	@echo "  >  Cleaning build cache"
	GOBIN=$(GOBIN) go clean

go-test:
	@echo "  >  Runing unit tests"
	GOBIN=$(GOBIN) go test -cover -race -v ./...

go-functional:
	@echo "  >  Runing functional tests"
	GOBIN=$(GOBIN) TEST_CONFIG=$(CONFIG_FILE) go test -race -tags=functional -v ./pkg/tests/functional

go-integration:
	@echo "  >  Runing integration tests"
	GOBIN=$(GOBIN) TEST_CONFIG=$(CONFIG_FILE) go test -race -tags=integration -v ./pkg/tests/integration

go-fmt:
	@echo "  >  Format all go files"
	GOBIN=$(GOBIN) gofmt -w ${GOFMT_FILES}

go-gen-coins:
	@echo "  >  Generating coin file"
	COIN_FILE=$(COIN_FILE) COIN_GO_FILE=$(COIN_GO_FILE) GOBIN=$(GOBIN) go run -tags=coins $(GEN_COIN_FILE)

go-gen-docs:
	@echo "  >  Generating swagger files"
	swag init -g ./cmd/api/main.go -o ./docs

go-goreleaser:
	@echo "  >  Releasing a new version"
	GOBIN=$(GOBIN) goreleaser --rm-dist

go-vet:
	@echo "  >  Running go vet"
	GOBIN=$(GOBIN) go vet ./...

go-lint:
	@echo "  >  Running golint"
	GOBIN=$(GOBIN) golint ./...

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECT_NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
