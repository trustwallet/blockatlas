package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/trustwallet/blockatlas/pkg/logger"
	"time"
)

var (
	GormDb *gorm.DB
)

func Setup(uri string) error {
	var err error
	GormDb, err = gorm.Open("postgres", uri)
	if err != nil {
		return err
	}

	//GormDb.AutoMigrate(
	//	&models.User{},
	//	&models.UserRewards{},
	//	&models.UserDevices{},
	//	&models.UserReferrals{},
	//	&models.Device{},
	//	&models.Referral{},
	//	&models.Reward{},
	//	&models.CoinStatusHistory{},
	//	&models.Token{},
	//	&models.Notification{},
	//)

	//// Setup. Make this async in the future
	//SetupRewards()

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

func FindByQuery(query *gorm.DB, result interface{}) error {
	if err := query.Find(result).Error; err != nil {
		return err
	}
	return nil
}
