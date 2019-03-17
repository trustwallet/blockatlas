# Block Atlas by Trust Wallet

Clean explorer API for every crypto. WIP.

### Supported Coins

<a href="https://binance.org" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/714.png" width="32" /></a>
<a href="https://nimiq.com/" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/242.png" width="32" /></a>
<a href="https://ripple.com/" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/144.png" width="32" /></a>

### Setup

__Quick start:__ (with config file)

```shell
# Create a workspace
mkdir blockatlas; cd blockatlas

# Install the Go toolchain
       apt install golang-go
# or # brew install go

# Download and build Block Atlas
go get trustwallet.com/blockatlas

# Use default config
cp ~/go/src/trustwallet.com/blockatlas/config_sample.yml ./config.yml

# Run app!
~/go/bin/blockatlas
```

__Detailed: Config File__

By default, `config.yml` is loaded from the working directory.
Live reload is supported across the app.

Example:
```yaml
# App settings
port: 8080
gin:
  mode: release

# Enabled endpoints
platforms:
  - binance
  - nimiq

# Coin specific options
binance:
  api: https://testnet-dex.binance.org/api/v1
nimiq:
  rpc: http://localhost:8648
#...
```

__Detailed: Environment__

If no `config.yml` is found, the app can run solely on environment variables.
Every config option is also available under the `ATLAS_` prefix.

Example:
```shell
ATLAS_PORT=9999 \
ATLAS_PLATFORMS="binance nimiq" \
ATLAS_BINANCE_API=https://testnet-dex.binance.org/api/v1 \
ATLAS_NIMIQ_API=http://localhost:8648 \
blockatlas
```

Other supported platforms:
 * [Heroku](http://heroku.com) (`Procfile` present)
 * Docker (`Dockerfile` present, Docker Hub image planned)
