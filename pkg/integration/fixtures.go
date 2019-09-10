// +build integration

package integration

import (
	"encoding/json"
	"fmt"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"io/ioutil"
	"path/filepath"
	"strings"
)

const (
	fixtures = "fixtures.json"
	exclude  = "exclude.json"
)

type Fixture map[string]map[string]string
type Exclude []string

var f Fixture
var e Exclude

func init() {
	var err error
	f, err = getFixtures()
	if err != nil {
		logger.Panic(err)
	}
	e, err = getExcludeApis()
	if err != nil {
		logger.Panic(err)
	}
}

func getFixtures() (Fixture, error) {
	b, err := getFile(fixtures)
	if err != nil {
		return nil, err
	}
	var r Fixture
	err = json.Unmarshal(b[:], &r)
	return r, err
}

func isExcluded(path string) bool {
	return contains(e, path)
}

func getExcludeApis() (Exclude, error) {
	b, err := getFile(exclude)
	if err != nil {
		return nil, err
	}
	var r Exclude
	err = json.Unmarshal(b[:], &r)
	return r, err
}

func getFile(file string) ([]byte, error) {
	golden := filepath.Join("testdata", file)
	return ioutil.ReadFile(golden)
}

func getCoin(path string) coin.Coin {
	for _, c := range coin.Coins {
		if strings.Contains(path, fmt.Sprintf("/%s/", c.Handle)) {
			return c
		}
	}
	return coin.Coin{}
}

func addFixtures(path string) string {
	c := getCoin(path)
	if (c == coin.Coin{}) {
		return path
	}
	fix, ok := f[c.Handle]
	if !ok {
		return strings.Replace(path, ":address", c.SampleAddr, -1)
	}
	if _, ok := fix["address"]; !ok {
		return strings.Replace(path, ":address", c.SampleAddr, -1)
	}
	result := path
	for key, value := range fix {
		result = strings.Replace(result, ":"+key, value, -1)
	}
	return result
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
