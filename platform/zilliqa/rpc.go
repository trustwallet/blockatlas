package zilliqa

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"strconv"
	"sync"
)

type RpcClient struct {
	blockatlas.Request
}

func (c *RpcClient) GetBlockchainInfo() (info *ChainInfo, err error) {
	err = c.RpcCall(&info, "GetBlockchainInfo", nil)
	return
}

func (c *RpcClient) GetTx(hash string) (tx Tx, err error) {
	err = c.RpcCall(&tx, "GetTransaction", []string{hash})
	return
}

func (c *RpcClient) GetTxInBlock(number int64) (txs []Tx, err error) {
	strNumber := strconv.FormatInt(number, 10)
	var results [][]string
	err = c.RpcCall(&results, "GetTransactionsForTxBlock", []string{strNumber})

	var wg sync.WaitGroup
	out := make(chan Tx)
	for _, ids := range results {
		for _, id := range ids {
			wg.Add(1)
			go func(id string) {
				defer wg.Done()
				tx, errTx := c.GetTx(id)
				if errTx != nil {
					return
				}
				out <- tx
			}(id)
		}
	}
	go func() {
		for tx := range out {
			txs = append(txs, tx)
		}
	}()
	wg.Wait()
	close(out)
	return
}
