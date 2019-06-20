package waves

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/util"
)

// MakeSetup returns a function used to register an Waves-based platform route
func MakeSetup(coinIndex uint, platform string) func(gin.IRouter) {
	apiKey := platform + ".api"

	client := Client{
		HTTPClient: http.DefaultClient,
	}

	return func(router gin.IRouter) {
		router.Use(util.RequireConfig(apiKey))
		router.Use(func(c *gin.Context) {
			client.BaseURL = viper.GetString(apiKey)
			c.Next()
		})
		router.GET("/:address", func(c *gin.Context) {
			GetTransactions(c, coinIndex, &client)
		})
	}
}

func GetTransactions(c *gin.Context, coinIndex uint, client *Client) {
	address := c.Param("address")
	var err error

	addressTxs, err := client.GetTxs(address, 25, "")

	if apiError(c, err) {
		return
	}

	var txs []blockatlas.Tx
	for _, srcTx := range addressTxs {
		// support only transfer transaction
		if srcTx.Type == 4 {
			txs = AppendTxs(txs, &srcTx, coinIndex, client)
		}
	}

	page := blockatlas.TxPage(txs)
	page.Sort()
	c.JSON(http.StatusOK, &page)
}

func AppendTxs(in []blockatlas.Tx, srcTx *Transaction, coinIndex uint, client *Client) (out []blockatlas.Tx) {
	out = in
	baseTx, ok := extractBase(srcTx, coinIndex)
	if !ok {
		return
	}

	// Waves transaction
	if len(srcTx.AssetId) == 0 {
		baseTx.Meta = blockatlas.Transfer{
			Value: blockatlas.Amount(srcTx.Amount),
		}
		out = append(out, baseTx)
	} else {
		// Token transaction
		tokenInfo, err := client.GetTokenInfo(srcTx.AssetId)
		if err != nil {
			return
		}
		baseTx.Meta = blockatlas.NativeTokenTransfer{
			Name:     tokenInfo.Description,
			Symbol:   tokenInfo.Name,
			TokenID:  srcTx.AssetId,
			Decimals: tokenInfo.Decimals,
			Value:    blockatlas.Amount(srcTx.Amount),
			From:     srcTx.Sender,
			To:       srcTx.Recipient,
		}
		out = append(out, baseTx)
	}
	return
}

func extractBase(srcTx *Transaction, coinIndex uint) (base blockatlas.Tx, ok bool) {
	var status string
	status = blockatlas.StatusCompleted

	base = blockatlas.Tx{
		ID:     srcTx.Id,
		Coin:   coinIndex,
		From:   srcTx.Sender,
		To:     srcTx.Recipient,
		Fee:    blockatlas.Amount(srcTx.Fee),
		Date:   int64(srcTx.Timestamp),
		Block:  srcTx.Block,
		Status: status,
	}
	return base, true
}

func apiError(c *gin.Context, err error) bool {
	if err != nil {
		logrus.WithError(err).Errorf("Unhandled error")
		c.AbortWithStatus(http.StatusInternalServerError)
		return true
	}
	return false
}
