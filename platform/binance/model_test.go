package binance

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var newOrderDataDst = Data{OrderData: OrderData{
	Symbol:   "AWC-986_BNB",
	Base:     "AWC-986",
	Quote:    "BNB",
	Quantity: 2.0,
	Price:    0.00324939,
}}

var cancelOrderDataDst = Data{OrderData: OrderData{
	Symbol:   "GTO-908_BNB",
	Base:     "GTO-908",
	Quote:    "BNB",
	Quantity: 1.0,
	Price:    0.00104716,
}}

func TestTx_getData(t *testing.T) {
	tests := []struct {
		name string
		Data string
		want Data
	}{
		{
			"new order",
			"{\"orderData\":{\"symbol\":\"AWC-986_BNB\",\"orderType\":\"limit\",\"side\":\"buy\",\"price\":0.00324939,\"quantity\":2.00000000,\"timeInForce\":\"GTE\",\"orderId\":\"D13BAF4BD6638FA3AAD6EBCA0E4BEEA73DF4D519-30\"}}",
			newOrderDataDst,
		},
		{
			"cancel order",
			"{\"orderData\":{\"symbol\":\"GTO-908_BNB\",\"orderType\":\"limit\",\"side\":\"buy\",\"price\":0.00104716,\"quantity\":1.00000000,\"timeInForce\":\"GTE\",\"orderId\":\"D13BAF4BD6638FA3AAD6EBCA0E4BEEA73DF4D519-28\"}}",
			cancelOrderDataDst,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := &Tx{Data: tt.Data}
			got, _ := tx.getData()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestConvertValue(t *testing.T) {
	tests := []struct {
		name       string
		value      interface{}
		wantResult float64
		wantError  bool
	}{
		{"test string 1", "9", 9, false},
		{"test number 1", 9, 9, false},
		{"test string 2", "9380938973", 9380938973, false},
		{"test number 2", 9380938973, 9380938973, false},
		{"test string 3", "0.0000003", 0.0000003, false},
		{"test number 3", 0.0000003, 0.0000003, false},
		{"test string 4", "0.44", 0.44, false},
		{"test number 4", 0.44, 0.44, false},
		{"test string 5", "3334", 3334, false},
		{"test number 5", 3334, 3334, false},
		{"test error", time.Time{}, 3334, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := convertValue(tt.value)
			if tt.wantError {
				assert.False(t, ok)
				return
			}
			assert.True(t, ok)
			assert.Equal(t, tt.wantResult, got)
		})
	}
}

func Test_removeFloatPoint(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  int64
	}{
		{"test float 1", 0.0034, 340000},
		{"test float 2", 0.00000013, 13},
		{"test float 3", 0.938984, 93898400},
		{"test float 4", 0.1, 10000000},
		{"test int 1", 12, 1200000000},
		{"test int 2", 2333333333, 233333333300000000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeFloatPoint(tt.value); got != tt.want {
				t.Errorf("removeFloatPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

// currently bad string value in balance object will be parsed as 0 and if others will be not zero - we will return it
// we can ignore object, if any value is not correct (negative, too big, empty, not a number)
func Test_isZeroBalance(t *testing.T) {
	type testZeroStruct struct {
		name    string
		balance Balance
		want    bool
	}
	// all combinations of 3 variables with 2 possible value 0 or 1 is 2^3 = 8
	tests := []testZeroStruct{
		{"1",
			Balance{"0.00000000", "0.00000000", "0.00000000", "BNB"},
			true,
		},
		{"2",
			Balance{"0.00000000", "0", "0.00000001", "BNB"},
			false,
		},
		{"3",
			Balance{"0.00000000", "0.00000001", "0.00000000", "BNB"},
			false,
		},
		{"4",
			Balance{"0.00000000", "0.00000001", "0.00000001", "BNB"},
			false,
		},
		{"5",
			Balance{"0.00000001", "0.00000000", "0.00000000", "BNB"},
			false,
		},
		{"6",
			Balance{"0.00000001", "0.00000000", "0.00000001", "BNB"},
			false,
		},
		{"7",
			Balance{"0.00000001", "0.00000001", "0.00000000", "BNB"},
			false,
		},
		{"8",
			Balance{"0.00000001", "0.00000001", "0.00000001", "BNB"},
			false,
		},
		{"Negative",
			Balance{"-0.00000001", "0.00000001", "0.00000001", "BNB"},
			false,
		},
		{"Bad others are 0",
			Balance{"f", // "f" will be parsed as 0 by Parse Float
				"0.0000000", "0.0000000", "BNB"},
			true,
		},
		{"Bad others are not 0",
			Balance{"f", // "f" will be parsed as 0 by Parse Float
				"0.0000001", "0.0000000", "BNB"},
			false,
		},
		{"Empty others are not 0",
			Balance{"", // "" will be parsed as 0 by Parse Float
				"0.00000001", "0.00000001", "BNB"},
			false,
		},
		{"Empty  others are 0",
			Balance{"", // "" will be parsed as 0 by Parse Float
				"0.00000000", "0.00000000", "BNB"},
			true,
		},
		{"Big",
			Balance{"9999999999999999999999999999999999999999999999999999999999999999999999999999999999" +
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
				"99999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999",
				// free will be +Inf and converted as zero
				"0.00000000", "0.00000000", "BNB"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.balance.isZeroBalance(); got != tt.want {
				t.Errorf("isZeroBalance() = %v, want %v, name %v", got, tt.want, tt.name)
			}
		})
	}
}
