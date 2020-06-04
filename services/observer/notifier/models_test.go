package notifier

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"sort"
	"testing"
)

var (
	nativeTokenTransfer = blockatlas.Tx{
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
	tokenTransfer = blockatlas.Tx{
		ID:       "0xbcd1a43e796de4035e5e2991d8db332958e36031d54cb1d3a08d2cb790e338c4",
		Coin:     60,
		From:     "0x08777CB1e80F45642752662B04886Df2d271E049",
		To:       "0xdd974D5C2e2928deA5F71b9825b8b646686BD200",
		Fee:      "52473000000000",
		Date:     1585169424,
		Block:    9742705,
		Status:   "completed",
		Sequence: 149,
		Type:     "token_transfer",
		Meta: blockatlas.TokenTransfer{
			Name:     "Kyber Network Crystal",
			Symbol:   "KNC",
			TokenID:  "0xdd974D5C2e2928deA5F71b9825b8b646686BD200",
			Decimals: 18,
			Value:    "100000000000000",
			From:     "0x08777CB1e80F45642752662B04886Df2d271E049",
			To:       "0x38d45371993eEc84f38FEDf93C646aA2D2267CEA",
		},
	}
	transfer = blockatlas.Tx{
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
	utxoTransfer = blockatlas.Tx{
		ID:   "zpub6ruK9k6YGm8BRHWvTiQcrEPnFkuRDJhR7mPYzV2LDvjpLa5CuGgrhCYVZjMGcLcFqv9b2WvsFtY2Gb3xq8NVq8qhk9veozrA2W9QaWtihrC",
		Coin: coin.BTC,
		Inputs: []blockatlas.TxOutput{
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
		Outputs: []blockatlas.TxOutput{
			{
				Address: "bc1qjcslq88cht8llqmh3aqshjx9we9msv386jvxl6",
				Value:   "3",
			},
		},
		From:   "bc1qhn03cww757mnnlpkdvvfkaydxqygm86nvkm92h",
		To:     "bc1qjcslq88cht8llqmh3aqshjx9we9msv386jvxl6",
		Fee:    "125000",
		Date:   1555117625,
		Block:  592400,
		Status: blockatlas.StatusCompleted,
		Memo:   "test",
		Meta: blockatlas.Transfer{
			Value:    "10000000000000",
			Decimals: 8,
			Symbol:   "BNB",
		},
	}
)

func Test_containsAddress(t *testing.T) {
	assert.True(t, containsAddress(tokenTransfer, "0x08777CB1e80F45642752662B04886Df2d271E049"))
	assert.False(t, containsAddress(tokenTransfer, "0xdd974D5C2e2928deA5F71b9825b8b646686BD200"))
	assert.True(t, containsAddress(tokenTransfer, "0x38d45371993eEc84f38FEDf93C646aA2D2267CEA"))
	assert.False(t, containsAddress(tokenTransfer, "0xdd974D5C2e2928deA5F71b9825b8b646686BD200"))

	assert.True(t, containsAddress(transfer, "tbnb1fhr04azuhcj0dulm7ka40y0cqjlafwae9k9gk2"))
	assert.False(t, containsAddress(transfer, "1681EE543FB4B5A628EF21D746E031F018E226D127044A4F9BA5EE2542A44556"))
	assert.True(t, containsAddress(transfer, "tbnb1fhr04azuhcj0dulm7ka40y0cqjlafwae9k9gk2"))

	assert.True(t, containsAddress(nativeTokenTransfer, "tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a"))
	assert.False(t, containsAddress(nativeTokenTransfer, "95CF63FAA27579A9B6AF84EF8B2DFEAC29627479E9C98E7F5AE4535E213FA4C9"))
	assert.True(t, containsAddress(nativeTokenTransfer, "tbnb12hlquylu78cjylk5zshxpdj6hf3t0tahwjt3ex"))

	assert.True(t, containsAddress(utxoTransfer, "bc1qhn03cww757mnnlpkdvvfkaydxqygm86nvkm92h"))
	assert.False(t, containsAddress(utxoTransfer, "bc1qc7ekqf2t0elfsmtgr2mgd7da2up4vgq8uqk2nh"))
	assert.False(t, containsAddress(utxoTransfer, "bc1qv454wacvnenr3hzzldjqn8cgfltdlxwe96h737"))
	assert.False(t, containsAddress(utxoTransfer, "zpub6ruK9k6YGm8BRHWvTiQcrEPnFkuRDJhR7mPYzV2LDvjpLa5CuGgrhCYVZjMGcLcFqv9b2WvsFtY2Gb3xq8NVq8qhk9veozrA2W9QaWtihrC"))
	assert.True(t, containsAddress(utxoTransfer, "bc1qjcslq88cht8llqmh3aqshjx9we9msv386jvxl6"))
}

func Test_findTransactionsByAddress(t *testing.T) {
	res := findTransactionsByAddress([]blockatlas.Tx{nativeTokenTransfer, tokenTransfer}, "tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a")
	sort.Slice(res, func(i, j int) bool {
		return res[i].ID < res[j].ID
	})
	assert.Equal(t, []blockatlas.Tx{nativeTokenTransfer}, res)

	resFail := findTransactionsByAddress([]blockatlas.Tx{nativeTokenTransfer, tokenTransfer}, "tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced")
	assert.Equal(t, []blockatlas.Tx{}, resFail)
}

func Test_buildNotificationsByAddress(t *testing.T) {
	notifications := buildNotificationsByAddress("tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a", []blockatlas.Tx{nativeTokenTransfer, tokenTransfer}, context.Background())
	sort.Slice(notifications, func(i, j int) bool {
		return notifications[i].Action < notifications[j].Action
	})
	nativeTokenTransfer.Direction = blockatlas.DirectionOutgoing
	assert.Equal(t, nativeTokenTransfer, notifications[0].Result)
}
