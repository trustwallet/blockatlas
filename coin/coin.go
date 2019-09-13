package coin

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"gopkg.in/yaml.v2"
	"os"
)

//go:generate rm -f slip44.go
//go:generate go run gen.go

var Coins map[uint]Coin

// Coin is the native currency of a blockchain
type Coin struct {
	ID               uint   `yaml:"id" json:"id"`                             // SLIP-44 ID (e.g. 242)
	Handle           string `yaml:"handle" json:"handle"`                     // Trust Wallet handle (e.g. nimiq)
	Symbol           string `yaml:"symbol" json:"symbol"`                     // Symbol of native currency
	Title            string `yaml:"name" json:"name"`                         // Full name of native currency
	Decimals         uint   `yaml:"decimals" json:"decimals"`                 // Number of decimals
	BlockTime        int    `yaml:"blockTime" json:"blockTime"`               // Average time between blocks (ms)
	MinConfirmations int64  `yaml:"minConfirmations" json:"minConfirmations"` // Number of confirmations before parsing a block
	SampleAddr       string `yaml:"sampleAddress" json:"sampleAddress"`       // Random address seen on chain
}

func (c Coin) String() string {
	return fmt.Sprintf("[%s] %s (#%d)", c.Symbol, c.Title, c.ID)
}

func Load(coinPath string) {
	err := load(coinPath)
	if err != nil {
		logger.Fatal("Failed to load coins at path", err, logger.Params{"coinPath": coinPath})
	}
}

func load(coinPath string) error {
	coin, err := os.Open(coinPath)
	if err != nil {
		return err
	}

	var coinList []Coin

	dec := yaml.NewDecoder(coin)
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
	println("Coins loaded successfully")

	return nil
}
