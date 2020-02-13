// +build functional

package functional

import (
	"encoding/json"
	"fmt"
	"github.com/Pantani/httpexpect"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"net/http"
	"sync"
	"testing"
	"time"
)

const (
	baseUrl = "http://localhost%s"
)

type Client struct {
	baseUrl string
	e       *httpexpect.Expect
	t       *testing.T
}

func newClient(t *testing.T, port string) *Client {
	client := httpexpect.WithConfig(httpexpect.Config{
		BaseURL: getBaseUrl(port),
		Client: &http.Client{
			Jar:     httpexpect.NewJar(),
			Timeout: time.Second * 30,
		},
		// use fatal failures
		Reporter: httpexpect.NewAssertReporter(t),
		// use verbose logging
		Printers: []httpexpect.Printer{},
	})
	return &Client{
		baseUrl: getBaseUrl(port),
		e:       client,
		t:       t,
	}
}

func (c *Client) testGet(route string, query string) {
	request := c.e.GET(route).WithURL(c.baseUrl)
	request.WithQueryString(query)
	response := request.Expect()

	if response == nil || response.Raw() == nil {
		logger.Error("Invalid response", logger.Params{"response": response, "route": route, "query": query})
	}
	if response.Raw().StatusCode != http.StatusOK {
		logger.Error("Invalid status code", logger.Params{"code": response.Raw().Status, "route": route, "query": query})
	}
	response.Status(http.StatusOK)
}

func (c *Client) testPost(route string, body interface{}) {
	request := c.e.POST(route).WithURL(c.baseUrl)
	if body == nil {
		request.WithText("[]")
	} else {
		b, err := json.Marshal(body)
		if err == nil && b != nil {
			request.WithText(string(b))
		}
	}
	response := request.Expect()
	if response == nil || response.Raw() == nil {
		logger.Error("Invalid response", logger.Params{"code": response.Raw().Status, "route": route})
	}
	if response.Raw().StatusCode != http.StatusOK {
		bodyJson, _ := json.Marshal(body)
		logger.Error("Invalid status code", logger.Params{"code": response.Raw().Status, "route": route, "body": bodyJson})
	}
	response.Status(http.StatusOK)
}

func (c *Client) doTests(method, path string, wg *sync.WaitGroup) {
	defer wg.Done()
	if isExcluded(path) {
		return
	}
	url := addCoinFixtures(path)
	switch method {
	case "GET":
		tests := getQueryTests(path)
		for _, query := range tests {
			c.testGet(url, query)
		}
	case "POST":
		tests := getBodyTests(path)
		for _, body := range tests {
			c.testPost(url, body)
		}
	}
}

func getBaseUrl(port string) string {
	return fmt.Sprintf(baseUrl, port)
}
