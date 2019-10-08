package sql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/trustwallet/blockatlas/pkg/errors"
)

type MySql struct {
	PgSql
}

func (db *MySql) Init(host string) error {
	client, err := gorm.Open("mysql", host)
	if err != nil {
		return errors.E(err, "mysql connection failed")
	}
	db.client = client
	return nil
}
