package platform

import (
	"fmt"

	"github.com/trustwallet/blockatlas/platform/kava"

	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
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
)

const (
	allPlatformsHandle = "all"
)

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

func getAllHandlers() blockatlas.Platforms {
	return blockatlas.Platforms{
		coin.Fio().Handle:          fio.Init(GetApiVar(coin.FIO)),
		coin.Aion().Handle:         aion.Init(GetApiVar(coin.AION)),
		coin.Icon().Handle:         icon.Init(GetApiVar(coin.ICON)),
		coin.Tron().Handle:         tron.Init(GetApiVar(coin.TRON), GetVar("tron.explorer")),
		coin.Nano().Handle:         nano.Init(GetApiVar(coin.NANO)),
		coin.Nimiq().Handle:        nimiq.Init(GetApiVar(coin.NIMIQ)),
		coin.Iotex().Handle:        iotex.Init(GetApiVar(coin.IOTEX)),
		coin.Theta().Handle:        theta.Init(GetApiVar(coin.THETA)),
		coin.Waves().Handle:        waves.Init(GetApiVar(coin.WAVES)),
		coin.Ripple().Handle:       ripple.Init(GetApiVar(coin.RIPPLE)),
		coin.Harmony().Handle:      harmony.Init(GetApiVar(coin.HARMONY)),
		coin.Vechain().Handle:      vechain.Init(GetApiVar(coin.VECHAIN)),
		coin.Nebulas().Handle:      nebulas.Init(GetApiVar(coin.NEBULAS)),
		coin.Ontology().Handle:     ontology.Init(GetApiVar(coin.ONTOLOGY)),
		coin.Algorand().Handle:     algorand.Init(GetApiVar(coin.ALGORAND)),
		coin.Aeternity().Handle:    aeternity.Init(GetApiVar(coin.AETERNITY)),
		coin.Solana().Handle:       solana.Init(GetApiVar(coin.SOLANA)),
		coin.Tezos().Handle:        tezos.Init(GetApiVar(coin.TEZOS), GetRpcVar(coin.TEZOS)),
		coin.Binance().Handle:      binance.Init(GetApiVar(coin.BINANCE)),
		coin.Zilliqa().Handle:      zilliqa.Init(GetApiVar(coin.ZILLIQA), GetVar("zilliqa.key"), GetRpcVar(coin.ZILLIQA), GetVar("zilliqa.lookup")),
		coin.Kusama().Handle:       polkadot.Init(coin.KUSAMA, GetApiVar(coin.KUSAMA)),
		coin.Polkadot().Handle:     polkadot.Init(coin.POLKADOT, GetApiVar(coin.POLKADOT)),
		coin.Stellar().Handle:      stellar.Init(coin.STELLAR, GetApiVar(coin.STELLAR)),
		coin.Kin().Handle:          stellar.Init(coin.KIN, GetApiVar(coin.KIN)),
		coin.Cosmos().Handle:       cosmos.Init(coin.COSMOS, GetApiVar(coin.COSMOS)),
		coin.Kava().Handle:         kava.Init(coin.KAVA, GetApiVar(coin.KAVA)),
		coin.Bitcoin().Handle:      bitcoin.Init(coin.BITCOIN, GetApiVar(coin.BITCOIN)),
		coin.Litecoin().Handle:     bitcoin.Init(coin.LITECOIN, GetApiVar(coin.LITECOIN)),
		coin.Bitcoincash().Handle:  bitcoin.Init(coin.BITCOINCASH, GetApiVar(coin.BITCOINCASH)),
		coin.Zcash().Handle:        bitcoin.Init(coin.ZCASH, GetApiVar(coin.ZCASH)),
		coin.Zcoin().Handle:        bitcoin.Init(coin.ZCOIN, GetApiVar(coin.ZCOIN)),
		coin.Viacoin().Handle:      bitcoin.Init(coin.VIACOIN, GetApiVar(coin.VIACOIN)),
		coin.Ravencoin().Handle:    bitcoin.Init(coin.RAVENCOIN, GetApiVar(coin.RAVENCOIN)),
		coin.Groestlcoin().Handle:  bitcoin.Init(coin.GROESTLCOIN, GetApiVar(coin.GROESTLCOIN)),
		coin.Zelcash().Handle:      bitcoin.Init(coin.ZELCASH, GetApiVar(coin.ZELCASH)),
		coin.Decred().Handle:       bitcoin.Init(coin.DECRED, GetApiVar(coin.DECRED)),
		coin.Digibyte().Handle:     bitcoin.Init(coin.DIGIBYTE, GetApiVar(coin.DIGIBYTE)),
		coin.Dash().Handle:         bitcoin.Init(coin.DASH, GetApiVar(coin.DASH)),
		coin.Doge().Handle:         bitcoin.Init(coin.DOGE, GetApiVar(coin.DOGE)),
		coin.Qtum().Handle:         bitcoin.Init(coin.QTUM, GetApiVar(coin.QTUM)),
		coin.Gochain().Handle:      ethereum.Init(coin.GOCHAIN, GetApiVar(coin.GOCHAIN), GetRpcVar(coin.GOCHAIN)),
		coin.Thundertoken().Handle: ethereum.Init(coin.THUNDERTOKEN, GetApiVar(coin.THUNDERTOKEN), GetRpcVar(coin.THUNDERTOKEN)),
		coin.Classic().Handle:      ethereum.Init(coin.CLASSIC, GetApiVar(coin.CLASSIC), GetRpcVar(coin.CLASSIC)),
		coin.Poa().Handle:          ethereum.Init(coin.POA, GetApiVar(coin.POA), GetRpcVar(coin.POA)),
		coin.Callisto().Handle:     ethereum.Init(coin.CALLISTO, GetApiVar(coin.CALLISTO), GetRpcVar(coin.CALLISTO)),
		coin.Wanchain().Handle:     ethereum.Init(coin.WANCHAIN, GetApiVar(coin.WANCHAIN), GetRpcVar(coin.WANCHAIN)),
		coin.Tomochain().Handle:    ethereum.Init(coin.TOMOCHAIN, GetApiVar(coin.TOMOCHAIN), GetRpcVar(coin.TOMOCHAIN)),
		coin.Bsc().Handle:          ethereum.InitWithBlockbook(coin.BSC, GetApiVar(coin.BSC), GetRpcVar(coin.BSC)),
		coin.Ethereum().Handle:     ethereum.InitWitCollection(coin.ETHEREUM, GetApiVar(coin.ETHEREUM), GetRpcVar(coin.ETHEREUM), GetVar("ethereum.blockbook_api"), GetVar("ethereum.collections_api"), GetVar("ethereum.collections_api_key")),
		coin.Near().Handle:         near.Init(GetApiVar(coin.NEAR)),
		coin.Elrond().Handle:       elrond.Init(coin.ELROND, GetApiVar(coin.ELROND)),
	}
}

func getCollectionsHandlers() blockatlas.CollectionsAPIs {
	return blockatlas.CollectionsAPIs{
		coin.ETHEREUM: ethereum.InitWitCollection(coin.ETHEREUM, GetApiVar(coin.ETHEREUM), GetRpcVar(coin.ETHEREUM), GetVar("ethereum.blockbook_api"), GetVar("ethereum.collections_api"), GetVar("ethereum.collections_api_key")),
	}
}

func getNamingHandlers() map[uint]blockatlas.NamingServiceAPI {
	return map[uint]blockatlas.NamingServiceAPI{
		coin.ETHEREUM: ethereum.Init(coin.ETHEREUM, GetApiVar(coin.ETHEREUM), GetRpcVar(coin.ETHEREUM)),
		coin.FIO:      fio.Init(GetApiVar(coin.FIO)),
		coin.ZILLIQA:  zilliqa.Init(GetApiVar(coin.ZILLIQA), GetVar("zilliqa.key"), GetRpcVar(coin.ZILLIQA), GetVar("zilliqa.lookup")),
	}
}
