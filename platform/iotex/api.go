package iotex

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"strconv"
	"time"

	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client = Client{blockatlas.InitClient(viper.GetString("iotex.api"))}
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.IOTX]
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetLatestBlock()
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	var normalized []blockatlas.Tx
	txs, err := p.client.GetTxsInBlock(num)
	if err != nil {
		return nil, err
	}

	for _, action := range txs {
		tx := Normalize(action)
		if tx != nil {
			normalized = append(normalized, *tx)
		}
	}

	return &blockatlas.Block{
		Number: num,
		Txs:    normalized,
	}, nil
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {

	txs := make([]blockatlas.Tx, 0)
	var start int64

	totalTrx, err := p.client.GetAddressTotalTransactions(address)
	if err != nil {
		return nil, err
	}

	if totalTrx >= blockatlas.TxPerPage {
		start = totalTrx - blockatlas.TxPerPage
	}

	actions, err := p.client.GetTxsOfAddress(address, start)
	if err != nil {
		return nil, err
	}

	for _, srcTx := range actions.ActionInfo {
		tx := Normalize(srcTx)
		if tx != nil {
			txs = append(txs, *tx)
		}
	}

	return txs, nil
}

// Normalize converts an Iotex transaction into the generic model
func Normalize(trx *ActionInfo) *blockatlas.Tx {
	if trx.Action == nil {
		return nil
	}
	if trx.Action.Core == nil {
		return nil
	}
	if trx.Action.Core.Transfer == nil {
		return nil
	}

	date, err := time.Parse(time.RFC3339, trx.Timestamp)
	if err != nil {
		return nil
	}
	height, err := strconv.ParseInt(trx.BlkHeight, 10, 64)
	if err != nil {
		return nil
	}
	if height <= 0 {
		return nil
	}
	nonce, err := strconv.ParseInt(trx.Action.Core.Nonce, 10, 64)
	if err != nil {
		return nil
	}

	return &blockatlas.Tx{
		ID:       trx.ActHash,
		Coin:     coin.IOTX,
		From:     trx.Sender,
		To:       trx.Action.Core.Transfer.Recipient,
		Fee:      blockatlas.Amount(trx.GasFee),
		Date:     date.Unix(),
		Block:    uint64(height),
		Status:   blockatlas.StatusCompleted,
		Sequence: uint64(nonce),
		Type:     blockatlas.TxTransfer,
		Meta: blockatlas.Transfer{
			Value:    trx.Action.Core.Transfer.Amount,
			Symbol:   coin.Coins[coin.IOTX].Symbol,
			Decimals: coin.Coins[coin.IOTX].Decimals,
		},
	}
}

func (p *Platform) GetValidators() (blockatlas.ValidatorPage, error) {
	return p.client.GetValidators()
}

func (p *Platform) GetDelegations(address string) (blockatlas.DelegationsPage, error) {
	return p.client.GetDelegations(address)
}

func (p *Platform) GetDetails() blockatlas.StakingDetails {
	return blockatlas.StakingDetails{
		Reward:        blockatlas.StakingReward{Annual: 0},
		MinimumAmount: blockatlas.Amount("100000000000000000000"),
		LockTime:      259200,
		Type:          blockatlas.DelegationTypeDelegate,
	}
}

func (p *Platform) UndelegatedBalance(address string) (string, error) {
	account, err := p.client.GetAccount(address)
	if err != nil {
		return "0", err
	}

	return account.AccountMeta.Balance, nil
}
