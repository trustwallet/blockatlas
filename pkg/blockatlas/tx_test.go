package blockatlas

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"testing"
)

var transferDst1 = Tx{
	ID:     "1681EE543FB4B5A628EF21D746E031F018E226D127044A4F9BA5EE2542A44555",
	Coin:   coin.BNB,
	From:   "tbnb1fhr04azuhcj0dulm7ka40y0cqjlafwae9k9gk2",
	To:     "tbnb1sylyjw032eajr9cyllp26n04300qzzre38qyv5",
	Fee:    "125000",
	Date:   1555049867,
	Block:  7761368,
	Status: StatusCompleted,
	Memo:   "test",
	Meta: Transfer{
		Value:    "10000000000000",
		Decimals: 8,
		Symbol:   "BNB",
	},
}

var nativeTransferDst1 = Tx{
	ID:     "95CF63FAA27579A9B6AF84EF8B2DFEAC29627479E9C98E7F5AE4535E213FA4C9",
	Coin:   coin.BNB,
	From:   "tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a",
	To:     "tbnb12hlquylu78cjylk5zshxpdj6hf3t0tahwjt3ex",
	Fee:    "125000",
	Date:   1555117625,
	Block:  7928667,
	Status: StatusCompleted,
	Memo:   "test",
	Meta: NativeTokenTransfer{
		TokenID:  "YLC-D8B",
		Symbol:   "YLC",
		Value:    "210572645",
		Decimals: 8,
		From:     "tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a",
		To:       "tbnb12hlquylu78cjylk5zshxpdj6hf3t0tahwjt3ex",
	},
}

var utxoTransferDst1 = Tx{
	ID:   "zpub6ruK9k6YGm8BRHWvTiQcrEPnFkuRDJhR7mPYzV2LDvjpLa5CuGgrhCYVZjMGcLcFqv9b2WvsFtY2Gb3xq8NVq8qhk9veozrA2W9QaWtihrC",
	Coin: coin.BTC,
	Inputs: []TxOutput{
		{
			Address: "bc1qhn03cww757mnnlpkdvvfkaydxqygm86nvkm92h",
			Value:   "1",
		},
		{
			Address: "bc1qc7ekqf2t0elfsmtgr2mgd7da2up4vgq8uqk2nh",
			Value:   "1",
		},
		{
			Address: "bc1qv454wacvnenr3hzzldjqn8cgfltdlxwe96h737",
			Value:   "1",
		},
	},
	Outputs: []TxOutput{
		{
			Address: "bc1qjcslq88cht8llqmh3aqshjx9we9msv386jvxl6",
			Value:   "3",
		},
	},
	Fee:    "125000",
	Date:   1555117625,
	Block:  592400,
	Status: StatusCompleted,
	Memo:   "test",
}

var utxoTransferDst2 = Tx{
	ID:   "zpub6ruK9k6YGm8BRHWvTiQcrEPnFkuRDJhR7mPYzV2LDvjpLa5CuGgrhCYVZjMGcLcFqv9b2WvsFtY2Gb3xq8NVq8qhk9veozrA2W9QaWtihrC",
	Coin: coin.BTC,
	Inputs: []TxOutput{
		{
			Address: "bc1q6e8sdxlgc7ekqkqyevtrx8wshfv7sg66z3z6ce",
			Value:   "4",
		},
		{
			Address: "bc1q7nn4txus4g6fc5v7d2tha35ely8mfpd8qvv6eg",
			Value:   "2",
		},
	},
	Outputs: []TxOutput{
		{
			Address: "bc1q2fpry7zwqh575huc9urwfdvjtuvz508wez56ff",
			Value:   "3",
		},
		{
			Address: "bc1qk3yj6h79qw7tnsg4durc9sd5fpd3qt0p0m8u5p",
			Value:   "1",
		},
		{
			Address: "bc1qm8836plkzft2rhh23z6j8s9s8fxrzd4zag95z8",
			Value:   "2",
		},
	},
	Fee:    "125000",
	Date:   1555117625,
	Block:  592400,
	Status: StatusCompleted,
	Memo:   "test",
}

func TestTxSet_Add(t *testing.T) {
	set := TxSet{}
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

func TestTx_GetUtxoAddresses(t *testing.T) {
	assert.Equal(t, utxoTransferDst1.GetUtxoAddresses(), []string{"bc1qhn03cww757mnnlpkdvvfkaydxqygm86nvkm92h", "bc1qc7ekqf2t0elfsmtgr2mgd7da2up4vgq8uqk2nh", "bc1qv454wacvnenr3hzzldjqn8cgfltdlxwe96h737", "bc1qjcslq88cht8llqmh3aqshjx9we9msv386jvxl6"})
	assert.Equal(t, utxoTransferDst2.GetUtxoAddresses(), []string{"bc1q6e8sdxlgc7ekqkqyevtrx8wshfv7sg66z3z6ce", "bc1q7nn4txus4g6fc5v7d2tha35ely8mfpd8qvv6eg", "bc1q2fpry7zwqh575huc9urwfdvjtuvz508wez56ff", "bc1qk3yj6h79qw7tnsg4durc9sd5fpd3qt0p0m8u5p", "bc1qm8836plkzft2rhh23z6j8s9s8fxrzd4zag95z8"})
}
