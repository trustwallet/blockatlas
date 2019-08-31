// +build integration

package integration

import (
	"github.com/trustwallet/blockatlas/integration/config"
	"github.com/trustwallet/blockatlas/integration/tester"
	"testing"
)

func TestApis(t *testing.T) {
	config.InitConfig()
	tester.Tester(t)
}
