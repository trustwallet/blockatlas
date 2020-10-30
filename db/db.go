package db

import (
	"context"
	"errors"
	"time"

	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/pkg/logger"

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
	cfg := &gorm.Config{}

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
