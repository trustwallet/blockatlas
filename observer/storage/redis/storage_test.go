package redis

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var webhook1Addr = "http://apple.com/push"

var webhook2Addr = "http://trustwallet.com/webhook"

func TestRedisStorage_WebhookAdd(t *testing.T) {
	result := add([]string{webhook1Addr}, []string{webhook2Addr})

	assert.Equal(t, result, []string{webhook1Addr, webhook2Addr})

	result = add(nil, []string{webhook2Addr})

	assert.Equal(t, result, []string{webhook2Addr})

	result = add([]string{webhook1Addr}, []string{webhook1Addr})

	assert.Equal(t, result, []string{webhook1Addr})

	result = add([]string{webhook2Addr}, nil)

	assert.Equal(t, result, []string{webhook2Addr})
}

func TestRedisStorage_WebhookDelete(t *testing.T) {
	result := remove([]string{webhook1Addr, webhook2Addr}, []string{webhook2Addr})

	assert.Equal(t, result, []string{webhook1Addr})

	result = remove([]string{webhook2Addr}, nil)

	assert.Equal(t, result, []string{webhook2Addr})

	result = remove([]string{webhook1Addr}, []string{webhook1Addr})

	assert.Equal(t, result, []string{})

	result = remove([]string{webhook2Addr}, []string{})

	assert.Equal(t, result, []string{webhook2Addr})
}
