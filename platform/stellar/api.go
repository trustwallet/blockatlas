package stellar

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stellar/go/clients/horizon"
	"net/http"
	"strconv"
	"sync"
	"time"
	"trustwallet.com/blockatlas/models"
	"trustwallet.com/blockatlas/util"
)

var client = horizon.Client {
	HTTP: http.DefaultClient,
}

func Setup(router gin.IRouter) {
	router.Use(util.RequireConfig("stellar.api"))
	router.Use(func(c *gin.Context) {
		client.URL = viper.GetString("stellar.api")
		c.Next()
	})
	router.GET("/:address", getTransactions)
}


func getTransactions(c *gin.Context) {
	acc, err := client.LoadAccount(c.Param("address"))
	if apiError(c, err) {
		return
	}

	ctxt, _ := context.WithTimeout(context.Background(), time.Second)

	var txMut sync.Mutex
	var txs []models.BasicTx

	err = client.StreamTransactions(ctxt, acc.ID, nil, func(tx horizon.Transaction) {
		txMut.Lock()
		defer txMut.Unlock()

		txs = append(txs, models.BasicTx{
			Kind:  models.TxBasic,
			Id:    tx.Hash,
			From:  tx.Account,
			To:    "TODO",
			Fee:   strconv.FormatInt(int64(tx.FeePaid), 10),
			Value: "TODO",
		})
	})
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	<-ctxt.Done()
	txMut.Lock()
	c.JSON(http.StatusOK, txs)
	txMut.Unlock()
}

func apiError(c *gin.Context, err error) bool {
	if err != nil {
		logrus.WithError(err).Warning("Stellar API request failed")
		c.String(http.StatusTeapot, "error: todo more descriptive")
		return true
	}
	return false
}
