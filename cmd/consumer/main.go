package main

import (
	"context"
	"github.com/trustwallet/blockatlas/services/subscriber"
	"github.com/trustwallet/blockatlas/services/tokenindexer"
	"github.com/trustwallet/blockatlas/services/tokensearcher"
	"time"

	"github.com/trustwallet/blockatlas/services/notifier"

	log "github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/config"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/mq"
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
	var err error
	mqClient, err = new_mq.New(
		config.Default.Observer.Rabbitmq.URL,
		config.Default.Observer.Rabbitmq.Consumer.PrefetchCount,
		ctx,
	)
	if err != nil {
		log.Fatal("MQ init: ", err)
	}

	database, err = db.New(
		config.Default.Postgres.URL,
		config.Default.Postgres.Log,
	)
	if err != nil {
		log.Fatal("Postgres init: ", err)
	}
	go database.RestoreConnectionWorker(time.Second*10, config.Default.Postgres.URL)

	time.Sleep(time.Millisecond)
}

func main() {
	defer mqClient.Close()

	mqClient.AddStream(&notifier.NotifierConsumer{Database: database})
	mqClient.AddStream(&tokenindexer.TokenIndexerConsumer{Database: database})
	mqClient.AddStream(&tokensearcher.TokenSearcherConsumer{Database: database})
	mqClient.AddStream(&subscriber.TransactionSubscriberConsumer{Database: database})
	mqClient.AddStream(&subscriber.TokenSubscriberConsumer{Database: database})

	internal.SetupGracefulShutdownForObserver()
	cancel()
}
