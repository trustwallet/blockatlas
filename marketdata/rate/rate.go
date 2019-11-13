package rate

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/storage"
	"time"
)

const (
	defaultUpdateTime = time.Second * 20
)

type Rate struct {
	blockatlas.Request
	Id         string
	UpdateTime time.Duration
	Storage    storage.Market
}

func (r *Rate) GetUpdateTime() time.Duration {
	return r.UpdateTime
}

func (r *Rate) GetId() string {
	return r.Id
}

func (r *Rate) GetLogType() string {
	return "market-rate"
}

func (r *Rate) Init(storage storage.Market) error {
	logger.Info("Init Provider", logger.Params{"rate": r.GetId()})
	if len(r.Id) == 0 {
		return errors.E("Provider: Id cannot be empty")
	}

	if storage == nil {
		return errors.E("Provider: Storage cannot be nil")
	}
	r.Storage = storage

	if r.UpdateTime == 0 {
		r.UpdateTime = defaultUpdateTime
	}
	return nil
}
