package filecoin

import (
	"github.com/trustwallet/blockatlas/platform/filecoin/explorer"
	"github.com/trustwallet/golibs/txtype"
)

const messageMethod = "Send"

func (p *Platform) GetTxsByAddress(address string) (txtype.TxPage, error) {
	res, err := p.explorer.GetMessagesByAddress(address, txtype.TxPerPage)
	if err != nil {
		return nil, err
	}
	normalized := make([]txtype.Tx, 0)
	for _, message := range res.Messages {
		// skip non transfer messages
		if message.Method != messageMethod {
			continue
		}
		normalized = append(normalized, p.NormalizeMessage(message, address))
	}

	return normalized, nil
}

func (p *Platform) NormalizeMessage(message explorer.Message, address string) txtype.Tx {
	status := txtype.StatusCompleted
	if message.Receipt.ExitCode != 0 {
		status = txtype.StatusError
	}
	direction := txtype.DirectionOutgoing
	if message.From != address {
		direction = txtype.DirectionIncoming
	}
	return txtype.Tx{
		ID:        message.Cid,
		Coin:      p.Coin().ID,
		From:      message.From,
		To:        message.To,
		Date:      message.Timestamp,
		Block:     message.Height,
		Status:    status,
		Sequence:  message.Nonce,
		Type:      txtype.TxTransfer,
		Direction: direction,
		Meta: txtype.Transfer{
			Value:    txtype.Amount(message.Value),
			Symbol:   p.Coin().Symbol,
			Decimals: p.Coin().Decimals,
		},
	}
}
