package solana

import (
	"encoding/json"
	"testing"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/mock"
	"github.com/trustwallet/golibs/txtype"

	"github.com/stretchr/testify/assert"
)

var (
	currentValidators, _ = mock.JsonStringFromFilePath("mocks/" + "currentValidators.json")
	expectedValidators   = []blockatlas.Validator{
		{
			Status: true,
			ID:     "2Afu38M1KaSfDBpjZjnJb9BSWP6YkBkoPiBfnFedD7JW",
			Details: blockatlas.StakingDetails{
				Reward:        blockatlas.StakingReward{Annual: 0},
				MinimumAmount: txtype.Amount("2282881"),
				LockTime:      0,
				Type:          blockatlas.DelegationTypeDelegate,
			},
		},
		{
			Status: true,
			ID:     "5CgQubGD1uwodwCe5UXDADbC69SiqXR8qq6pDMSm7ut5",
			Details: blockatlas.StakingDetails{
				Reward:        blockatlas.StakingReward{Annual: 0},
				MinimumAmount: txtype.Amount("2282881"),
				LockTime:      0,
				Type:          blockatlas.DelegationTypeDelegate,
			},
		},
	}

	keyedStakeAccount = KeyedAccount{
		Account: Account{
			Data:       "mrWMHx6j3BkmepJy67XLycC7LVeiq2NBESfV2YNmZvY62xT5jTgKnMzRBQheYtAuajncAniTEmU8QxgkpnytnXynTrMSJN4p6ihefU5cobkyCeSwMKugKGuBbyDLjQoUMu6BUKjDTFvjpJUHFCgz1Vaa8HSVscUqqRcioByf3owMUwHmEYsF8vuouLAqEQmo61wFkKfZELxLrhBbi2PQQZucryrnNDKXV4DY3oegLy9aMnMDZUeoDtDPPiJeM2F1Trh8ZkH1sQL6sQ5V",
			Executable: false,
			Lamports:   100000000000,
			Owner:      "Stake11111111111111111111111111111111111111",
			RentEpoch:  80,
		},
		Pubkey: "EgR17fgGmwQQaMZPsuJdk9oHw2xY8TJQj3Bp44o24mar",
	}

	stakeState = StakeData{
		State:                2,
		RentExemptReserve:    2282880,
		AuthorizedStaker:     arrayOfPubkey("B52Da5MCyTcyVJEsR9RUnbf715YuBAJMxCEEPzyZXgvY"),
		AuthorizedWithdrawer: arrayOfPubkey("B52Da5MCyTcyVJEsR9RUnbf715YuBAJMxCEEPzyZXgvY"),
		UnixTimestamp:        0,
		LockupEpoch:          0,
		Custodian:            arrayOfPubkey("11111111111111111111111111111111"),
		VoterPubkey:          arrayOfPubkey("5CgQubGD1uwodwCe5UXDADbC69SiqXR8qq6pDMSm7ut5"),
		Stake:                99997717120,
		ActivationEpoch:      79,
		DeactivationEpoch:    ^uint64(0),
		WarmupCooldownRate:   0.25,
		CreditsObserved:      21143,
	}

	deactivatedStakeState = StakeData{
		State:                2,
		RentExemptReserve:    2282880,
		AuthorizedStaker:     arrayOfPubkey("B52Da5MCyTcyVJEsR9RUnbf715YuBAJMxCEEPzyZXgvY"),
		AuthorizedWithdrawer: arrayOfPubkey("B52Da5MCyTcyVJEsR9RUnbf715YuBAJMxCEEPzyZXgvY"),
		UnixTimestamp:        0,
		LockupEpoch:          0,
		Custodian:            arrayOfPubkey("11111111111111111111111111111111"),
		VoterPubkey:          arrayOfPubkey("5CgQubGD1uwodwCe5UXDADbC69SiqXR8qq6pDMSm7ut5"),
		Stake:                99997717120,
		ActivationEpoch:      70,
		DeactivationEpoch:    78,
		WarmupCooldownRate:   0.25,
		CreditsObserved:      21143,
	}

	unpublishedValidatorStakeState = StakeData{
		State:                2,
		RentExemptReserve:    2282880,
		AuthorizedStaker:     arrayOfPubkey("B52Da5MCyTcyVJEsR9RUnbf715YuBAJMxCEEPzyZXgvY"),
		AuthorizedWithdrawer: arrayOfPubkey("B52Da5MCyTcyVJEsR9RUnbf715YuBAJMxCEEPzyZXgvY"),
		UnixTimestamp:        0,
		LockupEpoch:          0,
		Custodian:            arrayOfPubkey("11111111111111111111111111111111"),
		VoterPubkey:          arrayOfPubkey("BNTmegvdXzNVyc3UMTWSMSfJUryjr3fXEVErtdqrfs6y"),
		Stake:                99997717120,
		ActivationEpoch:      70,
		DeactivationEpoch:    78,
		WarmupCooldownRate:   0.25,
		CreditsObserved:      21143,
	}

	stakeValidator = blockatlas.StakeValidator{
		ID:     "5CgQubGD1uwodwCe5UXDADbC69SiqXR8qq6pDMSm7ut5",
		Status: true,
		Info: blockatlas.StakeValidatorInfo{
			Name:        "Certus One",
			Description: "Stake and earn rewards with the most secure and stable validator. Winner of the Game of Stakes. Operated by Certus One Inc. By delegating, you confirm that you are aware of the risk of slashing and that Certus One Inc is not liable for any potential damages to your investment.",
			Image:       "https://assets.trustwalletapp.com/blockchains/solana/validators/assets/2Afu38M1KaSfDBpjZjnJb9BSWP6YkBkoPiBfnFedD7JW/logo.png",
			Website:     "https://certus.one",
		},
		Details: getDetails(),
	}

	validatorMap = blockatlas.ValidatorMap{
		"5CgQubGD1uwodwCe5UXDADbC69SiqXR8qq6pDMSm7ut5": stakeValidator,
	}

	epochInfo = EpochInfo{
		AbsoluteSlot: 165763,
		Epoch:        80,
		SlotIndex:    1923,
		SlotsInEpoch: 2048,
	}

	delegation = blockatlas.DelegationsPage{
		{
			Delegator: stakeValidator,
			Value:     "99997717120",
			Status:    blockatlas.DelegationStatusActive,
		},
	}
)

func TestNormalizeValidator(t *testing.T) {
	var validators []VoteAccount
	minimumBalanceForRentExemption := uint64(2282880) // Arbitrary value; minimumBalanceForRentExemption is calculated from account data size
	err := json.Unmarshal([]byte(currentValidators), &validators)
	assert.Nil(t, err)
	for i, v := range validators {
		result := normalizeValidator(v, minimumBalanceForRentExemption)
		assert.Equal(t, expectedValidators[i], result)
	}
}

func TestParseStakeData(t *testing.T) {

	result, err := parseStakeData(keyedStakeAccount.Account)
	assert.Nil(t, err)
	assert.Equal(t, result, stakeState)
}

func TestIsAuthorized(t *testing.T) {
	address := "B52Da5MCyTcyVJEsR9RUnbf715YuBAJMxCEEPzyZXgvY"
	account, _ := parseStakeData(keyedStakeAccount.Account)
	assert.True(t, isAuthorized(account, address))

	address = "BADaddresscyVJEsR9RUnbf715YuBAJMxCEEPzyZXgvY"
	assert.False(t, isAuthorized(account, address))
}

func TestNormalizeDelegations(t *testing.T) {
	stakeAccounts := []StakeData{stakeState, deactivatedStakeState, unpublishedValidatorStakeState}
	result, err := NormalizeDelegations(stakeAccounts, validatorMap, epochInfo)
	assert.NoError(t, err)
	assert.Equal(t, delegation, result)
}
