package redis

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"

	"testing"
)

const ethCoin = coin.ETH
const addr1 = "0xde0B295669a9FD93d5F28D9Ec85E40f4cb697BAe"

var webhook1Addr = "http://apple.com/push"
var webhook1 = []string{webhook1Addr}

const addr2 = "0xEA674fdDe714fd979de3EdF0F56AA9716B898ec8"

var webhook2Addr = "http://trustwallet.com/webhook"
var webhook2 = []string{webhook2Addr}

func TestRedisStorage_Add(t *testing.T) {
	result := add([]string{webhook1Addr}, []string{webhook2Addr})

	assert.Equal(t, result, []string{webhook1Addr, webhook2Addr})

	result = add(nil, []string{webhook2Addr})

	assert.Equal(t, result, []string{webhook2Addr})

	result = add([]string{webhook1Addr}, []string{webhook1Addr})

	assert.Equal(t, result, []string{webhook1Addr})

	result = add([]string{webhook2Addr}, nil)

	assert.Equal(t, result, []string{webhook2Addr})
}
