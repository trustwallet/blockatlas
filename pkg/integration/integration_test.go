// +build integration

package integration

import (
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
	"github.com/trustwallet/blockatlas/pkg/integration/coins"
	"github.com/trustwallet/blockatlas/pkg/integration/config"
	"github.com/trustwallet/blockatlas/pkg/integration/tester"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestApis(t *testing.T) {
	config.InitConfig()
	tests, err := tester.GetTests()
	if err != nil {
		t.Error(err)
		return
	}
	coins, err := coins.GetCoins()
	if err != nil {
		t.Error(err)
		return
	}
	for _, test := range tests {
		doTests(t, test, coins)
	}
}

func doTests(t *testing.T, tests []tester.HttpTest, coins []coins.Coin) {
	c := tester.NewClient(t)
	for _, test := range tests {
		var results = make([]tester.HttpResult, 0)
		var wg sync.WaitGroup
		wg.Add(len(coins))
		for _, coin := range coins {
			switch strings.ToUpper(test.Method) {
			case "GET":
				go func(handle, addr string, t tester.HttpTest) {
					defer wg.Done()
					r := c.TestGet(handle, addr, t)
					results = append(results, r)
				}(coin.Handle, coin.SampleAddr, test)
			case "POST":
				go func(handle, addr string, t tester.HttpTest) {
					defer wg.Done()
					r := c.TestPost(handle, addr, t)
					results = append(results, r)
				}(coin.Handle, coin.SampleAddr, test)
			default:
				t.Error(fmt.Sprintf("Unrecognized method: %s", test.Method))
				wg.Done()
			}
		}
		wg.Wait()
		renderTable(results)
	}
}

func renderTable(results []tester.HttpResult) {
	var data []table.Row
	for _, r := range results {
		data = append(data, table.Row{
			r.Version, r.Coin, r.Method, r.Path, r.Status, r.Elapsed.String(),
		})
	}

	statusTransformer := text.Transformer(func(val interface{}) string {
		if v, ok := val.(int); ok && v == 200 {
			return text.FgGreen.Sprint(val)
		}
		return text.FgRed.Sprint(val)
	})

	t := table.NewWriter()
	t.SetColumnConfigs([]table.ColumnConfig{
		{
			Name:        "Status",
			Transformer: statusTransformer,
		},
	})
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Version", "Coin", "Method", "Path", "Status", "Time"})
	t.AppendRows(data)
	t.Render()
}
