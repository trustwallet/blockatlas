package main

import (
	"github.com/trustwallet/blockatlas/observer"
	"github.com/trustwallet/blockatlas/platform/ethereum"
	"github.com/trustwallet/blockatlas/platform/ripple"
	"net/http"
)

func startBlockObservers() {
	dispatcher := observer.Dispatcher{
		Client: http.DefaultClient,
	}

	ethereum.SetupObserver(&dispatcher, 3)
	go ethereum.ObserveNewBlocks()

	ripple.SetupObserver(&dispatcher)
	go ripple.ObserveNewBlocs()
}
