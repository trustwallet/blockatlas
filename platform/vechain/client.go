package vechain

import(
	"github.com/sirupsen/logrus"
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"strings"
)

type Client struct {
	HTTPClient *http.Client
	URL        string
}

func (c *Client) GetAddressTransactions(address string) ([]Tx, error) {
	url := fmt.Sprintf("%s/logs/transfer", c.URL)

	addressEscaped, _ := json.Marshal(address)

	payload := fmt.Sprintf(`{
		"range": {
			"unit": "block",
			"from": 0,
			"to": 9000000
		},
		"options": {
			"offset": 0,
			"limit": 25
		},
		"criteriaSet": [
			{
				"sender": %[1]s
			},
			{
				"recipient": %[1]s
			}
		],
		"order": "desc"
	}`, string(addressEscaped))

	resp, err := c.HTTPClient.Post(url, "application/json", strings.NewReader(payload))
	if err != nil {
		logrus.WithError(err).Error("VeChain: Failed HTTP get transactions")
		return nil, err
	}
	defer resp.Body.Close()

	body, errBody := ioutil.ReadAll(resp.Body)
	var transactions []Tx
	if errBody != nil {
		logrus.WithError(err).Error("VeChain: Error decode transaction response body")
		return nil, err
	}

	errUnm := json.Unmarshal(body, &transactions)
	if errUnm != nil {
		logrus.WithError(err).Error("VeChain: Error Unmarshal transaction response body")
		return nil, err
	}

	return transactions, nil
}

func (c *Client) GetTransactionId(cn chan<- TxId, id string) {
	defer wg.Done()

	url := fmt.Sprintf("%s/transactions/%s", c.URL, id)

	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		logrus.WithError(err).Error("VeChain: Failed HTTP get transaction")
	}
	defer resp.Body.Close()

	var transaction TxId
	err = json.NewDecoder(resp.Body).Decode(&transaction)

	if err != nil {
		logrus.WithError(err).Error("VeChain: Error decode transaction response body")
	} else {
		cn <- transaction
	}
}

func (c *Client) GetTransacionReceipt(cn chan <- TxReceipt, id string) {
	defer wg.Done()

	url := fmt.Sprintf("%s/transactions/%s/receipt", c.URL, id)

	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		logrus.WithError(err).Error("VeChain: Failed HTTP get transaction receipt")
	}
	defer resp.Body.Close()

	var receipt TxReceipt

	err = json.NewDecoder(resp.Body).Decode(&receipt)
	if err != nil {
		logrus.WithError(err).Error("VeChain: Error decode transaction receipt response body")
	} else {
		cn <- receipt
	}
}

