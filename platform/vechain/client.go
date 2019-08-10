package vechain

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas"
	"io/ioutil"
	"net/http"
)

type Client struct {
	HTTPClient *http.Client
	URL        string
}

func (c *Client) GetCurrentBlockInfo() (*CurrentBlockInfo, error) {
	uri := fmt.Sprintf("%s/clientInit", c.URL)

	resp, err := c.HTTPClient.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := getHTTPError(resp, "GetCurrentBlockInfo"); err != nil {
		return nil, err
	}

	cbi := new(CurrentBlockInfo)
	err = json.NewDecoder(resp.Body).Decode(cbi)
	return cbi, nil
}

func (c *Client) GetBlockByNumber(num int64) (*Block, error) {
	uri := fmt.Sprintf("%s/blocks/%d", c.URL, num)

	resp, err := c.HTTPClient.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := getHTTPError(resp, "GetBlockByNumber"); err != nil {
		return nil, err
	}

	block := new(Block)
	err = json.NewDecoder(resp.Body).Decode(block)
	return block, nil
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

func (c *Client) GetTransactionById(id string) (*NativeTransaction, error) {
	url := fmt.Sprintf("%s/transactions/%s", c.URL, id)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var receipt NativeTransaction
	err = json.NewDecoder(resp.Body).Decode(&receipt)
	if err != nil {
		return nil, err
	}

	return &receipt, nil
}

func getHTTPError(res *http.Response, desc string) error {
	switch res.StatusCode {
	case http.StatusBadRequest, http.StatusNotFound:
		return getAPIError(res, desc)
	case http.StatusOK:
		return nil
	default:
		return fmt.Errorf("%s", res.Status)
	}
}

func getAPIError(res *http.Response, desc string) error {
	var sErr Error
	err := json.NewDecoder(res.Body).Decode(&sErr)
	if err != nil {
		logrus.WithError(err).Error("VeChain: Failed to decode error response")
		return blockatlas.ErrSourceConn
	}

	switch sErr.Message {
	case "address is not valid":
		return blockatlas.ErrInvalidAddr
	}

	logrus.WithFields(logrus.Fields{
		"status":  res.StatusCode,
		"code":    sErr.Code,
		"message": sErr.Message,
	}).Error("VeChain: Failed to " + desc)
	return blockatlas.ErrSourceConn
}
