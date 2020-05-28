package compound

import (
	"encoding/json"
	"testing"

	"github.com/trustwallet/blockatlas/lendingproto/model"
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
		if !CompareJSON(resJSON, expectedJSON) {
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
			if !CompareJSON(resJSON, tt.responseJSON) {
				t.Errorf("Wrong result, %v vs %v", resJSON, tt.responseJSON)
			}
		})
	}
}

func TestGetAccountLendingContracts(t *testing.T) {
	tests := []struct {
		name    string
		reqJson string
		expJson string
	}{
		{"40", `{"addresses":["0x12340000"]}`, `[{"address":"0x12340000","contracts":[{"asset":"USDC","term":0,"start_amount":"200.0000000000","current_amount":"200.4500000000","end_amount_estimate":"200.4500000000","current_apr":1.67,"start_time":0,"end_time":0},{"asset":"DAI","term":0,"start_amount":"300.0000000000","current_amount":"300.8500000000","end_amount_estimate":"300.8500000000","current_apr":1.32,"start_time":0,"end_time":0}]}]`},
		{"60", `{"addresses":["0x12360000"]}`, `[{"address":"0x12360000","contracts":[{"asset":"USDC","term":0,"start_amount":"999.0000000000","current_amount":"1001.2500000000","end_amount_estimate":"1001.2500000000","current_apr":1.67,"start_time":0,"end_time":0}]}]`},
		{"40_usdc", `{"addresses":["0x12340000"], "assets":["USDC"]}`, `[{"address":"0x12340000","contracts":[{"asset":"USDC","term":0,"start_amount":"200.0000000000","current_amount":"200.4500000000","end_amount_estimate":"200.4500000000","current_apr":1.67,"start_time":0,"end_time":0}]}]`},
		{"40_60", `{"addresses":["0x12340000", "0x12360000"],"assets":[]}`, `[{"address":"0x12340000","contracts":[{"asset":"USDC","term":0,"start_amount":"200.0000000000","current_amount":"200.4500000000","end_amount_estimate":"200.4500000000","current_apr":1.67,"start_time":0,"end_time":0},{"asset":"DAI","term":0,"start_amount":"300.0000000000","current_amount":"300.8500000000","end_amount_estimate":"300.8500000000","current_apr":1.32,"start_time":0,"end_time":0}]},{"address":"0x12360000","contracts":[{"asset":"USDC","term":0,"start_amount":"999.0000000000","current_amount":"1001.2500000000","end_amount_estimate":"1001.2500000000","current_apr":1.67,"start_time":0,"end_time":0}]}]`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req model.AccountRequest
			if err := json.Unmarshal([]byte(tt.reqJson), &req); err != nil {
				t.Errorf("Input json unmarshal error %v", err)
			}
			res, err := GetAccountLendingContracts(req)
			if err != nil {
				t.Errorf("Unexpected error %v %v", tt.name, err)
			}
			b, _ := json.Marshal(res)
			resJSON := string(b)
			if !CompareJSON(resJSON, tt.expJson) {
				t.Errorf("Wrong result, %v %v vs %v", tt.name, resJSON, tt.expJson)
			}
		})
	}
}
