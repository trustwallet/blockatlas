package cosmos

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	services "github.com/trustwallet/blockatlas/services/assets"
	"github.com/trustwallet/blockatlas/util"
	"sort"
	"strconv"
	"time"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client = Client{blockatlas.InitClient(viper.GetString("cosmos.api"))}
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.ATOM]
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	srcTxs, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}

	txs := NormalizeTxs(srcTxs, len(srcTxs))
	return &blockatlas.Block{
		Number: num,
		Txs:    txs,
	}, nil
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.CurrentBlockNumber()
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	srcTxes := make([]Tx, 0)

	tagsList := []string{"recipient", "sender", "delegator", "destination-validator"}

	for _, tag := range tagsList {
		responseTxes, _ := p.client.GetAddrTxes(address, tag)
		srcTxes = append(srcTxes, responseTxes...)
	}

	normalisedTxs := make([]blockatlas.Tx, 0)

	for _, srcTx := range srcTxes {
		normalisedInputTx, ok := Normalize(&srcTx)
		if ok {
			normalisedTxs = append(normalisedTxs, normalisedInputTx)
		}
	}

	sort.Slice(normalisedTxs, func(i, j int) bool {
		return normalisedTxs[i].Date > normalisedTxs[j].Date
	})

	return normalisedTxs, nil
}

func (p *Platform) GetValidators() (blockatlas.ValidatorPage, error) {
	results := make(blockatlas.ValidatorPage, 0)
	validators, err := p.client.GetValidators()
	if err != nil {
		return nil, err
	}
	pool, err := p.client.GetPool()
	if err != nil {
		return nil, err
	}

	inflation, err := p.client.GetInflation()
	if err != nil {
		return nil, err
	}

	for _, validator := range validators {
		results = append(results, normalizeValidator(validator, pool, inflation, p.Coin()))
	}

	return results, nil
}

func (p *Platform) GetDetails() blockatlas.StakingDetails {
	//TODO: Find a way to have a dynamic
	return blockatlas.StakingDetails{
		Reward:        blockatlas.StakingReward{Annual: 11},
		MinimumAmount: blockatlas.Amount("0"),
		LockTime:      1814400,
		Type:          blockatlas.DelegationTypeDelegate,
	}
}

func (p *Platform) GetDelegations(address string) (blockatlas.DelegationsPage, error) {
	results := make(blockatlas.DelegationsPage, 0)
	delegations, err := p.client.GetDelegations(address)
	if err != nil {
		return nil, err
	}
	unbondingDelegations, err := p.client.GetUnbondingDelegations(address)
	if err != nil {
		return nil, err
	}
	validators, err := services.GetValidatorsMap(p)
	if err != nil {
		return nil, err
	}
	results = append(results, NormalizeDelegations(delegations, validators)...)
	results = append(results, NormalizeUnbondingDelegations(unbondingDelegations, validators)...)

	return results, nil
}

func (p *Platform) UndelegatedBalance(address string) (string, error) {
	account, err := p.client.GetAccount(address)
	if err != nil {
		return "0", err
	}
	for _, coin := range account.Value.Coins {
		if coin.Denom == "uatom" {
			return coin.Amount, nil
		}
	}
	return "0", nil
}

func NormalizeDelegations(delegations []Delegation, validators blockatlas.ValidatorMap) []blockatlas.Delegation {
	results := make([]blockatlas.Delegation, 0)
	for _, v := range delegations {
		validator, ok := validators[v.ValidatorAddress]
		if !ok {
			logger.Error("Validator not found", validator)
			continue
		}
		delegation := blockatlas.Delegation{
			Delegator: validator,
			Value:     v.Value(),
			Status:    blockatlas.DelegationStatusActive,
		}
		results = append(results, delegation)
	}
	return results
}

func NormalizeUnbondingDelegations(delegations []UnbondingDelegation, validators blockatlas.ValidatorMap) []blockatlas.Delegation {
	results := make([]blockatlas.Delegation, 0)
	for _, v := range delegations {
		for _, entry := range v.Entries {
			validator, ok := validators[v.ValidatorAddress]
			if !ok {
				logger.Error("Validator not found", validator)
				continue
			}
			t, _ := time.Parse(time.RFC3339, entry.CompletionTime)
			delegation := blockatlas.Delegation{
				Delegator: validator,
				Value:     entry.Balance,
				Status:    blockatlas.DelegationStatusPending,
				Metadata: blockatlas.DelegationMetaDataPending{
					AvailableDate: uint(t.Unix()),
				},
			}
			results = append(results, delegation)
		}
	}
	return results
}

// NormalizeTxs converts multiple Cosmos transactions
func NormalizeTxs(srcTxs []Tx, pageSize int) (txs []blockatlas.Tx) {
	for _, srcTx := range srcTxs {
		tx, ok := Normalize(&srcTx)
		if !ok || len(txs) >= pageSize {
			continue
		}
		txs = append(txs, tx)
	}
	return
}

// Normalize converts an Cosmos transaction into the generic model
func Normalize(srcTx *Tx) (tx blockatlas.Tx, ok bool) {
	date, _ := time.Parse("2006-01-02T15:04:05Z", srcTx.Date)
	block, _ := strconv.ParseUint(srcTx.Block, 10, 64)
	// Sometimes fees can be null objects (in the case of no fees e.g. F044F91441C460EDCD90E0063A65356676B7B20684D94C731CF4FAB204035B41)
	var fee string
	if len(srcTx.Data.Contents.Fee.FeeAmount) == 0 {
		fee = "0"
	} else {
		fee, _ = util.DecimalToSatoshis(srcTx.Data.Contents.Fee.FeeAmount[0].Quantity)
	}

	tx = blockatlas.Tx{
		ID:    srcTx.ID,
		Coin:  coin.ATOM,
		Date:  date.Unix(),
		Fee:   blockatlas.Amount(fee),
		Block: block,
		Memo:  srcTx.Data.Contents.Memo,
	}

	if len(srcTx.Data.Contents.Message) > 0 {
		msg := srcTx.Data.Contents.Message[0]
		switch msg.Value.(type) {
		case MessageValueTransfer:
			transfer := msg.Value.(MessageValueTransfer)
			fillTransfer(&tx, transfer)
			return tx, true
		case MessageValueDelegate:
			delegate := msg.Value.(MessageValueDelegate)
			fillDelegate(&tx, delegate, msg.Type)
			return tx, true
		}
	}

	return tx, false
}

func fillTransfer(tx *blockatlas.Tx, transfer MessageValueTransfer) {
	value, _ := util.DecimalToSatoshis(transfer.Amount[0].Quantity)

	tx.From = transfer.FromAddr
	tx.To = transfer.ToAddr

	tx.Meta = blockatlas.Transfer{
		Value:    blockatlas.Amount(value),
		Symbol:   coin.Coins[coin.ATOM].Symbol,
		Decimals: coin.Coins[coin.ATOM].Decimals,
	}
}

func fillDelegate(tx *blockatlas.Tx, delegate MessageValueDelegate, msgType string) {
	value, _ := util.DecimalToSatoshis(delegate.Amount.Quantity)

	tx.From = delegate.DelegatorAddr
	tx.To = delegate.ValidatorAddr

	title := ""
	switch msgType {
	case MsgDelegate:
		title = blockatlas.AnyActionDelegation
	case MsgUndelegate:
		title = blockatlas.AnyActionUndelegation
	}
	tx.Meta = blockatlas.AnyAction{
		Coin:     coin.ATOM,
		Title:    title,
		Key:      blockatlas.KeyStakeDelegate,
		Name:     "ATOM",
		Symbol:   coin.Coins[coin.ATOM].Symbol,
		Decimals: coin.Coins[coin.ATOM].Decimals,
		Value:    blockatlas.Amount(value),
	}
}

func normalizeValidator(v Validator, p StakingPool, inflation float64, c coin.Coin) (validator blockatlas.Validator) {
	reward := CalculateAnnualReward(p, inflation, v)
	return blockatlas.Validator{
		Status: v.Status == 2,
		ID:     v.Address,
		Details: blockatlas.StakingDetails{
			Reward:        blockatlas.StakingReward{Annual: reward},
			MinimumAmount: blockatlas.Amount("0"),
			LockTime:      1814400,
			Type:          blockatlas.DelegationTypeDelegate,
		},
	}
}

func CalculateAnnualReward(p StakingPool, inflation float64, validator Validator) float64 {
	notBondedTokens, err := strconv.ParseFloat(p.NotBondedTokens, 32)
	if err != nil {
		return 0
	}

	bondedTokens, err := strconv.ParseFloat(p.BondedTokens, 32)
	if err != nil {
		return 0
	}

	commission, err := strconv.ParseFloat(validator.Commission.Rate, 32)
	if err != nil {
		return 0
	}
	result := (notBondedTokens + bondedTokens) / bondedTokens * inflation
	return (result - (result * commission)) * 100
}
