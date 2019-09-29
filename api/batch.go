package api

import (
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/platform"
)

type AddressBatchRequest struct {
	Coin    uint   `json:"coin"`
	Address string `json:"address"`
}

type AddressesRequest []AddressBatchRequest

func makeStakingDelegationsBatchRoute(router gin.IRouter) {
	router.POST("/staking/delegations/", func(c *gin.Context) {
		var reqs AddressesRequest
		if c.BindJSON(&reqs) != nil {
			return
		}

		if len(reqs) == 0 {
			RenderSuccess(c, EmptyResponse)
			return
		}

		batch := make(blockatlas.DelegationsBatchPage, 0)
		for _, r := range reqs {
			d := blockatlas.DelegationsBatch{
				Coin:    r.Coin,
				Address: r.Address,
			}

			c := coin.Coins[r.Coin]
			p := platform.StakeAPIs[c.Handle]
			delegations, err := p.GetDelegations(r.Address)
			if err != nil {
				d.Error = err.Error()
			} else {
				d.Delegations = delegations
			}
			batch = append(batch, d)
		}
		RenderSuccess(c, blockatlas.DocsResponse{Docs: batch})
	})
}
