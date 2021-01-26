package nimiq

import (
	"encoding/json"
	"fmt"
	"math"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/mock"
	"github.com/trustwallet/golibs/types"
)

var (
	basicSrc, _   = mock.JsonStringFromFilePath("mocks/" + "tx.json")
	pendingSrc, _ = mock.JsonStringFromFilePath("mocks/" + "pending_tx.json")
	basicDst      = types.Tx{
		ID:    "8b219949f4c1dfe9e7a9cdc5dbbc507e40dc16f44a1a5182ed6125c9a6891a50",
		Coin:  coin.NIMIQ,
		From:  "NQ69 9A4A MB83 HXDQ 4J46 BH5R 4JFF QMA9 C3GN",
		To:    "NQ15 MLJN 23YB 8FBM 61TN 7LYG 2212 LVBG 4V19",
		Fee:   "138",
		Date:  1538924505,
		Block: 252575,
		Meta: types.Transfer{
			Value:    "10000000000000",
			Symbol:   "NIM",
			Decimals: 5,
		},
	}
	pendingDst = types.Tx{
		ID:    "79719d16f3f347cc98c35cd7a9af708cdce97de578b5135c5ae4393fd7920d61",
		Coin:  coin.NIMIQ,
		From:  "NQ74 SJ0Q 49T1 4XQL KABH 1RUC 8DPT 9F0U 9P0B",
		To:    "NQ97 18BJ 33YV QGHQ BV2K 56V4 12CH J8TD S9S3",
		Fee:   "300",
		Date:  666666, // special placholder value
		Block: 0,
		Meta: types.Transfer{
			Value:    "100000",
			Symbol:   "NIM",
			Decimals: 5,
		},
	}
)

func TestNormalizeTx1(t *testing.T) {
	now := time.Now().Unix()
	tests := []struct {
		name  string
		srcTx string
		want  types.Tx
	}{
		{"test transaction", basicSrc, basicDst},
		{"test pending transaction", pendingSrc, pendingDst},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var srcTx Tx
			err := json.Unmarshal([]byte(tt.srcTx), &srcTx)
			if err != nil {
				fmt.Println(tt.srcTx)
				t.Error(err)
				return
			}
			got := NormalizeTx(&srcTx)
			// special handling for current date, if around now, replace with special value
			if math.Abs(float64(got.Date-now)) < 30 {
				got.Date = 666666
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNormalizeTxs_Ordering(t *testing.T) {
	var srcTxs []Tx
	_ = mock.JsonModelFromFilePath("mocks/getTransactionsByAddress_50.json", &srcTxs)
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
