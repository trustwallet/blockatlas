package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"time"
)

type Instance struct {
	Gorm *gorm.DB
}

func New(uri string) (*Instance, error) {
	g, err := gorm.Open("postgres", uri)
	if err != nil {
		return nil, err
	}

	g.AutoMigrate(
		&models.Subscription{},
		&models.SubscriptionData{},
		&models.Tracker{},
	)

	i := &Instance{Gorm: g}

	return i, nil
}

func RestoreConnectionWorker(database *Instance, timeout time.Duration, uri string) {
	logger.Info("Run PG RestoreConnectionWorker")
	for {
		if err := database.Gorm.DB().Ping(); err != nil {
			for {
				logger.Warn("PG is not available now")
				logger.Warn("Trying to connect to PG...")
				database.Gorm, err = gorm.Open("postgres", uri)
				if err != nil {
					logger.Warn("PG is still unavailable:", err.Error())
					time.Sleep(timeout)
					continue
				} else {
					logger.Info("PG connection restored")
					break
				}
			}
		}
		time.Sleep(timeout)
	}
}
