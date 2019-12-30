package polkadot

import (
	"reflect"
	"testing"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func TestParseAmount(t *testing.T) {
	type args struct {
		source   string
		decimals uint
	}
	tests := []struct {
		name string
		args args
		want blockatlas.Amount
	}{
		{
			name: "KSM 1",
			args: args{
				source:   "0.01",
				decimals: 12,
			},
			want: "10000000000",
		},
		{
			name: "KSM 2",
			args: args{
				source:   "210",
				decimals: 12,
			},
			want: "210000000000000",
		},
		{
			name: "ETH",
			args: args{
				source:   "0.000000116639015088",
				decimals: 18,
			},
			want: "116639015088",
		},
		{
			name: "BTC",
			args: args{
				source:   "0.17368000",
				decimals: 8,
			},
			want: "17368000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseAmount(tt.args.source, tt.args.decimals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}
