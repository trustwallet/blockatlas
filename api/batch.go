package api

import (
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/ginutils"
	"github.com/trustwallet/blockatlas/pkg/ginutils/gincache"
	"github.com/trustwallet/blockatlas/platform"
	"strconv"
	"time"
)

type AddressBatchRequest struct {
	Address string `json:"address"`
	CoinBatchRequest
}

type CoinBatchRequest struct {
	Coin uint `json:"coin"`
}

type ENSBatchRequest struct {
	Coins []uint64 `json:"coins"`
	Name  string   `json:"name"`
}

type AddressesRequest []AddressBatchRequest
type CoinsRequest []CoinBatchRequest

// @Summary Get Multiple Stake Delegations
// @ID batch_delegations
// @Description Get Stake Delegations for multiple coins
// @Accept json
// @Produce json
// @Tags Platform-Staking
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
			c, ok := coin.Coins[r.Coin]
			if !ok {
				continue
			}
			p, ok := platform.StakeAPIs[c.Handle]
			if !ok {
				continue
			}
			delegation, err := getDelegationResponse(p, r.Address)
			if err != nil {
				continue
			}
			batch = append(batch, delegation)
		}
		ginutils.RenderSuccess(c, blockatlas.DocsResponse{Docs: batch})
	})
}

// @Summary Get Multiple Stake Delegations
// @ID batch_delegations
// @Description Get Stake Delegations for multiple coins
// @Accept json
// @Produce json
// @Tags Platform-Staking
// @Param delegations body api.AddressesRequest true "Validators addresses and coins"
// @Success 200 {object} blockatlas.DelegationsBatchPage
// @Router /v2/staking/list [post]
func makeStakingDelegationsSimpleBatchRoute(router gin.IRouter) {
	router.POST("/staking/list", gincache.CacheMiddleware(time.Hour*24, func(c *gin.Context) {
		var reqs CoinsRequest
		if err := c.BindJSON(&reqs); err != nil {
			ginutils.ErrorResponse(c).Message(err.Error()).Render()
			return
		}

		batch := make(blockatlas.StakingBatchPage, 0)
		for _, r := range reqs {
			c, ok := coin.Coins[r.Coin]
			if !ok {
				continue
			}
			p, ok := platform.StakeAPIs[c.Handle]
			if !ok {
				continue
			}
			staking := getStakingResponse(p)
			batch = append(batch, staking)
		}
		ginutils.RenderSuccess(c, blockatlas.DocsResponse{Docs: batch})
	}))
}

// @Description Get collection categories
// @ID collection_categories_v2
// @Summary Get list of collections from a specific coin and addresses
// @Accept json
// @Produce json
// @Tags Platform-Collections
// @Param data body string true "Payload" default({"60": ["0xb3624367b1ab37daef42e1a3a2ced012359659b0"]})
// @Success 200 {object} blockatlas.DocsResponse
// @Router /v2/collectibles/categories [post]
//TODO: remove once most of the clients will be updated (deadline: March 17th)
func oldMakeCategoriesBatchRoute(router gin.IRouter) {
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
				collections, err := p.OldGetCollections(address)
				if err != nil {
					continue
				}
				batch = append(batch, collections...)
			}
		}
		ginutils.RenderSuccess(c, batch)
	})
}

// @Description Get collection categories
// @ID collection_categories_v3
// @Summary Get list of collections from a specific coin and addresses
// @Accept json
// @Produce json
// @Tags Platform-Collections
// @Param data body string true "Payload" default({"60": ["0xb3624367b1ab37daef42e1a3a2ced012359659b0"]})
// @Success 200 {object} blockatlas.DocsResponse
// @Router /v3/collectibles/categories [post]
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
