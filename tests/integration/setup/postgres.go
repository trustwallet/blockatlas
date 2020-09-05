package setup

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/ory/dockertest"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/db/models"
	"log"
)

const (
	pgUser = "user"
	pgPass = "pass"
	pgDB   = "my_db"
)

var (
	pgResource     *dockertest.Resource
	pgContainerENV = []string{
		"POSTGRES_USER=" + pgUser,
		"POSTGRES_PASSWORD=" + pgPass,
		"POSTGRES_DB=" + pgDB,
	}

	tables = []interface{}{
		&models.AssetSubscription{},
		&models.NotificationSubscription{},
		&models.Tracker{},
		&models.AddressToAssetAssociation{},
		&models.Asset{},
		&models.Address{},
	}

	uri string
)

func runPgContainerAndInitConnection() (*db.Instance, error) {
	pool := runPgContainer()
	var (
		dbConn *db.Instance
		err    error
	)
	if err := pool.Retry(func() error {
		dbConn, err = db.New(uri, uri, "test", false)
		return err
	}); err != nil {
		return nil, err
	}
	dbConn.Gorm.Table("address_to_asset_associations").
		AddForeignKey("address_id", "addresses(id)", "RESTRICT", "RESTRICT").
		AddForeignKey("asset_id", "assets(id)", "RESTRICT", "RESTRICT")

	dbConn.Gorm.Table("notification_subscriptions").
		AddForeignKey("address_id", "addresses(id)", "RESTRICT", "RESTRICT")

	dbConn.Gorm.Table("asset_subscriptions").
		AddForeignKey("address_id", "addresses(id)", "RESTRICT", "RESTRICT")

	autoMigrate(dbConn.Gorm)
	return dbConn, nil
}

func CleanupPgContainer(dbConn *gorm.DB) {
	dbConn.DropTable(tables...)
	autoMigrate(dbConn)
}

func autoMigrate(dbConn *gorm.DB) {
	dbConn.AutoMigrate(tables...)
}

func stopPgContainer() error {
	return pgResource.Close()
}

func runPgContainer() *dockertest.Pool {
	var err error
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	pgResource, err = pool.Run("postgres", "latest", pgContainerENV)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	uri = fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable",
		pgUser, pgPass, pgResource.GetPort("5432/tcp"), pgDB,
	)
	return pool
}
