package tron

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	services "github.com/trustwallet/blockatlas/services/assets"
	"strconv"
	"sync"
	"time"
)

type Platform struct {
	client Client
}

const Annual = 4.32

func (p *Platform) Init() error {
	p.client = Client{blockatlas.InitClient(viper.GetString("tron.api"))}
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.TRX]
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	Txs, err := p.client.GetTxsOfAddress(address, "")
	if err != nil && len(Txs) == 0 {
		return nil, err
	}

	var txs []blockatlas.Tx
	for _, srcTx := range Txs {
		tx, ok := Normalize(&srcTx)
		if ok {
			txs = append(txs, tx)
		}
	}

	return txs, nil
}

func (p *Platform) GetTokenTxsByAddress(address, token string) (blockatlas.TxPage, error) {
	tokenTxs, err := p.client.GetTxsOfAddress(address, token)
	if err != nil {
		return nil, err
	}

	if len(tokenTxs) == 0 {
		return nil, err
	}

	var tokenInfo AssetInfo
	info, err := p.client.GetTokenInfo(token)
	if err != nil && len(info.Data) == 0 {
		return nil, err
	}

	tokenInfo = info.Data[0]

	var txs []blockatlas.Tx
	for _, trx := range tokenTxs {
		tx, err := NormalizeTokenTransfer(&trx, tokenInfo)
		if err != nil {
			logger.Error(err)
			continue
		}
		txs = append(txs, tx)
	}

	return txs, nil
}

func NormalizeTokenTransfer(srcTx *Tx, tokenInfo AssetInfo) (tx blockatlas.Tx, e error) {
	if len(srcTx.Data.Contracts) == 0 {
		return tx, errors.E("token transfer without contract", errors.TypePlatformApi,
			errors.Params{"tokenInfo": tokenInfo, "tx": tx}).PushToSentry()
	}
	contract := &srcTx.Data.Contracts[0]

	switch contract.Parameter.(type) {
	case TransferAssetContract:
		transfer := contract.Parameter.(TransferAssetContract)
		from, err := HexToAddress(transfer.Value.OwnerAddress)
		if err != nil {
			return tx, err
		}
		to, err := HexToAddress(transfer.Value.ToAddress)
		if err != nil {
			return tx, err
		}

		return blockatlas.Tx{
			ID:   srcTx.ID,
			Coin: coin.TRX,
			Date: srcTx.BlockTime / 1000,
			Fee:  "0",
			From: from,
			To:   to,
			Meta: blockatlas.TokenTransfer{
				Name:     tokenInfo.Name,
				Symbol:   tokenInfo.Symbol,
				TokenID:  tokenInfo.ID,
				Decimals: tokenInfo.Decimals,
				Value:    transfer.Value.Amount,
				From:     from,
				To:       to,
			},
		}, nil
	default:
		return tx, nil
	}
}

func (p *Platform) GetValidators() (blockatlas.ValidatorPage, error) {
	results := make(blockatlas.ValidatorPage, 0)
	validators, err := p.client.GetValidators()

	if err != nil {
		return results, err
	}

	for _, v := range validators.Witnesses {
		if val, ok := normalizeValidator(v); ok {
			results = append(results, val)
		}
	}

	return results, nil
}

func (p *Platform) GetDetails() blockatlas.StakingDetails {
	return getDetails()
}

func getDetails() blockatlas.StakingDetails {
	return blockatlas.StakingDetails{
		Reward:        blockatlas.StakingReward{Annual: Annual},
		MinimumAmount: blockatlas.Amount("1000000"),
		LockTime:      259200,
		Type:          blockatlas.DelegationTypeDelegate,
	}
}

func normalizeValidator(v Validator) (validator blockatlas.Validator, ok bool) {
	address, err := HexToAddress(v.Address)
	if err != nil {
		return validator, false
	}

	return blockatlas.Validator{
		Status:  true,
		ID:      address,
		Details: getDetails(),
	}, true
}

/// Normalize converts a Tron transaction into the generic model
func Normalize(srcTx *Tx) (tx blockatlas.Tx, ok bool) {
	if len(srcTx.Data.Contracts) < 1 {
		return tx, false
	}

	// TODO Support multiple transfers in a single transaction
	contract := &srcTx.Data.Contracts[0]
	switch contract.Parameter.(type) {
	case TransferContract:
		transfer := contract.Parameter.(TransferContract)
		from, err := HexToAddress(transfer.Value.OwnerAddress)
		if err != nil {
			return tx, false
		}
		to, err := HexToAddress(transfer.Value.ToAddress)
		if err != nil {
			return tx, false
		}

		return blockatlas.Tx{
			ID:   srcTx.ID,
			Coin: coin.TRX,
			Date: srcTx.BlockTime / 1000,
			From: from,
			To:   to,
			Fee:  "0",
			Meta: blockatlas.Transfer{
				Value:    transfer.Value.Amount,
				Symbol:   coin.Coins[coin.TRX].Symbol,
				Decimals: coin.Coins[coin.TRX].Decimals,
			},
		}, true
	default:
		return tx, false
	}
}

func (p *Platform) GetTokenListByAddress(address string) (blockatlas.TokenPage, error) {
	tokens, err := p.client.GetAccount(address)
	if err != nil {
		return nil, err
	}

	tokenPage := make([]blockatlas.Token, 0)
	var tokenIDs []string
	if len(tokens.Data) == 0 {
		return tokenPage, nil
	}

	for _, v := range tokens.Data[0].AssetsV2 {
		tokenIDs = append(tokenIDs, v.Key)
	}
	tokensInfoChan := make(chan *Asset, len(tokenIDs))

	var wg sync.WaitGroup
	wg.Add(len(tokenIDs))
	for _, id := range tokenIDs {
		go func(id string) {
			defer wg.Done()
			info, err := p.client.GetTokenInfo(id)
			if err != nil {
				logger.Error("GetTokenInfo", err)
			}
			tokensInfoChan <- info
		}(id)
	}
	wg.Wait()
	close(tokensInfoChan)

	tokensInfoMap := make(map[string]AssetInfo)
	for info := range tokensInfoChan {
		if len(info.Data) == 0 {
			continue
		}
		tokensInfoMap[info.Data[0].ID] = info.Data[0]
	}

	for _, v := range tokens.Data[0].AssetsV2 {
		tokenPage = append(tokenPage, NormalizeToken(tokensInfoMap[v.Key]))
	}

	return tokenPage, nil
}

func NormalizeToken(info AssetInfo) blockatlas.Token {
	return blockatlas.Token{
		Name:     info.Name,
		Symbol:   info.Symbol,
		TokenID:  info.ID,
		Coin:     coin.TRX,
		Decimals: info.Decimals,
		Type:     blockatlas.TokenTypeTRC10,
	}
}

func (p *Platform) GetDelegations(address string) (blockatlas.DelegationsPage, error) {
	results := make(blockatlas.DelegationsPage, 0)
	votes, err := p.client.GetAccountVotes(address)
	if err != nil {
		return nil, err
	}
	validators, err := services.GetValidatorsMap(p)
	if err != nil {
		return nil, err
	}
	results = append(results, NormalizeDelegations(votes, validators)...)
	return results, nil
}

func (p *Platform) UndelegatedBalance(address string) (string, error) {
	account, err := p.client.GetAccount(address)
	if err != nil {
		return "0", err
	}

	for _, data := range account.Data {
		return strconv.FormatUint(uint64(data.Balance), 10), nil
	}
	return "0", nil
}

func NormalizeDelegations(data *AccountData, validators blockatlas.ValidatorMap) []blockatlas.Delegation {
	results := make([]blockatlas.Delegation, 0)
	for _, v := range data.Votes {
		validator, ok := validators[v.VoteAddress]
		if !ok {
			logger.Error("Validator not found", validator)
			continue
		}
		delegation := blockatlas.Delegation{
			Delegator: validator,
			Value:     strconv.Itoa(v.VoteCount * 1000000),
			Status:    blockatlas.DelegationStatusActive,
		}
		for _, f := range data.Frozen {
			t2 := time.Now().UnixNano() / int64(time.Millisecond)
			if f.ExpireTime > t2 {
				delegation.Status = blockatlas.DelegationStatusPending
			}
		}
		results = append(results, delegation)
	}
	return results
}
