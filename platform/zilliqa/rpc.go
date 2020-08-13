package zilliqa

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
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
