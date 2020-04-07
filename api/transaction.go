package api

import (
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/api/middleware"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"net/http"
	"sort"
)

// @Summary Get Transactions
// @ID tx_v1
// @Description Get transactions from the address
// @Accept json
// @Produce json
// @Tags Transactions
// @Param coin path string true "the coin name" default(tezos)
// @Param address path string true "the query address" default(tz1WCd2jm4uSt4vntk4vSuUWoZQGhLcDuR9q)
// @Failure 500 {object} middleware.ApiError
// @Router /v1/{coin}/{address} [get]
func makeTxRouteV1(router gin.IRouter, api blockatlas.Platform) {
	makeTxRoute(router, api, "/:address")
}

// @Summary Get Transactions
// @ID tx_v2
// @Description Get transactions from the address
// @Accept json
// @Produce json
// @Tags Transactions
// @Param coin path string true "the coin name" default(tezos)
// @Param address path string true "the query address" default(tz1WCd2jm4uSt4vntk4vSuUWoZQGhLcDuR9q)
// @Success 200 {object} blockatlas.TxPage
// @Failure 500 {object} middleware.ApiError
// @Router /v2/{coin}/transactions/{address} [get]
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

		var txs []blockatlas.Tx
		var err error
		switch {
		case token == "" && txAPI != nil:
			txs, err = txAPI.GetTxsByAddress(address)
		case token != "" && tokenTxAPI != nil:
			txs, err = tokenTxAPI.GetTokenTxsByAddress(address, token)
		default:
			emptyPage(c)
			return
		}

		if err != nil {
			errResp := middleware.ErrorResponse(c)
			switch {
			case err == blockatlas.ErrInvalidAddr:
				errResp.Params(http.StatusBadRequest, "Invalid address")
			case err == blockatlas.ErrNotFound:
				errResp.Params(http.StatusNotFound, "No such address")
			case err == blockatlas.ErrSourceConn:
				errResp.Params(http.StatusServiceUnavailable, "Lost connection to blockchain")
			}
			errResp.Render()
			return
		}

		page := make(blockatlas.TxPage, 0)
		for _, tx := range txs {
			if tx.Direction != "" {
				goto AddTx
			}
			tx.Direction = blockatlas.DirectionOutgoing
			if tx.To == address {
				tx.Direction = blockatlas.DirectionIncoming
				if tx.From == address {
					tx.Direction = blockatlas.DirectionSelf
				}
			}
		AddTx:
			page = append(page, tx)
		}
		if len(page) > blockatlas.TxPerPage {
			page = page[0:blockatlas.TxPerPage]
		}
		sort.Sort(page)
		middleware.RenderSuccess(c, &page)
	})
}

// @Summary Get Tokens
// @ID tokens
// @Description Get tokens from the address
// @Accept json
// @Produce json
// @Tags Transactions
// @Param coin path string true "the coin name" default(ethereum)
// @Param address path string true "the query address" default(0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB)
// @Success 200 {object} blockatlas.CollectionPage
// @Failure 500 {object} middleware.ApiError
// @Router /v2/{coin}/tokens/{address} [get]
func makeTokenRoute(router gin.IRouter, api blockatlas.Platform) {
	var tokenAPI blockatlas.TokenAPI
	tokenAPI, _ = api.(blockatlas.TokenAPI)

	if tokenAPI == nil {
		return
	}

	router.GET("/tokens/:address", func(c *gin.Context) {
		address := c.Param("address")
		if address == "" {
			emptyPage(c)
			return
		}

		tl, err := tokenAPI.GetTokenListByAddress(address)
		if err != nil {
			middleware.ErrorResponse(c).Message(err.Error()).Render()
			return
		}

		middleware.RenderSuccess(c, blockatlas.DocsResponse{Docs: tl})
	})
}
