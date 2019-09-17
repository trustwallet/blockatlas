package ethereum

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetValidParameter(t *testing.T) {
	var tests = []struct {
		first  string
		second string
		result string
	}{
		{"trust", "wallet", "trust"},
		{"", "wallet", "wallet"},
		{"trust", "", "trust"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("getValidParameter %d", i), func(t *testing.T) {
			s := getValidParameter(tt.first, tt.second)
			if s != tt.result {
				t.Errorf("got %q, want %q", s, tt.result)
			}
		})
	}
}

func TestCreateCollectionId(t *testing.T) {
	var tests = []struct {
		address string
		slug    string
		result  string
	}{
		{"0x5574Cd97", "trust", "0x5574Cd97---trust"},
		{"0x5574Cd97", "", "0x5574Cd97---"},
		{"", "trust", "---trust"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("createCollectionId %d", i), func(t *testing.T) {
			s := createCollectionId(tt.address, tt.slug)
			if s != tt.result {
				t.Errorf("got %q, want %q", s, tt.result)
			}
		})
	}
}

func TestGetCollectionId(t *testing.T) {
	var tests = []struct {
		collectionId string
		result       string
	}{
		{"0x5574Cd97---trust", "trust"},
		{"0x5574Cd97---", ""},
		{"---", ""},
		{"trust", "trust"},
		{"---trust", "trust"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("getCollectionId %d", i), func(t *testing.T) {
			s := getCollectionId(tt.collectionId)
			if s != tt.result {
				t.Errorf("got %q, want %q", s, tt.result)
			}
		})
	}
}

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
		{[]Collection{c1, c2, c3}, "0xfaafdc07907ff5120a76b34b731b278c38d6043c", &c1},
		{[]Collection{c1, c2, c3}, "cryptokitties", &c2},
		{[]Collection{c1, c2}, "age-of-rust", nil},
		{[]Collection{c1, c2, c3}, "age-of-rust", &c3},
		{[]Collection{c1, c2}, "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB", &c2},
		{[]Collection{c1}, "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB", nil},
		{[]Collection{c1, c3}, "enjin", &c1},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("searchCollection %d", i), func(t *testing.T) {
			s := searchCollection(tt.collections, tt.collectibleID)
			assert.EqualValues(t, s, tt.result)
		})
	}

}
