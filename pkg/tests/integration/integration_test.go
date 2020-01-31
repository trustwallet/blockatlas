// +build integration

package integration

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/config"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/platform/ontology"
	"os"
	"testing"
)

var (
	testBlock = blockatlas.Block{
		Number: 7677564,
		ID:     "168d35ae9333f1d53ee0c124b44d268701df001df1313b388d001a5808f66d01",
		Txs: []blockatlas.Tx{
			{
				ID:     "736fab4fa13435f201bc90a43ca5cd8c324ec88d6048fedb136f267371daee39",
				Block:  7677564,
				Status: blockatlas.StatusCompleted,
				Date:   1580115134,
				Coin:   coin.Ontology().ID,
			},
		},
	}
	blockNum = 7677564
)

func TestOntology(t *testing.T) {
	config.LoadConfig(os.Getenv("TEST_CONFIG"))
	p := &ontology.Platform{}
	_ = p.Init()
	testCurrentBlockNumber(p, t)
	testGetBlockByNumber(p, t)
}

func testCurrentBlockNumber(p *ontology.Platform, t *testing.T) {
	resp, err := p.CurrentBlockNumber()
	if err != nil {
		t.Error(err)
	}
	if resp < 0 {
		t.Error("block is < 0")
	}
}

func testGetBlockByNumber(p *ontology.Platform, t *testing.T) {
	resp, err := p.GetBlockByNumber(int64(blockNum))
	if err != nil {
		t.Error(err)
	}

	isSame := resp.ID == testBlock.ID &&
		resp.Number == testBlock.Number &&
		resp.Txs[0].ID == testBlock.Txs[0].ID &&
		resp.Txs[0].Block == testBlock.Txs[0].Block &&
		resp.Txs[0].Status == testBlock.Txs[0].Status &&
		resp.Txs[0].Date == testBlock.Txs[0].Date &&
		resp.Txs[0].Coin == testBlock.Txs[0].Coin

	if !isSame {
		t.Error("Block is not the same")
	}
}
