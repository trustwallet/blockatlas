package nimiq

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"sort"
	"testing"
	"time"
)

const (
	basicSrc = `
{
	"hash": "8b219949f4c1dfe9e7a9cdc5dbbc507e40dc16f44a1a5182ed6125c9a6891a50",
	"blockHash": "ab36a0909c6ed5761a984ef261d9c3456b7c1aea6a52d531c5bf2518526a32e6",
	"blockNumber": 252575,
	"timestamp": 1538924505,
	"confirmations": 271245,
	"transactionIndex": 37,
	"from": "4a88aaad038f9b8248865c4b9249efc554960e16",
	"fromAddress": "NQ69 9A4A MB83 HXDQ 4J46 BH5R 4JFF QMA9 C3GN",
	"to": "ad25610feb43d75307763d3f010822a757027429",
	"toAddress": "NQ15 MLJN 23YB 8FBM 61TN 7LYG 2212 LVBG 4V19",
	"value": 10000000000000,
	"fee": 138,
	"data": null,
	"flags": 0
}
`
	pendingSrc = `
{
  "data": null,
  "fee": 300,
  "flags": 0,
  "from": "d48182276127b149a9710e78c436fb4bc1c4dc0b",
  "fromAddress": "NQ74 SJ0Q 49T1 4XQL KABH 1RUC 8DPT 9F0U 9P0B",
  "hash": "79719d16f3f347cc98c35cd7a9af708cdce97de578b5135c5ae4393fd7920d61",
  "to": "0a17218ffdc42385f45329ba4089919236dd2743",
  "toAddress": "NQ97 18BJ 33YV QGHQ BV2K 56V4 12CH J8TD S9S3",
  "value": 100000
}
`
)

var (
	basicDst = blockatlas.Tx{
		ID:    "8b219949f4c1dfe9e7a9cdc5dbbc507e40dc16f44a1a5182ed6125c9a6891a50",
		Coin:  coin.NIM,
		From:  "NQ69 9A4A MB83 HXDQ 4J46 BH5R 4JFF QMA9 C3GN",
		To:    "NQ15 MLJN 23YB 8FBM 61TN 7LYG 2212 LVBG 4V19",
		Fee:   "138",
		Date:  1538924505,
		Block: 252575,
		Meta: blockatlas.Transfer{
			Value:    "10000000000000",
			Symbol:   "NIM",
			Decimals: 5,
		},
	}
	pendingDst = blockatlas.Tx{
		ID:    "79719d16f3f347cc98c35cd7a9af708cdce97de578b5135c5ae4393fd7920d61",
		Coin:  coin.NIM,
		From:  "NQ74 SJ0Q 49T1 4XQL KABH 1RUC 8DPT 9F0U 9P0B",
		To:    "NQ97 18BJ 33YV QGHQ BV2K 56V4 12CH J8TD S9S3",
		Fee:   "300",
		Date:  time.Now().Unix(),
		Block: 0,
		Meta: blockatlas.Transfer{
			Value:    "100000",
			Symbol:   "NIM",
			Decimals: 5,
		},
	}
)

func TestNormalizeTx1(t *testing.T) {
	tests := []struct {
		name  string
		srcTx string
		want  blockatlas.Tx
	}{
		{"test transaction", basicSrc, basicDst},
		{"test pending transaction", pendingSrc, pendingDst},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var srcTx Tx
			err := json.Unmarshal([]byte(tt.srcTx), &srcTx)
			if err != nil {
				t.Error(err)
				return
			}
			got := NormalizeTx(&srcTx)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNormalizeTxs_Ordering(t *testing.T) {
	_, goFile, _, _ := runtime.Caller(0)
	testFilePath := filepath.Join(filepath.Dir(goFile), "tests", "getTransactionsByAddress_50.json")
	testFile, err := ioutil.ReadFile(testFilePath)
	if err != nil {
		t.Fatal(err)
	}
	var srcTxs []Tx
	if err := json.Unmarshal(testFile, &srcTxs); err != nil {
		t.Fatal(err)
	}
	txs := NormalizeTxs(srcTxs)
	if len(txs) != 4 {
		t.Fatalf("Unexpected count: %d", len(txs))
	}
	sorted := sort.SliceIsSorted(txs, func(i, j int) bool {
		return txs[i].Block > txs[j].Block
	})
	if !sorted {
		t.Fatal("Transactions not sorted")
	}
}
