package bakingbad

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Client struct {
	blockatlas.Request
}

func (c *Client) GetBakers() (baker *[]Baker, err error) {
	err = c.Get(&baker, "v2/bakers", nil)
	return
}
