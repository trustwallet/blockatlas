// +build integration

package integration

import (
	"fmt"
	"github.com/gavv/httpexpect"
	"github.com/sirupsen/logrus"
	"net/http"
	"sync"
	"testing"
	"time"
)

const (
	baseUrl = "http://localhost%s"
	schema  = `{
		"docs": "array"
	}`
)

type Client struct {
	baseUrl string
	e       *httpexpect.Expect
	t       *testing.T
}

func newClient(t *testing.T, port string) *Client {
	http := httpexpect.WithConfig(httpexpect.Config{
		BaseURL: getBaseUrl(port),
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
		baseUrl: getBaseUrl(port),
		e:       http,
		t:       t,
	}
}

func (c *Client) testGet(url string) {
	request := c.e.GET(url).WithURL(c.baseUrl)
	t := time.Now()
	response := request.Expect()
	timeTrack(url, t)
	response.JSON().Schema(schema)
	response.Status(http.StatusOK)
}

func (c *Client) doTests(path string, wg *sync.WaitGroup) {
	defer wg.Done()
	url := addFixtures(path)
	c.testGet(url)
}

func getBaseUrl(port string) string {
	return fmt.Sprintf(baseUrl, port)
}

func timeTrack(url string, t time.Time) {
	logrus.WithFields(logrus.Fields{
		"url":  url,
		"time": time.Since(t).String(),
	}).Info("Test")
}
