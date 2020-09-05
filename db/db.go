package db

import (
	"context"
	"errors"
	gormbulk "github.com/t-tiger/gorm-bulk-insert"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"go.elastic.co/apm/module/apmgorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"reflect"
	"time"
)

type Instance struct {
	Gorm *gorm.DB
}

const batchCount = 1000

func New(uri, env string) (*Instance, error) {
	var (
		g   *gorm.DB
		err error
	)
	if env == "prod" {
		g, err = apmgorm.Open("postgres", uri)
	} else {
		g, err = gorm.Open(postgres.Open(uri), &gorm.Config{})
	}

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

// Example:
// postgres.BulkInsert(DBWrite, []models.User{...})
func BulkInsert(db *gorm.DB, dbModels interface{}) error {
	interfaceSlice, err := getInterfaceSlice(dbModels)
	if err != nil {
		return err
	}
	batchList := getInterfaceSliceBatch(interfaceSlice, batchCount)
	for _, batch := range batchList {
		err := gormbulk.BulkInsert(db, batch, len(batch))
		if err != nil {
			return err
		}
	}
	return nil
}

func getInterfaceSliceBatch(values []interface{}, sizeUint uint) [][]interface{} {
	size := int(sizeUint)
	resultLength := (len(values) + size - 1) / size
	result := make([][]interface{}, resultLength)
	lo, hi := 0, size
	for i := range result {
		if hi > len(values) {
			hi = len(values)
		}
		result[i] = values[lo:hi:hi]
		lo, hi = hi, hi+size
	}
	return result
}

func getInterfaceSlice(slice interface{}) ([]interface{}, error) {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		return nil, errors.New("InterfaceSlice() given a non-slice type")
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret, nil
}
