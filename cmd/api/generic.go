package api

import (
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas"
	"net/http"
)

func makeGenericAPI(router gin.IRouter, api blockatlas.TxAPI) {
	router.GET("/:address", func(c *gin.Context) {
		genericTxsByAddress(c, api)
	})
}

func genericTxsByAddress(c *gin.Context, api blockatlas.TxAPI) {
	address := c.Param("address")
	page, err := api.GetTxsByAddress(address)

	// Error responses
	switch {
	case err == blockatlas.ErrInvalidAddr:
		c.String(http.StatusBadRequest, "Invalid address")
		return
	case err == blockatlas.ErrNotFound:
		c.String(http.StatusNotFound, "No such address")
		return
	case err == blockatlas.ErrSourceConn:
		c.String(http.StatusServiceUnavailable, "Lost connection to blockchain")
		return
	case err != nil:
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	page.Sort()
	c.JSON(http.StatusOK, &page)
}
