package aeternity

import (
	"encoding/base64"
	"strings"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/numbers"
)

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	addressTxs, err := p.client.GetTxs(address, blockatlas.TxPerPage)
	if err != nil {
		return nil, err
	}

	var txs []blockatlas.Tx
	for _, srcTx := range addressTxs {
		tx, err := NormalizeTx(&srcTx)
		if err != nil {
			continue
		}
		txs = append(txs, tx)
	}
	return txs, nil
}

func NormalizeTx(srcTx *Transaction) (blockatlas.Tx, error) {
	txValue := srcTx.TxValue
	decimals := coin.Coins[coin.AE].Decimals
	amountFloat, err := txValue.Amount.Float64()
	if err != nil {
		return blockatlas.Tx{}, err
	}
	amount := numbers.Float64toString(amountFloat)
	return blockatlas.Tx{
		ID:       srcTx.Hash,
		Coin:     coin.AE,
		From:     txValue.Sender,
		To:       txValue.Recipient,
		Fee:      blockatlas.Amount(txValue.Fee),
		Date:     srcTx.Timestamp / 1000,
		Block:    srcTx.BlockHeight,
		Memo:     getPayload(txValue.Payload),
		Status:   blockatlas.StatusCompleted,
		Sequence: txValue.Nonce,
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(amount),
			Symbol:   coin.Coins[coin.AE].Symbol,
			Decimals: decimals,
		},
	}, nil
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
