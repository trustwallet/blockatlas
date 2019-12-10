package sql

import (
	"github.com/jinzhu/gorm"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/storage/util"
)

type sql struct {
	Client *gorm.DB
}

type Handler func(db *gorm.DB, value interface{}) error

func Get(db *gorm.DB, value interface{}) error {
	err := db.Where(value).Take(value).Error
	if err != nil {
		return errors.E(err, util.ErrNotFound, errors.Params{"method": "Get", "value": value})
	}
	return nil
}

func Find(db *gorm.DB, out interface{}, where ...interface{}) error {
	err := db.Find(out, where...).Error
	if err != nil {
		return errors.E(err, util.ErrNotFound, errors.Params{"method": "Find", "out": out, "where": where})
	}
	return nil
}

func Save(db *gorm.DB, value interface{}) error {
	err := db.Save(value).Error
	if err != nil {
		return errors.E(err, util.ErrNotUpdated, errors.Params{"method": "Save", "value": value})
	}
	return nil
}

func Add(db *gorm.DB, value interface{}) error {
	err := db.Where(value).FirstOrCreate(value).Error
	if err != nil {
		return errors.E(err, util.ErrNotStored, errors.Params{"method": "Add", "value": value})
	}
	return nil
}

func (db *sql) AddMany(values ...interface{}) error {
	return db.Batch(true, Add, values...)
}

func (db *sql) MustAddMany(values ...interface{}) error {
	return db.Batch(false, Add, values...)
}

func Delete(db *gorm.DB, value interface{}) error {
	err := Get(db, value)
	if err != nil {
		return err
	}
	err = db.Delete(value).Error
	if err != nil {
		return errors.E(err, util.ErrNotDeleted)
	}
	return nil
}

func (db *sql) DeleteMany(values ...interface{}) error {
	return db.Batch(true, Delete, values...)
}

func (db *sql) MustDeleteMany(values ...interface{}) error {
	return db.Batch(false, Delete, values...)
}

func (db *sql) Batch(rollback bool, handler Handler, values ...interface{}) error {
	tx := db.Client.Begin()
	for _, value := range values {
		err := handler(tx, value)
		if err != nil {
			if rollback {
				tx.Rollback()
				return err
			}
		}
	}
	return tx.Commit().Error
}
