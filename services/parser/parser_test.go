package parser

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/coin"
)

var (
	wantedMockedNumber int64
)

func TestFetchBlocks(t *testing.T) {
	params := Params{
		Ctx:                    nil,
		Api:                    getMockedBlockAPI(),
		TransactionsQueue:      "",
		TokenTransactionsQueue: []mq.Queue{""},
		ParsingBlocksInterval:  0,
		FetchBlocksTimeout:     0,
		BacklogCount:           0,
		MaxBacklogBlocks:       0,
		StopChannel:            nil,
		Database:               nil,
	}
	blocks := FetchBlocks(params, 0, 100)
	assert.Equal(t, len(blocks), 100)
}

func TestParser_getBlockByNumberWithRetry(t *testing.T) {
	block, err := getBlockByNumberWithRetry(3, time.Millisecond*1, getBlock, 1, "")
	if err != nil {
		t.Error(err)
	}

	if block == nil {
		t.Error("block is nil")
	}
}

func TestParser_getBlockByNumberWithRetry_Error(t *testing.T) {
	now := time.Now()
	block, err := getBlockByNumberWithRetry(2, time.Millisecond*2, getBlock, 0, "")
	elapsed := time.Since(now)
	if err == nil {
		t.Error("getBlockByNumberWithRetry method need fail")
	}

	if block != nil {
		t.Error("block object need be nil")
	}

	if elapsed > time.Millisecond*7 {
		t.Error("Thundering Herd prevent doesn't work")
	}
}

func getBlock(num int64) (*blockatlas.Block, error) {
	if num == 0 {
		return nil, errors.New("test")
	}
	return &blockatlas.Block{}, nil
}

func getMockedBlockAPI() blockatlas.BlockAPI {
	p := Platform{CoinIndex: 60}
	return &p
}

type Platform struct {
	CoinIndex uint
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return wantedMockedNumber, nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[p.CoinIndex]
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	return &blockatlas.Block{}, nil
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
