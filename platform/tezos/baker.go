package tezos

import (
	"math"
	"strconv"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/services/assets"
	"github.com/trustwallet/golibs/coin"
)

type BakerClient struct {
	blockatlas.Request
}

func (c *BakerClient) GetBakers() (validators blockatlas.StakeValidators, err error) {
	var bakers []Baker
	err = c.Get(&bakers, "v2/bakers", nil)
	if err != nil {
		return
	}
	assetsValidators, err := assets.GetchValidatorsInfo(coin.Tezos())
	if err != nil {
		return
	}
	validatorMap := assetsValidators.ToMap()
	for _, baker := range bakers {
		if _, ok := validatorMap[baker.Address]; ok {
			validators = append(validators, NormalizeStakeValidator(baker))
		}
	}
	return
}

func NormalizeStakeValidator(baker Baker) blockatlas.StakeValidator {
	status := true
	if baker.FreeSpace < 0 || baker.ServiceHealth != "active" || !baker.OpenForDelegation {
		status = false
	}

	return blockatlas.StakeValidator{
		ID:     baker.Address,
		Status: status,
		Info: blockatlas.StakeValidatorInfo{
			Name:  baker.Name,
			Image: baker.Logo,
		},
		Details: blockatlas.StakingDetails{
			Reward: blockatlas.StakingReward{
				Annual: math.Round(baker.EstimatedRoi*10000) / 100,
			},
			LockTime:      LockTime,
			MinimumAmount: blockatlas.Amount(strconv.FormatUint(baker.MinDelegation, 10)),
			Type:          blockatlas.DelegationTypeDelegate,
		},
	}
}
