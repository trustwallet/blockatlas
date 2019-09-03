// +build integration

package integration

import (
	"github.com/trustwallet/blockatlas/pkg/integration/config"
	"github.com/trustwallet/blockatlas/pkg/integration/tester"
	"testing"
)

func TestApis(t *testing.T) {
	config.InitConfig()
	apis, err := tester.GetApis()
	if err != nil {
		t.Error(err)
		return
	}
	coins, err := tester.GetCoins()
	if err != nil {
		t.Error(err)
		return
	}
	for _, coin := range coins {
		tester.DoTests(t, apis, coin)
	}
}
