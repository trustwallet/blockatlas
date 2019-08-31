package observer

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/platform/bitcoin"
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
	emitted := make(map[string]bool)
	platform := bitcoin.UtxoPlatform(o.Coin)
	for _, sub := range subs {
		txs := txMap[sub.Address].Txs()
		for _, tx := range txs {
			if _, ok := emitted[tx.ID]; ok {
				continue
			}
			xpub, _ := o.Storage.GetXpubFromAddress(o.Coin, sub.Address)
			if len(xpub) != 0 {
				addressSet := mapset.NewSet(xpub)
				direction := platform.InferDirection(&tx, addressSet)
				value := platform.InferValue(&tx, direction, addressSet)

				tx.Direction = direction
				tx.Meta = blockatlas.Transfer{
					Value:    value,
					Symbol:   coin.Coins[o.Coin].Symbol,
					Decimals: coin.Coins[o.Coin].Decimals,
				}
			}
			emitted[tx.ID] = true
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
