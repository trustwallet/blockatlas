package fio

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Client struct {
	blockatlas.Request
}

func InitClient(baseUrl string) Client {
	return Client{
		Request: blockatlas.Request{
			HttpClient:   blockatlas.DefaultClient,
			ErrorHandler: blockatlas.DefaultErrorHandler,
			BaseUrl:      baseUrl,
		},
	}
}
