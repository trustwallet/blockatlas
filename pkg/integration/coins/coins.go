// +build integration

package coins

import (
	"encoding/json"
	"fmt"
	config "github.com/trustwallet/blockatlas/pkg/integration/config"
	"io/ioutil"
	"net/http"
)

type Coin struct {
	ID         uint   `json:"id"`
	Handle     string `json:"handle"`
	Symbol     string `json:"symbol"`
	Title      string `json:"name"`
	Decimals   uint   `json:"decimals"`
	BlockTime  int    `json:"blockTime"`
	SampleAddr string `json:"sampleAddress"`
}

func GetCoins() ([]Coin, error) {
	url := fmt.Sprintf("%s/%s", config.Configuration.Server.Url, config.Configuration.Server.Coin_Path)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	read, err := ioutil.ReadAll(resp.Body)

	var coins []Coin
	err = json.Unmarshal(read, &coins)
	if err != nil {
		return nil, err
	}

	return coins, err
}
