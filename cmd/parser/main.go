package main

import (
	"context"
	"fmt"
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
	"github.com/trustwallet/golibs/network/mq"
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

	internal.InitRabbitMQ(config.Default.Observer.Rabbitmq.URL)

	platform.Init(config.Default.Platform)

	if err := internal.RawTransactionsExchange.Declare("topic"); err != nil {
		log.Fatal(err)
	}

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
	defer mq.Close()
	var (
		wg          sync.WaitGroup
		coinCancel  = make(map[string]context.CancelFunc)
		stopChannel = make(chan<- struct{}, len(platform.BlockAPIs))
	)
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

		coinCancel[coin.Handle] = cancel

		params := parser.Params{
			Ctx:                   ctx,
			Api:                   api,
			TransactionsExchange:  internal.RawTransactionsExchange,
			ParsingBlocksInterval: pollInterval,
			FetchBlocksTimeout:    fetchBlocksInterval,
			BacklogCount:          backlogCount,
			MaxBacklogBlocks:      maxBackLogBlocks,
			StopChannel:           stopChannel,
			Database:              database,
		}

		go parser.RunParser(params)

		log.WithFields(log.Fields{
			"coin":                     api.Coin().Handle,
			"interval":                 pollInterval,
			"backlog":                  backlogCount,
			"Max backlog":              maxBackLogBlocks,
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
