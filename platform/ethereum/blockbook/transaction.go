package blockbook

import (
	"github.com/trustwallet/blockatlas/coin"
	Address "github.com/trustwallet/blockatlas/pkg/address"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (c *Client) GetTransactions(address string, coinIndex uint) (blockatlas.TxPage, error) {
	page, err := c.GetTxs(address)
	if err != nil {
		return nil, err
	}
	return NormalizePage(page, address, "", coinIndex), nil
}

func (c *Client) GetTokenTxs(address, token string, coinIndex uint) (blockatlas.TxPage, error) {
	page, err := c.GetTxsWithContract(address, token)
	if err != nil {
		return nil, err
	}
	return NormalizePage(page, address, token, coinIndex), nil
}

func NormalizePage(srcPage *Page, address, token string, coinIndex uint) blockatlas.TxPage {
	var txs []blockatlas.Tx
	normalizedAddr := Address.EIP55Checksum(address)
	var normalizedToken string
	if token != "" {
		normalizedToken = Address.EIP55Checksum(token)
	}
	for _, srcTx := range srcPage.Transactions {
		if srcTx.BlockHeight < 0 {
			continue // Skip pending tx
		}
		tx := normalizeTxWithAddress(&srcTx, normalizedAddr, normalizedToken, coinIndex)
		txs = append(txs, tx)
	}
	return txs
}

func normalizeTx(srcTx *Transaction, coinIndex uint) blockatlas.Tx {
	status, errReason := srcTx.EthereumSpecific.GetStatus()
	normalized := blockatlas.Tx{
		ID:       srcTx.TxID,
		Coin:     coinIndex,
		From:     srcTx.FromAddress(),
		To:       srcTx.ToAddress(),
		Fee:      blockatlas.Amount(srcTx.Fees),
		Date:     srcTx.BlockTime,
		Block:    normalizeBlockHeight(srcTx.BlockHeight),
		Status:   status,
		Error:    errReason,
		Sequence: srcTx.EthereumSpecific.Nonce,
	}
	fillMeta(&normalized, srcTx, coinIndex)
	return normalized
}

func normalizeTxWithAddress(srcTx *Transaction, address, token string, coinIndex uint) blockatlas.Tx {
	normalized := normalizeTx(srcTx, coinIndex)
	normalized.Direction = GetDirection(address, normalized.From, normalized.To)
	fillMetaWithAddress(&normalized, srcTx, address, token, coinIndex)
	return normalized
}

func normalizeBlockHeight(height int64) uint64 {
	if height < 0 {
		return uint64(0)
	}
	return uint64(height)
}

func fillMeta(final *blockatlas.Tx, tx *Transaction, coinIndex uint) {
	if ok := fillTokenTransfer(final, tx, coinIndex); !ok {
		fillTransferOrContract(final, tx, coinIndex)
	}
}

func fillMetaWithAddress(final *blockatlas.Tx, tx *Transaction, address, token string, coinIndex uint) {
	if ok := fillTokenTransferWithAddress(final, tx, address, token, coinIndex); !ok {
		fillTransferOrContract(final, tx, coinIndex)
	}
}

func fillTokenTransfer(final *blockatlas.Tx, tx *Transaction, coinIndex uint) bool {
	if len(tx.TokenTransfers) == 1 {
		transfer := tx.TokenTransfers[0]
		final.Meta = blockatlas.TokenTransfer{
			Name:     transfer.Name,
			Symbol:   transfer.Symbol,
			TokenID:  transfer.Token,
			Decimals: transfer.Decimals,
			Value:    blockatlas.Amount(transfer.Value),
			From:     transfer.From,
			To:       transfer.To,
		}
		return true
	}
	return false
}

func fillTokenTransferWithAddress(final *blockatlas.Tx, tx *Transaction, address, token string, coinIndex uint) bool {
	if len(tx.TokenTransfers) > 0 {
		for _, transfer := range tx.TokenTransfers {
			if transfer.To == address || transfer.From == address {
				// filter token if specified
				if token != "" {
					if token != transfer.Token {
						continue
					}
				}
				direction := GetDirection(address, transfer.From, transfer.To)
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
				return true
			}
		}
	}
	return false
}

func fillTransferOrContract(final *blockatlas.Tx, tx *Transaction, coinIndex uint) {
	gasUsed := tx.EthereumSpecific.GasUsed
	if gasUsed != nil && gasUsed.Int64() == 21000 {
		final.Meta = blockatlas.Transfer{
			Value:    blockatlas.Amount(tx.Value),
			Symbol:   coin.Coins[coinIndex].Symbol,
			Decimals: coin.Coins[coinIndex].Decimals,
		}
		return
	}
	final.Meta = blockatlas.ContractCall{
		Input: "0x", // FIXME blockbook api doesn't return tx data field
		Value: tx.Value,
	}
}
