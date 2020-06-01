package tron

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/address"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
)

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	Txs, err := p.client.GetTxsOfAddress(address, "")
	if err != nil && len(Txs) == 0 {
		return nil, err
	}

	txs := make(blockatlas.TxPage, 0)
	for _, srcTx := range Txs {
		tx, err := Normalize(srcTx)
		if err != nil {
			continue
		}

		if len(srcTx.RawData.Contracts) > 0 && srcTx.RawData.Contracts[0].Type == TransferContract {
			txs = append(txs, *tx)
		} else {
			continue
		}
	}

	return txs, nil
}

func (p *Platform) GetTokenTxsByAddress(address, token string) (blockatlas.TxPage, error) {
	tokenTxs, err := p.client.GetTxsOfAddress(address, token)
	if err != nil {
		return nil, errors.E(err, "TRON: failed to get token from address", errors.TypePlatformApi,
			errors.Params{"address": address, "token": token})
	}

	txs := make(blockatlas.TxPage, 0)

	if len(tokenTxs) == 0 {
		return txs, nil
	}

	info, err := p.client.GetTokenInfo(token)
	if err != nil || len(info.Data) == 0 {
		return nil, errors.E(err, "TRON: failed to get token info", errors.TypePlatformApi,
			errors.Params{"address": address, "token": token})
	}

	for _, srcTx := range tokenTxs {
		tx, err := Normalize(srcTx)
		if err != nil {
			logger.Error(err)
			continue
		}
		setTokenMeta(tx, srcTx, info.Data[0])
		txs = append(txs, *tx)
	}

	return txs, nil
}

func setTokenMeta(tx *blockatlas.Tx, srcTx Tx, tokenInfo AssetInfo) {
	transfer := srcTx.RawData.Contracts[0].Parameter.Value
	tx.Meta = blockatlas.TokenTransfer{
		Name:     tokenInfo.Name,
		Symbol:   tokenInfo.Symbol,
		TokenID:  tokenInfo.ID,
		Decimals: tokenInfo.Decimals,
		Value:    transfer.Amount,
		From:     tx.From,
		To:       tx.To,
	}
}

/// Normalize converts a Tron transaction into the generic model
func Normalize(srcTx Tx) (*blockatlas.Tx, error) {
	if len(srcTx.RawData.Contracts) == 0 {
		return nil, errors.E("TRON: transfer without contract", errors.TypePlatformApi,
			errors.Params{"tx": srcTx})
	}

	contract := srcTx.RawData.Contracts[0]
	if contract.Type != TransferContract && contract.Type != TransferAssetContract {
		return nil, errors.E("TRON: invalid contract transfer", errors.TypePlatformApi,
			errors.Params{"tx": srcTx, "type": contract.Type})
	}

	transfer := contract.Parameter.Value
	from, err := address.HexToBase58(transfer.OwnerAddress)
	if err != nil {
		return nil, errors.E(err, "TRON: failed to get from address", errors.TypePlatformApi,
			errors.Params{"tx": srcTx})
	}
	to, err := address.HexToBase58(transfer.ToAddress)
	if err != nil {
		return nil, errors.E(err, "TRON: failed to get to address", errors.TypePlatformApi,
			errors.Params{"tx": srcTx})
	}

	return &blockatlas.Tx{
		ID:     srcTx.ID,
		Coin:   coin.TRX,
		Date:   srcTx.BlockTime / 1000,
		From:   from,
		To:     to,
		Fee:    "0", // TODO get fee
		Block:  0,   // TODO get block
		Status: blockatlas.StatusCompleted, // TODO determine status
		Meta: blockatlas.Transfer{
			Value:    transfer.Amount,
			Symbol:   coin.Coins[coin.TRX].Symbol,
			Decimals: coin.Coins[coin.TRX].Decimals,
		},
	}, nil
}
