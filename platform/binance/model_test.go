package binance

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"strconv"
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

//

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
		{"test float 3", "0.00037500", "37500"}, // standard transfer fee
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

func Test_getError(t *testing.T) {
	tests := []struct {
		name   string
		trx    ExplorerTxs
		expect string
	}{
		{"Should not have error message", ExplorerTxs{Code: 0}, ""},
		{"Should have error message", ExplorerTxs{Code: 1}, "error"},
		{"Should have error message", ExplorerTxs{Code: -1}, "error"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tt.trx.getError(), tt.expect)
		})
	}
}

func Test_QuantityTransferType(t *testing.T) {
	tests := []struct {
		name   string
		trx    ExplorerTxs
		expect QuantityTransfer
	}{
		{"Should be multi transfer", ExplorerTxs{HasChildren: 1}, MultiTransferOperation},
		{"Should be single transfer", ExplorerTxs{HasChildren: 0}, SingleTransferOperation},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.trx.QuantityTransferType(), tt.expect)
		})
	}
}

func TestExplorerTxs_getDirection(t *testing.T) {
	const (
		addr1 = "bnb14cjy0yl23xkf0hnw3ql295v8nghqstvlzkvqpl"
		addr2 = "bnb1mr5f97rx5wnkfcakx9fcpvljmx2s6kwqc08yur"
	)

	tests := []struct {
		name    string
		address string
		trx     ExplorerTxs
		expect  blockatlas.Direction
	}{
		{"getDirection should be self send", addr1, ExplorerTxs{FromAddr: addr1, ToAddr: addr1}, blockatlas.DirectionSelf},
		{"getDirection should be incoming", addr1, ExplorerTxs{FromAddr: addr2, ToAddr: addr1}, blockatlas.DirectionIncoming},
		{"getDirection should be outgoing", addr1, ExplorerTxs{FromAddr: addr1, ToAddr: addr2}, blockatlas.DirectionOutgoing},
		{"getDirection should be empty", "", ExplorerTxs{FromAddr: addr1, ToAddr: addr2}, blockatlas.Direction("")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.trx.getDirection(tt.address), tt.expect)
		})
	}
}

func Test_getDexFee(t *testing.T) {
	tests := []struct {
		name      string
		trx       ExplorerTxs
		expectFee blockatlas.Amount
	}{
		{"Should have zero fee", ExplorerTxs{TxFee: 0}, blockatlas.Amount(0)},
		{"Should have zero fee", ExplorerTxs{TxFee: 0.0}, blockatlas.Amount(0)},
		{"Should have standard fee", ExplorerTxs{TxFee: 0.00037500}, blockatlas.Amount(strconv.Itoa(37500))},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			fee := tt.trx.getDexFee()
			assert.Equal(t, fee, tt.expectFee)
		})
	}
}
