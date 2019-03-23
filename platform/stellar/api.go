package stellar

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stellar/go/clients/horizon"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/platform/stellar/source"
	"github.com/trustwallet/blockatlas/util"
	"net/http"
	"strconv"
	"time"
)

var stellarClient = source.Client{
	Client: horizon.Client{
		HTTP: &http.Client{
			Timeout: 2 * time.Second,
		},
	},
}

func Setup(router gin.IRouter) {
	router.Use(util.RequireConfig("stellar.api"))
	router.Use(func(c *gin.Context) {
		stellarClient.URL = viper.GetString("stellar.api")
		c.Next()
	})
	router.GET("/:address", func(c *gin.Context) {
		GetTransactions(c, &stellarClient)
	})
}

func GetTransactions(c *gin.Context, client *source.Client) {
	s, err := client.GetTxsOfAddress(c.Param("address"))
	if apiError(c, err) {
		return
	}

	txs := make([]models.Tx, 0)
	for _, srcTx := range s {
		txs = append(txs, models.Tx{
			Id:   srcTx.Tx.Hash,
			Date: srcTx.Tx.LedgerCloseTime.Unix(),
			From: srcTx.Tx.Account,
			To:   srcTx.Payment.Destination.Address(),
			Fee:  strconv.FormatInt(int64(srcTx.Tx.FeePaid), 10),
			Meta: models.Transfer{
				Name:     coin.XLM.Title,
				Symbol:   coin.XLM.Symbol,
				Decimals: coin.XLM.Decimals,
				Value:    strconv.FormatInt(int64(srcTx.Payment.Amount), 10),
			},
		})
	}

	page := models.Response(txs)
	page.Sort()
	c.JSON(http.StatusOK, &page)
}

func apiError(c *gin.Context, err error) bool {
	if hErr, ok := err.(*horizon.Error); ok {
		if hErr.Problem.Type == "https://stellar.org/horizon-errors/bad_request" {
			c.String(http.StatusBadRequest, "Bad request!")
			return true
		} else {
			c.String(http.StatusBadRequest, hErr.Problem.Type)
			return true
		}
	}
	if err != nil {
		logrus.WithError(err).Warning("Stellar API request failed")
		c.String(http.StatusBadGateway, "Stellar API request failed")
		return true
	}
	return false
}
