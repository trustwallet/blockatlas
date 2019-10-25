package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	CoinType "github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/platform"
)

var TLDMapping = map[string]uint64{}

// @Summary Lookup .eth / .zil addresses
// @ID lookup
// @Description Lookup ENS/ZNS to find registered addresses
// @Produce json
// @Tags ns
// @Param name query string empty "string name"
// @Param coin query string 60 "string coin"
// @Success 200 {object} blockatlas.Resolved
// @Failure 500 {object} api.ApiError
// @Router /ns/lookup [get]
func MakeLookupRoute(router gin.IRouter) {
	ns := router.Group("/ns")
	ns.GET("/lookup", handleLookup)

	TLDMapping[".eth"] = CoinType.ETH
	TLDMapping[".zil"] = CoinType.ZIL
}

func handleLookup(c *gin.Context) {
	name := c.Query("name")
	coinQuery := c.Query("coin")

	if name == "" {
		RenderError(c, http.StatusBadRequest, "name query is missing")
		return
	}
	coin, err := strconv.ParseUint(coinQuery, 10, 64)
	if err != nil {
		RenderError(c, http.StatusBadRequest, "coin query is invalid")
		return
	}
	name = strings.ToLower(name)

	for tld, id := range TLDMapping {
		if strings.HasSuffix(name, tld) {
			api := platform.NamingAPIs[id]
			result, err := api.Lookup(coin, name)
			if err != nil {
				RenderError(c, http.StatusInternalServerError, err.Error())
				return
			} else {
				RenderSuccess(c, result)
				return
			}
		}
	}
	RenderSuccess(c, blockatlas.Resolved{Result: "", Coin: coin})
}
