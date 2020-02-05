// +build integration

package ontology

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/platform/ontology"
	"testing"
)

var (
	testBlock = blockatlas.Block{
		Number: 7707834,
		Txs: []blockatlas.Tx{
			{
				From:   "AbEeCHUWpzQaxUN7G1a83N3P2XtVLuMLaE",
				To:     "ASLbwuar3ZTbUbLPnCgjGUw2WHhMfvJJtx",
				ID:     "266d9d7282a5601bf6cb8fc5368a76a2aa54f45731a063a699a692487bcbd0cb",
				Fee:    "10000000",
				Block:  7707834,
				Status: blockatlas.StatusCompleted,
				Date:   1580481541,
				Coin:   coin.Ontology().ID,
				Type:   blockatlas.TxNativeTokenTransfer,
				Meta: blockatlas.AnyAction{
					Name:     "Claim Rewards",
					Symbol:   "ONG",
					TokenID:  "ong",
					Decimals: 9,
					Value:    "51000000000000",
					Title:    blockatlas.AnyActionClaimRewards,
					Key:      blockatlas.KeyStakeClaimRewards,
				},
			}, {
				From:   "ANdrA47zDXUu8MCkMdD3FYPmpSNGYeAvKz",
				To:     "ASLbwuar3ZTbUbLPnCgjGUw2WHhMfvJJtx",
				ID:     "2935268c5715f1f2015ba828681c39399dedbe7a24ed628ef7b85d9aac8045fd",
				Fee:    "10000000",
				Block:  7707834,
				Status: blockatlas.StatusCompleted,
				Date:   1580481541,
				Coin:   coin.Ontology().ID,
				Type:   blockatlas.TxNativeTokenTransfer,
				Meta: blockatlas.AnyAction{
					Name:     "Claim Rewards",
					Symbol:   "ONG",
					TokenID:  "ong",
					Decimals: 9,
					Value:    "113200000000",
					Title:    blockatlas.AnyActionClaimRewards,
					Key:      blockatlas.KeyStakeClaimRewards,
				},
			}, {
				From:   "Abg2gs6pfpQu82jXbm8EYGiipRBvf9ktVS",
				To:     "ASLbwuar3ZTbUbLPnCgjGUw2WHhMfvJJtx",
				ID:     "40976edc1306b0e5f55b90c8d3ca248bb544e5ebbadb02be6146ba0a0de402c3",
				Fee:    "10000000",
				Block:  7707834,
				Status: blockatlas.StatusCompleted,
				Date:   1580481541,
				Coin:   coin.Ontology().ID,
				Type:   blockatlas.TxNativeTokenTransfer,
				Meta: blockatlas.AnyAction{
					Name:     "Claim Rewards",
					Symbol:   "ONG",
					TokenID:  "ong",
					Decimals: 9,
					Value:    "10949000000000",
					Title:    blockatlas.AnyActionClaimRewards,
					Key:      blockatlas.KeyStakeClaimRewards,
				},
			},
		},
	}
)

const (
	blockNum = 7707834
)

func TestOntology(t *testing.T) {
	t.Run("test ontology", func(t *testing.T) {
		p := &ontology.Platform{}
		_ = p.Init()
		testCurrentBlockNumber(p, t)
		testGetBlockByNumber(p, t)
	})
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

	isSame := resp.Number == testBlock.Number &&
		resp.Txs[0].Block == testBlock.Txs[0].Block &&
		resp.Txs[0].Status == testBlock.Txs[0].Status &&
		resp.Txs[0].Date == testBlock.Txs[0].Date &&
		resp.Txs[0].Coin == testBlock.Txs[0].Coin
	if !isSame {
		t.Errorf("Block is not the same")
	}

	// check that we have tx hashes of parsed block
	txMap := map[string]bool{}
	for _, tx := range resp.Txs {
		txMap[tx.ID] = true
	}
	delete(txMap, "40976edc1306b0e5f55b90c8d3ca248bb544e5ebbadb02be6146ba0a0de402c3")
	delete(txMap, "266d9d7282a5601bf6cb8fc5368a76a2aa54f45731a063a699a692487bcbd0cb")
	delete(txMap, "2935268c5715f1f2015ba828681c39399dedbe7a24ed628ef7b85d9aac8045fd")

	if len(txMap) > 0 {
		t.Errorf("Block is not the same: %v", txMap)
	}
}
