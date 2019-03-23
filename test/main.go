package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/models"
	"net/http"
	"os"
	"strings"
	"time"
)

var failedFlag = 0

var addresses = map[string]string {
	"binance": "tbnb12hlquylu78cjylk5zshxpdj6hf3t0tahwjt3ex",
	"nimiq":   "NQ86 2H8F YGU5 RM77 QSN9 LYLH C56A CYYR 0MLA",
	"ripple":  "rMQ98K56yXJbDGv49ZSmW51sLn94Xe1mu1",
	"stellar": "GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX",
}

func main() {
	if len(os.Args) != 2 {
		logrus.Fatal("Usage: ./test <base_url>")
	}
	b := os.Args[1]

	logrus.SetOutput(os.Stdout)
	http.DefaultClient.Timeout = 5 * time.Second

	for ns, test := range addresses {
		runTest(ns, test, b)
	}

	os.Exit(failedFlag)
}

func log(endpoint string) *logrus.Entry {
	return logrus.WithField("@platform", endpoint)
}

func runTest(endpoint string, address string, baseUrl string) {
	start := time.Now()

	defer func() {
		if r := recover(); r != nil {
			log(endpoint).
				WithField("error", r).
				Error("Endpoint failed")
			failedFlag = 1
		}

		log(endpoint).WithField("time", time.Since(start)).Info("Endpoint tested")
	}()

	log(endpoint).Info("Testing endpoint")
	test(endpoint, address, baseUrl)
	log(endpoint).Info("Endpoint works")
}

func test(endpoint string, address string, baseUrl string) {
	res, err := http.Get(fmt.Sprintf("%s/v1/%s/%s", baseUrl, endpoint, address))
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
	var model []models.Tx
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&model)
	if err != nil {
		panic(err)
	}

	if len(model) == 0 {
		log(endpoint).Warning("No transactions")
		return
	}

	// Enumerate transactions
	var lastTime = ^uint64(0)
	for _, tx := range model {
		point := tx.Date

		if uint64(point) <= lastTime {
			lastTime = uint64(point)
		} else {
			panic("Transactions not in chronological order")
		}
	}

	// Pretty-print first transaction to console
	if len(model) > 0 {
		pretty, err := json.MarshalIndent(model[0], "", "\t")
		if err != nil {
			panic("Can't serialize transaction " + err.Error())
		}
		os.Stdout.Write(pretty)
		fmt.Println()
	}
}
