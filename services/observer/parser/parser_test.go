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

var wantedMockedNumber int64

func Test_getBlocksInterval(t *testing.T) {
	p := Parser{
		BlockAPI:              getMockedBlockAPI(),
		Storage:               getMockedRedis(t),
		ParsingBlocksInterval: time.Minute,
		BacklogCount:          10,
		MaxBacklogBlocks:      100,
	}
	latestParsedBlock := int64(100)
	wantedMockedNumber = 110
	lastParsedBlock, currentBlock, err := p.getBlocksIntervalToFetch()
	assert.Nil(t, err)
	assert.Equal(t, wantedMockedNumber, currentBlock)
	assert.Equal(t, latestParsedBlock, lastParsedBlock)
}

func Test_addLatestParsedBlock(t *testing.T) {
	p := Parser{
		BlockAPI:              getMockedBlockAPI(),
		Storage:               getMockedRedis(t),
		ParsingBlocksInterval: time.Minute,
		BacklogCount:          10,
		MaxBacklogBlocks:      100,
	}
	//err := p.SaveLastParsedBlock()
	//assert.Nil(t, err)
	//block, err := p.Storage.GetLastParsedBlockNumber(p.coin)
	//assert.Nil(t, err)
	//assert.Equal(t, int64(1), block)
	//
	//err = p.addLatestParsedBlock()
	//assert.Nil(t, err)
	//blockTwo, err := p.Storage.GetLastParsedBlockNumber(p.coin)
	//assert.Nil(t, err)
	//assert.Equal(t, int64(2), blockTwo)
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
