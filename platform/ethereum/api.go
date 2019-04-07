package ethereum

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/platform/ethereum/source"
	"github.com/trustwallet/blockatlas/util"
	"net/http"
	"strconv"
)

var client = source.Client{
	HttpClient: http.DefaultClient,
}

func Setup(router gin.IRouter) {
	router.Use(util.RequireConfig("ethereum.api"))
	router.Use(func(c *gin.Context) {
		client.RpcUrl = viper.GetString("ethereum.api")
		c.Next()
	})
	router.GET("/:address", getTransactions)
}

func getTransactions(c *gin.Context) {
	contract := c.Query("contract")
	var srcPage *source.Page
	var err error

	if contract != "" {
		srcPage, err = client.GetTxsWithContract(
			c.Param("address"), c.Query("contract"))
	} else {
		srcPage, err = client.GetTxs(c.Param("address"))
	}

	if apiError(c, err) {
		return
	}

	var txs []models.Tx
	if contract != "" {
		for _, srcTx := range srcPage.Docs {
			txs = AppendTokenTxs(txs, &srcTx)
		}
	} else {
		for _, srcTx := range srcPage.Docs {
			txs = AppendTxs(txs, &srcTx)
		}
	}

	page := models.Response(txs)
	page.Sort()
	c.JSON(http.StatusOK, &page)
}

func extractBase(srcTx *source.Doc) (base models.Tx, ok bool) {
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

	base = models.Tx{
		Id:     srcTx.Id,
		Coin:   srcTx.Coin, // SLIP-0044
		From:   srcTx.From,
		To:     srcTx.To,
		Fee:    models.Amount(srcTx.Gas),
		Date:   unix,
		Block:  srcTx.BlockNumber,
		Status: status,
		Error:  errReason,
	}
	return base, true
}

func AppendTxs(in []models.Tx, srcTx *source.Doc) (out []models.Tx) {
	out = in
	baseTx, ok := extractBase(srcTx)
	if !ok {
		return
	}

	// Native ETH transaction
	if srcTx.Value != "0" {
		transferTx := baseTx
		transferTx.Meta = models.Transfer{
			Value: models.Amount(srcTx.Value),
		}
		out = append(out, transferTx)
	}

	// Contract call
	if srcTx.Input != "" && srcTx.Input != "0x" {
		contractTx := baseTx
		contractTx.Meta = models.ContractCall{
			Input: srcTx.Input,
		}
		out = append(out, contractTx)
	}

	return
}

func AppendTokenTxs(in []models.Tx, srcTx *source.Doc) (out []models.Tx) {
	out = in
	baseTx, ok := extractBase(srcTx)
	if !ok {
		return
	}

	if len(srcTx.Ops) == 0 {
		return
	}
	op := &srcTx.Ops[0]

	if op.Type == "token_transfer" {
		tokenTx    := baseTx
		tokenTx.To   = op.To
		tokenTx.Meta = models.TokenTransfer{
			Name:     op.Contract.Name,
			Symbol:   op.Contract.Symbol,
			Contract: op.Contract.Address,
			Decimals: op.Contract.Decimals,
			Value:    models.Amount(op.Value),
		}
		out = append(out, tokenTx)
	}
	return
}

func apiError(c *gin.Context, err error) bool {
	if err != nil {
		logrus.WithError(err).Errorf("Unhandled error")
		c.AbortWithStatus(http.StatusInternalServerError)
		return true
	}
	return false
}
