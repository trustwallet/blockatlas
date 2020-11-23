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

func (c Client) getTipSetByHeight(height int64) (ChainHeadResponse, error) {
	var result ChainHeadResponse
	params := []interface{}{
		height, nil,
	}
	err := c.RpcCall(&result, "Filecoin.ChainGetTipSetByHeight", params)
	if err != nil {
		return ChainHeadResponse{}, err
	}
	return result, nil
}

func (c Client) getBlockMessage(cid string) (BlockMessageResponse, error) {
	var result BlockMessageResponse
	params := []interface{}{
		map[string]interface{}{
			"/": cid,
		},
	}
	err := c.RpcCall(&result, "Filecoin.ChainGetBlockMessages", params)
	if err != nil {
		return BlockMessageResponse{}, err
	}
	return result, nil
}
