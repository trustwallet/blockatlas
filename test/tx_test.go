package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"testing"
)

var transferDst = blockatlas.Tx{
	ID:     "1681EE543FB4B5A628EF21D746E031F018E226D127044A4F9BA5EE2542A44555",
	Coin:   coin.BNB,
	From:   "tbnb1fhr04azuhcj0dulm7ka40y0cqjlafwae9k9gk2",
	To:     "tbnb1sylyjw032eajr9cyllp26n04300qzzre38qyv5",
	Fee:    "125000",
	Date:   1555049867,
	Block:  7761368,
	Status: blockatlas.StatusCompleted,
	Memo:   "test",
	Meta: blockatlas.Transfer{
		Value:    "10000000000000",
		Decimals: 8,
		Symbol:   "BNB",
	},
}

var nativeTransferDst = blockatlas.Tx{
	ID:     "95CF63FAA27579A9B6AF84EF8B2DFEAC29627479E9C98E7F5AE4535E213FA4C9",
	Coin:   coin.BNB,
	From:   "tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a",
	To:     "tbnb12hlquylu78cjylk5zshxpdj6hf3t0tahwjt3ex",
	Fee:    "125000",
	Date:   1555117625,
	Block:  7928667,
	Status: blockatlas.StatusCompleted,
	Memo:   "test",
	Meta: blockatlas.NativeTokenTransfer{
		TokenID:  "YLC-D8B",
		Symbol:   "YLC",
		Value:    "210572645",
		Decimals: 8,
		From:     "tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a",
		To:       "tbnb12hlquylu78cjylk5zshxpdj6hf3t0tahwjt3ex",
	},
}

func TestTxSet_Add(t *testing.T) {
	set := blockatlas.TxSet{}
	set.Add(transferDst)
	var txs = set.Txs()
	assert.Equal(t, txs[0].ID, transferDst.ID)
	set.Add(transferDst)
	assert.Equal(t, set.Size(), 1)
	set.Add(nativeTransferDst)
	assert.Equal(t, set.Size(), 2)
}

func TestTx_GetAddresses(t *testing.T) {
	assert.Equal(t, transferDst.GetAddresses(), []string{"tbnb1fhr04azuhcj0dulm7ka40y0cqjlafwae9k9gk2", "tbnb1sylyjw032eajr9cyllp26n04300qzzre38qyv5"})
	assert.Equal(t, nativeTransferDst.GetAddresses(), []string{"tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a", "tbnb12hlquylu78cjylk5zshxpdj6hf3t0tahwjt3ex"})
}
