package coin

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
)

//go:generate rm -f slip44.go
//go:generate go run gen.go

var Coins map[uint]Coin

// Coin is the native currency of a blockchain
type Coin struct {
	ID         uint   `yaml:"id"`         // SLIP-44 ID (e.g. 242)
	Handle     string `yaml:"trust"`      // Trust Wallet handle (e.g. nimiq)
	Symbol     string `yaml:"symbol"`     // Symbol of native currency
	Title      string `yaml:"name"`       // Full name of native currency
	Website    string `yaml:"link"`       // Project website
	Decimals   uint   `yaml:"decimals"`   // Number of decimals
	BlockTime  int    `yaml:"blockTime"`  // Average time between blocks (ms)
	SampleAddr string `yaml:"sample"`     // Random address seen on chain
}

func (c Coin) String() string {
	return fmt.Sprintf("[%s] %s (#%d)", c.Symbol, c.Title, c.ID)
}

func Load(fPath string) {
	err := load(fPath)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to load coins")
	}
}

func load(fPath string) error {
	f, err := os.Open(fPath)
	if err != nil {
		return err
	}

	var coinList []Coin

	dec := yaml.NewDecoder(f)
	err = dec.Decode(&coinList)
	if err != nil {
		return err
	}

	Coins = make(map[uint]Coin)
	for _, coin := range coinList {
		if coin.Handle == "" {
			return fmt.Errorf("coin %d has no handle", coin.ID)
		}
		Coins[coin.ID] = coin
	}

	return nil
}
