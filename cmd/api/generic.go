package api

import (
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas"
	"net/http"
)

func makeTxRoute(router gin.IRouter, api blockatlas.Platform) {
	var txAPI blockatlas.TxAPI
	var tokenTxAPI blockatlas.TokenTxAPI
	txAPI, _ = api.(blockatlas.TxAPI)
	tokenTxAPI, _ = api.(blockatlas.TokenTxAPI)

	if txAPI == nil && tokenTxAPI == nil {
		return
	}

	router.GET("/:address", func(c *gin.Context) {
		address := c.Param("address")
		if address == "" {
			emptyPage(c)
			return
		}
		token := c.Query("token")

		var page blockatlas.TxPage
		var err error
		switch {
		case token == "" && txAPI != nil:
			page, err = txAPI.GetTxsByAddress(address)
		case token != "" && tokenTxAPI != nil:
			page, err = tokenTxAPI.GetTokenTxsByAddress(address, token)
		default:
			emptyPage(c)
			return
		}

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
	})
}

func emptyPage(c *gin.Context) {
	c.JSON(http.StatusOK, blockatlas.TxPage(nil))
}
