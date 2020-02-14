package api

import (
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/ginutils"
	"github.com/trustwallet/blockatlas/pkg/ginutils/gincache"
	"github.com/trustwallet/blockatlas/pkg/logger"
	services "github.com/trustwallet/blockatlas/services/assets"
	"time"
)

// @Summary Get Validators
// @ID validators
// @Description Get validators from the address
// @Accept json
// @Produce json
// @Tags Platform-Staking
// @Param coin path string true "the coin name" default(cosmos)
// @Success 200 {object} blockatlas.DocsResponse
// @Failure 500 {object} ginutils.ApiError
// @Router /v2/{coin}/staking/validators [get]
func makeStakingValidatorsRoute(router gin.IRouter, api blockatlas.Platform) {
	var stakingAPI blockatlas.StakeAPI
	stakingAPI, _ = api.(blockatlas.StakeAPI)

	if stakingAPI == nil {
		return
	}

	router.GET("/staking/validators", gincache.CacheMiddleware(time.Hour, func(c *gin.Context) {
		results, err := services.GetValidators(stakingAPI)
		if err != nil {
			logger.Error(err)
			ginutils.ErrorResponse(c).Message(err.Error()).Render()
			return
		}
		ginutils.RenderSuccess(c, blockatlas.DocsResponse{Docs: results})
	}))
}


// @Summary Get Stake Delegations
// @ID delegations
// @Description Get stake delegations from the address
// @Accept json
// @Produce json
// @Tags Platform-Staking
// @Param coin path string true "the coin name" default(tron)
// @Param address path string true "the query address" default(TPJYCz8ppZNyvw7pTwmjajcx4Kk1MmEUhD)
// @Success 200 {object} blockatlas.DelegationResponse
// @Failure 500 {object} ginutils.ApiError
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
			ginutils.ErrorResponse(c).Message(err.Error()).Render()
			return
		}

		ginutils.RenderSuccess(c, response)
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
