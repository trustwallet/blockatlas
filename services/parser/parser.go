package parser

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"sync/atomic"

	"github.com/trustwallet/blockatlas/db/models"

	"github.com/getsentry/raven-go"

	"math/rand"
	"sort"
	"sync"
	"time"

	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/network/mq"
	"github.com/trustwallet/golibs/numbers"
	"github.com/trustwallet/golibs/types"

	log "github.com/sirupsen/logrus"
)

type (
	Params struct {
		Api                                       blockatlas.BlockAPI
		TransactionsExchange                      mq.Exchange
		ParsingBlocksInterval, FetchBlocksTimeout time.Duration
		MaxBlocks                                 int64
		StopChannel                               chan<- struct{}
		Database                                  *db.Instance
	}

	GetBlockByNumber func(num int64) (*types.Block, error)

	stop struct {
		error
	}
)

func RunParser(params Params, ctx context.Context) {
	log.Info("------------------------------------------------------------")
	for {
		select {
		case <-ctx.Done():
			log.Info(fmt.Sprintf("Parser of %s stopped parsing blocks", params.Api.Coin().Handle))
			params.StopChannel <- struct{}{}
			return
		default:
			parse(params)
		}
		log.Info("------------------------------------------------------------")
	}
}

func GetInterval(value int, minInterval, maxInterval time.Duration) time.Duration {
	interval := time.Duration(value) * time.Millisecond
	pMin := numbers.Max(minInterval.Nanoseconds(), interval.Nanoseconds())
	pMax := numbers.Min(int(maxInterval.Nanoseconds()), int(pMin))
	return time.Duration(pMax)
}

func parse(params Params) {
	coinTracker, err := params.Database.GetLastParsedBlockNumber(params.Api.Coin().Handle)
	if err != nil {
		time.Sleep(params.ParsingBlocksInterval)
		return
	}
	if !coinTracker.Enabled {
		time.Sleep(params.ParsingBlocksInterval)
		return
	}

	lastParsedBlock, currentBlock, err := GetBlocksIntervalToFetch(params, coinTracker)
	if err != nil {
		time.Sleep(params.ParsingBlocksInterval)
		return
	}

	if lastParsedBlock > currentBlock {
		time.Sleep(params.ParsingBlocksInterval)
		return
	}

	blocks, err := FetchBlocks(params, lastParsedBlock, currentBlock)
	if err != nil {
		time.Sleep(params.ParsingBlocksInterval)
		return
	}

	err = SaveLastParsedBlock(params, blocks)
	if err != nil {
		log.WithFields(log.Fields{
			"operation":       "run SaveLastParsedBlock",
			"coin":            params.Api.Coin().Handle,
			"blocks":          blocks,
			"lastParsedBlock": lastParsedBlock,
			"currentBlock":    currentBlock,
			"tags":            raven.Tags{{Key: "coin", Value: params.Api.Coin().Handle}},
		}).Error(err)
		time.Sleep(params.ParsingBlocksInterval)
		return
	}

	var txs types.Txs
	for _, block := range blocks {
		txs = append(txs, block.Txs...)
	}
	txs = txs.FilterTransactionsByMemo()

	err = publish(params, txs)
	if err != nil {
		log.WithFields(log.Fields{
			"coin":         params.Api.Coin().Handle,
			"transactions": len(txs),
			"error":        err,
		}).Info("Publish Error")
	}

	log.WithFields(log.Fields{
		"coin":         params.Api.Coin().Handle,
		"transactions": len(txs),
	}).Info("Published transactions")

	log.WithFields(log.Fields{"coin": params.Api.Coin().Handle}).Info("End of parse step")
}

func GetBlocksIntervalToFetch(params Params, tracker models.Tracker) (int64, int64, error) {
	lastParsedBlock := tracker.Height
	currentBlock, err := params.Api.CurrentBlockNumber()
	if err != nil {
		return 0, 0, errors.New(err.Error() + "Polling failed: source didn't return chain head number. lastParsedBlock: " + strconv.Itoa(int(lastParsedBlock)))
	}
	currentBlock -= params.Api.Coin().MinConfirmations

	return GetNextBlocksToParse(lastParsedBlock, currentBlock, params.MaxBlocks)
}

func GetNextBlocksToParse(lastParsedBlock int64, currentBlock int64, maxBlocks int64) (int64, int64, error) {
	if lastParsedBlock == currentBlock {
		return lastParsedBlock, currentBlock, nil
	}
	if lastParsedBlock > currentBlock {
		return lastParsedBlock, lastParsedBlock, nil
	}

	var endParseBlock = currentBlock
	var nextBlock = lastParsedBlock + 1

	if currentBlock-lastParsedBlock > maxBlocks {
		endParseBlock = nextBlock + maxBlocks
	}

	return nextBlock, endParseBlock + 1, nil
}

func FetchBlocks(params Params, lastParsedBlock, currentBlock int64) ([]types.Block, error) {
	if lastParsedBlock == currentBlock {
		log.WithFields(log.Fields{
			"current_block": lastParsedBlock,
			"coin":          params.Api.Coin().Handle,
		}).Info("No new blocks")
		return nil, errors.New("no new blocks")
	}

	blocksCount := currentBlock - lastParsedBlock
	if blocksCount < 0 {
		log.WithFields(log.Fields{"coin": params.Api.Coin().Handle}).Error("Current block is 0")
		return nil, errors.New("current block is 0")
	}

	var (
		blocksChan = make(chan types.Block, blocksCount)
		errorsChan = make(chan error, blocksCount)
		totalCount int32
		wg         sync.WaitGroup
	)

	for i := lastParsedBlock; i <= currentBlock-1; i++ {
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
		log.WithFields(log.Fields{
			"coin":   params.Api.Coin().Handle,
			"count":  len(errorsList),
			"blocks": errorsList,
			"tags": raven.Tags{
				{Key: "coin", Value: params.Api.Coin().Handle},
			},
		}).Error("Fetch Blocks Errors")

		return []types.Block{}, fmt.Errorf("unable to fetch blocks: %d: %d", lastParsedBlock, currentBlock)
	}

	blocks := make([]types.Block, 0, len(blocksChan))
	for block := range blocksChan {
		blocks = append(blocks, block)
	}

	log.WithFields(log.Fields{
		"from":  lastParsedBlock,
		"to":    currentBlock - 1,
		"total": totalCount,
		"coin":  params.Api.Coin().Handle},
	).Info("Fetched blocks batch")

	return blocks, nil
}

func fetchBlock(api blockatlas.BlockAPI, num int64, blocksChan chan<- types.Block) error {
	block, err := getBlockByNumberWithRetry(5, time.Second*5, api.GetBlockByNumber, num, api.Coin().Symbol)
	if err != nil {
		return fmt.Errorf("%d", num)
	}
	blocksChan <- *block
	return nil
}

func SaveLastParsedBlock(params Params, blocks []types.Block) error {
	if len(blocks) == 0 {
		return nil
	}

	sort.Slice(blocks, func(i, j int) bool {
		return blocks[i].Number < blocks[j].Number
	})
	if len(blocks)-1 < 0 {
		return fmt.Errorf("cannot get last block number for %s", params.Api.Coin().Handle)
	}

	lastBlockNumber := blocks[len(blocks)-1].Number
	if lastBlockNumber <= 0 {
		return fmt.Errorf("parser of %s failed to save last block, lastBlockNumber <= 0: %d", params.Api.Coin().Handle, lastBlockNumber)
	}
	err := params.Database.SetLastParsedBlockNumber(params.Api.Coin().Handle, lastBlockNumber)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"block": lastBlockNumber,
		"coin":  params.Api.Coin().Handle,
	}).Info("Save last parsed block")

	return nil
}

func publish(params Params, transactions types.Txs) error {

	if len(transactions) == 0 {
		return nil
	}

	body, err := json.Marshal(transactions)
	if err != nil {
		log.WithFields(log.Fields{"operation": "publish marshal", "transactions": transactions, "coin": params.Api.Coin().Handle}).Error(err)
		return err
	}
	return params.TransactionsExchange.Publish(body)
}

func getBlockByNumberWithRetry(attempts int, sleep time.Duration, getBlockByNumber GetBlockByNumber, n int64, symbol string) (*types.Block, error) {
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

			log.WithFields(log.Fields{
				"number":   n,
				"attempts": attempts,
				"sleep":    sleep.String(),
				"symbol":   symbol},
			).Warn("retry GetBlockByNumber")

			time.Sleep(sleep)
			return getBlockByNumberWithRetry(attempts, sleep*2, getBlockByNumber, n, symbol)
		}
	}
	return r, err
}
