package internal

import (
	"net/http"
	"strconv"

	"github.com/getsentry/raven-go"

	log "github.com/sirupsen/logrus"

	"github.com/trustwallet/golibs/client"
)

type Client struct {
	client.Request
}

var errorHandler = func(res *http.Response, uri string) error {
	if res.StatusCode == http.StatusOK {
		return nil
	}
	log.WithFields(log.Fields{
		"tags": raven.Tags{
			{Key: "status_code", Value: strconv.Itoa(res.StatusCode)},
			{Key: "host", Value: res.Request.Host},
			{Key: "url", Value: uri},
		},
		"fingerprint": []string{"client_errors"},
	}).Error("Client Errors")

	return nil
}

func InitClient(url string) client.Request {
	return client.Request{
		Headers:      map[string]string{},
		HttpClient:   client.DefaultClient,
		ErrorHandler: errorHandler,
		BaseUrl:      url,
	}
}

func InitJSONClient(baseUrl string) client.Request {
	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}
	return client.Request{
		Headers:      headers,
		HttpClient:   client.DefaultClient,
		ErrorHandler: errorHandler,
		BaseUrl:      baseUrl,
	}
}
