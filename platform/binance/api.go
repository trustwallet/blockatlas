package binance

import (
	"fmt"
	"github.com/trustwallet/blockatlas"
	"net/http"

	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/util"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client.BaseURL = viper.GetString("binance.api")
	p.client.RPCBaseURL = viper.GetString("binance.rpc")
	p.client.HTTPClient = http.DefaultClient
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.BNB]
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	// No native function to get height in explorer API
	// Workaround: Request list of blocks
	// and return number of the newest one
	list, err := p.client.GetBlockList(1)
	if err != nil {
		return 0, err
	}
	if len(list.BlockArray) == 0 {
		return 0, fmt.Errorf("no block descriptor found")
	}
	return list.BlockArray[0].BlockHeight, nil
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	srcTxs, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}
	// TODO: Only returns BNB transactions for now
	txs := NormalizeTxs(srcTxs.Txs, "", "", p)
	return &blockatlas.Block{
		Number: num,
		Txs:    txs,
	}, nil
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	// Endpoint supports queries without token query parameter
	return p.GetTokenTxsByAddress(address, "")
}

func (p *Platform) GetTokenTxsByAddress(address, token string) (blockatlas.TxPage, error) {
	srcTxs, err := p.client.GetTxsOfAddress(address, token)
	if err != nil {
		return nil, err
	}
	return NormalizeTxs(srcTxs.Txs, token, address, p), nil
}

// NormalizeTx converts a Binance transaction into the generic model
func NormalizeTx(srcTx *Tx, token, address string, p *Platform) (tx blockatlas.Tx, ok bool) {
	value := util.DecimalExp(string(srcTx.Value), 8)
	fee := util.DecimalExp(string(srcTx.Fee), 8)
	hash := srcTx.Hash
	asset := srcTx.Asset

	tx = blockatlas.Tx{
		ID:    hash,
		Coin:  coin.BNB,
		Date:  srcTx.Timestamp / 1000,
		Fee:   blockatlas.Amount(fee),
		Block: srcTx.BlockHeight,
		Memo:  srcTx.Memo,
	}

	// Condition for native transfer (BNB)
	if asset == "BNB" && srcTx.Type == TRANSFER && token == "" {
		tx.From = srcTx.FromAddr
		tx.To = srcTx.ToAddr
		tx.Meta = blockatlas.Transfer{
			Value: blockatlas.Amount(value),
		}

		return tx, true
	}

	// Condition for native token transfer
	if asset == token && srcTx.Type == TRANSFER && srcTx.FromAddr != "" {
		tx.From = srcTx.FromAddr
		tx.To = srcTx.ToAddr
		tx.Meta = blockatlas.NativeTokenTransfer{
			TokenID:  asset,
			Symbol:   srcTx.MappedAsset,
			Value:    blockatlas.Amount(value),
			Decimals: 8,
			From:     srcTx.FromAddr,
			To:       srcTx.ToAddr,
		}

		return tx, true
	}

	// Conditin where explorer does not return sender/recepient for multisend transaction
	if (srcTx.FromAddr == "" || srcTx.ToAddr == "") && srcTx.Type == TRANSFER {
		println("hash :", hash)
		receipt, _ := p.client.GetTransactionReceipt(hash)
		zeroMsgValue := receipt.TxReceipts.Value.Msg[0].MsgValue
		zeroInput := zeroMsgValue.Inputs[0]
		outputs := zeroMsgValue.Outputs
		zeroOutputAdress := outputs[0].Address

		// Condition for native transfer in multisend transaction
		if zeroInput.Coins[0].Denom == "BNB" {
			if zeroInput.Address == address {
				tx.From = address
				tx.To = zeroOutputAdress  // Pick 0 index as main receipient
				tx.Meta = blockatlas.Transfer{
					Value: blockatlas.Amount(zeroInput.Coins[0].Amount),
				}
				return tx, true
			}

			coin := getCoinOutput(outputs, address)
			tx.To = address
			tx.From = zeroOutputAdress
			tx.Meta = blockatlas.Transfer{
				Value: blockatlas.Amount(coin.Amount),
			}
			return tx, true
		}

		// Condition for token_transfer in multisend transaction
		if zeroInput.Coins[0].Denom != "BNB" {
			if zeroInput.Address == address {
				tx.From = address
				tx.To = zeroOutputAdress  // Pick 0 index as main receipient
				tx.Meta = blockatlas.TokenTransfer{
					Name: "", // TODO: Replace with actual name
					Symbol: zeroInput.Coins[0].Denom,
					TokenID: "", // TODO: Replace with actual tokenID
					Decimals: 8,
					From: address,
					To: zeroOutputAdress,
					Value: blockatlas.Amount(zeroInput.Coins[0].Amount),
				}
				return tx, true
			}

			coin := getCoinOutput(outputs, address)
			tx.From = zeroOutputAdress
			tx.To = address
			tx.Meta = blockatlas.TokenTransfer{
				Name: "", // TODO: Replace with actual name
				Symbol: coin.Denom,
				TokenID: "", // TODO: Replace with actual tokenID
				Decimals: 8,
				From: zeroOutputAdress,
				To: address,
				Value: blockatlas.Amount(coin.Amount),

			}

			return tx, true
		}
	}
	

	return tx, false
}

// NormalizeTxs converts multiple Binance transactions
func NormalizeTxs(srcTxs []Tx, token, address string, p *Platform) (txs []blockatlas.Tx) {
	for _, srcTx := range srcTxs {
		tx, ok := NormalizeTx(&srcTx, token, address, p)
		if !ok || len(txs) >= blockatlas.TxPerPage {
			continue
		}
		txs = append(txs, tx)
	}
	return
}

func getCoinOutput(outputs []Output, address string) Coin {
	var coin Coin
	for _, out := range outputs {
		if out.Address == address {
			coin = out.Coins[0]
			continue
		}
	}

	return coin
}
