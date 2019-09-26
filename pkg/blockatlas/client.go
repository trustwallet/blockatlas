package blockatlas

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/metrics"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Request struct {
	BaseUrl      string
	HttpClient   *http.Client
	ErrorHandler func(res *http.Response, uri string) error
}

var DefaultClient = &http.Client{
	Timeout: time.Second * 15,
}

var DefaultErrorHandler = func(res *http.Response, uri string) error {
	return nil
}

func (r *Request) Get(result interface{}, path string, query url.Values) error {
	var queryStr = ""
	if query != nil {
		queryStr = query.Encode()
	}
	uri := strings.Join([]string{r.getBase(path), queryStr}, "?")
	return r.Execute("GET", uri, nil, result)
}

func (r *Request) Post(result interface{}, path string, body interface{}) error {
	buf, err := getBody(body)
	if err != nil {
		return err
	}
	uri := r.getBase(path)
	return r.Execute("POST", uri, buf, result)
}

func (r *Request) Execute(method string, url string, body io.Reader, result interface{}) error {
	start := time.Now()

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}
	res, err := r.HttpClient.Do(req)
	if err != nil {
		return err
	}
	go metrics.GetMetrics(res.Status, url, method, start)

	err = r.ErrorHandler(res, url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(result)
	return err
}

func (r *Request) getBase(path string) string {
	return fmt.Sprintf("%s/%s", r.BaseUrl, path)
}

func getBody(body interface{}) (buf io.ReadWriter, err error) {
	if body != nil {
		buf = new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return
		}
	}
	return
}
