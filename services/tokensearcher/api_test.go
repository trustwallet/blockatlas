package tokensearcher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
