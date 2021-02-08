package viewblock

import (
	"testing"
)

func TestTx_NonceValue(t *testing.T) {
	tests := []struct {
		name  string
		nonce interface{}
		want  uint64
	}{
		{"test int", 0, 0},
		{"test float", 3.4, 3},
		{"test string", "33", 33},
		{"test error string", "test", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := Tx{
				Nonce: tt.nonce,
			}
			if got := tx.NonceValue(); got != tt.want {
				t.Errorf("NonceValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
