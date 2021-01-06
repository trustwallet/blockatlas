package binance

import "testing"

func Test_isZeroBalance(t *testing.T) {
	type testZeroStruct struct {
		name    string
		balance TokenBalance
		want    bool
	}
	tests := []testZeroStruct{
		{"1", TokenBalance{"0.00000000", "0.00000000", "0.00000000", "BNB"}, true},
		{"2", TokenBalance{"0.00000000", "0", "0.00000001", "BNB"}, false},
		{"3", TokenBalance{"0.00000000", "0.00000001", "0.00000000", "BNB"}, false},
		{"4", TokenBalance{"0.00000000", "0.00000001", "0.00000001", "BNB"}, false},
		{"5", TokenBalance{"0.00000001", "0.00000000", "0.00000000", "BNB"}, false},
		{"6", TokenBalance{"0.00000001", "0.00000000", "0.00000001", "BNB"}, false},
		{"7", TokenBalance{"0.00000001", "0.00000001", "0.00000000", "BNB"}, false},
		{"8", TokenBalance{"0.00000001", "0.00000001", "0.00000001", "BNB"}, false},
		{"Negative", TokenBalance{"-0.00000001", "0.00000001", "0.00000001", "BNB"}, false},
		{"Bad others are 0", TokenBalance{"f", "0.0000000", "0.0000000", "BNB"}, false},
		{"Bad others are not 0", TokenBalance{"f", "0.0000001", "0.0000000", "BNB"}, false},
		{"Empty others are not 0", TokenBalance{"", "0.00000001", "0.00000001", "BNB"}, false},
		{"Empty others are 0", TokenBalance{"", "0.00000000", "0.00000000", "BNB"}, false},
		{"Big", TokenBalance{"9999999999999999999999999999999999999999999999999999999999999999999999999999999999" +
			"9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "0.00000000", "0.00000000", "BNB"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.balance.isAllZeroBalance(); got != tt.want {
				t.Errorf("isAllZeroBalance() = %v, want %v, name %v", got, tt.want, tt.name)
			}
		})
	}
}
