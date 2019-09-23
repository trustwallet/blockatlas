package vechain

import (
	"bytes"
	"encoding/json"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

const transferReceipt = `{
    "block": 2620166,
    "id": "0x2b8776bd4679fa2afa28b55d66d4f6c7c77522fc878ce294d25e32475b704517",
    "nonce": "0x3657a2025b11f27f",
    "origin": "0xB853D6a965fBC047aaA9f04d774d53861d7eD653",
	"timestamp": 1556569300,
	"receipt": {
		"paid": "0x1236efcbcbb340000",
		"reverted": false
	}
}`

const transferFailedReceipt = `{
    "block": 2620166,
    "id": "0x2b8776bd4679fa2afa28b55d66d4f6c7c77522fc878ce294d25e32475b704517",
    "nonce": "0x3657a2025b11f27f",
    "origin": "0xB853D6a965fBC047aaA9f04d774d53861d7eD653",
	"timestamp": 1556569300,
	"receipt": {
		"paid": "0x1236efcbcbb340000",
		"reverted": false
	}
}`

const transferClause = `{
	"to": "0xdA623049A13df5c8a24f0d7713F4adD4ab136B1f",
	"value": "0x29bde5885d7ac80000"
}`

const tokenTransfer = `
{
	"amount": "0x00000000000000000000000000000000000000000000000d8d726b7177a80000",
	"block": 2620166,
	"origin": "0xB853D6a965fBC047aaA9f04d774d53861d7eD653",
	"receiver": "0x9f3742C2c2fE66c7FCa08d77d2262C22e3D56ac8",
	"timestamp": 1556569300,
	"txId": "0x2b8776bd4679fa2afa28b55d66d4f6c7c77522fc878ce294d25e32475b704517"
}
`

const transaction = `
{  
   "block":2620166,
   "blockRef":"0x003578d93e73a9ca",
   "chainTag":74,
   "clauses":[  
      {  
         "data":"0xa9059cbb0000000000000000000000009f3742c2c2fe66c7fca08d77d2262c22e3d56ac8000000000000000000000000000000000000000000000008ac7230489e800000",
         "numValue":0,
         "to":"0x0000000000000000000000000000456e65726779",
         "txClauseIndex":0,
         "txId":"0x2b8776bd4679fa2afa28b55d66d4f6c7c77522fc878ce294d25e32475b704517",
         "value":"0x0"
      }
   ],
   "expiration":720,
   "gas":160000,
   "gasPriceCoef":0,
   "id":"0x2b8776bd4679fa2afa28b55d66d4f6c7c77522fc878ce294d25e32475b704517",
   "meta":{  
      "blockID":"0x003578db7b662faecc743f3a401515eef5baebe16c27e635f79bcfca3b8a39dc",
      "blockNumber":3504347,
      "blockTimestamp":1565458520
   },
   "nonce":"0x8e60abce86ae",
   "numClauses":1,
   "origin":"0xb853d6a965fbc047aaa9f04d774d53861d7ed653",
   "receipt":{  
      "gasPayer":"0xa760bdcbf6c2935d2f1591a38f23251619f802ad",
      "gasUsed":36582,
      "meta":{  
         "blockID":"0x003578db7b662faecc743f3a401515eef5baebe16c27e635f79bcfca3b8a39dc",
         "blockNumber":3504347,
         "blockTimestamp":1565458520,
         "txID":"0x2b8776bd4679fa2afa28b55d66d4f6c7c77522fc878ce294d25e32475b704517",
         "txOrigin":"0xb853d6a965fbc047aaa9f04d774d53861d7ed653"
      },
      "outputs":[  
         {  
            "events":[  
               {  
                  "address":"0x0000000000000000000000000000456e65726779",
                  "topics":[  
                     "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
                     "0x000000000000000000000000b853d6a965fbc047aaa9f04d774d53861d7ed653",
                     "0x0000000000000000000000009f3742c2c2fe66c7fca08d77d2262c22e3d56ac8"
                  ],
                  "data":"0x00000000000000000000000000000000000000000000000d8d726b7177a80000"
               }
            ],
            "transfers":[  
               {  
                  "sender":"0xb853d6a965fbc047aaa9f04d774d53861d7ed653",
                  "recipient":"0xda623049a13df5c8a24f0d7713f4add4ab136b1f",
                  "amount":"0x29bde5885d7ac80000"
               }
            ]
         }
      ],
      "paid":"0x1236efcbcbb340000",
      "reverted":false,
      "reward":"0x984d9c8dd8008000"
   },
   "reverted":0,
   "size":191,
   "timestamp":1556569300,
   "totalValue":0
}
`

var expectedTransferTrx = blockatlas.Tx{
	ID:       "0x2b8776bd4679fa2afa28b55d66d4f6c7c77522fc878ce294d25e32475b704517",
	Coin:     coin.VET,
	From:     "0xB853D6a965fBC047aaA9f04d774d53861d7eD653",
	To:       "0xdA623049A13df5c8a24f0d7713F4adD4ab136B1f",
	Fee:      "21000000000000000000",
	Date:     1556569300,
	Type:     "transfer",
	Status:   blockatlas.StatusCompleted,
	Block:    2620166,
	Sequence: 2620166,
	Meta: blockatlas.Transfer{
		Value:    "770000000000000000000",
		Decimals: 18,
		Symbol:   "VET",
	},
}

var expectedVeThorTrx = blockatlas.Tx{
	ID:       "0x2b8776bd4679fa2afa28b55d66d4f6c7c77522fc878ce294d25e32475b704517",
	Coin:     coin.VET,
	From:     "0xB853D6a965fBC047aaA9f04d774d53861d7eD653",
	To:       "0x9f3742C2c2fE66c7FCa08d77d2262C22e3D56ac8",
	Fee:      "21000000000000000000",
	Date:     1556569300,
	Type:     blockatlas.TxNativeTokenTransfer,
	Status:   blockatlas.StatusCompleted,
	Sequence: 2620166,
	Block:    2620166,
	Meta: blockatlas.NativeTokenTransfer{
		Name:     GasName,
		Symbol:   GasSymbol,
		TokenID:  GasContract,
		Decimals: 18,
		Value:    "250000000000000000000",
		From:     "0xB853D6a965fBC047aaA9f04d774d53861d7eD653",
		To:       "0x9f3742C2c2fE66c7FCa08d77d2262C22e3D56ac8",
	},
}

func TestNormalizeTransfer(t *testing.T) {
	var tests = []struct {
		Receipt  string
		Clause   string
		Expected blockatlas.Tx
	}{
		{transferReceipt, transferClause, expectedTransferTrx},
		// {transferTrx, transferReceipt, transferOutput, address, VeThorContract, expectedVeThorTrx},
	}

	for _, test := range tests {
		var receipt TransferReceipt
		var clause Clause

		// Unmarshal(*t, test.Receipt, &receipt)
		rErr := json.Unmarshal([]byte(test.Receipt), &receipt)
		if rErr != nil {
			t.Fatal(rErr)
		}

		cErr := json.Unmarshal([]byte(test.Clause), &clause)
		if cErr != nil {
			t.Fatal(cErr)
		}

		var readyTx blockatlas.Tx
		normTx, err := NormalizeTransfer(&receipt, &clause)
		if err != nil {
			t.Fatal("VeChain: Can't normalize transfer", readyTx)
		}
		readyTx = normTx

		actual, err := json.Marshal(&readyTx)
		if err != nil {
			t.Fatal(err)
		}

		expected, err := json.Marshal(&test.Expected)
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(actual, expected) {
			println(string(actual))
			println(string(expected))
			t.Error("Transactions not equal")
		}
	}
}

func TestNormalizeTokenTransfer(t *testing.T) {
	var tests = []struct {
		Receipt  string
		Transfer string
		Expected blockatlas.Tx
	}{
		{transferFailedReceipt, tokenTransfer, expectedVeThorTrx},
	}

	for _, test := range tests {
		var receipt TransferReceipt
		var tt TokenTransfer

		// Unmarshal(*t, test.Receipt, &receipt)
		rErr := json.Unmarshal([]byte(test.Receipt), &receipt)
		if rErr != nil {
			t.Fatal(rErr)
		}

		ttErr := json.Unmarshal([]byte(test.Transfer), &tt)
		if ttErr != nil {
			t.Fatal(ttErr)
		}

		var readyTx blockatlas.Tx
		normTx, err := NormalizeTokenTransfer(&tt, &receipt)
		if err != nil {
			t.Fatal("VeChain: Can't normalize token transfer", readyTx)
		}
		readyTx = normTx

		actual, err := json.Marshal(&readyTx)
		if err != nil {
			t.Fatal(err)
		}

		expected, err := json.Marshal(&test.Expected)
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(actual, expected) {
			println(string(actual))
			println(string(expected))
			t.Error("Transactions not equal")
		}
	}
}

func TestNormalizeTransaction(t *testing.T) {
	var tests = []struct {
		Transaction         string
		ExpectedTransaction []blockatlas.Tx
	}{
		{transaction, []blockatlas.Tx{expectedVeThorTrx, expectedTransferTrx}},
	}

	for _, test := range tests {
		var transaction NativeTransaction

		tErr := json.Unmarshal([]byte(test.Transaction), &transaction)
		if tErr != nil {
			t.Fatal(tErr)
		}

		var readyTxs []blockatlas.Tx

		ntxs := NormalizeBlockTransactions(&transaction)
		readyTxs = append(readyTxs, ntxs...)

		actual, err := json.Marshal(&readyTxs)
		if err != nil {
			t.Fatal(err)
		}

		expectedTransactions, err := json.Marshal(&test.ExpectedTransaction)
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(actual, expectedTransactions) {
			println(string(actual))
			println(string(expectedTransactions))
			t.Error("Transactions not equal")
		}
	}
}
