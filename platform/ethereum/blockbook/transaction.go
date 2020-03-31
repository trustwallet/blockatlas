package blockbook

import (
	"strings"

	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (c *Client) GetTransactions(address string, coinIndex uint) (blockatlas.TxPage, error) {
	page, err := c.GetTxs(address)
	if err != nil {
		return nil, err
	}
	return normalizePage(page, address, "", coinIndex), nil
}

func (c *Client) GetTokenTxs(address, token string, coinIndex uint) (blockatlas.TxPage, error) {
	page, err := c.GetTxsWithContract(address, token)
	if err != nil {
		return nil, err
	}
	return normalizePage(page, address, token, coinIndex), nil
}

func normalizePage(srcPage *Page, address, token string, coinIndex uint) blockatlas.TxPage {
	var txs []blockatlas.Tx
	for _, srcTx := range srcPage.Transactions {
		tx := normalizeTxWithAddress(&srcTx, address, token, coinIndex)
		txs = append(txs, tx)
	}
	return blockatlas.TxPage(txs)
}

func normalizeTx(srcTx *Transaction, coinIndex uint) blockatlas.Tx {
	status, errReason := getStatus(srcTx.EthereumSpecific)
	from := getFrom(srcTx)
	to := getTo(srcTx)

	return blockatlas.Tx{
		ID:       srcTx.TxID,
		Coin:     coinIndex,
		From:     from,
		To:       to,
		Fee:      blockatlas.Amount(srcTx.Fees),
		Date:     srcTx.BlockTime,
		Block:    srcTx.BlockHeight,
		Status:   status,
		Error:    errReason,
		Sequence: srcTx.EthereumSpecific.Nonce,
	}
}

func normalizeTxWithAddress(srcTx *Transaction, address, token string, coinIndex uint) blockatlas.Tx {
	normalized := normalizeTx(srcTx, coinIndex)
	normalized.Direction = getDirection(address, normalized.From, normalized.To)
	fillMeta(&normalized, srcTx, address, token, coinIndex)
	return normalized
}

func getStatus(specific *EthereumSpecific) (blockatlas.Status, string) {
	switch specific.Status {
	case -1:
		return blockatlas.StatusPending, ""
	case 0:
		return blockatlas.StatusError, "Error"
	case 1:
		return blockatlas.StatusCompleted, ""
	default:
		return blockatlas.StatusError, "Unable to define transaction status"
	}
}

func getFrom(srcTx *Transaction) string {
	if len(srcTx.Vin) > 0 {
		return srcTx.Vin[0].Addresses[0]
	}
	return ""
}

func getTo(srcTx *Transaction) string {
	if len(srcTx.Vout) > 0 {
		return srcTx.Vout[0].Addresses[0]
	}
	return ""
}

func getDirection(address, from, to string) blockatlas.Direction {
	if address == from && address == to {
		return blockatlas.DirectionSelf
	}
	if address == from {
		return blockatlas.DirectionOutgoing
	} else {
		return blockatlas.DirectionIncoming
	}
}

func fillMeta(final *blockatlas.Tx, tx *Transaction, address, token string, coinIndex uint) {
	if len(tx.TokenTransfers) > 0 {
		for _, transfer := range tx.TokenTransfers {
			if transfer.To == address || transfer.From == address {
				// filter token if specified
				if token != "" {
					if strings.ToLower(token) != strings.ToLower(transfer.Token) {
						continue
					}
				}
				direction := getDirection(address, transfer.From, transfer.To)
				metadata := blockatlas.TokenTransfer{
					Name:     transfer.Name,
					Symbol:   transfer.Symbol,
					TokenID:  transfer.Token,
					Decimals: transfer.Decimals,
					Value:    blockatlas.Amount(transfer.Value),
				}
				if direction == blockatlas.DirectionSelf {
					metadata.From = address
					metadata.To = address
				} else if direction == blockatlas.DirectionOutgoing {
					metadata.From = address
					metadata.To = transfer.To
				} else {
					metadata.From = transfer.From
					metadata.To = address
				}
				final.Direction = direction
				final.Meta = metadata
				return
			}
		}
	} else {
		if tx.EthereumSpecific.GasLimit.Uint64() == 21000 {
			final.Meta = blockatlas.Transfer{
				Value:    blockatlas.Amount(tx.Value),
				Symbol:   coin.Coins[coinIndex].Symbol,
				Decimals: coin.Coins[coinIndex].Decimals,
			}
			return
		}
	}
	final.Meta = blockatlas.ContractCall{
		Input: "0x", // FIXME blockbook api doesn't return tx data field
		Value: tx.Value,
	}
}
