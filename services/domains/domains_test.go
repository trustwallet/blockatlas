package domains

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pkg/errors"
)

func checkGetTLD(t *testing.T, name string, expectedTLD string, expectedError error) {
	name = strings.ToLower(name)
	tld, err := getTLD(name)
	assert.Equal(t, expectedTLD, tld)
	if expectedError == nil {
		assert.Nil(t, err)
	} else {
		assert.NotNil(t, err)
	}
}

func Test_getTLD(t *testing.T) {
	checkGetTLD(t, "vitalik.eth", ".eth", nil)
	checkGetTLD(t, "vitalik.ens", ".ens", nil)
	checkGetTLD(t, "ourxyzwallet.xyz", ".xyz", nil)
	checkGetTLD(t, "Cameron.Kred", ".kred", nil)
	checkGetTLD(t, "btc.zil", ".zil", nil)
	checkGetTLD(t, "btc.crypto", ".crypto", nil)
	checkGetTLD(t, "nick@fiotestnet", "@fiotestnet", nil)
	checkGetTLD(t, "a", "", errors.E("No TLD found in name"))  // no tld
	checkGetTLD(t, "a.", "", errors.E("No TLD found in name")) // empty tld
	checkGetTLD(t, "a@b.c", ".c", nil)
	checkGetTLD(t, "a.b@c", "@c", nil)
}
