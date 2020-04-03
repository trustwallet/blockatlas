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
	normalized_addr := Address.EIP55Checksum(address)
	normalized_token := ""
	if token != "" {
		normalized_token = Address.EIP55Checksum(token)
	}
	for _, srcTx := range srcPage.Transactions {
		tx := normalizeTxWithAddress(&srcTx, normalized_addr, normalized_token, coinIndex)
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
		Block:    srcTx.BlockHeight,
		Status:   status,
		Error:    errReason,
		Sequence: srcTx.EthereumSpecific.Nonce,
	}
	FillMeta(&normalized, srcTx, coinIndex)
	return normalized
}

func normalizeTxWithAddress(srcTx *Transaction, address, token string, coinIndex uint) blockatlas.Tx {
	normalized := normalizeTx(srcTx, coinIndex)
	normalized.Direction = GetDirection(address, normalized.From, normalized.To)
	FillMetaWithAddress(&normalized, srcTx, address, token, coinIndex)
	return normalized
}

func FillMeta(final *blockatlas.Tx, tx *Transaction, coinIndex uint) {
	if len(tx.TokenTransfers) == 1 {
		transfer := tx.TokenTransfers[0]
		if transfer.Token == tx.FromAddress() && transfer.Token == tx.ToAddress() {
			final.Meta = blockatlas.TokenTransfer{
				Name:     transfer.Name,
				Symbol:   transfer.Symbol,
				TokenID:  transfer.Token,
				Decimals: transfer.Decimals,
				Value:    blockatlas.Amount(transfer.Value),
				From:     transfer.From,
				To:       transfer.To,
			}
			return
		}
	}
	fillMeta(final, tx, coinIndex)
}

func FillMetaWithAddress(final *blockatlas.Tx, tx *Transaction, address, token string, coinIndex uint) {
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
				return
			}
		}
	}
	fillMeta(final, tx, coinIndex)
}

func fillMeta(final *blockatlas.Tx, tx *Transaction, coinIndex uint) {
	if tx.EthereumSpecific.GasUsed.Uint64() == 21000 {
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
