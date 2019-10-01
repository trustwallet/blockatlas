package aeternity

import (
	"encoding/base64"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"strings"

	"github.com/spf13/viper"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client = Client{blockatlas.InitClient(viper.GetString("aeternity.api"))}
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.AE]
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	addressTxs, err := p.client.GetTxs(address, blockatlas.TxPerPage)
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
	txValue := srcTx.TxValue
	return blockatlas.Tx{
		ID:       srcTx.Hash,
		Coin:     coin.AE,
		From:     txValue.Sender,
		To:       txValue.Recipient,
		Fee:      blockatlas.Amount(txValue.Fee),
		Date:     int64(srcTx.Timestamp) / 1000,
		Block:    srcTx.BlockHeight,
		Memo:     getPayload(txValue.Payload),
		Status:   blockatlas.StatusCompleted,
		Sequence: txValue.Nonce,
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(txValue.Amount),
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
