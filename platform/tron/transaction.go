package tron

import (
	"errors"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/types"
)

func (p *Platform) GetTxsByAddress(address string) (types.Txs, error) {
	Txs, err := p.gridClient.fetchTxsOfAddress(address, "")
	if err != nil && len(Txs) == 0 {
		return nil, err
	}

	txs := make(types.Txs, 0)
	for _, srcTx := range Txs {
		tx, err := normalize(srcTx)
		if err != nil {
			continue
		}

		if len(srcTx.Data.Contracts) > 0 && srcTx.Data.Contracts[0].Type == TransferContract {
			txs = append(txs, *tx)
		} else {
			continue
		}
	}

	return txs, nil
}

func (p *Platform) GetTokenTxsByAddress(address, token string) (types.Txs, error) {
	unknownTokenType := errors.New("unknownTokenType")
	tokenType := getTokenType(token)

	switch tokenType {
	case types.TRC10:
		txs, err := p.fetchTransactionsForTRC10Tokens(address, token)
		if err != nil {
			return nil, err
		}
		return txs, nil
	case types.TRC20:
		trc20Transactions, err := p.client.fetchTRC20Transactions(address)
		if err != nil {
			return nil, err
		}
		return normalizeTRC20Transactions(trc20Transactions), nil
	default:
		return nil, unknownTokenType
	}
}

func getTokenType(token string) types.TokenType {
	_, err := strconv.Atoi(token)
	if err != nil {
		return types.TRC20
	} else {
		return types.TRC10
	}
}

func addTokenMeta(tx *types.Tx, srcTx Tx, tokenInfo AssetInfo) {
	transfer := srcTx.Data.Contracts[0].Parameter.Value
	tx.Meta = types.TokenTransfer{
		Name:     tokenInfo.Name,
		Symbol:   strings.ToUpper(tokenInfo.Symbol),
		TokenID:  strconv.Itoa(int(tokenInfo.ID)),
		Decimals: tokenInfo.Decimals,
		Value:    transfer.Amount,
		From:     tx.From,
		To:       tx.To,
	}
}

func (p *Platform) fetchTransactionsForTRC10Tokens(address, token string) (types.Txs, error) {
	txs := make(types.Txs, 0)

	tokenTxs, err := p.gridClient.fetchTxsOfAddress(address, token)
	if err != nil {
		return nil, err
	}

	info, err := p.gridClient.fetchTokenInfo(token)
	if err != nil {
		return nil, err
	}
	for _, srcTx := range tokenTxs {
		tx, err := normalize(srcTx)
		if err != nil {
			log.Error(err)
			continue
		}
		if info.Data != nil && len(info.Data) > 0 {
			addTokenMeta(tx, srcTx, info.Data[0])
		}

		txs = append(txs, *tx)
	}
	return txs, nil
}

func normalizeTRC20Transactions(transactions TRC20Transactions) types.Txs {
	txs := make(types.Txs, 0, len(transactions.Data))
	for _, rawTx := range transactions.Data {
		tx := types.Tx{
			ID:     rawTx.TransactionID,
			Coin:   coin.TRON,
			Date:   rawTx.BlockTimestamp / 1000,
			From:   rawTx.From,
			To:     rawTx.To,
			Fee:    "0",
			Block:  0,
			Status: types.StatusCompleted,
			Meta: types.TokenTransfer{
				Name:     rawTx.TokenInfo.Name,
				Symbol:   rawTx.TokenInfo.Symbol,
				TokenID:  rawTx.TokenInfo.Address,
				Decimals: uint(rawTx.TokenInfo.Decimals),
				Value:    types.Amount(rawTx.Value),
				From:     rawTx.From,
				To:       rawTx.To,
			},
		}
		txs = append(txs, tx)
	}
	return txs
}

func normalize(srcTx Tx) (*types.Tx, error) {
	if len(srcTx.Data.Contracts) == 0 {
		return nil, errors.New("no contracts")
	}

	contract := srcTx.Data.Contracts[0]
	if contract.Type != TransferContract && contract.Type != TransferAssetContract {
		return nil, errors.New("TRON: invalid contract transfer")
	}

	transfer := contract.Parameter.Value
	from, err := HexToAddress(transfer.OwnerAddress)
	if err != nil {
		return nil, err
	}
	to, err := HexToAddress(transfer.ToAddress)
	if err != nil {
		return nil, err
	}

	return &types.Tx{
		ID:     srcTx.ID,
		Coin:   coin.TRON,
		Date:   srcTx.BlockTime / 1000,
		From:   from,
		To:     to,
		Fee:    "0",
		Block:  0,
		Status: types.StatusCompleted,
		Meta: types.Transfer{
			Value:    transfer.Amount,
			Symbol:   coin.Tron().Symbol,
			Decimals: coin.Tron().Decimals,
		},
	}, nil
}
