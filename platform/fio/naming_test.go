package fio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanHandle(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"vitalik@trust", true},
		{"vitalik@trustwallet", true},
		{"vitalik@binance", true},
		{"vitalik@fiomembers", true},
		{"vitalik@TRUST", true},
		{"vitalik@Trust", true},
		{"vitalik@somedomain", true},
		{"vitalik@x", true},
		{"v@trust", true},
		{"@trust", true},
		{"vitalik", false},
		{"vitalik@", false},
	}
	p := Init("")
	for _, tt := range tests {
		res := p.CanHandle(tt.name)
		assert.Equal(t, tt.want, res)
	}
}
