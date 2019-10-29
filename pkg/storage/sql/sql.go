package sql

import (
	"github.com/jinzhu/gorm"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/storage/util"
)

type sql struct {
	Client *gorm.DB
	Tx     *gorm.DB
}

type Handler func(value interface{}) error

func (db *sql) GeTx() *gorm.DB {
	tx := db.Client
	if db.Tx != nil {
		tx = db.Tx
	}
	return tx
}

func (db *sql) Get(value interface{}) error {
	tx := db.GeTx()
	err := tx.Where(value).Take(value).Error
	if err != nil {
		return errors.E(err, util.ErrNotFound, errors.Params{"method": "Get", "value": value})
	}
	return nil
}

func (db *sql) Find(out interface{}, where ...interface{}) error {
	tx := db.GeTx()
	err := tx.Find(out, where...).Error
	if err != nil {
		return errors.E(err, util.ErrNotFound, errors.Params{"method": "Find", "out": out, "where": where})
	}
	return nil
}

func (db *sql) Save(value interface{}) error {
	tx := db.GeTx()
	err := tx.Save(value).Error
	if err != nil {
		return errors.E(err, util.ErrNotUpdated, errors.Params{"method": "Save", "value": value})
	}
	return nil
}

func (db *sql) Add(value interface{}) error {
	tx := db.GeTx()
	err := tx.Where(value).FirstOrCreate(value).Error
	if err != nil {
		return errors.E(err, util.ErrNotStored, errors.Params{"method": "Add", "value": value})
	}
	return nil
}

func (db *sql) AddMany(values ...interface{}) error {
	return db.Batch(true, db.Add, values...)
}

func (db *sql) MustAddMany(values ...interface{}) error {
	return db.Batch(false, db.Add, values...)
}

func (db *sql) Delete(value interface{}) error {
	err := db.Get(value)
	if err != nil {
		return err
	}
	tx := db.GeTx()
	err = tx.Delete(value).Error
	if err != nil {
		return errors.E(err, util.ErrNotDeleted)
	}
	return nil
}

func (db *sql) DeleteMany(values ...interface{}) error {
	return db.Batch(true, db.Delete, values...)
}

func (db *sql) MustDeleteMany(values ...interface{}) error {
	return db.Batch(false, db.Delete, values...)
}

func (db *sql) Batch(rollback bool, handler Handler, values ...interface{}) error {
	db.Tx = db.Client.Begin()
	defer func(db *sql) {
		db.Tx = nil
	}(db)
	for _, value := range values {
		err := handler(value)
		if err != nil {
			if rollback {
				db.Tx.Rollback()
				return err
			}
		}
	}
	return db.Tx.Commit().Error
}
