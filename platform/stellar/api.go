package stellar

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/util"
	"strconv"
	"time"
)

type Platform struct {
	client    Client
	CoinIndex uint
}

func (p *Platform) Init() error {
	handle := coin.Coins[p.CoinIndex].Handle
	api := fmt.Sprintf("%s.api", handle)
	p.client = Client{blockatlas.InitClient(viper.GetString(api))}
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[p.CoinIndex]
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	payments, err := p.client.GetTxsOfAddress(address)
	if err != nil {
		return nil, err
	}
	return NormalizePayments(payments, p.CoinIndex), nil
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.CurrentBlockNumber()
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	if srcBlock, err := p.client.GetBlockByNumber(num); err == nil {
		block := NormalizeBlock(srcBlock, p.CoinIndex)
		return &block, nil
	} else {
		return nil, err
	}
}

func NormalizeBlock(block *Block, nativeCoinIndex uint) blockatlas.Block {
	return blockatlas.Block{
		ID:     block.Ledger.Id,
		Number: block.Ledger.Sequence,
		Txs:    NormalizePayments(block.Payments, nativeCoinIndex),
	}
}

func NormalizePayments(payments []Payment, nativeCoinIndex uint) (txs []blockatlas.Tx) {
	for _, payment := range payments {
		tx, ok := Normalize(&payment, nativeCoinIndex)
		if !ok {
			continue
		}
		txs = append(txs, tx)
	}
	return txs
}

// Normalize converts a Stellar-based transaction into the generic model
func Normalize(payment *Payment, nativeCoinIndex uint) (tx blockatlas.Tx, ok bool) {
	switch payment.Type {
	case "payment":
		if payment.AssetType != "native" {
			return tx, false
		}
	case "create_account":
		break
	default:
		return tx, false
	}
	id, err := strconv.ParseUint(payment.ID, 10, 64)
	if err != nil {
		return tx, false
	}
	date, err := time.Parse("2006-01-02T15:04:05Z", payment.CreatedAt)
	if err != nil {
		return tx, false
	}
	var value, from, to string
	if payment.Amount != "" {
		value, err = util.DecimalToSatoshis(payment.Amount)
		from = payment.From
		to = payment.To
	} else if payment.StartingBalance != "" {
		value, err = util.DecimalToSatoshis(payment.StartingBalance)
		from = payment.Funder
		to = payment.Account
	} else {
		return tx, false
	}
	if err != nil {
		return tx, false
	}
	return blockatlas.Tx{
		ID:   payment.TransactionHash,
		Coin: nativeCoinIndex,
		From: from,
		To:   to,
		// https://www.stellar.org/developers/guides/concepts/fees.html
		// Fee fixed at 100 stroops
		Fee:   "100",
		Date:  date.Unix(),
		Block: id,
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(value),
			Symbol:   coin.Coins[nativeCoinIndex].Symbol,
			Decimals: coin.Coins[nativeCoinIndex].Decimals,
		},
	}, true
}
