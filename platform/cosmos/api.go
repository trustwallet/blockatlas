package cosmos

import (
	"github.com/trustwallet/blockatlas"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/util"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client.BaseURL = viper.GetString("cosmos.api")
	p.client.HTTPClient = http.DefaultClient
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.ATOM]
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	srcTxs, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}

	txs := NormalizeTxs(srcTxs, "", len(srcTxs))
	return &blockatlas.Block{
		Number: num,
		Txs:    txs,
	}, nil
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.CurrentBlockNumber()
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	inputTxes, _ := p.client.GetAddrTxes(address, "inputs")
	outputTxes, _ := p.client.GetAddrTxes(address, "outputs")

	normalisedTxes := make([]blockatlas.Tx, 0)

	for _, inputTx := range inputTxes {
		normalisedInputTx, ok := Normalize(&inputTx)
		if ok {
			normalisedTxes = append(normalisedTxes, normalisedInputTx)
		}
	}
	for _, outputTx := range outputTxes {
		normalisedOutputTx, ok := Normalize(&outputTx)
		if ok {
			normalisedTxes = append(normalisedTxes, normalisedOutputTx)
		}
	}

	sort.Slice(normalisedTxes, func(i, j int) bool {
		return normalisedTxes[i].Date > normalisedTxes[j].Date
	})

	return normalisedTxes, nil
}

func (p *Platform) GetValidators() (blockatlas.ValidatorPage, error) {
	results := make([]blockatlas.StakeValidator, 0)
	validators, _ := p.client.GetValidators()

	for _, validator := range validators {
		results = append(results, normalizeValidator(validator))
	}

	return results, nil
}

func NormalizeTxs(srcTxs []Tx, token string, pageSize int) (txs []blockatlas.Tx) {
	for _, srcTx := range srcTxs {
		tx, ok := Normalize(&srcTx)
		if !ok || len(txs) >= pageSize {
			continue
		}
		txs = append(txs, tx)
	}
	return
}

// Normalize converts an Cosmos transaction into the generic model
func Normalize(srcTx *Tx) (tx blockatlas.Tx, ok bool) {
	date, _ := time.Parse("2006-01-02T15:04:05Z", srcTx.Date)
	block, _ := strconv.ParseUint(srcTx.Block, 10, 64)
	// Sometimes fees can be null objects (in the case of no fees e.g. F044F91441C460EDCD90E0063A65356676B7B20684D94C731CF4FAB204035B41)
	var fee string
	if len(srcTx.Data.Contents.Fee.FeeAmount) == 0 {
		fee = "0"
	} else {
		fee, _ = util.DecimalToSatoshis(srcTx.Data.Contents.Fee.FeeAmount[0].Quantity)
	}

	tx = blockatlas.Tx{
		ID:    srcTx.ID,
		Coin:  coin.ATOM,
		Date:  date.Unix(),
		Fee:   blockatlas.Amount(fee),
		Block: block,
		Memo:  srcTx.Data.Contents.Memo,
	}

	if len(srcTx.Data.Contents.Message) > 0 {
		msg := srcTx.Data.Contents.Message[0]
		switch msg.Value.(type) {
		case MessageValueTransfer:
			transfer := msg.Value.(MessageValueTransfer)
			value, _ := util.DecimalToSatoshis(transfer.Amount[0].Quantity)

			tx.From = transfer.FromAddr
			tx.To = transfer.ToAddr

			tx.Meta = blockatlas.Transfer{
				Value:    blockatlas.Amount(value),
				Symbol:   coin.Coins[coin.ATOM].Symbol,
				Decimals: coin.Coins[coin.ATOM].Decimals,
			}
			return tx, true
		case MessageValueDelegate:
			delegate := msg.Value.(MessageValueDelegate)
			value, _ := util.DecimalToSatoshis(delegate.Amount.Quantity)

			tx.From = delegate.DelegatorAddr
			tx.To = delegate.ValidatorAddr

			title := blockatlas.AnyActionDelegation
			if msg.Type == CosmosMsgUndelegate {
				title = blockatlas.AnyActionUndelegation
			}
			tx.Meta = blockatlas.AnyAction{
				Coin:     coin.ATOM,
				Title:    title,
				Key:      blockatlas.KeyStakeDelegate,
				Name:     "",
				Symbol:   coin.Coins[coin.ATOM].Symbol,
				Decimals: coin.Coins[coin.ATOM].Decimals,
				Value:    blockatlas.Amount(value),
			}
			return tx, true
		}
	}

	return tx, false
}

func normalizeValidator(v CosmosValidator) (validator blockatlas.StakeValidator) {
	return blockatlas.StakeValidator{
		Info: blockatlas.StakeValidatorInfo{
			Website:     v.Description.Website,
			Name:        v.Description.Moniker,
			Description: v.Description.Description,
		},
		Status:    bool(v.Status == 2),
		Address:   v.Operator_Address,
		PublicKey: v.Consensus_Pubkey,
	}
}
