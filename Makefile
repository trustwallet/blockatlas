#! /usr/bin/make -f

# Project variables.
PACKAGE := github.com/trustwallet/blockatlas
VERSION := $(shell git describe --tags 2>/dev/null || git describe --all)
BUILD := $(shell git rev-parse --short HEAD)
DATETIME := $(shell date +"%Y.%m.%d-%H:%M:%S")
PROJECT_NAME := $(shell basename "$(PWD)")
API := api
CONSUMER := consumer
PARSER := parser
COIN_FILE := coin/coins.yml
COIN_GO_FILE := coin/coins.go
GEN_COIN_FILE := coin/gen.go
DOCKER_LOCAL_DB_IMAGE_NAME := test_db
DOCKER_LOCAL_MQ_IMAGE_NAME := mq
DOCKER_LOCAL_DB_USER :=user
DOCKER_LOCAL_DB_PASS :=pass
DOCKER_LOCAL_DB := blockatlas

# Go related variables.
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
GOPKG := $(.)
# A valid GOPATH is required to use the `go get` command.
# If $GOPATH is not specified, $HOME/go will be used by default
GOPATH := $(if $(GOPATH),$(GOPATH),~/go)

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
PID_API := /tmp/.$(PROJECT_NAME).$(API).pid
PID_CONSUMER := /tmp/.$(PROJECT_NAME).$(CONSUMER).pid
PID_PARSER := /tmp/.$(PROJECT_NAME).$(PARSER).pid
PID_MOCKSERVER := /tmp/.$(PROJECT_NAME).mockserver.pid
# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

## install: Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar
install: go-get

## start: Start API, Observer and Sync in development mode.
start:
	@bash -c "$(MAKE) clean compile start-api start-parser start-consumer"

## start-api: Start platform api in development mode.
start-api: stop
	@echo "  >  Starting $(PROJECT_NAME)"
	@-$(GOBIN)/$(API)/api -c $(CONFIG_FILE) 2>&1 & echo $$! > $(PID_API)
	@cat $(PID_API) | sed "/^/s/^/  \>  Api PID: /"
	@echo "  >  Error log: $(STDERR)"

# start-platform-api-mock: Start API.  Similar to start-platform-api, but uses config file with mock URLs, and port 8437.
start-api-mock: stop start-mockserver
	@echo "  >  Starting $(PROJECT_NAME) API"
	@-$(GOBIN)/$(API)/api -p 8437 -c $(CONFIG_MOCK_FILE) 2>&1 & echo $$! > $(PID_API)
	@cat $(PID_API) | sed "/^/s/^/  \>  Mock PID: /"
	@echo "  >  Error log: $(STDERR)"

## start-parser: Start observer-parser in development mode.
start-parser: stop
	@echo "  >  Starting $(PROJECT_NAME)"
	@-$(GOBIN)/$(PARSER)/parser -c $(CONFIG_FILE) 2>&1 & echo $$! > $(PID_PARSER)
	@cat $(PID_PARSER) | sed "/^/s/^/  \>  Parser PID: /"
	@echo "  >  Error log: $(STDERR)"

## start-consumer: Start observer-consumer in development mode.
start-consumer: stop
	@echo "  >  Starting $(PROJECT_NAME)"
	@-$(GOBIN)/$(CONSUMER)/consumer -c $(CONFIG_FILE) 2>&1 & echo $$! > $(PID_CONSUMER)
	@cat $(PID_CONSUMER) | sed "/^/s/^/  \>  consumer PID: /"
	@echo "  >  Error log: $(STDERR)"

## stop: Stop development mode.
stop:
	@-touch $(PID_API) $(PID_CONSUMER) $(PID_PARSER)
	@-kill `cat $(PID_API)` 2> /dev/null || true
	@-kill `cat $(PID_CONSUMER)` 2> /dev/null || true
	@-kill `cat $(PID_PARSER)` 2> /dev/null || true
	@-rm $(PID_API) $(PID_CONSUMER) $(PID_PARSER)

stop-mockserver:
	@-touch $(PID_MOCKSERVER)
	@kill `cat $(PID_MOCKSERVER)` 2> /dev/null || true
	@rm $(PID_MOCKSERVER)

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

## start-mockserver: Start Mockserver with mocks of external services.  Test that it is operational (nasty case if port is taken).
start-mockserver: stop-mockserver
	@echo "  >  Starting Mockserver"
	GOBIN=$(GOBIN) go build -o $(GOBIN)/mockserver/mockserver ./mock/mockserver
	@-./bin/mockserver/mockserver & echo $$! > $(PID_MOCKSERVER)
	@echo "  >  Mockserver started with PID: " `cat $(PID_MOCKSERVER)`
	@sleep 1
	# Check that mock is running, by making a test with simple call (e.g. may fail due to unavailable port)
	@newman run tests/postman/blockatlas.postman_collection.json --folder mock-healthcheck --env-var "host=http://localhost:8437"

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


install-swag:
ifeq (,$(wildcard test -f $(GOPATH)/bin/swag))
	@echo "  >  Installing swagger"
	@-bash -c "go get github.com/swaggo/swag/cmd/swag"
endif

swag: install-swag
	@bash -c "$(GOPATH)/bin/swag init --parseDependency -g ./cmd/api/main.go -o ./docs"

## install-newman: Install Postman Newman for tests.
install-newman:
ifeq (,$(shell which newman))
	@echo "  >  Installing Postman Newman"
	@-sudo npm install -g newman
endif

## newman-mocked: Run mocked Postman Newman tests.
newman-mocked: install-newman go-compile
	@bash -c "$(MAKE) newman-mocked-params host=http://localhost:8437"

## newman-mocked-params: Run mocked Postman Newman tests, after starting platform api.
## The host parameter is required.
## E.g.: $ make newman-mocked-params test=transaction host=http://localhost:8437
newman-mocked-params: start-api-mock
ifeq (,$(test))
	@bash -c "$(MAKE) newman-run test=transaction host=$(host) && \
			  $(MAKE) newman-run test=staking host=$(host) && \
			  $(MAKE) newman-run test=token host=$(host) && \
			  $(MAKE) newman-run test=collection host=$(host)"
	@bash -c "$(MAKE) stop"
else
	@bash -c "$(MAKE) newman-run test=$(test) host=$(host)"
	@bash -c "$(MAKE) stop"
endif

## newman: Run Postman Newman test, the host parameter is required, and you can specify the name of the test do you wanna run (transaction, token, staking, collection, healthcheck, observer). e.g $ make newman test=staking host=http://localhost:8420
newman: install-newman
ifeq (,$(test))
	@bash -c "$(MAKE) newman-run test=transaction host=$(host)"
	@bash -c "$(MAKE) newman-run test=token host=$(host)"
	@bash -c "$(MAKE) newman-run test=staking host=$(host)"
	@bash -c "$(MAKE) newman-run test=collection host=$(host)"
	@bash -c "$(MAKE) newman-run test=healthcheck host=$(host)"
else
	@bash -c "$(MAKE) newman-run test=$(test) host=$(host)"
endif

## newman-run: Run chosen Newman tests. See newman target.
newman-run:
ifeq (,$(host))
	@echo "  >  Host parameter is missing. e.g: make newman test=staking host=http://localhost:8437"
	@exit 1
endif
	@echo "  >  Running $(test) tests"
	@newman run tests/postman/blockatlas.postman_collection.json --folder $(test) -d tests/postman/$(test)_data.json --env-var "host=$(host)"

go-compile: go-get go-build

go-build: go-build-api go-build-consumer go-build-parser

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

go-build-consumer:
	@echo "  >  Building consumer binary..."
	GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(CONSUMER)/consumer ./cmd/$(CONSUMER)

go-build-parser:
	@echo "  >  Building parser binary..."
	GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(PARSER)/parser ./cmd/$(PARSER)

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
