package nimiq

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"strconv"
)

type Client struct {
	blockatlas.Request
}

func (c *Client) GetTxsOfAddress(address string, count int) (tx []Tx, err error) {
	err = c.RpcCall(&tx, "getTransactionsByAddress", []string{address, strconv.Itoa(count)})
	return
}

func (c *Client) CurrentBlockNumber() (num int64, err error) {
	err = c.RpcCall(&num, "blockNumber", []string{})
	return
}

func (c *Client) GetBlockByNumber(num int64) (b *Block, err error) {
	n := strconv.Itoa(int(num))
	err = c.RpcCall(&b, "getBlockByNumber", []string{n, "true"})
	return
}
