package compound

import (
	"fmt"
	"strconv"
	"time"

	"github.com/trustwallet/blockatlas/lendingproto/model"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

// Compound Lending Provider
// Compound does not use fixed terms, only open-ended, but structure is done to support different predefined terms.

type Provider struct {
	client Client
}

func Init(api string) *Provider {
	return &Provider{
		client: Client{blockatlas.InitClient(api)},
	}
}

// GetProviderInfo return static info about the lending provider, such as name and asset classes supported.
func (p *Provider) GetProviderInfo() (model.LendingProvider, error) {
	// Note: should be cached
	return model.LendingProvider{
		ID: "compound",
		Info: model.LendingProviderInfo{
			ID:          "compound",
			Description: "Compound Decentralized Finance Protocol",
			Image:       "https://compound.finance/images/compound-logo.svg",
			Website:     "https://compound.finance",
		},
		Assets: p.getTokensNormalized(),
	}, nil
}

// GetCurrentLendingRates return current estimated yield rates for assets.  Rates are annualized.  Rates vary over time.
// assets: List asset IDs to consider, or empty for all
// Note: can use the CTokenRequest compound API
func (p *Provider) GetCurrentLendingRates(assets []string) (model.LendingRates, error) {
	if len(assets) == 0 {
		// empty filter, means any; get all available assets
		tokens := p.getTokensCached()
		for t := range tokens {
			assets = append(assets, t)
		}
	}
	return p.getCurrentLendingRatesForAssets(assets)
}

// GetAccountLendingContracts return current contract details for a given address.
// req.Assets: List asset IDs to consider, or empty for all
func (p *Provider) GetAccountLendingContracts(req model.AccountRequest) (*[]model.AccountLendingContracts, error) {
	accounts, err := p.client.GetAccounts(req.Addresses)
	if err != nil {
		return nil, nil
	}
	ret := []model.AccountLendingContracts{}
	for _, acc := range accounts {
		ret1 := model.AccountLendingContracts{Address: acc.Address, Contracts: []model.LendingContract{}}
		for _, t := range acc.Tokens {
			asset := t.Symbol
			if len(req.Assets) > 0 && !sliceContains(asset, req.Assets) {
				continue
			}
			// APR: no info, take general current APR
			var apr float64 = 0
			if assetInfo, err := p.getCurrentLendingRatesForAsset(asset); err == nil {
				apr = assetInfo.MaxAPR
			}
			ret1.Contracts = append(ret1.Contracts, model.LendingContract{
				Asset: t.Symbol,
				Term:  0,
				// startAmount: not available in API, derive as currentAmount - interest earn
				StartAmount:       strconv.FormatFloat(t.SupplyBalanceUnderlying-t.SupplyInterest, 'f', 10, 64),
				CurrentAmount:     strconv.FormatFloat(t.SupplyBalanceUnderlying, 'f', 10, 64),
				EndAmountEstimate: strconv.FormatFloat(t.SupplyBalanceUnderlying, 'f', 10, 64),
				CurrentAPR:        apr,
				// startTime: no info
				StartTime:   0, // no info
				CurrentTime: model.Time(time.Now().Unix()),
				EndTime:     0, // no info
			})
		}
		ret = append(ret, ret1)
	}
	return &ret, nil
}

func (p *Provider) getTokensNormalized() []model.AssetClass {
	// In compound all assets are updated with each ETH block, about each 15 seconds.  There are no predefined terms.
	tokens := p.getTokensCached()
	res := []model.AssetClass{}
	for s, t := range tokens {
		res = append(res, model.AssetClass{
			Symbol:         s,
			Chain:          "ETH",
			Description:    t.Name,
			YieldFrequency: 15,
			Terms:          []model.Term{},
		})
	}
	return res
}

var (
	_cachedTokens    map[string]CToken = make(map[string]CToken, 30)
	_cachedTokenTime time.Time
)

// Returns a info on tokens, map by symbol
// Cached for few seconds
func (p *Provider) getTokensCached() map[string]CToken {
	now := time.Now()
	if now.Sub(_cachedTokenTime) < 30000 && len(_cachedTokens) > 0 {
		// cached and recent
		return _cachedTokens
	}
	// rertieve and cache
	_cachedTokens = make(map[string]CToken, 30)
	res, err := p.client.GetTokens([]string{})
	now = time.Now()
	if err != nil {
		return _cachedTokens
	}
	for _, t := range res.CToken {
		_cachedTokens[t.UnderlyingSymbol] = t
	}
	_cachedTokenTime = now
	return _cachedTokens
}

/*
func (p *Provider) getTokens() map[string]tokenInfo {
	ret := make(map[string]tokenInfo)
	res, err := p.client.GetTokens([]string{})
	if err != nil {
		return ret
	}
	for _, t := range res.CToken {
		ret[t.UnderlyingSymbol] = tokenInfo{t.TokenAddress, t.Name}
	}
	return ret
}

// Note: should work from cached data
func (p *Provider) addressOfToken(symbol string) (string, bool) {
	tokens := p.getTokens()
	tokenInfo, ok := tokens[symbol]
	if !ok {
		return "", false
	}
	return tokenInfo.address, true
}
*/

func (p *Provider) getCurrentLendingRatesForAssets(assets []string) ([]model.LendingAssetRates, error) {
	ret := []model.LendingAssetRates{}
	tokens := p.getTokensCached()
	// group by asset (symbol)
	currSymbol := ""
	var ret1 *model.LendingAssetRates = nil
	for _, t := range tokens {
		symbol := t.UnderlyingSymbol
		if !sliceContains(symbol, assets) {
			continue
		}
		if len(currSymbol) > 0 && currSymbol != symbol && ret1 != nil {
			// close previous
			enrichAssetRatesWithMax(ret1)
			ret = append(ret, *ret1)
			ret1 = nil
			currSymbol = ""
		}
		if len(currSymbol) == 0 {
			// start new
			currSymbol = symbol
			ret1 = &model.LendingAssetRates{Asset: currSymbol, TermRates: []model.LendingTermAPR{}, MaxAPR: 0}
		}
		apr, err := strconv.ParseFloat(t.SupplyRate.Value, 64)
		if err != nil {
			apr = 0
		} else {
			apr = 100.0 * apr
		}
		ret1.TermRates = append(ret1.TermRates, model.LendingTermAPR{Term: 0.00017, APR: apr})
	}
	if ret1 != nil {
		// close previous
		enrichAssetRatesWithMax(ret1)
		ret = append(ret, *ret1)
		ret1 = nil
		currSymbol = ""
	}
	return ret, nil
}

func sliceContains(elem string, slice []string) bool {
	for _, e := range slice {
		if e == elem {
			return true
		}
	}
	return false
}

func (p *Provider) getCurrentLendingRatesForAsset(asset string) (model.LendingAssetRates, error) {
	ret := model.LendingAssetRates{Asset: asset, TermRates: []model.LendingTermAPR{}, MaxAPR: 0}
	tokens := p.getTokensCached()
	token, ok := tokens[asset]
	if !ok {
		return ret, fmt.Errorf("Token not found %v", asset)
	}
	apr, err := strconv.ParseFloat(token.SupplyRate.Value, 64)
	if err != nil {
		apr = 0
	} else {
		apr = 100.0 * apr
	}
	ret.TermRates = append(ret.TermRates, model.LendingTermAPR{Term: 0.00017, APR: apr})
	enrichAssetRatesWithMax(&ret)
	return ret, nil
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
