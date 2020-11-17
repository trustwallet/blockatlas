package filecoin

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/coin"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	response, err := p.client.getBlockHeight()
	if err != nil {
		return 0, err
	}
	return int64(response.Height), nil
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	chainHeadResponse, err := p.client.getTipSetByHeight(num)
	if err != nil {
		return nil, err
	}
	blockResponses := make([]BlockMessageResponse, 0, len(chainHeadResponse.getCids()))
	for _, cid := range chainHeadResponse.getCids() {
		blockResponse, err := p.client.getBlockMessage(cid)
		if err != nil {
			return nil, err
		}
		blockResponses = append(blockResponses, blockResponse)
	}

	return normalizeBlockResponses(uint64(chainHeadResponse.Height), blockResponses), nil
}

func normalizeBlockResponses(num uint64, responses []BlockMessageResponse) *blockatlas.Block {
	var result blockatlas.Block
	result.Number = int64(num)
	for _, resp := range responses {
		for _, msg := range resp.SecpkMessages {
			tx := normalizeBlockTx(num, msg)
			result.Txs = append(result.Txs, tx)
		}
	}
	return &result
}

func normalizeBlockTx(num uint64, msg SecpkMessage) blockatlas.Tx {
	return blockatlas.Tx{
		Coin:     coin.Filecoin().ID,
		From:     msg.Message.From,
		To:       msg.Message.To,
		Fee:      "0",
		Block:    num,
		Status:   blockatlas.StatusCompleted,
		Sequence: uint64(msg.Message.Nonce),
		Type:     blockatlas.TxTransfer,
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(msg.Message.Value),
			Symbol:   coin.Filecoin().Symbol,
			Decimals: coin.Filecoin().Decimals,
		},
	}
}
