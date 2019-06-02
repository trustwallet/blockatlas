package binance

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/util"
)

var client = Client{
	HTTPClient: http.DefaultClient,
}

// Setup registers the Binance DEX route
func Setup(router gin.IRouter) {
	router.Use(util.RequireConfig("binance.api"))
	router.Use(util.RequireConfig("binance.rpc"))
	router.Use(func(c *gin.Context) {
		client.ExplorerBaseURL = viper.GetString("binance.api")
		client.RPCBaseURL = viper.GetString("binance.rpc")
		c.Next()
	})
	router.GET("/:address", getTransactions)
}

func getTransactions(c *gin.Context) {
	token := c.Query("token")
	address := c.Param("address")

	transactions, err := client.GetTxsOfAddress(address, token)
	if apiError(c, err) {
		return
	}

	var txs []models.Tx
	for _, srcTx := range transactions.Txs {
		tx, ok := Normalize(&srcTx, token, address)
		if !ok || len(txs) >= models.TxPerPage {
			continue
		}

		txs = append(txs, tx)
	}
	page := models.Response(txs)
	page.Sort()
	c.JSON(http.StatusOK, &page)
}

// Normalize converts a Binance transaction into the generic model
func Normalize(srcTx *Tx, token, address string) (tx models.Tx, ok bool) {
	hash := srcTx.Hash
	value := util.DecimalExp(string(srcTx.Value), 8)
	fee := util.DecimalExp(string(srcTx.Fee), 8)

	tx = models.Tx{
		ID:    hash,
		Coin:  coin.BNB,
		Date:  srcTx.Timestamp / 1000,
		Fee:   models.Amount(fee),
		Block: srcTx.BlockHeight,
		Memo:  srcTx.Memo,
	}

	// Condition for native transfer (BNB)
	if srcTx.Asset == "BNB" && srcTx.Type == "TRANSFER" && token == "" {
		tx.From = srcTx.FromAddr
		tx.To = srcTx.ToAddr
		tx.Meta = models.Transfer{
			Value: models.Amount(value),
		}
		return tx, true
	}

	// Condiiton for native token transfer
	if srcTx.Asset == token && srcTx.Type == "TRANSFER" {
		tx.From = srcTx.FromAddr
		tx.To = srcTx.ToAddr
		tx.Meta = models.NativeTokenTransfer{
			TokenID:  srcTx.Asset,
			Symbol:   srcTx.MappedAsset,
			Value:    models.Amount(value),
			Decimals: 8,
			From:     srcTx.FromAddr,
			To:       srcTx.ToAddr,
		}

		return tx, true
	}

	
	// Condition for native transfer
	// Condition for multisend
	receipt, _ := client.getTransactionReceipt(hash)

	outputs := receipt.TxReceipts.Value.Msg[0].MsgValue.Outputs
	zeroInput := receipt.TxReceipts.Value.Msg[0].MsgValue.Inputs[0]
	zeroOutputAdress := outputs[0].Address
	if (srcTx.FromAddr == "" || srcTx.ToAddr == "") && zeroInput.Coins[0].Denom == "BNB" {
		if zeroInput.Address == address {
			tx.From = address
			tx.To = zeroOutputAdress  // Pick 0 index as main receipient
			tx.Meta = models.Transfer{
				Value: models.Amount(zeroInput.Coins[0].Amount),
			}
			return tx, true
		}

		tx.To = address
		tx.From = zeroOutputAdress
		tx.Meta = models.Transfer{
			Value: models.Amount(getAmount(outputs, address)),
		}
		return tx, true
	}

	return tx, false
}

func getAmount(outputs []Output, address string) string {
	for _, out := range outputs {
		if out.Address == address {
			return out.Coins[0].Amount
		}
	}

	return "0"
}

func apiError(c *gin.Context, err error) bool {
	if err == models.ErrNotFound {
		c.String(http.StatusNotFound, err.Error())
		return true
	}
	if err == models.ErrInvalidAddr {
		c.String(http.StatusBadRequest, err.Error())
		return true
	}
	if err == models.ErrSourceConn {
		c.String(http.StatusBadGateway, "connection to Binance API failed")
		return true
	}
	if _, ok := err.(*Error); ok {
		c.String(http.StatusBadGateway, "Binance API returned an error")
		return true
	}
	if err != nil {
		logrus.WithError(err).Errorf("Unhandled error: %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return true
	}
	return false
}
