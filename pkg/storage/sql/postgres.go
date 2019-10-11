package sql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"time"
)

type PgSql struct {
	sql
}

func (db *PgSql) Init(host string, conns int) error {
	client, err := gorm.Open("postgres", host)
	if err != nil {
		return errors.E(err, "postgress connection failed").PushToSentry()
	}
	client.DB().SetMaxIdleConns(conns * 2)
	client.DB().SetConnMaxLifetime(time.Second * 5)
	db.Client = client
	return nil
}
