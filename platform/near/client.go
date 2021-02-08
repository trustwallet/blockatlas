package near

import "github.com/trustwallet/golibs/client"

type Client struct {
	client.Request
}

func (c *Client) GetLasteBlock() (int64, error) {
	var block Block
	err := c.RpcCall(&block, "block", map[string]string{"finality": "final"})
	if err != nil {
		return 0, err
	}
	return int64(block.Header.Height), nil
}

func (c *Client) GetTxsInBlock(num int64) (result ChunkDetail, err error) {
	var block Block
	err = c.RpcCall(&block, "block", map[string]int64{"block_id": num})
	if err != nil || len(block.Chunks) == 0 {
		return
	}

	var chunk ChunkDetail
	err = c.RpcCall(&chunk, "chunk", []string{block.Chunks[0].Hash})
	if err != nil {
		return
	}
	chunk.Header.Timestamp = block.Header.Timestamp
	return chunk, nil
}
