package endpoint

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

// @Summary Get Block
// @ID block_v2
// @Description Get Block information
// @Accept json
// @Produce json
// @Tags Transactions
// @Param coin path string true "the coin name" default(zilliqa)
// @Param address path string true "the query address" default(850321)
// @Failure 500 {object} ErrorResponse
// @Router /v2/{coin}/blocks/{block} [get]
func GetBlock(c *gin.Context, blockAPI blockatlas.BlockAPI) {
	blockString := c.Param("block")
	blockNumber, err := strconv.ParseUint(blockString, 10, 32)

	if err != nil || blockNumber < 1 {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(errors.New("invalid block number")))
		return
	}

	block, err := blockAPI.GetBlockByNumber(int64(blockNumber))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(errors.New("block number not found")))
		return
	}

	c.JSON(http.StatusOK, &block)
}
