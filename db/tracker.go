package db

import (
	"context"
	"github.com/trustwallet/blockatlas/db/models"
	"go.elastic.co/apm/module/apmgorm"
	"sync"
)

var memoryCache heightBlockMap

func init() {
	memoryCache.m = make(map[string]int64)
}

type heightBlockMap struct {
	m map[string]int64
	sync.RWMutex
}

func (hbm *heightBlockMap) SetHeight(coin string, b int64) {
	hbm.Lock()
	defer hbm.Unlock()
	hbm.m[coin] = b
}

func (hbm *heightBlockMap) GetHeight(coin string) (int64, bool) {
	hbm.RLock()
	defer hbm.RUnlock()
	b, ok := hbm.m[coin]
	return b, ok
}

func (i *Instance) GetLastParsedBlockNumber(coin string, ctx context.Context) (int64, error) {
	height, ok := memoryCache.GetHeight(coin)
	if ok {
		return height, nil
	}
	var tracker models.Tracker
	g := apmgorm.WithContext(ctx, i.GormRead)
	if err := g.Where(models.Tracker{Coin: coin}).Find(&tracker).Error; err != nil {
		return 0, nil
	}
	return tracker.Height, nil
}

func (i *Instance) SetLastParsedBlockNumber(coin string, num int64, ctx context.Context) error {
	memoryCache.SetHeight(coin, num)
	tracker := models.Tracker{
		Coin:   coin,
		Height: num,
	}
	g := apmgorm.WithContext(ctx, i.Gorm)
	return g.
		Set("gorm:insert_option", "ON CONFLICT (coin) DO UPDATE SET height = excluded.height, updated_at = excluded.updated_at").
		Where(models.Tracker{Coin: coin}).
		Create(&tracker).Error
}
