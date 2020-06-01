package compound

import (
	"fmt"
	"strconv"
	"time"

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

var (
	_providerName                       = "compound"
	_cachedTokens     map[string]CToken = make(map[string]CToken, 30)
	_cachedTokenTime  time.Time
	_cacheValiditySec = 15
)

func (p *Provider) Name() string {
	return _providerName
}

// GetProviderInfo return static info about the lending provider, such as name and asset classes supported.
func (p *Provider) GetProviderInfo() (blockatlas.LendingProvider, error) {
	return blockatlas.LendingProvider{
		ID: "compound",
		Info: blockatlas.LendingProviderInfo{
			ID:          _providerName,
			Description: "Compound Decentralized Finance Protocol",
			Image:       "https://compound.finance/images/compound-logo.svg",
			Website:     "https://compound.finance",
		},
		Assets: p.getTokensNormalized(),
	}, nil
}

// GetCurrentLendingRates return current estimated yield rates for assets.  Rates are annualized.  Rates vary over time.
// assets: List asset IDs to consider, or empty for all
func (p *Provider) GetCurrentLendingRates(assets []string) (blockatlas.LendingRates, error) {
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
func (p *Provider) GetAccountLendingContracts(req blockatlas.AccountRequest) (*[]blockatlas.AccountLendingContracts, error) {
	accounts, err := p.client.GetAccounts(req.Addresses)
	if err != nil {
		return nil, err
	}
	ret := []blockatlas.AccountLendingContracts{}
	for _, acc := range accounts {
		ret1 := blockatlas.AccountLendingContracts{Address: acc.Address, Contracts: []blockatlas.LendingContract{}}
		for _, t := range acc.Tokens {
			asset := t.Symbol
			if len(req.Assets) > 0 && !sliceContains(asset, req.Assets) {
				continue // not requested, skip
			}
			// APR: no info, take general current APR
			var apr float64 = 0
			if assetInfo, err := p.getCurrentLendingRatesForAsset(asset); err == nil {
				apr = assetInfo.MaxAPR
			}
			ret1.Contracts = append(ret1.Contracts, blockatlas.LendingContract{
				Asset: t.Symbol,
				Term:  0,
				// startAmount: not available in API, derive as currentAmount - interest earn
				StartAmount:       strconv.FormatFloat(t.SupplyBalanceUnderlying-t.SupplyInterest, 'f', 10, 64),
				CurrentAmount:     strconv.FormatFloat(t.SupplyBalanceUnderlying, 'f', 10, 64),
				EndAmountEstimate: strconv.FormatFloat(t.SupplyBalanceUnderlying, 'f', 10, 64),
				CurrentAPR:        apr,
				// startTime: no info
				StartTime:   0, // no info
				CurrentTime: blockatlas.Time(time.Now().Unix()),
				EndTime:     0, // no info
			})
		}
		ret = append(ret, ret1)
	}
	return &ret, nil
}

func (p *Provider) getTokensNormalized() []blockatlas.AssetClass {
	// In compound all assets are updated with each ETH block, about each 15 seconds.  There are no predefined terms.
	tokens := p.getTokensCached()
	res := []blockatlas.AssetClass{}
	for s, t := range tokens {
		res = append(res, blockatlas.AssetClass{
			Symbol:         s,
			Chain:          "ETH",
			Description:    t.Name,
			YieldFrequency: 15,
			Terms:          []blockatlas.Term{},
		})
	}
	return res
}

// Returns a info on tokens, map by symbol
// Cached for _cacheValiditySec seconds
func (p *Provider) getTokensCached() map[string]CToken {
	now := time.Now()
	if now.Sub(_cachedTokenTime) < time.Duration(_cacheValiditySec*1000) && len(_cachedTokens) > 0 {
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

func (p *Provider) getCurrentLendingRatesForAssets(assets []string) ([]blockatlas.LendingAssetRates, error) {
	ret := []blockatlas.LendingAssetRates{}
	tokens := p.getTokensCached()
	// group by asset (symbol)
	currSymbol := ""
	var ret1 *blockatlas.LendingAssetRates = nil
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
			ret1 = &blockatlas.LendingAssetRates{Asset: currSymbol, TermRates: []blockatlas.LendingTermAPR{}, MaxAPR: 0}
		}
		apr := aprOfToken(&t)
		ret1.TermRates = append(ret1.TermRates, blockatlas.LendingTermAPR{Term: 0.00017, APR: apr})
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

func (p *Provider) getCurrentLendingRatesForAsset(asset string) (blockatlas.LendingAssetRates, error) {
	ret := blockatlas.LendingAssetRates{Asset: asset, TermRates: []blockatlas.LendingTermAPR{}, MaxAPR: 0}
	tokens := p.getTokensCached()
	token, ok := tokens[asset]
	if !ok {
		return ret, fmt.Errorf("Token not found %v", asset)
	}
	apr := aprOfToken(&token)
	ret.TermRates = append(ret.TermRates, blockatlas.LendingTermAPR{Term: 0.00017, APR: apr})
	enrichAssetRatesWithMax(&ret)
	return ret, nil
}

func enrichAssetRatesWithMax(rates *blockatlas.LendingAssetRates) {
	var max float64 = 0
	for _, r := range rates.TermRates {
		if r.APR > max {
			max = r.APR
		}
	}
	rates.MaxAPR = max
}

func aprOfToken(token *CToken) float64 {
	apr, err := strconv.ParseFloat(token.SupplyRate.Value, 64)
	if err != nil {
		return 0
	}
	return 100.0 * apr
}

func sliceContains(elem string, slice []string) bool {
	for _, e := range slice {
		if e == elem {
			return true
		}
	}
	return false
}
