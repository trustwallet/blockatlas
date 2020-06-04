package zilliqa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatch(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"vitalik.zil", true},
		{"vitalik.crypto", true},
		{"vitalik.ZIL", true},
		{"vitalik.Zil", true},
		{"vitalik.wrongdomain", false},
		{"v.zil", true},
		{".zil", true},
		{"vitalik", false},
		{"vitalik.", false},
	}
	p := Init("", "", "", "")
	for _, tt := range tests {
		res := p.Match(tt.name)
		assert.Equal(t, tt.want, res)
	}
}
