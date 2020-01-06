package terra

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client = Client{blockatlas.InitClient(viper.GetString(p.ConfigKey()))}
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.LUNA]
}

func (p *Platform) ConfigKey() string {
	return fmt.Sprintf("%s.api", p.Coin().Handle)
}

// GetBlockByNumber returns txs with block number
func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	srcTxs, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}

	txs := p.NormalizeTxs(srcTxs.Txs)
	return &blockatlas.Block{
		Number: num,
		Txs:    txs,
	}, nil
}

// CurrentBlockNumber returns current block number
func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.CurrentBlockNumber()
}

// GetTxsByAddress returns send/receive txs filtered by address
func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	srcTxs, err := p.client.GetAddrTxs(address)
	if err != nil {
		return nil, err
	}
	return p.NormalizeTxs(srcTxs.Txs), nil
}

// NormalizeTxs converts multiple Cosmos transactions
func (p *Platform) NormalizeTxs(srcTxs []Tx) blockatlas.TxPage {
	txMap := make(map[string]bool)
	txs := make(blockatlas.TxPage, 0)
	for _, srcTx := range srcTxs {
		_, ok := txMap[srcTx.ID]
		if ok {
			continue
		}
		normalisedInputTx, ok := p.Normalize(&srcTx)
		if ok {
			txMap[srcTx.ID] = true
			txs = append(txs, normalisedInputTx)
		}
	}
	return txs
}

// Normalize converts an Cosmos transaction into the generic model
func (p *Platform) Normalize(srcTx *Tx) (tx blockatlas.Tx, ok bool) {
	date, err := time.Parse("2006-01-02T15:04:05Z", srcTx.Date)
	if err != nil {
		return blockatlas.Tx{}, false
	}
	block, err := strconv.ParseUint(srcTx.Block, 10, 64)
	if err != nil {
		return blockatlas.Tx{}, false
	}

	status := blockatlas.StatusCompleted
	// https://github.com/cosmos/cosmos-sdk/blob/95ddc242ad024ca78a359a13122dade6f14fd676/types/errors/errors.go#L19
	if srcTx.Code > 0 {
		status = blockatlas.StatusFailed
	}

	tx = blockatlas.Tx{
		ID:     srcTx.ID,
		Coin:   p.Coin().ID,
		Date:   date.Unix(),
		Status: status,
		Fee:    "0",
		Block:  block,
		Memo:   srcTx.Data.Contents.Memo,
	}

	if len(srcTx.Data.Contents.Message) == 0 {
		return tx, false
	}

	fees := srcTx.Data.Contents.Fee.FeeAmount
	msg := srcTx.Data.Contents.Message[0]
	switch msg.Value.(type) {
	case MessageValueTransfer:
		transfer := msg.Value.(MessageValueTransfer)
		p.fillTransfer(&tx, transfer, fees)
		return tx, true
	case MessageValueDelegate:
		delegate := msg.Value.(MessageValueDelegate)
		p.fillDelegate(&tx, delegate, srcTx.Events, msg.Type, fees)
		return tx, true
	}
	return tx, false
}

func (p *Platform) fillTransfer(tx *blockatlas.Tx, msg MessageValueTransfer, feeAmounts Amounts) {
	if len(msg.Amount) == 0 {
		return
	}

	tx.From = msg.FromAddr
	tx.To = msg.ToAddr
	tx.Type = blockatlas.TxMultiCoinTransfer
	tx.Meta = blockatlas.MultiCoinTransfer{
		Coins: Amounts(msg.Amount).toCoins(),
		Fees:  feeAmounts.toCoins(),
	}

	return
}

func (p *Platform) fillDelegate(tx *blockatlas.Tx, delegate MessageValueDelegate, events Events, msgType TxType, feeAmounts Amounts) {
	coins := Amounts{delegate.Amount}.toCoins()

	tx.From = delegate.DelegatorAddr
	tx.To = delegate.ValidatorAddr
	tx.Type = blockatlas.TxMultiCoinAnyAction

	key := blockatlas.KeyStakeDelegate
	title := blockatlas.KeyTitle("")
	switch msgType {
	case MsgDelegate:
		tx.Direction = blockatlas.DirectionOutgoing
		title = blockatlas.AnyActionDelegation
	case MsgUndelegate:
		tx.Direction = blockatlas.DirectionIncoming
		title = blockatlas.AnyActionUndelegation
	case MsgWithdrawDelegationReward:
		tx.Direction = blockatlas.DirectionIncoming
		title = blockatlas.AnyActionClaimRewards
		key = blockatlas.KeyStakeClaimRewards
		rewards := events.GetWithdrawRewardValue()

		coins = rewards.toCoins()
	}

	var fees []blockatlas.Transfer
	for _, coin := range feeAmounts {
		fees = append(fees, blockatlas.Transfer{
			Symbol:   DenomMap[coin.Denom],
			Decimals: p.Coin().Decimals,
			Value:    blockatlas.Amount(coin.Quantity),
		})
	}

	tx.Meta = blockatlas.MultiCoinAnyAction{
		Title: title,
		Key:   key,
		Coins: coins,
		Fees:  feeAmounts.toCoins(),
	}
}
