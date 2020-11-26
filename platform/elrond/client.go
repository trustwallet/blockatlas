package elrond

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Client struct {
	blockatlas.Request
}

func (c *Client) CurrentBlockNumber() (num int64, err error) {
	var networkStatus NetworkStatus
	path := fmt.Sprintf("network/status/%s", metachainID)
	err = c.getResponse(&networkStatus, path, nil)
	if err != nil {
		return 0, err
	}

	latestNonce := networkStatus.Status.Nonce

	return int64(latestNonce), nil
}

func (c *Client) GetBlockByNumber(height int64) (*blockatlas.Block, error) {
	var blockRes BlockResponse

	path := fmt.Sprintf("hyperblock/by-nonce/%d", uint64(height))
	err := c.getResponse(&blockRes, path, nil)
	if err != nil {
		return nil, err
	}

	block := blockRes.Block
	txs := NormalizeTxs(block.Transactions, "")

	return &blockatlas.Block{
		Number: int64(block.Nonce),
		ID:     block.Hash,
		Txs:    txs,
	}, nil
}

func (c *Client) GetTxsOfAddress(address string) (blockatlas.TxPage, error) {
	var txPage TransactionsPage
	// TODO: enable pagination of Elrond transactions in the future.
	// TODO: currently Elrond only fetches the most recent 20 transactions.
	path := fmt.Sprintf("address/%s/transactions", address)
	err := c.getResponse(&txPage, path, nil)
	if err != nil {
		return nil, err
	}

	txs := NormalizeTxs(txPage.Transactions, address)

	return txs, nil
}

func (c *Client) getResponse(result interface{}, path string, query url.Values) error {
	var genericResponse GenericResponse
	if err := c.Get(&genericResponse, path, query); err != nil {
		return err
	}

	if genericResponse.Code != "successful" {
		return fmt.Errorf("%s", genericResponse.Error)
	}

	return json.Unmarshal(genericResponse.Data, &result)
}
