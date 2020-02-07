// +build integration

package integration

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/config"
	"github.com/trustwallet/blockatlas/pkg/tests/integration/bitcoin"
	"github.com/trustwallet/blockatlas/pkg/tests/integration/domains"
	"github.com/trustwallet/blockatlas/pkg/tests/integration/ontology"
	"github.com/trustwallet/blockatlas/platform"
	"os"
	"testing"
)

func Test(t *testing.T) {
	configPath := os.Getenv("TEST_CONFIG")
	if configPath == "" {
		config.LoadConfig("../../../config.yml")
	} else {
		config.LoadConfig(configPath)
	}
	platform.Init(viper.GetString("platform"))

	// Add your integration tests here
	ontology.TestOntology(t)
	bitcoin.TestBitcoin(t)
	domains.TestDomains(t)
}
