package fio

import (
	"net/http"

	"github.com/trustwallet/blockatlas"
)

type Client struct {
	Request blockatlas.Request
	URL     string
}

func InitClient(baseUrl string) Client {
	return Client{
		Request: blockatlas.Request{
			HttpClient: blockatlas.DefaultClient,
			ErrorHandler: func(res *http.Response, uri string) error {
				return nil
			},
		},
		URL: baseUrl,
	}
}
