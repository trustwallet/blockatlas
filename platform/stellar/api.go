package stellar

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/numbers"
	"strconv"
	"sync"
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

	return p.NormalizePayments(payments), nil
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.CurrentBlockNumber()
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	if srcBlock, err := p.client.GetBlockByNumber(num); err == nil {
		block := p.NormalizeBlock(srcBlock)
		return &block, nil
	} else {
		return nil, err
	}
}

func (p *Platform) NormalizeBlock(block *Block) blockatlas.Block {
	return blockatlas.Block{
		ID:     block.Ledger.Id,
		Number: block.Ledger.Sequence,
		Txs:    p.NormalizePayments(block.Payments),
	}
}

func (p *Platform) NormalizePayments(payments []Payment) (txs []blockatlas.Tx) {
	var (
		wg      sync.WaitGroup
		txsChan = make(chan blockatlas.Tx, len(payments))
	)

	for _, payment := range payments {
		wg.Add(1)
		go func(pay Payment) {
			defer wg.Done()

			txHash, err := p.client.GetTxHash(pay.TransactionHash)
			if err != nil {
				return
			}

			tx, ok := Normalize(&pay, p.CoinIndex, txHash)
			if !ok {
				return
			}

			txsChan <- tx

		}(payment)
	}
	wg.Wait()
	close(txsChan)

	for tx := range txsChan {
		txs = append(txs, tx)
	}

	return
}

// Normalize converts a Stellar-based transaction into the generic model
func Normalize(payment *Payment, nativeCoinIndex uint, hash TxHash) (tx blockatlas.Tx, ok bool) {
	switch payment.Type {
	case PaymentType:
		if payment.AssetType != Native {
			return tx, false
		}
	case CreateAccount:
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
		value, err = numbers.DecimalToSatoshis(payment.Amount)
		from = payment.From
		to = payment.To
	} else if payment.StartingBalance != "" { // When transfer to new account
		value, err = numbers.DecimalToSatoshis(payment.StartingBalance)
		from = payment.Funder
		to = payment.Account
	} else {
		return tx, false
	}
	if err != nil {
		return tx, false
	}
	return blockatlas.Tx{
		ID:    payment.TransactionHash,
		Coin:  nativeCoinIndex,
		From:  from,
		To:    to,
		Fee:   FixedFee,
		Date:  date.Unix(),
		Memo:  hash.Memo,
		Block: id,
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(value),
			Symbol:   coin.Coins[nativeCoinIndex].Symbol,
			Decimals: coin.Coins[nativeCoinIndex].Decimals,
		},
	}, true
}
