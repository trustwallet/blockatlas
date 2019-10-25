package storage

import (
	"github.com/trustwallet/blockatlas/pkg/errors"
)

func (s *Storage) Lookup(coin uint, addresses ...string) (observers []Subscription, err error) {
	if len(addresses) == 0 {
		err = errors.E("cannot look up an empty list", errors.Params{"coin": coin}).PushToSentry()
		return
	}
	s.Client.
		Table("subscriptions s1").
		Select("DISTINCT s1.coin, s1.address, s1.webhook, s1.xpub, s1.created_at").
		Joins("LEFT JOIN subscriptions s2 ON s2.xpub = s1.xpub").
		Where("s1.address IN (?)", addresses).
		Or("s2.address IN (?)", addresses).
		Find(&observers)
	return
}

func (s *Storage) AddSubscriptions(subscriptions []interface{}) error {
	return s.MustAddMany(subscriptions...)
}

func (s *Storage) DeleteSubscriptions(subscriptions []interface{}) error {
	return s.MustDeleteMany(subscriptions...)
}
