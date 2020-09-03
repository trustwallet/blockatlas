package db

import (
	"errors"
	"github.com/jinzhu/gorm"
	gormbulk "github.com/t-tiger/gorm-bulk-insert"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"go.elastic.co/apm/module/apmgorm"
	_ "go.elastic.co/apm/module/apmgorm/dialects/postgres"
	"reflect"
	"time"
)

type Instance struct {
	Gorm     *gorm.DB
	GormRead *gorm.DB
}

const batchCount = 1000

func New(uri, readURI, env string, mode bool) (*Instance, error) {
	var (
		g   *gorm.DB
		rg  *gorm.DB
		err error
	)
	if env == "prod" {
		g, err = apmgorm.Open("postgres", uri)
	} else {
		g, err = gorm.Open("postgres", uri)
	}

	if err != nil {
		return nil, err
	}

	if env == "prod" {
		rg, err = apmgorm.Open("postgres", readURI)
	} else {
		rg, err = gorm.Open("postgres", readURI)
	}

	if err != nil {
		return nil, err
	}

	g.AutoMigrate(
		&models.NotificationSubscription{},
		&models.Tracker{},
		&models.AddressToAssetAssociation{},
		&models.Asset{},
		&models.AssetSubscription{},
		&models.Address{},
	)
	g.Table("address_to_asset_associations").
		AddForeignKey("address_id", "addresses(id)", "RESTRICT", "RESTRICT").
		AddForeignKey("asset_id", "assets(id)", "RESTRICT", "RESTRICT")

	g.Table("notification_subscriptions").
		AddForeignKey("address_id", "addresses(id)", "RESTRICT", "RESTRICT")

	g.Table("asset_subscriptions").
		AddForeignKey("address_id", "addresses(id)", "RESTRICT", "RESTRICT")

	g.LogMode(mode)
	rg.LogMode(mode)

	i := &Instance{Gorm: g, GormRead: rg}

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
