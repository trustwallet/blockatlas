package waves

import (
	"encoding/json"
	"fmt"
	"github.com/mr-tron/base58"
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas"
	"net/http"
)

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
}

func (c *Client) GetTxs(address string, limit int, after string) ([]Transaction, error) {
	return c.getTxs(fmt.Sprintf("%s/transactions/address/%s/limit/%d?after=%s",
		c.BaseURL,
		address,
		limit,
		after))
}

func (c *Client) getTxs(uri string) ([]Transaction, error) {
	req, _ := http.NewRequest("GET", uri, nil)

	res, err := c.HTTPClient.Do((req))
	if err != nil {
		logrus.WithError(err).Error("Waves: Failed to get transactions")
		return nil, blockatlas.ErrSourceConn
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("http %s", res.Status)
	}

	txsArrays := new([][]Transaction)
	err = json.NewDecoder(res.Body).Decode(txsArrays)
	if err != nil {
		return nil, err
	}
	txsObj := *txsArrays
	txs := txsObj[0]

	var result []Transaction
	for _, tx := range txs {
		// support only transfer transaction
		if tx.Type == 4 {
			if len(tx.AssetId) != 0 {
				tokenInfo, err := c.GetTokenInfo(tx.AssetId)
				if err != nil {
					return nil, err
				}
				tx.Asset = tokenInfo
			}
		}
		attachmentBytes, err := base58.DecodeAlphabet(tx.Attachment, base58.BTCAlphabet)
		if err != nil {
			return nil, err
		}
		tx.Attachment = string(attachmentBytes)
		result = append(result, tx)
	}

	return result, nil
}

func (c *Client) GetTokenInfo(tokenId string) (*TokenInfo, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/assets/details/%s",
		c.BaseURL,
		tokenId), nil)

	res, err := c.HTTPClient.Do((req))
	if err != nil {
		logrus.WithError(err).Error("Waves: Failed to get token info")
		return nil, blockatlas.ErrSourceConn
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("http %s", res.Status)
	}

	tokenInfo := new(TokenInfo)
	err = json.NewDecoder(res.Body).Decode(tokenInfo)
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}
