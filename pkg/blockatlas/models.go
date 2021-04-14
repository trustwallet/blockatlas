package blockatlas

import "encoding/json"

type (
	ResultsResponse struct {
		Results interface{} `json:"docs"`
	}
)

func MapJsonObject(from interface{}, to interface{}) error {
	bytes, err := json.Marshal(from)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, to)
}
