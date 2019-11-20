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

		result, err := HandleLookup(name, []uint64{coin})
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

	TLDMapping[".eth"] = CoinType.ETH
	TLDMapping[".xyz"] = CoinType.ETH
	TLDMapping[".luxe"] = CoinType.ETH
	TLDMapping[".zil"] = CoinType.ZIL
}

func HandleLookup(name string, coins []uint64) (result []blockatlas.Resolved, err error) {
	if name == "" {
		return nil, errors.E("name is missing")
	}
	if len(coins) == 0 {
		return nil, errors.E("coins are missing")
	}

	name = strings.ToLower(name)
	for tld, id := range TLDMapping {
		if strings.HasSuffix(name, tld) {
			api := platform.NamingAPIs[id]
			result, err = api.Lookup(coins, name)
			if err != nil {
				return
			}
			return
		}
	}
	return nil, errors.E("name not found", errors.Params{"name": name})
}
