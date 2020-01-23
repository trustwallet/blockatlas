package observer

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/platform/bitcoin"
	"github.com/trustwallet/blockatlas/storage"
)

type Event struct {
	Subscription blockatlas.Subscription
	Tx           *blockatlas.Tx
}

type Observer struct {
	Storage storage.Addresses
	Coin    uint
}

func (o *Observer) Execute(blocks <-chan *blockatlas.Block) <-chan Event {
	events := make(chan Event)
	go o.run(events, blocks)
	return events
}

func (o *Observer) run(events chan<- Event, blocks <-chan *blockatlas.Block) {
	for block := range blocks {
		o.processBlock(events, block)
	}
}

func (o *Observer) processBlock(events chan<- Event, block *blockatlas.Block) {
	txMap := GetTxs(block)
	if len(txMap) == 0 {
		return
	}

	// Build list of unique addresses
	var addresses []string
	for address := range txMap {
		if len(address) == 0 {
			continue
		}
		addresses = append(addresses, address)
	}

	// Lookup subscriptions
	subs, err := o.Storage.Lookup(o.Coin, addresses)
	if err != nil || len(subs) == 0 {
		return
	}
	for _, sub := range subs {
		tx, ok := txMap[sub.Address]
		if !ok {
			continue
		}
		for _, tx := range tx.Txs() {
			tx.Direction = getDirection(tx, sub.Address)
			inferUtxoValue(&tx, sub.Address, o.Coin)
			events <- Event{
				Subscription: sub,
				Tx:           &tx,
			}
		}
	}
}

func GetTxs(block *blockatlas.Block) map[string]*blockatlas.TxSet {
	txMap := make(map[string]*blockatlas.TxSet)
	for i := 0; i < len(block.Txs); i++ {
		addresses := block.Txs[i].GetAddresses()
		addresses = append(addresses, block.Txs[i].GetUtxoAddresses()...)
		for _, address := range addresses {
			if txMap[address] == nil {
				txMap[address] = new(blockatlas.TxSet)
			}
			txMap[address].Add(&block.Txs[i])
		}
	}
	return txMap
}

func getDirection(tx blockatlas.Tx, address string) blockatlas.Direction {
	if len(tx.Inputs) > 0 && len(tx.Outputs) > 0 {
		addressSet := mapset.NewSet(address)
		return bitcoin.InferDirection(&tx, addressSet)
	}
	if address == tx.To {
		if tx.From == tx.To {
			return blockatlas.DirectionSelf
		}
		return blockatlas.DirectionIncoming
	}
	return blockatlas.DirectionOutgoing
}

func inferUtxoValue(tx *blockatlas.Tx, address string, coinIndex uint) {
	if len(tx.Inputs) > 0 && len(tx.Outputs) > 0 {
		addressSet := mapset.NewSet(address)
		value := bitcoin.InferValue(tx, tx.Direction, addressSet)
		tx.Meta = blockatlas.Transfer{
			Value:    value,
			Symbol:   coin.Coins[coinIndex].Symbol,
			Decimals: coin.Coins[coinIndex].Decimals,
		}
	}
}
