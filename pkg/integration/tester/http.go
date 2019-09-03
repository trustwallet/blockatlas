// +build integration

package tester

import (
	"github.com/gavv/httpexpect"
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/pkg/integration/config"
	"net/http"
	"sync"
	"testing"
	"time"
)

type Client struct {
	e *httpexpect.Expect
	t *testing.T
}

func NewClient(t *testing.T) *Client {
	http := httpexpect.WithConfig(httpexpect.Config{
		BaseURL: config.Configuration.Server.Url,
		Client: &http.Client{
			Jar:     httpexpect.NewJar(),
			Timeout: time.Second * 30,
		},
		// use fatal failures
		Reporter: httpexpect.NewRequireReporter(t),
		// use verbose logging
		Printers: []httpexpect.Printer{
			httpexpect.NewCurlPrinter(t),
		},
	})
	return &Client{
		http,
		t,
	}
}

func (c *Client) TestGet(coin, address string, test Api, wg *sync.WaitGroup) {
	defer wg.Done()

	path, err := getParameters(test.Path, coin, address)
	if err != nil {
		c.t.Error(err)
	}
	request := c.e.GET(path).WithURL(config.Configuration.Server.Url)

	t := time.Now()
	response := request.Expect()
	timeTrack(coin, address, path, t)

	response.Text().Schema(test.Schema)
	response.Status(http.StatusOK)
}

func DoTests(t *testing.T, apis map[string]Api, coin Coin) {
	c := NewClient(t)

	var wg sync.WaitGroup
	for _, coinApi := range coin.Apis {
		api, ok := apis[coinApi]
		if !ok {
			t.Errorf("invalid api %s for coin %s", coinApi, coin.Handle)
			continue
		}
		for _, addr := range coin.Addresses {
			wg.Add(1)
			go c.TestGet(coin.Handle, addr, api, &wg)
		}
	}
	wg.Wait()
}

func timeTrack(coin, address, path string, t time.Time) {
	logrus.WithFields(logrus.Fields{
		"coin":    coin,
		"address": address,
		"path":    path,
		"time":    time.Since(t).String(),
	}).Info("Test")
}
