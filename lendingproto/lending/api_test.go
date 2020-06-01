package lending

import (
	"encoding/json"
	"testing"

	"github.com/trustwallet/blockatlas/lendingproto/compound"
	"github.com/trustwallet/blockatlas/lendingproto/model"
)

func TestGetProviders(t *testing.T) {
	res, err := GetProviders()
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	expResJSON := `[{"id":"compound","info":{"id":"compound","description":"Compound Decentralized Finance Protocol","image":"https://compound.finance/images/compound-logo.svg","website":"https://compound.finance"},"assets":[{"symbol":"DAI","chain":"ETH","description":"Compound DAI","yield_freq":15,"terms":[]},{"symbol":"USDC","chain":"ETH","description":"Compound USD Coin","yield_freq":15,"terms":[]},{"symbol":"ETH","chain":"ETH","description":"Compound Ether","yield_freq":15,"terms":[]},{"symbol":"WBTC","chain":"ETH","description":"Compound Wrapped Bitcoin","yield_freq":15,"terms":[]}]}]`
	b, _ := json.Marshal(res)
	resJSON := string(b)
	if !compound.CompareJSON(resJSON, expResJSON, "") {
		t.Errorf("Wrong result, %v vs %v", resJSON, expResJSON)
	}
}

func TestGetRates(t *testing.T) {
	tests := []struct {
		name     string
		provider string
		assets   []string
		expJSON  string
	}{
		{"all", "compound", []string{}, `{"provider":"compound","rates":[{"asset":"WBTC","term_rates":[{"term":0.00017,"apr":2.1999999999999997}],"max_apr":2.1999999999999997},{"asset":"DAI","term_rates":[{"term":0.00017,"apr":1.32}],"max_apr":1.32},{"asset":"USDC","term_rates":[{"term":0.00017,"apr":1.67}],"max_apr":1.67},{"asset":"ETH","term_rates":[{"term":0.00017,"apr":0.8500000000000001}],"max_apr":0.8500000000000001}]}`},
		{"usdc", "compound", []string{"USDC"}, `{"provider":"compound","rates":[{"asset":"USDC","term_rates":[{"term":0.00017,"apr":1.67}],"max_apr":1.67}]}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := GetRates(tt.provider, model.RatesRequest{Assets: tt.assets})
			if err != nil {
				t.Errorf("Unexpected error %v %v", tt.name, err)
			}
			b, _ := json.Marshal(res)
			resJSON := string(b)
			if !compound.CompareJSON(resJSON, tt.expJSON, "") {
				t.Errorf("Wrong result, %v %v vs %v", tt.name, resJSON, tt.expJSON)
			}
		})
	}
}

func TestGetAccount(t *testing.T) {
	tests := []struct {
		name     string
		provider string
		req      model.AccountRequest
		expJSON  string
	}{
		{"40", "compound", model.AccountRequest{Addresses: []string{"0x12340000"}}, `[{"address":"0x12340000","contracts":[{"asset":"USDC","term":0,"start_amount":"200.0000000000","current_amount":"200.4500000000","end_amount_estimate":"200.4500000000","current_apr":1.67,"start_time":0,"end_time":0},{"asset":"DAI","term":0,"start_amount":"300.0000000000","current_amount":"300.8500000000","end_amount_estimate":"300.8500000000","current_apr":1.32,"start_time":0,"end_time":0}]}]`},
		{"60", "compound", model.AccountRequest{Addresses: []string{"0x12360000"}}, `[{"address":"0x12360000","contracts":[{"asset":"USDC","term":0,"start_amount":"999.0000000000","current_amount":"1001.2500000000","end_amount_estimate":"1001.2500000000","current_apr":1.67,"start_time":0,"end_time":0}]}]`},
		{"40_usdc", "compound", model.AccountRequest{Addresses: []string{"0x12340000"}, Assets: []string{"USDC"}}, `[{"address":"0x12340000","contracts":[{"asset":"USDC","term":0,"start_amount":"200.0000000000","current_amount":"200.4500000000","end_amount_estimate":"200.4500000000","current_apr":1.67,"start_time":0,"end_time":0}]}]`},
		{"40_60", "compound", model.AccountRequest{Addresses: []string{"0x12340000", "0x12360000"}}, `[{"address":"0x12340000","contracts":[{"asset":"USDC","term":0,"start_amount":"200.0000000000","current_amount":"200.4500000000","end_amount_estimate":"200.4500000000","current_apr":1.67,"start_time":0,"end_time":0},{"asset":"DAI","term":0,"start_amount":"300.0000000000","current_amount":"300.8500000000","end_amount_estimate":"300.8500000000","current_apr":1.32,"start_time":0,"end_time":0}]},{"address":"0x12360000","contracts":[{"asset":"USDC","term":0,"start_amount":"999.0000000000","current_amount":"1001.2500000000","end_amount_estimate":"1001.2500000000","current_apr":1.67,"start_time":0,"end_time":0}]}]`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := GetAccounts(tt.provider, tt.req)
			if err != nil {
				t.Errorf("Unexpected error %v %v", tt.name, err)
			}
			b, _ := json.Marshal(res)
			resJSON := string(b)
			if !compound.CompareJSON(resJSON, tt.expJSON, "current_time") {
				t.Errorf("Wrong result, %v %v vs %v", tt.name, resJSON, tt.expJSON)
			}
		})
	}
}
