package tron

import (
	"strconv"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/types"
)

func (p *Platform) GetTokenListByAddress(address string) (types.TokenPage, error) {
	tokens, err := p.client.fetchAccount(address)
	if err != nil {
		return nil, err
	}
	tokenPage := make(types.TokenPage, 0)
	if len(tokens.Data) == 0 {
		return tokenPage, nil
	}

	var tokenIds []string
	for _, v := range tokens.Data[0].AssetsV2 {
		tokenIds = append(tokenIds, v.Key)
	}

	tokensChan := p.getTokens(tokenIds)
	for info := range tokensChan {
		tokenPage = append(tokenPage, info)
	}

	trc20Tokens, err := p.explorerClient.fetchAllTRC20Tokens(address)
	if err != nil {
		log.Error("Explorer error" + err.Error())
	}

	for _, t := range trc20Tokens {
		tokenPage = append(tokenPage, types.Token{
			Name:     t.Name,
			Symbol:   strings.ToUpper(t.Symbol),
			Decimals: uint(t.Decimals),
			TokenID:  t.ContractAddress,
			Coin:     coin.Tron().ID,
			Type:     types.TRC20,
		})
	}

	return tokenPage, nil
}

func (p *Platform) getTokens(ids []string) chan types.Token {
	tkChan := make(chan types.Token, len(ids))
	var wg sync.WaitGroup
	for _, id := range ids {
		wg.Add(1)
		go func(i string, c chan types.Token) {
			defer wg.Done()
			_ = p.getTokensChannel(i, c)
		}(id, tkChan)
	}
	wg.Wait()
	close(tkChan)
	return tkChan
}

func (p *Platform) getTokensChannel(id string, tkChan chan types.Token) error {
	info, err := p.client.fetchTokenInfo(id)
	if err != nil || len(info.Data) == 0 {
		return err
	}
	asset := NormalizeToken(info.Data[0])
	tkChan <- asset
	return nil
}

func NormalizeToken(info AssetInfo) types.Token {
	return types.Token{
		Name:     info.Name,
		Symbol:   strings.ToUpper(info.Symbol),
		TokenID:  strconv.Itoa(int(info.ID)),
		Coin:     coin.TRX,
		Decimals: info.Decimals,
		Type:     types.TRC10,
	}
}
