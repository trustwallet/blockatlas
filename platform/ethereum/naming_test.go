package ethereum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanHandle(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"vitalik.eth", true},
		{"vitalik.xyz", true},
		{"vitalik.luxe", true},
		{"vitalik.kred", true},
		{"vitalik.ETH", true},
		{"vitalik.Eth", true},
		{"vitalik.wrongdomain", false},
		{"v.eth", true},
		{".eth", true},
		{"vitalik", false},
		{"vitalik.", false},
	}
	p := Init(0, "", "")
	for _, tt := range tests {
		res := p.CanHandle(tt.name)
		assert.Equal(t, tt.want, res)
	}
}
