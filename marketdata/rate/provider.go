package rate

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/storage"
	"time"
)

type Provider interface {
	Init(storage.Market) error
	FetchLatestRates() (blockatlas.Rates, error)
	GetUpdateTime() time.Duration
	GetId() string
	GetType() string
}
