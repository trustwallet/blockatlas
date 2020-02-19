package platform

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/platform/aeternity"
	"github.com/trustwallet/blockatlas/platform/aion"
	"github.com/trustwallet/blockatlas/platform/algorand"
	"github.com/trustwallet/blockatlas/platform/binance"
	"github.com/trustwallet/blockatlas/platform/bitcoin"
	"github.com/trustwallet/blockatlas/platform/cosmos"
	"github.com/trustwallet/blockatlas/platform/ethereum"
	"github.com/trustwallet/blockatlas/platform/fio"
	"github.com/trustwallet/blockatlas/platform/harmony"
	"github.com/trustwallet/blockatlas/platform/icon"
	"github.com/trustwallet/blockatlas/platform/iotex"
	"github.com/trustwallet/blockatlas/platform/nano"
	"github.com/trustwallet/blockatlas/platform/nebulas"
	"github.com/trustwallet/blockatlas/platform/nimiq"
	"github.com/trustwallet/blockatlas/platform/ontology"
	"github.com/trustwallet/blockatlas/platform/polkadot"
	"github.com/trustwallet/blockatlas/platform/ripple"
	"github.com/trustwallet/blockatlas/platform/stellar"
	"github.com/trustwallet/blockatlas/platform/tezos"
	"github.com/trustwallet/blockatlas/platform/theta"
	"github.com/trustwallet/blockatlas/platform/tron"
	"github.com/trustwallet/blockatlas/platform/vechain"
	"github.com/trustwallet/blockatlas/platform/waves"
	"github.com/trustwallet/blockatlas/platform/zilliqa"
)

func getVar(name string) string {
	return viper.GetString(name)
}

func getApiVar(coinId uint) string {
	varName := fmt.Sprintf("%s.api", coin.Coins[coinId].Handle)
	return getVar(varName)
}

func getRpcVar(coinId uint) string {
	varName := fmt.Sprintf("%s.rpc", coin.Coins[coinId].Handle)
	return getVar(varName)
}

func getPlatformMap() blockatlas.Platforms {
	return blockatlas.Platforms{
		coin.Fio().Handle:          fio.Init(getVar("fio.api")),
		coin.Aion().Handle:         aion.Init(getVar("aion.api")),
		coin.Icon().Handle:         icon.Init(getVar("icon.api")),
		coin.Tron().Handle:         tron.Init(getVar("tron.api")),
		coin.Nano().Handle:         nano.Init(getVar("nano.api")),
		coin.Nimiq().Handle:        nimiq.Init(getVar("nimiq.api")),
		coin.Iotex().Handle:        iotex.Init(getVar("iotex.api")),
		coin.Theta().Handle:        theta.Init(getVar("theta.api")),
		coin.Waves().Handle:        waves.Init(getVar("waves.api")),
		coin.Ripple().Handle:       ripple.Init(getVar("ripple.api")),
		coin.Harmony().Handle:      harmony.Init(getVar("harmony.api")),
		coin.Vechain().Handle:      vechain.Init(getVar("vechain.api")),
		coin.Nebulas().Handle:      nebulas.Init(getVar("nebulas.api")),
		coin.Ontology().Handle:     ontology.Init(getVar("ontology.api")),
		coin.Algorand().Handle:     algorand.Init(getVar("algorand.api")),
		coin.Aeternity().Handle:    aeternity.Init(getVar("aeternity.api")),
		coin.Tezos().Handle:        tezos.Init(getVar("tezos.api"), getVar("tezos.rpc")),
		coin.Binance().Handle:      binance.Init(getVar("binance.api"), getVar("binance.dex")),
		coin.Zilliqa().Handle:      zilliqa.Init(getVar("zilliqa.api"), getVar("zilliqa.key"), getVar("zilliqa.rpc"), getVar("zilliqa.lookup")),
		coin.Kusama().Handle:       polkadot.Init(coin.KSM, getApiVar(coin.KSM)),
		coin.Stellar().Handle:      stellar.Init(coin.XLM, getApiVar(coin.XLM)),
		coin.Kin().Handle:          stellar.Init(coin.KIN, getApiVar(coin.KIN)),
		coin.Cosmos().Handle:       cosmos.Init(coin.ATOM, getApiVar(coin.ATOM)),
		coin.Kava().Handle:         cosmos.Init(coin.KAVA, getApiVar(coin.KAVA)),
		coin.Bitcoin().Handle:      bitcoin.Init(coin.BTC, getApiVar(coin.BTC)),
		coin.Litecoin().Handle:     bitcoin.Init(coin.LTC, getApiVar(coin.LTC)),
		coin.Bitcoincash().Handle:  bitcoin.Init(coin.BCH, getApiVar(coin.BCH)),
		coin.Zcash().Handle:        bitcoin.Init(coin.ZEC, getApiVar(coin.ZEC)),
		coin.Zcoin().Handle:        bitcoin.Init(coin.XZC, getApiVar(coin.XZC)),
		coin.Viacoin().Handle:      bitcoin.Init(coin.VIA, getApiVar(coin.VIA)),
		coin.Ravencoin().Handle:    bitcoin.Init(coin.RVN, getApiVar(coin.RVN)),
		coin.Groestlcoin().Handle:  bitcoin.Init(coin.GRS, getApiVar(coin.GRS)),
		coin.Zelcash().Handle:      bitcoin.Init(coin.ZEL, getApiVar(coin.ZEL)),
		coin.Decred().Handle:       bitcoin.Init(coin.DCR, getApiVar(coin.DCR)),
		coin.Digibyte().Handle:     bitcoin.Init(coin.DGB, getApiVar(coin.DGB)),
		coin.Dash().Handle:         bitcoin.Init(coin.DASH, getApiVar(coin.DASH)),
		coin.Doge().Handle:         bitcoin.Init(coin.DOGE, getApiVar(coin.DOGE)),
		coin.Qtum().Handle:         bitcoin.Init(coin.QTUM, getApiVar(coin.QTUM)),
		coin.Gochain().Handle:      ethereum.Init(coin.GO, getApiVar(coin.GO), getRpcVar(coin.GO)),
		coin.Thundertoken().Handle: ethereum.Init(coin.TT, getApiVar(coin.TT), getRpcVar(coin.TT)),
		coin.Classic().Handle:      ethereum.Init(coin.ETC, getApiVar(coin.ETC), getRpcVar(coin.ETC)),
		coin.Poa().Handle:          ethereum.Init(coin.POA, getApiVar(coin.POA), getRpcVar(coin.POA)),
		coin.Callisto().Handle:     ethereum.Init(coin.CLO, getApiVar(coin.CLO), getRpcVar(coin.CLO)),
		coin.Wanchain().Handle:     ethereum.Init(coin.WAN, getApiVar(coin.WAN), getRpcVar(coin.WAN)),
		coin.Tomochain().Handle:    ethereum.Init(coin.TOMO, getApiVar(coin.TOMO), getRpcVar(coin.TOMO)),
		coin.Ethereum().Handle:     ethereum.InitWitCollection(coin.ETH, getApiVar(coin.ETH), getRpcVar(coin.ETH), getVar("ethereum.collections_api"), getVar("ethereum.collections_api_key")),
	}
}
