#! /usr/bin/make -f

# Project variables.
PACKAGE := github.com/trustwallet/blockatlas
VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
DATETIME := $(shell date +"%Y.%m.%d-%H:%M:%S")
PROJECT_NAME := $(shell basename "$(PWD)")
API_SERVICE := platform_api
OBSERVER_NOTIFIER := observer_notifier
OBSERVER_PARSER := observer_parser
OBSERVER_SUBSCRIBER := observer_subscriber
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
CONFIG_MOCK_FILE=$(GOBASE)/configmock.yml

# Go files
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

# Use linker flags to provide version/build settings
LDFLAGS=-ldflags "-X=$(PACKAGE)/build.Version=$(VERSION) -X=$(PACKAGE)/build.Build=$(BUILD) -X=$(PACKAGE)/build.Date=$(DATETIME)"

# Redirect error output to a file, so we can show it in development mode.
STDERR := /tmp/.$(PROJECT_NAME)-stderr.txt

# PID file will keep the process id of the server
PID_API := /tmp/.$(PROJECT_NAME).$(API_SERVICE).pid
PID_OBSERVER_NOTIFIER := /tmp/.$(PROJECT_NAME).$(OBSERVER_NOTIFIER).pid
PID_OBSERVER_PARSER := /tmp/.$(PROJECT_NAME).$(OBSERVER_PARSER).pid
PID_OBSERVER_SUBSCRIBER := /tmp/.$(PROJECT_NAME).$(OBSERVER_SUBSCRIBER).pid
PID_SWAGGER_API := /tmp/.$(PROJECT_NAME).$(SWAGGER_API).pid
PID_DYSON := /tmp/.$(PROJECT_NAME).dyson.pid
# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

## install: Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar
install: go-get

## start: Start API, Observer and Sync in development mode.
start:
	@bash -c "$(MAKE) clean compile start-platform-api start-observer-parser start-observer-notifier start-observer-subscriber"

## start-api: Start platform api in development mode.
start-platform-api: stop
	@echo "  >  Starting $(PROJECT_NAME)"
	@-$(GOBIN)/$(API_SERVICE)/platform_api -c $(CONFIG_FILE) 2>&1 & echo $$! > $(PID_API)
	@cat $(PID_API) | sed "/^/s/^/  \>  API PID: /"
	@echo "  >  Error log: $(STDERR)"

# start-platform-api-mock: Start API in development mode.  Similar to start-platform-api, but uses config file with mock URLs
start-platform-api-mock: stop start-mock-dyson
	@echo "  >  Starting $(PROJECT_NAME) API"
	@-$(GOBIN)/$(API_SERVICE)/platform_api -c $(CONFIG_MOCK_FILE) 2>&1 & echo $$! > $(PID_API)
	@cat $(PID_API) | sed "/^/s/^/  \>  API PID: /"
	@echo "  >  Error log: $(STDERR)"

## start-swagger-api: Start swagger api in development mode.
start-swagger-api: stop
	@echo "  >  Starting $(PROJECT_NAME)"
	@-$(GOBIN)/$(SWAGGER_API)/swagger_api -c $(CONFIG_FILE) 2>&1 & echo $$! > $(PID_SWAGGER_API)
	@cat $(PID_SWAGGER_API) | sed "/^/s/^/  \>  Sync PID: /"
	@echo "  >  Error log: $(STDERR)"

## start-observer-parser: Start observer-parser in development mode.
start-observer-parser: stop
	@echo "  >  Starting $(PROJECT_NAME)"
	@-$(GOBIN)/$(OBSERVER_PARSER)/observer_parser -c $(CONFIG_FILE) 2>&1 & echo $$! > $(PID_OBSERVER_PARSER)
	@cat $(PID_OBSERVER_PARSER) | sed "/^/s/^/  \>  Sync PID: /"
	@echo "  >  Error log: $(STDERR)"

## start-observer-notifier: Start observer-notifier in development mode.
start-observer-notifier: stop
	@echo "  >  Starting $(PROJECT_NAME)"
	@-$(GOBIN)/$(OBSERVER_NOTIFIER)/observer_notifier -c $(CONFIG_FILE) 2>&1 & echo $$! > $(PID_OBSERVER_NOTIFIER)
	@cat $(PID_OBSERVER_NOTIFIER) | sed "/^/s/^/  \>  Sync PID: /"
	@echo "  >  Error log: $(STDERR)"

## start-observer-subscriber: Start observer-subscriber in development mode.
start-observer-subscriber: stop
	@echo "  >  Starting $(PROJECT_NAME)"
	@-$(GOBIN)/$(OBSERVER_SUBSCRIBER)/observer_subscriber -c $(CONFIG_FILE) 2>&1 & echo $$! > $(PID_OBSERVER_SUBSCRIBER)
	@cat $(PID_OBSERVER_SUBSCRIBER) | sed "/^/s/^/  \>  Sync PID: /"
	@echo "  >  Error log: $(STDERR)"

## stop: Stop development mode.
stop:
	@-touch $(PID_API) $(PID_OBSERVER_NOTIFIER) $(PID_OBSERVER_PARSER) $(PID_OBSERVER_SUBSCRIBER) $(PID_SWAGGER_API) $(PID_DYSON)
	@-kill `cat $(PID_API)` 2> /dev/null || true
	@-kill `cat $(PID_OBSERVER_NOTIFIER)` 2> /dev/null || true
	@-kill `cat $(PID_OBSERVER_PARSER)` 2> /dev/null || true
	@-kill `cat $(PID_OBSERVER_SUBSCRIBER)` 2> /dev/null || true
	@-kill `cat $(PID_SWAGGER_API)` 2> /dev/null || true
	@-kill `cat $(PID_DYSON)` 2> /dev/null || true
	@-rm $(PID_API) $(PID_OBSERVER_NOTIFIER) $(PID_OBSERVER_PARSER) $(PID_OBSERVER_SUBSCRIBER) $(PID_SWAGGER_API) $(PID_DYSON)

stop-dyson:
	@-touch $(PID_DYSON)
	@kill `cat $(PID_DYSON)` 2> /dev/null || true
	@rm $(PID_DYSON)

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

## integration: Run all integration tests.
integration: go-integration

## start-mock-dyson: Start Dyson with mocks of external services
start-mock-dyson: stop-dyson
	@echo "  >  Starting Dyson with mocks"
	@-dyson  mock/ext-api-dyson & echo $$! > $(PID_DYSON)
	@echo "  >  Dyson started with PID: " `cat $(PID_DYSON)`

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
lint: go-lint-install go-lint

## docs: Generate swagger docs.
docs: go-gen-docs

## install-newman: Install Postman Newman for tests.
install-newman:
ifeq (,$(shell which newman))
	@echo "  >  Installing Postman Newman"
	@-sudo npm install -g newman
endif

## newman-mocked: Run mocked Postman Newman tests.
newman-mocked: install-newman install-dyson go-compile
	@bash -c "$(MAKE) newman-mocked-params host=http://localhost:8420"

## newman-mocked-params: Run mocked Postman Newman tests, after starting platform api.
## The host parameter is required.
## E.g.: $ make newman-mocked-params test=domain host=http://localhost:8420
newman-mocked-params: start-platform-api-mock
ifeq (,$(test))
	@bash -c "$(MAKE) newman-run test=transaction host=$(host) && \
	          $(MAKE) newman-run test=domain host=$(host) && \
			  $(MAKE) newman-run test=staking host=$(host)"
	#not-mocked-yet: $(MAKE) newman-run test=token host=$(host) && \
	#not-mocked-yet: $(MAKE) newman-run test=collection host=$(host) &&
	@bash -c "$(MAKE) stop"
else
	@bash -c "$(MAKE) newman-run test=$(test) host=$(host)"
	@bash -c "$(MAKE) stop"
endif

## newman: Run Postman Newman test, the host parameter is required, and you can specify the name of the test do you wanna run (transaction, token, staking, collection, domain, healthcheck, observer). e.g $ make newman test=staking host=http://localhost:8420
newman: install-newman
ifeq (,$(test))
	@bash -c "$(MAKE) newman-run test=transaction host=$(host)"
	@bash -c "$(MAKE) newman-run test=token host=$(host)"
	@bash -c "$(MAKE) newman-run test=staking host=$(host)"
	@bash -c "$(MAKE) newman-run test=collection host=$(host)"
	@bash -c "$(MAKE) newman-run test=domain host=$(host)"
	@bash -c "$(MAKE) newman-run test=healthcheck host=$(host)"
else
	@bash -c "$(MAKE) newman-run test=$(test) host=$(host)"
endif

## newman-run: Run chosen Newman tests. See newman target.
newman-run:
ifeq (,$(host))
	@echo "  >  Host parameter is missing. e.g: make newman test=staking host=http://localhost:8420"
	@exit 1
endif
	@echo "  >  Running $(test) tests"
	@newman run tests/postman/Blockatlas.postman_collection.json --folder $(test) -d tests/postman/$(test)_data.json --env-var "host=$(host)"

## install-dyson: Install Dyson for mocked tests.
install-dyson:
ifeq (,$(shell which dyson))
	@echo "  >  Installing Dyson"
	@-sudo npm install -g dyson
endif

go-compile: go-get go-build

go-build: go-build-platform-api go-build-observer-notifier go-build-observer-parser go-build-observer-subscriber go-build-swagger-api

go-build-platform-api:
	@echo "  >  Building platform_api binary..."
	GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(API_SERVICE)/platform_api ./cmd/$(API_SERVICE)

go-build-observer-notifier:
	@echo "  >  Building observer_notifier binary..."
	GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(OBSERVER_NOTIFIER)/observer_notifier ./cmd/$(OBSERVER_NOTIFIER)

go-build-observer-parser:
	@echo "  >  Building observer_parser binary..."
	GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(OBSERVER_PARSER)/observer_parser ./cmd/$(OBSERVER_PARSER)

go-build-observer-subscriber:
	@echo "  >  Building observer_subscriber binary..."
	GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(OBSERVER_SUBSCRIBER)/observer_subscriber ./cmd/$(OBSERVER_SUBSCRIBER)

go-build-swagger-api:
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

go-gen-docs:
	@echo "  >  Generating swagger files"
	swag init -g ./cmd/platform_api/main.go -o ./docs

go-goreleaser:
	@echo "  >  Releasing a new version"
	GOBIN=$(GOBIN) scripts/goreleaser --rm-dist

go-vet:
	@echo "  >  Running go vet"
	GOBIN=$(GOBIN) go vet ./...

go-lint-install:
	@echo "  >  Installing golint"
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s

go-lint:
	@echo "  >  Running golint"
	bin/golangci-lint run --timeout=2m

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECT_NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
