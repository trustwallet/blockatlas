package rpc

import (
	"reflect"
	"testing"
)

func TestBlockTxs_txs(t *testing.T) {
	tests := []struct {
		name string
		b    BlockTxs
		want []string
	}{
		{"test 1 tx", BlockTxs{{"tx1"}, {}}, []string{"tx1"}},
		{"test 2 txs  1", BlockTxs{{"tx1"}, {"tx2"}}, []string{"tx1", "tx2"}},
		{"test 2 txs 2", BlockTxs{{"tx1", "tx2"}}, []string{"tx1", "tx2"}},
		{"test 3 txs 1", BlockTxs{{"tx1", "tx2"}, {"tx3"}}, []string{"tx1", "tx2", "tx3"}},
		{"test 3 txs 2", BlockTxs{{"tx1"}, {"tx2"}, {"tx3"}}, []string{"tx1", "tx2", "tx3"}},
		{"test 4 txs", BlockTxs{{"tx1", "tx2"}, {"tx3"}, {"tx4"}}, []string{"tx1", "tx2", "tx3", "tx4"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.txs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("txs() = %v, want %v", got, tt.want)
			}
		})
	}
}
