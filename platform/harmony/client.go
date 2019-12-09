package harmony

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/util"
	"strconv"
)

type Client struct {
	blockatlas.Request
}

func (c *Client) GetTxsOfAddress(address string) (txPage *TxResult, err error) {
	params := []interface{}{
		map[string] interface{}{"address": address, "fullTx": true},
	}
	err = c.RpcCall(&txPage, "hmy_getTransactionsHistory", params)
	return
}

func (c *Client) CurrentBlockNumber() (int64, error) {
	var nodeInfo string
	err := c.RpcCall(&nodeInfo, "hmy_blockNumber", nil)
	if err != nil {
		return 0, err
	}
	decimalBlock, _ := hexToInt(nodeInfo)
	return int64(decimalBlock), nil
}

func hexToInt(hex string) (uint64, error) {
	nonceStr, err := util.HexToDecimal(hex)
	if err != nil {
		return 0, err
	}
	return strconv.ParseUint(nonceStr, 10, 64)
}
