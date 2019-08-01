package blockatlas

import (
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

	err = json.NewDecoder(res.Body).Decode(result)
	if err != nil {
		return err
	}

	return nil
}
