package docker_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"

	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/services/observer/notifier"
	"github.com/trustwallet/blockatlas/services/observer/parser"
	"github.com/trustwallet/blockatlas/storage"
	"github.com/trustwallet/blockatlas/tests/docker_test/setup"
	"sync"
	"testing"
	"time"
)

type TestsCounter struct {
	M       sync.Mutex
	Counter int
}

var (
	globalTestsCounter TestsCounter
	stopChan           = make(chan struct{})
	globalTesting      *testing.T
)

func TestParserFetchAndPublishBlock_NormalCase(t *testing.T) {
	globalTesting = t
	p := setupParser()

	err := p.FetchAndPublishBlocks(0, 100)
	assert.Nil(t, err)

	go mq.ConfirmedBlocks.RunConsumer(ConsumerToTestAmountOfBlocks, nil)

	<-stopChan
}

func TestParserRun(t *testing.T) {
	globalTesting = t
	p := setupParser()

}

func getMockedBlockAPI() blockatlas.BlockAPI {
	p := Platform{CoinIndex: 60}
	return &p
}

type Platform struct {
	CoinIndex uint
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return 0, nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[p.CoinIndex]
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	return &blockatlas.Block{
		Number: 0,
		ID:     "0x",
		Txs:    nil,
	}, nil
}

func ConsumerToTestAmountOfBlocks(delivery amqp.Delivery, s storage.Addresses) {
	var block blockatlas.Block
	if err := json.Unmarshal(delivery.Body, &block); err != nil {
		logger.Error(err)
		return
	}
	err := delivery.Ack(false)
	if err != nil {
		logger.Error(err)
	}

	globalTestsCounter.M.Lock()
	globalTestsCounter.Counter++
	globalTestsCounter.M.Unlock()

	globalTestsCounter.M.Lock()
	val := globalTestsCounter.Counter
	globalTestsCounter.M.Unlock()

	assert.Equal(globalTesting, int(block.Number), 0)
	assert.Equal(globalTesting, block.ID, "0x")
	if val == 100 {
		stopChan <- struct{}{}
	}
}

func setupParser() *parser.Parser {
	globalTestsCounter = TestsCounter{
		M:       sync.Mutex{},
		Counter: 0,
	}

	api := getMockedBlockAPI()

	if err := mq.ConfirmedBlocks.Declare(); err != nil {
		logger.Fatal(err)
	}

	minTime := time.Second
	maxTime := time.Second * 2
	maxBatchBlocksAmount := 10

	pollInterval := notifier.GetInterval(0, minTime, maxTime)

	backlogCount := 50

	return &parser.Parser{
		BlockAPI:              api,
		Storage:               setup.Cache,
		ParsingBlocksInterval: pollInterval,
		BacklogCount:          backlogCount,
		MaxBacklogBlocks:      int64(maxBatchBlocksAmount),
	}
}
