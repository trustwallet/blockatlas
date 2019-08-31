package tester

import (
	"bytes"
	"html/template"
)

func getParameters(url, address string) (string, error) {
	tpl := template.New("url")
	tpl, err := tpl.Parse(url)
	if err != nil {
		return "", err
	}

	data := struct {
		Address string
	}{
		address,
	}

	var out bytes.Buffer
	err = tpl.Execute(&out, data)
	if err != nil {
		return "", err
	}
	u := out.String()
	return u, nil
}
