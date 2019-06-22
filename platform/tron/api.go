package tron

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"net/http"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client.BaseURL = viper.GetString("tron.api")
	p.client.Token = viper.GetString("tron.token")
	p.client.HTTPClient = http.DefaultClient
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.TRX]
}

func (p *Platform) RegisterRoutes(router gin.IRouter) {
	router.GET("/:address", func(c *gin.Context) {
		p.getTransactions(c)
	})
}

/// Normalize converts a Tron transaction into the generic model
func Normalize(srcTx *Tx) (tx blockatlas.Tx, ok bool) {
	if len(srcTx.Data.Contracts) < 1 {
		return tx, false
	}

	// TODO Support multiple transfers in a single transaction
	contract := &srcTx.Data.Contracts[0]
	switch contract.Parameter.(type) {
	case TransferContract:
		transfer := contract.Parameter.(TransferContract)

		from, err := HexToAddress(transfer.Value.OwnerAddress)
		if err != nil {
			return tx, false
		}
		to, err := HexToAddress(transfer.Value.ToAddress)
		if err != nil {
			return tx, false
		}

		return blockatlas.Tx{
			ID:   srcTx.ID,
			Coin: coin.TRX,
			Date: srcTx.Data.Timestamp / 1000,
			From: from,
			To:   to,
			Fee:  "0",
			Meta: blockatlas.Transfer{
				Value: transfer.Value.Amount,
			},
		}, true
	default:
		return tx, false
	}
}

func (p *Platform) getTransactions(c *gin.Context) {
	srcTxs, err := p.client.GetTxsOfAddress(c.Param("address"))
	if err != nil {
		logrus.WithError(err).
			Errorf("Tron: Failed to get transactions for %s", c.Param("address"))
		// TODO AbortWithError
	}

	var txs []blockatlas.Tx
	for _, srcTx := range srcTxs {
		tx, ok := Normalize(&srcTx)
		if ok {
			txs = append(txs, tx)
		}
	}

	page := blockatlas.TxPage(txs)
	page.Sort()
	c.JSON(http.StatusOK, &page)
}
