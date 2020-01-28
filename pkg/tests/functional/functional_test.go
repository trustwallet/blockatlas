// +build functional

package functional

import (
	"os"
	"sync"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/cmd"
	"github.com/trustwallet/blockatlas/config"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/platform"
)

func TestApis(t *testing.T) {
	_ = os.Setenv("ATLAS_GIN_MODE", "debug")
	config.LoadConfig(os.Getenv("TEST_CONFIG"))

	logger.InitLogger()
	platform.Init()

	p := ":8080"
	c := make(chan *gin.Engine)
	go func() {
		cmd.RunApi(p, c)
	}()
	e := <-c
	time.Sleep(time.Second * 2)

	var wg sync.WaitGroup
	cl := newClient(t, p)
	for _, r := range e.Routes() {
		wg.Add(1)
		t.Run(r.Path, func(t *testing.T) {
			go cl.doTests(r.Method, r.Path, &wg)
		})
	}
	wg.Wait()
}
