// +build integration

package docker_test

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/services/observer/notifier"
	"github.com/trustwallet/blockatlas/storage"
	"github.com/trustwallet/blockatlas/tests/docker_test/setup"
	"testing"
)

var (
	block = blockatlas.Block{
		Number: 100,
		ID:     "",
		Txs: []blockatlas.Tx{
			{
				ID:     "95CF63FAA27579A9B6AF84EF8B2DFEAC29627479E9C98E7F5AE4535E213FA4C9",
				Coin:   coin.BNB,
				From:   "tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a",
				To:     "tbnb12hlquylu78cjylk5zshxpdj6hf3t0tahwjt3ex",
				Fee:    "125000",
				Date:   1555117625,
				Block:  7928667,
				Status: blockatlas.StatusCompleted,
				Memo:   "test",
				Meta: blockatlas.NativeTokenTransfer{
					TokenID:  "YLC-D8B",
					Symbol:   "YLC",
					Value:    "210572645",
					Decimals: 8,
					From:     "tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a",
					To:       "tbnb12hlquylu78cjylk5zshxpdj6hf3t0tahwjt3ex",
				},
			},
		},
	}
	notifierCounter TestsCounter

	stopChan2      = make(chan struct{})
	globalTesting2 *testing.T
)

func TestNotifier(t *testing.T) {

	globalTesting2 = t

	if err := mq.ConfirmedBlocks.Declare(); err != nil {
		assert.Nil(t, err)
	}

	if err := mq.Transactions.Declare(); err != nil {
		assert.Nil(t, err)
	}

	err := setup.Cache.AddSubscriptions([]blockatlas.Subscription{{Coin: 714, Address: "tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a", GUID: "guid_test"}})
	assert.Nil(t, err)

	for i := 0; i < 101; i++ {
		err := produceBlock(block)
		assert.Nil(t, err)
	}

	go mq.ConfirmedBlocks.RunConsumer(notifier.RunNotifier, setup.Cache)
	go mq.Transactions.RunConsumer(ConsumerToTestTransactions, setup.Cache)

	<-stopChan2
}

func ConsumerToTestTransactions(delivery amqp.Delivery, s storage.Addresses) {
	var event notifier.DispatchEvent
	if err := json.Unmarshal(delivery.Body, &event); err != nil {
		assert.Nil(globalTesting2, err)
		return
	}
	err := delivery.Ack(false)
	if err != nil {
		assert.Nil(globalTesting2, err)
	}

	memo := blockatlas.NativeTokenTransfer{
		Name:     "",
		TokenID:  "YLC-D8B",
		Symbol:   "YLC",
		Value:    "210572645",
		Decimals: 8,
		From:     "tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a",
		To:       "tbnb12hlquylu78cjylk5zshxpdj6hf3t0tahwjt3ex",
	}

	assert.Equal(globalTesting2, notifier.DispatchEvent{
		Action: blockatlas.TxNativeTokenTransfer,
		Result: &blockatlas.Tx{
			Type:      blockatlas.TxNativeTokenTransfer,
			Direction: "outgoing",
			ID:        "95CF63FAA27579A9B6AF84EF8B2DFEAC29627479E9C98E7F5AE4535E213FA4C9",
			Coin:      coin.BNB,
			From:      "tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a",
			To:        "tbnb12hlquylu78cjylk5zshxpdj6hf3t0tahwjt3ex",
			Fee:       "125000",
			Date:      1555117625,
			Block:     7928667,
			Status:    blockatlas.StatusCompleted,
			Memo:      "test",
			Meta:      &memo,
		},
		GUID: "guid_test",
	}, event)

	notifierCounter.M.Lock()
	notifierCounter.Counter++
	notifierCounter.M.Unlock()

	notifierCounter.M.Lock()
	val := notifierCounter.Counter
	notifierCounter.M.Unlock()

	if val == 100 {
		stopChan2 <- struct{}{}
	}

}
func produceBlock(block blockatlas.Block) error {
	body, err := json.Marshal(block)
	if err != nil {
		return err
	}
	return mq.ConfirmedBlocks.Publish(body)
}
