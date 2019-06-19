package iotex

import(
	"github.com/trustwallet/blockatlas"
	"net/http"
	"strconv"
	"time"

	"github.com/spf13/viper"

	"github.com/trustwallet/blockatlas/coin"
)

const Handle = "iotex"

var client = Client{
	HTTPClient : http.DefaultClient,
}

type Platform struct {
	client Client
}

func (p *Platform) Handle() string {
	return Handle
}

func (p *Platform) Init() error {
	p.client.BaseURL = viper.GetString("iotex.api")
	p.client.HTTPClient = http.DefaultClient
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.IOTX]
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	var start int64

	totalTrx, err := client.GetAddressTotalTransactions(address)

	if totalTrx >= blockatlas.TxPerPage {
		start = totalTrx - blockatlas.TxPerPage
	}

	trxs, err := client.GetTxsOfAddress(address, start)
	if err != nil {
		return nil, err
	}

	var txs []blockatlas.Tx
	for _, srcTx := range trxs.ActionInfo {
		tx, ok := Normalize(srcTx)
		if !ok || len(txs) >= blockatlas.TxPerPage {
			continue
		}
		txs = append(txs, tx)
	}
	
	return txs, nil
}

// Normalize converts an Iotex transaction into the generic model
func Normalize(trx *ActionInfo) (blockatlas.Tx, bool) {
	date, err := time.Parse(time.RFC3339, trx.Timestamp)
	if err != nil {
		return blockatlas.Tx{
			Coin: coin.IOTX,
			Status: blockatlas.StatusFailed,
			Error: err.Error(),
		}, false
	}
	height, err := strconv.ParseInt(trx.BlkHeight, 10, 64)
	if err != nil {
		return blockatlas.Tx{
			Coin: coin.IOTX,
			Status: blockatlas.StatusFailed,
			Error: err.Error(),
		}, false
	}
	if height <= 0 {
		return blockatlas.Tx{
			Coin: coin.IOTX,
			Status: blockatlas.StatusFailed,
			Error: "invalid block height",
		}, false
	}
	nonce, err := strconv.ParseInt(trx.Action.Core.Nonce, 10, 64)
	if err != nil {
		return blockatlas.Tx{
			Coin: coin.IOTX,
			Status: blockatlas.StatusFailed,
			Error: err.Error(),
		}, false
	}

	return blockatlas.Tx{
		ID       : trx.ActHash,
		Coin     : coin.IOTX,
		From     : trx.Sender,
		To       : trx.Action.Core.Transfer.Recipient,
		Fee      : blockatlas.Amount(trx.GasFee),
		Date     : date.Unix(),
		Block    : uint64(height),
		Status   : blockatlas.StatusCompleted,
		Sequence : uint64(nonce),
		Type     : blockatlas.TxTransfer,
		Meta     : blockatlas.Transfer{
			Value : blockatlas.Amount(trx.Action.Core.Transfer.Amount),
		},
	}, true
}
