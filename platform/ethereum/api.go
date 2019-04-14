package ethereum

import (
	"math/big"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/util"
)

func MakeSetup(coinIndex uint, platform string) func(gin.IRouter) {
	apiKey := platform + ".api"

	client := Client{
		HttpClient: http.DefaultClient,
	}

	return func(router gin.IRouter) {
		router.Use(util.RequireConfig(apiKey))
		router.Use(func(c *gin.Context) {
			client.RpcUrl = viper.GetString(apiKey)
			c.Next()
		})
		router.GET("/:address", func(c *gin.Context) {
			GetTransactions(c, coinIndex, &client)
		})
	}
}

func GetTransactions(c *gin.Context, coinIndex uint, client *Client) {
	token := c.Query("token")
	var srcPage *Page
	var err error

	if token != "" {
		srcPage, err = client.GetTxsWithContract(
			c.Param("address"), token)
	} else {
		srcPage, err = client.GetTxs(c.Param("address"))
	}

	if apiError(c, err) {
		return
	}

	var txs []models.Tx
	for _, srcTx := range srcPage.Docs {
		txs = AppendTxs(txs, &srcTx, coinIndex)
	}

	page := models.Response(txs)
	page.Sort()
	c.JSON(http.StatusOK, &page)
}

func extractBase(srcTx *Doc, coinIndex uint) (base models.Tx, ok bool) {
	var status, errReason string
	if srcTx.Error == "" {
		status = models.StatusCompleted
	} else {
		status = models.StatusFailed
		errReason = srcTx.Error
	}

	unix, err := strconv.ParseInt(srcTx.TimeStamp, 10, 64)
	if err != nil {
		return base, false
	}

	fee := calcFee(srcTx.GasPrice, srcTx.GasUsed)

	base = models.Tx{
		Id:       srcTx.Id,
		Coin:     coinIndex,
		From:     srcTx.From,
		To:       srcTx.To,
		Fee:      models.Amount(fee),
		Date:     unix,
		Block:    srcTx.BlockNumber,
		Status:   status,
		Error:    errReason,
		Sequence: srcTx.Nonce,
	}
	return base, true
}

func AppendTxs(in []models.Tx, srcTx *Doc, coinIndex uint) (out []models.Tx) {
	out = in
	baseTx, ok := extractBase(srcTx, coinIndex)
	if !ok {
		return
	}

	// Native ETH transaction
	if len(srcTx.Ops) == 0 && srcTx.Input == "0x" {
		transferTx := baseTx
		transferTx.Meta = models.Transfer{
			Value: models.Amount(srcTx.Value),
		}
		out = append(out, transferTx)
	}

	// Smart Contract Call
	if len(srcTx.Ops) == 0 && srcTx.Input != "0x" {
		contractTx := baseTx
		contractTx.Meta = models.ContractCall{
			Input: srcTx.Input,
			Value: srcTx.Value,
		}
		out = append(out, contractTx)
	}

	if len(srcTx.Ops) == 0 {
		return
	}
	op := &srcTx.Ops[0]

	if op.Type == "token_transfer" {
		tokenTx := baseTx
		tokenTx.To = op.To

		tokenTx.Meta = models.TokenTransfer{
			Name:     op.Contract.Name,
			Symbol:   op.Contract.Symbol,
			TokenID:  op.Contract.Address,
			Decimals: op.Contract.Decimals,
			Value:    models.Amount(op.Value),
			From:     op.From,
			To:       op.To,
		}
		out = append(out, tokenTx)
	}
	return
}

func calcFee(gasPrice string, gasUsed string) string {
	var gasPriceBig, gasUsedBig, feeBig big.Int

	gasPriceBig.SetString(gasPrice, 10)
	gasUsedBig.SetString(gasUsed, 10)

	feeBig.Mul(&gasPriceBig, &gasUsedBig)

	return feeBig.String()
}

func apiError(c *gin.Context, err error) bool {
	if err != nil {
		logrus.WithError(err).Errorf("Unhandled error")
		c.AbortWithStatus(http.StatusInternalServerError)
		return true
	}
	return false
}
