package sql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"time"
)

type PgSql struct {
	sql
}

func (db *PgSql) Init(host string) {
	client, err := gorm.Open("postgres", host)
	if err != nil {
		logger.Fatal(err, "postgress connection failed")
	}
	client.DB().SetMaxIdleConns(5)
	client.DB().SetMaxOpenConns(50)
	client.DB().SetConnMaxLifetime(time.Second * 20)
	db.Client = client
}

func (db *PgSql) IsReady() bool {
	return db.Client != nil
}
