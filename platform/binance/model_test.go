package binance

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	//"github.com/stretchr/testify/assert"
	//"sort"
	"gotest.tools/assert"
	"testing"
)

func Test_isZeroBalance(t *testing.T) {
	type testZeroStruct struct {
		name    string
		balance Balance
		want    bool
	}
	// all combinations of 3 variables with 2 possible value 0 or 1 is 2^3 = 8
	tests := []testZeroStruct{
		{"1", Balance{"0.00000000", "0.00000000", "0.00000000", "BNB"}, true},
		{"2", Balance{"0.00000000", "0", "0.00000001", "BNB"}, false},
		{"3", Balance{"0.00000000", "0.00000001", "0.00000000", "BNB"}, false},
		{"4", Balance{"0.00000000", "0.00000001", "0.00000001", "BNB"}, false},
		{"5", Balance{"0.00000001", "0.00000000", "0.00000000", "BNB"}, false},
		{"6", Balance{"0.00000001", "0.00000000", "0.00000001", "BNB"}, false},
		{"7", Balance{"0.00000001", "0.00000001", "0.00000000", "BNB"}, false},
		{"8", Balance{"0.00000001", "0.00000001", "0.00000001", "BNB"}, false},
		{"Negative", Balance{"-0.00000001", "0.00000001", "0.00000001", "BNB"}, false},
		{"Bad others are 0", Balance{"f", "0.0000000", "0.0000000", "BNB"}, false},
		{"Bad others are not 0", Balance{"f", "0.0000001", "0.0000000", "BNB"}, false},
		{"Empty others are not 0", Balance{"", "0.00000001", "0.00000001", "BNB"}, false},
		{"Empty others are 0", Balance{"", "0.00000000", "0.00000000", "BNB"}, false},
		{"Big", Balance{"9999999999999999999999999999999999999999999999999999999999999999999999999999999999" +
			"9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999" +
			"9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999" +
			"9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999" +
			"9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999" +
			"9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999" +
			"9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999" +
			"9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999" +
			"9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999" +
			"9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999" +
			"9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999" +
			"9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999" +
			"9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999" +
			"9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999" +
			"9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999" +
			"9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999" +
			"9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999" +
			"99999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "0.00000000", "0.00000000", "BNB"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.balance.isAllZeroBalance(); got != tt.want {
				t.Errorf("isAllZeroBalance() = %v, want %v, name %v", got, tt.want, tt.name)
			}
		})
	}
}

var (
	subTxDst = SubTx{
		Hash:     "C29D822EFBC0C91656D1C5870BA55922F3A72A25BC8415B32D1D1AD0C85142F5",
		Height:   63591484,
		Type:     "TRANSFER",
		Value:    "0.00000001",
		Asset:    "BNB",
		FromAddr: "bnb1nm4n03x00gw0x6v784jzryyp6wxnjaxswr3xm8",
		ToAddr:   "bnb1eff4hzx4lfsun3px5walchdy4vek4n0njcdzyn",
		Fee:      "0.0006",
	}
	subTxTokenDst = SubTx{
		Hash:     "C29D822EFBC0C91656D1C5870BA55922F3A72A25BC8415B32D1D1AD0C85142F5",
		Height:   63591485,
		Type:     "TRANSFER",
		Value:    "0.000064",
		Asset:    "AERGO-46B",
		FromAddr: "bnb1nm4n03x00gw0x6v784jzryyp6wxnjaxswr3xm8",
		ToAddr:   "bnb1eff4hzx4lfsun3px5walchdy4vek4n0njcdzyn",
		Fee:      "0.0006",
	}
	txDst = Tx{
		TxHash:      "C29D822EFBC0C91656D1C5870BA55922F3A72A25BC8415B32D1D1AD0C85142F5",
		BlockHeight: 63591484,
		Type:        "TRANSFER",
		Value:       "0.00000001",
		Asset:       "BNB",
		FromAddr:    "bnb1nm4n03x00gw0x6v784jzryyp6wxnjaxswr3xm8",
		ToAddr:      "bnb1eff4hzx4lfsun3px5walchdy4vek4n0njcdzyn",
		Fee:         "0.0006",
	}
	txTokenDst = Tx{
		TxHash:      "C29D822EFBC0C91656D1C5870BA55922F3A72A25BC8415B32D1D1AD0C85142F5",
		BlockHeight: 63591485,
		Type:        "TRANSFER",
		Value:       "0.000064",
		Asset:       "AERGO-46B",
		FromAddr:    "bnb1nm4n03x00gw0x6v784jzryyp6wxnjaxswr3xm8",
		ToAddr:      "bnb1eff4hzx4lfsun3px5walchdy4vek4n0njcdzyn",
		Fee:         "0.0006",
	}
)

//func TestSubTx_toTx(t *testing.T) {
//	tests := []struct {
//		name  string
//		subTx SubTx
//		want  Tx
//	}{
//		{"test conversion subTx to Tx", subTxTokenDst,
//			Tx{
//				Hash:        "C29D822EFBC0C91656D1C5870BA55922F3A72A25BC8415B32D1D1AD0C85142F5",
//				BlockHeight: 63591485,
//				Type:        TxTransfer,
//				FromAddr:    "bnb1nm4n03x00gw0x6v784jzryyp6wxnjaxswr3xm8",
//				ToAddr:      "bnb1eff4hzx4lfsun3px5walchdy4vek4n0njcdzyn",
//				Asset:       "AERGO-46B",
//				Fee:         "0.0006",
//				Value:       "0.000064",
//				SubTxsDto:   SubTxsDto{},
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got := tt.subTx.toTx()
//			assert.Equal(t, tt.want, got, "conversion failed")
//		})
//	}
//}

//func TestSubTxs_getTxs(t *testing.T) {
//	txDuplicatedTokenDst := txTokenDst
//	txDuplicatedTokenDst.Value = "0.000128"
//	txDuplicatedDst := txDst
//	txDuplicatedDst.Value = "0.00000002"
//	tests := []struct {
//		name    string
//		subTxs  SubTxs
//		wantTxs []Tx
//	}{
//		{"test empty", SubTxs{}, nil},
//		{"test subTx transfer", SubTxs{subTxDst}, []Tx{txDst}},
//		{"test subTx token transfer", SubTxs{subTxTokenDst}, []Tx{txTokenDst}},
//		{"test subTx and token transfer", SubTxs{subTxDst, subTxTokenDst}, []Tx{txDst, txTokenDst}},
//		{"test duplicate subTx token transfer", SubTxs{subTxTokenDst, subTxTokenDst}, []Tx{txDuplicatedTokenDst}},
//		{"test duplicate subTx", SubTxs{subTxDst, subTxDst}, []Tx{txDuplicatedDst}},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			gotTxs := tt.subTxs.getTxs()
//			sort.Slice(gotTxs, func(i, j int) bool {
//				return gotTxs[i].BlockHeight < gotTxs[j].BlockHeight
//			})
//			assert.Equal(t, tt.wantTxs, gotTxs, "get txs from subTxs failed")
//		})
//	}
//}

func TestTx_containAddress(t *testing.T) {
	type fields struct {
		FromAddr string
		ToAddr   string
	}
	tests := []struct {
		name    string
		fields  fields
		address string
		want    bool
	}{
		{"test from address valid", fields{FromAddr: "bnb1nm4n03x00gw0x6v784jzryyp6wxnjaxswr3xm8", ToAddr: "bnb1eff4hzx4lfsun3px5walchdy4vek4n0njcdzyn"}, "bnb1nm4n03x00gw0x6v784jzryyp6wxnjaxswr3xm8", true},
		{"test to address valid", fields{FromAddr: "bnb1nm4n03x00gw0x6v784jzryyp6wxnjaxswr3xm8", ToAddr: "bnb1eff4hzx4lfsun3px5walchdy4vek4n0njcdzyn"}, "bnb1eff4hzx4lfsun3px5walchdy4vek4n0njcdzyn", true},
		{"test no address valid", fields{FromAddr: "bnb1nm4n03x00gw0x6v784jzryyp6wxnjaxswr3xm8", ToAddr: "bnb1eff4hzx4lfsun3px5walchdy4vek4n0njcdzyn"}, "tbnb1qxm48ndhmh7su0r7zgwmwkltuqgly57jdf8yf8", false},
		{"test empty address", fields{FromAddr: "bnb1nm4n03x00gw0x6v784jzryyp6wxnjaxswr3xm8", ToAddr: "bnb1eff4hzx4lfsun3px5walchdy4vek4n0njcdzyn"}, "", true},
		{"test empty address without from", fields{FromAddr: "", ToAddr: "bnb1eff4hzx4lfsun3px5walchdy4vek4n0njcdzyn"}, "", true},
		{"test empty address without to", fields{FromAddr: "bnb1nm4n03x00gw0x6v784jzryyp6wxnjaxswr3xm8", ToAddr: ""}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := &Tx{
				FromAddr: tt.fields.FromAddr,
				ToAddr:   tt.fields.ToAddr,
			}
			if got := tx.containAddress(tt.address); got != tt.want {
				t.Errorf("containAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getFee(t *testing.T) {
	tests := []struct {
		name string
		fee  string
		want string
	}{
		{"test empty", "", "0"},
		{"test error", "test", "0"},
		{"test float 1", "444.5", "44450000000"},
		{"test float 2", "0.00000001", "1"},
		{"test float 3", "0.00037500", "37500"}, // standard fee
		{"test int", "3", "300000000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := &Tx{Fee: tt.fee}
			if got := tx.getFee(); got != tt.want {
				t.Errorf("getFee() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_blockTimestamp(t *testing.T) {
	tests := []struct {
		trx    Tx
		expect int64
	}{
		{Tx{Timestamp: "2020-03-16T05:34:38.947Z"}, 1584336878},
		{Tx{Timestamp: ""}, 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tt.trx.blockTimestamp(), tt.expect)
		})
	}
}

func Test_getStatus(t *testing.T) {
	tests := []struct {
		name   string
		trx    Tx
		expect blockatlas.Status
	}{
		{"Should have status completed", Tx{Code: 0}, blockatlas.StatusCompleted},
		{"Should have status error", Tx{Code: 1}, blockatlas.StatusError},
		{"Should have status error", Tx{Code: -1}, blockatlas.StatusError},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tt.trx.getStatus(), tt.expect)
		})
	}
}
