package sql

import (
	"github.com/jinzhu/gorm"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/pkg/storage/util"
)

type sql struct {
	Client *gorm.DB
}

type Handler func(value interface{}) error

func (db *sql) Get(value interface{}) error {
	err := db.Client.Last(value).Error
	if err != nil {
		return errors.E(err, util.ErrNotFound).PushToSentry()
	}
	return nil
}

func (db *sql) Find(out interface{}, where ...interface{}) error {
	err := db.Client.Find(out, where...).Error
	if err != nil {
		return errors.E(err, util.ErrNotFound).PushToSentry()
	}
	return nil
}

func (db *sql) Save(value interface{}) error {
	err := db.Client.Save(value).Error
	if err != nil {
		return errors.E(err, util.ErrNotUpdated).PushToSentry()
	}
	return nil
}

func (db *sql) Add(value interface{}) error {
	err := db.Client.Create(value).Error
	if err != nil {
		return errors.E(err, util.ErrNotStored).PushToSentry()
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
	err := db.Client.Delete(value).Error
	if err != nil {
		return errors.E(err, util.ErrNotDeleted).PushToSentry()
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
			logger.Error(err)
		}
	}
	return tx.Commit().Error
}
