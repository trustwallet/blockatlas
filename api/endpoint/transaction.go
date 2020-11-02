package endpoint

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"net/http"
)

// @Summary Get Transactions
// @ID tx_v2
// @Description Get transactions from the address
// @Accept json
// @Produce json
// @Tags Transactions
// @Param coin path string true "the coin name" default(tezos)
// @Param address path string true "the query address" default(tz1WCd2jm4uSt4vntk4vSuUWoZQGhLcDuR9q)
// @Failure 500 {object} ErrorResponse
// @Router /v1/{coin}/{address} [get]
// @Router /v2/{coin}/transactions/{address} [get]
func GetTransactionsHistory(c *gin.Context, txAPI blockatlas.TxAPI, tokenTxAPI blockatlas.TokenTxAPI) {
	address := c.Param("address")
	if address == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(blockatlas.ErrInvalidAddr))
		return
	}
	token := c.Query("token")

	var (
		txs []blockatlas.Tx
		err error
	)

	switch {
	case token == "" && txAPI != nil:
		txs, err = txAPI.GetTxsByAddress(address)
	case token != "" && tokenTxAPI != nil:
		txs, err = tokenTxAPI.GetTokenTxsByAddress(address, token)
	default:
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			errorResponse(errors.New("Failed to find api for that coin")),
		)
		return
	}

	if err != nil {
		switch err {
		case blockatlas.ErrInvalidAddr:
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				errorResponse(blockatlas.ErrInvalidAddr),
			)
			return
		case blockatlas.ErrNotFound:
			c.AbortWithStatusJSON(
				http.StatusNotFound,
				errorResponse(blockatlas.ErrNotFound),
			)
			return
		case blockatlas.ErrSourceConn:
			c.AbortWithStatusJSON(
				http.StatusServiceUnavailable,
				errorResponse(blockatlas.ErrSourceConn),
			)
			return
		default:
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				errorResponse(err),
			)
			return
		}
	}
	var (
		page        = make(blockatlas.TxPage, 0)
		filteredTxs = blockatlas.Txs(txs).FilterUniqueID().SortByDate()
	)
	for _, tx := range filteredTxs {
		tx.Direction = tx.GetTransactionDirection(address)
		page = append(page, tx)
	}

	page = page.FilterTransactionsByMemo()
	if token != "" {
		page = page.FilterTransactionsByToken(token)
	}

	if len(page) > blockatlas.TxPerPage {
		page = page[0:blockatlas.TxPerPage]
	}

	c.JSON(http.StatusOK, &page)
}

// @Summary Get Transactions by XPUB
// @ID tx_xpub_v2
// @Description Get transactions from XPUB address
// @Accept json
// @Produce json
// @Tags Transactions
// @Param coin path string true "the coin name" default(bitcoin)
// @Param xpub path string true "the xpub key" default(zpub6ruK9k6YGm8BRHWvTiQcrEPnFkuRDJhR7mPYzV2LDvjpLa5CuGgrhCYVZjMGcLcFqv9b2WvsFtY2Gb3xq8NVq8qhk9veozrA2W9QaWtihrC)
// @Failure 500 {object} ErrorResponse
// @Router /v1/{coin}/{address} [get]
// @Router /v2/{coin}/transactions/xpub/{xpub} [get]
func GetTransactionsByXpub(c *gin.Context, api blockatlas.TxUtxoAPI) {
	xPubKey := c.Param("xpub")
	if xPubKey == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(blockatlas.ErrInvalidKey))
		return
	}

	txs, err := api.GetTxsByXpub(xPubKey)
	if err != nil {
		switch err {
		case blockatlas.ErrInvalidKey:
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				errorResponse(blockatlas.ErrInvalidKey),
			)
			return
		case blockatlas.ErrNotFound:
			c.AbortWithStatusJSON(
				http.StatusNotFound,
				errorResponse(blockatlas.ErrNotFound),
			)
			return
		case blockatlas.ErrSourceConn:
			c.AbortWithStatusJSON(
				http.StatusServiceUnavailable,
				errorResponse(blockatlas.ErrSourceConn),
			)
			return
		default:
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				errorResponse(err),
			)
			return
		}
	}
	var (
		filteredTxs = blockatlas.Txs(txs).FilterUniqueID().SortByDate()
		page        = blockatlas.TxPage(filteredTxs)
	)
	page = page.FilterTransactionsByMemo()
	if len(page) > blockatlas.TxPerPage {
		page = page[0:blockatlas.TxPerPage]
	}

	c.JSON(http.StatusOK, &page)
}
