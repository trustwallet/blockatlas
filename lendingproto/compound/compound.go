package compound

import (
	"fmt"
	"strconv"
	"time"

	"github.com/trustwallet/blockatlas/lendingproto/model"
)

// Compound Lending Provider
// Compound does not use fixed terms, only open-ended, but structure is done to support different predefined terms.

// GetProviderInfo return static info about the lending provider, such as name and asset classes supported.
func GetProviderInfo() (model.LendingProvider, error) {
	// Note: should be cached
	return model.LendingProvider{
		"compound",
		model.LendingProviderInfo{
			"compound",
			"Compound Decentralized Finance Protocol",
			"https://compound.finance/images/compound-logo.svg",
			"https://compound.finance",
		},
		getTokensNormalized(),
	}, nil
}

// GetCurrentLendingRates return current estimated yield rates for assets.  Rates are annualized.  Rates vary over time.
// assets: List asset IDs to consider, or empty for all
// Note: can use the CTokenRequest compound API
func GetCurrentLendingRates(assets []string) (model.LendingRates, error) {
	res := model.LendingRates{}
	if len(assets) == 0 {
		// empty filter, get all available assets
		tokens := getTokens()
		for t := range tokens {
			assets = append(assets, t)
		}
	}
	for _, asset := range assets {
		if rates, err := getCurrentLendingRatesForAsset(asset); err == nil {
			res = append(res, rates)
		}
	}
	return res, nil
}

// GetAccountLendingContracts return current contract details for a given address.
// req.Assets: List asset IDs to consider, or empty for all
func GetAccountLendingContracts(req model.AccountRequest) (*[]model.AccountLendingContracts, error) {
	now := model.Time(time.Now().Unix())
	res := []model.AccountLendingContracts{}
	if len(req.Addresses) == 0 {
		return nil, fmt.Errorf("Missing addresses")
	}
	for _, address := range req.Addresses {
		res1 := model.AccountLendingContracts{address, []model.LendingContract{}}
		contracts, _ := CMockAccount(CMAccountRequest{[]string{address}})
		for _, sc := range contracts.Account {
			for _, t := range sc.Tokens {
				asset := t.Symbol
				if !matchAsset(asset, req.Assets) {
					continue
				}
				// APR: no info, take general current APR
				var apr float64 = 0
				if assetInfo, err := getCurrentLendingRatesForAsset(asset); err == nil {
					apr = assetInfo.MaxAPR
				}
				res1.Contracts = append(res1.Contracts, model.LendingContract{
					t.Symbol,
					0, // term
					// startAmount: not available in API, derive as currentAmount - interest earn
					strconv.FormatFloat(t.SupplyBalanceUnderlying-t.SupplyInterest, 'f', 10, 64),
					strconv.FormatFloat(t.SupplyBalanceUnderlying, 'f', 10, 64),
					strconv.FormatFloat(t.SupplyBalanceUnderlying, 'f', 10, 64),
					apr,
					// startTime: no info, use current time
					now,
					now,
					now,
				})
			}
		}
		res = append(res, res1)
	}
	return &res, nil
}

type tokenInfo struct {
	address string
	name    string
}

func getTokensNormalized() []model.AssetClass {
	// In compound all assets are updated with each ETH block, about each 15 seconds.  There are no predefined terms.
	tokens := getTokens()
	res := []model.AssetClass{}
	for s, t := range tokens {
		res = append(res, model.AssetClass{s, "ETH", t.name, 15, []model.Term{}})
	}
	return res
}

// Returns a info on tokens, map by symbol
// Note: this should be cached
func getTokens() map[string]tokenInfo {
	res := make(map[string]tokenInfo)
	tokens, err := CMockCToken([]string{})
	if err != nil {
		return res
	}
	for _, t := range tokens.CToken {
		res[t.UnderlyingSymbol] = tokenInfo{t.TokenAddress, t.Name}
	}
	return res
}

// Note: should work from cached data
func addressOfToken(symbol string) (string, bool) {
	tokens := getTokens()
	tokenInfo, ok := tokens[symbol]
	if !ok {
		return "", false
	}
	return tokenInfo.address, true
}

func getCurrentLendingRatesForAsset(asset string) (model.LendingAssetRates, error) {
	res := model.LendingAssetRates{asset, []model.LendingTermAPR{}, 0}
	address, ok := addressOfToken(asset)
	if !ok {
		return res, fmt.Errorf("Token not found %v", asset)
	}
	tokens, err := CMockCToken([]string{address})
	if err != nil {
		return res, fmt.Errorf("Token not found %v", asset)
	}
	for _, t := range tokens.CToken {
		apr, err := strconv.ParseFloat(t.SupplyRate, 64)
		if err != nil {
			apr = 0
		} else {
			apr = 100.0 * apr
		}
		res.TermRates = append(res.TermRates, model.LendingTermAPR{0.00017, apr})
	}
	enrichAssetRatesWithMax(&res)
	return res, nil
}

func enrichAssetRatesWithMax(rates *model.LendingAssetRates) {
	var max float64 = 0
	for _, r := range rates.TermRates {
		if r.APR > max {
			max = r.APR
		}
	}
	rates.MaxAPR = max
}

func matchAsset(asset string, assets []string) bool {
	if len(assets) == 0 {
		return true
	}
	for _, a := range assets {
		if asset == a {
			return true
		}
	}
	return false
}
