package observer

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
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
	if len(txMap) == 0 {
		return
	}

	// Build list of unique addresses
	var addresses []string
	for address := range txMap {
		addresses = append(addresses, address)
	}
	// Lookup subscriptions
	subs, err := o.Storage.Lookup(o.Coin, addresses...)
	if err != nil {
		return
	}

	// Emit events
	emittedUtxo := make(map[string]blockatlas.Direction)
	platform := bitcoin.UtxoPlatform(o.Coin)
	for _, sub := range subs {

		txs := txMap[sub.Address].Txs()
		for _, tx := range txs {

			xpub, err := o.Storage.GetXpubFromAddress(o.Coin, sub.Address)
			if err == nil && len(xpub) > 0 {
				xpubAddresses, err := o.Storage.GetAddressFromXpub(o.Coin, xpub)
				if err != nil {
					continue
				}
				addressSet := mapset.NewSet()
				for _, addr := range xpubAddresses {
					addressSet.Add(addr)
				}
				direction := platform.InferDirection(&tx, addressSet)
				value := platform.InferValue(&tx, direction, addressSet)

				tx.Direction = direction
				tx.Meta = blockatlas.Transfer{
					Value:    value,
					Symbol:   coin.Coins[o.Coin].Symbol,
					Decimals: coin.Coins[o.Coin].Decimals,
				}

				if d, ok := emittedUtxo[tx.ID]; ok {
					if d == tx.Direction || d == blockatlas.DirectionSelf {
						continue
					}
					emittedUtxo[tx.ID] = blockatlas.DirectionSelf
				} else {
					emittedUtxo[tx.ID] = tx.Direction
				}
			}
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
