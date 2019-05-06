package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"net/http"
	"os"
	"strings"
	"time"
)

var failedFlag = 0

// Entry describes a test
type Entry struct {
	// Coin index
	Index uint
	// Test address
	Addr  string
}

var addresses = map[string]Entry{
	"binance":          {coin.BNB, "tbnb1fhr04azuhcj0dulm7ka40y0cqjlafwae9k9gk2"},
	"nimiq":            {coin.NIM, "NQ86 2H8F YGU5 RM77 QSN9 LYLH C56A CYYR 0MLA"},
	"ripple":           {coin.XRP, "rMQ98K56yXJbDGv49ZSmW51sLn94Xe1mu1"},
	"stellar":          {coin.XLM, "GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX"},
	"kin":              {coin.KIN, "GBHKUZ7C2SZ5N3X2S7O6TT6LNUWSEA2BXMSR5GTTSR6VZARSVAXIQNGH"},
	"tezos":            {coin.XTZ, "tz1WCd2jm4uSt4vntk4vSuUWoZQGhLcDuR9q"},
	"ethereum":         {coin.ETH, "0xfc10cab6a50a1ab10c56983c80cc82afc6559cf1"},
	"classic":          {coin.ETC, "0xf3524415b6D873205B4c3Cda783527b2aC4daAA9"},
	"poa":              {coin.POA, "0x1fddEc96688e0538A316C64dcFd211c491ECf0d8"},
	"callisto":         {coin.CLO, "0x39ec1c88a7a7c1a575e8c8f42eff7630d9278179"},
	"gochain":          {coin.GO,  "0x76c2F81716A8D198a00502Ae9a59126418899FDe"},
	"wanchain":         {coin.WAN, "0x36cEdc3A9d969306AF4F7CA2b83ABBf74095914d"},
	"tomochain":        {coin.TOMO, "0x7daa83030e3086477b79b6e757ca8608899fe783"},
	"aion":             {coin.AION, "0xa07981da70ce919e1db5f051c3c386eb526e6ce8b9e2bfd56e3f3d754b0a17f3"},
	"thundertoken":     {coin.TT,  "0x0ad80a408eac4f17ba0a9de8a12d8736f60700c3"},
	"icon":             {coin.ICX, "hxee691e7bccc4eb11fee922896e9f51490e62b12e"},
	"tron":             {coin.TRX, "TMuA6YqfCeX8EhbfYEg5y7S4DqzSJireY9"},
	"vechain":          {coin.VET, "0xB5e883349e68aB59307d1604555AC890fAC47128"},
	"theta":            {coin.THETA, "0xac0eeb6ee3e32e2c74e14ac74155063e4f4f981f"},
}

func main() {
	if len(os.Args) != 2 {
		logrus.Fatal("Usage: ./test <base_url>")
	}
	baseURL := os.Args[1]

	logrus.SetOutput(os.Stdout)
	http.DefaultClient.Timeout = 5 * time.Second

	supportedEndpoints, err := supportedEndpoints(baseURL)
	if err != nil {
		logrus.WithError(err).Error("Failed to get supported platforms")
		os.Exit(1)
	}

	var supported = make(map[string]bool)
	for _, ns := range supportedEndpoints {
		supported[ns] = true
	}

	for ns, test := range addresses {
		if !supported[ns] {
			log(ns).Warning("Platform not enabled at server, skipping")
		} else {
			runTest(ns, &test, baseURL)
		}
	}

	os.Exit(failedFlag)
}

func log(endpoint string) *logrus.Entry {
	return logrus.WithField("@platform", endpoint)
}

func runTest(endpoint string, entry *Entry, baseURL string) {
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
	test(endpoint, entry, baseURL)
	log(endpoint).Info("Endpoint works")
}

func test(endpoint string, entry *Entry, baseURL string) {
	res, err := http.Get(fmt.Sprintf("%s/v1/%s/%s", baseURL, endpoint, entry.Addr))
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
		Docs []models.Tx `json:"docs"`
	}
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&model)
	if err != nil {
		panic(err)
	}

	if len(model.Docs) == 0 {
		log(endpoint).Warning("No transactions")
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

		if tx.Coin != entry.Index {
			panic("Wrong coin index")
		}
	}

	// Pretty-print first transaction to console
	if len(model.Docs) > 0 {
		pretty, err := json.MarshalIndent(model.Docs[0], "", "\t")
		if err != nil {
			panic("Can't serialize transaction " + err.Error())
		}
		os.Stdout.Write(pretty)
		fmt.Println()
	}
}

func supportedEndpoints(baseURL string) (endpoints []string, err error) {
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
