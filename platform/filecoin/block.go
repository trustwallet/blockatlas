package filecoin

import (
	"github.com/trustwallet/blockatlas/platform/filecoin/rpc"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/types"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	response, err := p.client.GetBlockHeight()
	if err != nil {
		return 0, err
	}
	return int64(response.Height), nil
}

func (p *Platform) GetBlockByNumber(num int64) (*types.Block, error) {
	chainHeadResponse, err := p.client.GetTipSetByHeight(num)
	if err != nil {
		return nil, err
	}
	blockResponses := make([]rpc.BlockMessageResponse, 0, len(chainHeadResponse.GetCids()))
	for _, cid := range chainHeadResponse.GetCids() {
		blockResponse, err := p.client.GetBlockMessage(cid)
		if err != nil {
			return nil, err
		}
		blockResponses = append(blockResponses, blockResponse)
	}

	return normalizeBlockResponses(uint64(chainHeadResponse.Height), uint64(chainHeadResponse.GetTimestamp()), blockResponses), nil
}

func normalizeBlockResponses(num, timestamp uint64, responses []rpc.BlockMessageResponse) *types.Block {
	var result types.Block
	result.Number = int64(num)
	for _, resp := range responses {
		for _, msg := range resp.SecpkMessages {
			tx := normalizeBlockTx(num, timestamp, msg)
			result.Txs = append(result.Txs, tx)
		}
	}
	return &result
}

func normalizeBlockTx(num, timestamp uint64, msg rpc.SecpkMessage) types.Tx {
	return types.Tx{
		Coin: coin.Filecoin().ID,
		From: msg.Message.From,
		To:   msg.Message.To,
		// todo: use StateGetReceipt + https://documenter.getpostman.com/view/4872192/SWLh5mUd?version=latest
		Fee:      "0",
		Block:    num,
		Date:     int64(timestamp),
		Status:   types.StatusCompleted,
		Sequence: uint64(msg.Message.Nonce),
		Type:     types.TxTransfer,
		Memo:     "",
		Meta: types.Transfer{
			Value:    types.Amount(msg.Message.Value),
			Symbol:   coin.Filecoin().Symbol,
			Decimals: coin.Filecoin().Decimals,
		},
	}
}
