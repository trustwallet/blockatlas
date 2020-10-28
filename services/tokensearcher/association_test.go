package tokensearcher

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

func Test_assetsMap(t *testing.T) {
	tx1 := blockatlas.Tx{
		Coin: 60,
		From: "A",
		To:   "B",
		Meta: blockatlas.NativeTokenTransfer{
			TokenID: "ABC",
			From:    "A",
			To:      "C",
		},
	}

	tx2 := blockatlas.Tx{
		Coin: 60,
		From: "D",
		To:   "V",
		Meta: blockatlas.TokenTransfer{
			TokenID: "EFG",
			From:    "D",
			To:      "F",
		},
	}

	tx3 := blockatlas.Tx{
		Coin: 60,
		From: "Q",
		To:   "L",
		Meta: blockatlas.AnyAction{
			TokenID: "HIJ",
		},
	}

	result := assetsMap(blockatlas.Txs{tx1, tx2, tx3}, "60")
	assert.Equal(t, result["60_A"], []models.Asset{{Asset: "c60_tABC", Type: "ERC20", Coin: 60}})
	assert.Equal(t, result["60_C"], []models.Asset{{Asset: "c60_tABC", Type: "ERC20", Coin: 60}})
	assert.Equal(t, result["60_D"], []models.Asset{{Asset: "c60_tEFG", Type: "ERC20", Coin: 60}})
	assert.Equal(t, result["60_F"], []models.Asset{{Asset: "c60_tEFG", Type: "ERC20", Coin: 60}})
	assert.Equal(t, result["60_Q"], []models.Asset{{Asset: "c60_tHIJ", Type: "ERC20", Coin: 60}})
	assert.Equal(t, result["60_L"], []models.Asset{{Asset: "c60_tHIJ", Type: "ERC20", Coin: 60}})
}

func Test_associationsToAdd(t *testing.T) {
	o := make(map[string][]models.Asset)
	n := make(map[string][]models.Asset)

	o["A"] = []models.Asset{{Asset: "1"}, {Asset: "2"}, {Asset: "3"}}
	o["B"] = []models.Asset{{Asset: "3"}, {Asset: "4"}, {Asset: "5"}}

	n["A"] = []models.Asset{{Asset: "1"}, {Asset: "2"}, {Asset: "5"}}
	n["B"] = []models.Asset{{Asset: "3"}, {Asset: "9"}, {Asset: "8"}}

	result := associationsToAdd(o, n)

	assert.Equal(t, result["A"], []models.Asset{{Asset: "5"}})
	assert.Equal(t, result["B"], []models.Asset{{Asset: "9"}, {Asset: "8"}})
}

func Test_newAssociationsForAddress(t *testing.T) {
	o := []models.Asset{{Asset: "1"}, {Asset: "2"}, {Asset: "3"}}
	n := []models.Asset{{Asset: "1"}, {Asset: "2"}, {Asset: "3"}, {Asset: "4"}, {Asset: "5"}}

	result := newAssociationsForAddress(o, n)
	assert.Equal(t, result, []models.Asset{{Asset: "4"}, {Asset: "5"}})

	o = []models.Asset{{Asset: "1"}, {Asset: "2"}, {Asset: "3"}}
	n = []models.Asset{{Asset: "1"}, {Asset: "2"}, {Asset: "3"}}

	result = newAssociationsForAddress(o, n)
	assert.Equal(t, len(result), len([]string{}))

	o = []models.Asset{{Asset: "1"}, {Asset: "2"}, {Asset: "3"}}
	n = []models.Asset{{Asset: "1"}, {Asset: "2"}}

	result = newAssociationsForAddress(o, n)
	assert.Equal(t, len(result), len([]string{}))
}

func Test_fromModelToAssociation(t *testing.T) {
	a := []models.AddressToAssetAssociation{
		{Address: models.Address{Address: "A"}, Asset: models.Asset{Asset: "1"}},
		{Address: models.Address{Address: "A"}, Asset: models.Asset{Asset: "2"}},
		{Address: models.Address{Address: "A"}, Asset: models.Asset{Asset: "3"}},
		{Address: models.Address{Address: "B"}, Asset: models.Asset{Asset: "2"}},
		{Address: models.Address{Address: "B"}, Asset: models.Asset{Asset: "3"}},
		{Address: models.Address{Address: "B"}, Asset: models.Asset{Asset: "4"}},
	}

	result := fromModelToAssociation(a)
	assert.Equal(t, result["A"], []models.Asset{{Asset: "1"}, {Asset: "2"}, {Asset: "3"}})
	assert.Equal(t, result["B"], []models.Asset{{Asset: "2"}, {Asset: "3"}, {Asset: "4"}})
}
