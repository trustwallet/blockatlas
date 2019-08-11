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

type Client struct {
	HTTPClient *http.Client
	URL        string
}

func (c *Client) GetCurrentBlockInfo() (cbi *CurrentBlockInfo, err error) {
	err = client.Request(c.HTTPClient, c.URL, "clientInit", url.Values{}, &cbi)

	return cbi, err
}

func (c *Client) GetBlockByNumber(num int64) (block *Block, err error) {
	path := fmt.Sprintf("blocks/%d", num)

	err = client.Request(c.HTTPClient, c.URL, path, url.Values{}, &block)

	return block, err
}

func (c *Client) GetTransactions(address string) (TransferTx, error) {
	var transfers TransferTx

	url := fmt.Sprintf("%s/transactions?address=%s&count=25&offset=0", c.URL, address)
	resp, err := c.HTTPClient.Get(url)
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

func (c *Client) GetTransactionReceipt(id string) (receipt *TransferReceipt, err error) {
	path := fmt.Sprintf("transactions/%s", id)

	err = client.Request(c.HTTPClient, c.URL, path, url.Values{}, &receipt)

	return receipt, err
}

func (c *Client) GetTransactionById(id string) (transaction *NativeTransaction, err error) {
	path := fmt.Sprintf("transactions/%s", id)

	err = client.Request(c.HTTPClient, c.URL, path, url.Values{}, &transaction)

	return transaction, err
}
