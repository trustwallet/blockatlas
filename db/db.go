package db

import (
	"context"
	"errors"
	"gorm.io/gorm/logger"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/db/models"

	gocache "github.com/patrickmn/go-cache"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

type Instance struct {
	Gorm        *gorm.DB
	MemoryCache *gocache.Cache
}

// By gorm-bulk-insert author:
// "Depending on the number of variables included, 2000 to 3000 is recommended."
const batchCount = 3000

func New(uri, readUri string, logMode bool) (*Instance, error) {
	cfg := &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)}

	db, err := gorm.Open(postgres.Open(uri), cfg)
	if err != nil {
		return nil, err
	}

	err = db.Use(dbresolver.Register(dbresolver.Config{
		Replicas: []gorm.Dialector{
			postgres.Open(readUri),
		},
	}))
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

func (i *Instance) RestoreConnectionWorker(ctx context.Context, timeout time.Duration, uri string) {
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

func (i *Instance) MemorySet(key string, data []byte, exp time.Duration, ctx context.Context) error {
	i.MemoryCache.Set(key, data, exp)
	return nil
}

func (i *Instance) MemoryGet(key string, ctx context.Context) ([]byte, error) {
	res, ok := i.MemoryCache.Get(key)
	if !ok {
		return nil, errors.New("not found")
	}
	return res.([]byte), nil
}
