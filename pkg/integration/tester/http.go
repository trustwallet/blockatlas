// +build integration

package tester

import (
	"fmt"
	"github.com/gavv/httpexpect"
	"github.com/trustwallet/blockatlas/pkg/integration/config"
	"net/http"
	"testing"
	"time"
)

type HttpResult struct {
	Coin    string
	Method  string
	Version string
	Path    string
	Status  int
	Elapsed time.Duration
}

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
		},
	})
	return &Client{
		http,
		t,
	}
}

func (c *Client) TestPost(coin, address string, test HttpTest) HttpResult {
	url := getBaseUrl(test.Version, coin, test.Path)
	request := c.e.POST(url)
	if test.Body != nil {
		request.WithJSON(test.Body)
	}
	t := time.Now()
	response := request.Expect()
	elapsed := time.Since(t)
	status := response.Raw().StatusCode

	return HttpResult{
		Coin:    coin,
		Method:  test.Method,
		Version: test.Version,
		Path:    test.Path,
		Status:  status,
		Elapsed: elapsed,
	}
}

func (c *Client) TestGet(coin, address string, test HttpTest) HttpResult {
	q, err := getParameters(test.QueryString, address)
	if err != nil {
		c.t.Error(err)
	}
	url := getParameterUrl(test.Version, coin, test.Path, q)

	request := c.e.GET(url)
	t := time.Now()
	response := request.Expect()
	elapsed := time.Since(t)
	status := response.Raw().StatusCode

	return HttpResult{
		Coin:    coin,
		Method:  test.Method,
		Version: test.Version,
		Path:    test.Path,
		Status:  status,
		Elapsed: elapsed,
	}
}

func getParameterUrl(version, coin, path, params string) string {
	return fmt.Sprintf("%s%s", getBaseUrl(version, coin, path), params)
}

func getBaseUrl(version, coin, path string) string {
	return fmt.Sprintf("%s/%s/%s%s", config.Configuration.Server.Url, version, coin, path)
}
