package tron

import (
	"strconv"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/tokentype"
	"github.com/trustwallet/golibs/txtype"
)

func (p *Platform) GetTokenListByAddress(address string) (txtype.TokenPage, error) {
	tokens, err := p.client.fetchAccount(address)
	if err != nil {
		return nil, err
	}
	tokenPage := make(txtype.TokenPage, 0)
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
		tokenPage = append(tokenPage, txtype.Token{
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

func (p *Platform) getTokens(ids []string) chan txtype.Token {
	tkChan := make(chan txtype.Token, len(ids))
	var wg sync.WaitGroup
	for _, id := range ids {
		wg.Add(1)
		go func(i string, c chan txtype.Token) {
			defer wg.Done()
			_ = p.getTokensChannel(i, c)
		}(id, tkChan)
	}
	wg.Wait()
	close(tkChan)
	return tkChan
}

func (p *Platform) getTokensChannel(id string, tkChan chan txtype.Token) error {
	info, err := p.client.fetchTokenInfo(id)
	if err != nil || len(info.Data) == 0 {
		return err
	}
	asset := NormalizeToken(info.Data[0])
	tkChan <- asset
	return nil
}

func NormalizeToken(info AssetInfo) txtype.Token {
	return txtype.Token{
		Name:     info.Name,
		Symbol:   strings.ToUpper(info.Symbol),
		TokenID:  strconv.Itoa(int(info.ID)),
		Coin:     coin.TRX,
		Decimals: info.Decimals,
		Type:     tokentype.TRC10,
	}
}
