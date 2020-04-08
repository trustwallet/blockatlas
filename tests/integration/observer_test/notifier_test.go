// +build integration

package observer_test

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/services/observer/notifier"
	"github.com/trustwallet/blockatlas/tests/integration/setup"
	"testing"
	"time"
)

var (
	txs = blockatlas.Txs{
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
	}
)

func TestNotifier(t *testing.T) {
	setup.CleanupPgContainer(database.Gorm)

	err := database.AddSubscriptions(1, []models.SubscriptionData{{Coin: 714, Address: "tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a", SubscriptionId: 1}})
	assert.Nil(t, err)

	err = produceTxs(txs)
	assert.Nil(t, err)

	ctx, cancel := context.WithCancel(context.Background())

	go mq.RunConsumerForChannelWithCancelAndDbConn(notifier.RunNotifier, rawTransactionsChannel, database, ctx)
	time.Sleep(time.Second * 3)
	msg := transactionsChannel.GetMessage()
	ConsumerToTestTransactions(msg, t)
	cancel()
}

func ConsumerToTestTransactions(delivery amqp.Delivery, t *testing.T) {
	var notifications []notifier.TransactionNotification
	if err := json.Unmarshal(delivery.Body, &notifications); err != nil {
		assert.Nil(t, err)
		return
	}
	err := delivery.Ack(false)
	if err != nil {
		assert.Nil(t, err)
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

	assert.Equal(t, notifier.TransactionNotification{
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
		Id: 1,
	}, notifications[0])

	return
}

func produceTxs(txs blockatlas.Txs) error {
	body, err := json.Marshal(txs)
	if err != nil {
		return err
	}
	return mq.RawTransactions.Publish(body)
}
