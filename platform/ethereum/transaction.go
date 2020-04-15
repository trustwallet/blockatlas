package ethereum

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
)

func (p *Platform) RegisterRoutes(router gin.IRouter) {
	router.GET("/:address", func(c *gin.Context) {
		p.getTransactions(c)
	})
}

func (p *Platform) getTransactions(c *gin.Context) {
	token := c.Query("token")
	address := c.Param("address")
	var page blockatlas.TxPage
	var err error

	if token != "" {
		page, err = p.client.GetTokenTxs(address, token, p.CoinIndex)
	} else {
		page, err = p.client.GetTransactions(address, p.CoinIndex)
	}

	if apiError(c, err) {
		return
	}

	sort.Sort(page)
	c.JSON(http.StatusOK, &page)
}

func apiError(c *gin.Context, err error) bool {
	if err != nil {
		logger.Error(err, "Unhandled error")
		c.AbortWithStatus(http.StatusInternalServerError)
		return true
	}
	return false
}
