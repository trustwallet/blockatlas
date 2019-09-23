package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/util"
	"net/http"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/coin"
)

var failedFlag int32 = 0
var baseURL string
var requireAll bool
var concurrency int

var app = cobra.Command{
	Use:   "test <base_url>",
	Short: "Test a live API",
	Long:  "Test a live API by requesting the sample addresses found in coin list",
	Args:  cobra.ExactArgs(1),
	Run:   run,
}

func init() {
	flags := app.Flags()
	flags.BoolVarP(&requireAll, "all", "a", false, "Don't skip platforms not supported server-side")
	flags.IntVarP(&concurrency, "concurrency", "c", 8, "Tests to run at once")
}

func main() {
	err := app.Execute()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(_ *cobra.Command, args []string) {
	baseURL = args[0]

	logrus.SetOutput(os.Stdout)
	http.DefaultClient.Timeout = 5 * time.Second

	supportedEndpoints, err := supportedEndpoints()
	if err != nil {
		logger.Error(err, "Failed to get supported platforms")
		os.Exit(1)
	}

	var supported = make(map[string]bool)
	for _, ns := range supportedEndpoints {
		supported[ns] = true
	}

	logger.Info("Running goroutines tests", logger.Params{"goroutines": concurrency})

	var wg sync.WaitGroup
	sem := util.NewSemaphore(concurrency)

	var tests []coin.Coin

	for _, c := range coin.Coins {
		if !supported[c.Handle] {
			if requireAll {
				log(&c).Error("Platform not enabled at server but required")
				atomic.StoreInt32(&failedFlag, 1)
			} else {
				log(&c).Warning("Platform not enabled at server, skipping")
			}
			continue
		}
		tests = append(tests, c)
	}
	logger.Info("Platforms to test", logger.Params{"count": len(supportedEndpoints)})

	wg.Add(len(tests))
	for _, c := range tests {
		go runTest(c, sem, &wg)
	}

	wg.Wait()

	failed := atomic.LoadInt32(&failedFlag)
	if failed == 1 {
		logger.Fatal("Test failed")
	} else {
		logger.Info("Test passed")
	}
}

func log(c *coin.Coin) *logrus.Entry {
	return logrus.WithField("@platform", c.Handle)
}

func runTest(c coin.Coin, sem *util.Semaphore, wg *sync.WaitGroup) {
	defer wg.Done()
	sem.Acquire()
	defer sem.Release()

	start := time.Now()

	defer func() {
		if r := recover(); r != nil {
			log(&c).
				WithField("error", r).
				Error("Endpoint failed")
			atomic.StoreInt32(&failedFlag, 1)
		}

		log(&c).WithField("time", time.Since(start)).Info("Endpoint tested")
	}()

	test(&c)
	log(&c).Info("Endpoint works")
}

func test(c *coin.Coin) {
	res, err := http.Get(fmt.Sprintf("%s/v1/%s/%s", baseURL, c.Handle, c.SampleAddr))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		panic("Status " + res.Status)
	}

	if !strings.HasPrefix(res.Header.Get("Content-Type"), "application/json") {
		panic("Unexpected Content-Type " + res.Header.Get("Content-Type"))
	}

	// Parse model and read into buffer
	var model struct {
		Docs []blockatlas.Tx `json:"docs"`
	}
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&model)
	if err != nil {
		panic(err)
	}

	if len(model.Docs) == 0 {
		log(c).Warning("No transactions")
		return
	}

	// Enumerate transactions
	var lastTime = ^uint64(0)
	for _, tx := range model.Docs {
		point := tx.Date

		if uint64(point) <= lastTime {
			lastTime = uint64(point)
		} else {
			panic("Transactions not in chronological order")
		}

		if tx.Coin != c.ID {
			panic("Wrong coin index")
		}
	}
}

func supportedEndpoints() (endpoints []string, err error) {
	var data struct {
		Endpoints []string `json:"endpoints"`
	}
	res, err := http.Get(fmt.Sprintf("%s/v1/", baseURL))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&data)
	if err != nil {
		return nil, err
	}
	return data.Endpoints, nil
}
