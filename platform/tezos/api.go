package tezos

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"time"
)

type Platform struct {
	client Client
}

const Annual = 7.0

func (p *Platform) Init() error {
	p.client = InitClient(viper.GetString("tezos.api"), viper.GetString("tezos.rpc"))
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.XTZ]
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	s, err := p.client.GetTxsOfAddress(address)
	if err != nil {
		return nil, err
	}

	txs := NormalizeTxs(s)

	return txs, nil
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetCurrentBlock()
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	if srcBlock, err := p.client.GetBlockByNumber(num); err == nil {
		txs := NormalizeTxs(srcBlock)
		return &blockatlas.Block{
			Number: num,
			Txs:    txs,
		}, nil
	} else {
		return nil, err
	}
}

func NormalizeTxs(srcTxs []Tx) (txs []blockatlas.Tx) {
	for _, srcTx := range srcTxs {
		tx, ok := Normalize(&srcTx)
		if !ok || len(txs) >= blockatlas.TxPerPage {
			continue
		}
		txs = append(txs, tx)
	}
	return txs
}

func (p *Platform) GetValidators() (blockatlas.ValidatorPage, error) {
	results := make(blockatlas.ValidatorPage, 0)
	validators, err := p.client.GetValidators()

	if err != nil {
		return results, err
	}

	for _, v := range validators {
		results = append(results, normalizeValidator(v))
	}

	return results, nil
}

func normalizeValidator(v Validator) (validator blockatlas.Validator) {
	// How to calculate Tezos APR? I have no idea. Tezos team does not know either. let's assume it's around 7% - no way to calculate in decentralized manner
	// Delegation rewards distributed by the validators manually, it's up to them to do it.

	return blockatlas.Validator{
		Status: true,
		ID:     v.Address,
		Reward: blockatlas.StakingReward{Annual: Annual},
	}
}

// Normalize converts a Tezos transaction into the generic model
func Normalize(srcTx *Tx) (tx blockatlas.Tx, ok bool) {
	if srcTx.Type.Kind != "manager" {
		return tx, false
	}
	if len(srcTx.Type.Operations) < 1 {
		return tx, false
	}

	op := srcTx.Type.Operations[0]

	date, err := time.Parse("2006-01-02T15:04:05Z", op.Timestamp)
	var unix int64
	if err != nil {
		unix = 0
	} else {
		unix = date.Unix()
	}

	if op.Kind != "transaction" {
		return tx, false
	}
	var status blockatlas.Status
	var errMsg string
	if !op.Failed {
		status = blockatlas.StatusCompleted
	} else {
		status = blockatlas.StatusFailed
		errMsg = "transaction failed"
	}
	return blockatlas.Tx{
		ID:    srcTx.Hash,
		Coin:  coin.XTZ,
		Date:  unix,
		From:  op.Src.Tz,
		To:    op.Dest.Tz,
		Fee:   op.Fee,
		Block: op.OpLevel,
		Meta: blockatlas.Transfer{
			Value:    op.Amount,
			Symbol:   coin.Coins[coin.XTZ].Symbol,
			Decimals: coin.Coins[coin.XTZ].Decimals,
		},
		Status: status,
		Error:  errMsg,
	}, true
}
