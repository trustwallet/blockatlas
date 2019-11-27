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
