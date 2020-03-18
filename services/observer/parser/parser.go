package parser

import (
	"encoding/json"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/storage"
	"golang.org/x/sync/errgroup"
	"math/rand"
	"sort"
	"time"
)

type (
	Params struct {
		ParsingBlocksInterval time.Duration
		BacklogCount          int
		MaxBacklogBlocks      int64
		Coin                  uint
	}

	GetBlockByNumber func(num int64) (*blockatlas.Block, error)

	stop struct {
		error
	}
)

func RunParser(api blockatlas.BlockAPI, storage storage.Tracker, config Params) {
	logger.Info("------------------------------------------------------------")
	for {
		lastParsedBlock, currentBlock, err := getBlocksIntervalToFetch(api, storage, config)
		if err != nil {
			logger.Error(err)
		}

		blocks, err := FetchBlocks(api, lastParsedBlock, currentBlock)
		if err != nil {
			logger.Error(err)
		}

		err = SaveLastParsedBlock(storage, config, blocks)
		if err != nil {
			logger.Error(err)
		}

		err = PublishBlocks(blocks)
		if err != nil {
			logger.Error(err)
		}

		time.Sleep(config.ParsingBlocksInterval)
	}
}

func FetchBlocks(api blockatlas.BlockAPI, lastParsedBlock, currentBlock int64) ([]blockatlas.Block, error) {
	if lastParsedBlock == currentBlock {
		logger.Info("No new blocks", logger.Params{"last": lastParsedBlock, "coin": api.Coin().ID, "time": time.Now().Unix()})
		logger.Info("------------------------------------------------------------")
		return nil, nil
	}

	blocksChan := make(chan blockatlas.Block, currentBlock)

	var g errgroup.Group
	for i := lastParsedBlock + 1; i <= currentBlock; i++ {
		i := i
		g.Go(func() error {
			return fetchBlock(api, i, blocksChan)
		})
	}
	if err := g.Wait(); err != nil {
		logger.Error(err)
	}
	close(blocksChan)

	blocksList := make([]blockatlas.Block, 0, len(blocksChan))
	for block := range blocksChan {
		blocksList = append(blocksList, block)
	}

	logger.Info("Fetched blocks batch", logger.Params{"from": lastParsedBlock, "to": currentBlock, "total": len(blocksList)})
	return blocksList, nil
}

func fetchBlock(api blockatlas.BlockAPI, num int64, blocksChan chan<- blockatlas.Block) error {
	block, err := getBlockByNumberWithRetry(5, time.Second*5, api.GetBlockByNumber, num)
	if err != nil {
		return errors.E(err, "Fetching failed", logger.Params{"block": num})
	}
	//logger.Info("Fetched", logger.Params{"block": num, "txs_amount": len(block.Txs)})
	blocksChan <- *block
	return nil
}

func SaveLastParsedBlock(storage storage.Tracker, config Params, blocks []blockatlas.Block) error {
	if len(blocks) == 0 {
		return nil
	}

	sort.Slice(blocks, func(i, j int) bool {
		return blocks[i].Number < blocks[j].Number
	})

	lastBlockNumber := blocks[len(blocks)-1].Number

	err := storage.SetLastParsedBlockNumber(config.Coin, lastBlockNumber)
	if err != nil {
		return err
	}

	logger.Info(err, "Save last parsed block", logger.Params{"block": lastBlockNumber, "coin": config.Coin})
	return nil
}

func PublishBlocks(blocks []blockatlas.Block) error {
	if len(blocks) == 0 {
		return nil
	}
	var txsAmount int
	for _, block := range blocks {
		if len(block.Txs) == 0 {
			continue
		}
		txsAmount += len(block.Txs)
		if err := publishBlock(block); err != nil {
			logger.Error(err)
		}
	}
	logger.Info("Published blocks batch", logger.Params{"blocks": len(blocks), "txs": txsAmount})
	logger.Info("------------------------------------------------------------")
	return nil
}

func publishBlock(block blockatlas.Block) error {
	body, err := json.Marshal(block)
	if err != nil {
		return err
	}
	//logger.Info(err, "Published", logger.Params{"block": block.Number})
	return mq.ConfirmedBlocks.Publish(body)
}

func getBlocksIntervalToFetch(api blockatlas.BlockAPI, storage storage.Tracker, config Params) (int64, int64, error) {
	lastParsedBlock, err := storage.GetLastParsedBlockNumber(config.Coin)
	if err != nil {
		return 0, 0, errors.E(err, "Polling failed: tracker didn't return last known block number")
	}
	currentBlock, err := api.CurrentBlockNumber()
	currentBlock -= api.Coin().MinConfirmations
	if err != nil {
		return 0, 0, errors.E(err, "Polling failed: source didn't return chain head number")
	}

	if currentBlock-lastParsedBlock > int64(config.BacklogCount) {
		lastParsedBlock = currentBlock - int64(config.BacklogCount)
	}

	if currentBlock-lastParsedBlock > config.MaxBacklogBlocks {
		lastParsedBlock = currentBlock - config.MaxBacklogBlocks
	}

	return lastParsedBlock, currentBlock, nil
}

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
