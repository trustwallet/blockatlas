package main

import (
	"context"
	"fmt"
	"github.com/trustwallet/blockatlas/config"
	"github.com/trustwallet/blockatlas/services/spamfilter"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/platform"
	"github.com/trustwallet/blockatlas/services/parser"
)

const (
	defaultConfigPath = "../../config.yml"
)

var (
	ctx      context.Context
	cancel   context.CancelFunc
	database *db.Instance
)

func init() {
	ctx, cancel = context.WithCancel(context.Background())
	_, confPath := internal.ParseArgs("", defaultConfigPath)

	internal.InitConfig(confPath)

	internal.InitRabbitMQ(
		config.Default.Observer.Rabbitmq.URL,
		config.Default.Observer.Rabbitmq.Consumer.PrefetchCount,
	)

	platform.Init(config.Default.Platform)
	spamfilter.SpamList = config.Default.SpamWords

	if err := mq.RawTransactions.Declare(); err != nil {
		log.Fatal(err)
	}

	if err := mq.RawTransactionsTokenIndexer.Declare(); err != nil {
		log.Fatal(err)
	}

	if len(platform.BlockAPIs) == 0 {
		log.Fatal("No APIs to observe")
	}

	var err error
	database, err = db.New(config.Default.Postgres.URL, config.Default.Postgres.Read.URL,
		config.Default.Postgres.Log)
	if err != nil {
		log.Fatal(err)
	}
	go database.RestoreConnectionWorker(ctx, time.Second*10, config.Default.Postgres.URL)

	time.Sleep(time.Millisecond)
}

func main() {
	defer mq.Close()
	var (
		wg          sync.WaitGroup
		coinCancel  = make(map[string]context.CancelFunc)
		stopChannel = make(chan<- struct{}, len(platform.BlockAPIs))
	)
	txsBatchLimit := config.Default.Observer.TxsBatchLimit
	backlogTime := config.Default.Observer.Backlog
	minInterval := config.Default.Observer.BlockPoll.Min
	maxInterval := config.Default.Observer.BlockPoll.Max
	fetchBlocksInterval := config.Default.Observer.FetchBlocksInterval
	maxBackLogBlocks := config.Default.Observer.BacklogMaxBlocks

	go mq.FatalWorker(time.Second * 10)

	wg.Add(len(platform.BlockAPIs))
	for _, api := range platform.BlockAPIs {
		time.Sleep(time.Millisecond * 5)
		coin := api.Coin()
		pollInterval := parser.GetInterval(coin.BlockTime, minInterval, maxInterval)

		var backlogCount int
		if coin.BlockTime == 0 {
			backlogCount = 50
			log.WithFields(log.Fields{"coin": coin.Handle}).Warn("Unknown block time")
		} else {
			backlogCount = int(backlogTime / pollInterval)
		}

		// do not allow
		if txsBatchLimit < parser.MinTxsBatchLimit {
			txsBatchLimit = parser.MinTxsBatchLimit
		}

		coinCancel[coin.Handle] = cancel

		params := parser.Params{
			Ctx: ctx,
			Api: api,
			Queue: []mq.Queue{
				mq.RawTransactions,
				mq.RawTransactionsSearcher,
				mq.RawTransactionsTokenIndexer,
			},
			ParsingBlocksInterval: pollInterval,
			FetchBlocksTimeout:    fetchBlocksInterval,
			BacklogCount:          backlogCount,
			MaxBacklogBlocks:      maxBackLogBlocks,
			StopChannel:           stopChannel,
			TxBatchLimit:          txsBatchLimit,
			Database:              database,
		}

		go parser.RunParser(params)

		log.WithFields(log.Fields{
			"interval":                 pollInterval,
			"backlog":                  backlogCount,
			"Max backlog":              maxBackLogBlocks,
			"Txs Batch limit":          txsBatchLimit,
			"Fetching blocks interval": fetchBlocksInterval,
		}).Info("Parser params")

		wg.Done()
	}

	wg.Wait()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	cancel()
	log.Info("Shutdown parser ...")
	for coin, cancel := range coinCancel {
		log.Info(fmt.Sprintf("Starting to stop %s parser...", coin))
		cancel()
	}
	for {
		if len(stopChannel) == len(platform.BlockAPIs) {
			log.Info("All parsers are stopped")
			break
		}
	}

	log.Info("Exiting gracefully")
}
