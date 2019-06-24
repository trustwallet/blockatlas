package waves

import (
	"fmt"
	"github.com/trustwallet/blockatlas/coin"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas"
)

type Platform struct {
	client    Client
	CoinIndex uint
	HandleStr string
}

func (p *Platform) Handle() string {
	return p.HandleStr
}

func (p *Platform) Init() error {
	p.client.BaseURL = viper.GetString(fmt.Sprintf("%s.api", p.HandleStr))
	p.client.HTTPClient = http.DefaultClient
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[p.CoinIndex]
}

func (p *Platform) RegisterRoutes(router gin.IRouter) {
	router.GET("/:address", func(c *gin.Context) {
		p.getTransactions(c)
	})
}

func (p *Platform) getTransactions(c *gin.Context) {
	address := c.Param("address")
	var err error

	addressTxs, err := p.client.GetTxs(address, 25, "")

	if apiError(c, err) {
		return
	}

	var txs []blockatlas.Tx
	for _, srcTx := range addressTxs {
		txs = AppendTxs(txs, &srcTx, p.CoinIndex)
	}

	page := blockatlas.TxPage(txs)
	page.Sort()
	c.JSON(http.StatusOK, &page)
}

func AppendTxs(in []blockatlas.Tx, srcTx *Transaction, coinIndex uint) (out []blockatlas.Tx) {
	out = in
	baseTx, ok := extractBase(srcTx, coinIndex)
	if !ok {
		return
	}

	// Waves transaction
	if len(srcTx.AssetId) == 0 {
		baseTx.Meta = blockatlas.Transfer{
			Value: blockatlas.Amount(strconv.Itoa(int(srcTx.Amount))),
		}
		out = append(out, baseTx)
	} else {
		// Token transaction
		baseTx.Meta = blockatlas.NativeTokenTransfer{
			Name:     srcTx.Asset.Description,
			Symbol:   srcTx.Asset.Name,
			TokenID:  srcTx.AssetId,
			Decimals: srcTx.Asset.Decimals,
			Value:    blockatlas.Amount(strconv.Itoa(int(srcTx.Amount))),
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
		Fee:    blockatlas.Amount(strconv.Itoa(int(srcTx.Fee))),
		Date:   int64(srcTx.Timestamp),
		Block:  srcTx.Block,
		Memo:   srcTx.Attachment,
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
