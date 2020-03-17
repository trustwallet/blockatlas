package parser

import (
	"encoding/json"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/storage"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type Parser struct {
	BlockAPI                 blockatlas.BlockAPI
	LatestParsedBlockTracker storage.Tracker
	ParsingBlocksInterval    time.Duration
	BacklogCount             int
	MaxBacklogBlocks         int64

	logParams   logger.Params
	mu          sync.Mutex
	coin        uint
	blockNumber int64
}

func (s *Parser) Run() {
	coin := s.BlockAPI.Coin()
	s.coin = coin.ID
	s.logParams = logger.Params{"platform": coin.Handle}
	conns := viper.GetInt("observer.stream_conns")
	if conns == 0 {
		logger.Fatal("observer.stream_conns is 0")
	}
	for {
		lastParsedBlock, currentBlock, err := s.getBlocksInterval()
		if err != nil {
			logger.Error(err)
		}
		s.fetchAndPublishBlocks(lastParsedBlock, currentBlock)
		time.Sleep(s.ParsingBlocksInterval)
	}
}

func (s *Parser) getBlocksInterval() (int64, int64, error) {
	lastParsedBlock, err := s.LatestParsedBlockTracker.GetBlockNumber(s.coin)
	if err != nil {
		return 0, 0, errors.E(err, "Polling failed: tracker didn't return last known block number", s.logParams)
	}
	currentBlock, err := s.BlockAPI.CurrentBlockNumber()
	currentBlock -= s.BlockAPI.Coin().MinConfirmations
	if err != nil {
		return 0, 0, errors.E(err, "Polling failed: source didn't return chain head number", s.logParams)
	}

	if currentBlock-lastParsedBlock > int64(s.BacklogCount) {
		lastParsedBlock = currentBlock - int64(s.BacklogCount)
	}

	if currentBlock-lastParsedBlock > s.MaxBacklogBlocks {
		lastParsedBlock = currentBlock - s.MaxBacklogBlocks
	}

	return lastParsedBlock, currentBlock, nil
}

func (s *Parser) fetchAndPublishBlocks(lastParsedBlock, currentBlock int64) {
	atomic.StoreInt64(&s.blockNumber, lastParsedBlock)
	var wg sync.WaitGroup

	for i := lastParsedBlock + 1; i <= currentBlock; i++ {
		wg.Add(1)
		go s.fetchAndPublishBlock(i, &wg)
	}
	wg.Wait()
}

func (s *Parser) fetchAndPublishBlock(num int64, wg *sync.WaitGroup) {
	defer wg.Done()

	block, err := getBlockByNumberWithRetry(5, time.Second*5, s.BlockAPI.GetBlockByNumber, num)
	if err != nil {
		logger.Error(err, "Polling failed: could not get block", s.logParams, logger.Params{"block": num})
		return
	}
	logger.Info(err, "Got new block", s.logParams, logger.Params{"block": num, "txs": len(block.Txs)})

	// Ignore block if it's not marked as parsed (set to redis) to prevent double notifications
	// TODO: add retry
	if err := s.addLatestParsedBlock(); err != nil {
		logger.Error("addLatestParsedBlock failed", s.logParams, logger.Params{"block": num, "coin": s.coin, "err": err})
		return
	}
	if err := s.publishBlock(*block); err != nil {
		logger.Error(err)
	}
}

func (s *Parser) addLatestParsedBlock() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	newNum := atomic.AddInt64(&s.blockNumber, 1)
	err := s.LatestParsedBlockTracker.SetBlockNumber(s.coin, newNum)
	if err != nil {
		return err
	}
	return nil
}

func (s *Parser) publishBlock(block blockatlas.Block) error {
	body, err := json.Marshal(block)
	if err != nil {
		return err
	}
	return mq.ConfirmedBlocks.Publish(body)
}

type (
	GetBlockByNumber func(num int64) (*blockatlas.Block, error)

	stop struct {
		error
	}
)

func getBlockByNumberWithRetry(attempts int, sleep time.Duration, getBlockByNumber GetBlockByNumber, n int64) (*blockatlas.Block, error) {
	r, err := getBlockByNumber(n)
	if err != nil {
		if s, ok := err.(stop); ok {
			return nil, s.error
		}
		if attempts--; attempts > 0 {
			// Add some randomness to prevent creating a Thundering Herd
			jitter := time.Duration(rand.Int63n(int64(sleep)))
			sleep = sleep + jitter/2

			logger.Info("retry GetBlockByNumber",
				logger.Params{
					"number":   n,
					"attempts": attempts,
					"sleep":    sleep.String(),
				},
			)

			time.Sleep(sleep)
			return getBlockByNumberWithRetry(attempts, sleep*2, getBlockByNumber, n)
		}
	}
	return r, err
}
