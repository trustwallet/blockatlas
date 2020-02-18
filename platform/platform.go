package platform

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/platform/cosmos"
	"github.com/trustwallet/blockatlas/platform/polkadot"

	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/platform/aeternity"
	"github.com/trustwallet/blockatlas/platform/aion"
	"github.com/trustwallet/blockatlas/platform/algorand"
	"github.com/trustwallet/blockatlas/platform/binance"
	"github.com/trustwallet/blockatlas/platform/bitcoin"
	"github.com/trustwallet/blockatlas/platform/ethereum"
	"github.com/trustwallet/blockatlas/platform/fio"
	"github.com/trustwallet/blockatlas/platform/harmony"
	"github.com/trustwallet/blockatlas/platform/icon"
	"github.com/trustwallet/blockatlas/platform/iotex"
	"github.com/trustwallet/blockatlas/platform/nano"
	"github.com/trustwallet/blockatlas/platform/nebulas"
	"github.com/trustwallet/blockatlas/platform/nimiq"
	"github.com/trustwallet/blockatlas/platform/ontology"
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
	return fmt.Sprintf("%s.api", coin.Coins[coinId].Handle)
}

func getPlatformMap() blockatlas.Platforms {
	return blockatlas.Platforms{
		coin.Binance().Handle:      binance.Init(getVar("binance.api"), getVar("binance.dex")),
		coin.Nimiq().Handle:        nimiq.Init(getVar("nimiq.api")),
		coin.Ripple().Handle:       ripple.Init(getVar("ripple.api")),
		coin.Tezos().Handle:        tezos.Init(getVar("tezos.api"), getVar("tezos.rpc")),
		coin.Aion().Handle:         aion.Init(getVar("aion.api")),
		coin.Icon().Handle:         icon.Init(getVar("icon.api")),
		coin.Iotex().Handle:        iotex.Init(getVar("iotex.api")),
		coin.Ontology().Handle:     ontology.Init(getVar("ontology.api")),
		coin.Theta().Handle:        theta.Init(getVar("theta.api")),
		coin.Tron().Handle:         tron.Init(getVar("tron.api")),
		coin.Vechain().Handle:      vechain.Init(getVar("vechain.api")),
		coin.Zilliqa().Handle:      zilliqa.Init(getVar("zilliqa.api"), getVar("zilliqa.key"), getVar("zilliqa.rpc"), getVar("zilliqa.lookup")),
		coin.Waves().Handle:        waves.Init(getVar("waves.api")),
		coin.Aeternity().Handle:    aeternity.Init(getVar("aeternity.api")),
		coin.Nebulas().Handle:      nebulas.Init(getVar("nebulas.api")),
		coin.Fio().Handle:          fio.Init(getVar("fio.api")),
		coin.Algorand().Handle:     algorand.Init(getVar("algorand.api")),
		coin.Nano().Handle:         nano.Init(getVar("nano.api")),
		coin.Harmony().Handle:      harmony.Init(getVar("harmony.api")),
		coin.Kusama().Handle:       polkadot.Init(getApiVar(coin.KSM)),
		coin.Cosmos().Handle:       cosmos.Init(getApiVar(coin.ATOM)),
		coin.Kava().Handle:         cosmos.Init(getApiVar(coin.KAVA)),
		coin.Stellar().Handle:      stellar.Init(getApiVar(coin.XLM)),
		coin.Kin().Handle:          stellar.Init(getApiVar(coin.KIN)),
		coin.Ethereum().Handle:     ethereum.InitWitCollection(getApiVar(coin.ETH), getVar("ethereum.rpc"), getVar("ethereum.collections_api"), getVar("ethereum.collections_api_key")),
		coin.Classic().Handle:      ethereum.Init(getApiVar(coin.ETC), ""),
		coin.Poa().Handle:          ethereum.Init(getApiVar(coin.POA), ""),
		coin.Callisto().Handle:     ethereum.Init(getApiVar(coin.CLO), ""),
		coin.Gochain().Handle:      ethereum.Init(getApiVar(coin.GO), ""),
		coin.Wanchain().Handle:     ethereum.Init(getApiVar(coin.WAN), ""),
		coin.Tomochain().Handle:    ethereum.Init(getApiVar(coin.TOMO), ""),
		coin.Thundertoken().Handle: ethereum.Init(getApiVar(coin.TT), ""),
		coin.Bitcoin().Handle:      bitcoin.Init(getApiVar(coin.BTC)),
		coin.Litecoin().Handle:     bitcoin.Init(getApiVar(coin.LTC)),
		coin.Bitcoincash().Handle:  bitcoin.Init(getApiVar(coin.BCH)),
		coin.Dash().Handle:         bitcoin.Init(getApiVar(coin.DASH)),
		coin.Doge().Handle:         bitcoin.Init(getApiVar(coin.DOGE)),
		coin.Zcash().Handle:        bitcoin.Init(getApiVar(coin.ZEC)),
		coin.Zcoin().Handle:        bitcoin.Init(getApiVar(coin.XZC)),
		coin.Viacoin().Handle:      bitcoin.Init(getApiVar(coin.VIA)),
		coin.Ravencoin().Handle:    bitcoin.Init(getApiVar(coin.RVN)),
		coin.Qtum().Handle:         bitcoin.Init(getApiVar(coin.QTUM)),
		coin.Groestlcoin().Handle:  bitcoin.Init(getApiVar(coin.GRS)),
		coin.Zelcash().Handle:      bitcoin.Init(getApiVar(coin.ZEL)),
		coin.Decred().Handle:       bitcoin.Init(getApiVar(coin.DCR)),
		coin.Digibyte().Handle:     bitcoin.Init(getApiVar(coin.DGB)),
	}
}
