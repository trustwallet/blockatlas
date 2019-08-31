package integration

import (
	"fmt"
	"github.com/trustwallet/blockatlas/integration/config"
	log "github.com/trustwallet/blockatlas/integration/logger"
	"github.com/trustwallet/blockatlas/integration/tester"
	"strings"
	"sync"
	"testing"
)

func TestApis(t *testing.T) {
	config.InitConfig()
	files, err := tester.GetFilesFromFolder()
	if err != nil {
		log.Error(err)
		return
	}
	log.GetFiles(len(files))
	coins, err := GetCoins()
	if err != nil {
		log.Error(err)
		return
	}
	log.GetCoins(len(coins))
	for _, f := range files {
		log.FileTest(f)
		tests, err := tester.GetHttpTests(f)
		if err != nil {
			log.Error(err)
			continue
		}
		for _, test := range tests {
			log.Test(test.Path, test.Method, test.HttpCode)
			doTests(t, test, coins)
		}
	}
}

func doTests(t *testing.T, test tester.HttpTest, coins []Coin) {
	c := tester.NewClient(t)
	var wg sync.WaitGroup
	wg.Add(len(coins))
	for _, coin := range coins {
		switch strings.ToUpper(test.Method) {
		case "GET":
			go c.TestGet(coin.Handle, coin.SampleAddr, test, &wg)
		case "POST":
			go c.TestPost(coin.Handle, coin.SampleAddr, test, &wg)
		default:
			log.Error(fmt.Sprintf("Unrecognized method: %s", test.Method))
			wg.Done()
		}
	}
	wg.Wait()
}
