package theta

import(
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/util"
	"net/http"
	"strconv"
)

var client = Client{
	HTTPClient: http.DefaultClient,
}

// Setup registers for THETA route
func Setup(router gin.IRouter) {
	router.Use(util.RequireConfig("theta.api"))
	router.Use(func(c *gin.Context) {
		client.BaseURL = viper.GetString("theta.api")
		c.Next()
	})
	router.GET("/:address", getTransactions)
}

// Get transactions for THETA address
func getTransactions(c *gin.Context) {
	address := c.Param("address")
	token := c.Query("token")
	
	trx, err := client.FetchAddressTransactions(address)
	if apiError(c, err) {
		return
	}

	var txsNormalized []models.Tx
	for _, tr := range trx {
		if tr.Type == SendTransaction {
			if tx, ok := Normalize(&tr, address, token); ok && len(txsNormalized) < models.TxPerPage {
				txsNormalized = append(txsNormalized, tx)
			}
		}
	}

	page := models.Response(txsNormalized)
	page.Sort()
	c.JSON(http.StatusOK, &page)
}

func Normalize(trx *Tx, address, token string) (tx models.Tx, ok bool) {
	time, _ := strconv.ParseInt(trx.Timestamp, 10, 64)
	block, _ := strconv.ParseUint(trx.BlockHeight, 10, 64)

	tx = models.Tx{
		ID: trx.Hash,
		Coin: coin.THETA,
		Fee: models.Amount(trx.Data.Fee.Tfuelwei),
		Date:  time,
		Block: block,
		Sequence: block,
	}

	input := trx.Data.Inputs[0]
	output := trx.Data.Outputs[0]
	sequence, _ := strconv.ParseUint(input.Sequence, 10, 64)

	// Condition for transfer THETA trnafer
	if address != "" && token == "" && output.Coins.Tfuelwei == "0" {
		tx.From = input.Address
		tx.To = output.Address
		tx.Sequence = sequence
		tx.Type = models.TxTransfer
		tx.Meta = models.Transfer{
			Value: models.Amount(output.Coins.Thetawei),
		}

		return tx, true
	}

	// Condition for transfer Theta Fuel (TFUEL)
	if address != "" && token == "tfuel" && output.Coins.Thetawei == "0" {
		from := input.Address
		to := output.Address
		tx.From = from
		tx.To = to
		tx.Sequence = sequence
		tx.Type = models.TxNativeTokenTransfer
		tx.Meta = models.NativeTokenTransfer{
			Name: "Theta Fuel",
			Symbol: "TFUEL",
			TokenID: "tfuel",
			Decimals: 18,
			Value: models.Amount(output.Coins.Tfuelwei),
			From: from,
			To: to,
		}

		return tx, true
	}

	return tx, false
}

func apiError(c *gin.Context, err error) bool {
	if err != nil {
		logrus.WithError(err).Errorf("Unhandled error: %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return true
	}
	return false
}