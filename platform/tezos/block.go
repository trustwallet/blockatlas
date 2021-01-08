package tezos

import (
	"encoding/json"

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
				if tx, err := mapTransaction(content); err == nil && tx.Kind == TxTypeTransaction {
					if normalized, err := NormalizeRpcTransaction(tx, block.Header); err == nil {
						normalized.ID = op.Hash
						txs = append(txs, normalized)
					}
				}
			}
		}
	}

	return &types.Block{Number: block.Header.Level, Txs: txs}, nil
}

func NormalizeRpcTransaction(tx RpcTransaction, header RpcBlockHeader) (types.Tx, error) {
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

	coin := coin.Tezos()
	return types.Tx{
		Coin:   coin.ID,
		From:   tx.Source,
		To:     tx.Destination,
		Fee:    types.Amount(tx.Fee),
		Date:   date.Unix(),
		Block:  uint64(header.Level),
		Status: status,
		Type:   types.TxTransfer,
		Meta: types.Transfer{
			Value:    types.Amount(tx.Amount),
			Symbol:   coin.Symbol,
			Decimals: coin.Decimals,
		},
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
