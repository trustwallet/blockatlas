package binance

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)
const (
	addr1 = "bnb13a7gyv5zl57c0rzeu0henx6d0tzspvrrakxxtv"
	addr2 = "bnb1t6tnm2rckd3pfptngj6u8466v3ah4fcdu78n5y"
)

const (
	bnbTransferResponse = `
        {
            "txHash": "73176E5BFA5856AEAB9BAB1F3030E6F2B2F274324052E84562BE9BE70E1AAEE7",
            "blockHeight": 74821444,
            "txType": "TRANSFER",
            "timeStamp": "2020-03-16T05:34:38.947Z",
            "fromAddr": "bnb13a7gyv5zl57c0rzeu0henx6d0tzspvrrakxxtv",
            "toAddr": "bnb1t6tnm2rckd3pfptngj6u8466v3ah4fcdu78n5y",
            "value": "10.00000000",
            "txAsset": "BNB",
            "txFee": "0.00037500",
            "proposalId": null,
            "txAge": 1868055,
            "orderId": null,
            "code": 0,
            "data": null,
            "confirmBlocks": 0,
            "memo": "bnb-transfer",
            "source": 1,
            "sequence": 158
        }`

	bep2TransferResponse = `
		{
            "txHash": "73176E5BFA5856AEAB9BAB1F3030E6F2B2F274324052E84562BE9BE70E1AAEE7",
            "blockHeight": 74821444,
            "txType": "TRANSFER",
            "timeStamp": "2020-03-16T05:34:38.947Z",
            "fromAddr": "bnb13a7gyv5zl57c0rzeu0henx6d0tzspvrrakxxtv",
            "toAddr": "bnb1t6tnm2rckd3pfptngj6u8466v3ah4fcdu78n5y",
            "value": "50.00000000",
            "txAsset": "TWT-8C2",
            "txFee": "0.00037500",
            "proposalId": null,
            "txAge": 276251,
            "orderId": null,
            "code": 0,
            "data": null,
            "confirmBlocks": 0,
            "memo": "bep2-transfer",
            "source": 0,
            "sequence": 158
        }`
)

var (
	expectBNBTransfer = blockatlas.Tx{
		ID:        "73176E5BFA5856AEAB9BAB1F3030E6F2B2F274324052E84562BE9BE70E1AAEE7",
		Coin:      714,
		From:      addr1,
		To:        addr2,
		Fee:       "37500",
		Date:      1584336878,
		Block:     74821444,
		Status:    blockatlas.StatusCompleted,
		Error:     "",
		Sequence:  158,
		Type:      blockatlas.TxTransfer,
		Direction: blockatlas.DirectionOutgoing,
		Memo:      "bnb-transfer",
		Meta: blockatlas.Transfer{
			Value:    "1000000000",
			Symbol:   "BNB",
			Decimals: 8,
		},
	}

	expectBEP2Transfer = blockatlas.Tx{
		ID:        "73176E5BFA5856AEAB9BAB1F3030E6F2B2F274324052E84562BE9BE70E1AAEE7",
		Coin:      714,
		From:      addr1,
		To:        addr2,
		Fee:       "37500",
		Date:      1584336878,
		Block:     74821444,
		Status:    blockatlas.StatusCompleted,
		Error:     "",
		Sequence:  158,
		Type:      blockatlas.TxNativeTokenTransfer,
		Direction: blockatlas.DirectionIncoming,
		Memo:      "bep2-transfer",
		Meta: blockatlas.NativeTokenTransfer{
			Decimals: 8,
			From:     addr1,
			Name:     "",
			Symbol:   "TWT",
			To:       addr2,
			TokenID:  "TWT-8C2",
			Value:    "5000000000",
		},
	}
)

func TestNormalizeTxs(t *testing.T) {
	type testTxs struct {
		name        string
		apiResponse string
		expected    []blockatlas.Tx
		address     string
	}
	testTxsList := []testTxs{
		{name: "BNB transfer", apiResponse: bnbTransferResponse, expected: []blockatlas.Tx{expectBNBTransfer}, address: addr1},
		{name: "BEP2 transfer", apiResponse: bep2TransferResponse, expected: []blockatlas.Tx{expectBEP2Transfer}, address: addr2},
	}

	for _, testTxsInstance := range testTxsList {
		t.Run(testTxsInstance.name, func(t *testing.T) {
			var srcTxs []Tx
			err := json.Unmarshal([]byte(convertJsonToArray(testTxsInstance.apiResponse)), &srcTxs)
			assert.Nil(t, err)
			txs := NormalizeTxs(srcTxs, testTxsInstance.address)
			assert.Equal(t, testTxsInstance.expected, txs, "tx don't equal")
		})
	}
}

func TestTokenSymbol(t *testing.T) {
	assert.Equal(t, "UGAS", TokenSymbol("UGAS"))
	assert.Equal(t, "UGAS", TokenSymbol("UGAS-B0C"))
}

func convertJsonToArray(jsonString string) string {
	return "[" + jsonString + "]"
}
