package elrond

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/trustwallet/golibs/client"
	"github.com/trustwallet/golibs/types"
)

type Client struct {
	client.Request
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

func (c *Client) GetBlockByNumber(height int64) (*types.Block, error) {
	var blockRes BlockResponse

	path := fmt.Sprintf("hyperblock/by-nonce/%d", uint64(height))
	err := c.getResponse(&blockRes, path, nil)
	if err != nil {
		return nil, err
	}

	block := blockRes.Block
	txs := NormalizeTxs(block.Transactions, "", blockRes.Block)

	return &types.Block{
		Number: int64(block.Nonce),
		Txs:    txs,
	}, nil
}

func (c *Client) GetTxsOfAddress(address string) (types.Txs, error) {
	var txPage TransactionsPage
	// TODO: enable pagination of Elrond transactions in the future.
	// TODO: currently Elrond only fetches the most recent 20 transactions.
	path := fmt.Sprintf("address/%s/transactions", address)
	err := c.getResponse(&txPage, path, nil)
	if err != nil {
		return nil, err
	}

	txs := NormalizeTxs(txPage.Transactions, address, Block{})

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
