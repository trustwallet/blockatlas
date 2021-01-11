package nimiq

import (
	"strconv"

	"github.com/trustwallet/golibs/client"
	"github.com/trustwallet/golibs/types"
)

type Client struct {
	client.Request
}

func (c *Client) GetTxsOfAddress(address string) (tx []Tx, err error) {
	err = c.RpcCall(&tx, "getTransactionsByAddress", []string{address, strconv.Itoa(types.TxPerPage)})
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
