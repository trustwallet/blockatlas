package harmony

import (
	"fmt"
	"strconv"

	"github.com/trustwallet/golibs/client"
	"github.com/trustwallet/golibs/numbers"
)

type Client struct {
	client.Request
}

func (c *Client) GetTxsOfAddress(address string) (txPage *TxResult, err error) {
	params := []interface{}{
		map[string]interface{}{
			"address": address,
			"fullTx":  true,
		},
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

func (c *Client) GetBalance(address string) (string, error) {
	var result string
	err := rpcCallStub(c, &result, "hmy_getBalance", []interface{}{address, "latest"})

	if err != nil {
		return "0", err
	}
	balance, err := numbers.HexToDecimal(result)
	if err != nil {
		return "0", err
	}
	return balance, nil
}

func hexToInt(hex string) (uint64, error) {
	nonceStr, err := numbers.HexToDecimal(hex)
	if err != nil {
		return 0, err
	}
	return strconv.ParseUint(nonceStr, 10, 64)
}

// rpcCallStub is can be overwritten by the unit test
var rpcCallStub = func(c *Client, result interface{}, method string, params interface{}) error {
	return c.RpcCall(result, method, params)
}
