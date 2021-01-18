package vechain

import (
	"fmt"
	"strings"

	"github.com/trustwallet/golibs/client"
)

type Client struct {
	client.Request
}

func (c *Client) GetCurrentBlock() (int64, error) {
	var b Block
	err := c.Get(&b, "blocks/best", nil)
	return b.Number, err
}

func (c *Client) GetBlockByNumber(num int64) (block Block, err error) {
	path := fmt.Sprintf("blocks/%d", num)
	err = c.Get(&block, path, nil)
	return
}

func (c *Client) GetTransactions(address string, block int64) (txs []LogTransfer, err error) {
	err = c.Post(&txs, "logs/transfer", LogRequest{
		Options: Options{Offset: 0, Limit: 15},
		CriteriaSet: []CriteriaSet{
			{Sender: address},
			{Recipient: address},
		},
		Range: Range{Unit: rangeUnit, From: 0, To: block},
		Order: "desc",
	})
	return
}

// Get events related to address and VIP180 contract address
func (c *Client) GetLogsEvent(address, token string, block int64) (txs []LogEvent, err error) {
	tokenHex := getFilter(address)
	err = c.Post(&txs, "logs/event", LogRequest{
		Options: Options{Offset: 0, Limit: 10},
		CriteriaSet: []CriteriaSet{
			{Address: token, Topic1: tokenHex},
			{Address: token, Topic2: tokenHex},
		},
		Range: Range{Unit: rangeUnit, From: 0, To: block},
		Order: "desc",
	})
	return
}

func (c *Client) GetTransactionByID(id string) (transaction Tx, err error) {
	path := fmt.Sprintf("transactions/%s", id)
	err = c.Get(&transaction, path, nil)
	return
}

func (c *Client) GetTransactionReceiptByID(id string) (transaction TxReceipt, err error) {
	path := fmt.Sprintf("transactions/%s/receipt", id)
	err = c.Get(&transaction, path, nil)
	return
}

func (c *Client) GetAccount(address string) (account Account, err error) {
	err = c.Get(&account, "accounts/"+address, nil)
	return
}

// Creates hex based on address as required for topic criteria
// 0xB5e883349e68aB59307d1604555AC890fAC47128 => 0x000000000000000000000000B5e883349e68aB59307d1604555AC890fAC47128
func getFilter(hex string) string {
	hexStr := strings.TrimPrefix(hex, "0x")
	return fmt.Sprintf(filterPrefix, hexStr)
}
