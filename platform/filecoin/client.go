package filecoin

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/platform/filecoin/explorer"
	"github.com/trustwallet/blockatlas/platform/filecoin/rpc"
)

type Client struct {
	blockatlas.Request
}

func (c Client) getBlockHeight() (rpc.ChainHeadResponse, error) {
	var result rpc.ChainHeadResponse
	err := c.RpcCall(&result, "Filecoin.ChainHead", nil)
	if err != nil {
		return rpc.ChainHeadResponse{}, err
	}
	return result, nil
}

func (c Client) getTipSetByHeight(height int64) (rpc.ChainHeadResponse, error) {
	var result rpc.ChainHeadResponse
	params := []interface{}{
		height, nil,
	}
	err := c.RpcCall(&result, "Filecoin.ChainGetTipSetByHeight", params)
	if err != nil {
		return rpc.ChainHeadResponse{}, err
	}
	return result, nil
}

func (c Client) getBlockMessage(cid string) (rpc.BlockMessageResponse, error) {
	var result rpc.BlockMessageResponse
	params := []interface{}{
		map[string]interface{}{
			"/": cid,
		},
	}
	err := c.RpcCall(&result, "Filecoin.ChainGetBlockMessages", params)
	if err != nil {
		return rpc.BlockMessageResponse{}, err
	}
	return result, nil
}

func (c Client) getMessagesByAddress(address string, pageSize int) (res explorer.Response, err error) {
	path := fmt.Sprintf("/v1/address/%s/messages", address)
	query := url.Values{"pageSize": {strconv.Itoa(pageSize)}}
	err = c.Get(&res, path, query)
	if err != nil {
		return res, err
	}
	return
}
