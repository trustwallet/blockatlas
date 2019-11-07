package observer

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/platform/bitcoin"
	"github.com/trustwallet/blockatlas/storage"
)

type Event struct {
	Subscription storage.Subscription
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
	xpubs := make(map[string][]string)
	for address := range txMap {
		if len(address) == 0 {
			continue
		}
		// Verify we already have this xpub
		xpub, xpubAddresses, err := o.Storage.GetXpubFromAddress(o.Coin, address)
		if err == nil && len(xpub) > 0 {
			// Add xpub in addresses list for lookup
			addresses = append(addresses, xpub)
			// Temp cache for xpub addresses
			xpubs[xpub] = xpubAddresses
			// Save txMap for this xpub
			txMap[xpub] = txMap[address]
			continue
		}
		addresses = append(addresses, address)
	}

	// Lookup subscriptions
	subs, err := o.Storage.Lookup(o.Coin, addresses)
	if err != nil || len(subs) == 0 {
		return
	}

	// Emit events
	emittedUtxo := make(map[string]blockatlas.Direction)
	// Get utxo platform to infer the direction and value
	platform := bitcoin.UtxoPlatform(o.Coin)
	for _, sub := range subs {
		tx, ok := txMap[sub.Address]
		if !ok {
			continue
		}
		// Verify the tx is for xpub
		xpubAddresses, ok := xpubs[sub.Address]
		for _, tx := range tx.Txs() {
			if sub.Address == tx.To {
				tx.Direction = blockatlas.DirectionIncoming
			} else if sub.Address == tx.From {
				tx.Direction = blockatlas.DirectionOutgoing
			}
			if ok {
				// Create a mapset for xpub addresses
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
