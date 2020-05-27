package main

import (
	"encoding/json"
	"testing"
)

func TestGetProviderInfo(t *testing.T) {
	expectedJSON := `{"id":"compound","info":{"id":"compound","description":"Compound Decentralized Finance Protocol","image":"https://compound.finance/images/compound-logo.svg","website":"https://compound.finance"},"assets":[{"symbol":"DAI","chain":"ETH","description":"Compound DAI","yield_freq":15,"terms":[]},{"symbol":"USDC","chain":"ETH","description":"Compound USD Coin","yield_freq":15,"terms":[]},{"symbol":"ETH","chain":"ETH","description":"Compound Ether","yield_freq":15,"terms":[]},{"symbol":"WBTC","chain":"ETH","description":"Compound Wrapped Bitcoin","yield_freq":15,"terms":[]}]}`
	t.Run("all", func(t *testing.T) {
		res, err := GetProviderInfo()
		if err != nil {
			t.Errorf("Unexpected error %v", err)
		}
		b, _ := json.Marshal(res)
		resJSON := string(b)
		if !compareJSON(resJSON, expectedJSON) {
			t.Errorf("Wrong result, %v vs %v", resJSON, expectedJSON)
		}
	})
}

func TestGetCurrentLendingRates(t *testing.T) {
	tests := []struct {
		name         string
		inputAssets  []string
		responseJSON string
	}{
		{"wbtc", []string{"WBTC"}, `[{"asset":"WBTC","term_rates":[{"term":0.00017,"apr":2.1999999999999997}],"max_apr":2.1999999999999997}]`},
		{"all", []string{}, `[{"asset":"ETH","term_rates":[{"term":0.00017,"apr":0.8500000000000001}],"max_apr":0.8500000000000001},{"asset":"WBTC","term_rates":[{"term":0.00017,"apr":2.1999999999999997}],"max_apr":2.1999999999999997},{"asset":"DAI","term_rates":[{"term":0.00017,"apr":1.32}],"max_apr":1.32},{"asset":"USDC","term_rates":[{"term":0.00017,"apr":1.67}],"max_apr":1.67}]`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := GetCurrentLendingRates(tt.inputAssets)
			if err != nil {
				t.Errorf("Unexpected error %v", err)
			}
			b, _ := json.Marshal(res)
			resJSON := string(b)
			if !compareJSON(resJSON, tt.responseJSON) {
				t.Errorf("Wrong result, %v vs %v", resJSON, tt.responseJSON)
			}
		})
	}
}
