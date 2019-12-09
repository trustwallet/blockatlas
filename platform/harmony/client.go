package harmony

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/util"
	"strconv"
)

type Client struct {
	blockatlas.Request
}

func (c *Client) GetTxsOfAddress(address string) (txPage *TxResult, err error) {
	type params struct {
		address string `json:"address"`
		fullTx  bool   `json:"fullTx"`
	}
	err = c.RpcCall(&txPage, "hmy_getTransactionsHistory", []params{{address: address, fullTx: true}})
	return
}

func (c *Client) CurrentBlockNumber() (int64, error) {
	var nodeInfo string
	err := c.RpcCall(&nodeInfo, "hmy_blockNumber", nil)
	decimalBlock, err := hexToInt(nodeInfo)
	if err != nil {
		return 0, err
	}
	return int64(decimalBlock), nil
}

func (c *Client) GetBlockByNumber(num int64) (info BlockInfo, err error) {
	n := fmt.Sprintf("0x%x", num)
	err = c.RpcCall(&info, "hmy_getBlockByNumber", []interface{}{n, true})
	return
}

func hexToInt(hex string) (uint64, error) {
	nonceStr, err := util.HexToDecimal(hex)
	if err != nil {
		return 0, err
	}
	return strconv.ParseUint(nonceStr, 10, 64)
}
