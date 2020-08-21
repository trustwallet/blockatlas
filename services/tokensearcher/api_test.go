package tokensearcher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getAddressesFromRequest(t *testing.T) {
	request := make(map[string][]string)
	request["60"] = []string{"1", "2", "3"}
	r := getAddressesFromRequest(request)
	assert.Equal(t, []string{"60_1", "60_2", "60_3"}, r)
}

func Test_getAddressesToRegisterByCoin(t *testing.T) {
	addressFromDB := make(map[string][]string)
	addressFromDB["60_a"] = []string{"1", "2", "3"}
	addressFromDB["714_b"] = []string{"1", "3"}

	addressesFromRequest := []string{"60_a", "714_b", "118_c"}
	result := getAddressesToRegisterByCoin(addressFromDB, addressesFromRequest)

	c, ok := result[118]
	assert.True(t, ok)
	assert.Equal(t, []string{"c"}, c)
}

func Test_getCoinIDFromAddress(t *testing.T) {
	address, id, ok := getCoinIDFromAddress("60_a")
	assert.Equal(t, "a", address)
	assert.Equal(t, uint(60), id)
	assert.True(t, ok)
}

func Test_getAssetsToResponse(t *testing.T) {
	addressFromDB := make(map[string][]string)
	addressFromDB["60_a"] = []string{"1", "2", "3"}
	addressFromDB["714_b"] = []string{"1", "3"}

	addressFromNodes := make(map[string][]string)
	addressFromNodes["118_c"] = []string{"1", "2", "3"}

	addressesFromRequest := []string{"60_a", "714_b", "118_c"}

	result := getAssetsToResponse(addressFromDB, addressFromNodes, addressesFromRequest)
	assert.NotNil(t, result)

	assert.Equal(t, []string{"1", "2", "3"}, result["60_a"])
	assert.Equal(t, []string{"1", "3"}, result["714_b"])
	assert.Equal(t, []string{"1", "2", "3"}, result["118_c"])
}
