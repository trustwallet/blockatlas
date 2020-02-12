package rate

import (
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/storage"
)

const (
	defaultUpdateTime = "5m"
)

type Rate struct {
	Id         string
	UpdateTime string
	Storage    storage.Market
}

func (r *Rate) GetUpdateTime() string {
	return r.UpdateTime
}

func (r *Rate) GetId() string {
	return r.Id
}

func (r *Rate) GetLogType() string {
	return "market-rate"
}

func (r *Rate) Init(storage storage.Market) error {
	logger.Info("Init Market Rate Provider", logger.Params{"rate": r.GetId()})
	if len(r.Id) == 0 {
		return errors.E("Market Rate: Id cannot be empty")
	}

	if storage == nil {
		return errors.E("Market Rate: Storage cannot be nil")
	}
	r.Storage = storage

	if len(r.UpdateTime) == 0 {
		r.UpdateTime = defaultUpdateTime
	}
	return nil
}
