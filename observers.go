package main

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/observer"
	"github.com/trustwallet/blockatlas/platform/ethereum"
	"github.com/trustwallet/blockatlas/platform/ripple"
	"net/http"
	"time"
)

func startBlockObservers() {
	dispatcher := observer.Dispatcher{
		Client: http.DefaultClient,
	}

	ethereum.SetupObserver(&dispatcher, time.Duration(viper.GetInt("ethereum.interval")))
	go ethereum.ObserveNewBlocks()

	ripple.SetupObserver(&dispatcher)
	go ripple.ObserveNewBlocs()
}
