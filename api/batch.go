package api

import (
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/ginutils"
	"github.com/trustwallet/blockatlas/platform"
)

type AddressBatchRequest struct {
	Coin    uint   `json:"coin"`
	Address string `json:"address"`
}

type AddressesRequest []AddressBatchRequest

// @Summary Get Multiple Stake Delegations
// @ID batch_delegations
// @Description Get Stake Delegations for multiple coins
// @Accept json
// @Produce json
// @Tags platform,staking
// @Param delegations body api.AddressesRequest true "Validators addresses and coins"
// @Success 200 {object} blockatlas.DelegationsBatchPage
// @Router /v2/staking/delegations [post]
func makeStakingDelegationsBatchRoute(router gin.IRouter) {
	router.POST("/staking/delegations", func(c *gin.Context) {
		var reqs AddressesRequest
		if err := c.BindJSON(&reqs); err != nil {
			ginutils.ErrorResponse(c).Message(err.Error()).Render()
			return
		}

		batch := make(blockatlas.DelegationsBatchPage, 0)
		for _, r := range reqs {
			c := coin.Coins[r.Coin]

			d := blockatlas.DelegationsBatch{
				Coin:    c.External(),
				Address: r.Address,
			}

			p := platform.StakeAPIs[c.Handle]
			delegations, err := p.GetDelegations(r.Address)
			balance, err := p.UndelegatedBalance(r.Address)

			if err != nil {
				d.Error = err.Error()
			} else {
				d.Delegations = delegations
				d.Details = p.GetDetails()
				d.Balance = balance
			}
			batch = append(batch, d)
		}
		ginutils.RenderSuccess(c, blockatlas.DocsResponse{Docs: batch})
	})
}
