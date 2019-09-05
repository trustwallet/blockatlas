package cosmos

import (
	"strconv"

	"github.com/trustwallet/blockatlas"

	"github.com/loomnetwork/blockatlas/coin"
	"github.com/spf13/viper"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client = InitClient(viper.GetString("loom.rpc"))
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.LOOM]
}

func (p *Platform) GetValidators() (blockatlas.ValidatorPage, error) {
	results := make(blockatlas.ValidatorPage, 0)
	validators, err := p.client.GetValidators()
	if err != nil {
		return results, nil
	}
	pool, err := p.client.GetPool()
	if err != nil {
		return results, nil
	}

	inflation, err := p.client.GetInflation()
	if err != nil {
		return results, nil
	}

	for _, validator := range validators {
		results = append(results, normalizeValidator(validator, pool, inflation, p.Coin()))
	}

	return results, nil
}

func normalizeValidator(v Validator, p StakingPool, inflation float64, c coin.Coin) (validator blockatlas.Validator) {

	reward := CalculateAnnualReward(p, inflation, v)

	return blockatlas.Validator{
		Status: bool(v.Status == 2),
		ID:     v.Address,
		Reward: blockatlas.StakingReward{Annual: reward},
	}
}

func CalculateAnnualReward(p StakingPool, inflation float64, validator Validator) float64 {

	notBondedTokens, err := strconv.ParseFloat(string(p.NotBondedTokens), 32)

	if err != nil {
		return 0
	}

	bondedTokens, err := strconv.ParseFloat(string(p.BondedTokens), 32)
	if err != nil {
		return 0
	}

	commission, err := strconv.ParseFloat(string(validator.Commission.Rate), 32)
	if err != nil {
		return 0
	}

	result := (notBondedTokens + bondedTokens) / bondedTokens * inflation

	return (result - (result * commission)) * 100
}
