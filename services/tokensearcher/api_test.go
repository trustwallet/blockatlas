package tokensearcher

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

func Test_getAddressesFromRequest(t *testing.T) {
	request := Request{
		AddressesByCoin: make(map[string][]string),
		From:            0,
	}
	request.AddressesByCoin["60"] = []string{"1", "2", "3"}
	r := getAddressesFromRequest(request)
	assert.Equal(t, []string{"60_1", "60_2", "60_3"}, r)
}

func Test_getUnsubscribedAddresses(t *testing.T) {
	r := getUnsubscribedAddresses([]string{"60_1"}, []string{"60_1", "714_2", "714_22", "118_3"})
	assert.Equal(t, []string{"2", "22"}, r[714])
	assert.Equal(t, []string{"3"}, r[118])
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

func Test_getAssetsForAddressesFromNodes(t *testing.T) {
	apis := make(map[uint]blockatlas.TokensAPI)
	mock0 := mockedTokenAPI{WantedToken: "ABC", WantedCoin: 0}
	mock60 := mockedTokenAPI{WantedToken: "XYZ", WantedCoin: 60}
	apis[0] = mock0
	apis[60] = mock60

	addresses := make(map[uint][]string)
	addresses[0] = []string{"A", "B", "C"}
	addresses[60] = []string{"X", "Y", "Z"}
	result := getAssetsByAddressFromNodes(addresses, apis)
	assert.NotNil(t, result)
}

type mockedTokenAPI struct {
	WantedCoin  uint
	WantedToken string
}

func (m mockedTokenAPI) GetTokenListByAddress(address string) (blockatlas.TokenPage, error) {
	if address == "Y" {
		return nil, errors.New("failed")
	}
	tk := blockatlas.Token{
		Name:     "",
		Symbol:   "",
		Decimals: 0,
		TokenID:  m.WantedToken,
		Coin:     m.WantedCoin,
		Type:     "",
	}
	return blockatlas.TokenPage{tk}, nil
}
func (m mockedTokenAPI) Coin() coin.Coin {
	return coin.Coin{ID: m.WantedCoin}
}
