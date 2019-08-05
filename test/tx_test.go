package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/observer"
	"testing"
)

var transferDst1 = blockatlas.Tx{
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

var transferDst2 = blockatlas.Tx{
	ID:     "1681EE543FB4B5A628EF21D746E031F018E226D127044A4F9BA5EE2542A44556",
	Coin:   coin.BNB,
	From:   "tbnb1fhr04azuhcj0dulm7ka40y0cqjlafwae9k9gk2",
	To:     "tbnb1fhr04azuhcj0dulm7ka40y0cqjlafwae9k9gk2",
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

var nativeTransferDst1 = blockatlas.Tx{
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

var nativeTransferDst2 = blockatlas.Tx{
	ID:     "95CF63FAA27579A9B6AF84EF8B2DFEAC29627479E9C98E7F5AE4535E213FA4D0",
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

var txsBlock = blockatlas.Block{
	Number: 12345,
	ID:     "12345",
	Txs: []blockatlas.Tx{
		transferDst1,
		transferDst2,
		nativeTransferDst1,
		nativeTransferDst2,
	},
}

func TestGetTxs(t *testing.T) {
	txs := observer.GetTxs(&txsBlock)
	assert.Equal(t, len(txs), 4)
	assert.Equal(t, txs["tbnb1fhr04azuhcj0dulm7ka40y0cqjlafwae9k9gk2"].Size(), 2)
	assert.Equal(t, txs["tbnb1sylyjw032eajr9cyllp26n04300qzzre38qyv5"].Size(), 1)
	assert.Equal(t, txs["tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a"].Size(), 2)
	assert.Equal(t, txs["tbnb12hlquylu78cjylk5zshxpdj6hf3t0tahwjt3ex"].Size(), 2)
}

func TestTxSet_Add(t *testing.T) {
	set := blockatlas.TxSet{}
	set.Add(&transferDst1)
	var txs = set.Txs()
	assert.Equal(t, txs[0].ID, transferDst1.ID)
	set.Add(&transferDst1)
	assert.Equal(t, set.Size(), 1)
	set.Add(&nativeTransferDst1)
	assert.Equal(t, set.Size(), 2)
}

func TestTx_GetAddresses(t *testing.T) {
	assert.Equal(t, transferDst1.GetAddresses(), []string{"tbnb1fhr04azuhcj0dulm7ka40y0cqjlafwae9k9gk2", "tbnb1sylyjw032eajr9cyllp26n04300qzzre38qyv5"})
	assert.Equal(t, nativeTransferDst1.GetAddresses(), []string{"tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a", "tbnb12hlquylu78cjylk5zshxpdj6hf3t0tahwjt3ex"})
}
