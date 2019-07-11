package api

import (
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas"
	"net/http"
)

func makeTxRouteV1(router gin.IRouter, api blockatlas.Platform) {
	makeTxRoute(router, api, "/:address")
}

func makeTxRouteV2(router gin.IRouter, api blockatlas.Platform) {
	makeTxRoute(router, api, "/transactions/:address")
}

func makeTxRoute(router gin.IRouter, api blockatlas.Platform, path string) {
	var txAPI blockatlas.TxAPI
	var tokenTxAPI blockatlas.TokenTxAPI
	txAPI, _ = api.(blockatlas.TxAPI)
	tokenTxAPI, _ = api.(blockatlas.TokenTxAPI)

	if txAPI == nil && tokenTxAPI == nil {
		return
	}

	router.GET(path, func(c *gin.Context) {
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

func makeStakingRoute(router gin.IRouter, api blockatlas.Platform) {
	var stakingAPI blockatlas.StakeAPI
	stakingAPI, _ = api.(blockatlas.StakeAPI)

	if stakingAPI == nil {
		return
	}

	router.GET("/staking/validators", func(c *gin.Context) {

		validators, err := stakingAPI.GetValidators()

		if err != nil {
			c.JSON(http.StatusServiceUnavailable, err)
		}

		c.JSON(http.StatusOK, blockatlas.DocsResponse{Docs: validators})
	})
}

func makeCollectionRoute(router gin.IRouter, api blockatlas.Platform) {
	var collectionAPI blockatlas.CollectionAPI
	collectionAPI, _ = api.(blockatlas.CollectionAPI)

	if collectionAPI == nil {
		return
	}

	router.GET("/collections/:owner/", func(c *gin.Context) {
		collections, err := collectionAPI.GetCollections(c.Param("owner"))

		if err != nil {
			c.JSON(http.StatusServiceUnavailable, err)
		}

		c.JSON(http.StatusOK, collections)
	})

	router.GET("/collections/:owner/collection/:contract", func(c *gin.Context) {
		collectibles, err := collectionAPI.GetCollectibles(c.Param("owner"), c.Param("contract"))

		if err != nil {
			c.JSON(http.StatusServiceUnavailable, err)
		}

		c.JSON(http.StatusOK, collectibles)
	})
}

func emptyPage(c *gin.Context) {
	var page blockatlas.TxPage
	c.JSON(http.StatusOK, &page)
}
