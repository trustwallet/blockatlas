package setup

import (
	"fmt"
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
		&models.Subscription{},
		&models.SubscriptionData{},
		&models.Tracker{},
	}

	uri string
)

func runPgContainerAndInitConnection() error {
	pool := runPgContainer()
	if err := pool.Retry(func() error {
		return db.Setup(uri)
	}); err != nil {
		return err
	}

	autoMigrate()
	return nil
}

func CleanupPgContainer() {
	db.GormDb.DropTable(tables...)
	autoMigrate()
}

func autoMigrate() {
	db.GormDb.AutoMigrate(tables...)
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
