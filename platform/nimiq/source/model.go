package source

import "errors"


/* "hash": "745e19018e785cd8f05219578cceb6620d32f9c500ea1e4e9c0e416216984fe7",
"blockHash": "119657593bf6ac9e2b2a46ce28cd36a016ee0277f585235297c6a1e01b918a24",
"blockNumber": 79569,
"timestamp": 1528487555,
"confirmations": 156156,
"transactionIndex": 0,
"from": "4a88aaad038f9b8248865c4b9249efc554960e16",
"fromAddress": "NQ69 9A4A MB83 HXDQ 4J46 BH5R 4JFF QMA9 C3GN",
"to": "ad25610feb43d75307763d3f010822a757027429",
"toAddress": "NQ15 MLJN 23YB 8FBM 61TN 7LYG 2212 LVBG 4V19",
"value": 8000000000000,
"fee": 0,
"data": null,
"flags": 0 */

type Tx struct {
	Hash          string `json:"hash"`
	BlockHash     string `json:"blockHash"`
	BlockNumber   uint64 `json:"blockNumber"`
	Timestamp     string `json:"timestamp"`
	Confirmations int    `json:"confirmations"`
	TxIndex       int    `json:"transactionIndex"`
	FromAddress   string `json:"fromAddress"`
	ToAddress     string `json:"toAddress"`
	Value         uint64 `json:"value"`
	Fee           uint64 `json:"fee"`
}

var ErrSourceConn  = errors.New("connection to servers failed")
var ErrInvalidAddr = errors.New("invalid address")
