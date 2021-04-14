package cosmos

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/mock"
	"github.com/trustwallet/golibs/types"
)

var (
	transferSrc, _       = mock.JsonStringFromFilePath("mocks/" + "transfer.json")
	transferSrcKava, _   = mock.JsonStringFromFilePath("mocks/" + "transfer_kava.json")
	failedTransferSrc, _ = mock.JsonStringFromFilePath("mocks/" + "transfer_failed.json")
	delegateSrc, _       = mock.JsonStringFromFilePath("mocks/" + "delegate_tx.json")
	unDelegateSrc, _     = mock.JsonStringFromFilePath("mocks/" + "undelegate_tx.json")
	claimRewardSrc1, _   = mock.JsonStringFromFilePath("mocks/" + "claim_1.json")
	claimRewardSrc2, _   = mock.JsonStringFromFilePath("mocks/" + "claim_2.json")

	atom = coin.Cosmos()
	kava = coin.Kava()

	transferDst = types.Tx{
		ID:     "E19B011D20D862DA0BEA7F24E3BC6DFF666EE6E044FCD9BD95B073478086DBB6",
		Coin:   atom.ID,
		From:   "cosmos1rw62phusuv9vzraezr55k0vsqssvz6ed52zyrl",
		To:     "cosmos1nynns8ex9fq6sjjfj8k79ymkdz4sqth06xexae",
		Fee:    "1",
		Date:   1556992677,
		Block:  151980,
		Status: types.StatusCompleted,
		Type:   types.TxTransfer,
		Meta: types.Transfer{
			Value:    "2271999999",
			Symbol:   coin.Cosmos().Symbol,
			Decimals: 6,
		},
	}

	transferDstKava = types.Tx{
		ID:     "E19B011D20D862DA0BEA7F24E3BC6DFF666EE6E044FCD9BD95B073478086DBB6",
		Coin:   kava.ID,
		From:   "kava17wcggpjx007uc09s8y4hwrj8f228mlwe0n0upn",
		To:     "kava1z89utvygweg5l56fsk8ak7t6hh88fd0agl98n0",
		Fee:    "1",
		Date:   1556992677,
		Block:  151980,
		Status: types.StatusCompleted,
		Type:   types.TxTransfer,
		Meta: types.Transfer{
			Value:    "2271999999",
			Symbol:   kava.Symbol,
			Decimals: 6,
		},
	}

	delegateDst = types.Tx{
		ID:        "11078091D1D5BD84F4275B6CEE02170428944DB0E8EEC37E980551435F6D04C7",
		Coin:      atom.ID,
		From:      "cosmos1237l0vauhw78qtwq045jd24ay4urpec6r3xfn3",
		To:        "cosmosvaloper12w6tynmjzq4l8zdla3v4x0jt8lt4rcz5gk7zg2",
		Fee:       "5000",
		Date:      1564632616,
		Block:     1258202,
		Status:    types.StatusCompleted,
		Type:      types.TxAnyAction,
		Direction: types.DirectionOutgoing,
		Meta: types.AnyAction{
			Coin:     atom.ID,
			Title:    types.AnyActionDelegation,
			Key:      types.KeyStakeDelegate,
			Name:     atom.Name,
			Symbol:   atom.Symbol,
			Decimals: atom.Decimals,
			Value:    "49920",
		},
	}

	unDelegateDst = types.Tx{
		ID:        "A1EC36741FEF681F4A77B8F6032AD081100EE5ECB4CC76AEAC2174BC6B871CFE",
		Coin:      atom.ID,
		From:      "cosmos137rrp4p8n0nqcft0mwc62tdnyhhzf80knv5t94",
		To:        "cosmosvaloper1te8nxpc2myjfrhaty0dnzdhs5ahdh5agzuym9v",
		Fee:       "5000",
		Date:      1564624521,
		Block:     1257037,
		Status:    types.StatusCompleted,
		Type:      types.TxAnyAction,
		Direction: types.DirectionIncoming,
		Meta: types.AnyAction{
			Coin:     atom.ID,
			Title:    types.AnyActionUndelegation,
			Key:      types.KeyStakeDelegate,
			Name:     atom.Name,
			Symbol:   atom.Symbol,
			Decimals: atom.Decimals,
			Value:    "5100000000",
		},
	}

	claimRewardDst2 = types.Tx{
		ID:        "082BA88EC055A7C343A353297EAC104CE87C659E0DDD84621C9AC3C284232800",
		Coin:      atom.ID,
		From:      "cosmos1y6yvdel7zys8x60gz9067fjpcpygsn62ae9x46",
		To:        "cosmosvaloper12w6tynmjzq4l8zdla3v4x0jt8lt4rcz5gk7zg2",
		Fee:       "0",
		Date:      1576462863,
		Block:     54561,
		Status:    types.StatusCompleted,
		Type:      types.TxAnyAction,
		Direction: types.DirectionIncoming,
		Memo:      "复投",
		Meta: types.AnyAction{
			Coin:     atom.ID,
			Title:    types.AnyActionClaimRewards,
			Key:      types.KeyStakeClaimRewards,
			Name:     atom.Name,
			Symbol:   atom.Symbol,
			Decimals: atom.Decimals,
			Value:    "2692701",
		},
	}

	claimRewardDst1 = types.Tx{
		ID:        "C382DCFDC30E2DA294421DAEAD5862F118592A7B000EE91F6BEF8452A1F525D7",
		Coin:      atom.ID,
		From:      "cosmos1cxehfdhfm96ljpktdxsj0k6xp9gtuheghwgqug",
		To:        "cosmosvaloper1ptyzewnns2kn37ewtmv6ppsvhdnmeapvtfc9y5",
		Fee:       "1000",
		Date:      1576638273,
		Block:     79678,
		Status:    types.StatusCompleted,
		Type:      types.TxAnyAction,
		Direction: types.DirectionIncoming,
		Memo:      "",
		Meta: types.AnyAction{
			Coin:     atom.ID,
			Title:    types.AnyActionClaimRewards,
			Key:      types.KeyStakeClaimRewards,
			Name:     atom.Name,
			Symbol:   atom.Symbol,
			Decimals: atom.Decimals,
			Value:    "86278",
		},
	}

	failedTransferDst = types.Tx{
		ID:     "5E78C65A8C1A6C8239EBBBBF2E42020E6ADBA8037EDEA83BF88E1A9159CF13B8",
		Coin:   atom.ID,
		From:   "cosmos1shpfyt7psrff2ux7nznxvj6f7gq59fcqng5mku",
		To:     "cosmos1za4pu5gxm80fg6sx0956f88l2sx7jfg2vf7nlc",
		Fee:    "2000",
		Date:   1576120902,
		Block:  5552,
		Status: types.StatusError,
		Type:   types.TxTransfer,
		Memo:   "UniCoins registration rewards",
		Meta: types.Transfer{
			Value:    "100000",
			Symbol:   coin.Cosmos().Symbol,
			Decimals: 6,
		},
	}
)

type test struct {
	name     string
	platform Platform
	Data     string
	want     types.Tx
}

func TestNormalize(t *testing.T) {

	cosmos := Platform{CoinIndex: atom.ID}
	kava := Platform{CoinIndex: coin.KAVA}

	tests := []test{
		{
			"test transfer tx",
			cosmos,
			transferSrc,
			transferDst,
		},
		{
			"test delegate tx",
			cosmos,
			delegateSrc,
			delegateDst,
		},
		{
			"test undelegate tx",
			cosmos,
			unDelegateSrc,
			unDelegateDst,
		},
		{
			"test claimReward tx 1",
			cosmos,
			claimRewardSrc1,
			claimRewardDst1,
		},
		{
			"test claimReward tx 2",
			cosmos,
			claimRewardSrc2,
			claimRewardDst2,
		},
		{
			"test failed tx",
			cosmos,
			failedTransferSrc,
			failedTransferDst,
		},
		{
			"test kava transfer tx",
			kava,
			transferSrcKava,
			transferDstKava,
		},
	}
	for _, tt := range tests {
		testNormalize(t, tt)
	}
}

func testNormalize(t *testing.T, tt test) {
	t.Run(tt.name, func(t *testing.T) {
		var srcTx Tx
		err := json.Unmarshal([]byte(tt.Data), &srcTx)
		assert.Nil(t, err)
		tx, ok := tt.platform.Normalize(&srcTx)
		assert.True(t, ok)
		assert.Equal(t, tt.want, tx, "transfer: tx don't equal")
	})
}
