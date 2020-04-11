package db

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestHeightBlockMap_SetHeight(t *testing.T) {
	db, mock := setupDB(t)
	defer db.Close()
	mock.ExpectBegin()
	mock.ExpectQuery(
		regexp.QuoteMeta(
			`INSERT INTO "trackers" ("coin","height") VALUES ($1,$2) ON CONFLICT (coin) DO UPDATE SET height = excluded.height RETURNING "trackers"."coin"`)).WithArgs("bitcoin", 1).WillReturnRows(sqlmock.NewRows([]string{"id"}).
		AddRow("id"))
	mock.ExpectCommit()
	i := Instance{Gorm: db}

	assert.Nil(t, i.SetLastParsedBlockNumber("bitcoin", 1))
}

func TestHeightBlockMap_GetHeight(t *testing.T) {
	db, mock := setupDB(t)
	defer db.Close()
	mock.ExpectBegin()
	mock.ExpectQuery(
		regexp.QuoteMeta(
			`INSERT INTO "trackers" ("coin","height") VALUES ($1,$2) ON CONFLICT (coin) DO UPDATE SET height = excluded.height RETURNING "trackers"."coin"`)).WithArgs("bitcoin", 1).WillReturnRows(sqlmock.NewRows([]string{"id"}).
		AddRow("id"))
	mock.ExpectCommit()
	i := Instance{Gorm: db}

	assert.Nil(t, i.SetLastParsedBlockNumber("bitcoin", 1))
	block, err := i.GetLastParsedBlockNumber("bitcoin")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), block)

	mock.ExpectQuery(
		regexp.QuoteMeta(
			`SELECT * FROM "trackers"  WHERE ("trackers"."coin" = $1`)).WithArgs("ethereum").WillReturnRows(sqlmock.NewRows([]string{"coin", "height"}).
		AddRow("ethereum", 1))

	b, err := i.GetLastParsedBlockNumber("ethereum")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), b)

}

func setupDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when sqlmock", err)
	}

	d, err := gorm.Open("postgres", db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	d.LogMode(true)
	return d, mock
}
