package db

import (
	"github.com/trustwallet/blockatlas/db/models"
	"gorm.io/gorm/clause"
)

func (i *Instance) GetLastParsedBlockNumbers() ([]models.Tracker, error) {
	var trackers []models.Tracker
	if err := i.Gorm.
		Find(&trackers).Error; err != nil {
		return trackers, nil
	}
	return trackers, nil
}

func (i *Instance) GetLastParsedBlockNumber(coin string) (models.Tracker, error) {
	var tracker models.Tracker
	if err := i.Gorm.
		Find(&tracker, "coin = ?", coin).Error; err != nil {
		return tracker, err
	}
	return tracker, nil
}

func (i *Instance) SetLastParsedBlockNumber(coin string, num int64) error {
	tracker := models.Tracker{
		Coin:   coin,
		Height: num,
	}
	return i.Gorm.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{
				Name: "coin",
			},
		},
		DoUpdates: clause.AssignmentColumns([]string{"height", "updated_at"}),
	}).Create(&tracker).Error
}
