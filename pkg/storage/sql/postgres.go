package cache

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/storage"
)

type PgSql struct {
	storage.Db
	client *gorm.DB
}

func (db *PgSql) Init(host string) error {
	client, err := gorm.Open("postgres", host)
	if err != nil {
		return errors.E(err, "postgress connection test failed")
	}
	db.client = client
	return nil
}

func (db *PgSql) Into(value interface{}) error {
	q, err := db.SQLQuery()
	if err != nil {
		return err
	}
	rows, err := db.client.Query(q)
	if err != nil {
		return err
	}
	return nil
}

func (db *PgSql) Update(value interface{}) error {
	db.client.Update(value)
	return nil
}

func (db *PgSql) Add(value interface{}) error {
	db.client.Create(value)
	return nil
}

func (db *PgSql) Delete() error {
	db.client.Delete(db.EntityName)
	return nil
}
