package filecoin

type ChainHeadResponse struct {
	Cids []struct {
		Cid string `json:"/"`
	} `json:"Cids"`
	Height int `json:"Height"`
}
