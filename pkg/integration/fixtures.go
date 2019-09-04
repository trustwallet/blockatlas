// +build integration

package integration

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/coin"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	fixtures = "fixtures.json"
)

type Fixture map[string]map[string]string

var f Fixture

func init() {
	var err error
	f, err = getFixtures()
	if err != nil {
		logrus.Panic(err)
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

func getFile(file string) ([]byte, error) {
	golden := filepath.Join("testdata", file)
	return ioutil.ReadFile(golden)
}

func getCoin(path string) coin.Coin {
	for _, coin := range coin.Coins {
		if strings.Contains(path, coin.Handle) {
			return coin
		}
	}
	return coin.Coin{}
}

func addFixtures(path string) string {
	c := getCoin(path)
	if (c == coin.Coin{}) {
		return path
	}
	fix, ok := f[strconv.Itoa(int(c.ID))]
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
