package endpoint

import (
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/api/model"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"net/http"
)

// @Summary Get Tokens
// @ID tokens
// @Description Get tokens from the address
// @Accept json
// @Produce json
// @Tags Transactions
// @Param coin path string true "the coin name" default(ethereum)
// @Param address path string true "the query address" default(0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB)
// @Success 200 {object} blockatlas.CollectionPage
// @Failure 500 {object} middleware.ApiError
// @Router /v2/{coin}/tokens/{address} [get]
func GetTokensByAddress(c *gin.Context, tokenAPI blockatlas.TokenAPI) {
	address := c.Param("address")
	if address == "" {
		EmptyPage(c)
		return
	}

	result, err := tokenAPI.GetTokenListByAddress(address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.CreateErrorResponse(model.InternalFail, err))
		return
	}
	c.JSON(http.StatusOK, blockatlas.DocsResponse{Docs: &result})
}
