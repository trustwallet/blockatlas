package tezos

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	services "github.com/trustwallet/blockatlas/services/assets"
)

type Platform struct {
	client    Client
	rpcClient RpcClient
}

const Annual = 6.09

func (p *Platform) Init() error {
	p.client = Client{blockatlas.InitClient(viper.GetString("tezos.api"))}
	p.client.SetTimeout(25)
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

func (p *Platform) GetDelegations(address string) (blockatlas.DelegationsPage, error) {
	delegations, err := p.client.GetDelegations(address)
	if err != nil {
		return nil, err
	}
	if len(delegations) == 0 {
		return make(blockatlas.DelegationsPage, 0), nil
	}
	validators, err := services.GetValidatorsMap(p)
	if err != nil {
		return nil, err
	}
	delegatedBalance := p.rpcClient.GetDelegatedBalance(address)
	return NormalizeDelegation(delegations[0], delegatedBalance, validators)
}

func NormalizeDelegation(delegation TxDelegation, delegatedBalance string, validators blockatlas.ValidatorMap) (blockatlas.DelegationsPage, error) {
	validator, ok := validators[delegation.Delegation.Delegate]
	if !ok {
		return nil, errors.E("Validator not found",
			errors.Params{"Address": delegation.Delegation.Source, "Delegate": delegation.Delegation.Delegate})
	}
	return blockatlas.DelegationsPage{
		{
			Delegator: validator,
			Value:     delegatedBalance,
			Status:    blockatlas.DelegationStatusActive,
		},
	}, nil
}

func NormalizeTxs(srcTxs []Transaction) (txs []blockatlas.Tx) {
	for _, srcTx := range srcTxs {
		tx, ok := NormalizeTx(srcTx)
		if !ok {
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

func (p *Platform) UndelegatedBalance(address string) (string, error) {
	return p.rpcClient.GetBalance(address), nil
}

func getDetails() blockatlas.StakingDetails {
	return blockatlas.StakingDetails{
		Reward:        blockatlas.StakingReward{Annual: Annual},
		MinimumAmount: "0",
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
func NormalizeTx(srcTx Transaction) (tx blockatlas.Tx, ok bool) {
	if srcTx.Tx.Kind != "transaction" {
		return tx, false
	}

	var status blockatlas.Status
	var errMsg string
	if srcTx.Tx.Status == "applied" {
		status = blockatlas.StatusCompleted
	} else {
		status = blockatlas.StatusFailed
		errMsg = "transaction failed"
	}
	return blockatlas.Tx{
		ID:    srcTx.Op.OpHash,
		Coin:  coin.XTZ,
		Date:  srcTx.Op.BlockTimestamp.Unix(),
		From:  srcTx.Tx.Source,
		To:    srcTx.Tx.Destination,
		Fee:   blockatlas.Amount(srcTx.Tx.Fee),
		Block: srcTx.Op.BlockLevel,
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(srcTx.Tx.Amount),
			Symbol:   coin.Coins[coin.XTZ].Symbol,
			Decimals: coin.Coins[coin.XTZ].Decimals,
		},
		Status: status,
		Error:  errMsg,
	}, true
}
