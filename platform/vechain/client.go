package vechain

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type Client struct {
	HTTPClient *http.Client
	URL        string
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
		logrus.WithError(err).Error("VeChain: Failed HTTP get token trasnfer transactions")
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

func (c *Client) GetTransactionReceipt(id string) (*TransferReceipt, error) {
	url := fmt.Sprintf("%s/transactions/%s", c.URL, id)
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
