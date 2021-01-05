package rpc

import "github.com/trustwallet/golibs/client"

type Client struct {
	client.Request
}

func (c Client) GetBlockHeight() (ChainHeadResponse, error) {
	var result ChainHeadResponse
	err := c.RpcCall(&result, "Filecoin.ChainHead", nil)
	if err != nil {
		return ChainHeadResponse{}, err
	}
	return result, nil
}

func (c Client) GetTipSetByHeight(height int64) (ChainHeadResponse, error) {
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

func (c Client) GetBlockMessage(cid string) (BlockMessageResponse, error) {
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
