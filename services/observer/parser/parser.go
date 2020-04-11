package parser

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"sync/atomic"

	"github.com/trustwallet/blockatlas/pkg/logger"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type (
	Params struct {
		Ctx                                       context.Context
		Api                                       blockatlas.BlockAPI
		Queue                                     mq.Queue
		ParsingBlocksInterval, FetchBlocksTimeout time.Duration
		BacklogCount                              int
		MaxBacklogBlocks                          int64
		StopChannel                               chan<- struct{}
		TxBatchLimit                              uint
		Database                                  *db.Instance
	}

	GetBlockByNumber func(num int64) (*blockatlas.Block, error)

	stop struct {
		error
	}

	transactionsBatch struct {
		sync.Mutex
		blockatlas.Txs
	}
)

const MinTxsBatchLimit = 500

func RunParser(params Params) {
	logger.Info("------------------------------------------------------------")
	for {
		select {
		case <-params.Ctx.Done():
			logger.Info(fmt.Sprintf("Parser of %s stopped parsing blocks", params.Api.Coin().Handle))
			params.StopChannel <- struct{}{}
			return
		default:
			lastParsedBlock, currentBlock, err := GetBlocksIntervalToFetch(params)
			if err != nil || lastParsedBlock > currentBlock {
				logger.Error(err, logger.Params{"coin": params.Api.Coin().Handle})
				time.Sleep(params.ParsingBlocksInterval)
				continue
			}

			blocks := FetchBlocks(params, lastParsedBlock, currentBlock)

			err = SaveLastParsedBlock(params, blocks)
			if err != nil {
				logger.Error(err, logger.Params{"coin": params.Api.Coin().Handle})
				time.Sleep(params.ParsingBlocksInterval)
				continue
			}

			txs := ConvertToBatch(blocks)

			PublishTransactionsBatch(params, txs)

			time.Sleep(params.ParsingBlocksInterval)
		}
	}
}

func GetBlocksIntervalToFetch(params Params) (int64, int64, error) {
	lastParsedBlock, err := params.Database.GetLastParsedBlockNumber(params.Api.Coin().Handle)
	if err != nil {
		return 0, 0, errors.E(err, "Polling failed: tracker didn't return last known block number")
	}
	currentBlock, err := params.Api.CurrentBlockNumber()
	currentBlock -= params.Api.Coin().MinConfirmations
	if err != nil {
		return 0, 0, errors.E(err, "Polling failed: source didn't return chain head number")
	}

	if currentBlock-lastParsedBlock > int64(params.BacklogCount) {
		lastParsedBlock = currentBlock - int64(params.BacklogCount)
	}

	if currentBlock-lastParsedBlock > params.MaxBacklogBlocks {
		lastParsedBlock = currentBlock - params.MaxBacklogBlocks
	}

	return lastParsedBlock, currentBlock, nil
}

func FetchBlocks(params Params, lastParsedBlock, currentBlock int64) []blockatlas.Block {
	if lastParsedBlock == currentBlock {
		logger.Info("No new blocks", logger.Params{"last": lastParsedBlock, "coin": params.Api.Coin().ID, "time": time.Now().Unix()})
		return nil
	}

	blocksCount := currentBlock - lastParsedBlock
	if blocksCount < 0 {
		logger.Error("Current block is 0", logger.Params{"coin": params.Api.Coin().Handle})
		return nil
	}

	var (
		blocksChan = make(chan blockatlas.Block, blocksCount)
		errorsChan = make(chan error, blocksCount)
		totalCount int32
		wg         sync.WaitGroup
	)

	for i := lastParsedBlock + 1; i <= currentBlock; i++ {
		wg.Add(1)
		time.Sleep(params.FetchBlocksTimeout)
		go func(i int64, wg *sync.WaitGroup) {
			defer wg.Done()
			err := fetchBlock(params.Api, i, blocksChan)
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
	return blocksList
}

func fetchBlock(api blockatlas.BlockAPI, num int64, blocksChan chan<- blockatlas.Block) error {
	block, err := getBlockByNumberWithRetry(5, time.Second*5, api.GetBlockByNumber, num)
	if err != nil {
		return errors.E(fmt.Sprintf("%d", num))
	}
	blocksChan <- *block
	return nil
}

func SaveLastParsedBlock(params Params, blocks []blockatlas.Block) error {
	if len(blocks) == 0 {
		return nil
	}

	sort.Slice(blocks, func(i, j int) bool {
		return blocks[i].Number < blocks[j].Number
	})

	lastBlockNumber := blocks[len(blocks)-1].Number
	if lastBlockNumber <= 0 {
		return errors.E(fmt.Sprintf("Parser of %s failed to save last block, lastBlockNumber <= 0", params.Api.Coin().Handle))
	}
	err := params.Database.SetLastParsedBlockNumber(params.Api.Coin().Handle, lastBlockNumber)
	if err != nil {
		return err
	}

	logger.Info(err, "Save last parsed block", logger.Params{"block": lastBlockNumber, "coin": params.Api.Coin().Handle})
	return nil
}

func ConvertToBatch(blocks []blockatlas.Block) blockatlas.Txs {
	if len(blocks) == 0 {
		return nil
	}

	var (
		txsBatch transactionsBatch
		wg       sync.WaitGroup
	)

	for _, block := range blocks {
		wg.Add(1)
		go func(block blockatlas.Block, wg *sync.WaitGroup) {
			defer wg.Done()
			txsBatch.fillBatch(block.Txs)
		}(block, &wg)
	}
	wg.Wait()

	if len(txsBatch.Txs) == 0 {
		logger.Info("Blocks converted to transactions batch, there is no transactions", logger.Params{"blocks": len(blocks)})
		return nil
	}

	logger.Info("Blocks converted to transactions batch", logger.Params{"blocks": len(blocks), "txs": len(txsBatch.Txs)})
	return txsBatch.Txs
}

func PublishTransactionsBatch(params Params, txs blockatlas.Txs) {
	if len(txs) == 0 {
		logger.Info("------------------------------------------------------------")
		return
	}

	batches := getTxsBatches(txs, params.TxBatchLimit)

	var wg sync.WaitGroup
	for _, batch := range batches {
		wg.Add(1)
		go publish(params, batch, &wg)
	}
	wg.Wait()

	logger.Info("Published transactions batch", logger.Params{"txs": len(txs), "batchCount": len(batches)})
	logger.Info("------------------------------------------------------------")
}

func getTxsBatches(txs blockatlas.Txs, sizeUint uint) []blockatlas.Txs {
	size := int(sizeUint)
	resultLength := (len(txs) + size - 1) / size
	result := make([]blockatlas.Txs, resultLength)
	lo, hi := 0, size
	for i := range result {
		if hi > len(txs) {
			hi = len(txs)
		}
		result[i] = txs[lo:hi:hi]
		lo, hi = hi, hi+size
	}
	return result
}

func publish(params Params, txs blockatlas.Txs, wg *sync.WaitGroup) {
	defer wg.Done()
	body, err := json.Marshal(txs)
	if err != nil {
		logger.Error(err, logger.Params{"coin": params.Api.Coin().Handle})
		return
	}
	err = params.Queue.Publish(body)
	if err != nil {
		logger.Error(err, logger.Params{"coin": params.Api.Coin().Handle})
		return
	}
}

func getBlockByNumberWithRetry(attempts int, sleep time.Duration, getBlockByNumber GetBlockByNumber, n int64) (*blockatlas.Block, error) {
	r, err := getBlockByNumber(n)
	if err != nil {
		if s, ok := err.(stop); ok {
			return nil, s.error
		}
		if attempts--; attempts > 0 {
			// Add some randomness to prevent creating a Thundering Herd
			rand.Seed(time.Now().UnixNano())
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

func (t *transactionsBatch) fillBatch(transactions blockatlas.Txs) {
	t.Lock()
	defer t.Unlock()
	if len(transactions) == 0 {
		return
	}
	t.Txs = append(t.Txs, transactions...)
}
