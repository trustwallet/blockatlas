// +build coins
//go:generate go run coinGenerator.go

package main

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v2"
	"html/template"
	"log"
	"os"
)

const (
	coinFile  = "coins.yml"
	filename  = "coins.go"
	codeBegin = `package coin

var coins = []Coin{`
	codeEnd = `
}`
	codeCoinTemplate = `
	{
		ID:               {{.ID}},
		Handle:           "{{.Handle}}",
		Symbol:           "{{.Symbol}}",
		Title:            "{{.Title}}",
		Decimals:         {{.Decimals}},
		BlockTime:        {{.BlockTime}},
		MinConfirmations: {{.MinConfirmations}},
		SampleAddr:       "{{.SampleAddr}}",
	},`
)

type Coin struct {
	ID               uint   `yaml:"id"`
	Handle           string `yaml:"handle"`
	Symbol           string `yaml:"symbol"`
	Title            string `yaml:"name"`
	Decimals         uint   `yaml:"decimals"`
	BlockTime        int    `yaml:"blockTime"`
	MinConfirmations int64  `yaml:"minConfirmations"`
	SampleAddr       string `yaml:"sampleAddress"`
}

func main() {
	var coinList []Coin
	coin, err := os.Open(coinFile)
	dec := yaml.NewDecoder(coin)
	err = dec.Decode(&coinList)
	if err != nil {
		log.Panic(err)
		return
	}
	code := generateCode(coinList)
	fmt.Println(code)
	saveGoFile(code)
}

func saveGoFile(code string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Panic(err)
		return
	}
	l, err := f.WriteString(code)
	if err != nil {
		log.Panic(err)
		f.Close()
		return
	}
	fmt.Println(l, "Go file save successfully!!!")
	err = f.Close()
	if err != nil {
		log.Panic(err)
		return
	}
}

func generateCode(coinList []Coin) string {
	code := codeBegin
	for _, coin := range coinList {
		temp, err := fillTemplate(coin)
		if err != nil {
			continue
		}
		code += temp
	}
	code += codeEnd
	return code
}

func fillTemplate(coin Coin) (string, error) {
	tpl := template.New("url")
	tpl, err := tpl.Parse(codeCoinTemplate)
	if err != nil {
		return "", err
	}
	var out bytes.Buffer
	err = tpl.Execute(&out, coin)
	if err != nil {
		return "", err
	}
	u := out.String()
	return u, nil
}
