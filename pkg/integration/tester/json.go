// +build integration

package tester

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

type HttpTest struct {
	Version     string                 `json:"version"`
	Path        string                 `json:"path"`
	Method      string                 `json:"method"`
	QueryString string                 `json:"query_string"`
	HttpCode    int                    `json:"http_code"`
	Body        map[string]interface{} `json:"body,omitempty"`
}

func GetTests() ([][]HttpTest, error) {
	files, err := GetFilesFromFolder()
	if err != nil {
		return nil, err
	}
	tests := make([][]HttpTest, 0)
	for _, f := range files {
		t, err := GetHttpTests(f)
		if err != nil {
			continue
		}
		tests = append(tests, t)
	}
	return tests, nil
}

func GetFilesFromFolder() ([]string, error) {
	files, err := ioutil.ReadDir("testdata")
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
	golden := filepath.Join("testdata", file)
	byteValue, err := ioutil.ReadFile(golden)
	if err != nil {
		return nil, err
	}
	var repo []HttpTest
	err = json.Unmarshal(byteValue[:], &repo)
	return repo, err
}
