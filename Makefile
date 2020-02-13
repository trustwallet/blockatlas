#! /usr/bin/make -f

# Project variables.
VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
PROJECT_NAME := $(shell basename "$(PWD)")
API_SERVICE := platform_api
OBSERVER_SERVICE := platform_observer
OBSERVER_API := observer_api
MARKET_SERVICE := market_observer
MARKET_API := market_api
SWAGGER_API := swagger_api
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
PID_OBSERVER_API := /tmp/.$(PROJECT_NAME).$(OBSERVER_API).pid
PID_MARKET := /tmp/.$(PROJECT_NAME).$(MARKET_SERVICE).pid
PID_MARKET_API := /tmp/.$(PROJECT_NAME).$(MARKET_API).pid
PID_SWAGGER_API := /tmp/.$(PROJECT_NAME).$(SWAGGER_API).pid
# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

## install: Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar
install: go-get

## start: Start API, Observer and Sync in development mode.
start:
	@bash -c "$(MAKE) clean compile start-platform-api start-platform-observer start-observer-api start-market-observer start-market-api"

## start-api: Start API in development mode.
start-platform-api: stop
	@echo "  >  Starting $(PROJECT_NAME) API"
	@-$(GOBIN)/$(API_SERVICE)/platform_api -c $(CONFIG_FILE) 2>&1 & echo $$! > $(PID_API)
	@cat $(PID_API) | sed "/^/s/^/  \>  API PID: /"
	@echo "  >  Error log: $(STDERR)"

## start-observer: Start Observer in development mode.
start-platform-observer: stop
	@echo "  >  Starting $(PROJECT_NAME) Observer"
	@-$(GOBIN)/$(OBSERVER_SERVICE)/platform_observer -c $(CONFIG_FILE) 2>&1 & echo $$! > $(PID_OBSERVER)
	@cat $(PID_OBSERVER) | sed "/^/s/^/  \>  Observer PID: /"
	@echo "  >  Error log: $(STDERR)"

## start-observer: Start Observer in development mode.
start-observer-api: stop
	@echo "  >  Starting $(PROJECT_NAME) Observer"
	@-$(GOBIN)/$(OBSERVER_API)/observer_api -c $(CONFIG_FILE) 2>&1 & echo $$! > $(PID_OBSERVER_API)
	@cat $(PID_OBSERVER_API) | sed "/^/s/^/  \>  Observer PID: /"
	@echo "  >  Error log: $(STDERR)"

## start-sync-market: Start Sync market in development mode.
start-market-observer: stop
	@echo "  >  Starting $(PROJECT_NAME) Sync"
	@-$(GOBIN)/$(MARKET_SERVICE)/market_observer -c $(CONFIG_FILE) 2>&1 & echo $$! > $(PID_MARKET)
	@cat $(PID_MARKET) | sed "/^/s/^/  \>  Sync PID: /"
	@echo "  >  Error log: $(STDERR)"

## start-sync-market-api: Start Sync market api in development mode.
start-market-api: stop
	@echo "  >  Starting $(PROJECT_NAME) Sync API"
	@-$(GOBIN)/$(MARKET_API)/market_api -c $(CONFIG_FILE) 2>&1 & echo $$! > $(PID_MARKET_API)
	@cat $(PID_MARKET_API) | sed "/^/s/^/  \>  Sync PID: /"
	@echo "  >  Error log: $(STDERR)"

## start-sync-market-api: Start Sync market api in development mode.
start-swagger-api: stop
	@echo "  >  Starting $(PROJECT_NAME) Sync API"
	@-$(GOBIN)/$(SWAGGER_API)/swagger_api -c $(CONFIG_FILE) 2>&1 & echo $$! > $(PID_SWAGGER_API)
	@cat $(PID_SWAGGER_API) | sed "/^/s/^/  \>  Sync PID: /"
	@echo "  >  Error log: $(STDERR)"

## stop: Stop development mode.
stop:
	@-touch $(PID_API) $(PID_OBSERVER) $(PID_OBSERVER_API) $(PID_MARKET) $(PID_MARKET_API) $(PID_SWAGGER_API)
	@-kill `cat $(PID_API)` 2> /dev/null || true
	@-kill `cat $(PID_OBSERVER)` 2> /dev/null || true
	@-kill `cat $(PID_OBSERVER_API)` 2> /dev/null || true
	@-kill `cat $(PID_MARKET)` 2> /dev/null || true
	@-kill `cat $(PID_MARKET_API)` 2> /dev/null || true
	@-kill `cat $(PID_SWAGGER_API)` 2> /dev/null || true
	@-rm $(PID_API) $(PID_OBSERVER) $(PID_OBSERVER_API) $(PID_MARKET) $(PID_MARKET_API) $(PID_SWAGGER_API)

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
ifeq (,$(shell which newman))
	@echo "  >  Installing Postman Newman"
	@-npm install -g newman
endif

## newman: Run Postman Newman test, the host parameter is required, and you can specify the name of the test do you wanna run (transaction, token, staking, collection, domain, healthcheck, observer). e.g $ make newman test=staking host=http//localhost
newman: install-newman
	@echo "  >  Runing $(test) tests"
ifeq (,$(host))
	@echo "  >  Host parameter is missing. e.g: make newman test=staking host=http://localhost:8420"
	@exit 1
endif
ifeq (,$(test))
	@bash -c "$(MAKE) newman test=transaction host=$(host)"
	@bash -c "$(MAKE) newman test=token host=$(host)"
	@bash -c "$(MAKE) newman test=staking host=$(host)"
	@bash -c "$(MAKE) newman test=collection host=$(host)"
	@bash -c "$(MAKE) newman test=domain host=$(host)"
	@bash -c "$(MAKE) newman test=healthcheck host=$(host)"
	@bash -c "$(MAKE) newman test=observer host=$(host)"
	@bash -c "$(MAKE) newman test=market host=$(host)"
else
	@newman run pkg/tests/postman/Blockatlas.postman_collection.json --folder $(test) -d pkg/tests/postman/$(test)_data.json --env-var "host=$(host)"
endif

go-compile: go-get go-build

go-build:
	@echo "  >  Building platform_api binary..."
	GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(API_SERVICE)/platform_api ./cmd/$(API_SERVICE)
	@echo "  >  Building market_observer binary..."
	GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(MARKET_SERVICE)/market_observer ./cmd/$(MARKET_SERVICE)
	@echo "  >  Building market_api binary..."
	GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(MARKET_API)/market_api ./cmd/$(MARKET_API)
	@echo "  >  Building platform_observer binary..."
	GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(OBSERVER_SERVICE)/platform_observer ./cmd/$(OBSERVER_SERVICE)
	@echo "  >  Building observer_api binary..."
	GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(OBSERVER_API)/observer_api ./cmd/$(OBSERVER_API)
	@echo "  >  Building swagger_api binary..."
	GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(SWAGGER_API)/swagger_api ./cmd/$(SWAGGER_API)

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
	swag init -g ./cmd/platform_api/main.go -o ./docs

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
