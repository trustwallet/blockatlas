package binance

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
		wantResult int64
	}{
		{"test string 1", "9", 900000000},
		{"test string 2", "9380938973", 938093897300000000},
		{"test string 3", "0.00000000003", 0},
		{"test string 4", "0.0000003", 30},
		{"test int 1", 32424234, 3242423400000000},
		{"test int 2", 34, 3400000000},
		{"test float 1", 2233.222, 223322200000},
		{"test float 2", 0.00000000003, 0},
		{"test float 3", 0.0000003, 30},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := ConvertValue(tt.value)
			if got != tt.wantResult {
				t.Errorf("ConvertValue() got = %v, want %v", got, tt.wantResult)
			}
		})
	}
}
