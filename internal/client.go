package internal

import (
	"github.com/trustwallet/golibs/client"
	"github.com/trustwallet/golibs/network/middleware"
)

type Client struct {
	client.Request
}

func InitClient(url string) client.Request {
	return client.Request{
		Headers:      map[string]string{},
		HttpClient:   client.DefaultClient,
		ErrorHandler: middleware.SentryErrorHandler,
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
		ErrorHandler: middleware.SentryErrorHandler,
		BaseUrl:      baseUrl,
	}
}
