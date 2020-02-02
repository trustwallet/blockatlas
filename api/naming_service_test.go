package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func checkGetTLD(t *testing.T, name string, expectedTLD string, expectedOk bool) {
	tld, ok := getTLD(name)
	assert.Equal(t, expectedTLD, tld)
	assert.Equal(t, expectedOk, ok)
}

func Test_getTLD(t *testing.T) {
	checkGetTLD(t, "vitalik.eth", ".eth", true)
	checkGetTLD(t, "vitalik.ens", ".ens", true)
	checkGetTLD(t, "ourxyzwallet.xyz", ".xyz", true)
	checkGetTLD(t, "btc.zil", ".zil", true)
	checkGetTLD(t, "btc.crypto", ".crypto", true)
	checkGetTLD(t, "nick@fiotestnet", "@fiotestnet", true)
	checkGetTLD(t, "a", "", false) // no tld
	checkGetTLD(t, "a@b.c", ".c", true)
	checkGetTLD(t, "a.b@c", ".b@c", true)
}
