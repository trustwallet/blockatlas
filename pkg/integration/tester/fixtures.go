// +build integration

package tester

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

type Coin struct {
	Id        int      `json:"id"`
	Handle    string   `json:"handle"`
	Apis      []string `json:"apis"`
	Addresses []string `json:"addresses"`
}

type Api struct {
	Name   string                 `json:"name"`
	Path   string                 `json:"path"`
	Schema map[string]interface{} `json:"schema,omitempty"`
}

const (
	apis  = "apis.json"
	coins = "coins.json"
)

func GetCoins() ([]Coin, error) {
	b, err := getFile(coins)
	if err != nil {
		return nil, err
	}
	var coins []Coin
	err = json.Unmarshal(b[:], &coins)
	return coins, err
}

func GetApis() (map[string]Api, error) {
	b, err := getFile(apis)
	if err != nil {
		return nil, err
	}

	apis := make(map[string]Api)
	var r []Api
	err = json.Unmarshal(b[:], &r)
	for _, api := range r {
		apis[api.Name] = api
	}
	return apis, err
}

func getFile(file string) ([]byte, error) {
	golden := filepath.Join("testdata", file)
	return ioutil.ReadFile(golden)
}
