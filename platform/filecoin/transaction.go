package filecoin

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/platform/filecoin/explorer"
)

const messageMethod = "Send"

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	res, err := p.explorer.GetMessagesByAddress(address, blockatlas.TxPerPage)
	if err != nil {
		return nil, err
	}
	normalized := make([]blockatlas.Tx, 0)
	for _, message := range res.Messages {
		// skip non transfer messages
		if message.Method != messageMethod {
			continue
		}
		normalized = append(normalized, p.NormalizeMessage(message, address))
	}

	return normalized, nil
}

func (p *Platform) NormalizeMessage(message explorer.Message, address string) blockatlas.Tx {
	status := blockatlas.StatusCompleted
	if message.Receipt.ExitCode != 0 {
		status = blockatlas.StatusError
	}
	direction := blockatlas.DirectionOutgoing
	if message.From != address {
		direction = blockatlas.DirectionIncoming
	}
	return blockatlas.Tx{
		ID:        message.Cid,
		Coin:      p.Coin().ID,
		From:      message.From,
		To:        message.To,
		Date:      message.Timestamp,
		Block:     message.Height,
		Status:    status,
		Sequence:  message.Nonce,
		Type:      blockatlas.TxTransfer,
		Direction: direction,
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(message.Value),
			Symbol:   p.Coin().Symbol,
			Decimals: p.Coin().Decimals,
		},
	}
}
