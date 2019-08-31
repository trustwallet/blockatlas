package tester

import (
	"fmt"
	"github.com/gavv/httpexpect"
	"github.com/trustwallet/blockatlas/integration/config"
	log "github.com/trustwallet/blockatlas/integration/logger"
	"sync"
	"testing"
	"time"
)

type Client struct {
	e *httpexpect.Expect
	t *testing.T
}

func NewClient(t *testing.T) *Client {
	return &Client{
		httpexpect.New(t, config.Configuration.Server.Url),
		t,
	}
}

func (c *Client) TestPost(coin Coin, test HttpTest, wg *sync.WaitGroup) {
	defer wg.Done()

	url := getBaseUrl(test.Version, coin.Handle, test.Path)
	defer TimeTrack(coin.Symbol, test.Method, url, time.Now())
	request := c.e.POST(url)
	if test.Body != nil {
		request.WithJSON(test.Body)
	}
	response := request.Expect()
	response.Status(test.HttpCode)
}

func (c *Client) TestGet(coin Coin, test HttpTest, wg *sync.WaitGroup) {
	defer wg.Done()

	q, err := getParameters(test.QueryString, coin.SampleAddr)
	if err != nil {
		c.t.Error(err)
	}
	url := getParameterUrl(test.Version, coin.Handle, test.Path, q)
	defer TimeTrack(coin.Symbol, test.Method, url, time.Now())
	request := c.e.GET(url)
	response := request.Expect()
	response.Status(test.HttpCode)
}

func getParameterUrl(version, coin, path, params string) string {
	return fmt.Sprintf("%s%s", getBaseUrl(version, coin, path), params)
}

func getBaseUrl(version, coin, path string) string {
	return fmt.Sprintf("%s/%s/%s%s", config.Configuration.Server.Url, version, coin, path)
}

func TimeTrack(name, method, url string, start time.Time) time.Duration {
	elapsed := time.Since(start)
	log.TimeTrack(name, method, url, elapsed)
	return time.Since(start)
}
