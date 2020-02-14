package api

import (
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/services/domains"
	"net/http"
	"strconv"
	"strings"

	"github.com/trustwallet/blockatlas/pkg/ginutils"

	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type LookupBatchPage []blockatlas.Resolved

// @Summary Lookup .eth / .zil addresses
// @ID lookup
// @Description Lookup ENS/ZNS to find registered addresses
// @Produce json
// @Tags Platform-Naming-Service
// @Param name query string empty "string name"
// @Param coin query string 60 "string coin"
// @Success 200 {object} blockatlas.Resolved
// @Failure 500 {object} ginutils.ApiError
// @Router /ns/lookup [get]
func MakeLookupRoute(router gin.IRouter) {
	router.GET("/lookup", func(c *gin.Context) {
		name := c.Query("name")
		coinQuery := c.Query("coin")
		coin, err := strconv.ParseUint(coinQuery, 10, 64)
		if err != nil {
			ginutils.RenderError(c, http.StatusBadRequest, "coin query is invalid")
			return
		}

		result, err := domains.HandleLookup(name, []uint64{coin})
		if err != nil {
			ginutils.RenderError(c, http.StatusBadRequest, err.Error())
			return
		}
		if len(result) == 0 {
			ginutils.RenderError(c, http.StatusBadRequest, errors.E("name not found", errors.Params{"coin": coin, "name": name}).Error())
			return
		}
		ginutils.RenderSuccess(c, result[0])
	})
}

// @Summary Lookup .eth / .zil addresses
// @ID lookup
// @Description Lookup ENS/ZNS to find registered addresses for multiple coins
// @Produce json
// @Tags Platform-Naming-Service
// @Param name query string empty "string name"
// @Param coins query string true "List of coins"
// @Success 200 {array} blockatlas.Resolved
// @Failure 500 {object} ginutils.ApiError
// @Router /v2/ns/lookup [get]
func MakeLookupBatchRoute(router gin.IRouter) {
	router.GET("/lookup", func(c *gin.Context) {
		name := c.Query("name")
		coinsRaw := strings.Split(c.Query("coins"), ",")
		coins, err := sliceAtoi(coinsRaw)
		if err != nil {
			ginutils.RenderError(c, http.StatusBadRequest, "coin query is invalid")
			return
		}

		result, err := domains.HandleLookup(name, coins)
		if err != nil {
			ginutils.RenderError(c, http.StatusBadRequest, err.Error())
			return
		}
		if len(result) == 0 {
			ginutils.RenderError(c, http.StatusBadRequest, errors.E("name not found", errors.Params{"name": name}).Error())
			return
		}
		ginutils.RenderSuccess(c, result)
	})
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
