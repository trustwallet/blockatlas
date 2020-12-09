package db

import (
	"errors"
	"time"

	"gorm.io/gorm/logger"

	log "github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/db/models"

	gocache "github.com/patrickmn/go-cache"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Instance struct {
	Gorm        *gorm.DB
	MemoryCache *gocache.Cache
}

// By gorm-bulk-insert author:
// "Depending on the number of variables included, 2000 to 3000 is recommended."
const batchCount = 3000

func New(url string, log bool) (*Instance, error) {
	var logMode logger.LogLevel
	if log {
		logMode = logger.Info
	}

	cfg := &gorm.Config{Logger: logger.Default.LogMode(logMode)}

	db, err := gorm.Open(postgres.Open(url), cfg)
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.NotificationSubscription{},
		&models.Tracker{},
		&models.Asset{},
		&models.AssetSubscription{},
		&models.Address{},
		&models.AddressToAssetAssociation{},
	)
	if err != nil {
		return nil, err
	}
	mc := gocache.New(gocache.NoExpiration, gocache.NoExpiration)
	i := &Instance{Gorm: db, MemoryCache: mc}

	return i, nil
}

func (i *Instance) RestoreConnectionWorker(timeout time.Duration, uri string) {
	log.Info("Run PG RestoreConnectionWorker")

	for {
		if err := i.restoreConnection(uri); err != nil {
			log.Error("PG is not available now")
		}
		time.Sleep(timeout)
	}
}

func (i *Instance) restoreConnection(uri string) error {
	db, err := i.Gorm.DB()
	if err != nil {
		return err
	}

	log.Info("Run restoreConnection")

	if err = db.Ping(); err != nil {
		log.Warn("PG is not available now")
		log.Warn("Trying to connect to PG...")
		i.Gorm, err = gorm.Open(postgres.Open(uri), &gorm.Config{})
		if err != nil {
			return err
		}
		log.Info("PG connection restored")
	}
	return nil
}

func (i *Instance) MemorySet(key string, data []byte, exp time.Duration) error {
	i.MemoryCache.Set(key, data, exp)
	return nil
}

func (i *Instance) MemoryGet(key string) ([]byte, error) {
	res, ok := i.MemoryCache.Get(key)
	if !ok {
		return nil, errors.New("not found")
	}
	return res.([]byte), nil
}
