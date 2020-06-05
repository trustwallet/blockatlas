package address

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getTLD(t *testing.T) {
	tests := []struct {
		name, separator, wantTLD string
	}{
		{"vitalik.eth", ".", ".eth"},
		{"vitalik.ens", ".", ".ens"},
		{"ourxyzwallet.xyz", ".", ".xyz"},
		{"Cameron.Kred", ".", ".Kred"},
		{"btc.zil", ".", ".zil"},
		{"btc.crypto", ".", ".crypto"},
		{"nick@fiotestnet", "@", "@fiotestnet"},
		{"a", ".", ""},
		{"a.", ".", ""},
		{"a.b", ".", ".b"},
		{"a@b.c", ".", ".c"},
		{"a@b.c", "@", "@b.c"},
	}
	for _, tt := range tests {
		result := GetTLD(tt.name, tt.separator)
		assert.Equal(t, tt.wantTLD, result)
	}
}
