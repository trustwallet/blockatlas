package binance

import (
	"reflect"
	"testing"
)

var newOrderDataDst = Data{OrderData: OrderData{
	Symbol: "AWC-986_BNB",
	Base:   "AWC-986",
	Quote:  "BNB",
	Price:  0.00324939,
}}

var cancelOrderDataDst = Data{OrderData: OrderData{
	Symbol: "GTO-908_BNB",
	Base:   "GTO-908",
	Quote:  "BNB",
	Price:  0.00104716,
}}

func TestTx_getData(t *testing.T) {
	tests := []struct {
		name    string
		Data    string
		want    Data
		wantErr bool
	}{
		{
			"new order",
			"{\"orderData\":{\"symbol\":\"AWC-986_BNB\",\"orderType\":\"limit\",\"side\":\"buy\",\"price\":0.00324939,\"quantity\":2.00000000,\"timeInForce\":\"GTE\",\"orderId\":\"D13BAF4BD6638FA3AAD6EBCA0E4BEEA73DF4D519-30\"}}",
			newOrderDataDst,
			false,
		},
		{
			"cancel order",
			"{\"orderData\":{\"symbol\":\"GTO-908_BNB\",\"orderType\":\"limit\",\"side\":\"buy\",\"price\":0.00104716,\"quantity\":1.00000000,\"timeInForce\":\"GTE\",\"orderId\":\"D13BAF4BD6638FA3AAD6EBCA0E4BEEA73DF4D519-28\"}}",
			cancelOrderDataDst,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := &Tx{Data: tt.Data}
			got, err := tx.getData()
			if (err != nil) != tt.wantErr {
				t.Errorf("getData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getData() got = %v, want %v", got, tt.want)
			}
		})
	}
}
