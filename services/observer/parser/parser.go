package parser

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"sync/atomic"

	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/storage"
	"math/rand"
	"sort"
	"sync"
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

func RunParser(api blockatlas.BlockAPI, storage storage.Tracker, config Params, ctx context.Context) {
	logger.Info("------------------------------------------------------------")
	for {
		select {
		case <-ctx.Done():
			logger.Info(fmt.Sprintf("Parser of %d has been stopped", config.Coin))
			return
		default:
			lastParsedBlock, currentBlock, err := GetBlocksIntervalToFetch(api, storage, config)
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
}

func GetBlocksIntervalToFetch(api blockatlas.BlockAPI, storage storage.Tracker, config Params) (int64, int64, error) {
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

func FetchBlocks(api blockatlas.BlockAPI, lastParsedBlock, currentBlock int64) ([]blockatlas.Block, error) {
	if lastParsedBlock == currentBlock {
		logger.Info("No new blocks", logger.Params{"last": lastParsedBlock, "coin": api.Coin().ID, "time": time.Now().Unix()})
		logger.Info("------------------------------------------------------------")
		return nil, nil
	}

	var (
		blocksCount = currentBlock - lastParsedBlock
		blocksChan  = make(chan blockatlas.Block, blocksCount)
		errorsChan  = make(chan error, blocksCount)
		totalCount  int32
	)
	var wg sync.WaitGroup
	for i := lastParsedBlock + 1; i <= currentBlock; i++ {
		wg.Add(1)
		go func(i int64, wg *sync.WaitGroup) {
			defer wg.Done()
			err := fetchBlock(api, i, blocksChan)
			if err != nil {
				errorsChan <- err
				return
			}
			atomic.AddInt32(&totalCount, 1)
		}(i, &wg)
	}
	wg.Wait()
	close(errorsChan)
	close(blocksChan)

	if len(errorsChan) > 0 {
		var (
			errorsList = make([]error, 0, len(errorsChan))
		)
		for err := range errorsChan {
			errorsList = append(errorsList, err)
		}
		logger.Error("Fetch blocks errors", logger.Params{"count": len(errorsList), "blocks": errorsList})
	}

	blocksList := make([]blockatlas.Block, 0, len(blocksChan))
	for block := range blocksChan {
		blocksList = append(blocksList, block)
	}

	logger.Info("Fetched blocks batch", logger.Params{"from": lastParsedBlock, "to": currentBlock, "total": totalCount})
	return blocksList, nil
}

func fetchBlock(api blockatlas.BlockAPI, num int64, blocksChan chan<- blockatlas.Block) error {
	block, err := getBlockByNumberWithRetry(5, time.Second*5, api.GetBlockByNumber, num)
	if err != nil {
		return errors.E(fmt.Sprintf("%d", num))
	}
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

	var (
		//txsAmount, publishedBlocksCount int32
		txsBatch blockatlas.Txs
		//errorsChan                      = make(chan error, len(blocks))
		//wg                              sync.WaitGroup
	)

	for _, block := range blocks {
		if len(block.Txs) == 0 {
			continue
		}
		for _, tx := range block.Txs {
			txsBatch = append(txsBatch, tx)
		}
		//wg.Add(1)
		//go func(block blockatlas.Block, wg *sync.WaitGroup) {
		//	defer wg.Done()
		//
		//	if len(block.Txs) == 0 {
		//		return
		//	}
		//
		//	atomic.AddInt32(&txsAmount, int32(len(block.Txs)))
		//
		//	err := publishTxsBatch(block)
		//	if err != nil {
		//		errorsChan <- err
		//		return
		//	}
		//
		//	atomic.AddInt32(&publishedBlocksCount, 1)
		//}(block, &wg)
	}
	//wg.Wait()
	//close(errorsChan)

	//if len(errorsChan) > 0 {
	//	var (
	//		errorsList = make([]error, 0, len(errorsChan))
	//	)
	//	for err := range errorsChan {
	//		errorsList = append(errorsList, err)
	//	}
	//	logger.Error("Publish block errors", logger.Params{"errorsCount": len(errorsList), "errorsDetails": errorsList})
	//}

	err := publishTxsBatch(txsBatch)
	if err != nil {
		logger.Error(err)
	}
	logger.Info("Published blocks batch", logger.Params{"blocks": len(blocks), "txs": len(txsBatch)})
	logger.Info("------------------------------------------------------------")
	return nil
}

func publishTxsBatch(txs blockatlas.Txs) error {
	body, err := json.Marshal(txs)
	if err != nil {
		return err
	}
	return mq.ParsedTransactionsBatch.Publish(body)
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
