package solana

import (
	"bytes"
	"encoding/binary"
	"github.com/btcsuite/btcutil/base58"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	services "github.com/trustwallet/blockatlas/services/assets"
	"strconv"
)

func arrayOfPubkey(pubkey string) [32]byte {
	var array [32]byte
	copy(array[:], base58.Decode(pubkey))
	return array
}

func (p *Platform) GetValidators() (blockatlas.ValidatorPage, error) {
	results := make(blockatlas.ValidatorPage, 0)

	validators, err := p.client.GetCurrentVoteAccounts()
	if err != nil {
		return results, err
	}

	minimumBalance, err := p.client.GetMinimumBalanceForRentExemption()
	if err != nil {
		minimumBalance = 0
	}

	for _, v := range validators {
		results = append(results, normalizeValidator(v, minimumBalance))
	}
	return results, nil
}

func (p *Platform) GetDetails() blockatlas.StakingDetails {
	return getDetails()
}

func getDetails() blockatlas.StakingDetails {
	return blockatlas.StakingDetails{
		Reward:        blockatlas.StakingReward{Annual: 0},
		MinimumAmount: "0",
		LockTime:      0,
		Type:          blockatlas.DelegationTypeDelegate,
	}
}

func normalizeValidator(v VoteAccount, minimumBalance uint64) (validator blockatlas.Validator) {
	minimumAmount := strconv.FormatUint(minimumBalance+1, 10) // Must stake at least 1 lamport
	return blockatlas.Validator{
		Status: true,
		ID:     v.VotePubkey,
		Details: blockatlas.StakingDetails{
			Reward:        blockatlas.StakingReward{Annual: 0},
			MinimumAmount: blockatlas.Amount(minimumAmount),
			LockTime:      0,
			Type:          blockatlas.DelegationTypeDelegate,
		},
	}
}

func (p *Platform) GetDelegations(address string) (blockatlas.DelegationsPage, error) {
	accounts, err := p.client.GetStakeAccounts()
	if err != nil {
		return nil, err
	}

	stakeAccounts := make([]StakeState, 0)
	for _, keyedAccount := range accounts {
		account, err := parseStakeData(keyedAccount.Account)
		if err != nil {
			return nil, err
		}
		if isAuthorized(account, address) {
			stakeAccounts = append(stakeAccounts, account)
		}
	}
	if len(stakeAccounts) == 0 {
		return make(blockatlas.DelegationsPage, 0), nil
	}

	validators, err := services.GetValidatorsMap(p)
	if err != nil {
		return nil, err
	}

	epochInfo, err := p.client.GetEpochInfo()
	if err != nil {
		return nil, err
	}

	return NormalizeDelegations(stakeAccounts, validators, epochInfo)
}

func parseStakeData(account Account) (stakeAccount StakeState, err error) {
	buffer := base58.Decode(account.Data)
	r := bytes.NewReader(buffer)
	err = binary.Read(r, binary.LittleEndian, &stakeAccount)
	return
}

func isAuthorized(stakeAccount StakeState, address string) bool {
	return stakeAccount.AuthorizedStaker == arrayOfPubkey(address)
}

func (p *Platform) UndelegatedBalance(address string) (string, error) {
	account, err := p.client.GetAccount(address)
	if err != nil {
		return "0", err
	}
	return strconv.FormatUint(account.Lamports, 10), nil
}

func NormalizeDelegations(stakeAccounts []StakeState, validators blockatlas.ValidatorMap, epochInfo EpochInfo) (blockatlas.DelegationsPage, error) {
	results := make([]blockatlas.Delegation, 0)
	for _, stakeState := range stakeAccounts {
		votePubkey := base58.Encode(stakeState.VoterPubkey[:])
		validator, ok := validators[votePubkey]
		if !ok {
			return nil, errors.E("Validator not found",
				errors.Params{"Validator": votePubkey, "Platform": "solana"})
		}
		status := blockatlas.DelegationStatusPending
		if stakeState.ActivationEpoch > 0 && stakeState.ActivationEpoch <= epochInfo.Epoch {
			status = blockatlas.DelegationStatusActive
		}
		delegation := blockatlas.Delegation{
			Delegator: validator,
			Value:     strconv.FormatUint(stakeState.Stake, 10),
			Status:    status,
		}
		results = append(results, delegation)
	}
	return results, nil
}
