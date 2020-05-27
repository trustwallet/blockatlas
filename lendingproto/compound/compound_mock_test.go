package compound

import (
	"encoding/json"
	"testing"
)

func TestCMockAccount(t *testing.T) {
	tests := []struct {
		name         string
		input        CMAccountRequest
		responseJSON string
	}{
		{"40", CMAccountRequest{[]string{"0x12340000"}}, `{"accounts":[{"address":"0x12340000","tokens":[{"address":"0x6aabbcc000002","symbol":"USDC","supply_balance_underlying":200.45,"lifetime_supply_interest_accrued":0.45},{"address":"0x6aabbcc000001","symbol":"DAI","supply_balance_underlying":300.85,"lifetime_supply_interest_accrued":0.85}]}]}`},
		{"60", CMAccountRequest{[]string{"0x12360000"}}, `{"accounts":[{"address":"0x12360000","tokens":[{"address":"0x6aabbcc000002","symbol":"USDC","supply_balance_underlying":1001.25,"lifetime_supply_interest_accrued":2.25}]}]}`},
		{"none", CMAccountRequest{[]string{}}, `{"accounts":null}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := CMockAccount(tt.input)
			if err != nil {
				t.Errorf("Unexpected error %v", err)
			}
			b, _ := json.Marshal(res)
			resJSON := string(b)
			if resJSON != tt.responseJSON {
				t.Errorf("Wrong result, %v %v vs %v", tt.name, resJSON, tt.responseJSON)
			}
		})
	}
}

func TestCMockCToken(t *testing.T) {
	tests := []struct {
		name           string
		inputAddresses []string
		responseJSON   string
	}{
		{"usdc", []string{"0x6aabbcc000002"}, `{"cToken":[{"token_address":"0x6aabbcc000002","total_supply":"300000","exchange_rate":"1.4259","supply_rate":"0.0167","symbol":"cUSDC","name":"Compound USD Coin","underlying_symbol":"USDC","underlying_name":"Circle USD Coin"}]}`},
		{"all", []string{}, `{"cToken":[{"token_address":"0x6aabbcc000001","total_supply":"200000","exchange_rate":"1.5678","supply_rate":"0.0132","symbol":"cDAI","name":"Compound DAI","underlying_symbol":"DAI","underlying_name":"DAI"},{"token_address":"0x6aabbcc000002","total_supply":"300000","exchange_rate":"1.4259","supply_rate":"0.0167","symbol":"cUSDC","name":"Compound USD Coin","underlying_symbol":"USDC","underlying_name":"Circle USD Coin"},{"token_address":"0x6aabbcc000003","total_supply":"800000","exchange_rate":"1.2657","supply_rate":"0.0085","symbol":"cETH","name":"Compound Ether","underlying_symbol":"ETH","underlying_name":"Ether"},{"token_address":"0x6aabbcc000004","total_supply":"60000","exchange_rate":"1.0456","supply_rate":"0.0220","symbol":"cWBTC","name":"Compound Wrapped Bitcoin","underlying_symbol":"WBTC","underlying_name":"Wrpaped Bitcoin"}]}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := CMockCToken(tt.inputAddresses)
			if err != nil {
				t.Errorf("Unexpected error %v", err)
			}
			b, _ := json.Marshal(res)
			resJSON := string(b)
			if resJSON != tt.responseJSON {
				t.Errorf("Wrong result, %v vs %v", resJSON, tt.responseJSON)
			}
		})
	}
}
