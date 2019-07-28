package observer

import (
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas"
)

type Event struct {
	Subscription Subscription
	Tx           *blockatlas.Tx
}

type Observer struct {
	Storage Storage
	Coin    uint
}

func (o *Observer) Execute(blocks <-chan *blockatlas.Block) <-chan Event {
	if o.Coin == 0 {
		panic("coin ID not set")
	}
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
	// Order transactions in block by addresses
	txMap := make(map[string][]*blockatlas.Tx)
	for _, tx := range block.Txs {
		txMap[tx.From] = append(txMap[tx.From], &tx)
		txMap[tx.To] = append(txMap[tx.To], &tx)
	}

	// Build list of unique addresses
	var addresses []string
	for address := range txMap {
		addresses = append(addresses, address)
	}

	// Lookup subscriptions
	subs, err := o.Storage.Lookup(o.Coin, addresses...)
	if err != nil {
		logrus.WithError(err).Error("Failed to look up subscriptions")
		return
	}

	// Emit events
	for _, sub := range subs {
		txs := txMap[sub.Address]
		for _, tx := range txs {
			events <- Event{
				Subscription: sub,
				Tx:           tx,
			}
		}
	}
}
