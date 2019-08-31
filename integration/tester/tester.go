package tester

import (
	"fmt"
	log "github.com/trustwallet/blockatlas/integration/logger"
	"strings"
	"sync"
	"testing"
)

func Tester(t *testing.T) {
	files, err := GetFilesFromFolder()
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
		tests, err := GetHttpTests(f)
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

func doTests(t *testing.T, test HttpTest, coins []Coin) {
	c := NewClient(t)
	var wg sync.WaitGroup
	wg.Add(len(coins))
	for _, coin := range coins {
		switch strings.ToUpper(test.Method) {
		case "GET":
			go c.TestGet(coin, test, &wg)
		case "POST":
			go c.TestPost(coin, test, &wg)
		default:
			log.Error(fmt.Sprintf("Unrecognized method: %s", test.Method))
			wg.Done()
		}
	}
	wg.Wait()
}
