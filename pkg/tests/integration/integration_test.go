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
		Number: 7707834,
		ID:     "a5f3ee1a102df7196bb1e262a05435f260392fae6be676ae2c0a6147f8ecf94c",
		Txs: []blockatlas.Tx{
			{
				ID:     "266d9d7282a5601bf6cb8fc5368a76a2aa54f45731a063a699a692487bcbd0cb",
				Block:  7707834,
				Status: blockatlas.StatusCompleted,
				Date:   1580481541,
				Coin:   coin.Ontology().ID,
				Meta: blockatlas.NativeTokenTransfer{
					Name:     "Ontology Gas",
					Symbol:   "ONG",
					TokenID:  "ong",
					Decimals: 9,
					Value:    "51000000000000",
					From:     "AbEeCHUWpzQaxUN7G1a83N3P2XtVLuMLaE",
					To:       "ASLbwuar3ZTbUbLPnCgjGUw2WHhMfvJJtx",
				},
			},
		},
	}
	blockNum = 7707834
)

func TestOntology(t *testing.T) {
	configPath := os.Getenv("TEST_CONFIG")
	if configPath == "" {
		config.LoadConfig("../../../config.yml")
	} else {
		config.LoadConfig(configPath)
	}
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
		resp.Txs[0].Block == testBlock.Txs[0].Block &&
		resp.Txs[0].Status == testBlock.Txs[0].Status &&
		resp.Txs[0].Date == testBlock.Txs[0].Date &&
		resp.Txs[0].Coin == testBlock.Txs[0].Coin

	if isSame {
		// check that we have tx hashes of parsed block
		for _, tx := range resp.Txs {
			switch tx.ID {
			case "40976edc1306b0e5f55b90c8d3ca248bb544e5ebbadb02be6146ba0a0de402c3":
				isSame = true
			case "266d9d7282a5601bf6cb8fc5368a76a2aa54f45731a063a699a692487bcbd0cb":
				isSame = true
			case "2935268c5715f1f2015ba828681c39399dedbe7a24ed628ef7b85d9aac8045fd":
				isSame = true
			default:
				isSame = false
			}
		}
	}
	if !isSame {
		t.Error("Block is not the same")
	}
}
