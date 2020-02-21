# Block Atlas by Trust Wallet

![Go Version](https://img.shields.io/github/go-mod/go-version/TrustWallet/blockatlas)
[![GoDoc](https://godoc.org/github.com/TrustWallet/blockatlas?status.svg)](https://godoc.org/github.com/TrustWallet/blockatlas) 
[![Build Status](https://dev.azure.com/TrustWallet/Trust%20BlockAtlas/_apis/build/status/TrustWallet.blockatlas?branchName=master)](https://dev.azure.com/TrustWallet/Trust%20BlockAtlas/_build/latest?definitionId=27&branchName=master)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/43834b0c94ad4f6088629aa3e3bb5e94)](https://www.codacy.com/app/TrustWallet/blockatlas?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=TrustWallet/blockatlas&amp;utm_campaign=Badge_Grade)
[![Go Report Card](https://goreportcard.com/badge/trustwallet/blockatlas)](https://goreportcard.com/report/TrustWallet/blockatlas)
[![Docker](https://img.shields.io/docker/cloud/build/trustwallet/blockatlas.svg)](https://hub.docker.com/r/trustwallet/blockatlas)

> BlockAtlas is a clean explorer API and transaction observer for cryptocurrencies.

BlockAtlas connects to nodes or explorer APIs of the supported coins and maps transaction data,
account transaction history into a generic, easy to work with JSON format.
It is in production use at the [Trust Wallet app](https://trustwallet.com/), 
the official cryptocurrency wallet of Binance. Also is in production at the [BUTTON Wallet](https://buttonwallet.com), Telegram based non-custodial wallet.
The observer API watches the chain for new transactions and generates notifications by webhooks.

#### Supported Coins

<a href="https://binance.com" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/714.png" width="32" /></a>
<a href="https://nimiq.com" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/242.png" width="32" /></a>
<a href="https://ripple.com" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/144.png" width="32" /></a>
<a href="https://stellar.org" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/148.png" width="32" /></a>
<a href="https://kin.org" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/2017.png" width="32" /></a>
<a href="https://tezos.com" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/1729.png" width="32" /></a>
<a href="https://aion.network" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/425.png" width="32" /></a>
<a href="https://ethereum.org" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/60.png" width="32" /></a>
<a href="https://ethereumclassic.github.io" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/61.png" width="32" /></a>
<a href="https://poa.network" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/178.png" width="32" /></a>
<a href="https://callisto.network" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/820.png" width="32" /></a>
<a href="https://gochain.io" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/6060.png" width="32" /></a>
<a href="https://wanchain.org" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/5718350.png" width="32" /></a>
<a href="https://thundercore.com" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/1001.png" width="32" /></a>
<a href="https://icon.foundation" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/74.png" width="32" /></a>
<a href="https://tron.network" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/195.png" width="32" /></a>
<a href="https://vechain.org/" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/818.png" width="32" /></a>
<a href="https://www.thetatoken.org/" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/500.png" width="32" /></a>
<a href="https://cosmos.network/" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/118.png" width="32" /></a>
<a href="https://bitcoin.org/" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/0.png" width="32" /></a>
<a href="https://harmony.one/" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/harmony/info/logo.png" width="32" /></a>

## Setup

### Requirements

 * [Go Toolchain](https://golang.org/doc/install) versions 1.13+
 * [Redis](https://redis.io/topics/quickstart) instance (observer and markets)

### From Source

There are multiple services:

1. Platform API - to get transactions, staking, tokens, domain lookup for supported coins in common format
2. Observer API - to subscribe several addresses on different supported coins and receive webhook
3. Swagger API - swagger for all handlers of 1-3 APIs. You need to route requests to them on you own (nginx)

There are workers that are linked with Observer API and Market API:

5. Platform Observer - fetching latest blocks, parse them to common block specification, check subscribed addresses - send webhook. We use Redis to get information about subscribed addresses per coin with webhooks and caching latest block that was processed by observer

Observer API <-> Redis <-> Platform Observer

#### IMPORTANT

You can run platform API for specific coin only!
```shell
cd cmd/platform_api
ATLAS_PLATFORM=ethereum go run main.go
```
You will run platform API for Ethereum coin only. You can run 30 coins with 30 binaries for scalability and sustainability. Howevever, you can run all of them at once by using ```ATLAS_PLATFORM=all``` env param

It works the same for platoform_observer - you can run all observer at 1 binary or 30 coins per 30 binaries

```shell
# Download source to $GOPATH
go get -u github.com/trustwallet/blockatlas
cd $(go env GOPATH)/src/github.com/trustwallet/blockatlas

# Start platform_observer with the path to the config.yml ./ 
go build -o platform-observer-bin cmd/platform_observer/main.go && ./platform-observer-bin -c config.yml

# Start Platform API server at port 8420 with the path to the config.yml ./ 
go build -o platform-api-bin cmd/platform_api/main.go  && ./platform-api-bin -p 8420 -c config.yml

# Start Observer API server at port 8422 with the path to the config.yml ./ 
go build -o observer-api-bin cmd/observer-api/main.go  && ./observer-api-bin -p 8422 -c config.yml

# Startp Swagger API server at port 8422 with the path to the config.yml ./ 
go build -o swagger-api-bin cmd/swagger-api/main.go  && ./swagger-api-bin -p 8423
```

OR 

```shell
make go-build
```
Then
```shell
make start
```

### Docker

Build and run from local Dockerfile:

You should change `config.yml`:
```yaml
redis: redis://redis:6379
```
Then build:
```shell
docker-compose build
```

Run all services:
```shell
docker-compose up
```

If you need to start one service:
```shell
# Run only platform API 
docker-compose start platform_api
# Run only observer for addresses and api for it
docker-compose start platform_observer observer_api redis
# Run swagger api
docker-compose start swagger_api
```

## Configuration

Block Atlas can run just fine without configuration.
By default, all coins offering public RPC/explorer APIs are enabled.

If you want to use custom RPC endpoints, or enable coins without public RPC (like Nimiq),
you can configure Block Atlas over `config.yml` or environment variables.

#### Config File

Config is loaded from `config.yml` if it exists in the working directory.
The repository includes a [default config](./config.yml) for reference with all available config options.

Example (`config.yml`):

```yaml
nimiq:
  api: http://localhost:8648
#...
```

#### Environment

The rest gets loaded from environment variables.
Every config option is available under the `ATLAS_` prefix.
Nested keys are joined via `_` (Example `nimiq.api` => `NIMIQ_API`)

Example:

```shell
ATLAS_NIMIQ_API=http://localhost:8648
```

## Docs

Swagger API docs provided at path `/swagger/index.html`

#### Updating Docs

- After creating a new route, add comments to your API source code, [See Declarative Comments Format](https://swaggo.github.io/swaggo.io/declarative_comments_format/).
- Download Swag for Go by using:

    `$ go get -u github.com/swaggo/swag/cmd/swag`

- Run the Swag in your Go project root folder.

    `$ swag init -g ./cmd/platform_api/main.go -o ./docs`

## Contributing

If you'd like to add support for a new blockchain, feel free to file a pull request.
Note that most tokens that run on top of other chains are already supported and
don't require code changes (e.g. ERC-20).

The best way to submit feedback and report bugs is to open a GitHub issue.
Please be sure to include your operating system, version number, and
[steps](https://gist.github.com/nrollr/eb24336b8fb8e7ba5630) to reproduce reported bugs.

More resources for developers are in [CONTRIBUTING.md](CONTRIBUTING.md).
