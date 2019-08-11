package cosmos

import (
	"github.com/trustwallet/blockatlas"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/util"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client.BaseURL = viper.GetString("cosmos.api")
	p.client.HTTPClient = http.DefaultClient
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.ATOM]
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	inputTxes, _ := p.client.GetAddrTxes(address, "inputs")
	outputTxes, _ := p.client.GetAddrTxes(address, "outputs")

	normalisedTxes := make([]blockatlas.Tx, 0)

	for _, inputTx := range inputTxes {
		normalisedInputTx := Normalize(&inputTx)
		normalisedTxes = append(normalisedTxes, normalisedInputTx)
	}
	for _, outputTx := range outputTxes {
		normalisedOutputTx := Normalize(&outputTx)
		normalisedTxes = append(normalisedTxes, normalisedOutputTx)
	}

	sort.Slice(normalisedTxes, func(i, j int) bool {
		return normalisedTxes[i].Date > normalisedTxes[j].Date
	})

	return normalisedTxes, nil
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

// Normalize converts an Cosmos transaction into the generic model
func Normalize(srcTx *Tx) (tx blockatlas.Tx) {
	date, _ := time.Parse("2006-01-02T15:04:05Z", srcTx.Date)
	value, _ := util.DecimalToSatoshis(srcTx.Data.Contents.Message[0].Particulars.Amount[0].Quantity)
	block, _ := strconv.ParseUint(srcTx.Block, 10, 64)
	// Sometimes fees can be null objects (in the case of no fees e.g. F044F91441C460EDCD90E0063A65356676B7B20684D94C731CF4FAB204035B41)
	var fee string
	if len(srcTx.Data.Contents.Fee.FeeAmount) == 0 {
		fee = "0"
	} else {
		fee, _ = util.DecimalToSatoshis(srcTx.Data.Contents.Fee.FeeAmount[0].Quantity)
	}
	return blockatlas.Tx{
		ID:    srcTx.ID,
		Coin:  coin.ATOM,
		Date:  date.Unix(),
		From:  srcTx.Data.Contents.Message[0].Particulars.FromAddr, // This will need to be adjusted for multi-outputs, later
		To:    srcTx.Data.Contents.Message[0].Particulars.ToAddr,   // Likewise
		Fee:   blockatlas.Amount(fee),
		Block: block,
		Memo:  srcTx.Data.Contents.Memo,
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(value),
			Symbol:   coin.Coins[coin.ATOM].Symbol,
			Decimals: coin.Coins[coin.ATOM].Decimals,
		},
	}
}

func normalizeValidator(v CosmosValidator, p StakingPool, inflation float64, c coin.Coin) (validator blockatlas.Validator) {

	reward := CalculateAnnualReward(p, inflation, v)

	return blockatlas.Validator{
		Coin:   c,
		Status: bool(v.Status == 2),
		ID:     v.Operator_Address,
		Reward: blockatlas.StakingReward{Annual: reward},
	}
}

func CalculateAnnualReward(p StakingPool, inflation float64, validator CosmosValidator) float64 {

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
