package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/trustwallet/blockatlas/db/models"

	"github.com/trustwallet/blockatlas/pkg/logger"
	"time"
)

type Instance struct {
	DB gorm.DB
}

func Setup(uri string) (*gorm.DB, error) {
	dbConn, err := gorm.Open("postgres", uri)
	if err != nil {
		return dbConn, err
	}

	dbConn.AutoMigrate(
		&models.Subscription{},
		&models.SubscriptionData{},
		&models.Tracker{},
	)

	return dbConn, nil
}

func RestoreConnectionWorker(dbConn *gorm.DB, timeout time.Duration, uri string) {
	logger.Info("Run PG RestoreConnectionWorker")
	for {
		if err := dbConn.DB().Ping(); err != nil {
			for {
				logger.Warn("PG is not available now")
				logger.Warn("Trying to connect to PG...")
				dbConn, err = gorm.Open("postgres", uri)
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
