package zilliqa

import (
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"

	"net/http"

	"github.com/spf13/viper"
)

const Handle = "zilliqa"

type Platform struct {
	client Client
}

func (p *Platform) Handle() string {
	return Handle
}

func (p *Platform) Init() error {
	p.client.BaseURL = viper.GetString("zilliqa.api")
	p.client.APIKey = viper.GetString("zilliqa.key")
	p.client.HTTPClient = http.DefaultClient
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.ZIL]
}

func (p *Platform) RegisterRoutes(router gin.IRouter) {
	router.GET("/:address", func(c *gin.Context) {
		p.GetTxsByAddress(c)
	})
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	var normalized []blockatlas.Tx
	txs, err := p.client.GetTxsOfAddress(address)

	if err != nil {
		return nil, err
	}

	for _, srcTx := range txs {
		tx := Normalize(&srcTx)
		if len(normalized) >= blockatlas.TxPerPage {
			break
		}
		normalized = append(normalized, tx)
	}

	return normalized, nil
}

func Normalize(srcTx *Tx) (tx blockatlas.Tx) {
	tx = blockatlas.Tx{
		ID:       srcTx.Hash,
		Coin:     coin.ZIL,
		Date:     srcTx.Timestamp / 1000,
		From:     srcTx.From,
		To:       srcTx.To,
		Fee:      blockatlas.Amount(srcTx.Fee),
		Block:    srcTx.BlockHeight,
		Sequence: srcTx.Nonce,
		Meta:     blockatlas.Transfer{Value: blockatlas.Amount(srcTx.Value)},
	}
	if !srcTx.ReceiptSuccess {
		tx.Status = blockatlas.StatusFailed
	}
	return tx
}
