// +build integration

package integration

import (
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/cmd/api"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/config"
	"github.com/trustwallet/blockatlas/platform"
	"os"
	"sync"
	"testing"
	"time"
)

func TestApis(t *testing.T) {
	config.LoadConfig(os.Getenv("TEST_CONFIG"))
	coin.Load(os.Getenv("TEST_COINS"))
	platform.Init()

	p := ":8080"
	c := make(chan *gin.Engine)
	go func() {
		api.Run(p, c)
	}()
	e := <-c
	time.Sleep(time.Second * 2)

	var wg sync.WaitGroup
	cl := newClient(t, p)
	for _, r := range e.Routes() {
		wg.Add(1)
		go cl.doTests(r.Path, &wg)
	}
	wg.Wait()
}
