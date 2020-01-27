package bakingbad

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	baseURL        = "https://api.baking-bad.org"
	defaultTimeout = time.Second * 2
)

type (
	API struct {
		client *http.Client
	}
	Baker struct {
		Address      string  `json:"address"`
		FreeSpace    float64 `json:"freeSpace"`
		Fee          float64 `json:"fee"`
		EstimatedRoi float64 `json:"estimatedRoi"`
	}
)

func NewAPI() API {
	return API{
		client: &http.Client{
			Timeout: defaultTimeout,
		},
	}
}

func (api API) Bakers() (bakers []Baker, err error) {
	data, err := api.send("v2/bakers")
	if err != nil {
		return nil, fmt.Errorf("send: %s", err.Error())
	}
	err = json.Unmarshal(data, &bakers)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %s", err.Error())
	}
	return bakers, nil
}

func (api API) send(endpoint string) (data []byte, err error) {
	url := fmt.Sprintf("%s/%s", baseURL, endpoint)
	response, err := api.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("client.Get: %s", err.Error())
	}
	defer response.Body.Close()
	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll: %s", err.Error())
	}
	return data, nil
}
