package compound

import (
	"fmt"
	"strconv"
	"time"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

// Compound Lending Provider

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
		Type:   blockatlas.ProviderTypeLending,
		Assets: p.getAssetInfos(false),
	}, nil
}

// GetCurrentLendingRates return current estimated yield rates for assets.  Rates are annualized.  Rates vary over time.
// assets: List asset IDs to consider, or empty for all
func (p *Provider) GetCurrentLendingRates(assets []string) ([]blockatlas.AssetInfo, error) {
	if len(assets) == 0 {
		// empty filter, means any; get all available assets
		assets = p.getAssetSymbols()
	}
	return p.getAssetInfosFiltered(assets), nil
}

// GetAccountLendingContracts return current contract details for a given address.
// req.Assets: List asset IDs to consider, or empty for all
func (p *Provider) GetAccountLendingContracts(req blockatlas.AccountRequest) ([]blockatlas.AccountLendingContracts, error) {
	accounts, err := p.client.GetAccounts(req.Addresses)
	if err != nil {
		return nil, err
	}
	ret := []blockatlas.AccountLendingContracts{}
	for _, acc := range accounts {
		ret1 := blockatlas.AccountLendingContracts{Address: acc.Address, Contracts: []blockatlas.LendingContract{}}
		for _, t := range acc.Tokens {
			asset := p.getUnderlyingSymbol(t.Symbol)
			if len(req.Assets) > 0 && !sliceContains(asset, req.Assets) {
				continue // not requested, skip
			}
			assetInfo := blockatlas.AssetInfo{Symbol: asset}
			if ai, err := p.getAssetInfosForAsset(asset); err == nil {
				assetInfo = ai
			}
			ret1.Contracts = append(ret1.Contracts, blockatlas.LendingContract{
				Asset:         assetInfo,
				CurrentAmount: strconv.FormatFloat(asFloat(t.SupplyBalanceUnderlying.Value), 'f', 10, 64),
			})
		}
		ret = append(ret, ret1)
	}
	return ret, nil
}

func (p *Provider) getAssetSymbols() []string {
	tokens := p.getTokensCached()
	ret := []string{}
	for s := range tokens {
		ret = append(ret, s)
	}
	return ret
}

func getAssetInfo(t *CToken, includeMeta bool) blockatlas.AssetInfo {
	ret := blockatlas.AssetInfo{
		Symbol:         t.UnderlyingSymbol,
		Description:    t.Name,
		APY:            apyOfToken(t),
		YieldPeriod:    0,
		YieldFrequency: 15,
		TotalSupply:    t.TotalSupply.Value,
		MinimumAmount:  "0",
	}
	if includeMeta {
		ret.MetaInfo = blockatlas.AssetMetaInfo{
			DefiInfo: blockatlas.DefiAssetInfo{
				AssetToken: blockatlas.DefiTokenInfo{
					Symbol: t.UnderlyingSymbol,
					Chain:  Chain,
				},
				TechnicalToken: blockatlas.DefiTokenInfo{
					Symbol:          t.Symbol,
					Chain:           Chain,
					ContractAddress: t.TokenAddress,
				},
			},
		}
	}
	return ret
}

func (p *Provider) getAssetInfosForAsset(asset string) (blockatlas.AssetInfo, error) {
	tokens := p.getTokensCached()
	token, ok := tokens[asset]
	if !ok {
		return blockatlas.AssetInfo{}, fmt.Errorf("Token not found %v", asset)
	}
	return getAssetInfo(&token, true), nil
}

func (p *Provider) getAssetInfos(includeMeta bool) []blockatlas.AssetInfo {
	// In compound all assets are updated with each ETH block, about each 15 seconds.
	tokens := p.getTokensCached()
	res := []blockatlas.AssetInfo{}
	for _, t := range tokens {
		res = append(res, getAssetInfo(&t, includeMeta))
	}
	return res
}

func (p *Provider) getAssetInfosFiltered(assets []string) []blockatlas.AssetInfo {
	tokens := p.getTokensCached()
	res := []blockatlas.AssetInfo{}
	for s, t := range tokens {
		if !sliceContains(s, assets) {
			continue
		}
		res = append(res, getAssetInfo(&t, true))
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

func (p *Provider) getUnderlyingSymbol(symbol string) string {
	tokens := p.getTokensCached()
	for s := range tokens {
		if tokens[s].Symbol == symbol {
			return s
		}
	}
	return ""
}

func apyOfToken(token *CToken) float64 {
	return 100.0 * asFloat(token.SupplyRate.Value)
}

func asFloat(value string) float64 {
	valF, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}
	return valF
}

func sliceContains(elem string, slice []string) bool {
	for _, e := range slice {
		if e == elem {
			return true
		}
	}
	return false
}
