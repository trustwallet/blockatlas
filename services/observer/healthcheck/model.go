package healthcheck

type Info struct {
	LastParsedBlock int64 `json:"lastParsedBlock"`
	Healthy         bool  `json:"healthy"`
}
