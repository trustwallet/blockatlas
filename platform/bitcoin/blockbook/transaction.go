package blockbook

import (
	"strings"

	Address "github.com/trustwallet/golibs/address"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/txtype"
)

func (c *Client) GetTransactions(address string, coinIndex uint) (txtype.TxPage, error) {
	page, err := c.GetTxs(address)
	if err != nil {
		return nil, err
	}
	return NormalizePage(page, address, "", coinIndex), nil
}

func (c *Client) GetTokenTxs(address, token string, coinIndex uint) (txtype.TxPage, error) {
	page, err := c.GetTxsWithContract(address, token)
	if err != nil {
		return nil, err
	}
	return NormalizePage(page, address, token, coinIndex), nil
}

func NormalizePage(srcPage TransactionsList, address, token string, coinIndex uint) (txs txtype.TxPage) {
	normalizedAddr, err := Address.EIP55Checksum(address)
	if err != nil {
		return
	}
	var normalizedToken string
	if token != "" {
		normalizedToken, err = Address.EIP55Checksum(token)
		if err != nil {
			return
		}
	}
	for _, srcTx := range srcPage.Transactions {
		tx := normalizeTxWithAddress(&srcTx, normalizedAddr, normalizedToken, coinIndex)
		txs = append(txs, tx)
	}
	return txs
}

func normalizeTx(srcTx *Transaction, coinIndex uint) txtype.Tx {
	status, errReason := srcTx.EthereumSpecific.GetStatus()
	normalized := txtype.Tx{
		ID:       srcTx.ID,
		Coin:     coinIndex,
		From:     srcTx.FromAddress(),
		To:       srcTx.ToAddress(),
		Fee:      txtype.Amount(srcTx.GetFee()),
		Date:     srcTx.BlockTime,
		Block:    normalizeBlockHeight(srcTx.BlockHeight),
		Status:   status,
		Error:    errReason,
		Sequence: srcTx.EthereumSpecific.Nonce,
	}
	fillMeta(&normalized, srcTx, coinIndex)
	return normalized
}

func normalizeTxWithAddress(srcTx *Transaction, address, token string, coinIndex uint) txtype.Tx {
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

func fillMeta(final *txtype.Tx, tx *Transaction, coinIndex uint) {
	if ok := fillTokenTransfer(final, tx, coinIndex); !ok {
		fillTransferOrContract(final, tx, coinIndex)
	}
}

func fillMetaWithAddress(final *txtype.Tx, tx *Transaction, address, token string, coinIndex uint) {
	if ok := fillTokenTransferWithAddress(final, tx, address, token, coinIndex); !ok {
		fillTransferOrContract(final, tx, coinIndex)
	}
}

func fillTokenTransfer(final *txtype.Tx, tx *Transaction, coinIndex uint) bool {
	if len(tx.TokenTransfers) == 1 {
		transfer := tx.TokenTransfers[0]
		final.Meta = txtype.TokenTransfer{
			Name:     transfer.Name,
			Symbol:   transfer.Symbol,
			TokenID:  transfer.Token,
			Decimals: transfer.Decimals,
			Value:    txtype.Amount(transfer.Value),
			From:     transfer.From,
			To:       transfer.To,
		}
		return true
	}
	return false
}

func fillTokenTransferWithAddress(final *txtype.Tx, tx *Transaction, address, token string, coinIndex uint) bool {
	if len(tx.TokenTransfers) == 1 {
		transfer := tx.TokenTransfers[0]
		if transfer.To == address || transfer.From == address {
			// filter token if specified
			if token != "" {
				if token != transfer.Token {
					return false
				}
			}
			direction := GetDirection(address, transfer.From, transfer.To)
			metadata := txtype.TokenTransfer{
				Name:     transfer.Name,
				Symbol:   transfer.Symbol,
				TokenID:  transfer.Token,
				Decimals: transfer.Decimals,
				Value:    txtype.Amount(transfer.Value),
			}
			if direction == txtype.DirectionSelf {
				metadata.From = address
				metadata.To = address
			} else if direction == txtype.DirectionOutgoing {
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
	return false
}

func fillTransferOrContract(final *txtype.Tx, tx *Transaction, coinIndex uint) {
	gasUsed := tx.EthereumSpecific.GasUsed
	if gasUsed != nil && gasUsed.Int64() == 21000 {
		final.Meta = txtype.Transfer{
			Value:    txtype.Amount(tx.Value),
			Symbol:   coin.Coins[coinIndex].Symbol,
			Decimals: coin.Coins[coinIndex].Decimals,
		}
		return
	}
	data := tx.EthereumSpecific.Data
	if data == "" {
		// old node doesn't have data field
		final.Meta = txtype.ContractCall{
			Input: "0x",
			Value: tx.Value,
		}
	} else {
		if len(strings.TrimPrefix(data, "0x")) > 0 {
			final.Meta = txtype.ContractCall{
				Input: data,
				Value: tx.Value,
			}
		} else {
			final.Meta = txtype.Transfer{
				Value:    txtype.Amount(tx.Value),
				Symbol:   coin.Coins[coinIndex].Symbol,
				Decimals: coin.Coins[coinIndex].Decimals,
			}
		}
	}
}
