package semux

import (
	"bytes"
	"encoding/json"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"testing"
)

const getAccountResponseStr = `
{
  "message": "Request processed successfully",
  "result": {
    "address": "0x8197987c401a3466ad678b2080b24838ebd95b41",
    "available": "194960598272319",
    "locked": "0",
    "nonce": "1190",
    "pendingTransactionCount": 1,
    "transactionCount": 3797
  },
  "success": true
}
`

const getAccountTransactionsResponseStr = `
{
  "message": "Request processed successfully",
  "result": [
    {
      "blockNumber": "1333625",
      "data": "0x",
      "fee": "5000000",
      "from": "0x8197987c401a3466ad678b2080b24838ebd95b41",
      "hash": "0x76274d1e328882095ad0369ea5e5bdf2c3c233c8715365092dfab542d33a1142",
      "nonce": "1189",
      "timestamp": "1557392426624",
      "to": "0xf1c36b97b48a71ffbc17c687f71491324366f8c0",
      "type": "TRANSFER",
      "value": "10266598890"
    }
  ],
  "success": true
}
`

var basicDst = blockatlas.Tx{
	ID:    "0x76274d1e328882095ad0369ea5e5bdf2c3c233c8715365092dfab542d33a1142",
	Coin:  coin.SEM,
	From:  "0x8197987c401a3466ad678b2080b24838ebd95b41",
	To:    "0xf1c36b97b48a71ffbc17c687f71491324366f8c0",
	Fee:   "5000000",
	Date:  1557392426,
	Block: 1333625,
	Meta: blockatlas.Transfer{
		Value: "10266598890",
	},
}

func TestUnmarshalGetAccountResponse(t *testing.T) {
	var getAccountResponse GetAccountResponse
	err := json.Unmarshal([]byte(getAccountResponseStr), &getAccountResponse)
	if err != nil {
		t.Fatal(err)
	}

	if getAccountResponse.Result.TransactionCount != 3797 {
		t.Error("UnmarshalGetAccountResponse failed")
	}
}

func TestNormalize(t *testing.T) {
	var getAccountTransactionsResponse GetAccountTransactionsResponse
	err := json.Unmarshal([]byte(getAccountTransactionsResponseStr), &getAccountTransactionsResponse)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := Normalize(&getAccountTransactionsResponse.Result[0])
	if err != nil {
		t.Fatal(err)
	}

	resJSON, err := json.Marshal(&tx)
	if err != nil {
		t.Fatal(err)
	}

	dstJSON, err := json.Marshal(&basicDst)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(resJSON, dstJSON) {
		println(string(resJSON))
		println(string(dstJSON))
		t.Error("basic: tx don't equal")
	}
}
