package api

import (
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/ginutils"
	"github.com/trustwallet/blockatlas/platform"
	"strconv"
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
			p := platform.StakeAPIs[c.Handle]
			delegation, err := getDelegationResponse(p, r.Address)
			if err != nil {
				delegation = blockatlas.DelegationResponse{
					Address: r.Address,
					Coin:    c.External(),
					Error:   err,
				}
			}
			batch = append(batch, delegation)
		}
		ginutils.RenderSuccess(c, blockatlas.DocsResponse{Docs: batch})
	})
}

// @Description Get collection categories
// @ID collection_categories
// @Summary Get list of collections from a specific coin and addresses
// @Accept json
// @Produce json
// @Tags Collectibles
// @Param data body string true "Payload" default({"60": ["0xb3624367b1ab37daef42e1a3a2ced012359659b0"]})
// @Success 200 {object} blockatlas.DocsResponse
// @Router /v2/collectibles/categories [post]
func makeCategoriesBatchRoute(router gin.IRouter) {
	router.POST("/collectibles/categories", func(c *gin.Context) {
		var reqs map[string][]string
		if err := c.BindJSON(&reqs); err != nil {
			ginutils.ErrorResponse(c).Message(err.Error()).Render()
			return
		}

		batch := make(blockatlas.CollectionPage, 0)
		for key, addresses := range reqs {
			coinId, err := strconv.Atoi(key)
			if err != nil {
				continue
			}
			p, ok := platform.CollectionAPIs[uint(coinId)]
			if !ok {
				continue
			}
			for _, address := range addresses {
				collections, err := p.GetCollections(address)
				if err != nil {
					continue
				}
				batch = append(batch, collections...)
			}
		}
		ginutils.RenderSuccess(c, batch)
	})
}
