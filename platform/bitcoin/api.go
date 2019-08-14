package bitcoin

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/deckarep/golang-set"
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

	addressSet := mapset.NewSet()
	for _, token := range sourceTxs.Tokens {
		addressSet.Add(token.Name)
	}

	txs := NormalizeTxs(sourceTxs, p.CoinIndex, addressSet)
	return txs, nil
}

func (p *Platform) getTxsByAddress(address string) ([]blockatlas.Tx, error) {
	sourceTxs, err := p.client.GetTransactions(address)
	if err != nil {
		return []blockatlas.Tx{}, err
	}
	addressSet := mapset.NewSet()
	addressSet.Add(address)
	txs := NormalizeTxs(sourceTxs, p.CoinIndex, addressSet)
	return txs, nil
}

func NormalizeTxs(sourceTxs TransactionsList, coinIndex uint, addressSet mapset.Set) []blockatlas.Tx {
	var txs []blockatlas.Tx
	for _, transaction := range sourceTxs.Transactions {
		if tx, ok := NormalizeTransfer(&transaction, coinIndex); ok {
			tx.Direction = inferDirection(&tx, addressSet)
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

func inferDirection(tx *blockatlas.Tx, addressSet mapset.Set) (string) {
	inputSet := mapset.NewSet()
	for _, address := range tx.Inputs {
		inputSet.Add(address)
	}
	outputSet := mapset.NewSet()
	for _, address := range tx.Outputs {
		outputSet.Add(address)
	}
	intersect := addressSet.Intersect(inputSet)
	if intersect.Cardinality() == 0 {
		return blockatlas.DirectionIncoming
	} else {
		if outputSet.IsProperSubset(addressSet) {
			return blockatlas.DirectionSelf
		} else {
			return blockatlas.DirectionOutgoing
		}
	}
}
