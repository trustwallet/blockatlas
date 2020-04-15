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
)

const (
	allPlatformsHandle = "all"
)

var CollectionsWhitelist map[uint]bool

func GetVar(name string) string {
	return viper.GetString(name)
}

func GetApiVar(coinId uint) string {
	varName := fmt.Sprintf("%s.api", GetHandle(coinId))
	return GetVar(varName)
}

func GetRpcVar(coinId uint) string {
	varName := fmt.Sprintf("%s.rpc", GetHandle(coinId))
	return GetVar(varName)
}

func GetHandle(coinId uint) string {
	return coin.Coins[coinId].Handle
}

func getPlatformMap() blockatlas.Platforms {
	return blockatlas.Platforms{
		coin.Fio().Handle:          fio.Init(GetApiVar(coin.FIO)),
		coin.Aion().Handle:         aion.Init(GetApiVar(coin.AION)),
		coin.Icon().Handle:         icon.Init(GetApiVar(coin.ICX)),
		coin.Tron().Handle:         tron.Init(GetApiVar(coin.TRX)),
		coin.Nano().Handle:         nano.Init(GetApiVar(coin.NANO)),
		coin.Nimiq().Handle:        nimiq.Init(GetApiVar(coin.NIM)),
		coin.Iotex().Handle:        iotex.Init(GetApiVar(coin.IOTX)),
		coin.Theta().Handle:        theta.Init(GetApiVar(coin.THETA)),
		coin.Waves().Handle:        waves.Init(GetApiVar(coin.WAVES)),
		coin.Ripple().Handle:       ripple.Init(GetApiVar(coin.XRP)),
		coin.Harmony().Handle:      harmony.Init(GetApiVar(coin.ONE)),
		coin.Vechain().Handle:      vechain.Init(GetApiVar(coin.VET)),
		coin.Nebulas().Handle:      nebulas.Init(GetApiVar(coin.NAS)),
		coin.Ontology().Handle:     ontology.Init(GetApiVar(coin.ONT)),
		coin.Algorand().Handle:     algorand.Init(GetApiVar(coin.ALGO)),
		coin.Aeternity().Handle:    aeternity.Init(GetApiVar(coin.AE)),
		coin.Solana().Handle:       solana.Init(GetApiVar(coin.SOL)),
		coin.Tezos().Handle:        tezos.Init(GetApiVar(coin.XTZ), GetRpcVar(coin.XTZ)),
		coin.Binance().Handle:      binance.Init(GetApiVar(coin.BNB), GetVar("binance.dex")),
		coin.Zilliqa().Handle:      zilliqa.Init(GetApiVar(coin.ZIL), GetVar("zilliqa.key"), GetRpcVar(coin.ZIL), GetVar("zilliqa.lookup")),
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
		coin.Ethereum().Handle:     ethereum.InitWitCollection(coin.ETH, GetApiVar(coin.ETH), GetRpcVar(coin.ETH), GetVar("ethereum.blockbook_api"), GetVar("ethereum.collections_api"), GetVar("ethereum.collections_api_key")),
		coin.Near().Handle:         near.Init(GetApiVar(coin.NEAR)),
	}
}

func InitCollectionsWhitelist() {
	CollectionsWhitelist = make(map[uint]bool)
	CollectionsWhitelist[coin.Ethereum().ID] = true
}
