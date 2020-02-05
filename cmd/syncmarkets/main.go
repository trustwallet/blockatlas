package main

import (
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/storage"
	"github.com/trustwallet/blockatlas/syncmarkets"
)

const (
	defaultConfigPath = "../../config.yml"
)

var (
	confPath string
	cache    *storage.Storage
)

func init() {
	_, confPath, _, cache = internal.InitAPIWithRedis("", defaultConfigPath)
}

func main() {
	syncmarkets.InitRates(cache)
	syncmarkets.InitMarkets(cache)
	<-make(chan struct{})
}
