package stellar

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/util"
	"net/http"
	"strconv"
	"time"
)

type Platform struct {
	client Client
	CoinIndex uint
}

func (p *Platform) Init() error {
	handle := coin.Coins[p.CoinIndex].Handle
	p.client.API = viper.GetString(fmt.Sprintf("%s.api", handle))
	p.client.HTTP = &http.Client{
		Timeout: 2 * time.Second,
	}
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
	payments, err := p.client.GetTxsOfAddress(c.Param("address"))
	if apiError(c, err) {
		return
	}

	txs := make([]blockatlas.Tx, 0)
	for _, payment := range payments {
		tx, ok := Normalize(&payment, p.CoinIndex)
		if !ok {
			continue
		}
		txs = append(txs, tx)
	}

	page := blockatlas.TxPage(txs)
	page.Sort()
	c.JSON(http.StatusOK, &page)
}

func apiError(c *gin.Context, err error) bool {
	if err != nil {
		logrus.WithError(err).Warning("Stellar API request failed")
		c.String(http.StatusBadGateway, "Stellar API request failed")
		return true
	}
	return false
}

// Normalize converts a Stellar-based transaction into the generic model
func Normalize(payment *Payment, nativeCoinIndex uint) (tx blockatlas.Tx, ok bool) {
	switch payment.Type {
	case "payment":
		if payment.AssetType != "native" {
			return tx, false
		}
	case "create_account":
		break
	default:
		return tx, false
	}
	id, err := strconv.ParseUint(payment.ID, 10, 64)
	if err != nil {
		return tx, false
	}
	date, err := time.Parse("2006-01-02T15:04:05Z", payment.CreatedAt)
	if err != nil {
		return tx, false
	}
	var value, from, to string
	if payment.Amount != "" {
		value, err = util.DecimalToSatoshis(payment.Amount)
		from = payment.From
		to = payment.To
	} else if payment.StartingBalance != "" {
		value, err = util.DecimalToSatoshis(payment.StartingBalance)
		from = payment.Funder
		to = payment.Account
	} else {
		return tx, false
	}
	if err != nil {
		return tx, false
	}
	return blockatlas.Tx{
		ID:   payment.TransactionHash,
		Coin: nativeCoinIndex,
		From: from,
		To:   to,
		// https://www.stellar.org/developers/guides/concepts/fees.html
		// Fee fixed at 100 stroops
		Fee:   "100",
		Date:  date.Unix(),
		Block: id,
		Meta:  blockatlas.Transfer{
			Value: blockatlas.Amount(value),
		},
	}, true
}
