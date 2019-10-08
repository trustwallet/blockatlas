package storage

import (
	"errors"
)

var (
	ErrNotFound      = errors.New("storage: obj not found")
	ErrNotStored     = errors.New("storage: not stored")
	ErrNotDeleted    = errors.New("storage: not deleted")
	ErrNotSupport    = errors.New("storage: not support")
	ErrEmptyQuery    = errors.New("storage: missing query value")
	ErrEmptyEntity   = errors.New("storage: missing entity value")
	ErrEmptyDatabase = errors.New("storage: missing database value")
)

type Db struct {
	EntityName string
	QueryValue interface{}
}

type Storage interface {
	Init() error

	Entity(name string) *Db

	Query(query interface{}) *Db

	Into(value interface{}) error

	Add(value interface{}) error

	Update(value interface{}) error

	Delete(key string) error
}

func (r *Db) Entity(name string) *Db {
	r.EntityName = name
	return r
}

func (r *Db) Query(query interface{}) *Db {
	r.QueryValue = query
	return r
}

//func (r *Db) Into(value interface{}) error {
//	log.Panic(ErrNotSupport)
//	return nil
//}
//
//func (r *Db) Set(value interface{}) error {
//	log.Panic(ErrNotSupport)
//	return nil
//}
//
//func (r *Db) Delete(key string) error {
//	log.Panic(ErrNotSupport)
//	return nil
//}
