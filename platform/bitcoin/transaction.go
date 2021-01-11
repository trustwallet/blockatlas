package bitcoin

import (
	"sort"

	"github.com/trustwallet/blockatlas/platform/bitcoin/blockbook"

	mapset "github.com/deckarep/golang-set"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/numbers"
	"github.com/trustwallet/golibs/txtype"
)

func (p *Platform) GetTxsByAddress(address string) (txtype.TxPage, error) {
	txs, err := p.getTxsByAddress(address)
	if err != nil {
		return nil, err
	}
	txPage := txtype.TxPage(txs)
	sort.Sort(txPage)
	return txPage, nil
}

func (p *Platform) GetTxsByXpub(xpub string) (txtype.TxPage, error) {
	txs, err := p.getTxsByXpub(xpub)
	if err != nil {
		return nil, err
	}
	txPage := txtype.TxPage(txs)
	sort.Sort(txPage)
	return txPage, nil
}

func (p *Platform) getTxsByXpub(xpub string) ([]txtype.Tx, error) {
	sourceTxs, err := p.client.GetTransactionsByXpub(xpub)

	if err != nil {
		return []txtype.Tx{}, err
	}

	addressSet := mapset.NewSet()
	for _, token := range sourceTxs.Tokens {
		addressSet.Add(token.Name)
	}

	txs := normalizeTxs(sourceTxs, p.CoinIndex, addressSet)
	return txs, nil
}

func (p *Platform) getTxsByAddress(address string) ([]txtype.Tx, error) {
	sourceTxs, err := p.client.GetTxs(address)
	if err != nil {
		return []txtype.Tx{}, err
	}
	addressSet := mapset.NewSet()
	addressSet.Add(address)
	txs := normalizeTxs(sourceTxs, p.CoinIndex, addressSet)
	return txs, nil
}

func normalizeTxs(sourceTxs blockbook.TransactionsList, coinIndex uint, addressSet mapset.Set) []txtype.Tx {
	var txs []txtype.Tx
	for _, transaction := range sourceTxs.TransactionList() {
		if tx, ok := normalizeTransfer(transaction, coinIndex, addressSet); ok {
			txs = append(txs, tx)
		}
	}
	return txs
}

func normalizeTransfer(transaction blockbook.Transaction, coinIndex uint, addressSet mapset.Set) (tx txtype.Tx, ok bool) {
	tx = normalizeTransaction(transaction, coinIndex)
	direction := txtype.InferDirection(&tx, addressSet)
	value := txtype.InferValue(&tx, direction, addressSet)

	tx.Direction = direction
	tx.Meta = txtype.Transfer{
		Value:    value,
		Symbol:   coin.Coins[coinIndex].Symbol,
		Decimals: coin.Coins[coinIndex].Decimals,
	}

	return tx, true
}

func normalizeTransaction(tx blockbook.Transaction, coinIndex uint) txtype.Tx {
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
	amount := txtype.Amount(tx.Amount())
	fees := txtype.Amount(numbers.GetAmountValue(tx.Fees))

	return txtype.Tx{
		ID:       tx.ID,
		Coin:     coinIndex,
		From:     from,
		To:       to,
		Inputs:   inputs,
		Outputs:  outputs,
		Fee:      fees,
		Date:     tx.BlockTime,
		Type:     txtype.TxTransfer,
		Block:    tx.GetBlockHeight(),
		Status:   tx.GetStatus(),
		Sequence: 0,
		Meta: txtype.Transfer{
			Value:    amount,
			Symbol:   coin.Coins[coinIndex].Symbol,
			Decimals: coin.Coins[coinIndex].Decimals,
		},
	}
}

func parseOutputs(outputs []blockbook.Output) (addresses []txtype.TxOutput) {
	set := make(map[string]*txtype.TxOutput)
	var ordered []string
	for _, output := range outputs {
		for _, address := range output.OutputAddress() {
			if val, ok := set[address]; ok {
				value := numbers.AddAmount(string(val.Value), output.Value)
				val.Value = txtype.Amount(value)
			} else {
				amount := numbers.GetAmountValue(output.Value)
				set[address] = &txtype.TxOutput{
					Address: address,
					Value:   txtype.Amount(amount),
				}
				ordered = append(ordered, address)
			}
		}
	}
	for _, val := range ordered {
		addresses = append(addresses, *set[val])
	}
	return addresses
}
