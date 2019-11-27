package tron

import (
	"encoding/hex"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"sync"
)

type Platform struct {
	client Client
}

const Annual = 4.32

func (p *Platform) Init() error {
	p.client = Client{blockatlas.InitClient(viper.GetString("tron.api"))}
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.TRX]
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.CurrentBlockNumber()
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	block, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}

	txsChan := p.NormalizeBlockTxs(block.Txs)
	txs := make(blockatlas.TxPage, 0)
	for cTxs := range txsChan {
		txs = append(txs, cTxs)
	}

	return &blockatlas.Block{
		Number: num,
		Txs:    txs,
	}, nil
}

func (p *Platform) NormalizeBlockTxs(srcTxs []Tx) chan blockatlas.Tx {
	txChan := make(chan blockatlas.Tx, len(srcTxs))
	var wg sync.WaitGroup
	for _, srcTx := range srcTxs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			p.NormalizeBlockChannel(srcTx, txChan)
		}()
	}
	wg.Wait()
	close(txChan)
	return txChan
}

func (p *Platform) NormalizeBlockChannel(srcTx Tx, txChan chan blockatlas.Tx) {
	if len(srcTx.Data.Contracts) == 0 {
		return
	}

	tx, err := Normalize(srcTx)
	if err != nil {
		return
	}
	transfer := srcTx.Data.Contracts[0].Parameter.Value
	if len(transfer.AssetName) > 0 {
		assetName, err := hex.DecodeString(transfer.AssetName[:])
		if err == nil {
			info, err := p.client.GetTokenInfo(string(assetName))
			if err == nil && len(info.Data) > 0 {
				setTokenMeta(tx, srcTx, info.Data[0])
			}
		}
	}
	txChan <- *tx
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	Txs, err := p.client.GetTxsOfAddress(address, "")
	if err != nil && len(Txs) == 0 {
		return nil, err
	}

	txs := make(blockatlas.TxPage, 0)
	for _, srcTx := range Txs {
		tx, err := Normalize(srcTx)
		if err != nil {
			continue
		}
		txs = append(txs, *tx)
	}

	return txs, nil
}

func (p *Platform) GetTokenTxsByAddress(address, token string) (blockatlas.TxPage, error) {
	tokenTxs, err := p.client.GetTxsOfAddress(address, token)
	if err != nil {
		return nil, errors.E(err, "TRON: failed to get token from address", errors.TypePlatformApi,
			errors.Params{"address": address, "token": token}).PushToSentry()
	}

	txs := make(blockatlas.TxPage, 0)

	if len(tokenTxs) == 0 {
		return txs, nil
	}

	info, err := p.client.GetTokenInfo(token)
	if err != nil || len(info.Data) == 0 {
		return nil, errors.E(err, "TRON: failed to get token info", errors.TypePlatformApi,
			errors.Params{"address": address, "token": token}).PushToSentry()
	}

	for _, srcTx := range tokenTxs {
		tx, err := Normalize(srcTx)
		if err != nil {
			logger.Error(err)
			continue
		}
		setTokenMeta(tx, srcTx, info.Data[0])
		txs = append(txs, *tx)
	}

	return txs, nil
}

func (p *Platform) GetTokenListByAddress(address string) (blockatlas.TokenPage, error) {
	tokens, err := p.client.GetAccount(address)
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
	return tokenPage, nil
}

func (p *Platform) getTokens(ids []string) chan blockatlas.Token {
	tkChan := make(chan blockatlas.Token, len(ids))
	var wg sync.WaitGroup
	for _, id := range ids {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := p.getTokensChannel(id, tkChan)
			if err != nil {
				logger.Error(err)
			}
		}()
	}
	wg.Wait()
	close(tkChan)
	return tkChan
}

func (p *Platform) getTokensChannel(id string, tkChan chan blockatlas.Token) error {
	info, err := p.client.GetTokenInfo(id)
	if err != nil || len(info.Data) == 0 {
		logger.Error(err, "GetTokenInfo: invalid token")
		return err
	}
	asset := NormalizeToken(info.Data[0])
	tkChan <- asset
	return nil
}

func NormalizeToken(info AssetInfo) blockatlas.Token {
	return blockatlas.Token{
		Name:     info.Name,
		Symbol:   info.Symbol,
		TokenID:  info.ID,
		Coin:     coin.TRX,
		Decimals: info.Decimals,
		Type:     blockatlas.TokenTypeTRC10,
	}
}

func setTokenMeta(tx *blockatlas.Tx, srcTx Tx, tokenInfo AssetInfo) {
	transfer := srcTx.Data.Contracts[0].Parameter.Value
	tx.Meta = blockatlas.TokenTransfer{
		Name:     tokenInfo.Name,
		Symbol:   tokenInfo.Symbol,
		TokenID:  tokenInfo.ID,
		Decimals: tokenInfo.Decimals,
		Value:    transfer.Amount,
		From:     tx.From,
		To:       tx.To,
	}
}

/// Normalize converts a Tron transaction into the generic model
func Normalize(srcTx Tx) (*blockatlas.Tx, error) {
	if len(srcTx.Data.Contracts) == 0 {
		return nil, errors.E("TRON: transfer without contract", errors.TypePlatformApi,
			errors.Params{"tx": srcTx}).PushToSentry()
	}

	contract := srcTx.Data.Contracts[0]
	if contract.Type != TransferContract && contract.Type != TransferAssetContract {
		return nil, errors.E("TRON: invalid contract transfer", errors.TypePlatformApi,
			errors.Params{"tx": srcTx, "type": contract.Type})
	}

	transfer := contract.Parameter.Value
	from, err := HexToAddress(transfer.OwnerAddress)
	if err != nil {
		return nil, errors.E(err, "TRON: failed to get from address", errors.TypePlatformApi,
			errors.Params{"tx": srcTx}).PushToSentry()
	}
	to, err := HexToAddress(transfer.ToAddress)
	if err != nil {
		return nil, errors.E(err, "TRON: failed to get to address", errors.TypePlatformApi,
			errors.Params{"tx": srcTx}).PushToSentry()
	}

	return &blockatlas.Tx{
		ID:     srcTx.ID,
		Coin:   coin.TRX,
		Date:   srcTx.BlockTime / 1000,
		From:   from,
		To:     to,
		Fee:    "0",
		Block:  0,
		Status: blockatlas.StatusCompleted,
		Meta: blockatlas.Transfer{
			Value:    transfer.Amount,
			Symbol:   coin.Coins[coin.TRX].Symbol,
			Decimals: coin.Coins[coin.TRX].Decimals,
		},
	}, nil
}
