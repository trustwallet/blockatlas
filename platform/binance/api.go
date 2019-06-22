package binance

import (
	"fmt"
	"github.com/trustwallet/blockatlas"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/util"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client.BaseURL = viper.GetString("binance.api")
	p.client.HTTPClient = http.DefaultClient
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.BNB]
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	// No native function to get height in explorer API
	// Workaround: Request list of blocks
	// and return number of the newest one
	list, err := p.client.GetBlockList(1)
	if err != nil {
		return 0, err
	}
	if len(list.BlockArray) == 0 {
		return 0, fmt.Errorf("no block descriptor found")
	}
	return list.BlockArray[0].BlockHeight, nil
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	srcTxs, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}
	// TODO: Only returns BNB transactions for now
	txs := NormalizeTxs(srcTxs.Txs, "")
	return &blockatlas.Block{
		Number: num,
		Txs:    txs,
	}, nil
}

func (p *Platform) RegisterRoutes(router gin.IRouter) {
	router.GET("/:address", func(c *gin.Context) {
		p.getTransactions(c)
	})
}

func (p *Platform) getTransactions(c *gin.Context) {
	token := c.Query("token")

	srcTxs, err := p.client.GetTxsOfAddress(c.Param("address"), token)
	if apiError(c, err) {
		return
	}
	txs := NormalizeTxs(srcTxs.Txs, token)

	page := blockatlas.TxPage(txs)
	page.Sort()
	c.JSON(http.StatusOK, &page)
}

// NormalizeTx converts a Binance transaction into the generic model
func NormalizeTx(srcTx *Tx, token string) (tx blockatlas.Tx, ok bool) {
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

	// Condition for native token transfer
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

// NormalizeTxs converts multiple Binance transactions
func NormalizeTxs(srcTxs []Tx, token string) (txs []blockatlas.Tx) {
	for _, srcTx := range srcTxs {
		tx, ok := NormalizeTx(&srcTx, token)
		if !ok || len(txs) >= blockatlas.TxPerPage {
			continue
		}
		txs = append(txs, tx)
	}
	return
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
