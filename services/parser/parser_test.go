package parser

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"sync"
	"testing"
	"time"
)

var (
	wantedMockedNumber int64

	block = blockatlas.Block{
		Number: 110,
		ID:     "",
		Txs: []blockatlas.Tx{
			{
				ID:     "95CF63FAA27579A9B6AF84EF8B2DFEAC29627479E9C98E7F5AE4535E213FA4C9",
				Coin:   coin.BNB,
				From:   "tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a",
				To:     "tbnb12hlquylu78cjylk5zshxpdj6hf3t0tahwjt3ex",
				Fee:    "125000",
				Date:   1555117625,
				Block:  7928667,
				Status: blockatlas.StatusCompleted,
				Memo:   "test",
				Meta: blockatlas.NativeTokenTransfer{
					TokenID:  "YLC-D8B",
					Symbol:   "YLC",
					Value:    "210572645",
					Decimals: 8,
					From:     "tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a",
					To:       "tbnb12hlquylu78cjylk5zshxpdj6hf3t0tahwjt3ex",
				},
			},
		},
	}
)

func TestFetchBlocks(t *testing.T) {
	params := Params{
		Ctx:                   nil,
		Api:                   getMockedBlockAPI(),
		Queue:                 []mq.Queue{""},
		ParsingBlocksInterval: 0,
		FetchBlocksTimeout:    0,
		BacklogCount:          0,
		MaxBacklogBlocks:      0,
		StopChannel:           nil,
		TxBatchLimit:          0,
		Database:              nil,
	}
	blocks := FetchBlocks(params, 0, 100, context.Background())
	assert.Equal(t, len(blocks), 100)
}

func TestParser_ConvertToBatch(t *testing.T) {
	blocks := []blockatlas.Block{block, block, block, block}
	txs := ConvertToBatch(blocks, context.Background())
	assert.Equal(t, 4, len(txs))

	empty := []blockatlas.Block{}
	txsEmpty := ConvertToBatch(empty, context.Background())
	assert.Equal(t, 0, len(txsEmpty))
}

func TestParser_add(t *testing.T) {
	blocks := []blockatlas.Block{block, block, block, block}
	txs := ConvertToBatch(blocks, context.Background())

	batch := transactionsBatch{
		Mutex: sync.Mutex{},
		Txs:   txs,
	}

	batch.fillBatch(txs, context.Background())
	assert.Equal(t, 8, len(batch.Txs))

	batch.fillBatch(nil, context.Background())
	assert.Equal(t, 8, len(batch.Txs))
}

func TestParser_getBlockByNumberWithRetry(t *testing.T) {
	block, err := getBlockByNumberWithRetry(3, time.Millisecond*1, getBlock, 1, "", context.Background())
	if err != nil {
		t.Error(err)
	}

	if block == nil {
		t.Error("block is nil")
	}
}

func TestParser_getBlockByNumberWithRetry_Error(t *testing.T) {
	now := time.Now()
	block, err := getBlockByNumberWithRetry(2, time.Millisecond*2, getBlock, 0, "", context.Background())
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

func TestGetTxBatches(t *testing.T) {
	txs := make(blockatlas.Txs, 10000)
	batches := getTxsBatches(txs, 1000, context.Background())
	assert.Len(t, batches, 10)
	batches = getTxsBatches(txs, 100, context.Background())
	assert.Len(t, batches, 100)
	batches = getTxsBatches(txs, 500, context.Background())
	assert.Len(t, batches, 20)

	txs = make(blockatlas.Txs, 3800)
	batches = getTxsBatches(txs, 100, context.Background())
	assert.Len(t, batches, 38)
	batches = getTxsBatches(txs, 1000, context.Background())
	assert.Len(t, batches, 4)

	txs = make(blockatlas.Txs, 5000)
	batches = getTxsBatches(txs, 10000, context.Background())
	assert.Len(t, batches, 1)

	txs = make(blockatlas.Txs, 0)
	batches = getTxsBatches(txs, 100, context.Background())
	assert.Len(t, batches, 0)

	txs = make(blockatlas.Txs, 0)
	batches = getTxsBatches(txs, 100, context.Background())
	assert.Len(t, batches, 0)

	batches = getTxsBatches(nil, 100, context.Background())
	assert.Len(t, batches, 0)

	txs = make(blockatlas.Txs, 1000000)
	batches = getTxsBatches(txs, 5000, context.Background())
	assert.Len(t, batches, 200)
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
