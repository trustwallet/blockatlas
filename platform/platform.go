package platform

import (
	"github.com/trustwallet/blockatlas/config"

	"github.com/trustwallet/blockatlas/platform/filecoin"
	"github.com/trustwallet/blockatlas/platform/kava"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/platform/aeternity"
	"github.com/trustwallet/blockatlas/platform/aion"
	"github.com/trustwallet/blockatlas/platform/algorand"
	"github.com/trustwallet/blockatlas/platform/binance"
	"github.com/trustwallet/blockatlas/platform/bitcoin"
	"github.com/trustwallet/blockatlas/platform/cosmos"
	"github.com/trustwallet/blockatlas/platform/elrond"
	"github.com/trustwallet/blockatlas/platform/ethereum"
	"github.com/trustwallet/blockatlas/platform/fio"
	"github.com/trustwallet/blockatlas/platform/harmony"
	"github.com/trustwallet/blockatlas/platform/icon"
	"github.com/trustwallet/blockatlas/platform/iotex"
	"github.com/trustwallet/blockatlas/platform/nano"
	"github.com/trustwallet/blockatlas/platform/near"
	"github.com/trustwallet/blockatlas/platform/nebulas"
	"github.com/trustwallet/blockatlas/platform/nimiq"
	"github.com/trustwallet/blockatlas/platform/ontology"
	"github.com/trustwallet/blockatlas/platform/polkadot"
	"github.com/trustwallet/blockatlas/platform/ripple"
	"github.com/trustwallet/blockatlas/platform/solana"
	"github.com/trustwallet/blockatlas/platform/stellar"
	"github.com/trustwallet/blockatlas/platform/tezos"
	"github.com/trustwallet/blockatlas/platform/theta"
	"github.com/trustwallet/blockatlas/platform/tron"
	"github.com/trustwallet/blockatlas/platform/vechain"
	"github.com/trustwallet/blockatlas/platform/waves"
	"github.com/trustwallet/blockatlas/platform/zilliqa"
	"github.com/trustwallet/golibs/coin"
)

const (
	allPlatformsHandle = "all"
)

func GetHandle(coinId uint) string {
	return coin.Coins[coinId].Handle
}

func getAllHandlers() blockatlas.Platforms {
	return blockatlas.Platforms{
		coin.Fio().Handle:          fio.Init(config.Default.Fio.API),
		coin.Aion().Handle:         aion.Init(config.Default.Aion.API),
		coin.Icon().Handle:         icon.Init(config.Default.Icon.API),
		coin.Tron().Handle:         tron.Init(config.Default.Tron.API, config.Default.Tron.Key),
		coin.Nano().Handle:         nano.Init(config.Default.Nano.API),
		coin.Nimiq().Handle:        nimiq.Init(config.Default.Nimiq.API),
		coin.Iotex().Handle:        iotex.Init(config.Default.Iotex.API),
		coin.Theta().Handle:        theta.Init(config.Default.Theta.API),
		coin.Waves().Handle:        waves.Init(config.Default.Waves.API),
		coin.Ripple().Handle:       ripple.Init(config.Default.Ripple.API),
		coin.Harmony().Handle:      harmony.Init(config.Default.Harmony.API),
		coin.Vechain().Handle:      vechain.Init(config.Default.Vechain.API),
		coin.Nebulas().Handle:      nebulas.Init(config.Default.Nebulas.API),
		coin.Ontology().Handle:     ontology.Init(config.Default.Ontology.API),
		coin.Algorand().Handle:     algorand.Init(config.Default.Algorand.API, config.Default.Algorand.Key),
		coin.Aeternity().Handle:    aeternity.Init(config.Default.Aeternity.API),
		coin.Solana().Handle:       solana.Init(config.Default.Solana.API),
		coin.Tezos().Handle:        tezos.Init(config.Default.Tezos.API, config.Default.Tezos.RPC, config.Default.Tezos.Baker),
		coin.Binance().Handle:      binance.Init(config.Default.Binance.API, config.Default.Binance.Key, config.Default.Binance.StakingAPI),
		coin.Zilliqa().Handle:      zilliqa.Init(config.Default.Zilliqa.API, config.Default.Zilliqa.Key, config.Default.Zilliqa.RPC),
		coin.Polkadot().Handle:     polkadot.Init(coin.POLKADOT, config.Default.Polkadot.API),
		coin.Stellar().Handle:      stellar.Init(coin.STELLAR, config.Default.Stellar.API),
		coin.Cosmos().Handle:       cosmos.Init(coin.COSMOS, config.Default.Cosmos.API),
		coin.Kava().Handle:         kava.Init(coin.KAVA, config.Default.Kava.API),
		coin.Bitcoin().Handle:      bitcoin.Init(coin.BITCOIN, config.Default.Bitcoin.API),
		coin.Litecoin().Handle:     bitcoin.Init(coin.LITECOIN, config.Default.Litecoin.API),
		coin.Bitcoincash().Handle:  bitcoin.Init(coin.BITCOINCASH, config.Default.Bitcoincash.API),
		coin.Zcash().Handle:        bitcoin.Init(coin.ZCASH, config.Default.Zcash.API),
		coin.Zcoin().Handle:        bitcoin.Init(coin.ZCOIN, config.Default.Zcoin.API),
		coin.Viacoin().Handle:      bitcoin.Init(coin.VIACOIN, config.Default.Viacoin.API),
		coin.Ravencoin().Handle:    bitcoin.Init(coin.RAVENCOIN, config.Default.Ravencoin.API),
		coin.Groestlcoin().Handle:  bitcoin.Init(coin.GROESTLCOIN, config.Default.Groestlcoin.API),
		coin.Zelcash().Handle:      bitcoin.Init(coin.ZELCASH, config.Default.Zelcash.API),
		coin.Decred().Handle:       bitcoin.Init(coin.DECRED, config.Default.Decred.API),
		coin.Digibyte().Handle:     bitcoin.Init(coin.DIGIBYTE, config.Default.Digibyte.API),
		coin.Dash().Handle:         bitcoin.Init(coin.DASH, config.Default.Dash.API),
		coin.Doge().Handle:         bitcoin.Init(coin.DOGE, config.Default.Doge.API),
		coin.Qtum().Handle:         bitcoin.Init(coin.QTUM, config.Default.Qtum.API),
		coin.Gochain().Handle:      ethereum.InitWithBlockbook(coin.GOCHAIN, config.Default.Gochain.API),
		coin.Thundertoken().Handle: ethereum.InitWithBlockbook(coin.THUNDERTOKEN, config.Default.Thundertoken.API),
		coin.Classic().Handle:      ethereum.InitWithBlockbook(coin.CLASSIC, config.Default.Classic.API),
		coin.Poa().Handle:          ethereum.InitWithBlockbook(coin.POA, config.Default.Poa.API),
		coin.Callisto().Handle:     ethereum.InitWithBlockbook(coin.CALLISTO, config.Default.Callisto.API),
		coin.Wanchain().Handle:     ethereum.InitWithBlockbook(coin.WANCHAIN, config.Default.Wanchain.API),
		coin.Tomochain().Handle:    ethereum.InitWithBlockbook(coin.TOMOCHAIN, config.Default.Tomochain.API),
		coin.Smartchain().Handle:   ethereum.InitWithBounce(coin.SMARTCHAIN, config.Default.Smartchain.API, config.Default.Smartchain.CollectionsAPI),
		coin.Ethereum().Handle:     ethereum.InitWithOpenSea(coin.ETHEREUM, config.Default.Ethereum.API, config.Default.Ethereum.CollectionsAPI, config.Default.Ethereum.CollectionsKey),
		coin.Near().Handle:         near.Init(config.Default.Near.API),
		coin.Elrond().Handle:       elrond.Init(coin.ELROND, config.Default.Elrond.API),
		coin.Filecoin().Handle:     filecoin.Init(config.Default.Filecoin.API, config.Default.Filecoin.Explorer),
	}
}

func getCollectionsHandlers() blockatlas.CollectionsAPIs {
	return blockatlas.CollectionsAPIs{
		coin.ETHEREUM:   ethereum.InitWithOpenSea(coin.ETHEREUM, config.Default.Ethereum.API, config.Default.Ethereum.CollectionsAPI, config.Default.Ethereum.CollectionsKey),
		coin.SMARTCHAIN: ethereum.InitWithBounce(coin.SMARTCHAIN, config.Default.Smartchain.API, config.Default.Smartchain.CollectionsAPI),
	}
}
