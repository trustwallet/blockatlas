package endpoint

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/services/naming"
)

// @Summary Lookup .eth / .zil addresses
// @ID lookup
// @Description Lookup ENS/ZNS to find registered addresses
// @Produce json
// @Tags Naming
// @Param name query string empty "string name"
// @Param coin query string 60 "string coin"
// @Success 200 {object} blockatlas.Resolved
// @Failure 500 {object} ErrorResponse
// @Router /ns/lookup [get]
func GetAddressByCoinAndNaming(c *gin.Context) {
	name := c.Query("name")
	coinQuery := c.Query("coin")
	coin, err := strconv.ParseUint(coinQuery, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	result, err := naming.HandleLookup(name, []uint64{coin})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(err))
		return
	}
	if len(result) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(err))
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
// @Failure 500 {object} ErrorResponse
// @Router /v2/ns/lookup [get]
func GetAddressByCoinAndNamingBatch(c *gin.Context) {
	name := c.Query("name")
	coinsRaw := strings.Split(c.Query("coins"), ",")
	coins, err := sliceAtoi(coinsRaw)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	result, err := naming.HandleLookup(name, coins)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(err))
		return
	}
	if len(result) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(err))
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
