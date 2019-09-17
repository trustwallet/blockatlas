package stellar

import (
	"encoding/json"
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"net/http"
	"net/url"
)

type Client struct {
	HTTP *http.Client
	API  string
}

func (c *Client) GetTxsOfAddress(address string) (txs []Payment, err error) {
	path := fmt.Sprintf("%s/accounts/%s/payments?order=desc&limit=25",
		c.API, url.PathEscape(address))
	return c.getTxs(path)
}

func (c *Client) getTxs(path string) (txs []Payment, err error) {
	res, err := http.Get(path)
	if err != nil {
		return nil, errors.E(err, errors.TypePlatformRequest, errors.Params{"url": path, "platform": "stellar"})
	}
	defer res.Body.Close()

	var payments PaymentsPage
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&payments)
	if err != nil {
		return nil, errors.E(err, errors.TypePlatformUnmarshal, errors.Params{"url": path, "platform": "stellar"})
	}

	return payments.Embedded.Records, nil
}

func (c *Client) CurrentBlockNumber() (num int64, err error) {
	path := fmt.Sprintf("%s/ledgers?order=desc&limit=1", c.API)
	res, err := http.Get(path)
	if err != nil {
		return num, errors.E(err, errors.TypePlatformRequest, errors.Params{"url": path, "platform": "stellar"})
	}
	defer res.Body.Close()
	var ledgers LedgersPage
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&ledgers)
	if err != nil {
		return num, errors.E(err, errors.TypePlatformUnmarshal, errors.Params{"url": path, "platform": "stellar"})
	}

	return ledgers.Embedded.Records[0].Sequence, nil
}

func (c *Client) GetBlockByNumber(num int64) (block *Block, err error) {
	ledger, err := c.getLedger(num)
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("%s/ledgers/%d/payments?limit=100&order=desc", c.API, num)
	payments, err := c.getTxs(path)
	if err != nil {
		return nil, err
	}

	return &Block{Ledger: *ledger, Payments: payments}, nil
}

func (c *Client) getLedger(num int64) (ledger *Ledger, err error) {
	path := fmt.Sprintf("%s/ledgers/%d", c.API, num)
	res, err := http.Get(path)
	if err != nil {
		return nil, errors.E(err, errors.TypePlatformRequest, errors.Params{"url": path, "platform": "stellar"})
	}
	defer res.Body.Close()
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&ledger)
	if err != nil {
		return nil, errors.E(err, errors.TypePlatformUnmarshal, errors.Params{"url": path, "platform": "stellar"})
	}
	return ledger, nil
}
