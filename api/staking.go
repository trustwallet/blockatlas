package api

import (
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/api/middleware"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/platform"
	services "github.com/trustwallet/blockatlas/services/assets"
	"net/http"
	"time"
)

type (
	AddressBatchRequest struct {
		Address string `json:"address"`
		CoinBatchRequest
	}

	CoinBatchRequest struct {
		Coin uint `json:"coin"`
	}

	ENSBatchRequest struct {
		Coins []uint64 `json:"coins"`
		Name  string   `json:"name"`
	}

	AddressesRequest []AddressBatchRequest
	CoinsRequest     []CoinBatchRequest
)

// @Summary Get Multiple Stake Delegations
// @ID batch_delegations
// @Description Get Stake Delegations for multiple coins
// @Accept json
// @Produce json
// @Tags Staking
// @Param delegations body api.AddressesRequest true "Validators addresses and coins"
// @Success 200 {object} blockatlas.DelegationsBatchPage
// @Router /v2/staking/delegations [post]
func makeStakingDelegationsBatchRoute(router gin.IRouter) {
	router.POST("/staking/delegations", func(c *gin.Context) {
		var reqs AddressesRequest
		if err := c.BindJSON(&reqs); err != nil {
			c.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidQuery, err))
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
		c.JSON(http.StatusOK, blockatlas.DocsResponse{Docs: &batch})
	})
}

// @Summary Get Multiple Stake Delegations
// @ID batch_delegations
// @Description Get Stake Delegations for multiple coins
// @Accept json
// @Produce json
// @Tags Staking
// @Param delegations body api.AddressesRequest true "Validators addresses and coins"
// @Success 200 {object} blockatlas.DelegationsBatchPage
// @Router /v2/staking/list [post]
func makeStakingDelegationsSimpleBatchRoute(router gin.IRouter) {
	router.POST("/staking/list", middleware.CacheMiddleware(time.Hour*24, func(c *gin.Context) {
		var reqs CoinsRequest
		if err := c.BindJSON(&reqs); err != nil {
			c.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidQuery, err))
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
		c.JSON(http.StatusOK, blockatlas.DocsResponse{Docs: &batch})
	}))
}

// @Summary Get Validators
// @ID validators
// @Description Get validators from the address
// @Accept json
// @Produce json
// @Tags Staking
// @Param coin path string true "the coin name" default(cosmos)
// @Success 200 {object} blockatlas.DocsResponse
// @Failure 500 {object} middleware.ApiError
// @Router /v2/{coin}/staking/validators [get]
func makeStakingValidatorsRoute(router gin.IRouter, api blockatlas.Platform) {
	var stakingAPI blockatlas.StakeAPI
	stakingAPI, _ = api.(blockatlas.StakeAPI)

	if stakingAPI == nil {
		return
	}

	router.GET("/staking/validators", middleware.CacheMiddleware(time.Hour, func(c *gin.Context) {
		results, err := services.GetActiveValidators(stakingAPI)
		if err != nil {
			c.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalFail, err))
			return
		}
		c.JSON(http.StatusOK, blockatlas.DocsResponse{Docs: &results})
	}))
}

// @Summary Get Stake Delegations
// @ID delegations
// @Description Get stake delegations from the address
// @Accept json
// @Produce json
// @Tags Staking
// @Param coin path string true "the coin name" default(tron)
// @Param address path string true "the query address" default(TPJYCz8ppZNyvw7pTwmjajcx4Kk1MmEUhD)
// @Success 200 {object} blockatlas.DelegationResponse
// @Failure 500 {object} middleware.ApiError
// @Router /v2/{coin}/staking/delegations/{address} [get]
func makeStakingDelegationsRoute(router gin.IRouter, api blockatlas.Platform) {
	var stakingAPI blockatlas.StakeAPI
	stakingAPI, _ = api.(blockatlas.StakeAPI)

	if stakingAPI == nil {
		return
	}

	router.GET("/staking/delegations/:address", func(c *gin.Context) {
		response, err := getDelegationResponse(stakingAPI, c.Param("address"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalFail, err))
			return
		}
		c.JSON(http.StatusOK, &response)
	})
}

func getDelegationResponse(p blockatlas.StakeAPI, address string) (blockatlas.DelegationResponse, error) {
	delegations, err := p.GetDelegations(address)
	if err != nil {
		return blockatlas.DelegationResponse{
			StakingResponse: getStakingResponse(p),
		}, errors.E("Unable to fetch delegations list", err)
	}
	balance, err := p.UndelegatedBalance(address)
	if err != nil {
		return blockatlas.DelegationResponse{
			StakingResponse: getStakingResponse(p),
		}, errors.E("Unable to fetch undelegated balance", err)
	}
	return blockatlas.DelegationResponse{
		Balance:         balance,
		Delegations:     delegations,
		Address:         address,
		StakingResponse: getStakingResponse(p),
	}, nil
}

func getStakingResponse(p blockatlas.StakeAPI) blockatlas.StakingResponse {
	c := p.Coin()
	return blockatlas.StakingResponse{
		Coin:    c.External(),
		Details: p.GetDetails(),
	}
}
