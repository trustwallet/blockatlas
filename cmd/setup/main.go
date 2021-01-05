package main

import (
	"github.com/trustwallet/blockatlas/config"
	"github.com/trustwallet/golibs/network/middleware"
	"github.com/trustwallet/golibs/network/mq"

	log "github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/platform"
)

const (
	defaultConfigPath = "../../config.yml"
)

var (
	database *db.Instance
)

func init() {
	_, confPath := internal.ParseArgs("", defaultConfigPath)

	internal.InitConfig(confPath)
	internal.InitMQ(config.Default.Observer.Rabbitmq.URL)
	platform.Init(config.Default.Platform)

	var err error
	database, err = db.New(config.Default.Postgres.URL, config.Default.Postgres.Log)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Info("Start setup")

	if err := middleware.SetupSentry(config.Default.Sentry.DSN); err != nil {
		log.Error(err)
	}

	if err := db.Setup(database.Gorm); err != nil {
		log.Fatal(err)
	}

	if err := internal.RawTransactionsExchange.Declare("topic"); err != nil {
		log.Fatal(err)
	}

	queues := []mq.Queue{
		internal.TxNotifications,
		internal.RawTransactions,
		internal.Subscriptions,
		internal.SubscriptionsTokens,
		internal.RawTokens,
		internal.Subscriptions,
	}
	for _, queue := range queues {
		if err := queue.Declare(); err != nil {
			log.Fatal("Queue declare: ", queue, err)
		}
	}

	if err := internal.RawTransactionsExchange.Bind([]mq.Queue{internal.RawTokens, internal.RawTransactions}); err != nil {
		log.Fatal("Transactions Exchange bind: ", err)
	}

	log.Info("Finish setup")
}
