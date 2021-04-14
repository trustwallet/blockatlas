package bitcoin

import (
	"sort"

	"github.com/trustwallet/blockatlas/platform/bitcoin/blockbook"

	mapset "github.com/deckarep/golang-set"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/numbers"
	"github.com/trustwallet/golibs/types"
)

func (p *Platform) GetTxsByAddress(address string) (types.Txs, error) {
	txs, err := p.getTxsByAddress(address)
	if err != nil {
		return nil, err
	}
	sort.Sort(txs)
	return txs, nil
}

func (p *Platform) GetTxsByXpub(xpub string) (types.Txs, error) {
	txs, err := p.getTxsByXpub(xpub)
	if err != nil {
		return nil, err
	}
	sort.Sort(txs)
	return txs, nil
}

func (p *Platform) getTxsByXpub(xpub string) (types.Txs, error) {
	sourceTxs, err := p.client.GetTransactionsByXpub(xpub)

	if err != nil {
		return types.Txs{}, err
	}

	addressSet := mapset.NewSet()
	for _, token := range sourceTxs.Tokens {
		addressSet.Add(token.Name)
	}

	txs := normalizeTxs(sourceTxs, p.CoinIndex, addressSet)
	return txs, nil
}

func (p *Platform) getTxsByAddress(address string) (types.Txs, error) {
	sourceTxs, err := p.client.GetTxs(address)
	if err != nil {
		return types.Txs{}, err
	}
	addressSet := mapset.NewSet()
	addressSet.Add(address)
	txs := normalizeTxs(sourceTxs, p.CoinIndex, addressSet)
	return txs, nil
}

func normalizeTxs(sourceTxs blockbook.TransactionsList, coinIndex uint, addressSet mapset.Set) types.Txs {
	var txs types.Txs
	for _, transaction := range sourceTxs.TransactionList() {
		if tx, ok := normalizeTransfer(transaction, coinIndex, addressSet); ok {
			txs = append(txs, tx)
		}
	}
	return txs
}

func normalizeTransfer(transaction blockbook.Transaction, coinIndex uint, addressSet mapset.Set) (tx types.Tx, ok bool) {
	tx = normalizeTransaction(transaction, coinIndex)
	direction := types.InferDirection(&tx, addressSet)
	value := types.InferValue(&tx, direction, addressSet)

	tx.Direction = direction
	tx.Meta = types.Transfer{
		Value:    value,
		Symbol:   coin.Coins[coinIndex].Symbol,
		Decimals: coin.Coins[coinIndex].Decimals,
	}

	return tx, true
}

func normalizeTransaction(tx blockbook.Transaction, coinIndex uint) types.Tx {
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
	amount := types.Amount(tx.Amount())
	fees := types.Amount(numbers.GetAmountValue(tx.Fees))

	return types.Tx{
		ID:       tx.ID,
		Coin:     coinIndex,
		From:     from,
		To:       to,
		Inputs:   inputs,
		Outputs:  outputs,
		Fee:      fees,
		Date:     tx.BlockTime,
		Type:     types.TxTransfer,
		Block:    tx.GetBlockHeight(),
		Status:   tx.GetStatus(),
		Sequence: 0,
		Meta: types.Transfer{
			Value:    amount,
			Symbol:   coin.Coins[coinIndex].Symbol,
			Decimals: coin.Coins[coinIndex].Decimals,
		},
	}
}

func parseOutputs(outputs []blockbook.Output) (addresses []types.TxOutput) {
	set := make(map[string]*types.TxOutput)
	var ordered []string
	for _, output := range outputs {
		for _, address := range output.OutputAddress() {
			if val, ok := set[address]; ok {
				value := numbers.AddAmount(string(val.Value), output.Value)
				val.Value = types.Amount(value)
			} else {
				amount := numbers.GetAmountValue(output.Value)
				set[address] = &types.TxOutput{
					Address: address,
					Value:   types.Amount(amount),
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
