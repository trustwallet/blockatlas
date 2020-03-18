package blockatlas

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"testing"
)

var (
	transferDstOne = Tx{
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

	transferDst2 = Tx{
		ID:     "1681EE543FB4B5A628EF21D746E031F018E226D127044A4F9BA5EE2542A44556",
		Coin:   coin.BNB,
		From:   "tbnb1fhr04azuhcj0dulm7ka40y0cqjlafwae9k9gk2",
		To:     "tbnb1fhr04azuhcj0dulm7ka40y0cqjlafwae9k9gk2",
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

	nativeTransferDstOne = Tx{
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

	nativeTransferDst2 = Tx{
		ID:     "95CF63FAA27579A9B6AF84EF8B2DFEAC29627479E9C98E7F5AE4535E213FA4D0",
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
	txsBlock = Block{
		Number: 12345,
		ID:     "12345",
		Txs: []Tx{
			transferDstOne,
			transferDst2,
			nativeTransferDstOne,
			nativeTransferDst2,
		},
	}
)

func TestGetTxs(t *testing.T) {
	txs := txsBlock.GetTransactionsMap()
	assert.Equal(t, len(txs), 4)
	assert.Equal(t, txs["tbnb1fhr04azuhcj0dulm7ka40y0cqjlafwae9k9gk2"].Size(), 2)
	assert.Equal(t, txs["tbnb1sylyjw032eajr9cyllp26n04300qzzre38qyv5"].Size(), 1)
	assert.Equal(t, txs["tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a"].Size(), 2)
	assert.Equal(t, txs["tbnb12hlquylu78cjylk5zshxpdj6hf3t0tahwjt3ex"].Size(), 2)
}
