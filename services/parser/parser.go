package parser

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync/atomic"

	"math/rand"
	"sort"
	"sync"
	"time"

	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/network/mq"
	"github.com/trustwallet/golibs/numbers"

	log "github.com/sirupsen/logrus"
)

type (
	Params struct {
		Ctx                                       context.Context
		Api                                       blockatlas.BlockAPI
		TransactionsExchange                      mq.Exchange
		ParsingBlocksInterval, FetchBlocksTimeout time.Duration
		BacklogCount                              int
		MaxBacklogBlocks                          int64
		StopChannel                               chan<- struct{}
		Database                                  *db.Instance
	}

	GetBlockByNumber func(num int64) (*blockatlas.Block, error)

	stop struct {
		error
	}
)

func RunParser(params Params) {
	log.Info("------------------------------------------------------------")
	for {
		select {
		case <-params.Ctx.Done():
			log.Info(fmt.Sprintf("Parser of %s stopped parsing blocks", params.Api.Coin().Handle))
			params.StopChannel <- struct{}{}
			return
		default:
			parse(params)
			time.Sleep(params.ParsingBlocksInterval)
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
	lastParsedBlock, currentBlock, err := GetBlocksIntervalToFetch(params)
	if err != nil || lastParsedBlock > currentBlock {
		log.WithFields(log.Fields{"operation": "fetch GetBlocksIntervalToFetch", "lastParsedBlock": lastParsedBlock, "currentBlock": currentBlock, "coin": params.Api.Coin().Handle}).Error(err)
		time.Sleep(params.ParsingBlocksInterval)
		return
	}

	blocks := FetchBlocks(params, lastParsedBlock, currentBlock)

	err = SaveLastParsedBlock(params, blocks)
	if err != nil {
		log.WithFields(log.Fields{"operation": "run SaveLastParsedBlock", "coin": params.Api.Coin().Handle}).Error(err)
		time.Sleep(params.ParsingBlocksInterval)
		return
	}

	var txs []blockatlas.Tx
	for _, block := range blocks {
		txs = append(txs, block.Txs...)
	}
	txs = blockatlas.TxPage(txs).FilterTransactionsByMemo()

	publish(params, txs)

	log.WithFields(log.Fields{
		"coin":         params.Api.Coin().Handle,
		"transactions": len(txs),
	}).Info("Published transactions")

	log.WithFields(log.Fields{"coin": params.Api.Coin().Handle}).Info("End of parse step")
}

func GetBlocksIntervalToFetch(params Params) (int64, int64, error) {
	lastParsedBlock, err := params.Database.GetLastParsedBlockNumber(params.Api.Coin().Handle)
	if err != nil {
		return 0, 0, errors.New(err.Error() + " Polling failed: tracker didn't return last known block number")
	}
	currentBlock, err := params.Api.CurrentBlockNumber()
	currentBlock -= params.Api.Coin().MinConfirmations
	if err != nil {
		return 0, 0, errors.New(err.Error() + "Polling failed: source didn't return chain head number")
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
		log.WithFields(log.Fields{
			"block": lastParsedBlock,
			"coin":  params.Api.Coin().Handle,
		}).Info("No new blocks")
		return nil
	}

	blocksCount := currentBlock - lastParsedBlock
	if blocksCount < 0 {
		log.WithFields(log.Fields{"coin": params.Api.Coin().Handle}).Error("Current block is 0")
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
		log.WithFields(log.Fields{"coin": params.Api.Coin().Handle, "count": len(errorsList), "blocks": errorsList}).Error("Fetch blocks errors")
	}

	blocksList := make([]blockatlas.Block, 0, len(blocksChan))
	for block := range blocksChan {
		blocksList = append(blocksList, block)
	}

	log.WithFields(log.Fields{
		"from":  lastParsedBlock,
		"to":    currentBlock,
		"total": totalCount,
		"coin":  params.Api.Coin().Handle},
	).Info("Fetched blocks batch")

	return blocksList
}

func fetchBlock(api blockatlas.BlockAPI, num int64, blocksChan chan<- blockatlas.Block) error {
	block, err := getBlockByNumberWithRetry(5, time.Second*5, api.GetBlockByNumber, num, api.Coin().Symbol)
	if err != nil {
		return fmt.Errorf("%d", num)
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
	if len(blocks)-1 < 0 {
		return fmt.Errorf("cannot get last block number for %s", params.Api.Coin().Handle)
	}

	lastBlockNumber := blocks[len(blocks)-1].Number
	if lastBlockNumber <= 0 {
		return fmt.Errorf("parser of %s failed to save last block, lastBlockNumber <= 0", params.Api.Coin().Handle)
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

func publish(params Params, txs blockatlas.Txs) {

	if len(txs) == 0 {
		return
	}

	body, err := json.Marshal(txs)
	if err != nil {
		log.WithFields(log.Fields{"operation": "publish marshal", "coin": params.Api.Coin().Handle}).Error(err)
		return
	}

	// Notify transactions queue
	err = params.TransactionsExchange.Publish(body)
	if err != nil {
		log.WithFields(log.Fields{"operation": "publish transactionsQueue", "coin": params.Api.Coin().Handle}).Error(err)
		return
	}
}

func getBlockByNumberWithRetry(attempts int, sleep time.Duration, getBlockByNumber GetBlockByNumber, n int64, symbol string) (*blockatlas.Block, error) {
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
