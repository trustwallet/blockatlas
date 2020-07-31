package stake

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
)

func GetStakingReward(api blockatlas.StakeAPI) blockatlas.StakingReward {
	validators, err := api.GetActiveValidators()
	if err != nil {
		logger.Error("GetMaxAPR", logger.Params{"details": err, "platform": api.Coin().Symbol})
		return blockatlas.StakingReward{
			Annual: blockatlas.DefaultAnnualReward,
		}
	}

	return blockatlas.StakingReward{
		Annual: getMaxApr(validators),
	}
}

func getMaxApr(validators []blockatlas.StakeValidator) float64 {
	var max = 0.0
	for _, e := range validators {
		v := e.Details.Reward.Annual
		if v > max {
			max = v
		}
	}
	return max
}
