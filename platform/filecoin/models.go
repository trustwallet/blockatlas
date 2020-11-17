package filecoin

type ChainHeadResponse struct {
	Cids []struct {
		Cid string `json:"/"`
	} `json:"Cids"`
	Blocks []struct {
		Timestamp int `json:"Timestamp"`
	}
	Height int `json:"Height"`
}

type BlockMessageResponse struct {
	SecpkMessages []SecpkMessage `json:"SecpkMessages"`
}

type SecpkMessage struct {
	Message struct {
		Version    int         `json:"Version"`
		To         string      `json:"To"`
		From       string      `json:"From"`
		Nonce      int         `json:"Nonce"`
		Value      string      `json:"Value"`
		GasLimit   int         `json:"GasLimit"`
		GasFeeCap  string      `json:"GasFeeCap"`
		GasPremium string      `json:"GasPremium"`
		Method     int         `json:"Method"`
		Params     interface{} `json:"Params"`
	} `json:"Message"`
}

func (c ChainHeadResponse) getCids() []string {
	result := make([]string, 0, len(c.Cids))
	for _, cid := range c.Cids {
		result = append(result, cid.Cid)
	}
	return result
}

func (c ChainHeadResponse) getTimestamp() int64 {
	if len(c.Blocks) == 0 {
		return 0
	}
	return int64(c.Blocks[0].Timestamp)
}
