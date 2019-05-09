package semux

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/util"
	"net/http"
	"strconv"
)

var client *Client

// Setup registers the semux route
func Setup(router gin.IRouter) {
	router.Use(util.RequireConfig("semux.api"))
	router.Use(withClient)
	router.GET("/:address", getTransactions)
}

func getTransactions(c *gin.Context) {
	s, err := client.GetTxsOfAddress(c.Param("address"))
	if apiError(c, err) {
		return
	}

	txs := make([]models.Tx, len(s))
	for i, srcTx := range s {
		txs[i] = Normalize(&srcTx)
	}
	page := models.Response(txs)
	page.Sort()
	c.JSON(http.StatusOK, &page)
}

func withClient(c *gin.Context) {
	rpcURL := viper.GetString("semux.api")
	if client == nil || rpcURL != client.BaseURL {
		logrus.WithField("rpc", rpcURL).Info("Created semux RPC client")
		client = &Client{BaseURL: rpcURL}
		client.Init()
	}
	c.Next()
}

// Normalize converts a semux transaction into the generic model
func Normalize(srcTx *Tx) (tx models.Tx) {
	logrus.Info(srcTx)
	blockNumber, err := strconv.ParseInt(srcTx.BlockNumber, 10, 64)
	if err != nil {
		logrus.Error("Failed to convert TX blockNumber for Semux API")
		return models.Tx{}
	}

	date, err := strconv.ParseInt(srcTx.Timestamp, 10, 64)
	if err != nil {
		logrus.Error("Failed to convert TX timestamp for Semux API")
		return models.Tx{}
	}

	return models.Tx{
		ID:    srcTx.Hash,
		Coin:  coin.SEM,
		Date:  date / 1000,
		From:  srcTx.From,
		To:    srcTx.To,
		Fee:   srcTx.Fee,
		Block: uint64(blockNumber),
		Meta: models.Transfer{
			Value: srcTx.Value,
		},
	}
}

func apiError(c *gin.Context, err error) bool {
	if err == models.ErrInvalidAddr {
		c.String(http.StatusBadRequest, err.Error())
		return true
	}
	if err == models.ErrInvalidAddr {
		c.String(http.StatusBadGateway, "semux RPC returned an error")
		return true
	}
	if err != nil {
		logrus.WithError(err).Errorf("Unhandled error: %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return true
	}
	return false
}
