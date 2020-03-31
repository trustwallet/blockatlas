package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/trustwallet/blockatlas/db/models"

	"github.com/trustwallet/blockatlas/pkg/logger"
	"time"
)

var GormDb *gorm.DB

func Setup(uri string) error {
	var err error
	GormDb, err = gorm.Open("postgres", uri)
	if err != nil {
		return err
	}

	GormDb.AutoMigrate(
		&models.Subscription{},
		&models.SubscriptionData{},
		&models.Tracker{},
	)

	return nil
}

func RestoreConnectionWorker(timeout time.Duration, uri string) {
	logger.Info("Run PG RestoreConnectionWorker")
	for {
		if err := GormDb.DB().Ping(); err != nil {
			for {
				logger.Warn("PG is not available now")
				logger.Warn("Trying to connect to PG...")
				if err := Setup(uri); err != nil {
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
