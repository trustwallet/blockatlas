package main

import (
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/market"
	"github.com/trustwallet/blockatlas/storage"
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
	market.InitRates(cache)
	market.InitMarkets(cache)
	<-make(chan struct{})
}
