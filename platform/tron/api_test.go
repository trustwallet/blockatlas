package tron

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

const transferSrc = `
{
	"block_timestamp": 1564797900000,
	"raw_data": {
		"contract": [
			{
				"parameter": {
					"value": {
						"amount": 100666888000000,
						"owner_address": "4182dd6b9966724ae2fdc79b416c7588da67ff1b35",
						"to_address": "410583a68a3bcd86c25ab1bee482bac04a216b0261"
					}
				},
				"type": "TransferContract"
			}
		]
	},
	"txID": "24a10f7a503e78adc0d7e380b68005531b09e16b9e3f7b524e33f40985d287df"
}
`

const tokenTransferSrc = `
{
	"block_timestamp": 1564797900000,
	"raw_data": {
		"contract": [
			{
				"parameter": {
					"value": {
						"amount": 2776267,
						"asset_name": "1002000",
						"owner_address": "4182dd6b9966724ae2fdc79b416c7588da67ff1b35",
						"to_address": "410583a68a3bcd86c25ab1bee482bac04a216b0261"
					}
				},
				"type": "TransferAssetContract"
			}
		]
	},
	"txID": "24a10f7a503e78adc0d7e380b68005531b09e16b9e3f7b524e33f40985d287df"
}
`

var transferDst = blockatlas.Tx{
	ID:     "24a10f7a503e78adc0d7e380b68005531b09e16b9e3f7b524e33f40985d287df",
	Coin:   coin.TRX,
	From:   "TMuA6YqfCeX8EhbfYEg5y7S4DqzSJireY9",
	To:     "TAUN6FwrnwwmaEqYcckffC7wYmbaS6cBiX",
	Fee:    "0", // TODO
	Date:   1564797900,
	Block:  0, // TODO
	Status: blockatlas.StatusCompleted,
	Meta: blockatlas.Transfer{
		Value:    "100666888000000",
		Symbol:   "TRX",
		Decimals: 6,
	},
}

var tokenTransferDst = blockatlas.Tx{
	ID:     "24a10f7a503e78adc0d7e380b68005531b09e16b9e3f7b524e33f40985d287df",
	Coin:   coin.TRX,
	From:   "TMuA6YqfCeX8EhbfYEg5y7S4DqzSJireY9",
	To:     "TAUN6FwrnwwmaEqYcckffC7wYmbaS6cBiX",
	Fee:    "0", // TODO
	Date:   1564797900,
	Block:  0, // TODO
	Status: blockatlas.StatusCompleted,
	Meta: blockatlas.TokenTransfer{
		Name:     "BitTorrent",
		Symbol:   "BTT",
		TokenID:  "1002000",
		Decimals: 6,
		Value:    "2776267",
		From:     "TMuA6YqfCeX8EhbfYEg5y7S4DqzSJireY9",
		To:       "TAUN6FwrnwwmaEqYcckffC7wYmbaS6cBiX",
	},
}

var assetInfo = AssetInfo{Name: "BitTorrent", Symbol: "BTT", Decimals: 6, ID: "1002000"}

type test struct {
	name        string
	apiResponse string
	expected    *blockatlas.Tx
}

func TestNormalize(t *testing.T) {
	testNormalize(t, &test{
		name:        "transfer",
		apiResponse: transferSrc,
		expected:    &transferDst,
	})
}

func TestNormalizeTokenTransfer(t *testing.T) {
	testNormalizeTokenTransfer(t, &test{
		name:        "token transfer",
		apiResponse: tokenTransferSrc,
		expected:    &tokenTransferDst,
	})
}

func testNormalizeTokenTransfer(t *testing.T, _test *test) {
	var srcTx Tx
	err := json.Unmarshal([]byte(_test.apiResponse), &srcTx)
	if err != nil {
		t.Error(err)
		return
	}
	res, err := NormalizeTokenTransfer(&srcTx, assetInfo)
	if err != nil {
		t.Errorf("%s: tx could not be normalized", _test.name)
		return
	}

	actual, err := json.Marshal(&res)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := json.Marshal(_test.expected)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

func testNormalize(t *testing.T, _test *test) {
	var srcTx Tx
	err := json.Unmarshal([]byte(_test.apiResponse), &srcTx)
	if err != nil {
		t.Error(err)
		return
	}
	res, ok := Normalize(&srcTx)
	if !ok {
		t.Errorf("%s: tx could not be normalized", _test.name)
		return
	}

	actual, err := json.Marshal(&res)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := json.Marshal(&transferDst)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

var tokenDst = blockatlas.Token{
	Name:     "Test",
	Symbol:   "TST",
	Decimals: 8,
	TokenID:  "1",
	Coin:     195,
	Type:     "TRC10",
}

func TestNormalizeToken(t *testing.T) {
	asset := AssetInfo{Name: "Test", Symbol: "TST", ID: "1", Decimals: 8}
	actual := NormalizeToken(asset)
	assert.Equal(t, tokenDst, actual)
}

func TestNormalizeValidator(t *testing.T) {
	validator := Validator{Address: "414d1ef8673f916debb7e2515a8f3ecaf2611034aa"}

	actual, _ := normalizeValidator(validator)
	expected := blockatlas.Validator{
		ID:     "TGzz8gjYiYRqpfmDwnLxfgPuLVNmpCswVp",
		Status: true,
		Details: blockatlas.StakingDetails{
			Reward: blockatlas.StakingReward{
				Annual: Annual,
			},
			LockTime:      259200,
			MinimumAmount: "1000000",
			Type:          blockatlas.DelegationTypeDelegate,
		},
	}
	assert.Equal(t, expected, actual)
}

const delegationsSrc1 = `
{
  "address": "419241920da7d6bb487a33a6df3838e3d208f0b251",
  "frozen": [
	{
	  "expire_time": 10437262001000,
	  "frozen_balance": "35000000"
	}
  ],
  "votes": [
	{
	  "vote_address": "TGzz8gjYiYRqpfmDwnLxfgPuLVNmpCswVp",
	  "vote_count": 21
	},
	{
	  "vote_address": "TPMGfspxLQGom8sKutrbHcDKtHjRHFbGKw",
	  "vote_count": 5
	},
	{
	  "vote_address": "TPMGfspxLQGom8sKutrbHcDKtHjRHFbGKw",
	  "vote_count": 5
	}
  ]
}`

const delegationsSrc2 = `
{
  "address": "419241920da7d6bb487a33a6df3838e3d208f0b251",
  "frozen": [
	{
	  "expire_time": 1569465251000,
	  "frozen_balance": "5000000"
	}
  ],
  "votes": [
	{
	  "vote_address": "TPMGfspxLQGom8sKutrbHcDKtHjRHFbGKw",
	  "vote_count": 5
	}
  ]
}`

var validator1 = blockatlas.StakeValidator{
	ID:     "TGzz8gjYiYRqpfmDwnLxfgPuLVNmpCswVp",
	Status: true,
	Info: blockatlas.StakeValidatorInfo{
		Name:        "Sesameseed",
		Description: "Sesameseed is a blockchain community providing fair and transparent representation in delegated governance by rewarding Voters for their participation on Tron and Ontology.",
		Image:       "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/tron/validators/assets/tgzz8gjyiyrqpfmdwnlxfgpulvnmpcswvp/logo.png",
		Website:     "https://www.sesameseed.org",
	},
	Details: blockatlas.StakingDetails{
		Reward: blockatlas.StakingReward{
			Annual: 4.32,
		},
		LockTime:      259200,
		MinimumAmount: "1000000",
	},
}

var validator2 = blockatlas.StakeValidator{
	ID:     "TPMGfspxLQGom8sKutrbHcDKtHjRHFbGKw",
	Status: true,
	Info: blockatlas.StakeValidatorInfo{
		Name:        "InfStones",
		Description: "World's leading cloud infrastructure and staking as a service provicer for blockchains. Supernodes on EOS, TRON, VeChain, Ontology, LOOM, IOST and many other chains.",
		Image:       "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/tron/validators/assets/tpmgfspxlqgom8skutrbhcdkthjrhfbgkw/logo.png",
		Website:     "https://infstones.io/",
	},
	Details: blockatlas.StakingDetails{
		Reward: blockatlas.StakingReward{
			Annual: 4.32,
		},
		LockTime:      259200,
		MinimumAmount: "1000000",
	},
}

var delegation1 = blockatlas.Delegation{
	Delegator: validator1,
	Value:     "21000000",
	Status:    blockatlas.DelegationStatusPending,
}
var delegation2 = blockatlas.Delegation{
	Delegator: validator2,
	Value:     "5000000",
	Status:    blockatlas.DelegationStatusPending,
}
var delegation3 = blockatlas.Delegation{
	Delegator: validator2,
	Value:     "5000000",
	Status:    blockatlas.DelegationStatusPending,
}
var delegation4 = blockatlas.Delegation{
	Delegator: validator2,
	Value:     "5000000",
	Status:    blockatlas.DelegationStatusActive,
}

var validatorMap = blockatlas.ValidatorMap{
	"TGzz8gjYiYRqpfmDwnLxfgPuLVNmpCswVp": validator1,
	"TPMGfspxLQGom8sKutrbHcDKtHjRHFbGKw": validator2,
}

func TestNormalizeDelegations(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  []blockatlas.Delegation
	}{
		{"Status Pending", delegationsSrc1, []blockatlas.Delegation{delegation1, delegation2, delegation3}},
		{"Status Active", delegationsSrc2, []blockatlas.Delegation{delegation4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testNormalizeDelegations(t, tt.value, tt.want)
		})
	}
}

func testNormalizeDelegations(t *testing.T, value string, want []blockatlas.Delegation) {
	var accountData *AccountData
	err := json.Unmarshal([]byte(value), &accountData)
	assert.NoError(t, err)
	assert.NotNil(t, accountData)
	result := NormalizeDelegations(accountData, validatorMap)
	assert.Equal(t, result, want)
}
