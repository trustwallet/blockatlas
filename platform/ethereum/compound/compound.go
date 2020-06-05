package compound

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

// Compound Lending Provider

type CompoundClient interface {
	GetAccounts(addresses []string) ([]Account, error)
	GetCTokensCached(tokenAddresses []string, cacheExpiry time.Duration) (CTokenResponse, error)
}

type Provider struct {
	client CompoundClient
}

const (
	providerName                   = "compound"
	cacheValiditySec time.Duration = 15 * time.Second
)

func Init(api string) *Provider {
	return &Provider{
		client: &Client{blockatlas.InitClient(api)},
	}
}

func InitForTest(client CompoundClient) *Provider {
	return &Provider{client: client}
}

func (p *Provider) Name() string {
	return providerName
}

// GetProviderInfo return static info about the lending provider, such as name and asset classes supported.
func (p *Provider) GetProviderInfo() (blockatlas.LendingProvider, error) {
	provider := blockatlas.LendingProvider{
		ID: "compound",
		Info: blockatlas.StakeValidatorInfo{
			Name:        providerName,
			Description: "Compound Decentralized Finance Protocol",
			Image:       "https://compound.finance/images/compound-logo.svg",
			Website:     "https://compound.finance",
		},
		Type:   blockatlas.ProviderTypeLending,
		Assets: make([]blockatlas.AssetInfo, 0),
	}
	assets, err := p.getAssetInfos(false)
	if err == nil {
		provider.Assets = assets
	}
	return provider, nil
}

func (p *Provider) GetAssets() ([]blockatlas.AssetInfo, error) {
	return p.getAssetInfos(true)
}

// GetAccountLendingContracts return current contract details for a given address.
// req.Assets: List asset IDs to consider, or empty for all
func (p *Provider) GetAccountLendingContracts(req blockatlas.AccountRequest) ([]blockatlas.AccountLendingContracts, error) {
	accounts, err := p.client.GetAccounts(req.Addresses)
	if err != nil {
		return nil, err
	}
	ret := make([]blockatlas.AccountLendingContracts, 0)
	for _, acc := range accounts {
		ret1 := blockatlas.AccountLendingContracts{Address: acc.Address, Contracts: make([]blockatlas.LendingContract, 0)}
		for _, t := range acc.Tokens {
			asset := p.getSymbolByCSymbol(t.Symbol)
			if len(req.Assets) > 0 && !sliceContains(asset, req.Assets) {
				continue // not requested, skip
			}
			assetInfo := blockatlas.AssetInfo{Symbol: asset}
			if ai, err := p.getAssetInfoForSymbol(asset); err == nil {
				assetInfo = ai
			}
			ret1.Contracts = append(ret1.Contracts, blockatlas.LendingContract{
				Asset:         assetInfo,
				CurrentAmount: strconv.FormatFloat(asFloat(t.SupplyBalanceUnderlying.Value), 'f', 10, 64),
			})
		}
		sort.Slice(ret1.Contracts, func(a, b int) bool { return ret1.Contracts[a].CurrentAmount > ret1.Contracts[b].CurrentAmount })
		ret = append(ret, ret1)
	}
	sort.Slice(ret, func(a, b int) bool { return ret[a].Address > ret[b].Address })
	return ret, nil
}

func getAssetInfo(t *CToken, includeMeta bool) blockatlas.AssetInfo {
	ret := blockatlas.AssetInfo{
		Symbol:         t.UnderlyingSymbol,
		Description:    t.Name,
		APY:            apyOfToken(t),
		YieldPeriod:    0,
		YieldFrequency: 15,
		TotalSupply:    blockatlas.Amount(t.TotalSupply.Value),
		MinimumAmount:  blockatlas.Amount("0"),
		LockTime:       0,
	}
	if includeMeta {
		ret.MetaInfo = blockatlas.AssetMetaInfo{
			DefiInfo: &blockatlas.DefiAssetInfo{
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

func (p *Provider) getAssetInfoForSymbol(asset string) (blockatlas.AssetInfo, error) {
	tokens, err := p.getTokensCached()
	if err != nil {
		return blockatlas.AssetInfo{}, err
	}
	token, ok := tokens[strings.ToUpper(asset)]
	if !ok {
		return blockatlas.AssetInfo{}, fmt.Errorf("Token not found %v", asset)
	}
	return getAssetInfo(&token, true), nil
}

func (p *Provider) getAssetInfos(includeMeta bool) ([]blockatlas.AssetInfo, error) {
	res := make([]blockatlas.AssetInfo, 0)
	tokens, err := p.getTokensCached()
	if err != nil {
		return res, err
	}
	for _, t := range tokens {
		res = append(res, getAssetInfo(&t, includeMeta))
	}
	sort.Slice(res, func(i, j int) bool { return res[i].Symbol < res[j].Symbol })
	return res, nil
}

func (p *Provider) getTokensCached() (map[string]CToken, error) {
	tokens := make(map[string]CToken, 30)
	res, err := p.client.GetCTokensCached(make([]string, 0), cacheValiditySec)
	if err != nil {
		return tokens, err
	}
	for _, t := range res.CToken {
		tokens[strings.ToUpper(t.UnderlyingSymbol)] = t
	}
	return tokens, nil
}

func (p *Provider) getSymbolByCSymbol(symbol string) string {
	tokens, err := p.getTokensCached()
	if err == nil {
		for s := range tokens {
			if tokens[s].Symbol == symbol {
				return s
			}
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
