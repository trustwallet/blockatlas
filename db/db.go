package db

import (
	"errors"
	"time"

	"github.com/trustwallet/blockatlas/db/models"

	"gorm.io/gorm/logger"

	gocache "github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Instance struct {
	Gorm        *gorm.DB
	MemoryCache *gocache.Cache
}

func New(url string, log bool) (*Instance, error) {
	var logMode logger.LogLevel
	if log {
		logMode = logger.Info
	}

	cfg := &gorm.Config{Logger: logger.Default.LogMode(logMode), SkipDefaultTransaction: true}

	db, err := gorm.Open(postgres.Open(url), cfg)
	if err != nil {
		return nil, err
	}

	mc := gocache.New(gocache.NoExpiration, gocache.NoExpiration)
	i := &Instance{Gorm: db, MemoryCache: mc}

	return i, nil
}

func Setup(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Tracker{},
		&models.Asset{},
		&models.Subscription{},
		&models.SubscriptionsAssetAssociation{},
	)
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
