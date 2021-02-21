package aeternity

import (
	"encoding/base64"
	"strings"

	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/numbers"
	"github.com/trustwallet/golibs/types"
)

func (p *Platform) GetTxsByAddress(address string) (types.Txs, error) {
	addressTxs, err := p.client.GetTxs(address, types.TxPerPage)
	if err != nil {
		return nil, err
	}

	var txs types.Txs
	for _, srcTx := range addressTxs {
		tx, err := NormalizeTx(&srcTx)
		if err != nil {
			continue
		}
		txs = append(txs, tx)
	}
	return txs, nil
}

func NormalizeTx(srcTx *Transaction) (types.Tx, error) {
	txValue := srcTx.TxValue
	decimals := coin.Aeternity().Decimals
	amountFloat, err := txValue.Amount.Float64()
	if err != nil {
		return types.Tx{}, err
	}
	amount := numbers.Float64toString(amountFloat)
	return types.Tx{
		ID:       srcTx.Hash,
		Coin:     coin.AETERNITY,
		From:     txValue.Sender,
		To:       txValue.Recipient,
		Fee:      types.Amount(txValue.Fee),
		Date:     srcTx.Timestamp / 1000,
		Block:    srcTx.BlockHeight,
		Memo:     getPayload(txValue.Payload),
		Status:   types.StatusCompleted,
		Sequence: txValue.Nonce,
		Meta: types.Transfer{
			Value:    types.Amount(amount),
			Symbol:   coin.Aeternity().Symbol,
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
