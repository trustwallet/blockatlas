# Block Atlas by Trust Wallet

![Go Version](https://img.shields.io/github/go-mod/go-version/TrustWallet/blockatlas)
[![Build Status](https://dev.azure.com/TrustWallet/Trust%20BlockAtlas/_apis/build/status/TrustWallet.blockatlas?branchName=master)](https://dev.azure.com/TrustWallet/Trust%20BlockAtlas/_build/latest?definitionId=27&branchName=master)
[![codecov](https://codecov.io/gh/trustwallet/blockatlas/branch/master/graph/badge.svg)](https://codecov.io/gh/trustwallet/blockatlas)
[![Go Report Card](https://goreportcard.com/badge/trustwallet/blockatlas)](https://goreportcard.com/report/TrustWallet/blockatlas)

> BlockAtlas is a clean explorer API and transaction observer for cryptocurrencies.

BlockAtlas connects to nodes or explorer APIs of the supported coins and maps transaction data,
account transaction history into a generic, easy to work with JSON format.
It is in production use at the [Trust Wallet app](https://trustwallet.com/), 
the official cryptocurrency wallet of Binance. Also is in production at the [BUTTON Wallet](https://buttonwallet.com), Telegram based non-custodial wallet.
The observer API watches the chain for new transactions and generates notifications by guids.

#### Supported Coins

<a href="https://binance.com" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/binance/info/logo.png" width="32" /></a>
<a href="https://nimiq.com" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/nimiq/info/logo.png" width="32" /></a>
<a href="https://ripple.com" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/ripple/info/logo.png" width="32" /></a>
<a href="https://stellar.org" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/stellar/info/logo.png" width="32" /></a>
<a href="https://kin.org" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/kin/info/logo.png" width="32" /></a>
<a href="https://tezos.com" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/tezos/info/logo.png" width="32" /></a>
<a href="https://aion.network" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/aion/info/logo.png" width="32" /></a>
<a href="https://ethereum.org" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/ethereum/info/logo.png" width="32" /></a>
<a href="https://ethereumclassic.org/" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/classic/info/logo.png" width="32" /></a>
<a href="https://poa.network" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/poa/info/logo.png" width="32" /></a>
<a href="https://callisto.network" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/callisto/info/logo.png" width="32" /></a>
<a href="https://gochain.io" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/gochain/info/logo.png" width="32" /></a>
<a href="https://wanchain.org" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/wanchain/info/logo.png" width="32" /></a>
<a href="https://thundercore.com" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/thundertoken/info/logo.png" width="32" /></a>
<a href="https://icon.foundation" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/icon/info/logo.png" width="32" /></a>
<a href="https://tron.network" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/tron/info/logo.png" width="32" /></a>
<a href="https://vechain.org/" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/vechain/info/logo.png" width="32" /></a>
<a href="https://www.thetatoken.org/" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/theta/info/logo.png" width="32" /></a>
<a href="https://cosmos.network/" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/cosmos/info/logo.png" width="32" /></a>
<a href="https://bitcoin.org/" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/bitcoin/info/logo.png" width="32" /></a>
<a href="https://harmony.one/" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/harmony/info/logo.png" width="32" /></a>

## Architecture

#### NOTE
Currently Blockatlas is under active development and is not well documented. If you still want to run it on your own or help to contribute, **please** pay attention that currenlty integration, nemwan, functional tests are not working locally without all endpoints. We are fixing that issue and soon you will be able to test all the stuff locally

Blockatlas allows to:
- Get information about transactions, tokens, staking details, collectibles, crypto domains for supported coins.
- Subscribe for price notifications via Rabbit MQ

Platform API is independent service and can work with the specific blockchain only (like bitcoin, ethereum, etc)

Notifications:

(Observer Subscriber Producer) - Create new blockatlas.SubscriptionEvent [Not implemented at Atlas, write it on your own]

(Observer Subscriber) - Get subscriptions from queue, set them to the DB

(Observer Parser) - Parse the block, convert block to the transactions batch, send to queue

(Observer Notifier) - Check each transaction for having the same address as stored at DB, if so - send tx data and id to the next queue

(Observer Notifier Consumer) - Notify the user [Not implemented at Atlas, write it on your own]

```
New Subscriptions --(Rabbit MQ)--> Subscriber --> DB
                                                   |
                      Parser  --(Rabbit MQ)--> Notifier --(Rabbit MQ)--> Notifier Consumer --> User

```

The whole flow is not availible at Atlas repo. We will have integration tests with it. Also there will be examples of all instances soon.

## Setup

### Requirements

 * [Go Toolchain](https://golang.org/doc/install) versions 1.13+
 * [Postgres](https://www.postgresql.org/download) storing user subscriptions and latest parsed block number
 * [Rabbit MQ](https://www.rabbitmq.com/#getstarted) using to pass subscriptions and send transaction notifications

### From Source

#### IMPORTANT

You can run platform API for specific coin only!
```shell
cd cmd/platform_api
ATLAS_PLATFORM=ethereum go run main.go
```
You will run platform API for Ethereum coin only. You can run 30 coins with 30 binaries for scalability and sustainability. Howevever, you can run all of them at once by using ```ATLAS_PLATFORM=all``` env param

It works the same for observer_worker - you can run all observer at 1 binary or 30 coins per 30 binaries

```shell
# Download source to $GOPATH
go get -u github.com/trustwallet/blockatlas
cd $(go env GOPATH)/src/github.com/trustwallet/blockatlas

# Start observer_parser with the path to the config.yml ./ 
go build -o observer_parser-bin cmd/observer_parser/main.go && ./observer_parser-bin -c config.yml

# Start observer_notifier with the path to the config.yml ./ 
go build -o observer_notifier-bin cmd/observer_notifier/main.go && ./observer_notifier-bin -c config.yml

# Start observer_subscriber with the path to the config.yml ./ 
go build -o observer_subscriber-bin cmd/observer_subscriber/main.go && ./observer_subscriber-bin -c config.yml

# Start Platform API server at port 8420 with the path to the config.yml ./ 
go build -o platform-api-bin cmd/platform_api/main.go && ./platform-api-bin -p 8420 -c config.yml

# Startp Swagger API server at port 8422 with the path to the config.yml ./ 
go build -o swagger-api-bin cmd/swagger-api/main.go && ./swagger-api-bin -p 8423
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
