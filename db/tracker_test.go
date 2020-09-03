package db

import (
	"context"
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
			`INSERT INTO "trackers" ("updated_at","coin","height") VALUES ($1,$2,$3) ON CONFLICT (coin) DO UPDATE SET height = excluded.height, updated_at = excluded.updated_at RETURNING "trackers"."coin"`)).WithArgs(sqlmock.AnyArg(), "bitcoin", 1).WillReturnRows(sqlmock.NewRows([]string{"id"}).
		AddRow("id"))
	mock.ExpectCommit()
	i := Instance{Gorm: db, GormRead: db}

	assert.Nil(t, i.SetLastParsedBlockNumber("bitcoin", 1, context.Background()))
}

func TestHeightBlockMap_GetHeight(t *testing.T) {
	db, mock := setupDB(t)
	defer db.Close()
	mock.ExpectBegin()
	mock.ExpectQuery(
		regexp.QuoteMeta(
			`INSERT INTO "trackers" ("updated_at","coin","height") VALUES ($1,$2,$3) ON CONFLICT (coin) DO UPDATE SET height = excluded.height, updated_at = excluded.updated_at RETURNING "trackers"."coin"`)).WithArgs(sqlmock.AnyArg(), "bitcoin", 1).WillReturnRows(sqlmock.NewRows([]string{"id"}).
		AddRow("id"))
	mock.ExpectCommit()
	i := Instance{Gorm: db, GormRead: db}

	assert.Nil(t, i.SetLastParsedBlockNumber("bitcoin", 1, context.Background()))
	block, err := i.GetLastParsedBlockNumber("bitcoin", context.Background())
	assert.Nil(t, err)
	assert.Equal(t, int64(1), block)

	mock.ExpectQuery(
		regexp.QuoteMeta(
			`SELECT * FROM "trackers"  WHERE ("trackers"."coin" = $1`)).WithArgs("ethereum").WillReturnRows(sqlmock.NewRows([]string{"coin", "height"}).
		AddRow("ethereum", 1))

	b, err := i.GetLastParsedBlockNumber("ethereum", context.Background())
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
