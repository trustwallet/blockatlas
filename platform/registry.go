package platform

import (
	"fmt"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/platform/algorand"

	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/platform/aeternity"
	"github.com/trustwallet/blockatlas/platform/aion"
	"github.com/trustwallet/blockatlas/platform/binance"
	"github.com/trustwallet/blockatlas/platform/bitcoin"
	"github.com/trustwallet/blockatlas/platform/cosmos"
	"github.com/trustwallet/blockatlas/platform/ethereum"
	"github.com/trustwallet/blockatlas/platform/fio"
	"github.com/trustwallet/blockatlas/platform/icon"
	"github.com/trustwallet/blockatlas/platform/iotex"
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

var platformList = []blockatlas.Platform{
	&binance.Platform{},
	&nimiq.Platform{},
	&ripple.Platform{},
	&stellar.Platform{CoinIndex: coin.XLM},
	&stellar.Platform{CoinIndex: coin.KIN},
	&ethereum.Platform{CoinIndex: coin.ETH},
	&ethereum.Platform{CoinIndex: coin.ETC},
	&ethereum.Platform{CoinIndex: coin.POA},
	&ethereum.Platform{CoinIndex: coin.CLO},
	&ethereum.Platform{CoinIndex: coin.GO},
	&ethereum.Platform{CoinIndex: coin.WAN},
	&ethereum.Platform{CoinIndex: coin.TOMO},
	&ethereum.Platform{CoinIndex: coin.TT},
	&tezos.Platform{},
	&aion.Platform{},
	&cosmos.Platform{},
	&icon.Platform{},
	&iotex.Platform{},
	&ontology.Platform{},
	&theta.Platform{},
	&tron.Platform{},
	&vechain.Platform{},
	&zilliqa.Platform{},
	&waves.Platform{},
	&aeternity.Platform{},
	&bitcoin.Platform{CoinIndex: coin.BTC},
	&bitcoin.Platform{CoinIndex: coin.LTC},
	&bitcoin.Platform{CoinIndex: coin.BCH},
	&bitcoin.Platform{CoinIndex: coin.DASH},
	&bitcoin.Platform{CoinIndex: coin.DOGE},
	&bitcoin.Platform{CoinIndex: coin.ZEC},
	&bitcoin.Platform{CoinIndex: coin.XZC},
	&bitcoin.Platform{CoinIndex: coin.VIA},
	&bitcoin.Platform{CoinIndex: coin.RVN},
	&bitcoin.Platform{CoinIndex: coin.QTUM},
	&bitcoin.Platform{CoinIndex: coin.GRS},
	&bitcoin.Platform{CoinIndex: coin.ZEL},
	&bitcoin.Platform{CoinIndex: coin.DCR},
	&nebulas.Platform{},
	&fio.Platform{},
	&algorand.Platform{},
}

// Platforms contains all registered platforms by handle
var Platforms map[string]blockatlas.Platform

// BlockAPIs contain platforms with block services
var BlockAPIs map[string]blockatlas.BlockAPI

// StakeAPIs contain platforms with staking services
var StakeAPIs map[string]blockatlas.StakeAPI

// CustomAPIs contain platforms with custom HTTP services
var CustomAPIs map[string]blockatlas.CustomAPI

// NamingAPIs contain platforms which support naming services
var NamingAPIs map[uint64]blockatlas.NamingServiceAPI

func Init() {
	Platforms = make(map[string]blockatlas.Platform)
	BlockAPIs = make(map[string]blockatlas.BlockAPI)
	StakeAPIs = make(map[string]blockatlas.StakeAPI)
	CustomAPIs = make(map[string]blockatlas.CustomAPI)
	NamingAPIs = make(map[uint64]blockatlas.NamingServiceAPI)

	for _, platform := range platformList {
		handle := platform.Coin().Handle
		apiKey := fmt.Sprintf("%s.api", handle)

		if !viper.IsSet(apiKey) {
			continue
		}
		if viper.GetString(apiKey) == "" {
			continue
		}

		p := logger.Params{
			"platform": handle,
			"coin":     platform.Coin(),
		}

		if _, exists := Platforms[handle]; exists {
			logger.Fatal("Duplicate handle", p)
		}

		err := platform.Init()
		if err != nil {
			logger.Error("Failed to initialize API", err, p)
		}

		Platforms[handle] = platform

		if blockAPI, ok := platform.(blockatlas.BlockAPI); ok {
			BlockAPIs[handle] = blockAPI
		}
		if stakeAPI, ok := platform.(blockatlas.StakeAPI); ok {
			StakeAPIs[handle] = stakeAPI
		}
		if customAPI, ok := platform.(blockatlas.CustomAPI); ok {
			CustomAPIs[handle] = customAPI
		}
		if namingAPI, ok := platform.(blockatlas.NamingServiceAPI); ok {
			NamingAPIs[uint64(platform.Coin().ID)] = namingAPI
		}
	}
}
