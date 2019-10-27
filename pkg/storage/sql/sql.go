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
	err := db.Client.Where(value).Take(value).Error
	if err != nil {
		return errors.E(err, util.ErrNotFound, errors.Params{"method": "Get", "value": value})
	}
	return nil
}

func (db *sql) Find(out interface{}, where ...interface{}) error {
	err := db.Client.Find(out, where...).Error
	if err != nil {
		return errors.E(err, util.ErrNotFound, errors.Params{"method": "Find", "out": out, "where": where})
	}
	return nil
}

func (db *sql) Save(value interface{}) error {
	err := db.Client.Save(value).Error
	if err != nil {
		return errors.E(err, util.ErrNotUpdated, errors.Params{"method": "Save", "value": value})
	}
	return nil
}

func (db *sql) Add(value interface{}) error {
	err := db.Client.Where(value).FirstOrCreate(value).Error
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
	err = db.Client.Delete(value).Error
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
	tx := db.Client.Begin()
	for _, value := range values {
		err := handler(value)
		if err != nil {
			if rollback {
				tx.Rollback()
				return err
			}
		}
	}
	return tx.Commit().Error
}
