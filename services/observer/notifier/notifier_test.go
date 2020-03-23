package notifier

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

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
