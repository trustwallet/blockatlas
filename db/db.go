package db

import (
	"context"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Instance struct {
	Gorm *gorm.DB
}

func New(uri string, logMode bool) (*Instance, error) {
	cfg := &gorm.Config{}
	if logMode {
		cfg.Logger = gormlogger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			gormlogger.Config{
				SlowThreshold: time.Second,
				LogLevel:      gormlogger.Info,
				Colorful:      false,
			},
		)
	}
	g, err := gorm.Open(postgres.Open(uri), cfg)
	if err != nil {
		return nil, err
	}

	err = g.AutoMigrate(
		&models.NotificationSubscription{},
		&models.Tracker{},
		&models.AddressToAssetAssociation{},
		&models.Asset{},
		&models.AssetSubscription{},
		&models.Address{},
	)
	if err != nil {
		return nil, err
	}

	i := &Instance{Gorm: g}

	return i, nil
}

func (i *Instance) RestoreConnectionWorker(ctx context.Context, timeout time.Duration, uri string) {
	logger.Info("Run PG RestoreConnectionWorker")
	t := time.NewTicker(timeout)

	if err := i.restoreConnection(uri); err != nil {
		logger.Warn("PG is still unavailable:", err)
	}

	select {
	case <-ctx.Done():
		logger.Info("Ctx.Done RestoreConnectionWorker exit")
		return
	case <-t.C:
		if err := i.restoreConnection(uri); err != nil {
			logger.Warn("PG is still unavailable:", err)
		}
	}
}

func (i *Instance) restoreConnection(uri string) error {
	db, err := i.Gorm.DB()
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		logger.Warn("PG is not available now")
		logger.Warn("Trying to connect to PG...")
		i.Gorm, err = gorm.Open(postgres.Open(uri), &gorm.Config{})
		if err != nil {
			return err
		}
		logger.Info("PG connection restored")
	}
	return nil
}
