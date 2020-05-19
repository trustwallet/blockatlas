package numbers

import (
	"testing"
)

func Test_addAmount(t *testing.T) {
	type args struct {
		left  string
		right string
	}
	tests := []struct {
		name    string
		args    args
		wantSum string
	}{
		{"test zero + float", args{left: "0", right: "0.33333"}, "33333000"},
		{"test zero + int", args{left: "0", right: "333"}, "333"},
		{"test zero + zero", args{left: "0", right: "0"}, "0"},
		{"test int + float", args{left: "232", right: "0.222"}, "22200232"},
		{"test int + int", args{left: "661", right: "12"}, "673"},
		{"test int + zero", args{left: "131", right: "0"}, "131"},
		{"test float + float", args{left: "0.4141", right: "0.11211"}, "52621000"},
		{"test float + int", args{left: "3.111", right: "11"}, "311100011"},
		{"test float + zero", args{left: "0.455", right: "0"}, "45500000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := AddAmount(tt.args.left, tt.args.right); gotSum != tt.wantSum {
				t.Errorf("AddAmount() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

func Test_ToSatoshi(t *testing.T) {
	tests := []struct {
		name   string
		amount string
		want   int64
	}{
		{"test float", "0.33333", 33333000},
		{"test int", "3333", 333300000000},
		{"test zero", "0", 0},
		{"test error", "trust", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToSatoshi(tt.amount); got != tt.want {
				t.Errorf("ToSatoshi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getValue(t *testing.T) {
	tests := []struct {
		name   string
		amount string
		want   string
	}{
		{"test float", "0.33333", "33333000"},
		{"test int", "3333", "3333"},
		{"test zero", "0", "0"},
		{"test error", "trust", "0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAmountValue(tt.amount); got != tt.want {
				t.Errorf("GetAmountValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseAmount(t *testing.T) {
	tests := []struct {
		name   string
		amount string
		want   int64
	}{
		{"test float", "0.33333", 33333000},
		{"test int", "3333", 3333},
		{"test zero", "0", 0},
		{"test error", "trust", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseAmount(tt.amount); got != tt.want {
				t.Errorf("ParseAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}
