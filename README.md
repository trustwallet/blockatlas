# Block Atlas by Trust Wallet

Clean explorer API for every crypto. WIP.

__Supported Coins__

<a href="https://binance.org" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/714.png" width="32" /></a>
<a href="https://nimiq.com/" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/242.png" width="32" /></a>
<a href="https://ripple.com/" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/144.png" width="32" /></a>

## Setup

#### Quick start

Deploy it in less than 30 seconds!

__From Source__ (Go Toolchain required)

```shell
go get -u trustwallet.com/blockatlas
~/go/bin/blockatlas
```

__With Docker__

`docker run -it -p 8420:8420 trustwallet/blockatlas`

## Configuration

__Config File__

By default, `config.yml` is loaded from the working directory.
Live reload is supported across the app.

Example:
```yaml
# App settings
gin:
  mode: release

# Enabled endpoints
platforms:
  - binance
  - nimiq

# Custom coin options
nimiq:
  rpc: http://localhost:8648
#...
```

__Environment__

The rest gets loaded from the environment variables.
Every config option is available under the `ATLAS_` prefix.

Example:
```shell
ATLAS_PLATFORMS="binance nimiq" \
ATLAS_NIMIQ_API=http://localhost:8648 \
blockatlas
```

Supported platforms:
 * [Heroku](http://heroku.com)
 * Docker _via Dockerfile_
 * Docker _[via Hub](https://hub.docker.com/r/trustwallet/blockatlas): `trustwallet/blockatlas`_
