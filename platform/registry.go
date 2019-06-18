package platform

import (
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/platform/nimiq"
)

var platformList = []blockatlas.Platform{
	new(nimiq.Platform),
}

// Platforms contains all registered platforms by handle
var Platforms map[string]blockatlas.Platform

// TxAPIs contains platforms with transaction services
var TxAPIs map[string]blockatlas.TxAPI

// BlockAPIs contains platforms with block services
var BlockAPIs map[string]blockatlas.BlockAPI

// CustomAPIs contains platforms with custom HTTP services
var CustomAPIs map[string]blockatlas.CustomAPI

func init() {
	Platforms  = make(map[string]blockatlas.Platform)
	TxAPIs     = make(map[string]blockatlas.TxAPI)
	BlockAPIs  = make(map[string]blockatlas.BlockAPI)
	CustomAPIs = make(map[string]blockatlas.CustomAPI)

	for _, platform := range platformList {
		handle := platform.Handle()
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
