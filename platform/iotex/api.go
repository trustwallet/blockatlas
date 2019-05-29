package iotex

import(
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/util"
)

var client = Client{
	HTTPClient : http.DefaultClient,
}

// Setup registers the Iotex chain route
func Setup(router gin.IRouter) {
	router.Use(util.RequireConfig("iotex.api"))
	router.Use(func(c *gin.Context) {
		client.BaseURL = viper.GetString("iotex.api")
		c.Next()
	})
	router.GET("/:address", getTransactions)
}

func getTransactions(c *gin.Context) {
	trxs, err := client.GetTxsOfAddress(c.Param("address"))
	if apiError(c, err) {
		return
	}

	var txs []models.Tx
	for _, srcTx := range trxs.ActionInfo {
		tx, ok := Normalize(srcTx)
		if !ok || len(txs) >= models.TxPerPage {
			continue
		}
		txs = append(txs, tx)
	}
	page := models.Response(txs)
	page.Sort()
	c.JSON(http.StatusOK, &page)
}

// Normalize converts an Iotex transaction into the generic model
func Normalize(trx *ActionInfo) (models.Tx, bool) {
	date, err := time.Parse(time.RFC3339, trx.Timestamp)
	if err != nil {
		return models.Tx{
			Coin: coin.IOTX,
			Status: models.StatusFailed,
			Error: err.Error(),
		}, false
	}
	height, err := strconv.ParseInt(trx.BlkHeight, 10, 64)
	if err != nil {
		return models.Tx{
			Coin: coin.IOTX,
			Status: models.StatusFailed,
			Error: err.Error(),
		}, false
	}
	if height <= 0 {
		return models.Tx{
			Coin: coin.IOTX,
			Status: models.StatusFailed,
			Error: "invalid block height",
		}, false
	}
	nonce, err := strconv.ParseInt(trx.Action.Core.Nonce, 10, 64)
	if err != nil {
		return models.Tx{
			Coin: coin.IOTX,
			Status: models.StatusFailed,
			Error: err.Error(),
		}, false
	}

	return models.Tx{
		ID       : trx.ActHash,
		Coin     : coin.IOTX,
		From     : trx.Sender,
		To       : trx.Action.Core.Transfer.Recipient,
		Fee      : models.Amount(trx.GasFee),
		Date     : date.Unix(),
		Block    : uint64(height),
		Status   : models.StatusCompleted,
		Sequence : uint64(nonce),
		Type     : models.TxTransfer,
		Meta     : models.Transfer{
			Value : models.Amount(trx.Action.Core.Transfer.Amount),
		},
	}, true
}

func apiError(c *gin.Context, err error) bool {
	if err == models.ErrSourceConn {
		c.String(http.StatusBadGateway, "connection to IoTeX API failed")
		return true
	}
	if err == models.ErrNotFound {
		c.String(http.StatusNotFound, err.Error())
		return true
	}
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return true
	}
	return false
}
