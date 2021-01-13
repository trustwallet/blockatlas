package nano

import (
	"strconv"

	"github.com/trustwallet/golibs/client"
	"github.com/trustwallet/golibs/types"
)

type Client struct {
	client.Request
}

func (c *Client) GetAccountHistory(address string) (history AccountHistory, err error) {
	count := strconv.Itoa(types.TxPerPage)
	err = c.Post(&history, "", AccountHistoryRequest{Action: "account_history", Account: address, Count: count})
	return
}
