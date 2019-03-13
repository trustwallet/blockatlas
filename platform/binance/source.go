package binance

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

// TODO Headers, rate limiting
var source = http.DefaultClient

var ErrSourceConn = errors.New("connection to servers failed")
var ErrNotFound = errors.New("not found")

type SourceError struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

func (s *SourceError) Error() string {
	return fmt.Sprintf("%d: %s", s.Code, s.Message)
}

type SourceTx struct {
	BlockHeight   uint64 `json:"blockHeight"`
	Code          int    `json:"code"`
	ConfirmBlocks int    `json:"confirmBlocks"`
	Data          string `json:"data"`
	FromAddr      string `json:"fromAddr"`
	OrderId       string `json:"orderId"`
	Timestamp     string `json:"timeStamp"`
	ToAddr        string `json:"toAddr"`
	Age           int64  `json:"txAge"`
	Asset         string `json:"txAsset"`
	Fee           uint64 `json:"txFee"`
	Hash          string `json:"txHash"`
	Value         uint64 `json:"value"`
}

func sourceGetTx(rpcUrl string, txHash string) (*SourceTx, error) {
	uri := fmt.Sprintf("%s/tx?%s",
		rpcUrl,
		url.Values{"hash": {txHash}}.Encode())
	res, err := source.Get(uri)
	if err != nil {
		logrus.WithError(err).Error("Binance: Failed to get transaction")
		return nil, ErrSourceConn
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, ErrNotFound
	} else if res.StatusCode != http.StatusOK {
		var sErr SourceError
		err = json.NewDecoder(res.Body).Decode(&sErr)
		if err != nil {
			logrus.WithError(err).Error("Binance: Failed to get transaction")
			return nil, ErrSourceConn
		} else {
			return nil, &sErr
		}
	}

	stx := new(SourceTx)
	err = json.NewDecoder(res.Body).Decode(stx)
	return stx, nil
}
