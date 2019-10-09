package storage

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/pkg/storage/sql"
)

var Postgres sql.PgSql

func InitDatabases() {
	err := Postgres.Init(viper.GetString("observer.postgres"))
	if err != nil {
		logger.Fatal(errors.E(err), "Cannot connect to Postgres")
	}
}
