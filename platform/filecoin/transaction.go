package filecoin

import (
	"github.com/trustwallet/blockatlas/platform/filecoin/explorer"
	"github.com/trustwallet/golibs/types"
)

const messageMethod = "Send"

func (p *Platform) GetTxsByAddress(address string) (types.Txs, error) {
	res, err := p.explorer.GetMessagesByAddress(address, types.TxPerPage)
	if err != nil {
		return nil, err
	}
	normalized := make(types.Txs, 0)
	for _, message := range res.Messages {
		// skip non transfer messages
		if message.Method != messageMethod {
			continue
		}
		normalized = append(normalized, p.NormalizeMessage(message, address))
	}

	return normalized, nil
}

func (p *Platform) NormalizeMessage(message explorer.Message, address string) types.Tx {
	status := types.StatusCompleted
	if message.Receipt.ExitCode != 0 {
		status = types.StatusError
	}
	direction := types.DirectionOutgoing
	if message.From != address {
		direction = types.DirectionIncoming
	}
	return types.Tx{
		ID:        message.Cid,
		Coin:      p.Coin().ID,
		From:      message.From,
		To:        message.To,
		Date:      message.Timestamp,
		Block:     message.Height,
		Status:    status,
		Sequence:  message.Nonce,
		Type:      types.TxTransfer,
		Direction: direction,
		Meta: types.Transfer{
			Value:    types.Amount(message.Value),
			Symbol:   p.Coin().Symbol,
			Decimals: p.Coin().Decimals,
		},
	}
}
