package nimiq

import (
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/models"
	"github.com/ybbus/jsonrpc"
)

type Client struct {
	BaseURL   string
	rpcClient jsonrpc.RPCClient
}

func (c *Client) Init() {
	c.rpcClient = jsonrpc.NewClient(c.BaseURL)
}

func (c *Client) GetTxsOfAddress(address string) (txs []Tx, err error) {
	var res *jsonrpc.RPCResponse
	res, err = c.rpcClient.CallRaw(&jsonrpc.RPCRequest {
		Method: "getTransactionsByAddress",
		Params: []interface{}{ address, models.TxPerPage },
		ID: 42,
		JSONRPC: "2.0",
	})
	if jErr, ok := err.(*jsonrpc.RPCError); ok {
		if jErr.Code == 1 {
			return nil, models.ErrInvalidAddr
		} else {
			logrus.WithError(err).Error("Nimiq: Failed to get transactions")
			return nil, models.ErrSourceConn
		}
	} else if err != nil {
		return nil, err
	}
	res.GetObject(&txs)
	return
}
