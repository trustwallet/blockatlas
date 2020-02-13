package bitcoin

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/pkg/numbers"
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
