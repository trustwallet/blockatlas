package tester

import (
	"encoding/json"
	"fmt"
	"github.com/trustwallet/blockatlas/integration/config"
	log "github.com/trustwallet/blockatlas/integration/logger"
	"io/ioutil"
	"os"
)

type HttpTest struct {
	Version     string                 `json:"version"`
	Path        string                 `json:"path"`
	Method      string                 `json:"method"`
	QueryString string                 `json:"query_string"`
	HttpCode    int                    `json:"http_code"`
	Body        map[string]interface{} `json:"body,omitempty"`
}

func GetFilesFromFolder() ([]string, error) {
	files, err := ioutil.ReadDir(config.Configuration.Json.Path)
	if err != nil {
		return nil, err
	}

	list := make([]string, 0)
	for _, f := range files {
		list = append(list, f.Name())
	}
	return list, nil
}

func GetHttpTests(file string) ([]HttpTest, error) {
	jsonFile, err := os.Open(fmt.Sprintf("%s/%s", config.Configuration.Json.Path, file))
	if err != nil {
		return nil, err
	}
	defer func() {
		err := jsonFile.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var repo []HttpTest
	err = json.Unmarshal(byteValue[:], &repo)
	return repo, err
}
