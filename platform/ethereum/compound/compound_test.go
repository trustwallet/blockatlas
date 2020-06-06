package compound

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

const accountAddr1 string = "0x4383E7CB85743fcd95E29a528b5d1EFddAA9488f"

type TestClient struct {
}

func (c *TestClient) GetAccounts(addresses []string) ([]Account, error) {
	if len(addresses) == 1 {
		if addresses[0] == accountAddr1 {
			return []Account{
				Account{
					Address: accountAddr1,
					Tokens: []AccountCToken{
						AccountCToken{
							Address:                 "0x39aa39c021dfbae8fac545936693ac917d5e7563",
							Symbol:                  "cUSDC",
							SupplyBalanceUnderlying: precise("4.000097398638441"),
							SupplyInterest:          precise("0.000097398638441"),
						},
						AccountCToken{
							Address:                 "0xc11b1268c1a384e55c48c2391d8d480264a3a7f4",
							Symbol:                  "cWBTC",
							SupplyBalanceUnderlying: precise("0.10234"),
							SupplyInterest:          precise("0.01401"),
						},
					},
				},
			}, nil
		}
		return []Account{}, fmt.Errorf("Unknown address %v", addresses[0])
	}
	return []Account{}, fmt.Errorf("Zero or more than one address, %v", len(addresses))
}

func (c *TestClient) GetCTokensCached(tokenAddresses []string, cacheExpiry time.Duration) (CTokenResponse, error) {
	return CTokenResponse{
		CToken: []CToken{
			CToken{
				TokenAddress:     "0xc11b1268c1a384e55c48c2391d8d480264a3a7f4",
				TotalSupply:      precise("9562.9935"),
				ExchangeRate:     precise("0.0201"),
				SupplyRate:       precise("0.001717"),
				Symbol:           "cWBTC",
				Name:             "Compound Wrapped BTC",
				UnderlyingSymbol: "WBTC",
				UnderlyingName:   "Wrapped BTC",
			},
			CToken{
				TokenAddress:     "0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5",
				TotalSupply:      precise("15494642.5004"),
				ExchangeRate:     precise("0.0200"),
				SupplyRate:       precise("0.000074"),
				Symbol:           "cETH",
				Name:             "Compound Ether",
				UnderlyingSymbol: "ETH",
				UnderlyingName:   "Ether",
			},
			CToken{
				TokenAddress:     "0x39aa39c021dfbae8fac545936693ac917d5e7563",
				TotalSupply:      precise("1023919029.5655"),
				ExchangeRate:     precise("0.02104"),
				SupplyRate:       precise("0.016573"),
				Symbol:           "cUSDC",
				Name:             "Compound USD Coin",
				UnderlyingSymbol: "USDC",
				UnderlyingName:   "USD Coin",
			},
		},
	}, nil
}

func precise(value string) Precise {
	return Precise{Value: value}
}

func jsonToString(obj interface{}) string {
	b, _ := json.Marshal(obj)
	return string(b)
}

func TestGetProviderInfo(t *testing.T) {
	p := InitForTest(&TestClient{})
	res, err := p.GetProviderInfo()
	if err != nil {
		t.Errorf("Unexpected error %v", err.Error())
	}
	resJSON := jsonToString(res)
	expected := `{"id":"compound","info":{"name":"compound","description":"Compound Decentralized Finance Protocol","image":"https://compound.finance/images/compound-logo.svg","website":"https://compound.finance"},"type":"lending","assets":[{"symbol":"ETH","description":"Compound Ether","apy":0.0073999999999999995,"yield_freq":15,"total_supply":"15494642.5004","minimum_amount":"0","meta_info":{}},{"symbol":"USDC","description":"Compound USD Coin","apy":1.6573,"yield_freq":15,"total_supply":"1023919029.5655","minimum_amount":"0","meta_info":{}},{"symbol":"WBTC","description":"Compound Wrapped BTC","apy":0.1717,"yield_freq":15,"total_supply":"9562.9935","minimum_amount":"0","meta_info":{}}]}`
	if resJSON != expected {
		t.Errorf("Wrong result\n%v\n%v", resJSON, expected)
	}
}

func TestGetAssets(t *testing.T) {
	p := InitForTest(&TestClient{})
	res, err := p.GetAssets()
	if err != nil {
		t.Errorf("Unexpected error %v", err.Error())
	}
	resJSON := jsonToString(res)
	expected := `[{"symbol":"ETH","description":"Compound Ether","apy":0.0073999999999999995,"yield_freq":15,"total_supply":"15494642.5004","minimum_amount":"0","meta_info":{"defi_info":{"asset_token":{"symbol":"ETH","chain":"ETH"},"technical_token":{"symbol":"cETH","chain":"ETH","contract_address":"0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5"}}}},{"symbol":"USDC","description":"Compound USD Coin","apy":1.6573,"yield_freq":15,"total_supply":"1023919029.5655","minimum_amount":"0","meta_info":{"defi_info":{"asset_token":{"symbol":"USDC","chain":"ETH"},"technical_token":{"symbol":"cUSDC","chain":"ETH","contract_address":"0x39aa39c021dfbae8fac545936693ac917d5e7563"}}}},{"symbol":"WBTC","description":"Compound Wrapped BTC","apy":0.1717,"yield_freq":15,"total_supply":"9562.9935","minimum_amount":"0","meta_info":{"defi_info":{"asset_token":{"symbol":"WBTC","chain":"ETH"},"technical_token":{"symbol":"cWBTC","chain":"ETH","contract_address":"0xc11b1268c1a384e55c48c2391d8d480264a3a7f4"}}}}]`
	if resJSON != expected {
		t.Errorf("Wrong result\n%v\n%v", resJSON, expected)
	}
}

func TestGetAccountLendingContracts(t *testing.T) {
	tests := []struct {
		name     string
		request  blockatlas.AccountRequest
		expected string
		expError bool
	}{
		{
			name: "addr1 normal",
			request: blockatlas.AccountRequest{
				Addresses: []string{accountAddr1},
				Assets:    []string{},
			},
			expected: `[{"address":"0x4383E7CB85743fcd95E29a528b5d1EFddAA9488f","contracts":[{"asset":{"symbol":"USDC","description":"Compound USD Coin","apy":1.6573,"yield_freq":15,"total_supply":"1023919029.5655","minimum_amount":"0","meta_info":{"defi_info":{"asset_token":{"symbol":"USDC","chain":"ETH"},"technical_token":{"symbol":"cUSDC","chain":"ETH","contract_address":"0x39aa39c021dfbae8fac545936693ac917d5e7563"}}}},"current_amount":"4.0000973986"},{"asset":{"symbol":"WBTC","description":"Compound Wrapped BTC","apy":0.1717,"yield_freq":15,"total_supply":"9562.9935","minimum_amount":"0","meta_info":{"defi_info":{"asset_token":{"symbol":"WBTC","chain":"ETH"},"technical_token":{"symbol":"cWBTC","chain":"ETH","contract_address":"0xc11b1268c1a384e55c48c2391d8d480264a3a7f4"}}}},"current_amount":"0.1023400000"}]}]`,
			expError: false,
		},
		{
			name: "addr1 with asset",
			request: blockatlas.AccountRequest{
				Addresses: []string{accountAddr1},
				Assets:    []string{"USDC"},
			},
			expected: `[{"address":"0x4383E7CB85743fcd95E29a528b5d1EFddAA9488f","contracts":[{"asset":{"symbol":"USDC","description":"Compound USD Coin","apy":1.6573,"yield_freq":15,"total_supply":"1023919029.5655","minimum_amount":"0","meta_info":{"defi_info":{"asset_token":{"symbol":"USDC","chain":"ETH"},"technical_token":{"symbol":"cUSDC","chain":"ETH","contract_address":"0x39aa39c021dfbae8fac545936693ac917d5e7563"}}}},"current_amount":"4.0000973986"}]}]`,
			expError: false,
		},
		{
			name: "addr1 with assets",
			request: blockatlas.AccountRequest{
				Addresses: []string{accountAddr1},
				Assets:    []string{"WBTC", "USDC"},
			},
			expected: `[{"address":"0x4383E7CB85743fcd95E29a528b5d1EFddAA9488f","contracts":[{"asset":{"symbol":"USDC","description":"Compound USD Coin","apy":1.6573,"yield_freq":15,"total_supply":"1023919029.5655","minimum_amount":"0","meta_info":{"defi_info":{"asset_token":{"symbol":"USDC","chain":"ETH"},"technical_token":{"symbol":"cUSDC","chain":"ETH","contract_address":"0x39aa39c021dfbae8fac545936693ac917d5e7563"}}}},"current_amount":"4.0000973986"},{"asset":{"symbol":"WBTC","description":"Compound Wrapped BTC","apy":0.1717,"yield_freq":15,"total_supply":"9562.9935","minimum_amount":"0","meta_info":{"defi_info":{"asset_token":{"symbol":"WBTC","chain":"ETH"},"technical_token":{"symbol":"cWBTC","chain":"ETH","contract_address":"0xc11b1268c1a384e55c48c2391d8d480264a3a7f4"}}}},"current_amount":"0.1023400000"}]}]`,
			expError: false,
		},
		{
			name: "addr1 with wrong asset",
			request: blockatlas.AccountRequest{
				Addresses: []string{accountAddr1},
				Assets:    []string{"ETH"},
			},
			expected: `[{"address":"0x4383E7CB85743fcd95E29a528b5d1EFddAA9488f","contracts":[]}]`,
			expError: false,
		},
		{
			name: "wrong address",
			request: blockatlas.AccountRequest{
				Addresses: []string{"0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5"},
				Assets:    []string{},
			},
			expected: `null`,
			expError: true,
		},
	}
	p := InitForTest(&TestClient{})
	for _, tt := range tests {
		res, err := p.GetAccountLendingContracts(tt.request)
		if !tt.expError {
			if err != nil {
				t.Errorf("Unexpected error %v %v", tt.name, err.Error())
			}
			resJSON := jsonToString(res)
			if resJSON != tt.expected {
				t.Errorf("Wrong result %v\n%v\n%v", tt.name, resJSON, tt.expected)
			}
		} else {
			if err == nil {
				t.Errorf("Expected error, got none %v", tt.name)
			}
		}
	}
}

func TestProviderName(t *testing.T) {
	p := InitForTest(&TestClient{})
	expected := "compound"
	res := p.Name()
	if res != expected {
		t.Errorf("Wrong result %v %v", res, expected)
	}
}
