package bitcoin

import (
	"encoding/json"
	"fmt"
	"github.com/trustwallet/blockatlas/client"
	"net/http"
	"net/url"
)

type Client struct {
	HTTPClient *http.Client
	URL        string
}

func (c *Client) GetTransactions(address string) (TransactionsList, error) {
	var transfers TransactionsList

	path := fmt.Sprintf("address/%s", address)
	err := client.Request(c.HTTPClient, c.URL, path, url.Values{"details": {"txs"}}, &transfers)

	return transfers, err
}

func (c *Client) GetTransactionsByXpub(xpub string) (TransactionsList, error) {
	var transfers TransactionsList

	path := fmt.Sprintf("v2/xpub/%s", xpub)
	err := client.Request(c.HTTPClient, c.URL, path, url.Values{"details": {"txs"}}, &transfers)

	return transfers, err
}

func (c *Client) GetTransactionReceipt(id string) (*TransferReceipt, error) {
	url := fmt.Sprintf("%s/v2/tx/%s", c.URL, id)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var receipt TransferReceipt
	err = json.NewDecoder(resp.Body).Decode(&receipt)
	if err != nil {
		return nil, err
	}

	return &receipt, nil
}
