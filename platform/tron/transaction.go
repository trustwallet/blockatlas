package tron

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/address"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"strings"
)

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	txs, err := p.client.getTxsOfAddress(address)
	if err != nil || len(txs) == 0 {
		return nil, err
	}

	txPage := make(blockatlas.TxPage, 0)
	for _, t := range txs {
		if len(t.RawData.Contracts) > 0 && t.RawData.Contracts[0].Type == TransferContract {
			normalizedTx, err := normalizeTransfer(t)
			if err != nil {
				continue
			}
			txPage = append(txPage, *normalizedTx)
		} else {
			continue
		}
	}

	return txPage, nil
}

func (p *Platform) GetTokenTxsByAddress(address, token string) (blockatlas.TxPage, error) {
	tokenType, err := getTokenType(token)
	if err != nil {
		return nil, err
	}

	switch tokenType {
	case blockatlas.TokenTypeTRC10:
		return getTRC10Txs(address, token, p)
	case blockatlas.TokenTypeTRC20:
		return getTRC20Txs(address, token, p)
	default:
		return make(blockatlas.TxPage, 0), nil
	}
}

func getTRC10Txs(address, token string, p *Platform) (blockatlas.TxPage, error) {
	tokenTxs, err := p.client.getTxsOfAddress(address)
	if err != nil {
		return nil, errors.E(err, "TRON: failed to get token from address", errors.TypePlatformApi,
			errors.Params{"address": address, "token": token})
	}

	txs := make(blockatlas.TxPage, 0)
	if len(tokenTxs) == 0 {
		return txs, nil
	}

	trc10Txs := make([]Tx, 0)
	for _, tx := range tokenTxs {
		if len(tx.RawData.Contracts) > 0 && tx.RawData.Contracts[0].Parameter.Value.AssetName == token {
			trc10Txs = append(trc10Txs, tx)
		}
	}

	tInfo, err := p.client.getTokenInfo(token)
	if err != nil || len(tInfo.Data) == 0 {
		return nil, errors.E(err, "TRON: failed to get token info", errors.TypePlatformApi,
			errors.Params{"address": address, "token": token})
	}

	for _, trc10Tx := range trc10Txs {
		normalizedTx, err := normalizeTransfer(trc10Tx)
		if err != nil {
			logger.Error(err)
			continue
		}
		setTokenTransferMeta(normalizedTx, trc10Tx, tInfo.Data[0])
		txs = append(txs, *normalizedTx)
	}

	return txs, nil
}

func getTRC20Txs(address, token string, p *Platform) (blockatlas.TxPage, error) {
	txs, err := p.client.getTRC20TxsOfAddress(address, token, blockatlas.TxPerPage)
	if err != nil {
		return nil, errors.E(err, "TRON: failed to get token from address", errors.TypePlatformApi,
			errors.Params{"address": address, "token": token})
	}

	txPage := make(blockatlas.TxPage, 0)
	for _, t := range txs {
		if t.Type == Transfer {
			normalizedTx, err := normalizeTrc20Transfer(t)
			if err != nil {
				logger.Error(err)
				continue
			}
			txPage = append(txPage, *normalizedTx)
		}
	}

	return txPage, nil
}

func normalizeTrc20Transfer(d D) (*blockatlas.Tx, error) {
	return &blockatlas.Tx{
		ID:     d.TransactionId,
		Coin:   coin.TRX,
		Date:   d.BlockTimestamp / 1000,
		From:   d.From,
		To:     d.TokenInfo.Address,
		Fee:    "0",                        // TODO get fee
		Block:  0,                          // TODO get block
		Status: blockatlas.StatusCompleted, // TODO determine status
		Meta: blockatlas.TokenTransfer{
			Name:     d.TokenInfo.Name,
			Symbol:   d.TokenInfo.Symbol,
			TokenID:  d.TokenInfo.Address,
			Decimals: d.TokenInfo.Decimals,
			Value:    blockatlas.Amount(d.Value),
			From:     d.From,
			To:       d.To,
		},
	}, nil
}

func setTokenTransferMeta(tx *blockatlas.Tx, srcTx Tx, tokenInfo AssetInfo) {
	tx.Meta = blockatlas.TokenTransfer{
		Name:     tokenInfo.Name,
		Symbol:   tokenInfo.Symbol,
		TokenID:  tokenInfo.ID,
		Decimals: tokenInfo.Decimals,
		Value:    srcTx.RawData.Contracts[0].Parameter.Value.Amount,
		From:     tx.From,
		To:       tx.To,
	}
}

/// Normalize converts a Tron transaction into the generic model
func normalizeTransfer(srcTx Tx) (*blockatlas.Tx, error) {
	if len(srcTx.RawData.Contracts) == 0 {
		return nil, errors.E("TRON: transfer without contract", errors.TypePlatformApi, errors.Params{"tx": srcTx})
	}

	contract := srcTx.RawData.Contracts[0]
	if contract.Type != TransferContract && contract.Type != TransferAssetContract {
		return nil, errors.E("TRON: invalid contract transfer", errors.TypePlatformApi, errors.Params{"tx": srcTx, "type": contract.Type})
	}

	transfer := contract.Parameter.Value
	from, err := address.HexToBase58(transfer.OwnerAddress)
	if err != nil {
		return nil, errors.E(err, "TRON: failed to get from address", errors.TypePlatformApi, errors.Params{"tx": srcTx})
	}
	to, err := address.HexToBase58(transfer.ToAddress)
	if err != nil {
		return nil, errors.E(err, "TRON: failed to get to address", errors.TypePlatformApi, errors.Params{"tx": srcTx})
	}

	return &blockatlas.Tx{
		ID:     srcTx.ID,
		Coin:   coin.TRX,
		Date:   srcTx.BlockTime / 1000,
		From:   from,
		To:     to,
		Fee:    "0",                        // TODO get fee
		Block:  0,                          // TODO get block
		Status: blockatlas.StatusCompleted, // TODO determine status
		Meta: blockatlas.Transfer{
			Value:    transfer.Amount,
			Symbol:   coin.Coins[coin.TRX].Symbol,
			Decimals: coin.Coins[coin.TRX].Decimals,
		},
	}, nil
}

func getTokenType(token string) (blockatlas.TokenType, error) {
	if len(token) == 7 {
		return blockatlas.TokenTypeTRC10, nil
	}
	if len(token) == 34 && strings.HasPrefix(token, "T") {
		return blockatlas.TokenTypeTRC20, nil
	}
	return "", errors.E("token not TRC10 or TRC20", errors.Params{"token": token})
}
