package notifier

import (
	"encoding/json"
	mapset "github.com/deckarep/golang-set"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/observer/parser"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/pkg/numbers"
	"github.com/trustwallet/blockatlas/platform/bitcoin"
	"github.com/trustwallet/blockatlas/storage"
	"time"
)

type Event struct {
	Subscription blockatlas.Subscription
	Tx           *blockatlas.Tx
}

func ProcessBlock(delivery amqp.Delivery, s storage.Addresses) {
	var blockData parser.BlockData
	if err := json.Unmarshal(delivery.Body, &blockData); err != nil {
		logger.Error(err)
		return
	}

	txMap := GetTxs(blockData.Block)
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
	subs, err := s.Lookup(blockData.Coin, addresses)
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
			inferUtxoValue(&tx, sub.Address, blockData.Coin)
			dispatch(Event{
				Subscription: sub,
				Tx:           &tx,
			})
		}
	}
}

func GetTxs(block blockatlas.Block) map[string]*blockatlas.TxSet {
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
	switch meta := tx.Meta.(type) {
	case blockatlas.TokenTransfer:
		return determineDirection(address, meta.From, meta.To)
	case blockatlas.NativeTokenTransfer:
		return determineDirection(address, meta.From, meta.To)
	default:
		return determineDirection(address, tx.From, tx.To)
	}
}

func determineDirection(address, from, to string) blockatlas.Direction {
	if address == to {
		if from == to {
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

func GetInterval(value int, minInterval, maxInterval time.Duration) time.Duration {
	interval := time.Duration(value) * time.Millisecond
	pMin := numbers.Max(minInterval.Nanoseconds(), interval.Nanoseconds())
	pMax := numbers.Min(int(maxInterval.Nanoseconds()), int(pMin))
	return time.Duration(pMax)
}
