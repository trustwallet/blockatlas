package ontology

import (
	"fmt"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Client struct {
	blockatlas.Request
}

type blockListResponse struct {
	Action  string            `json:"Action"`
	Error   int               `json:"Error"`
	Desc    string            `json:"Desc"`
	Version string            `json:"Version"`
	Result  []blockListResult `json:"Result"`
}
type blockListResult struct {
	PrevBlock     string `json:"PrevBlock"`
	TxnsRoot      string `json:"TxnsRoot"`
	BlockTime     int    `json:"BlockTime"`
	NextBlock     string `json:"NextBlock"`
	BookKeeper    string `json:"BookKeeper"`
	TxnNum        int    `json:"TxnNum"`
	Height        int    `json:"Height"`
	Hash          string `json:"Hash"`
	ConsensusData string `json:"ConsensusData"`
	BlockSize     int    `json:"BlockSize"`
}

type blockByNumberResponse struct {
	Action  string `json:"Action"`
	Error   int    `json:"Error"`
	Desc    string `json:"Desc"`
	Version string `json:"Version"`
	Result  struct {
		PrevBlock  string `json:"PrevBlock"`
		TxnsRoot   string `json:"TxnsRoot"`
		BlockTime  int    `json:"BlockTime"`
		NextBlock  string `json:"NextBlock"`
		BookKeeper string `json:"BookKeeper"`
		TxnNum     int    `json:"TxnNum"`
		Height     int    `json:"Height"`
		TxnList    []struct {
			TxnTime     int    `json:"TxnTime"`
			ConfirmFlag int    `json:"ConfirmFlag"`
			TxnHash     string `json:"TxnHash"`
			Height      int    `json:"Height"`
		} `json:"TxnList"`
		Hash          string `json:"Hash"`
		ConsensusData string `json:"ConsensusData"`
		BlockSize     int    `json:"BlockSize"`
	} `json:"Result"`
}

// Explorer API max returned transactions per page
const TxPerPage = 20

func (c *Client) GetTxsOfAddress(address, assetName string) (txPage *TxPage, err error) {
	url := fmt.Sprintf("address/%s/%s/%d/1", address, assetName, TxPerPage)
	err = c.Get(&txPage, url, nil)
	return
}

func (c *Client) CurrentBlockNumber() (int64, error) {
	url := fmt.Sprintf("blocklist/%d", 1)
	var (
		response blockListResponse
		height   int64
	)
	err := c.Get(&response, url, nil)
	if err != nil {
		return 0, err
	}
	if len(response.Result) > 0 {
		height = (int64)(response.Result[0].Height)
	}
	return height, nil
}

func (c *Client) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	url := fmt.Sprintf("block/%d", num)
	var (
		response blockByNumberResponse
		block    blockatlas.Block
		txs      []blockatlas.Tx
	)
	err := c.Get(&response, url, nil)
	if err != nil {
		return nil, err
	}

	if response.Error == 0 {
		block.ID = response.Result.Hash
		block.Number = int64(response.Result.Height)
		for _, txn := range response.Result.TxnList {
			tx := new(blockatlas.Tx)
			tx.ID = txn.TxnHash
			tx.Block = uint64(txn.Height)
			if txn.ConfirmFlag == 1 {
				tx.Status = blockatlas.StatusCompleted
			}
			tx.Date = int64(txn.TxnTime)
			tx.Coin = coin.Ontology().ID
			txs = append(txs, *tx)
		}
		block.Txs = txs
	}
	return &block, nil
}
