package harmony

import (
	"fmt"
	"strconv"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/pkg/numbers"
)

type Client struct {
	blockatlas.Request
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

func (c *Client) GetValidators() (validators Validators, err error) {
	err = c.RpcCall(&validators.Validators, "hmy_getAllValidatorInformation", []interface{}{-1})
	if err != nil {
		logger.Error(err, "Harmony: Failed to get all validator addresses")
	}

	return
}

func (c *Client) GetDelegations(address string) (delegations Delegations, err error) {
	err = c.RpcCall(&delegations.List, "hmy_getDelegationsByDelegator", []interface{}{address})
	if err != nil {
		logger.Error(err, "Harmony: Failed to get delegations for address")
	}
	return
}

func (c *Client) GetBalance(address string) (string, error) {
	var result string
	err := c.RpcCall(&result, "hmy_getBalance", []interface{}{address, "latest"})
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
