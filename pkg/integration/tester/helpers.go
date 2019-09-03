// +build integration

package tester

import (
	"bytes"
	"html/template"
)

func getParameters(url, coin, address string) (string, error) {
	tpl := template.New("url")
	tpl, err := tpl.Parse(url)
	if err != nil {
		return "", err
	}

	data := struct {
		Address string
		Coin    string
	}{
		address,
		coin,
	}

	var out bytes.Buffer
	err = tpl.Execute(&out, data)
	if err != nil {
		return "", err
	}
	u := out.String()
	return u, nil
}
