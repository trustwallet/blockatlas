package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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

func RequestPost(c *http.Client, base string, path string, contentType string, body map[string]string, result interface{}) error {

	uri := fmt.Sprintf("%s/%s", base, path)

	requestBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	res, err := c.Post(uri, contentType, bytes.NewBuffer(requestBody))
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
