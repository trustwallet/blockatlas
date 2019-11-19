package api

import (
	"github.com/trustwallet/blockatlas/pkg/errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/trustwallet/blockatlas/pkg/ginutils"

	"github.com/gin-gonic/gin"
	CoinType "github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/platform"
)

var TLDMapping = map[string]uint64{}

type LookupBatchPage []blockatlas.Resolved

// @Summary Lookup .eth / .zil addresses
// @ID lookup
// @Description Lookup ENS/ZNS to find registered addresses
// @Produce json
// @Tags ns
// @Param name query string empty "string name"
// @Param coin query string 60 "string coin"
// @Success 200 {object} blockatlas.Resolved
// @Failure 500 {object} ginutils.ApiError
// @Router /ns/lookup [get]
func MakeLookupRoute(router gin.IRouter) {
	ns := router.Group("/ns")
	ns.GET("/lookup", func(c *gin.Context) {
		name := c.Query("name")
		coinQuery := c.Query("coin")
		coin, err := strconv.ParseUint(coinQuery, 10, 64)
		if err != nil {
			ginutils.RenderError(c, http.StatusBadRequest, "coin query is invalid")
			return
		}

		result, err := handleLookup(name, coin)
		if err != nil {
			ginutils.RenderError(c, http.StatusBadRequest, err.Error())
		}
		ginutils.RenderSuccess(c, result)
	})

	TLDMapping[".eth"] = CoinType.ETH
	TLDMapping[".xyz"] = CoinType.ETH
	TLDMapping[".luxe"] = CoinType.ETH
	TLDMapping[".zil"] = CoinType.ZIL
}

func handleLookup(name string, coin uint64) (blockatlas.Resolved, error) {
	result := blockatlas.Resolved{Result: "", Coin: coin}
	if name == "" {
		return result, errors.E("name query is missing")
	}

	name = strings.ToLower(name)
	for tld, id := range TLDMapping {
		if strings.HasSuffix(name, tld) {
			api := platform.NamingAPIs[id]
			resolved, err := api.Lookup(coin, name)
			if err != nil {
				return result, err
			} else {
				return resolved, nil
			}
		}
	}
	return result, nil
}
