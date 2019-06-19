package binance

import (
	"github.com/trustwallet/blockatlas"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/util"
)

const Handle = "binance"

type Platform struct {
	client Client
}

func (p *Platform) Handle() string {
	return Handle
}

func (p *Platform) Init() error {
	p.client.BaseURL = viper.GetString("ontology.api")
	p.client.HTTPClient = http.DefaultClient
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.ONT]
}

func (p *Platform) RegisterRoutes(router gin.IRouter) {
	router.GET("/:address", func(c *gin.Context) {
		p.getTransactions(c)
	})
}

func (p *Platform) getTransactions(c *gin.Context) {
	token := c.Query("token")

	transactions, err := p.client.GetTxsOfAddress(c.Param("address"), token)
	if apiError(c, err) {
		return
	}

	var txs []blockatlas.Tx
	for _, srcTx := range transactions.Txs {
		tx, ok := Normalize(&srcTx, token)
		if !ok || len(txs) >= blockatlas.TxPerPage {
			continue
		}

		txs = append(txs, tx)
	}
	page := blockatlas.TxPage(txs)
	page.Sort()
	c.JSON(http.StatusOK, &page)
}

// Normalize converts a Binance transaction into the generic model
func Normalize(srcTx *Tx, token string) (tx blockatlas.Tx, ok bool) {
	value := util.DecimalExp(string(srcTx.Value), 8)
	fee := util.DecimalExp(string(srcTx.Fee), 8)

	tx = blockatlas.Tx{
		ID:    srcTx.Hash,
		Coin:  coin.BNB,
		Date:  srcTx.Timestamp / 1000,
		From:  srcTx.FromAddr,
		To:    srcTx.ToAddr,
		Fee:   blockatlas.Amount(fee),
		Block: srcTx.BlockHeight,
		Memo:  srcTx.Memo,
	}

	// Condition for native transfer (BNB)
	if srcTx.Asset == "BNB" && srcTx.Type == "TRANSFER" && token == "" {
		tx.Meta = blockatlas.Transfer{
			Value: blockatlas.Amount(value),
		}
		return tx, true
	}

	// Condiiton for native token transfer
	if srcTx.Asset == token && srcTx.Type == "TRANSFER" {
		tx.Meta = blockatlas.NativeTokenTransfer{
			TokenID:  srcTx.Asset,
			Symbol:   srcTx.MappedAsset,
			Value:    blockatlas.Amount(value),
			Decimals: 8,
			From:     srcTx.FromAddr,
			To:       srcTx.ToAddr,
		}

		return tx, true
	}

	return tx, false
}

func apiError(c *gin.Context, err error) bool {
	if err == blockatlas.ErrNotFound {
		c.String(http.StatusNotFound, err.Error())
		return true
	}
	if err == blockatlas.ErrInvalidAddr {
		c.String(http.StatusBadRequest, err.Error())
		return true
	}
	if err == blockatlas.ErrSourceConn {
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
