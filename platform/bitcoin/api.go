package bitcoin

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"net/http"
)

type Platform struct {
	client    Client
	CoinIndex uint
}

func (p *Platform) Init() error {
	p.client.URL = viper.GetString("bitcoin.api")
	p.client.HTTPClient = http.DefaultClient
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[p.CoinIndex]
}

func (p *Platform) RegisterRoutes(router gin.IRouter) {
	router.GET("/xpub/:key", func(c *gin.Context) {
		p.handleXpubRoute(c)
	})
	router.GET("/address/:address", func(c *gin.Context) {
		p.handleAddressRoute(c)
	})
}

func (p *Platform) handleAddressRoute(c *gin.Context) {
	address := c.Param("address")
	txs, ok := p.getTxsByAddress(address)
	txPage := blockatlas.TxPage(txs)
	txPage.Sort()
	if ok != nil {
		c.JSON(http.StatusInternalServerError, ok)
	} else {
		c.JSON(http.StatusOK, &txPage)
	}
}

func (p *Platform) handleXpubRoute(c *gin.Context) {
	xpub := c.Param("key")
	txs, ok := p.getTxsByXPub(xpub)
	txPage := blockatlas.TxPage(txs)
	txPage.Sort()
	if ok != nil {
		c.JSON(http.StatusInternalServerError, ok)
	} else {
		c.JSON(http.StatusOK, &txPage)
	}
}

func (p *Platform) getTxsByXPub(xpub string) ([]blockatlas.Tx, error) {
	sourceTxs, err := p.client.GetTransactionsByXpub(xpub)

	if err != nil {
		return []blockatlas.Tx{}, err
	}

	return NormalizeTxs(sourceTxs, p.CoinIndex), nil
}

func (p *Platform) getTxsByAddress(address string) ([]blockatlas.Tx, error) {
	sourceTxs, err := p.client.GetTransactions(address)
	if err != nil {
		return []blockatlas.Tx{}, err
	}

	return NormalizeTxs(sourceTxs, p.CoinIndex), nil
}

func NormalizeTxs(sourceTxs TransactionsList, coinIndex uint) []blockatlas.Tx {
	var txs []blockatlas.Tx
	for _, transaction := range sourceTxs.Transactions {
		if tx, ok := NormalizeTransfer(&transaction, coinIndex); ok {
			txs = append(txs, tx)
		}
	}
	return txs
}

func NormalizeTransfer(transaction *Transaction, coinIndex uint) (tx blockatlas.Tx, ok bool) {
	inputs := parseOutputs(transaction.Vin)
	outputs := parseOutputs(transaction.Vout)
	from := ""
	to := ""

	if len(inputs) > 0 {
		from = inputs[0]
	}

	if len(outputs) > 0 {
		to = outputs[0]
	}

	return blockatlas.Tx{
		ID:       transaction.ID,
		Coin:     coinIndex,
		From:     from,
		To:       to,
		Inputs:   inputs,
		Outputs:  outputs,
		Fee:      blockatlas.Amount(transaction.Fees),
		Date:     int64(transaction.BlockTime),
		Type:     blockatlas.TxTransfer,
		Block:    transaction.BlockHeight,
		Sequence: 0,
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(transaction.Value),
			Symbol:   coin.Coins[coinIndex].Symbol,
			Decimals: coin.Coins[coinIndex].Decimals,
		},
	}, true
}

func parseOutputs(outputs []Output) (addresses []string) {
	set := make(map[string]bool)
	result := []string{}
	for _, output := range outputs {
		for _, address := range output.Addresses {
			if set[address] {
				continue
			}
			set[address] = true
			result = append(result, address)
		}
	}
	return result
}
