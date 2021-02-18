package solana

import (
	"github.com/trustwallet/golibs/client"
)

type Client struct {
	client.Request
}

func (c *Client) GetLasteBlock() (int64, error) {
	var epoch EpochInfo
	err := c.RpcCall(&epoch, "getEpochInfo", []string{})
	if err != nil {
		return 0, err
	}
	return int64(epoch.BlockHeight), nil
}

func (c *Client) GetEpochInfo() (epochInfo EpochInfo, err error) {
	err = c.RpcCall(&epochInfo, "getEpochInfo", []string{})
	return
}

func (c *Client) GetTransactionsByAddress(address string) ([]ConfirmedTransaction, error) {
	var signatures []ConfirmedSignature
	params := []interface{}{
		address,
		map[string]interface{}{"limit": 25},
	}
	err := c.RpcCall(&signatures, "getConfirmedSignaturesForAddress2", params)
	if err != nil {
		return nil, err
	}

	return c.GetTransactionSignatures(signatures)
}

func (c *Client) GetTransactionSignatures(signatures []ConfirmedSignature) ([]ConfirmedTransaction, error) {
	var txs []ConfirmedTransaction

	// check empty
	if len(signatures) == 0 {
		return txs, nil
	}

	// build batch request
	requests := make(client.RpcRequests, 0)
	for _, sig := range signatures {
		requests = append(requests, &client.RpcRequest{
			Method: "getConfirmedTransaction",
			Params: []string{
				sig.Signature,
				"jsonParsed",
			},
		})
	}

	responses, err := c.RpcBatchCall(requests)
	if err != nil {
		return txs, err
	}

	// convert to ConfirmedTransaction
	for _, response := range responses {
		var tx ConfirmedTransaction
		if err := response.GetObject(&tx); err == nil {
			txs = append(txs, tx)
		}
	}
	return txs, nil
}

func (c *Client) GetTransactionsInBlock(num int64) (block Block, err error) {
	err = c.RpcCall(&block, "getConfirmedBlock", []interface{}{num, "jsonParsed"})
	return
}
