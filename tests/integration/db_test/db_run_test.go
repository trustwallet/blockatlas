// +build integration

package db_test

import (
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/tests/integration/setup"
	"os"
	"testing"
)

var dbConn *gorm.DB

func TestMain(m *testing.M) {
	dbConn = setup.RunPgContainer()
	code := m.Run()
	setup.StopPgContainer()
	os.Exit(code)
}

func TestPgSetup(t *testing.T) {
	assert.NotNil(t, dbConn)
}
