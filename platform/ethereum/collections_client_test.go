package ethereum

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var c1 = Collection{
	Slug: "enjin",
	Contracts: []PrimaryAssetContract{{
		Address: "0xfaafdc07907ff5120a76b34b731b278c38d6043c",
	}},
}
var c2 = Collection{
	Slug: "cryptokitties",
	Contracts: []PrimaryAssetContract{{
		Address: "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
	}},
}
var c3 = Collection{
	Slug: "age-of-rust",
	Contracts: []PrimaryAssetContract{{
		Address: "0x0875BCab22dE3d02402bc38aEe4104e1239374a7",
	}},
}

func TestSearchCollection(t *testing.T) {
	var tests = []struct {
		collections   []Collection
		collectibleID string
		result        *Collection
	}{
		{[]Collection{c1, c2, c3}, "enjin", &c1},
		{[]Collection{c1, c2, c3}, "cryptokitties", &c2},
		{[]Collection{c1, c2}, "age-of-rust", nil},
		{[]Collection{c1, c2, c3}, "age-of-rust", &c3},
		{[]Collection{c1, c2}, "cryptokitties", &c2},
		{[]Collection{c1}, "age-of-rust", nil},
		{[]Collection{c1, c3}, "enjin", &c1},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("searchCollection %d", i), func(t *testing.T) {
			s := searchCollection(tt.collections, tt.collectibleID)
			assert.EqualValues(t, s, tt.result)
		})
	}

}
