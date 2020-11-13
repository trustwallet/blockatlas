package endpoint

import (
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/trustwallet/golibs/numbers"

	"errors"

	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/coin"
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
// @ID staking_v2_batch
// @Description Get Stake Delegations for multiple coins
// @Accept json
// @Produce json
// @Tags Staking
// @Param delegations body AddressesRequest true "Validators addresses and coins"
// @Success 200 {object} blockatlas.DelegationsBatchPage
// @Router /v2/staking/delegations [post]
func GetStakeDelegationsWithAllInfoForBatch(c *gin.Context, apis map[string]blockatlas.StakeAPI) {
	var reqs AddressesRequest
	if err := c.BindJSON(&reqs); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
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
		delegation.Delegations = sortDelegations(delegation.Delegations)
		batch = append(batch, delegation)
	}
	c.JSON(http.StatusOK, blockatlas.DocsResponse{Docs: &batch})
}

// @Summary Get Multiple Stake Delegations
// @ID staking_v2
// @Description Get Stake Delegations for multiple coins
// @Accept json
// @Produce json
// @Tags Staking
// @Param delegations body AddressesRequest true "Validators addresses and coins"
// @Success 200 {object} blockatlas.DelegationsBatchPage
// @Router /v2/staking/list [post]
func GetStakeInfoForBatch(c *gin.Context, apis map[string]blockatlas.StakeAPI) {
	var reqs CoinsRequest
	if err := c.BindJSON(&reqs); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
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

// @Summary Get staking info by coin ID
// @ID staking_v3
// @Description Get staking info by coin ID
// @Produce json
// @Tags Staking
// @Param coins query string true "List of coins"
// @Success 200 {array} blockatlas.DelegationsBatchPage
// @Failure 400 {object} ErrorResponse
// @Router /v3/staking/list [get]
func GetStakeInfoForCoins(c *gin.Context, apis map[string]blockatlas.StakeAPI) {
	coinsRequest := c.Query("coins")
	if coinsRequest == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(errors.New("empty coins list")))
		return
	}

	coinsRaw := strings.Split(coinsRequest, ",")

	coins, err := numbers.SliceAtoi(coinsRaw)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var reqs CoinsRequest
	for _, c := range coins {
		reqs = append(reqs, CoinBatchRequest{Coin: uint(c)})
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
// @ID validators_v2
// @Description Get validators from the address
// @Accept json
// @Produce json
// @Tags Staking
// @Param coin path string true "the coin name" default(cosmos)
// @Success 200 {object} blockatlas.DocsResponse
// @Failure 500 {object} ErrorResponse
// @Router /v2/{coin}/staking/validators [get]
func GetValidators(c *gin.Context, api blockatlas.StakeAPI) {
	results, err := api.GetActiveValidators()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
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
// @Failure 500 {object} ErrorResponse
// @Router /v2/{coin}/staking/delegations/{address} [get]
func GetStakingDelegationsForSpecificCoin(c *gin.Context, api blockatlas.StakeAPI) {
	result, err := getDelegationResponse(api, c.Param("address"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	result.Delegations = sortDelegations(result.Delegations)
	c.JSON(http.StatusOK, &result)
}

func getDelegationResponse(api blockatlas.StakeAPI, address string) (blockatlas.DelegationResponse, error) {
	delegations, err := api.GetDelegations(address)
	if err != nil {
		return blockatlas.DelegationResponse{
			StakingResponse: getStakingResponse(api),
			Address:         address,
		}, err
	}
	balance, err := api.UndelegatedBalance(address)
	if err != nil {
		return blockatlas.DelegationResponse{
			Delegations:     delegations,
			Address:         address,
			StakingResponse: getStakingResponse(api),
		}, err
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

func sortDelegations(delegations blockatlas.DelegationsPage) blockatlas.DelegationsPage {
	sort.Slice(delegations, func(i, j int) bool {
		iA, err := strconv.Atoi(delegations[i].Value)
		if err != nil {
			return false
		}
		jA, err := strconv.Atoi(delegations[j].Value)
		if err != nil {
			return false
		}
		return iA > jA
	})
	return delegations
}
