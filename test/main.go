package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/models"
	"io"
	"net/http"
	"os"
	"strconv"
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

	var buffer bytes.Buffer

	// Parse model and read into buffer
	var model models.Response
	dec := json.NewDecoder(io.TeeReader(res.Body, &buffer))
	err = dec.Decode(&model)
	if err != nil {
		panic(err)
	}

	if len(model.Docs) != model.Total {
		panic("Mismatch len(docs) vs total")
	}

	if len(model.Docs) == 0 {
		log(endpoint).Warning("No transactions")
		return
	}

	// Enumerate transactions
	var lastTime = ^uint64(0)
	for _, tx := range model.Docs {
		point, err := strconv.ParseUint(tx.Timestamp, 10, 64)
		if err != nil {
			panic(err)
		}

		if point <= lastTime {
			lastTime = point
		} else {
			panic("Transactions not in chronological order")
		}
	}

	// Pretty-print to console
	var msg struct {
		Docs []json.RawMessage
	}
	json.Unmarshal(buffer.Bytes(), &msg)
	pretty, _ := json.MarshalIndent(msg.Docs[0], "", "\t")
	os.Stdout.Write(pretty)
	fmt.Println()
}
