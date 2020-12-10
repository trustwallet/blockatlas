package main

import (
	"context"
	"fmt"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/services/notifier"
	"github.com/trustwallet/blockatlas/services/tokensearcher"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/trustwallet/blockatlas/config"

	log "github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/internal"
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
	mqClient *new_mq.Client
)

func init() {
	ctx, cancel = context.WithCancel(context.Background())
	_, confPath := internal.ParseArgs("", defaultConfigPath)

	internal.InitConfig(confPath)

	mqClient, _ = new_mq.New(
		config.Default.Observer.Rabbitmq.URL,
		config.Default.Observer.Rabbitmq.Consumer.PrefetchCount,
		ctx,
	)
	mqClient.AddPublish(&notifier.NotifierConsumer{
		Database: database,
		MQClient: mqClient,
	})
	mqClient.AddPublish(&tokensearcher.TokenSearcherConsumer{
		Database: database,
	})
	mqClient.AddPublish(&tokensearcher.TokenSearcherConsumer{
		Database: database,
	})
	platform.Init(config.Default.Platform)
	if len(platform.BlockAPIs) == 0 {
		log.Fatal("No APIs to observe")
	}

	var err error
	database, err = db.New(config.Default.Postgres.URL, config.Default.Postgres.Log)
	if err != nil {
		log.Fatal(err)
	}
	go database.RestoreConnectionWorker(time.Second*10, config.Default.Postgres.URL)

	time.Sleep(time.Millisecond)
}

func main() {
	defer mqClient.Close()
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
			Ctx:               ctx,
			Api:               api,
			TransactionsQueue: new_mq.RawTransactions,
			TokenTransactionsQueue: []new_mq.Queue{
				new_mq.RawTransactionsSearcher,
				new_mq.RawTransactionsTokenIndexer,
			},
			ParsingBlocksInterval: pollInterval,
			FetchBlocksTimeout:    fetchBlocksInterval,
			BacklogCount:          backlogCount,
			MaxBacklogBlocks:      maxBackLogBlocks,
			StopChannel:           stopChannel,
			TxBatchLimit:          txsBatchLimit,
			Database:              database,
			MQClient:              mqClient,
		}

		go parser.RunParser(params)

		log.WithFields(log.Fields{
			"coin":                     api.Coin().Handle,
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
