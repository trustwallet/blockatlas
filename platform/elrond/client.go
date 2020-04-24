package elrond

import (
	"fmt"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Client struct {
	blockatlas.Request
}

func (c *Client) GetCurrentBlock() (num int64, err error) {
	var latestNonce LatestNonce
	err = c.Get(&latestNonce, "block/latest-nonce", nil)
	if err != nil {
		return num, err
	}

	return int64(latestNonce.Nonce), nil
}

func (c *Client) GetBlockByNumber(height int64) (*blockatlas.Block, error) {
	var block Block

	path := fmt.Sprintf("block/meta/%d", uint64(height))
	err := c.Get(&block, path, nil)
	if err != nil {
		return nil, err
	}

	return &blockatlas.Block{
		Number: int64(block.Nonce),
		ID:     block.Hash,
		Txs:    NormalizeTxs(block.Transactions, ""),
	}, nil
}

func (c *Client) GetTxsOfAddress(address string) (blockatlas.TxPage, error) {
	var bulkTxs BulkTransactions
	path := fmt.Sprintf("address/%s/transactions", address)
	err := c.Get(&bulkTxs, path, nil)
	if err != nil {
		return nil, err
	}

	return NormalizeTxs(bulkTxs.Transactions, address), nil
}
