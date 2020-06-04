package tron

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"sync"
)

func (p *Platform) GetTokenListByAddress(address string) (blockatlas.TokenPage, error) {
	tokens, err := p.client.GetAccount(address)
	if err != nil {
		return nil, err
	}
	tokenPage := make(blockatlas.TokenPage, 0)
	if len(tokens.Data) == 0 {
		return tokenPage, nil
	}

	var trc10TokenIds []string
	for _, trc10 := range tokens.Data[0].AssetsV2 {
		if trc10.Value > 0 {
			trc10TokenIds = append(trc10TokenIds, trc10.Key)
		}
	}

	for trc10Info := range p.getTRC10Tokens(trc10TokenIds) {
		tokenPage = append(tokenPage, trc10Info)
	}

	var trc20TokenIds []string
	for _, value := range tokens.Data[0].TRC20 {
		for k := range value {
			if value[k] != "0" {
				trc20TokenIds = append(trc20TokenIds, k)
			}
			continue
		}
	}

	for trc20Info := range p.getTRC20Tokens(trc20TokenIds, address) {
		tokenPage = append(tokenPage, trc20Info)
	}

	return tokenPage, nil
}

func (p *Platform) getTRC10Tokens(ids []string) chan blockatlas.Token {
	tkChan := make(chan blockatlas.Token, len(ids))
	var wg sync.WaitGroup
	for _, id := range ids {
		wg.Add(1)
		go func(i string, c chan blockatlas.Token) {
			defer wg.Done()
			err := p.getTRC10TokenChannel(i, c)
			if err != nil {
				logger.Error(err)
			}
		}(id, tkChan)
	}
	wg.Wait()
	close(tkChan)
	return tkChan
}

func (p *Platform) getTRC10TokenChannel(id string, tkChan chan blockatlas.Token) error {
	info, err := p.client.getTokenInfo(id)
	if err != nil || len(info.Data) == 0 {
		logger.Error(err, "getTRC10TokenChannel: invalid or missing token")
		return err
	}

	asset := normalizeTRC10Token(info.Data[0])
	tkChan <- asset
	return nil
}

func (p *Platform) getTRC20Tokens(ids []string, address string) chan blockatlas.Token {
	tkChan := make(chan blockatlas.Token, len(ids))
	var wg sync.WaitGroup
	for _, id := range ids {
		wg.Add(1)
		go func(i string, a string, c chan blockatlas.Token) {
			defer wg.Done()
			err := p.getTRC20TokenChannel(i, a, c)
			if err != nil {
				logger.Error(err)
			}
		}(id, address, tkChan)
	}
	wg.Wait()
	close(tkChan)
	return tkChan
}

func (p *Platform) getTRC20TokenChannel(id, address string, tkChan chan blockatlas.Token) error {
	txs, err := p.client.getTRC20TxsOfAddress(address, id, 1)
	if err != nil || len(txs) == 0 {
		logger.Error("getTRC20TokenChannel: invalid or missing token", errors.Params{"err": err, "address": address, "token": id})
		return err
	}

	asset := normalizeTRC20Token(txs[0].TokenInfo)
	tkChan <- asset
	return nil
}

func normalizeTRC20Token(info trc20Info) blockatlas.Token {
	return blockatlas.Token{
		Name:     info.Name,
		Symbol:   info.Symbol,
		TokenID:  info.Address,
		Coin:     coin.TRX,
		Decimals: info.Decimals,
		Type:     blockatlas.TokenTypeTRC20,
	}
}

func normalizeTRC10Token(info AssetInfo) blockatlas.Token {
	return blockatlas.Token{
		Name:     info.Name,
		Symbol:   info.Symbol,
		TokenID:  info.ID,
		Coin:     coin.TRX,
		Decimals: info.Decimals,
		Type:     blockatlas.TokenTypeTRC10,
	}
}


