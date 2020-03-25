package models

import "github.com/jinzhu/gorm"

type Subscription struct {
	gorm.Model
	GUID    string `json:"guid" sql:"index" gorm:"ForeignKey:UserId; not null"`
	Coin    uint   `json:"coin" sql:"index"`
	Address string `json:"address" sql:"index"`
}
