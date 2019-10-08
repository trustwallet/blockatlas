package sql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/storage"
)

type PgSql struct {
	client *gorm.DB
}

func (db *PgSql) Init(host string) error {
	client, err := gorm.Open("postgres", host)
	if err != nil {
		return errors.E(err, "postgress connection failed")
	}
	db.client = client
	return nil
}

func (db *PgSql) Get(value interface{}) error {
	err := db.client.Last(value).Error
	if err != nil {
		return errors.E(err, storage.ErrNotFound)
	}
	return nil
}

func (db *PgSql) Update(value interface{}) error {
	if db.client.NewRecord(value) {
		return storage.ErrNotFound
	}
	err := db.client.Save(value).Error
	if err != nil {
		return errors.E(err, storage.ErrNotUpdated)
	}
	return nil
}

func (db *PgSql) Add(value interface{}) error {
	if !db.client.NewRecord(value) {
		return storage.ErrAlreadyExist
	}
	db.client.AutoMigrate(value)
	err := db.client.Create(value).Error
	if err != nil {
		return errors.E(err, storage.ErrNotStored)
	}
	return nil
}

func (db *PgSql) Delete(value interface{}) error {
	if db.client.NewRecord(value) {
		return storage.ErrNotFound
	}
	err := db.client.Delete(value).Error
	if err != nil {
		return errors.E(err, storage.ErrNotDeleted)
	}
	return nil
}
