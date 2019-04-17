package icon

import(
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/util"
	"net/http"
	"time"
	"fmt"
)

var client = Client{
	HTTPClient : http.DefaultClient,
}

func Setup(router gin.IRouter) {
	router.Use(util.RequireConfig("icon.api"))
	router.Use(func(c *gin.Context) {
		client.RPCUrl = viper.GetString("icon.api")
	})
	router.GET("/:address", getTransactions)
}

func getTransactions(c *gin.Context) {
	trxs, _ := client.GetAddressTransactions(c.Param("address"))

	nTrxs := make([]models.Tx, 0)
	for _, trx := range trxs {
		nTrx, ok := Normalize(&trx)
		if !ok {
			continue
		}
		nTrxs = append(nTrxs, nTrx)
	}

	page := models.Response(nTrxs)
	page.Sort()
	c.JSON(http.StatusOK, &page)
}

func Normalize(trx *Tx) (tx models.Tx, b bool) {
	date, err := time.Parse("2006-01-02T15:04:05.999Z0700", trx.CreateDate)
	if err != nil {
		fmt.Printf("%v\n", err)
		return tx, false
	}

	return models.Tx{
		Id     : trx.TxHash,
		Coin   : coin.ICX,
		From   : trx.FromAddr,
		To     : trx.ToAddr,
		Fee    : models.Amount(trx.Fee),
		Status : models.StatusCompleted,
		Date   : date.Unix(),
		Type   : models.TxTransfer,
		Block  : trx.Height,
		Meta: models.Transfer{
			Value : models.Amount(trx.Amount),
		},
	}, true
}