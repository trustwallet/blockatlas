package coin

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

//go:generate rm -f slip44.go
//go:generate go run gen.go

var Coins map[uint]Coin

// Coin is the native currency of a blockchain
type Coin struct {
	// SLIP-44 index
	Index    uint   `json:"index"`
	// Symbol of native currency
	Symbol   string `json:"symbol"`
	// Full name of native currency
	Title    string `json:"name"`
	// Project website
	Website  string `json:"link"`
	// Number of decimals
	Decimals uint   `json:"decimals"`
	// Average time between blocks
	BlockTime int `json:"blockTime"`
}

func (c Coin) String() string {
	return fmt.Sprintf("[%s] %s (#%d)", c.Symbol, c.Title, c.Index)
}

func Load(fPath string) {
	buf, err := ioutil.ReadFile(fPath)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to load coins")
	}
	var coinList []Coin
	err = json.Unmarshal(buf, &coinList)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to load coins")
	}
	Coins = make(map[uint]Coin)
	for _, coin := range coinList {
		Coins[coin.Index] = coin
	}
}
