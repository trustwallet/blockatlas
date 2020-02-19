package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/metrics"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Request struct {
	BaseUrl      string
	Headers      map[string]string
	HttpClient   *http.Client
	ErrorHandler func(res *http.Response, uri string) error
}

func (r *Request) SetTimeout(seconds time.Duration) {
	r.HttpClient.Timeout = time.Second * seconds
}

func InitClient(baseUrl string) Request {
	return Request{
		Headers:      make(map[string]string),
		HttpClient:   DefaultClient,
		ErrorHandler: DefaultErrorHandler,
		BaseUrl:      baseUrl,
	}
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
	uri := strings.Join([]string{r.GetBase(path), queryStr}, "?")
	return r.Execute("GET", uri, nil, result, errors.Params{"path": path})
}

func (r *Request) Post(result interface{}, path string, body interface{}) error {
	buf, err := GetBody(body)
	if err != nil {
		return err
	}
	uri := r.GetBase(path)
	return r.Execute("POST", uri, buf, result, errors.Params{"path": path})
}

func (r *Request) Execute(method string, url string, body io.Reader, result interface{}, params errors.Params) error {
	params["method"] = method
	start := time.Now()
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return errors.E(err, errors.TypePlatformRequest, params)
	}

	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}

	res, err := r.HttpClient.Do(req)
	if err != nil {
		return errors.E(err, errors.TypePlatformRequest, params)
	}
	go metrics.GetMetrics(res.Status, url, method, start)

	err = r.ErrorHandler(res, url)
	if err != nil {
		return errors.E(err, errors.TypePlatformError, params)
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.E(err, errors.TypePlatformUnmarshal, params)
	}
	err = json.Unmarshal(b, result)
	if err != nil {
		return errors.E(err, errors.TypePlatformUnmarshal, params)
	}
	return err
}

func (r *Request) GetBase(path string) string {
	if path == "" {
		return fmt.Sprintf("%s", r.BaseUrl)
	}
	return fmt.Sprintf("%s/%s", r.BaseUrl, path)
}

func GetBody(body interface{}) (buf io.ReadWriter, err error) {
	if body != nil {
		buf = new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(body)
	}
	return
}
