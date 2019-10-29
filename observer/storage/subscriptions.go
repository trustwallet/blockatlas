package storage

import (
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/storage/sql"
)

func (s *Storage) Lookup(addresses ...string) (observers []Subscription, err error) {
	if len(addresses) == 0 {
		err = errors.E("cannot look up an empty list")
		return
	}
	err = sql.Find(s.Client, &observers, "address = ?", addresses)
	return
}

func (s *Storage) AddSubscriptions(subscriptions []interface{}) error {
	return s.MustAddMany(subscriptions...)
}

func (s *Storage) DeleteSubscriptions(subscriptions []interface{}) error {
	return s.MustDeleteMany(subscriptions...)
}
