package aeternity

import (
	"encoding/base64"
	"github.com/trustwallet/blockatlas/coin"
	"net/http"
	"strings"

	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client.URL = viper.GetString("aeternity.api")
	p.client.HTTPClient = http.DefaultClient
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.AE]
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	addressTxs, err := p.client.GetTxs(address, 25)
	if err != nil {
		return nil, err
	}

	var txs []blockatlas.Tx
	for _, srcTx := range addressTxs {
		txs = append(txs, NormalizeTx(&srcTx))
	}
	return txs, nil
}

func NormalizeTx(srcTx *Transaction) blockatlas.Tx {
	return blockatlas.Tx{
		ID:     srcTx.Hash,
		Coin:   coin.AE,
		From:   srcTx.TxValue.Sender,
		To:     srcTx.TxValue.Recipient,
		Fee:    blockatlas.Amount(srcTx.TxValue.Fee),
		Date:   int64(srcTx.Timestamp) / 1000,
		Block:  srcTx.BlockHeight,
		Memo:   getPayload(srcTx.TxValue.Payload),
		Status: blockatlas.StatusCompleted,
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(srcTx.TxValue.Amount),
			Symbol:   coin.Coins[coin.AE].Symbol,
			Decimals: coin.Coins[coin.AE].Decimals,
		},
	}
}

func getPayload(encodedPayload string) string {
	payload := []byte(strings.Replace(encodedPayload, "ba_", "", 1))
	if len(payload) <= 8 {
		return ""
	} else {
		payload = payload[:len(payload)-8]
		data, err := base64.StdEncoding.DecodeString(string(payload))
		if err != nil || len(data) <= 4 {
			payload = []byte("")
		} else {
			payload = data
		}
		return string(payload)
	}
}
