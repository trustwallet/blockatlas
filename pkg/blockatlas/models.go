package blockatlas

type (
	DocsResponse struct {
		Docs interface{} `json:"docs"`
	}

	ResultsResponse struct {
		Total   int         `json:"total"`
		Results interface{} `json:"docs"`
	}
)
