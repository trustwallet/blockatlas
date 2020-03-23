package bitcoin

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/numbers"
	"net/http"
	"sort"
)

// @Summary Get xpub transactions
// @ID xpub
// @Description Get transactions from xpub address
// @Accept json
// @Produce json
// @Tags Transactions
// @Param coin path string true "the coin name" default(bitcoin)
// @Param xpub path string true "the xpub address" default(zpub6ruK9k6YGm8BRHWvTiQcrEPnFkuRDJhR7mPYzV2LDvjpLa5CuGgrhCYVZjMGcLcFqv9b2WvsFtY2Gb3xq8NVq8qhk9veozrA2W9QaWtihrC)
// @Success 200 {object} blockatlas.TxPage
// @Router /v1/{coin}/xpub/{xpub} [get]
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
	sort.Sort(txPage)
	if ok != nil {
		c.JSON(http.StatusInternalServerError, ok)
		return
	}
	c.JSON(http.StatusOK, &txPage)
}

func (p *Platform) handleXpubRoute(c *gin.Context) {
	xpub := c.Param("key")
	txs, ok := p.getTxsByXPub(xpub)
	txPage := blockatlas.TxPage(txs)
	sort.Sort(txPage)
	if ok != nil {
		c.JSON(http.StatusInternalServerError, ok)
		return
	}
	c.JSON(http.StatusOK, &txPage)
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

	txs := normalizeTxs(sourceTxs, p.CoinIndex, addressSet)
	return txs, nil
}

func (p *Platform) getTxsByAddress(address string) ([]blockatlas.Tx, error) {
	sourceTxs, err := p.client.GetTransactions(address)
	if err != nil {
		return []blockatlas.Tx{}, err
	}
	addressSet := mapset.NewSet()
	addressSet.Add(address)
	txs := normalizeTxs(sourceTxs, p.CoinIndex, addressSet)
	return txs, nil
}

func normalizeTxs(sourceTxs TransactionsList, coinIndex uint, addressSet mapset.Set) []blockatlas.Tx {
	var txs []blockatlas.Tx
	for _, transaction := range sourceTxs.TransactionList() {
		if tx, ok := normalizeTransfer(transaction, coinIndex, addressSet); ok {
			txs = append(txs, tx)
		}
	}
	return txs
}

func normalizeTransfer(transaction Transaction, coinIndex uint, addressSet mapset.Set) (tx blockatlas.Tx, ok bool) {
	tx = normalizeTransaction(transaction, coinIndex)
	direction := blockatlas.InferDirection(&tx, addressSet)
	value := blockatlas.InferValue(&tx, direction, addressSet)

	tx.Direction = direction
	tx.Meta = blockatlas.Transfer{
		Value:    value,
		Symbol:   coin.Coins[coinIndex].Symbol,
		Decimals: coin.Coins[coinIndex].Decimals,
	}

	return tx, true
}

func normalizeTransaction(tx Transaction, coinIndex uint) blockatlas.Tx {
	inputs := parseOutputs(tx.Vin)
	outputs := parseOutputs(tx.Vout)
	from := ""
	if len(inputs) > 0 {
		from = inputs[0].Address
	}

	to := ""
	if len(outputs) > 0 {
		to = outputs[0].Address
	}
	amount := blockatlas.Amount(tx.Amount())
	fees := blockatlas.Amount(numbers.GetAmountValue(tx.Fees))

	return blockatlas.Tx{
		ID:       tx.ID,
		Coin:     coinIndex,
		From:     from,
		To:       to,
		Inputs:   inputs,
		Outputs:  outputs,
		Fee:      fees,
		Date:     int64(tx.BlockTime),
		Type:     blockatlas.TxTransfer,
		Block:    tx.GetBlockHeight(),
		Status:   tx.getStatus(),
		Sequence: 0,
		Meta: blockatlas.Transfer{
			Value:    amount,
			Symbol:   coin.Coins[coinIndex].Symbol,
			Decimals: coin.Coins[coinIndex].Decimals,
		},
	}
}

func parseOutputs(outputs []Output) (addresses []blockatlas.TxOutput) {
	set := make(map[string]blockatlas.TxOutput)
	var ordered []string
	for _, output := range outputs {
		for _, address := range output.OutputAddress() {
			if val, ok := set[address]; ok {
				value := numbers.AddAmount(string(val.Value), output.Value)
				val.Value = blockatlas.Amount(value)
			} else {
				amount := numbers.GetAmountValue(output.Value)
				set[address] = blockatlas.TxOutput{
					Address: address,
					Value:   blockatlas.Amount(amount),
				}
				ordered = append(ordered, address)
			}
		}
	}
	for _, val := range ordered {
		addresses = append(addresses, set[val])
	}
	return addresses
}

func (transaction *Transaction) getStatus() blockatlas.Status {
	if transaction.Confirmations == 0 {
		return blockatlas.StatusPending
	}
	return blockatlas.StatusCompleted
}
