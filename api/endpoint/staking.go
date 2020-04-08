package endpoint

import (
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/api/model"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	services "github.com/trustwallet/blockatlas/services/assets"
	"net/http"
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
func GetStakeDelegationsWithAllInfoForBatch(c *gin.Context, apis map[string]blockatlas.StakeAPI) {
	var reqs AddressesRequest
	if err := c.BindJSON(&reqs); err != nil {
		c.JSON(http.StatusBadRequest, model.CreateErrorResponse(model.InvalidQuery, err))
		return
	}

	batch := make(blockatlas.DelegationsBatchPage, 0)
	for _, r := range reqs {
		requestCoin, ok := coin.Coins[r.Coin]
		if !ok {
			continue
		}
		p, ok := apis[requestCoin.Handle]
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
func GetStakeInfoForBatch(c *gin.Context, apis map[string]blockatlas.StakeAPI) {
	var reqs CoinsRequest
	if err := c.BindJSON(&reqs); err != nil {
		c.JSON(http.StatusBadRequest, model.CreateErrorResponse(model.InvalidQuery, err))
		return
	}

	batch := make(blockatlas.StakingBatchPage, 0)
	for _, r := range reqs {
		requestCoin, ok := coin.Coins[r.Coin]
		if !ok {
			continue
		}
		p, ok := apis[requestCoin.Handle]
		if !ok {
			continue
		}
		staking := getStakingResponse(p)
		batch = append(batch, staking)
	}
	c.JSON(http.StatusOK, blockatlas.DocsResponse{Docs: &batch})
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
func GetValidators(c *gin.Context, api blockatlas.StakeAPI) {
	results, err := services.GetActiveValidators(api)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.CreateErrorResponse(model.InternalFail, err))
		return
	}
	c.JSON(http.StatusOK, blockatlas.DocsResponse{Docs: &results})
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
func GetStakingDelegationsForSpecificCoin(c *gin.Context, api blockatlas.StakeAPI) {
	result, err := getDelegationResponse(api, c.Param("address"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.CreateErrorResponse(model.InternalFail, err))
		return
	}
	c.JSON(http.StatusOK, &result)
}

func getDelegationResponse(api blockatlas.StakeAPI, address string) (blockatlas.DelegationResponse, error) {
	delegations, err := api.GetDelegations(address)
	if err != nil {
		return blockatlas.DelegationResponse{
			StakingResponse: getStakingResponse(api),
			Address:         address,
		}, errors.E("Unable to fetch delegations list", err)
	}
	balance, err := api.UndelegatedBalance(address)
	if err != nil {
		return blockatlas.DelegationResponse{
			Delegations:     delegations,
			Address:         address,
			StakingResponse: getStakingResponse(api),
		}, errors.E("Unable to fetch undelegated balance", err)
	}
	return blockatlas.DelegationResponse{
		Balance:         balance,
		Delegations:     delegations,
		Address:         address,
		StakingResponse: getStakingResponse(api),
	}, nil
}

func getStakingResponse(api blockatlas.StakeAPI) blockatlas.StakingResponse {
	stakingCoin := api.Coin()
	return blockatlas.StakingResponse{
		Coin:    stakingCoin.External(),
		Details: api.GetDetails(),
	}
}
