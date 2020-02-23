// +build integration

package bitcoin

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/platform"
	"github.com/trustwallet/blockatlas/platform/bitcoin"
	"testing"
)

const (
	blockNum = 616070
)

func TestBitcoin(t *testing.T) {
	t.Run("test bitcoin", func(t *testing.T) {
		p := bitcoin.Init(coin.BTC, platform.GetApiVar(coin.BTC))
		testCurrentBlockNumber(p, t)
		testGetBlockByNumber(p, t)
	})
}

func testCurrentBlockNumber(p *bitcoin.Platform, t *testing.T) {
	resp, err := p.CurrentBlockNumber()
	if err != nil {
		t.Error(err)
	}
	if resp < 0 {
		t.Error("block is < 0")
	}
}

func testGetBlockByNumber(p *bitcoin.Platform, t *testing.T) {
	resp, err := p.GetBlockByNumber(int64(blockNum))
	if err != nil {
		t.Error(err)
	}

	expectedTxs := 2129
	if len(resp.Txs) != expectedTxs {
		t.Errorf("invalid tx count for block %d, expected %d", len(resp.Txs), expectedTxs)
	}
}
