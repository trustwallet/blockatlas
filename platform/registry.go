package platform

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/platform/aion"
	"github.com/trustwallet/blockatlas/platform/binance"
	"github.com/trustwallet/blockatlas/platform/cosmos"
	"github.com/trustwallet/blockatlas/platform/ethereum"
	"github.com/trustwallet/blockatlas/platform/icon"
	"github.com/trustwallet/blockatlas/platform/iotex"
	"github.com/trustwallet/blockatlas/platform/nimiq"
	"github.com/trustwallet/blockatlas/platform/ontology"
	"github.com/trustwallet/blockatlas/platform/ripple"
	"github.com/trustwallet/blockatlas/platform/semux"
	"github.com/trustwallet/blockatlas/platform/stellar"
	"github.com/trustwallet/blockatlas/platform/tezos"
	"github.com/trustwallet/blockatlas/platform/theta"
	"github.com/trustwallet/blockatlas/platform/tron"
	"github.com/trustwallet/blockatlas/platform/vechain"
	"github.com/trustwallet/blockatlas/platform/zilliqa"
	"github.com/trustwallet/blockatlas/platform/waves"
)

var platformList = []blockatlas.Platform{
	&binance.Platform{},
	&nimiq.Platform{},
	&ripple.Platform{},
	&stellar.Platform{ CoinIndex: coin.XLM },
	&stellar.Platform{ CoinIndex: coin.KIN },
	&ethereum.Platform{ CoinIndex: coin.ETH },
	&ethereum.Platform{ CoinIndex: coin.ETC },
	&ethereum.Platform{ CoinIndex: coin.POA },
	&ethereum.Platform{ CoinIndex: coin.CLO },
	&ethereum.Platform{ CoinIndex: coin.GO },
	&ethereum.Platform{ CoinIndex: coin.WAN },
	&ethereum.Platform{ CoinIndex: coin.TOMO },
	&ethereum.Platform{ CoinIndex: coin.TT },
	&tezos.Platform{},
	&aion.Platform{},
	&cosmos.Platform{},
	&icon.Platform{},
	&iotex.Platform{},
	&ontology.Platform{},
	&semux.Platform{},
	&theta.Platform{},
	&tron.Platform{},
	&vechain.Platform{},
	&zilliqa.Platform{},
	&waves.Platform{},
}

// Platforms contains all registered platforms by handle
var Platforms map[string]blockatlas.Platform

// TxAPIs contains platforms with transaction services
var TxAPIs map[string]blockatlas.TxAPI

// BlockAPIs contains platforms with block services
var BlockAPIs map[string]blockatlas.BlockAPI

// CustomAPIs contains platforms with custom HTTP services
var CustomAPIs map[string]blockatlas.CustomAPI

func Init() {
	Platforms  = make(map[string]blockatlas.Platform)
	TxAPIs     = make(map[string]blockatlas.TxAPI)
	BlockAPIs  = make(map[string]blockatlas.BlockAPI)
	CustomAPIs = make(map[string]blockatlas.CustomAPI)

	for _, platform := range platformList {
		handle := platform.Coin().Handle
		apiKey := fmt.Sprintf("%s.api", handle)

		if !viper.IsSet(apiKey) {
			continue
		}
		if viper.GetString(apiKey) == "" {
			continue
		}

		log := logrus.WithFields(logrus.Fields{
			"platform": handle,
			"coin": platform.Coin(),
		})

		if _, exists := Platforms[handle]; exists {
			log.Fatal("Duplicate handle")
		}

		err := platform.Init()
		if err != nil {
			log.WithError(err).Fatal("Failed to initialize API")
		}

		Platforms[handle] = platform

		if txAPI, ok := platform.(blockatlas.TxAPI); ok {
			TxAPIs[handle] = txAPI
		}
		if blockAPI, ok := platform.(blockatlas.BlockAPI); ok {
			BlockAPIs[handle] = blockAPI
		}
		if customAPI, ok := platform.(blockatlas.CustomAPI); ok {
			CustomAPIs[handle] = customAPI
		}
	}
}
