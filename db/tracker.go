package db

import (
	"github.com/jinzhu/gorm"
	"github.com/trustwallet/blockatlas/db/models"
	"sync"
)

var memoryCache heightBlockMap

func init() {
	memoryCache.m = make(map[uint]int64, 0)
}

type heightBlockMap struct {
	m map[uint]int64
	sync.RWMutex
}

func (hbm *heightBlockMap) SetHeight(coin uint, b int64) {
	hbm.Lock()
	defer hbm.Unlock()
	hbm.m[coin] = b
}

func (hbm *heightBlockMap) GetHeight(coin uint) (int64, bool) {
	hbm.RLock()
	defer hbm.RUnlock()
	b, ok := hbm.m[coin]
	return b, ok
}

func GetLastParsedBlockNumber(dbConn *gorm.DB, coin uint) (int64, error) {
	height, ok := memoryCache.GetHeight(coin)
	if ok {
		return height, nil
	}
	var tracker models.Tracker
	if err := dbConn.Where(models.Tracker{Coin: coin}).Find(&tracker).Error; err != nil {
		return 0, nil
	}
	return tracker.Height, nil
}

func SetLastParsedBlockNumber(dbConn *gorm.DB, coin uint, num int64) error {
	memoryCache.SetHeight(coin, num)
	tracker := models.Tracker{
		Coin:   coin,
		Height: num,
	}

	return dbConn.
		Set("gorm:insert_option", "ON CONFLICT (coin) DO UPDATE SET height = excluded.height").
		Where(models.Tracker{Coin: coin}).
		Create(&tracker).Error
}
