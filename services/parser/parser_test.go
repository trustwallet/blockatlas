package parser

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/types"
)

var (
	wantedMockedNumber int64
)

func TestFetchBlocks(t *testing.T) {
	params := Params{
		Api:                   getMockedBlockAPI(),
		TransactionsExchange:  "",
		ParsingBlocksInterval: 0,
		FetchBlocksTimeout:    0,
		MaxBlocks:             0,
		StopChannel:           nil,
		Database:              nil,
	}
	blocks, err := FetchBlocks(params, 0, 100)
	assert.Equal(t, len(blocks), 100)
	assert.Nil(t, err)
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

func getBlock(num int64) (*types.Block, error) {
	if num == 0 {
		return nil, errors.New("test")
	}
	return &types.Block{}, nil
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

func (p *Platform) GetBlockByNumber(num int64) (*types.Block, error) {
	return &types.Block{}, nil
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

func TestGetNextBlocksToParse(t *testing.T) {
	type args struct {
		lastParsedBlock int64
		currentBlock    int64
		maxBlocks       int64
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		want1   int64
		wantErr bool
	}{
		{
			"Test behind blocks",
			args{
				lastParsedBlock: 10,
				currentBlock:    25,
				maxBlocks:       3,
			},
			11,
			15,
			false,
		},
		{
			"Test when only 1 block to parse",
			args{
				lastParsedBlock: 10,
				currentBlock:    13,
				maxBlocks:       5,
			},
			11,
			14,
			false,
		},
		{
			"Test same block",
			args{
				lastParsedBlock: 10,
				currentBlock:    10,
				maxBlocks:       3,
			},
			10,
			10,
			false,
		},
		{
			"Last parsed block ahead",
			args{
				lastParsedBlock: 15,
				currentBlock:    10,
				maxBlocks:       3,
			},
			15,
			15,
			false,
		},
		{
			"Parse last block",
			args{
				lastParsedBlock: 15,
				currentBlock:    15,
				maxBlocks:       3,
			},
			15,
			15,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := GetNextBlocksToParse(tt.args.lastParsedBlock, tt.args.currentBlock, tt.args.maxBlocks)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNextBlocksToParse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetNextBlocksToParse() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetNextBlocksToParse() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
