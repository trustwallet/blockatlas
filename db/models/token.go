package models

import (
	"github.com/jinzhu/gorm"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Token struct {
	gorm.Model
	Address  string `gorm:"type:varchar(256);unique_index"`
	Coin     uint   `sql:"index"`
	Decimals int
	Name     string               `gorm:"type:varchar(64)" sql:"index"`
	Symbol   string               `gorm:"type:varchar(16)" sql:"index"`
	Type     blockatlas.TokenType `gorm:"type:varchar(16)" sql:"index"`
	TokenID  string               `gorm:"type:varchar(256)" sql:"index"`
	AssetID  string               `gorm:"type:varchar(256);unique_index"`
}
