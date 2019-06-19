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
)

var platformList = []blockatlas.Platform{
	&aion.Platform{},
	&binance.Platform{},
	&cosmos.Platform{},
	&ethereum.Platform{ CoinIndex: coin.ETH, HandleStr: "ethereum" },
	&ethereum.Platform{ CoinIndex: coin.ETC, HandleStr: "classic"},
	&icon.Platform{},
	&iotex.Platform{},
	&nimiq.Platform{},
	&ontology.Platform{},
	&ripple.Platform{},
	&semux.Platform{},
	&stellar.Platform{ CoinIndex: coin.XLM, HandleStr: "stellar" },
	&stellar.Platform{ CoinIndex: coin.KIN, HandleStr: "kin" },
	&tezos.Platform{},
	&theta.Platform{},
	&tron.Platform{},
	&vechain.Platform{},
	&zilliqa.Platform{},
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
		handle := platform.Handle()

		if !viper.IsSet(fmt.Sprintf("%s.api", handle)) {
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

		log.Info("Registered platform")
	}
}
