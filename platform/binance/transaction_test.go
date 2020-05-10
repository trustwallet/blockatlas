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
	bnbSingleExplorerTransferResponse = `
        {
            "txHash": "73176E5BFA5856AEAB9BAB1F3030E6F2B2F274324052E84562BE9BE70E1AAEE7",
            "blockHeight": 74821444,
            "txType": "TRANSFER",
            "timeStamp": 1588086370574,
            "fromAddr": "bnb13a7gyv5zl57c0rzeu0henx6d0tzspvrrakxxtv",
            "toAddr": "bnb1t6tnm2rckd3pfptngj6u8466v3ah4fcdu78n5y",
            "value": 10.00000000,
            "txAsset": "BNB",
            "txFee": 0.00037500,
            "txAge": 1868055,
            "code": 0,
			"log": "Msg 0: ",
            "confirmBlocks": 0,
			"memo": "bnb-transfer",
			"source": 1,
            "hasChildren": 0
        }`

	bep2SingleExplorerTransferResponse = `
		{
            "txHash": "73176E5BFA5856AEAB9BAB1F3030E6F2B2F274324052E84562BE9BE70E1AAEE7",
            "blockHeight": 74821444,
            "txType": "TRANSFER",
            "timeStamp": 1588086357686,
            "fromAddr": "bnb13a7gyv5zl57c0rzeu0henx6d0tzspvrrakxxtv",
            "toAddr": "bnb1t6tnm2rckd3pfptngj6u8466v3ah4fcdu78n5y",
            "value": 2800.00000000,
            "txAsset": "TWT-8C2",
            "txFee": 0.00037500,
            "txAge": 276251,
            "code": 0,
            "log": "Msg 0: ",
            "confirmBlocks": 0,
            "memo": "bep2-transfer",
            "source": 0,
            "hasChildren": 0
        }`
	bep2MultipleExplorerTransferResponse = `
		{
  			"txHash": "73176E5BFA5856AEAB9BAB1F3030E6F2B2F274324052E84562BE9BE70E1AAEE7",
  			"blockHeight": 74821444,
  			"txType": "TRANSFER",
  			"timeStamp": 1588086357686,
  			"txFee": 0.00037500,
  			"txAge": 2068619,
  			"code": 0,
  			"log": "Msg 0: ",
  			"confirmBlocks": 0,
 			"memo": "bep2-transfer",
  			"source": 0,
  			"hasChildren": 1,
  			"subTxsDto": [
   			 {
      			"amount": "280000000000",
      			"asset": "TWT-8C2",
      			"from": "bnb13a7gyv5zl57c0rzeu0henx6d0tzspvrrakxxtv",
      			"to": "bnb1t6tnm2rckd3pfptngj6u8466v3ah4fcdu78n5y"
  			 }
  			 ]
		}`
)

var (
	expectBnbSingleExplorerTransfer = blockatlas.Tx{
		ID:        "73176E5BFA5856AEAB9BAB1F3030E6F2B2F274324052E84562BE9BE70E1AAEE7",
		Coin:      714,
		From:      addr1,
		To:        addr2,
		Fee:       "37500",
		Date:      1588086370,
		Block:     74821444,
		Status:    blockatlas.StatusCompleted,
		Error:     "",
		Sequence:  0,
		Type:      blockatlas.TxTransfer,
		Direction: blockatlas.DirectionOutgoing,
		Memo:      "bnb-transfer",
		Meta: blockatlas.Transfer{
			Value:    "1000000000",
			Symbol:   "BNB",
			Decimals: 8,
		},
	}

	expectBEP2SingleExplorerTransfer = blockatlas.Tx{
		ID:        "73176E5BFA5856AEAB9BAB1F3030E6F2B2F274324052E84562BE9BE70E1AAEE7",
		Coin:      714,
		From:      addr1,
		To:        addr2,
		Fee:       "37500",
		Date:      1588086357,
		Block:     74821444,
		Status:    blockatlas.StatusCompleted,
		Error:     "",
		Sequence:  0,
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
			Value:    "280000000000",
		},
	}
)

func TestNormalizeTxs(t *testing.T) {
	type test struct {
		name, address, token, dexTxResponse string
		expected                            []blockatlas.Tx
	}
	tests := []test{
		{name: "BNB single transfer", dexTxResponse: bnbSingleExplorerTransferResponse, expected: []blockatlas.Tx{expectBnbSingleExplorerTransfer}, address: addr1, token: ""},
		{name: "BEP2 single transfer", dexTxResponse: bep2SingleExplorerTransferResponse, expected: []blockatlas.Tx{expectBEP2SingleExplorerTransfer}, address: addr2, token: "TWT-8C2"},
		{name: "BEP2 multiple transfer", dexTxResponse: bep2MultipleExplorerTransferResponse, expected: []blockatlas.Tx{expectBEP2SingleExplorerTransfer}, address: addr2, token: "TWT-8C2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var srcTx ExplorerTxs
			err := json.Unmarshal([]byte(tt.dexTxResponse), &srcTx)
			assert.Nil(t, err)
			actual := normalizeTx(srcTx, tt.address, tt.token)
			assert.Equal(t, tt.expected, actual, "tx don't equal")
		})
	}
}

func TestTokenSymbol(t *testing.T) {
	assert.Equal(t, "UGAS", tokenSymbol("UGAS"))
	assert.Equal(t, "UGAS", tokenSymbol("UGAS-B0C"))
}
