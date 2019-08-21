package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func Request(c *http.Client, base string, path string, params url.Values, result interface{}) error {

	uri := fmt.Sprintf("%s/%s?%s",
		base,
		path,
		params.Encode())

	res, err := c.Get(uri)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(result)
	if err != nil {
		return err
	}

	return nil
}

func PostRequest(c *http.Client, base string, path string, jsonBody string, result interface{}) error {
	// An example of the jsonBody is `{"key":"value"}`
	uri := fmt.Sprintf("%s/%s",
		base,
		path)

	b := strings.NewReader(jsonBody)
	res, err := c.Post(uri, "application/json", b)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(result)
	if err != nil {
		return err
	}

	return nil
}
