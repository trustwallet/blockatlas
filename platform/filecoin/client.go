package filecoin

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Client struct {
	blockatlas.Request
}

func (c Client) getBlockHeight() (ChainHeadResponse, error) {
	var result ChainHeadResponse
	err := c.RpcCall(&result, "Filecoin.ChainHead", nil)
	if err != nil {
		return ChainHeadResponse{}, err
	}
	return result, nil
}
