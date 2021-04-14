package endpoint

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/services/tokenindexer"
	"github.com/trustwallet/golibs/types"
)

// @Summary Get Tokens
// @ID tokens
// @Description Get tokens from the address
// @Accept json
// @Produce json
// @Tags Transactions
// @Param coin path string true "the coin name" default(ethereum)
// @Param address path string true "the query address" default(0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB)
// @Success 200 {object} []string
// @Failure 500 {object} ErrorResponse
// @Router /v2/{coin}/tokens/{address} [get]
func GetTokensByAddress(c *gin.Context, tokenAPI blockatlas.TokensAPI) {
	address := c.Param("address")
	if address == "" {
		c.JSON(http.StatusOK, []types.Token{})
		return
	}

	result, err := tokenAPI.GetTokenListByAddress(address)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetTokensIdsByAddress(c *gin.Context, tokenAPI blockatlas.TokensAPI) {
	address := c.Param("address")
	if address == "" {
		c.JSON(http.StatusOK, []string{})
		return
	}

	result, err := tokenAPI.GetTokenListIdsByAddress(address)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetTokensByAddressV3(c *gin.Context, instance tokenindexer.Instance) {
	var query tokenindexer.GetTokensByAddressRequest
	if err := c.Bind(&query); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	result, err := instance.GetTokensByAddress(query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Description Get new tokens
// @ID tokens_new_v3
// @Summary Get list of new tokens by coin from specific unix timstamp
// @Accept json
// @Produce json
// @Tags Transactions
// @Param from query int true "unix timestamp"
// @Success 200 {object} tokenindexer.Response
// @Router /v3/tokens/new [get]
func GetNewTokens(c *gin.Context, instance tokenindexer.Instance) {
	var request tokenindexer.Request
	fromRaw := c.Query("from")

	from, err := strconv.Atoi(fromRaw)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(errors.New("invalid from param")))
		return
	}
	request.From = int64(from)

	resp, err := instance.GetNewTokensRequest(request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	c.JSON(http.StatusOK, resp)
}
