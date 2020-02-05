// +build integration

package integration

import (
	"github.com/trustwallet/blockatlas/config"
	"github.com/trustwallet/blockatlas/pkg/tests/integration/domains"
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
	platform.Init()

	//ontology.TestOntology(t)
	//bitcoin.TestBitcoin(t)
	domains.TestDomains(t)
}
