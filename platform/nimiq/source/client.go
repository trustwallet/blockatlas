package source

import (
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/models"
	"github.com/ybbus/jsonrpc"
)

type Client struct {
	RpcClient jsonrpc.RPCClient
	RpcUrl    string
}

func NewClient(rpcUrl string) *Client {
	return &Client {
		RpcUrl:    rpcUrl,
		RpcClient: jsonrpc.NewClient(rpcUrl),
	}
}

func (c *Client) GetTxsOfAddress(address string) (txs []Tx, err error) {
	err = c.RpcClient.CallFor(&txs, "getTransactionsByAddress", address, models.TxPerPage)
	if jErr, ok := err.(*jsonrpc.RPCError); ok {
		if jErr.Code == 1 {
			return nil, ErrInvalidAddr
		} else {
			logrus.WithError(err).Error("Nimiq: Failed to get transactions")
			return nil, ErrSourceConn
		}
	} else if err != nil {
		return nil, err
	}
	return
}
