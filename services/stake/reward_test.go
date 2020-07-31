package stake

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

func Test_getMaxApr(t *testing.T) {
	type args struct {
		validators []blockatlas.StakeValidator
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"Test empty validators list",
			args{
				validators: []blockatlas.StakeValidator{},
			},
			0,
		},
		{
			"Test max APR validators list",
			args{
				validators: []blockatlas.StakeValidator{
					{
						Details: blockatlas.StakingDetails{
							Reward: blockatlas.StakingReward{
								Annual: 10,
							},
						},
					},
					{
						Details: blockatlas.StakingDetails{
							Reward: blockatlas.StakingReward{
								Annual: 5,
							},
						},
					},
				},
			},
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaxApr(tt.args.validators); got != tt.want {
				t.Errorf("getMaxApr() = %v, want %v", got, tt.want)
			}
		})
	}
}
