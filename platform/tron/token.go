package tron

import (
	log "github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/tokentype"
	"strings"
	"sync"
	"time"
)

func (p *Platform) GetTokenListByAddress(address string) (blockatlas.TokenPage, error) {
	tokens, err := p.client.fetchAccount(address)
	if err != nil {
		return nil, err
	}
	tokenPage := make(blockatlas.TokenPage, 0)
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
		tokenPage = append(tokenPage, blockatlas.Token{
			Name:     t.Name,
			Symbol:   strings.ToUpper(t.Symbol),
			Decimals: uint(t.Decimals),
			TokenID:  t.ContractAddress,
			Coin:     coin.Tron().ID,
			Type:     tokentype.TRC20,
		})
	}

	return tokenPage, nil
}

func (p *Platform) getTokens(ids []string) chan blockatlas.Token {
	tkChan := make(chan blockatlas.Token, len(ids))
	var wg sync.WaitGroup
	for _, id := range ids {
		wg.Add(1)
		go func(i string, c chan blockatlas.Token) {
			defer wg.Done()
			time.Sleep(time.Millisecond)
			err := p.getTokensChannel(i, c)
			if err != nil {
				log.Error("tron getTokens: " + i)
			}
		}(id, tkChan)
	}
	wg.Wait()
	close(tkChan)
	return tkChan
}

func (p *Platform) getTokensChannel(id string, tkChan chan blockatlas.Token) error {
	info, err := p.client.fetchTokenInfo(id)
	if err != nil || len(info.Data) == 0 {
		return err
	}
	asset := NormalizeToken(info.Data[0])
	tkChan <- asset
	return nil
}

func NormalizeToken(info AssetInfo) blockatlas.Token {
	return blockatlas.Token{
		Name:     info.Name,
		Symbol:   strings.ToUpper(info.Symbol),
		TokenID:  info.ID,
		Coin:     coin.TRX,
		Decimals: info.Decimals,
		Type:     tokentype.TRC10,
	}
}
