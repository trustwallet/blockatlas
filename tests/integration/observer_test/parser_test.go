// +build integration

package observer_test

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/services/observer/notifier"
	"github.com/trustwallet/blockatlas/services/observer/parser"
	"github.com/trustwallet/blockatlas/tests/integration/setup"
	"testing"
	"time"
)

func TestParserFetchAndPublishBlock_NormalCase(t *testing.T) {
	setup.CleanupPgContainer(database.Gorm)
	stopChan := make(chan struct{}, 1)

	params := setupParser(stopChan)
	params.Database = database

	ctx, cancel := context.WithCancel(context.Background())

	params.Ctx = ctx
	params.Queue = mq.RawTransactions

	go parser.RunParser(params)

	time.Sleep(time.Microsecond)
	ConsumerToTestAmountOfBlocks(rawTransactionsChannel.GetMessage(), t, cancel)
	<-stopChan
}

func getMockedBlockAPI() blockatlas.BlockAPI {
	p := Platform{CoinIndex: 60}
	return &p
}

type Platform struct {
	CoinIndex uint
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return int64(100), nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[p.CoinIndex]
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	if num < 101 {
		return &blockatlas.Block{
			Number: num,
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
		}, nil
	}
	return &blockatlas.Block{}, nil
}

func ConsumerToTestAmountOfBlocks(delivery amqp.Delivery, t *testing.T, cancelFunc context.CancelFunc) {
	var txs blockatlas.Txs
	if err := json.Unmarshal(delivery.Body, &txs); err != nil {
		logger.Error(err)
		return
	}
	err := delivery.Ack(false)
	if err != nil {
		logger.Error(err)
	}

	assert.Equal(t, len(txs), 50)
	cancelFunc()
}

func setupParser(stopChan chan struct{}) parser.Params {
	minTime := time.Second
	maxTime := time.Second * 2
	maxBatchBlocksAmount := 100

	pollInterval := notifier.GetInterval(0, minTime, maxTime)

	backlogCount := 50

	return parser.Params{
		Api:                   getMockedBlockAPI(),
		ParsingBlocksInterval: pollInterval,
		BacklogCount:          backlogCount,
		MaxBacklogBlocks:      int64(maxBatchBlocksAmount),
		TxBatchLimit:          100,
		StopChannel:           stopChan,
	}
}
