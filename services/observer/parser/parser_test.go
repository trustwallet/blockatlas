package parser

import (
	"errors"
	"fmt"
	"github.com/alicebob/miniredis"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/storage"
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

func Test_GetBlocksIntervalToFetch(t *testing.T) {
	params := Params{
		ParsingBlocksInterval: time.Minute,
		BacklogCount:          10,
		MaxBacklogBlocks:      100,
	}
	latestParsedBlock := int64(100)
	wantedMockedNumber = 110
	lastParsedBlock, currentBlock, err := GetBlocksIntervalToFetch(getMockedBlockAPI(), getMockedRedis(t), params)
	assert.Nil(t, err)
	assert.Equal(t, wantedMockedNumber, currentBlock)
	assert.Equal(t, latestParsedBlock, lastParsedBlock)
}

func TestFetchBlocks(t *testing.T) {
	blocks, err := FetchBlocks(getMockedBlockAPI(), 0, 100)
	assert.Nil(t, err)
	assert.Equal(t, len(blocks), 100)
}

func TestSaveLastParsedBlock(t *testing.T) {
	params := Params{
		ParsingBlocksInterval: time.Minute,
		BacklogCount:          10,
		MaxBacklogBlocks:      100,
	}
	blocks := make([]blockatlas.Block, 0)
	blocks = append(blocks, block)
	s := getMockedRedis(t)

	err := SaveLastParsedBlock(s, params, blocks)
	assert.Nil(t, err)

	lastParsedBlock, err := s.GetLastParsedBlockNumber(params.Coin)
	assert.Nil(t, err)
	assert.Equal(t, lastParsedBlock, int64(110))
}

func TestParser_getBlockByNumberWithRetry(t *testing.T) {
	block, err := getBlockByNumberWithRetry(3, time.Millisecond*1, getBlock, 1)
	if err != nil {
		t.Error(err)
	}

	if block == nil {
		t.Error("block is nil")
	}
}

func TestParser_getBlockByNumberWithRetry_Error(t *testing.T) {
	now := time.Now()
	block, err := getBlockByNumberWithRetry(2, time.Millisecond*2, getBlock, 0)
	elapsed := time.Since(now)
	if err == nil {
		t.Error("getBlockByNumberWithRetry method need fail")
	}

	if block != nil {
		t.Error("block object need be nil")
	}

	if elapsed > time.Millisecond*8 {
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

func getMockedRedis(t *testing.T) *storage.Storage {
	s, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}

	cache := storage.New()
	err = cache.Init(fmt.Sprintf("redis://%s", s.Addr()))
	if err != nil {
		logger.Fatal(err)
	}
	return cache
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
