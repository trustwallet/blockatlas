package numbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMin(t *testing.T) {
	assert.Equal(t, Min(1, 5), 1)
	assert.Equal(t, Min(22, 5), 5)
}

func TestMax(t *testing.T) {
	assert.Equal(t, Max(1, 5), int64(5))
	assert.Equal(t, Max(22, 5), int64(22))
}

func TestToDecimal(t *testing.T) {
	assert.Equal(t, ToDecimal("0", 18), "0")
	assert.Equal(t, ToDecimal("100", 1), "10")
	assert.Equal(t, ToDecimal("123123", 3), "123.123")
	assert.Equal(t, ToDecimal("10012000000000000", 12), "10012")
	assert.Equal(t, ToDecimal("123456789012345678901", 18), "123.4567890123")
	assert.Equal(t, ToDecimal("4618", 6), "0.004618")
	assert.Equal(t, ToDecimal("218218", 8), "0.00218218")
	assert.Equal(t, ToDecimal("212880628", 9), "0.212880628")
	assert.Equal(t, ToDecimal("4634460765323682", 18), "0.0046344608")
}

func TestFromDecimal(t *testing.T) {
	assert.Equal(t, FromDecimal("100.12"), "10012")
}

func TestToDecimalExp(t *testing.T) {
	assert.Equal(t, FromDecimalExp("10", 1), "100")
	assert.Equal(t, FromDecimalExp("100", 1), "1000")
	assert.Equal(t, FromDecimalExp("10012", 12), "10012000000000000")
	assert.Equal(t, FromDecimalExp("123.123", 3), "123123")
	//assert.Equal(t, FromDecimalExp("0.005170630816959669", 2), "") Need fix
	assert.Equal(t, FromDecimalExp("0.000180508184692364", 4), "1")
	assert.Equal(t, FromDecimalExp("0.004618071835862274", 6), "4618")
	assert.Equal(t, FromDecimalExp("0.00216013705800604", 8), "216013")
	assert.Equal(t, FromDecimalExp("0.002182187913804679", 8), "218218")
	assert.Equal(t, FromDecimalExp("0.21288062808828456", 9), "212880628")
	assert.Equal(t, FromDecimalExp("0.004634460765323682", 18), "4634460765323682")
	assert.Equal(t, FromDecimalExp("0.00000001", 8), "1")
	assert.Equal(t, FromDecimalExp("10.00000000", 8), "1000000000")
}

func TestFloat64toPrecision(t *testing.T) {
	assert.Equal(t, Float64toPrecision(3.643005, 4), 3.6430)
	assert.Equal(t, Float64toPrecision(9.8233168e-5, 4), 0.0001)
	assert.Equal(t, Float64toPrecision(0.8010, 4), 0.8010)
	assert.Equal(t, Float64toPrecision(26.5, 4), 26.5)
	assert.Equal(t, Float64toPrecision(3374, 4), 3374.0)
}

func TestGetInterval(t *testing.T) {
	min, _ := time.ParseDuration("2s")
	max, _ := time.ParseDuration("30s")
	type args struct {
		blockTime   int
		minInterval time.Duration
		maxInterval time.Duration
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{
			"test minimum",
			args{
				blockTime:   100,
				minInterval: min,
				maxInterval: max,
			},
			min,
		}, {
			"test maximum",
			args{
				blockTime:   600000,
				minInterval: min,
				maxInterval: max,
			},
			max,
		}, {
			"test right blocktime",
			args{
				blockTime:   5000,
				minInterval: min,
				maxInterval: max,
			},
			5000 * time.Millisecond,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetInterval(tt.args.blockTime, tt.args.minInterval, tt.args.maxInterval)
			assert.EqualValues(t, tt.want, got)
		})
	}
}
