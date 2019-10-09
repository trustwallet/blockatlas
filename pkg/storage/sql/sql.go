package sql

import (
	"github.com/jinzhu/gorm"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/storage/util"
)

type sql struct {
	client *gorm.DB
}

func (db *sql) Get(value interface{}) error {
	err := db.client.Last(value).Error
	if err != nil {
		return errors.E(err, util.ErrNotFound)
	}
	return nil
}

func (db *sql) Update(value interface{}) error {
	if db.client.NewRecord(value) {
		return util.ErrNotFound
	}
	err := db.client.Save(value).Error
	if err != nil {
		return errors.E(err, util.ErrNotUpdated)
	}
	return nil
}

func (db *sql) Add(value interface{}) error {
	if !db.client.NewRecord(value) {
		return util.ErrAlreadyExist
	}
	db.client.AutoMigrate(value)
	err := db.client.Create(value).Error
	if err != nil {
		return errors.E(err, util.ErrNotStored)
	}
	return nil
}

func (db *sql) Delete(value interface{}) error {
	if db.client.NewRecord(value) {
		return util.ErrNotFound
	}
	err := db.client.Delete(value).Error
	if err != nil {
		return errors.E(err, util.ErrNotDeleted)
	}
	return nil
}
