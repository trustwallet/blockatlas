package coin

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

const (
	coinFile = "../coins.yml"
	filename = "coins.go"
)

type TestCoin struct {
	ID               uint   `yaml:"id"`
	Handle           string `yaml:"handle"`
	Symbol           string `yaml:"symbol"`
	Title            string `yaml:"name"`
	Decimals         uint   `yaml:"decimals"`
	BlockTime        int    `yaml:"blockTime"`
	MinConfirmations int64  `yaml:"minConfirmations"`
	SampleAddr       string `yaml:"sampleAddress"`
}

func TestFilesExists(t *testing.T) {
	assert.True(t, assert.FileExists(t, coinFile))
	assert.True(t, assert.FileExists(t, filename))
}

func TestCoinFile(t *testing.T) {
	var coinList []TestCoin
	coin, err := os.Open(coinFile)
	dec := yaml.NewDecoder(coin)
	err = dec.Decode(&coinList)
	if err != nil {
		t.Error(err)
	}

	f, err := os.Open(filename)
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	code := string(b)

	for _, want := range coinList {
		got, ok := Coins[want.ID]
		assert.True(t, ok)
		assert.Equal(t, got.ID, want.ID)
		assert.Equal(t, got.Handle, want.Handle)
		assert.Equal(t, got.Symbol, want.Symbol)
		assert.Equal(t, got.Title, want.Title)
		assert.Equal(t, got.Decimals, want.Decimals)
		assert.Equal(t, got.BlockTime, want.BlockTime)
		assert.Equal(t, got.MinConfirmations, want.MinConfirmations)
		assert.Equal(t, got.SampleAddr, want.SampleAddr)

		s := strings.Title(want.Handle)
		method := fmt.Sprintf("func %s() Coin", s)
		assert.True(t, strings.Contains(code, method), "Coin method not found")
		enum := fmt.Sprintf("%s = %d", want.Symbol, want.ID)
		assert.True(t, strings.Contains(code, enum), "Coin enum not found")
	}
}
