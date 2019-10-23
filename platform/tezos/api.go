package tezos

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"math"
	"strconv"
	"time"
)

type Platform struct {
	client    Client
	rpcClient RpcClient
}

const Annual = 7.0

func (p *Platform) Init() error {
	p.client = Client{blockatlas.InitClient(viper.GetString("tezos.api"))}
	p.rpcClient = RpcClient{blockatlas.InitClient(viper.GetString("tezos.rpc"))}
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
	srcBlock, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}
	txs := NormalizeTxs(srcBlock)
	return &blockatlas.Block{
		Number: num,
		Txs:    txs,
	}, nil
}

func (p *Platform) GetDelegations(address string) (page blockatlas.DelegationsPage, err error) {
	//TODO https://github.com/trustwallet/blockatlas/issues/386
	return page, err
}

func NormalizeTxs(srcTxs []Tx) (txs []blockatlas.Tx) {
	for _, srcTx := range srcTxs {
		tx, ok := NormalizeTx(&srcTx)
		if !ok || len(txs) >= blockatlas.TxPerPage {
			continue
		}
		txs = append(txs, tx)
	}
	return txs
}

func (p *Platform) GetValidators() (blockatlas.ValidatorPage, error) {
	results := make(blockatlas.ValidatorPage, 0)
	validators, err := p.rpcClient.GetValidators()

	if err != nil {
		return results, err
	}

	for _, v := range validators {
		results = append(results, normalizeValidator(v))
	}

	return results, nil
}

func (p *Platform) GetDetails() blockatlas.StakingDetails {
	return getDetails()
}

func (p *Platform) GetBalance(address string) (string, error) {
	//TODO: Implement fetching balance for Tezos
	return "0", nil
}

func getDetails() blockatlas.StakingDetails {
	return blockatlas.StakingDetails{
		Reward:        blockatlas.StakingReward{Annual: Annual},
		MinimumAmount: blockatlas.Amount("0"),
		LockTime:      0,
		Type:          blockatlas.DelegationTypeDelegate,
	}
}

func normalizeValidator(v Validator) (validator blockatlas.Validator) {
	// How to calculate Tezos APR? I have no idea. Tezos team does not know either. let's assume it's around 7% - no way to calculate in decentralized manner
	// Delegation rewards distributed by the validators manually, it's up to them to do it.

	return blockatlas.Validator{
		Status:  true,
		ID:      v.Address,
		Details: getDetails(),
	}
}

// NormalizeTx converts a Tezos transaction into the generic model
func NormalizeTx(srcTx *Tx) (tx blockatlas.Tx, ok bool) {
	unix := int64(0)
	date, err := time.Parse("2006-01-02T15:04:05Z", srcTx.Time)
	if err == nil {
		unix = date.Unix()
	}

	if srcTx.Type != "transaction" {
		return tx, false
	}

	var status blockatlas.Status
	var errMsg string
	if srcTx.Success && srcTx.Status == "applied" {
		status = blockatlas.StatusCompleted
	} else {
		status = blockatlas.StatusFailed
		errMsg = "transaction failed"
	}

	decimals := coin.Coins[coin.XTZ].Decimals
	d := math.Pow10(int(decimals))
	v := srcTx.Volume * d
	volume := strconv.Itoa(int(v))
	return blockatlas.Tx{
		ID:    srcTx.Hash,
		Coin:  coin.XTZ,
		Date:  unix,
		From:  srcTx.Sender,
		To:    srcTx.Receiver,
		Fee:   blockatlas.Amount(strconv.Itoa(srcTx.Fee)),
		Block: srcTx.Height,
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(volume),
			Symbol:   coin.Coins[coin.XTZ].Symbol,
			Decimals: coin.Coins[coin.XTZ].Decimals,
		},
		Status: status,
		Error:  errMsg,
	}, true
}
