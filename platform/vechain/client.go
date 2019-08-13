package vechain

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/client"
	"io/ioutil"
	"net/http"
	"net/url"
)

//Client model contains client instance and base url
type Client struct {
	HTTPClient *http.Client
	URL        string
}

// GetCurrentBlockInfo get request function which returns current  blockchain status model
func (c *Client) GetCurrentBlockInfo() (cbi *CurrentBlockInfo, err error) {
	err = client.Request(c.HTTPClient, c.URL, "clientInit", url.Values{}, &cbi)

	return cbi, err
}

// GetBlockByNumber get request function which returns block model requested by number
func (c *Client) GetBlockByNumber(num int64) (block *Block, err error) {
	path := fmt.Sprintf("blocks/%d", num)

	err = client.Request(c.HTTPClient, c.URL, path, url.Values{}, &block)

	return block, err
}

// GetTransactions get request function which returns a VET transfer transactions for given address
func (c *Client) GetTransactions(address string) (TransferTx, error) {
	var transfers TransferTx

	path := fmt.Sprintf("%s/transactions?address=%s&count=25&offset=0", c.URL, address)
	resp, err := c.HTTPClient.Get(path)
	if err != nil {
		logrus.WithError(err).Error("VeChain: Failed HTTP get transactions")
		return transfers, err
	}
	defer resp.Body.Close()

	body, errBody := ioutil.ReadAll(resp.Body)
	if errBody != nil {
		logrus.WithError(err).Error("VeChain: Error decode transaction response body")
		return transfers, err
	}

	errUnm := json.Unmarshal(body, &transfers)
	if errUnm != nil {
		logrus.WithError(err).Error("VeChain: Error Unmarshal transaction response body")
		return transfers, err
	}

	return transfers, nil
}

// GetTokenTransfers get request function which returns a token transfer transactions for given address
func (c *Client) GetTokenTransfers(address string) (TokenTransferTxs, error) {
	var transfers TokenTransferTxs

	url := fmt.Sprintf("%s/tokenTransfers?address=%s&count=25&offset=0", c.URL, address)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		logrus.WithError(err).Error("VeChain: Failed HTTP get token transfer transactions")
		return transfers, err
	}
	defer resp.Body.Close()

	body, errBody := ioutil.ReadAll(resp.Body)
	if errBody != nil {
		logrus.WithError(err).Error("VeChain: Error decode token transfer transaction response body")
		return transfers, err
	}

	errUnm := json.Unmarshal(body, &transfers)
	if errUnm != nil {
		logrus.WithError(err).Error("VeChain: Error Unmarshal token transfer transaction response body")
		return transfers, err
	}

	return transfers, nil
}

// GetTransactionReceipt get request function which returns a transaction for given id and parses it to TransferReceipt
func (c *Client) GetTransactionReceipt(id string) (receipt *TransferReceipt, err error) {
	path := fmt.Sprintf("transactions/%s", id)

	err = client.Request(c.HTTPClient, c.URL, path, url.Values{}, &receipt)

	return receipt, err
}

// GetTransactionByID get request function which returns a transaction for given id and parses it to NativeTransaction
func (c *Client) GetTransactionByID(id string) (transaction *NativeTransaction, err error) {
	path := fmt.Sprintf("transactions/%s", id)

	err = client.Request(c.HTTPClient, c.URL, path, url.Values{}, &transaction)

	return transaction, err
}
