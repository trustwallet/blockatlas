package ethereum

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/platform/ethereum/source"
	"github.com/trustwallet/blockatlas/util"
	"net/http"
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
	router.GET("/:address/:token", getTransactionsOfContract)
}

func getTransactions(c *gin.Context) {
	page, err := client.GetTxs(c.Param("address"))
	sendResult(c, page, err)
}

func getTransactionsOfContract(c *gin.Context) {
	page, err := client.GetTxsWithContract(
		c.Param("address"), c.Query("contract"))
	sendResult(c, page, err)
}

func sendResult(c *gin.Context, page *source.Page, err error) {
	if apiError(c, err) {
		return
	}

	var txs []models.Tx
	for _, srcTx := range page.Docs {
		if len(srcTx.Ops) < 1 {
			continue
		}
		op := &srcTx.Ops[0]
		switch op.Type {
		case "token_transfer":
			baseTx := models.Tx{
				Id:    op.TxId,
				Coin:  op.Coin, // SLIP-0044
				From:  op.From,
				To:    op.To,
				Fee:   models.Amount(srcTx.Gas),
				Date:  srcTx.TimeStamp,
				Block: srcTx.BlockNumber,
			}

			tokenTx    := baseTx
			contractTx := baseTx

			tokenTx.Meta = models.TokenTransfer{
				Name:     op.Contract.Name,
				Symbol:   op.Contract.Symbol,
				Contract: op.Contract.Address,
				Decimals: op.Contract.Decimals,
				Value:    models.Amount(op.Value),
			}

			contractTx.Meta = models.ContractCall(ethContractMeta{
				Contract: op.Contract.Address,
				Input:    srcTx.Input,
			})

			txs = append(txs, tokenTx, contractTx)
		}
	}
}

func apiError(c *gin.Context, err error) bool {
	if err != nil {
		logrus.WithError(err).Errorf("Unhandled error: %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return true
	}
	return false
}

// TODO @vikmeup, let's refactor this to models.ContractCall?
type ethContractMeta struct {
	Contract string `json:"contract"`
	Input    string `json:"input"`
}
