package binance

import (
	//"github.com/stretchr/testify/assert"
	//"sort"
	"testing"
)

//var (
//	newOrderDataDst    = Data{OrderData: OrderData{Symbol: "AWC-986_BNB", Base: "AWC-986", Quote: "BNB", Quantity: 2.0, Price: 0.00324939}}
//	cancelOrderDataDst = Data{OrderData: OrderData{Symbol: "GTO-908_BNB", Base: "GTO-908", Quote: "BNB", Quantity: 1.0, Price: 0.00104716}}
//)

//func TestTx_getData(t *testing.T) {
//	tests := []struct {
//		name string
//		Data string
//		want Data
//	}{
//		{
//			"new order",
//			"{\"orderData\":{\"symbol\":\"AWC-986_BNB\",\"orderType\":\"limit\",\"side\":\"buy\",\"price\":0.00324939,\"quantity\":2.00000000,\"timeInForce\":\"GTE\",\"orderId\":\"D13BAF4BD6638FA3AAD6EBCA0E4BEEA73DF4D519-30\"}}",
//			newOrderDataDst,
//		},
//		{
//			"cancel order",
//			"{\"orderData\":{\"symbol\":\"GTO-908_BNB\",\"orderType\":\"limit\",\"side\":\"buy\",\"price\":0.00104716,\"quantity\":1.00000000,\"timeInForce\":\"GTE\",\"orderId\":\"D13BAF4BD6638FA3AAD6EBCA0E4BEEA73DF4D519-28\"}}",
//			cancelOrderDataDst,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tx := &Tx{Data: tt.Data}
//			got, _ := tx.getData()
//			assert.Equal(t, tt.want, got)
//		})
//	}
//}

//func TestConvertValue(t *testing.T) {
//	tests := []struct {
//		name       string
//		value      interface{}
//		wantResult float64
//		wantError  bool
//	}{
//		{"test string 1", "9", 9, false},
//		{"test number 1", 9, 9, false},
//		{"test string 2", "9380938973", 9380938973, false},
//		{"test number 2", 9380938973, 9380938973, false},
//		{"test string 3", "0.0000003", 0.0000003, false},
//		{"test number 3", 0.0000003, 0.0000003, false},
//		{"test string 4", "0.44", 0.44, false},
//		{"test number 4", 0.44, 0.44, false},
//		{"test string 5", "3334", 3334, false},
//		{"test number 5", 3334, 3334, false},
//		{"test error", time.Time{}, 3334, true},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, ok := convertValue(tt.value)
//			if tt.wantError {
//				assert.False(t, ok)
//				return
//			}
//			assert.True(t, ok)
//			assert.Equal(t, tt.wantResult, got)
//		})
//	}
//}

//func Test_removeFloatPoint(t *testing.T) {
//	tests := []struct {
//		name  string
//		value float64
//		want  int64
//	}{
//		{"test float 1", 0.0034, 340000},
//		{"test float 2", 0.00000013, 13},
//		{"test float 3", 0.938984, 93898400},
//		{"test float 4", 0.1, 10000000},
//		{"test int 1", 12, 1200000000},
//		{"test int 2", 2333333333, 233333333300000000},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := removeFloatPoint(tt.value); got != tt.want {
//				t.Errorf("removeFloatPoint() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

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
	txDst = TxBase{
		TxHash:        "C29D822EFBC0C91656D1C5870BA55922F3A72A25BC8415B32D1D1AD0C85142F5",
		BlockHeight: 63591484,
		Type:        "TRANSFER",
		Value:       "0.00000001",
		Asset:       "BNB",
		FromAddr:    "bnb1nm4n03x00gw0x6v784jzryyp6wxnjaxswr3xm8",
		ToAddr:      "bnb1eff4hzx4lfsun3px5walchdy4vek4n0njcdzyn",
		Fee:         "0.0006",
	}
	txTokenDst = TxBase{
		TxHash:        "C29D822EFBC0C91656D1C5870BA55922F3A72A25BC8415B32D1D1AD0C85142F5",
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
			tx := &TxBase{
				FromAddr: tt.fields.FromAddr,
				ToAddr:   tt.fields.ToAddr,
			}
			if got := tx.containAddress(tt.address); got != tt.want {
				t.Errorf("containAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTx_getFee(t *testing.T) {
	tests := []struct {
		name string
		fee  string
		want string
	}{
		{"test empty", "", "0"},
		{"test error", "test", "0"},
		{"test float 1", "444.5", "44450000000"},
		{"test float 2", "0.00000001", "1"},
		{"test int", "3", "300000000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := &TxV1{TxBase{Fee: tt.fee}}
			if got := tx.getFee(); got != tt.want {
				t.Errorf("getFee() = %v, want %v", got, tt.want)
			}
		})
	}
}
