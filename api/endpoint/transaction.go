package endpoint

import (
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/api/model"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"net/http"
)

// @Summary Get Transactions
// @ID tx_v1
// @Description Get transactions from the address
// @Accept json
// @Produce json
// @Tags Transactions
// @Param coin path string true "the coin name" default(tezos)
// @Param address path string true "the query address" default(tz1WCd2jm4uSt4vntk4vSuUWoZQGhLcDuR9q)
// @Failure 500 {object} model.ErrorResponse
// @Router /v1/{coin}/{address} [get]
func GetTransactionsHistory(c *gin.Context, txAPI blockatlas.TxAPI, tokenTxAPI blockatlas.TokenTxAPI) {
	address := c.Param("address")
	if address == "" {
		c.JSON(http.StatusBadRequest, model.CreateErrorResponse(model.InvalidQuery, blockatlas.ErrInvalidAddr))
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
		c.JSON(http.StatusInternalServerError,
			model.CreateErrorResponse(model.InternalFail, errors.E("Failed to find api for that coin")))
		return
	}

	if err != nil {
		switch err {
		case blockatlas.ErrInvalidAddr:
			c.JSON(http.StatusBadRequest,
				model.CreateErrorResponse(model.InvalidQuery, blockatlas.ErrInvalidAddr))
			return
		case blockatlas.ErrNotFound:
			c.JSON(http.StatusNotFound,
				model.CreateErrorResponse(model.RequestedDataNotFound, blockatlas.ErrNotFound))
			return
		case blockatlas.ErrSourceConn:
			c.JSON(http.StatusServiceUnavailable,
				model.CreateErrorResponse(model.InternalFail, blockatlas.ErrSourceConn))
			return
		default:
			c.JSON(http.StatusInternalServerError,
				model.CreateErrorResponse(model.Default, err))
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
	if len(page) > blockatlas.TxPerPage {
		page = page[0:blockatlas.TxPerPage]
	}
	c.JSON(http.StatusOK, &page)
}

// @Summary Get Transactions by XPUB
// @ID txxpub_v1
// @Description Get transactions from XPUB address
// @Accept json
// @Produce json
// @Tags Transactions
// @Param coin path string true "the coin name" default(bitcoin)
// @Param xpub path string true "the xpub key" default(zpub6ruK9k6YGm8BRHWvTiQcrEPnFkuRDJhR7mPYzV2LDvjpLa5CuGgrhCYVZjMGcLcFqv9b2WvsFtY2Gb3xq8NVq8qhk9veozrA2W9QaWtihrC)
// @Failure 500 {object} model.ErrorResponse
// @Router /v1/{coin}/xpub/{xpub} [get]
func GetTransactionsByXpub(c *gin.Context, api blockatlas.TxUtxoAPI) {
	xPubKey := c.Param("xpub")
	if xPubKey == "" {
		c.JSON(http.StatusBadRequest, model.CreateErrorResponse(model.InvalidQuery, blockatlas.ErrInvalidKey))
		return
	}

	txs, err := api.GetTxsByXpub(xPubKey)
	if err != nil {
		switch err {
		case blockatlas.ErrInvalidKey:
			c.JSON(http.StatusBadRequest,
				model.CreateErrorResponse(model.InvalidQuery, blockatlas.ErrInvalidKey))
			return
		case blockatlas.ErrNotFound:
			c.JSON(http.StatusNotFound,
				model.CreateErrorResponse(model.RequestedDataNotFound, blockatlas.ErrNotFound))
			return
		case blockatlas.ErrSourceConn:
			c.JSON(http.StatusServiceUnavailable,
				model.CreateErrorResponse(model.InternalFail, blockatlas.ErrSourceConn))
			return
		default:
			c.JSON(http.StatusInternalServerError,
				model.CreateErrorResponse(model.Default, err))
			return
		}
	}
	var (
		filteredTxs = blockatlas.Txs(txs).FilterUniqueID().SortByDate()
		page        = blockatlas.TxPage(filteredTxs)
	)

	if len(page) > blockatlas.TxPerPage {
		page = page[0:blockatlas.TxPerPage]
	}
	c.JSON(http.StatusOK, &page)
}
