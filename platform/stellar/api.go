package stellar

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/util"
	"net/http"
	"strconv"
	"time"
)

func MakeSetup(coinIndex uint, platform string) func(gin.IRouter) {
	return func(router gin.IRouter) {
		apiKey := platform + ".api"

		stellarClient := Client{
			HTTP: &http.Client{
				Timeout: 2 * time.Second,
			},
		}

		router.Use(util.RequireConfig(apiKey))
		router.Use(func(c *gin.Context) {
			stellarClient.API = viper.GetString(apiKey)
			c.Next()
		})
		router.GET("/:address", func(c *gin.Context) {
			GetTransactions(c, coinIndex, &stellarClient)
		})
	}
}

func GetTransactions(c *gin.Context, coinIndex uint, client *Client) {
	payments, err := client.GetTxsOfAddress(c.Param("address"))
	if apiError(c, err) {
		return
	}

	txs := make([]models.Tx, 0)
	for _, payment := range payments {
		tx, ok := Normalize(&payment, coinIndex)
		if !ok {
			continue
		}
		txs = append(txs, tx)
	}

	page := models.Response(txs)
	page.Sort()
	c.JSON(http.StatusOK, &page)
}

func apiError(c *gin.Context, err error) bool {
	if err != nil {
		logrus.WithError(err).Warning("Stellar API request failed")
		c.String(http.StatusBadGateway, "Stellar API request failed")
		return true
	}
	return false
}

func Normalize(payment *Payment, nativeCoinIndex uint) (tx models.Tx, ok bool) {
	switch payment.Type {
	case "payment":
		if payment.AssetType != "native" {
			return tx, false
		}
	case "create_account":
		break
	default:
		return tx, false
	}
	id, err := strconv.ParseUint(payment.ID, 10, 64)
	if err != nil {
		return tx, false
	}
	date, err := time.Parse("2006-01-02T15:04:05Z", payment.CreatedAt)
	if err != nil {
		return tx, false
	}
	var value, from, to string
	if payment.Amount != "" {
		value, err = util.DecimalToSatoshis(payment.Amount)
		from = payment.From
		to = payment.To
	} else if payment.StartingBalance != "" {
		value, err = util.DecimalToSatoshis(payment.StartingBalance)
		from = payment.Funder
		to = payment.Account
	} else {
		return tx, false
	}
	if err != nil {
		return tx, false
	}
	return models.Tx{
		Id:    payment.TransactionHash,
		Coin:  nativeCoinIndex,
		From:  from,
		To:    to,
		// https://www.stellar.org/developers/guides/concepts/fees.html
		// Fee fixed at 100 stroops
		Fee:   "100",
		Date:  date.Unix(),
		Block: id,
		Meta:  models.Transfer{
			Value: models.Amount(value),
		},
	}, true
}
