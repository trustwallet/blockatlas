package tezos

import (
	"encoding/json"
	"fmt"

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

	return ProcessRpcBlock(block, &p.rpcClient)
}

func ProcessRpcBlock(block RpcBlock, rpcClient IRpcClient) (*types.Block, error) {
	txs := types.Txs{}

	for _, ops := range block.Operations {
		for _, op := range ops {
			for _, content := range op.Contents {
				if tx, err := mapTransaction(content); err == nil {
					if !(tx.Kind == TxTypeDelegation || tx.Kind == TxTypeTransaction) {
						continue
					}

					balance := tx.Amount
					if len(balance) == 0 {
						account, err := rpcClient.GetAccountBalanceAtBlock(tx.Source, block.Header.Level)
						if err != nil {
							return nil, err
						}
						balance = account.Balance
					}

					normalized, err := NormalizeRpcTransaction(tx, block.Header, balance)
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

func NormalizeRpcTransaction(tx RpcTransaction, header RpcBlockHeader, balance string) (types.Tx, error) {

	var to string
	var txType types.TransactionType
	switch tx.Kind {
	case TxTypeTransaction:
		to = tx.Destination
		txType = types.TxTransfer
	case TxTypeDelegation:
		to = tx.Delegate
		txType = types.TxAnyAction
	default:
		return types.Tx{}, fmt.Errorf("not supported operation kind: %s", tx.Kind)
	}

	date, err := timefmt.Parse(header.Timestamp, "%Y-%m-%dT%H:%M:%SZ")
	if err != nil {
		return types.Tx{}, err
	}

	metadata, err := mapMetadat(tx, balance)
	if err != nil {
		return types.Tx{}, err
	}

	status := types.StatusCompleted
	if tx.Metadata.OperationResult.Status != TxStatusApplied {
		status = types.StatusError
	}

	return types.Tx{
		Coin:   coin.Tezos().ID,
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

func mapMetadat(tx RpcTransaction, balance string) (interface{}, error) {
	coin := coin.Tezos()

	switch tx.Kind {
	case TxTypeTransaction:
		return types.Transfer{
			Value:    types.Amount(balance),
			Symbol:   coin.Symbol,
			Decimals: coin.Decimals,
		}, nil
	case TxTypeDelegation:
		var title types.KeyTitle
		if len(tx.Delegate) == 0 {
			title = types.AnyActionUndelegation
		} else {
			title = types.AnyActionDelegation
		}

		return types.AnyAction{
			Coin:     coin.ID,
			Title:    title,
			Key:      types.KeyStakeDelegate,
			Name:     coin.Name,
			Symbol:   coin.Symbol,
			Decimals: coin.Decimals,
			Value:    types.Amount(balance),
		}, nil
	default:
		return nil, fmt.Errorf("not supported metadata for kind: %s", tx.Kind)
	}
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
