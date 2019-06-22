package theta

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"net/http"
	"strconv"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client.BaseURL = viper.GetString("theta.api")
	p.client.HTTPClient = http.DefaultClient
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.THETA]
}

func (p *Platform) RegisterRoutes(router gin.IRouter) {
	router.GET("/:address", func(c *gin.Context) {
		p.getTransactions(c)
	})
}

// Get transactions for THETA address
func (p *Platform) getTransactions(c *gin.Context) {
	address := c.Param("address")
	token := c.Query("token")
	
	trx, err := p.client.FetchAddressTransactions(address)
	if apiError(c, err) {
		return
	}

	var txsNormalized []blockatlas.Tx
	for _, tr := range trx {
		if tr.Type == SendTransaction {
			if tx, ok := Normalize(&tr, address, token); ok && len(txsNormalized) < blockatlas.TxPerPage {
				txsNormalized = append(txsNormalized, tx)
			}
		}
	}

	page := blockatlas.TxPage(txsNormalized)
	page.Sort()
	c.JSON(http.StatusOK, &page)
}

func Normalize(trx *Tx, address, token string) (tx blockatlas.Tx, ok bool) {
	time, _ := strconv.ParseInt(trx.Timestamp, 10, 64)
	block, _ := strconv.ParseUint(trx.BlockHeight, 10, 64)

	tx = blockatlas.Tx{
		ID: trx.Hash,
		Coin: coin.THETA,
		Fee: blockatlas.Amount(trx.Data.Fee.Tfuelwei),
		Date:  time,
		Block: block,
		Sequence: block,
	}

	input := trx.Data.Inputs[0]
	output := trx.Data.Outputs[0]
	sequence, _ := strconv.ParseUint(input.Sequence, 10, 64)

	// Condition for transfer THETA trnafer
	if address != "" && token == "" && output.Coins.Tfuelwei == "0" {
		tx.From = input.Address
		tx.To = output.Address
		tx.Sequence = sequence
		tx.Type = blockatlas.TxTransfer
		tx.Meta = blockatlas.Transfer{
			Value: blockatlas.Amount(output.Coins.Thetawei),
		}

		return tx, true
	}

	// Condition for transfer Theta Fuel (TFUEL)
	if address != "" && token == "tfuel" && output.Coins.Thetawei == "0" {
		from := input.Address
		to := output.Address
		tx.From = from
		tx.To = to
		tx.Sequence = sequence
		tx.Type = blockatlas.TxNativeTokenTransfer
		tx.Meta = blockatlas.NativeTokenTransfer{
			Name: "Theta Fuel",
			Symbol: "TFUEL",
			TokenID: "tfuel",
			Decimals: 18,
			Value: blockatlas.Amount(output.Coins.Tfuelwei),
			From: from,
			To: to,
		}

		return tx, true
	}

	return tx, false
}

func apiError(c *gin.Context, err error) bool {
	if err != nil {
		logrus.WithError(err).Errorf("Unhandled error: %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return true
	}
	return false
}
