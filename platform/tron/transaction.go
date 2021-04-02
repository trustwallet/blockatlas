package tron

import (
	"errors"
	"strconv"

	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/types"
)

func (p *Platform) GetTxsByAddress(address string) (types.Txs, error) {
	Txs, err := p.client.fetchTxsOfAddress(address, "")
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
		return types.Txs{}, nil
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
