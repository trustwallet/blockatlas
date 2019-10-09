package storage

import (
	"errors"
	"github.com/spf13/viper"
	er "github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/pkg/storage/sql"
)

var (
	ErrNotFound     = errors.New("storage: obj not found")
	ErrNotStored    = errors.New("storage: not stored")
	ErrNotUpdated   = errors.New("storage: not updated")
	ErrNotDeleted   = errors.New("storage: not deleted")
	ErrAlreadyExist = errors.New("storage: object already exist")
)

var Postgres sql.PgSql

func InitDatabases() {
	err := Postgres.Init(viper.GetString("observer.postgres"))
	if err != nil {
		logger.Fatal(er.E(err), "Cannot connect to Postgres")
	}
}
