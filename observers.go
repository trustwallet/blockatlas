package main

import (
	"github.com/trustwallet/blockatlas/observer"
	"github.com/trustwallet/blockatlas/platform/ethereum"
	"net/http"
)

func startBlockObservers() {
	dispatcher := observer.Dispatcher{
		Client: http.DefaultClient,
	}

	ethereum.Setup(&dispatcher, 3)
	go ethereum.ObserveNewBlocks()
}
