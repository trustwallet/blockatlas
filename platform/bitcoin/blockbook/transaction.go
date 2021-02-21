package blockbook

import (
	"strings"

	Address "github.com/trustwallet/golibs/address"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/types"
)

func (c *Client) GetTransactions(address string, coinIndex uint) (types.Txs, error) {
	page, err := c.GetTxs(address)
	if err != nil {
		return nil, err
	}
	return NormalizePage(page, address, "", coinIndex), nil
}

func (c *Client) GetTokenTxs(address, token string, coinIndex uint) (types.Txs, error) {
	page, err := c.GetTxsWithContract(address, token)
	if err != nil {
		return nil, err
	}
	return NormalizePage(page, address, token, coinIndex), nil
}

func NormalizePage(srcPage TransactionsList, address, token string, coinIndex uint) (txs types.Txs) {
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

func normalizeTx(srcTx *Transaction, coinIndex uint) types.Tx {
	status, errReason := srcTx.EthereumSpecific.GetStatus()
	normalized := types.Tx{
		ID:             srcTx.ID,
		Coin:           coinIndex,
		From:           srcTx.FromAddress(),
		To:             srcTx.ToAddress(),
		Fee:            types.Amount(srcTx.GetFee()),
		Date:           srcTx.BlockTime,
		Block:          normalizeBlockHeight(srcTx.BlockHeight),
		Status:         status,
		Error:          errReason,
		Sequence:       srcTx.EthereumSpecific.Nonce,
		TokenTransfers: normalizeTokenTransfers(srcTx.TokenTransfers),
	}
	fillMeta(&normalized, srcTx, coinIndex)
	return normalized
}

func normalizeTxWithAddress(srcTx *Transaction, address, token string, coinIndex uint) types.Tx {
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

func normalizeTokenTransfers(tokenTransfers []TokenTransfer) []types.TokenTransfer {
	result := make([]types.TokenTransfer, 0)
	for _, transfer := range tokenTransfers {
		result = append(result, types.TokenTransfer{
			Name:     transfer.Name,
			Symbol:   transfer.Symbol,
			TokenID:  transfer.Token,
			Decimals: transfer.Decimals,
			Value:    types.Amount(transfer.Value),
			From:     transfer.From,
			To:       transfer.To,
		})
	}
	return result
}

func fillMeta(final *types.Tx, tx *Transaction, coinIndex uint) {
	if ok := fillTokenTransfer(final, tx, coinIndex); !ok {
		fillTransferOrContract(final, tx, coinIndex)
	}
}

func fillMetaWithAddress(final *types.Tx, tx *Transaction, address, token string, coinIndex uint) {
	if ok := fillTokenTransferWithAddress(final, tx, address, token, coinIndex); !ok {
		fillTransferOrContract(final, tx, coinIndex)
	}
}

func fillTokenTransfer(final *types.Tx, tx *Transaction, coinIndex uint) bool {
	if len(tx.TokenTransfers) == 1 {
		transfer := tx.TokenTransfers[0]
		final.Meta = types.TokenTransfer{
			Name:     transfer.Name,
			Symbol:   transfer.Symbol,
			TokenID:  transfer.Token,
			Decimals: transfer.Decimals,
			Value:    types.Amount(transfer.Value),
			From:     transfer.From,
			To:       transfer.To,
		}
		final.TokenTransfers = []types.TokenTransfer{}
		return true
	}
	return false
}

func fillTokenTransferWithAddress(final *types.Tx, tx *Transaction, address, token string, coinIndex uint) bool {
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
			metadata := types.TokenTransfer{
				Name:     transfer.Name,
				Symbol:   transfer.Symbol,
				TokenID:  transfer.Token,
				Decimals: transfer.Decimals,
				Value:    types.Amount(transfer.Value),
			}
			if direction == types.DirectionSelf {
				metadata.From = address
				metadata.To = address
			} else if direction == types.DirectionOutgoing {
				metadata.From = address
				metadata.To = transfer.To
			} else {
				metadata.From = transfer.From
				metadata.To = address
			}
			final.Direction = direction
			final.Meta = metadata
			final.TokenTransfers = []types.TokenTransfer{}
			return true
		}
	}
	return false
}

func fillTransferOrContract(final *types.Tx, tx *Transaction, coinIndex uint) {
	gasUsed := tx.EthereumSpecific.GasUsed
	if gasUsed != nil && gasUsed.Int64() == 21000 {
		final.Meta = types.Transfer{
			Value:    types.Amount(tx.Value),
			Symbol:   coin.Coins[coinIndex].Symbol,
			Decimals: coin.Coins[coinIndex].Decimals,
		}
		return
	}
	data := tx.EthereumSpecific.Data
	if data == "" {
		// old node doesn't have data field
		final.Meta = types.ContractCall{
			Input: "0x",
			Value: tx.Value,
		}
	} else {
		if len(strings.TrimPrefix(data, "0x")) > 0 {
			final.Meta = types.ContractCall{
				Input: data,
				Value: tx.Value,
			}
		} else {
			final.Meta = types.Transfer{
				Value:    types.Amount(tx.Value),
				Symbol:   coin.Coins[coinIndex].Symbol,
				Decimals: coin.Coins[coinIndex].Decimals,
			}
		}
	}
}
