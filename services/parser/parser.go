package parser

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync/atomic"

	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/numbers"
	"go.elastic.co/apm"

	"math/rand"
	"sort"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type (
	Params struct {
		Ctx                                       context.Context
		Api                                       blockatlas.BlockAPI
		TransactionsQueue                         mq.Queue
		TokenTransactionsQueue                    []mq.Queue
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
	tx := apm.DefaultTracer.StartTransaction("parse", "app")
	defer tx.End()

	ctx := apm.ContextWithTransaction(context.Background(), tx)

	lastParsedBlock, currentBlock, err := GetBlocksIntervalToFetch(params, ctx)
	if err != nil || lastParsedBlock > currentBlock {
		log.WithFields(log.Fields{"coin": params.Api.Coin().Handle}).Error(err)
		time.Sleep(params.ParsingBlocksInterval)
		return
	}

	blocks := FetchBlocks(params, lastParsedBlock, currentBlock, ctx)

	err = SaveLastParsedBlock(params, blocks, ctx)
	if err != nil {
		log.WithFields(log.Fields{"coin": params.Api.Coin().Handle}).Error(err)
		time.Sleep(params.ParsingBlocksInterval)
		return
	}

	txs := ConvertToBatch(blocks, ctx)
	txs = blockatlas.Txs(blockatlas.TxPage(txs).FilterTransactionsByMemo())

	PublishTransactionsBatch(params, txs, ctx)

	log.WithFields(log.Fields{"coin": params.Api.Coin().Handle}).Info("End of parse step")
}

func GetBlocksIntervalToFetch(params Params, ctx context.Context) (int64, int64, error) {
	span, ctx := apm.StartSpan(ctx, "GetBlocksIntervalToFetch", "app")
	defer span.End()

	lastParsedBlock, err := params.Database.GetLastParsedBlockNumber(params.Api.Coin().Handle, ctx)
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

func FetchBlocks(params Params, lastParsedBlock, currentBlock int64, ctx context.Context) []blockatlas.Block {
	span, ctx := apm.StartSpan(ctx, "FetchBlocks", "app")
	defer span.End()

	if lastParsedBlock == currentBlock {
		log.WithFields(log.Fields{"last": lastParsedBlock, "coin": params.Api.Coin().Handle, "time": time.Now().Unix()}).
			Info("No new blocks")
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
			err := fetchBlock(params.Api, i, blocksChan, ctx)
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
		log.WithFields(log.Fields{"count": len(errorsList), "blocks": errorsList}).Error("Fetch blocks errors")
	}

	blocksList := make([]blockatlas.Block, 0, len(blocksChan))
	for block := range blocksChan {
		blocksList = append(blocksList, block)
	}

	log.WithFields(log.Fields{"from": lastParsedBlock, "to": currentBlock, "total": totalCount, "coin": params.Api.Coin().Handle}).
		Info("Fetched blocks batch")
	return blocksList
}

func fetchBlock(api blockatlas.BlockAPI, num int64, blocksChan chan<- blockatlas.Block, ctx context.Context) error {
	span, ctx := apm.StartSpan(ctx, "fetchBlock", "app")
	defer span.End()
	block, err := getBlockByNumberWithRetry(5, time.Second*5, api.GetBlockByNumber, num, api.Coin().Symbol, ctx)
	if err != nil {
		return fmt.Errorf("%d", num)
	}
	blocksChan <- *block
	return nil
}

func SaveLastParsedBlock(params Params, blocks []blockatlas.Block, ctx context.Context) error {
	span, ctx := apm.StartSpan(ctx, "SaveLastParsedBlock", "app")
	defer span.End()

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
		return fmt.Errorf("parser of %s failed to save last block, lastBlockNumber <= 0",
			params.Api.Coin().Handle)
	}
	err := params.Database.SetLastParsedBlockNumber(params.Api.Coin().Handle, lastBlockNumber, ctx)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{"block": lastBlockNumber, "coin": params.Api.Coin().Handle}).
		Info("Save last parsed block")
	return nil
}

func ConvertToBatch(blocks []blockatlas.Block, ctx context.Context) blockatlas.Txs {
	span, ctx := apm.StartSpan(ctx, "ConvertToBatch", "app")
	defer span.End()
	if len(blocks) == 0 {
		return nil
	}

	var txs []blockatlas.Tx

	for _, block := range blocks {
		txs = append(txs, block.Txs...)
	}

	return txs
}

func PublishTransactionsBatch(params Params, txs blockatlas.Txs, ctx context.Context) {
	span, ctx := apm.StartSpan(ctx, "PublishTransactionsBatch", "app")
	defer span.End()

	if len(txs) == 0 {
		return
	}

	batches := getTxsBatches(txs, params.TxBatchLimit, ctx)

	for _, batch := range batches {
		publish(params, batch, ctx)
	}

	log.WithFields(log.Fields{"txs": len(txs), "batchCount": len(batches)}).Info("Published transactions batch")
}

func getTxsBatches(txs blockatlas.Txs, sizeUint uint, ctx context.Context) []blockatlas.Txs {
	span, _ := apm.StartSpan(ctx, "getTxsBatches", "app")
	defer span.End()

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

func publish(params Params, txs blockatlas.Txs, ctx context.Context) {
	span, _ := apm.StartSpan(ctx, "publish", "app")
	defer span.End()

	body, err := json.Marshal(txs)
	if err != nil {
		log.WithFields(log.Fields{"coin": params.Api.Coin().Handle}).Error(err)
		return
	}

	// Notify transactions queue
	err = params.TransactionsQueue.Publish(body)
	if err != nil {
		log.WithFields(log.Fields{"coin": params.Api.Coin().Handle}).Error(err)
		return
	}

	// Notify token transfers queue, if conforms to TokensAPI protocol
	tokenTransfers := txs.FilterTransactionsByType([]blockatlas.TransactionType{
		blockatlas.TxTokenTransfer,
		blockatlas.TxNativeTokenTransfer,
	})

	if len(tokenTransfers) == 0 {
		return
	}

	tokenTransfersBody, err := json.Marshal(tokenTransfers)
	if err != nil {
		log.WithFields(log.Fields{"coin": params.Api.Coin().Handle}).Error(err)
		return
	}

	if _, ok := params.Api.(blockatlas.TokensAPI); ok {
		for _, q := range params.TokenTransactionsQueue {
			err = q.Publish(tokenTransfersBody)
			if err != nil {
				log.WithFields(log.Fields{"coin": params.Api.Coin().Handle}).Error(err)
				return
			}
		}
	}
}

func getBlockByNumberWithRetry(attempts int, sleep time.Duration, getBlockByNumber GetBlockByNumber, n int64, symbol string, ctx context.Context) (*blockatlas.Block, error) {
	span, ctx := apm.StartSpan(ctx, "getBlockByNumberWithRetry", "app")
	defer span.End()
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

			log.WithFields(log.Fields{"number": n, "attempts": attempts, "sleep": sleep.String(), "symbol": symbol}).
				Info("retry GetBlockByNumber")

			time.Sleep(sleep)
			return getBlockByNumberWithRetry(attempts, sleep*2, getBlockByNumber, n, symbol, ctx)
		}
	}
	return r, err
}
