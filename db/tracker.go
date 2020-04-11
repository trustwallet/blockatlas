package db

import (
	"github.com/trustwallet/blockatlas/db/models"
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

func (i *Instance) GetLastParsedBlockNumber(coin string) (int64, error) {
	height, ok := memoryCache.GetHeight(coin)
	if ok {
		return height, nil
	}
	var tracker models.Tracker
	if err := i.Gorm.Where(models.Tracker{Coin: coin}).Find(&tracker).Error; err != nil {
		return 0, nil
	}
	return tracker.Height, nil
}

func (i *Instance) SetLastParsedBlockNumber(coin string, num int64) error {
	memoryCache.SetHeight(coin, num)
	tracker := models.Tracker{
		Coin:   coin,
		Height: num,
	}

	return i.Gorm.
		Set("gorm:insert_option", "ON CONFLICT (coin) DO UPDATE SET height = excluded.height").
		Where(models.Tracker{Coin: coin}).
		Create(&tracker).Error
}
