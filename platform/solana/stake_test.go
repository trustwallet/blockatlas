package solana

import (
	"encoding/json"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"

	"github.com/stretchr/testify/assert"
)

const currentValidators = `
[
  {
    "activatedStake": 3733867423940,
    "commission": 100,
    "epochCredits": [
      [70,524224,516032],[71,532416,524224],[72,540608,532416],[73,548800,540608],[74,556992,548800],[75,565184,556992],[76,573376,565184],[77,581568,573376],[78,589760,581568],[79,597952,589760],[80,601055,597952]
    ],
    "epochVoteAccount": true,
    "lastVote": 601085,
    "nodePubkey": "boot1Z6jb15CLqpaMTn2CxktktwZpRAVAgHZEW6SxQ7",
    "rootSlot": 601054,
    "votePubkey": "2Afu38M1KaSfDBpjZjnJb9BSWP6YkBkoPiBfnFedD7JW"
  },
  {
    "activatedStake": 10540011934,
    "commission": 100 ,
    "epochCredits": [
      [78,4760,0],[79,12952,4760],[80,16055,12952]
    ],
    "epochVoteAccount": true,
    "lastVote": 601085,
    "nodePubkey": "B52Da5MCyTcyVJEsR9RUnbf715YuBAJMxCEEPzyZXgvY",
    "rootSlot": 601054,
    "votePubkey": "5CgQubGD1uwodwCe5UXDADbC69SiqXR8qq6pDMSm7ut5"
  }
]`

var expectedValidators = []blockatlas.Validator{
	blockatlas.Validator{
		Status: true,
		ID:     "2Afu38M1KaSfDBpjZjnJb9BSWP6YkBkoPiBfnFedD7JW",
		Details: blockatlas.StakingDetails{
			Reward:        blockatlas.StakingReward{Annual: 0},
			MinimumAmount: blockatlas.Amount("2282881"),
			LockTime:      0,
			Type:          blockatlas.DelegationTypeDelegate,
		},
	},
	blockatlas.Validator{
		Status: true,
		ID:     "5CgQubGD1uwodwCe5UXDADbC69SiqXR8qq6pDMSm7ut5",
		Details: blockatlas.StakingDetails{
			Reward:        blockatlas.StakingReward{Annual: 0},
			MinimumAmount: blockatlas.Amount("2282881"),
			LockTime:      0,
			Type:          blockatlas.DelegationTypeDelegate,
		},
	},
}

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

var keyedStakeAccount = KeyedAccount{
	Account: Account{
		Data:       "mrWMHx6j3BkmepJy67XLycC7LVeiq2NBESfV2YNmZvY62xT5jTgKnMzRBQheYtAuajncAniTEmU8QxgkpnytnXynTrMSJN4p6ihefU5cobkyCeSwMKugKGuBbyDLjQoUMu6BUKjDTFvjpJUHFCgz1Vaa8HSVscUqqRcioByf3owMUwHmEYsF8vuouLAqEQmo61wFkKfZELxLrhBbi2PQQZucryrnNDKXV4DY3oegLy9aMnMDZUeoDtDPPiJeM2F1Trh8ZkH1sQL6sQ5V",
		Executable: false,
		Lamports:   100000000000,
		Owner:      "Stake11111111111111111111111111111111111111",
		RentEpoch:  80,
	},
	Pubkey: "EgR17fgGmwQQaMZPsuJdk9oHw2xY8TJQj3Bp44o24mar",
}

var stakeState = StakeState{
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

var stakeValidator = blockatlas.StakeValidator{
	ID:     "5CgQubGD1uwodwCe5UXDADbC69SiqXR8qq6pDMSm7ut5",
	Status: true,
	Info: blockatlas.StakeValidatorInfo{
		Name:        "Certus One",
		Description: "Stake and earn rewards with the most secure and stable validator. Winner of the Game of Stakes. Operated by Certus One Inc. By delegating, you confirm that you are aware of the risk of slashing and that Certus One Inc is not liable for any potential damages to your investment.",
		Image:       "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/solana/validators/assets/2Afu38M1KaSfDBpjZjnJb9BSWP6YkBkoPiBfnFedD7JW/logo.png",
		Website:     "https://certus.one",
	},
	Details: getDetails(),
}

var validatorMap = blockatlas.ValidatorMap{
	"5CgQubGD1uwodwCe5UXDADbC69SiqXR8qq6pDMSm7ut5": stakeValidator,
}

var epochInfo = EpochInfo{
	AbsoluteSlot: 165763,
	Epoch:        80,
	SlotIndex:    1923,
	SlotsInEpoch: 2048,
}

var delegation = blockatlas.DelegationsPage{
	{
		Delegator: stakeValidator,
		Value:     "99997717120",
		Status:    blockatlas.DelegationStatusActive,
	},
}

func TestNormalizeDelegations(t *testing.T) {
	result, err := NormalizeDelegations([]StakeState{stakeState}, validatorMap, epochInfo)
	assert.NoError(t, err)
	assert.Equal(t, delegation, result)
}
