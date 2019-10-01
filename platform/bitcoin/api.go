package bitcoin

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"net/http"
	"strconv"
	"sync"

	mapset "github.com/deckarep/golang-set"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	//"sync"
)

type Platform struct {
	client    Client
	CoinIndex uint
}

func UtxoPlatform(index uint) *Platform {
	platform := &Platform{CoinIndex: index}
	err := platform.Init()
	if err != nil {
		logger.Panic("UtxoPlatform index error", err, logger.Params{"index": index})
	}
	return platform
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
// @Tags platform,tx
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

	txs := p.NormalizeTxs(sourceTxs, p.CoinIndex, addressSet)
	return txs, nil
}

func (p *Platform) getTxsByAddress(address string) ([]blockatlas.Tx, error) {
	sourceTxs, err := p.client.GetTransactions(address)
	if err != nil {
		return []blockatlas.Tx{}, err
	}
	addressSet := mapset.NewSet()
	addressSet.Add(address)
	txs := p.NormalizeTxs(sourceTxs, p.CoinIndex, addressSet)
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

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	block, err := p.client.GetTransactionsByBlock(num, 1)
	if err != nil {
		return nil, err
	}
	if block.Page < block.TotalPages {
		var wg sync.WaitGroup
		out := make(chan Block)
		for i := int64(2); i <= block.TotalPages; i++ {
			go p.client.GetTransactionsByBlockChan(num, i, out, &wg)
		}

		wg.Wait()
		defer close(out)
		for r := range out {
			block.Transactions = append(block.Transactions, r.Transactions...)
		}
	}
	var normalized []blockatlas.Tx
	for _, tx := range block.Transactions {
		normalized = append(normalized, NormalizeTransaction(&tx, p.CoinIndex))
	}
	return &blockatlas.Block{
		Number: num,
		ID:     block.Hash,
		Txs:    normalized,
	}, nil
}

func (p *Platform) NormalizeTxs(sourceTxs TransactionsList, coinIndex uint, addressSet mapset.Set) []blockatlas.Tx {
	var txs []blockatlas.Tx
	for _, transaction := range sourceTxs.Transactions {
		if tx, ok := p.NormalizeTransfer(&transaction, coinIndex, addressSet); ok {
			txs = append(txs, tx)
		}
	}
	return txs
}

func NormalizeTransaction(transaction *Transaction, coinIndex uint) blockatlas.Tx {
	inputs := parseOutputs(transaction.Vin)
	outputs := parseOutputs(transaction.Vout)
	from := ""
	if len(inputs) > 0 {
		from = inputs[0].Address
	}

	to := ""
	if len(outputs) > 0 {
		to = outputs[0].Address
	}

	status := blockatlas.StatusCompleted
	if transaction.Confirmations == 0 {
		status = blockatlas.StatusPending
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
		Status:   status,
		Sequence: 0,
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(transaction.Value),
			Symbol:   coin.Coins[coinIndex].Symbol,
			Decimals: coin.Coins[coinIndex].Decimals,
		},
	}
}

func (p *Platform) NormalizeTransfer(transaction *Transaction, coinIndex uint, addressSet mapset.Set) (tx blockatlas.Tx, ok bool) {
	tx = NormalizeTransaction(transaction, coinIndex)
	direction := p.InferDirection(&tx, addressSet)
	value := p.InferValue(&tx, direction, addressSet)

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
		for _, address := range output.Addresses {
			if val, ok := set[address]; ok {
				val.Value = AddAmount(string(val.Value), output.Value)
			} else {
				set[address] = blockatlas.TxOutput{
					Address: address,
					Value:   blockatlas.Amount(output.Value),
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

func AddAmount(left string, right string) (sum blockatlas.Amount) {
	amount1, _ := strconv.ParseInt(left, 10, 64)
	amount2, _ := strconv.ParseInt(right, 10, 64)
	return blockatlas.Amount(strconv.FormatInt(amount1+amount2, 10))
}

func (p *Platform) InferDirection(tx *blockatlas.Tx, addressSet mapset.Set) blockatlas.Direction {
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
	} else {
		if outputSet.IsProperSubset(addressSet) {
			return blockatlas.DirectionSelf
		} else {
			return blockatlas.DirectionOutgoing
		}
	}
}

func (p *Platform) InferValue(tx *blockatlas.Tx, direction blockatlas.Direction, addressSet mapset.Set) blockatlas.Amount {
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
			amount = AddAmount(string(amount), string(output.Value))
		}
		value = amount
	}
	return value
}
