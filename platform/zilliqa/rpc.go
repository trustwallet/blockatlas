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

func (c *RpcClient) GetTx(hash string) (tx TxRPC, err error) {
	err = c.RpcCall(&tx, "GetTransaction", []string{hash})
	return
}

func (c *RpcClient) GetTxInBlock(number int64) []Tx {
	strNumber := strconv.Itoa(int(number))
	txs := make([]Tx, 0)

	var results BlockTxs
	err := c.RpcCall(&results, "GetTransactionsForTxBlock", []string{strNumber})
	if err != nil {
		return txs
	}

	var wg sync.WaitGroup
	btxs := results.txs()
	out := make(chan Tx, len(btxs))
	wg.Add(len(btxs))
	for _, id := range btxs {
		go func(hash string, txChan chan Tx, wg *sync.WaitGroup) {
			defer wg.Done()
			tx, errTx := c.GetTx(hash)
			if errTx != nil {
				return
			}
			txChan <- tx.toTx()
		}(id, out, &wg)
	}
	wg.Wait()
	close(out)

	for tx := range out {
		txs = append(txs, tx)
	}
	return txs
}
