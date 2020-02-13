package market

import (
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/storage"
)

const (
	defaultUpdateTime = "5m"
)

type Market struct {
	Id         string
	UpdateTime string
	Storage    storage.Market
}

func (m *Market) GetId() string {
	return m.Id
}

func (m *Market) GetLogType() string {
	return "market-data"
}

func (m *Market) GetUpdateTime() string {
	return m.UpdateTime
}

func (m *Market) Init(storage storage.Market) error {
	logger.Info("Init Market Quote Provider", logger.Params{"market": m.GetId()})
	if len(m.Id) == 0 {
		return errors.E("Market Quote: Id cannot be empty")
	}

	if storage == nil {
		return errors.E("Market Quote: Storage cannot be nil")
	}
	m.Storage = storage

	if len(m.UpdateTime) == 0 {
		m.UpdateTime = defaultUpdateTime
	}
	return nil
}
