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
		coin.Tron().Handle:         tron.Init(config.Default.Tron.API, config.Default.Tron.Explorer),
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
		coin.Algorand().Handle:     algorand.Init(config.Default.Algorand.API),
		coin.Aeternity().Handle:    aeternity.Init(config.Default.Aeternity.API),
		coin.Solana().Handle:       solana.Init(config.Default.Solana.API),
		coin.Tezos().Handle:        tezos.Init(config.Default.Tezos.API, config.Default.Tezos.RPC),
		coin.Binance().Handle:      binance.Init(config.Default.Binance.API),
		coin.Zilliqa().Handle:      zilliqa.Init(config.Default.Zilliqa.API, config.Default.Zilliqa.Key, config.Default.Zilliqa.RPC),
		coin.Kusama().Handle:       polkadot.Init(coin.KSM, config.Default.Kusama.API),
		coin.Polkadot().Handle:     polkadot.Init(coin.DOT, config.Default.Polkadot.API),
		coin.Stellar().Handle:      stellar.Init(coin.XLM, config.Default.Stellar.API),
		coin.Kin().Handle:          stellar.Init(coin.KIN, config.Default.Kin.API),
		coin.Cosmos().Handle:       cosmos.Init(coin.ATOM, config.Default.Cosmos.API),
		coin.Kava().Handle:         kava.Init(coin.KAVA, config.Default.Kava.API),
		coin.Bitcoin().Handle:      bitcoin.Init(coin.BTC, config.Default.Bitcoin.API),
		coin.Litecoin().Handle:     bitcoin.Init(coin.LTC, config.Default.Litecoin.API),
		coin.Bitcoincash().Handle:  bitcoin.Init(coin.BCH, config.Default.Bitcoincash.API),
		coin.Zcash().Handle:        bitcoin.Init(coin.ZEC, config.Default.Zcash.API),
		coin.Zcoin().Handle:        bitcoin.Init(coin.XZC, config.Default.Zcoin.API),
		coin.Viacoin().Handle:      bitcoin.Init(coin.VIA, config.Default.Viacoin.API),
		coin.Ravencoin().Handle:    bitcoin.Init(coin.RVN, config.Default.Ravencoin.API),
		coin.Groestlcoin().Handle:  bitcoin.Init(coin.GRS, config.Default.Groestlcoin.API),
		coin.Zelcash().Handle:      bitcoin.Init(coin.ZEL, config.Default.Zelcash.API),
		coin.Decred().Handle:       bitcoin.Init(coin.DCR, config.Default.Decred.API),
		coin.Digibyte().Handle:     bitcoin.Init(coin.DGB, config.Default.Digibyte.API),
		coin.Dash().Handle:         bitcoin.Init(coin.DASH, config.Default.Dash.API),
		coin.Doge().Handle:         bitcoin.Init(coin.DOGE, config.Default.Doge.API),
		coin.Qtum().Handle:         bitcoin.Init(coin.QTUM, config.Default.Qtum.API),
		coin.Gochain().Handle:      ethereum.Init(coin.GO, config.Default.Gochain.API, config.Default.Gochain.RPC),
		coin.Thundertoken().Handle: ethereum.Init(coin.TT, config.Default.Thundertoken.API, config.Default.Thundertoken.RPC),
		coin.Classic().Handle:      ethereum.Init(coin.ETC, config.Default.Classic.API, config.Default.Classic.RPC),
		coin.Poa().Handle:          ethereum.Init(coin.POA, config.Default.Poa.API, config.Default.Poa.RPC),
		coin.Callisto().Handle:     ethereum.Init(coin.CLO, config.Default.Callisto.API, config.Default.Callisto.RPC),
		coin.Wanchain().Handle:     ethereum.Init(coin.WAN, config.Default.Wanchain.API, config.Default.Wanchain.RPC),
		coin.Tomochain().Handle:    ethereum.Init(coin.TOMO, config.Default.Tomochain.API, config.Default.Tomochain.RPC),
		coin.Bsc().Handle:          ethereum.InitWithBlockbook(coin.BSCLegacy, config.Default.BSC.API, config.Default.BSC.RPC),
		coin.Smartchain().Handle:   ethereum.InitWithBlockbook(coin.BSC, config.Default.Smartchain.API, config.Default.Smartchain.RPC),
		coin.Ethereum().Handle:     ethereum.InitWithCollection(coin.ETH, config.Default.Ethereum.RPC, config.Default.Ethereum.BlockbookAPI, config.Default.Ethereum.CollectionsAPI, config.Default.Ethereum.CollectionsKey),
		coin.Near().Handle:         near.Init(config.Default.Near.API),
		coin.Elrond().Handle:       elrond.Init(coin.ERD, config.Default.Elrond.API),
		coin.Filecoin().Handle:     filecoin.Init(config.Default.Filecoin.API),
	}
}

func getCollectionsHandlers() blockatlas.CollectionsAPIs {
	return blockatlas.CollectionsAPIs{
		coin.ETH: ethereum.InitWithCollection(coin.ETH, config.Default.Ethereum.RPC, config.Default.Ethereum.BlockbookAPI, config.Default.Ethereum.CollectionsAPI, config.Default.Ethereum.CollectionsKey),
	}
}
