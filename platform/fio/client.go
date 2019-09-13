package fio

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas"
	"net/http"
)

type Client struct {
	Request blockatlas.Request
	URL     string
}

func InitClient() Client {
	return Client{
		Request: blockatlas.Request{
			HttpClient: http.DefaultClient,
			ErrorHandler: func(res *http.Response, uri string) error {
				return nil
			},
		},
		URL: viper.GetString("fio.api"),
	}
}
