package tezos

import (
	"fmt"
	"strconv"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type BakerClient struct {
	blockatlas.Request
}

func (c *Client) GetBakers() (validators blockatlas.StakeValidators, err error) {
	var bakers []Baker
	path := fmt.Sprintf("/v2/bakers")
	err = c.Get(&bakers, path, nil)
	if err != nil {
		return nil, err
	}
	return NormalizeStakeValidators(bakers), nil
}

func NormalizeStakeValidators(bakers []Baker) (validators blockatlas.StakeValidators) {
	for _, baker := range bakers {
		status := true
		if baker.FreeSpace < 0 || baker.ServiceHealth != "active" || baker.OpenForDelegation == false {
			status = false
		}

		validator := blockatlas.StakeValidator{
			ID:     baker.Address,
			Status: status,
			Info: blockatlas.StakeValidatorInfo{
				Name:  baker.Name,
				Image: baker.Logo,
			},
			Details: blockatlas.StakingDetails{
				Reward: blockatlas.StakingReward{
					Annual: baker.EstimatedRoi,
				},
				LockTime:      LockTime,
				MinimumAmount: blockatlas.Amount(strconv.FormatUint(baker.MinDelegation, 10)),
				Type:          blockatlas.DelegationTypeDelegate,
			},
		}
		validators = append(validators, validator)
	}
	return
}
