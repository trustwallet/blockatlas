package sql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/trustwallet/blockatlas/pkg/errors"
)

type PgSql struct {
	sql
}

func (db *PgSql) Init(host string) error {
	client, err := gorm.Open("postgres", host)
	if err != nil {
		return errors.E(err, "postgress connection failed").PushToSentry()
	}
	db.Client = client
	return nil
}
