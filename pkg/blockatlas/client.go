package blockatlas

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
		return errors.E(err, errors.TypePlatformRequest, errors.Params{"url": url, "method": method}).PushToSentry()
	}

	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}

	res, err := r.HttpClient.Do(req)
	if err != nil {
		return errors.E(err, errors.TypePlatformRequest, errors.Params{"url": url, "method": method}).PushToSentry()
	}
	go metrics.GetMetrics(res.Status, url, method, start)

	err = r.ErrorHandler(res, url)
	if err != nil {
		return errors.E(err, errors.TypePlatformError, errors.Params{"url": url, "method": method}).PushToSentry()
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.E(err, errors.TypePlatformUnmarshal, errors.Params{"url": url, "method": method}).PushToSentry()
	}
	err = json.Unmarshal(b, result)
	if err != nil {
		return errors.E(err, errors.TypePlatformUnmarshal, errors.Params{"url": url, "method": method}).PushToSentry()
	}
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
