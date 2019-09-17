package blockatlas

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Request struct {
	HttpClient   *http.Client
	ErrorHandler func(res *http.Response, uri string) error
}

func (r *Request) Get(result interface{}, base string, path string, query url.Values) error {
	var queryStr = ""
	if query != nil {
		queryStr = query.Encode()
	}
	uri := strings.Join([]string{fmt.Sprintf("%s/%s", base, path), queryStr}, "?")
	return r.Execute("GET", uri, nil, result)
}

func (r *Request) Execute(method string, url string, body io.Reader, result interface{}) error {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}
	res, err := r.HttpClient.Do(req)
	if err != nil {
		return err
	}
	err = r.ErrorHandler(res, url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(result)
}
