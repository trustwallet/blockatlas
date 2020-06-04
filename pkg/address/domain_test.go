package address

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func checkGetTLD(t *testing.T, name string, separator string, expectedTLD string) {
	name = strings.ToLower(name)
	tld := GetTLD(name, separator)
	assert.Equal(t, expectedTLD, tld)
}

func Test_getTLD(t *testing.T) {
	tests := []struct {
		name, separator, wantTLD string
	}{
		{"vitalik.eth", ".", ".eth"},
		{"vitalik.ens", ".", ".ens"},
		{"ourxyzwallet.xyz", ".", ".xyz"},
		{"Cameron.Kred", ".", ".kred"},
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
		checkGetTLD(t, tt.name, tt.separator, tt.wantTLD)
	}
}
