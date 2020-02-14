package bitcoin

import (
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/pkg/numbers"
	"net/http"
	"sync"
)

type Platform struct {
	client    Client
	CoinIndex uint
}

func (p *Platform) Init() error {
	p.client = Client{blockatlas.InitClient(viper.GetString(p.ConfigKey()))}
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[p.CoinIndex]
}

func (p *Platform) ConfigKey() string {
	return fmt.Sprintf("%s.api", p.Coin().Handle)
}

// @Summary Get xpub transactions
// @ID xpub
// @Description Get transactions from xpub address
// @Accept json
// @Produce json
// @Tags Platform-Transactions
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
	txPage.Sort()
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
	txPage.Sort()
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

func (p *Platform) GetAddressesFromXpub(xpub string) ([]string, error) {
	tokens, err := p.client.GetAddressesFromXpub(xpub)
	addresses := make([]string, 0)
	for _, token := range tokens {
		addresses = append(addresses, token.Name)
	}
	return addresses, err
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	status, err := p.client.GetBlockNumber()
	return status.Backend.Blocks, err
}

func (p *Platform) GetAllBlockPages(total, num int64) []Transaction {
	txs := make([]Transaction, 0)
	if total <= 1 {
		return txs
	}

	start := int64(1)
	var wg sync.WaitGroup
	out := make(chan TransactionsList, int(total-start))
	wg.Add(int(total - start))
	for start < total {
		start++
		go func(page, num int64, out chan TransactionsList, wg *sync.WaitGroup) {
			defer wg.Done()
			block, err := p.client.GetTransactionsByBlock(num, page)
			if err != nil {
				logger.Error("GetTransactionsByBlockChan", err, logger.Params{"number": num, "page": page})
				return
			}
			out <- block
		}(start, num, out, &wg)
	}
	wg.Wait()
	close(out)
	for r := range out {
		txs = append(txs, r.TransactionList()...)
	}
	return txs
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	page := int64(1)
	block, err := p.client.GetTransactionsByBlock(num, page)
	if err != nil {
		return nil, err
	}
	txPages := p.GetAllBlockPages(block.TotalPages, num)
	txs := append(txPages, block.TransactionList()...)
	var normalized []blockatlas.Tx
	for _, tx := range txs {
		normalized = append(normalized, normalizeTransaction(tx, p.CoinIndex))
	}
	return &blockatlas.Block{
		Number: num,
		ID:     block.Hash,
		Txs:    normalized,
	}, nil
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

func normalizeTransfer(transaction Transaction, coinIndex uint, addressSet mapset.Set) (tx blockatlas.Tx, ok bool) {
	tx = normalizeTransaction(transaction, coinIndex)
	direction := InferDirection(&tx, addressSet)
	value := InferValue(&tx, direction, addressSet)

	tx.Direction = direction
	tx.Meta = blockatlas.Transfer{
		Value:    value,
		Symbol:   coin.Coins[coinIndex].Symbol,
		Decimals: coin.Coins[coinIndex].Decimals,
	}

	return tx, true
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

func InferDirection(tx *blockatlas.Tx, addressSet mapset.Set) blockatlas.Direction {
	inputSet := mapset.NewSet()
	for _, address := range tx.Inputs {
		inputSet.Add(address.Address)
	}
	outputSet := mapset.NewSet()
	for _, address := range tx.Outputs {
		outputSet.Add(address.Address)
	}
	intersect := addressSet.Intersect(inputSet)
	if intersect.Cardinality() == 0 {
		return blockatlas.DirectionIncoming
	}
	if outputSet.IsProperSubset(addressSet) || outputSet.Equal(inputSet) {
		return blockatlas.DirectionSelf
	}
	return blockatlas.DirectionOutgoing
}

func InferValue(tx *blockatlas.Tx, direction blockatlas.Direction, addressSet mapset.Set) blockatlas.Amount {
	value := blockatlas.Amount("0")
	if len(tx.Outputs) == 0 {
		return value
	}
	if direction == blockatlas.DirectionOutgoing || direction == blockatlas.DirectionSelf {
		value = tx.Outputs[0].Value
	} else if direction == blockatlas.DirectionIncoming {
		amount := value
		for _, output := range tx.Outputs {
			if !addressSet.Contains(output.Address) {
				continue
			}
			value := numbers.AddAmount(string(amount), string(output.Value))
			amount = blockatlas.Amount(value)
		}
		value = amount
	}
	return value
}

func (transaction *Transaction) getStatus() blockatlas.Status {
	if transaction.Confirmations == 0 {
		return blockatlas.StatusPending
	}
	return blockatlas.StatusCompleted
}
