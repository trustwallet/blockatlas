package sql

import (
	"github.com/jinzhu/gorm"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/storage/util"
)

type sql struct {
	Client *gorm.DB
}

type Handler func(value interface{}) error

func (db *sql) Get(value interface{}) error {
	err := db.Client.Last(value).Error
	if err != nil {
		return errors.E(err, util.ErrNotFound)
	}
	return nil
}

func (db *sql) Find(out interface{}, where ...interface{}) error {
	err := db.Client.Find(out, where...).Error
	if err != nil {
		return errors.E(err, util.ErrNotFound)
	}
	return nil
}

func (db *sql) CreateOrUpdate(value interface{}) error {
	if db.Client.NewRecord(value) {
		return db.Add(value)
	}
	return db.Update(value)
}

func (db *sql) Update(value interface{}) error {
	if db.Client.NewRecord(value) {
		return util.ErrNotFound
	}
	err := db.Client.Save(value).Error
	if err != nil {
		return errors.E(err, util.ErrNotUpdated)
	}
	return nil
}

func (db *sql) UpdateMany(values ...interface{}) error {
	return db.ToMany(db.Update, values...)
}

func (db *sql) Add(value interface{}) error {
	err := db.Client.Create(value).Error
	if err != nil {
		return errors.E(err, util.ErrNotStored)
	}
	return nil
}

func (db *sql) AddMany(values ...interface{}) error {
	return db.ToMany(db.Add, values...)
}

func (db *sql) Delete(value interface{}) error {
	err := db.Client.Delete(value).Error
	if err != nil {
		return errors.E(err, util.ErrNotDeleted)
	}
	return nil
}

func (db *sql) DeleteMany(values ...interface{}) error {
	return db.ToMany(db.Delete, values...)
}

func (db *sql) ToMany(handler Handler, values ...interface{}) error {
	tx := db.Client.Begin()
	for _, value := range values {
		err := handler(value)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}
