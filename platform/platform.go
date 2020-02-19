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

func GetVar(name string) string {
	return viper.GetString(name)
}

func GetApiVar(coinId uint) string {
	varName := fmt.Sprintf("%s.api", coin.Coins[coinId].Handle)
	return GetVar(varName)
}

func GetRpcVar(coinId uint) string {
	varName := fmt.Sprintf("%s.rpc", coin.Coins[coinId].Handle)
	return GetVar(varName)
}

func getPlatformMap() blockatlas.Platforms {
	return blockatlas.Platforms{
		coin.Fio().Handle:          fio.Init(GetVar("fio.api")),
		coin.Aion().Handle:         aion.Init(GetVar("aion.api")),
		coin.Icon().Handle:         icon.Init(GetVar("icon.api")),
		coin.Tron().Handle:         tron.Init(GetVar("tron.api")),
		coin.Nano().Handle:         nano.Init(GetVar("nano.api")),
		coin.Nimiq().Handle:        nimiq.Init(GetVar("nimiq.api")),
		coin.Iotex().Handle:        iotex.Init(GetVar("iotex.api")),
		coin.Theta().Handle:        theta.Init(GetVar("theta.api")),
		coin.Waves().Handle:        waves.Init(GetVar("waves.api")),
		coin.Ripple().Handle:       ripple.Init(GetVar("ripple.api")),
		coin.Harmony().Handle:      harmony.Init(GetVar("harmony.api")),
		coin.Vechain().Handle:      vechain.Init(GetVar("vechain.api")),
		coin.Nebulas().Handle:      nebulas.Init(GetVar("nebulas.api")),
		coin.Ontology().Handle:     ontology.Init(GetVar("ontology.api")),
		coin.Algorand().Handle:     algorand.Init(GetVar("algorand.api")),
		coin.Aeternity().Handle:    aeternity.Init(GetVar("aeternity.api")),
		coin.Tezos().Handle:        tezos.Init(GetVar("tezos.api"), GetVar("tezos.rpc")),
		coin.Binance().Handle:      binance.Init(GetVar("binance.api"), GetVar("binance.dex")),
		coin.Zilliqa().Handle:      zilliqa.Init(GetVar("zilliqa.api"), GetVar("zilliqa.key"), GetVar("zilliqa.rpc"), GetVar("zilliqa.lookup")),
		coin.Kusama().Handle:       polkadot.Init(coin.KSM, GetApiVar(coin.KSM)),
		coin.Stellar().Handle:      stellar.Init(coin.XLM, GetApiVar(coin.XLM)),
		coin.Kin().Handle:          stellar.Init(coin.KIN, GetApiVar(coin.KIN)),
		coin.Cosmos().Handle:       cosmos.Init(coin.ATOM, GetApiVar(coin.ATOM)),
		coin.Kava().Handle:         cosmos.Init(coin.KAVA, GetApiVar(coin.KAVA)),
		coin.Bitcoin().Handle:      bitcoin.Init(coin.BTC, GetApiVar(coin.BTC)),
		coin.Litecoin().Handle:     bitcoin.Init(coin.LTC, GetApiVar(coin.LTC)),
		coin.Bitcoincash().Handle:  bitcoin.Init(coin.BCH, GetApiVar(coin.BCH)),
		coin.Zcash().Handle:        bitcoin.Init(coin.ZEC, GetApiVar(coin.ZEC)),
		coin.Zcoin().Handle:        bitcoin.Init(coin.XZC, GetApiVar(coin.XZC)),
		coin.Viacoin().Handle:      bitcoin.Init(coin.VIA, GetApiVar(coin.VIA)),
		coin.Ravencoin().Handle:    bitcoin.Init(coin.RVN, GetApiVar(coin.RVN)),
		coin.Groestlcoin().Handle:  bitcoin.Init(coin.GRS, GetApiVar(coin.GRS)),
		coin.Zelcash().Handle:      bitcoin.Init(coin.ZEL, GetApiVar(coin.ZEL)),
		coin.Decred().Handle:       bitcoin.Init(coin.DCR, GetApiVar(coin.DCR)),
		coin.Digibyte().Handle:     bitcoin.Init(coin.DGB, GetApiVar(coin.DGB)),
		coin.Dash().Handle:         bitcoin.Init(coin.DASH, GetApiVar(coin.DASH)),
		coin.Doge().Handle:         bitcoin.Init(coin.DOGE, GetApiVar(coin.DOGE)),
		coin.Qtum().Handle:         bitcoin.Init(coin.QTUM, GetApiVar(coin.QTUM)),
		coin.Gochain().Handle:      ethereum.Init(coin.GO, GetApiVar(coin.GO), GetRpcVar(coin.GO)),
		coin.Thundertoken().Handle: ethereum.Init(coin.TT, GetApiVar(coin.TT), GetRpcVar(coin.TT)),
		coin.Classic().Handle:      ethereum.Init(coin.ETC, GetApiVar(coin.ETC), GetRpcVar(coin.ETC)),
		coin.Poa().Handle:          ethereum.Init(coin.POA, GetApiVar(coin.POA), GetRpcVar(coin.POA)),
		coin.Callisto().Handle:     ethereum.Init(coin.CLO, GetApiVar(coin.CLO), GetRpcVar(coin.CLO)),
		coin.Wanchain().Handle:     ethereum.Init(coin.WAN, GetApiVar(coin.WAN), GetRpcVar(coin.WAN)),
		coin.Tomochain().Handle:    ethereum.Init(coin.TOMO, GetApiVar(coin.TOMO), GetRpcVar(coin.TOMO)),
		coin.Ethereum().Handle:     ethereum.InitWitCollection(coin.ETH, GetApiVar(coin.ETH), GetRpcVar(coin.ETH), GetVar("ethereum.collections_api"), GetVar("ethereum.collections_api_key")),
	}
}
