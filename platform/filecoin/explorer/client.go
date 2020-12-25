package explorer

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Client struct {
	blockatlas.Request
}

func (c Client) GetMessagesByAddress(address string, pageSize int) (res Response, err error) {
	path := fmt.Sprintf("/v1/address/%s/messages", address)
	query := url.Values{"pageSize": {strconv.Itoa(pageSize)}}
	err = c.Get(&res, path, query)
	if err != nil {
		return res, err
	}
	return
}
