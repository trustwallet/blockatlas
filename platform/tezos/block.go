package tezos

import (
	"encoding/json"
	"errors"

	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/types"

	"github.com/itchyny/timefmt-go"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.rpcClient.GetBlockHead()
}

func (p *Platform) GetBlockByNumber(num int64) (*types.Block, error) {
	block, err := p.rpcClient.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}

	return NormalizeRpcBlock(block)
}

func NormalizeRpcBlock(block RpcBlock) (*types.Block, error) {
	txs := []types.Tx{}

	for _, ops := range block.Operations {
		for _, op := range ops {
			for _, content := range op.Contents {
				if tx, err := mapTransaction(content); err == nil {
					if !(tx.Kind == TxTypeDelegation || tx.Kind == TxTypeTransaction) {
						continue
					}
					normalized, err := NormalizeRpcTransaction(tx, block.Header)
					if err != nil {
						return nil, err
					}
					normalized.ID = op.Hash
					txs = append(txs, normalized)
				}
			}
		}
	}

	return &types.Block{Number: block.Header.Level, Txs: txs}, nil
}

func NormalizeRpcTransaction(tx RpcTransaction, header RpcBlockHeader) (types.Tx, error) {
	coin := coin.Tezos()

	var metadata interface{}
	var to string
	var txType types.TransactionType
	switch tx.Kind {
	case TxTypeTransaction:
		to = tx.Destination
		txType = types.TxTransfer
		metadata = types.Transfer{
			Value:    types.Amount(tx.Amount),
			Symbol:   coin.Symbol,
			Decimals: coin.Decimals,
		}
	case TxTypeDelegation:
		var title types.KeyTitle
		if len(tx.Delegate) == 0 {
			title = types.AnyActionUndelegation
		} else {
			title = types.AnyActionDelegation
			to = tx.Delegate
		}
		txType = types.TxAnyAction

		metadata = types.AnyAction{
			Coin:     coin.ID,
			Title:    title,
			Key:      types.KeyStakeDelegate,
			Name:     coin.Name,
			Symbol:   coin.Symbol,
			Decimals: coin.Decimals,
			Value:    "0",
		}
	default:
		return types.Tx{}, errors.New("not supported operation kind: " + tx.Kind)
	}

	date, err := timefmt.Parse(header.Timestamp, "%Y-%m-%dT%H:%M:%SZ")
	if err != nil {
		return types.Tx{}, err
	}

	var status types.Status
	if tx.Metadata.OperationResult.Status == TxStatusApplied {
		status = types.StatusCompleted
	} else {
		status = types.StatusError
	}

	return types.Tx{
		Coin:   coin.ID,
		From:   tx.Source,
		To:     to,
		Fee:    types.Amount(tx.Fee),
		Date:   date.Unix(),
		Block:  uint64(header.Level),
		Status: status,
		Type:   txType,
		Meta:   metadata,
	}, nil
}

func mapTransaction(content interface{}) (RpcTransaction, error) {
	bytes, err := json.Marshal(content)
	if err != nil {
		return RpcTransaction{}, err
	}

	var tx RpcTransaction
	err = json.Unmarshal(bytes, &tx)
	if err != nil {
		return RpcTransaction{}, err
	}

	return tx, nil
}
