package endpoint

import (
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/api/model"
	"github.com/trustwallet/blockatlas/services/domains"
	"net/http"
	"strconv"
	"strings"
)

// @Summary Lookup .eth / .zil addresses
// @ID lookup
// @Description Lookup ENS/ZNS to find registered addresses
// @Produce json
// @Tags Naming
// @Param name query string empty "string name"
// @Param coin query string 60 "string coin"
// @Success 200 {object} blockatlas.Resolved
// @Failure 500 {object} middleware.ApiError
// @Router /ns/lookup [get]
func GetAddressByCoinAndDomain(c *gin.Context) {
	name := c.Query("name")
	coinQuery := c.Query("coin")
	coin, err := strconv.ParseUint(coinQuery, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.CreateErrorResponse(model.InvalidQuery, err))
		return
	}
	result, err := domains.HandleLookup(name, []uint64{coin})
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.CreateErrorResponse(model.InternalFail, err))
		return
	}
	if len(result) == 0 {
		c.JSON(http.StatusNotFound, model.CreateErrorResponse(model.RequestedDataNotFound, err))
		return
	}
	c.JSON(http.StatusOK, result[0])
}

// @Summary Lookup .eth / .zil addresses
// @ID lookup
// @Description Lookup ENS/ZNS to find registered addresses for multiple coins
// @Produce json
// @Tags Naming
// @Param name query string empty "string name"
// @Param coins query string true "List of coins"
// @Success 200 {array} blockatlas.Resolved
// @Failure 500 {object} middleware.ApiError
// @Router /v2/ns/lookup [get]
func GetAddressByCoinAndDomainBatch(c *gin.Context) {
	name := c.Query("name")
	coinsRaw := strings.Split(c.Query("coins"), ",")
	coins, err := sliceAtoi(coinsRaw)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.CreateErrorResponse(model.InvalidQuery, err))
		return
	}
	result, err := domains.HandleLookup(name, coins)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.CreateErrorResponse(model.InternalFail, err))
		return
	}
	if len(result) == 0 {
		c.JSON(http.StatusNotFound, model.CreateErrorResponse(model.RequestedDataNotFound, err))
		return
	}
	c.JSON(http.StatusOK, &result)
}

func sliceAtoi(sa []string) ([]uint64, error) {
	si := make([]uint64, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.ParseUint(a, 10, 64)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}
