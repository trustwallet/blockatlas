package binance

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

const (
	transferTransaction = `
{
	"blockHeight": 7761368,
	"code": 0,
	"confirmBlocks": 2089441,
	"fromAddr": "tbnb1fhr04azuhcj0dulm7ka40y0cqjlafwae9k9gk2",
	"hasChildren": 0,
	"log": "Msg 0: ",
	"timeStamp": 1555049867552,
	"toAddr": "tbnb1sylyjw032eajr9cyllp26n04300qzzre38qyv5",
	"txAge": 836729,
	"txAsset": "BNB",
	"txFee": 0.00125,
	"txHash": "1681EE543FB4B5A628EF21D746E031F018E226D127044A4F9BA5EE2542A44555",
	"txType": "TRANSFER",
	"value": 100000,
	"memo": "test"
}`
	tokenTransferTransaction = `
{
	"blockHeight": 7928667,
	"code": 0,
	"confirmBlocks": 1922024,
	"fromAddr": "tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a",
	"hasChildren": 0,
	"log": "Msg 0: ",
	"timeStamp": 1555117625829,
	"toAddr": "tbnb12hlquylu78cjylk5zshxpdj6hf3t0tahwjt3ex",
	"txAge": 768924,
	"txAsset": "YLC-D8B",
	"txFee": 0.00125,
	"txHash": "95CF63FAA27579A9B6AF84EF8B2DFEAC29627479E9C98E7F5AE4535E213FA4C9",
	"txType": "TRANSFER",
	"value": 2.10572645,
	"memo": "test"
}`
	newOrderTransaction = `
{
      "txHash": "B0677F3436C1B1661E94D192B84B98AA42AC2485D9808357796EE501CBF794F7",
      "blockHeight": 10815565,
      "txType": "NEW_ORDER",
      "timeStamp": 1559689901929,
      "fromAddr": "bnb16ya67j7kvw8682kka09qujlw5u7lf4geqef0ku",
      "value": 0.00649878,
      "txAsset": "BNB",
      "txQuoteAsset": "AWC-986",
      "txFee": 0,
      "txAge": 14346340,
      "orderId": "D13BAF4BD6638FA3AAD6EBCA0E4BEEA73DF4D519-30",
      "data": "{\"orderData\":{\"symbol\":\"BNB_AWC-986\",\"orderType\":\"limit\",\"side\":\"buy\",\"price\":0.00324939,\"quantity\":2.00000000,\"timeInForce\":\"GTE\",\"orderId\":\"D13BAF4BD6638FA3AAD6EBCA0E4BEEA73DF4D519-30\"}}",
      "code": 0,
      "log": "Msg 0: ",
      "confirmBlocks": 0,
      "memo": "",
      "source": 0,
      "hasChildren": 0
}`
	cancelOrderTransaction = `
{
      "txHash": "F48DE755170C10F4A4C0E6836A708C33EEF9A7144800F25187D5F2349FD15A34",
      "blockHeight": 10815539,
      "txType": "CANCEL_ORDER",
      "timeStamp": 1559689892180,
      "fromAddr": "bnb16ya67j7kvw8682kka09qujlw5u7lf4geqef0ku",
      "txFee": 0,
      "txAge": 14346349,
      "data": "{\"orderData\":{\"symbol\":\"BNB_GTO-908\",\"orderType\":\"limit\",\"side\":\"buy\",\"price\":0.00104716,\"quantity\":1.00000000,\"timeInForce\":\"GTE\",\"orderId\":\"D13BAF4BD6638FA3AAD6EBCA0E4BEEA73DF4D519-28\"}}",
      "code": 0,
      "log": "Msg 0: ",
      "confirmBlocks": 0,
      "memo": "",
      "source": 0,
      "hasChildren": 0
}`
	multipleTx = `
{
  "txHash": "0C954A46D5AE90EBF9CB7E6F2EAC0E7C3E8DA2DA94B868962164A3AF9D54BEE8",
  "blockHeight": 64374278,
  "txType": "TRANSFER",
  "timeStamp": 1580128370826,
  "txFee": 0.0006,
  "txAge": 222591,
  "code": 0,
  "log": "Msg 0: ",
  "confirmBlocks": 553829,
  "memo": "",
  "source": 0,
  "sequence": 154231,
  "hasChildren": 1,
  "subTxsDto": {
    "totalNum": 2,
    "pageSize": 15,
    "subTxDtoList": [
      {
        "hash": "0C954A46D5AE90EBF9CB7E6F2EAC0E7C3E8DA2DA94B868962164A3AF9D54BEE8",
        "height": 64374278,
        "type": "TRANSFER",
        "value": 3269,
        "asset": "AERGO-46B",
        "fromAddr": "bnb1nm4n03x00gw0x6v784jzryyp6wxnjaxswr3xm8",
        "toAddr": "bnb1eff4hzx4lfsun3px5walchdy4vek4n0njcdzyn",
        "fee": 0.0006
      },
      {
        "hash": "0C954A46D5AE90EBF9CB7E6F2EAC0E7C3E8DA2DA94B868962164A3AF9D54BEE8",
        "height": 64374278,
        "type": "TRANSFER",
        "value": 1,
        "asset": "BNB",
        "fromAddr": "bnb1nm4n03x00gw0x6v784jzryyp6wxnjaxswr3xm8",
        "toAddr": "bnb1eff4hzx4lfsun3px5walchdy4vek4n0njcdzyn",
        "fee": null
      }
    ]
  }
}`
	multipleTwiceTx = `
{
  "txHash": "C29D822EFBC0C91656D1C5870BA55922F3A72A25BC8415B32D1D1AD0C85142F5",
  "blockHeight": 63591484,
  "txType": "TRANSFER",
  "timeStamp": 1580421001269,
  "txFee": 0.0006,
  "txAge": 18773,
  "code": 0,
  "log": "Msg 0: ",
  "confirmBlocks": 38965,
  "memo": "Trust Wallet Redeem",
  "source": 0,
  "sequence": 33,
  "hasChildren": 1,
  "subTxsDto": {
    "totalNum": 2,
    "pageSize": 15,
    "subTxDtoList": [
      {
        "hash": "C29D822EFBC0C91656D1C5870BA55922F3A72A25BC8415B32D1D1AD0C85142F5",
        "height": 63591484,
        "type": "TRANSFER",
        "value": 1e-8,
        "asset": "BNB",
        "fromAddr": "bnb1nm4n03x00gw0x6v784jzryyp6wxnjaxswr3xm8",
        "toAddr": "bnb1eff4hzx4lfsun3px5walchdy4vek4n0njcdzyn",
        "fee": 0.0006
      },
      {
        "hash": "C29D822EFBC0C91656D1C5870BA55922F3A72A25BC8415B32D1D1AD0C85142F5",
        "height": 63591484,
        "type": "TRANSFER",
        "value": 1e-8,
        "asset": "BNB",
        "fromAddr": "bnb1nm4n03x00gw0x6v784jzryyp6wxnjaxswr3xm8",
        "toAddr": "bnb1eff4hzx4lfsun3px5walchdy4vek4n0njcdzyn",
        "fee": null
      }
    ]
  }
}`
)

var (
	transferDst = blockatlas.Tx{
		ID:     "1681EE543FB4B5A628EF21D746E031F018E226D127044A4F9BA5EE2542A44555",
		Coin:   coin.BNB,
		From:   "tbnb1fhr04azuhcj0dulm7ka40y0cqjlafwae9k9gk2",
		To:     "tbnb1sylyjw032eajr9cyllp26n04300qzzre38qyv5",
		Fee:    "125000",
		Date:   1555049867,
		Block:  7761368,
		Status: blockatlas.StatusCompleted,
		Memo:   "test",
		Meta: blockatlas.Transfer{
			Value:    "10000000000000",
			Decimals: 8,
			Symbol:   "BNB",
		},
	}
	tokenTransferDst = blockatlas.Tx{
		ID:     "95CF63FAA27579A9B6AF84EF8B2DFEAC29627479E9C98E7F5AE4535E213FA4C9",
		Coin:   coin.BNB,
		From:   "tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a",
		To:     "tbnb12hlquylu78cjylk5zshxpdj6hf3t0tahwjt3ex",
		Fee:    "125000",
		Date:   1555117625,
		Block:  7928667,
		Status: blockatlas.StatusCompleted,
		Memo:   "test",
		Meta: blockatlas.NativeTokenTransfer{
			TokenID:  "YLC-D8B",
			Symbol:   "YLC",
			Value:    "210572645",
			Decimals: 8,
			From:     "tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a",
			To:       "tbnb12hlquylu78cjylk5zshxpdj6hf3t0tahwjt3ex",
		},
	}
	newOrderTransferDst = blockatlas.Tx{
		ID:     "B0677F3436C1B1661E94D192B84B98AA42AC2485D9808357796EE501CBF794F7",
		Coin:   coin.BNB,
		From:   "bnb16ya67j7kvw8682kka09qujlw5u7lf4geqef0ku",
		Fee:    "0",
		Date:   1559689901,
		Block:  10815565,
		Status: blockatlas.StatusCompleted,
		Meta: blockatlas.AnyAction{
			Coin:     coin.BNB,
			Title:    blockatlas.KeyTitlePlaceOrder,
			Key:      blockatlas.KeyPlaceOrder,
			TokenID:  "BNB_AWC-986",
			Name:     "AWC-986",
			Symbol:   "AWC",
			Value:    "649878",
			Decimals: 8,
		},
	}
	//TODO: temp test dst
	metaFreeNewOrderTransferDst = blockatlas.Tx{
		ID:     "B0677F3436C1B1661E94D192B84B98AA42AC2485D9808357796EE501CBF794F7",
		Coin:   coin.BNB,
		From:   "bnb16ya67j7kvw8682kka09qujlw5u7lf4geqef0ku",
		Fee:    "0",
		Date:   1559689901,
		Block:  10815565,
		Status: blockatlas.StatusCompleted,
		Meta:   nil,
	}
	cancelOrdeTransferDst = blockatlas.Tx{
		ID:     "F48DE755170C10F4A4C0E6836A708C33EEF9A7144800F25187D5F2349FD15A34",
		Coin:   coin.BNB,
		From:   "bnb16ya67j7kvw8682kka09qujlw5u7lf4geqef0ku",
		Fee:    "0",
		Date:   1559689892,
		Block:  10815539,
		Status: blockatlas.StatusCompleted,
		Meta: blockatlas.AnyAction{
			Coin:     coin.BNB,
			Title:    blockatlas.KeyTitleCancelOrder,
			Key:      blockatlas.KeyCancelOrder,
			TokenID:  "BNB_GTO-908",
			Name:     "GTO-908",
			Symbol:   "GTO",
			Value:    "104716",
			Decimals: 8,
		},
	}
	multipleTxDst = blockatlas.Tx{
		ID:     "0C954A46D5AE90EBF9CB7E6F2EAC0E7C3E8DA2DA94B868962164A3AF9D54BEE8",
		Coin:   coin.BNB,
		From:   "bnb1nm4n03x00gw0x6v784jzryyp6wxnjaxswr3xm8",
		To:     "bnb1eff4hzx4lfsun3px5walchdy4vek4n0njcdzyn",
		Fee:    "60000",
		Date:   1580128370,
		Block:  64374278,
		Status: blockatlas.StatusCompleted,
		Meta: blockatlas.NativeTokenTransfer{
			From:     "bnb1nm4n03x00gw0x6v784jzryyp6wxnjaxswr3xm8",
			To:       "bnb1eff4hzx4lfsun3px5walchdy4vek4n0njcdzyn",
			TokenID:  "AERGO-46B",
			Symbol:   "AERGO",
			Value:    "326900000000",
			Decimals: 8,
		},
	}
	multipleTwiceTxDst = blockatlas.Tx{
		ID:     "C29D822EFBC0C91656D1C5870BA55922F3A72A25BC8415B32D1D1AD0C85142F5",
		Coin:   coin.BNB,
		From:   "bnb1nm4n03x00gw0x6v784jzryyp6wxnjaxswr3xm8",
		To:     "bnb1eff4hzx4lfsun3px5walchdy4vek4n0njcdzyn",
		Fee:    "60000",
		Date:   1580421001,
		Block:  63591484,
		Status: blockatlas.StatusCompleted,
		Memo:   "Trust Wallet Redeem",
		Meta: blockatlas.Transfer{
			Value:    "2",
			Decimals: 8,
			Symbol:   "BNB",
		},
	}
	//TODO: temp test dst
	metaFreeCancelOrdeTransferDst = blockatlas.Tx{
		ID:     "F48DE755170C10F4A4C0E6836A708C33EEF9A7144800F25187D5F2349FD15A34",
		Coin:   coin.BNB,
		From:   "bnb16ya67j7kvw8682kka09qujlw5u7lf4geqef0ku",
		Fee:    "0",
		Date:   1559689892,
		Block:  10815539,
		Status: blockatlas.StatusCompleted,
		Meta:   nil,
	}
	baseTransferTx = blockatlas.Tx{
		ID:     "0C954A46D5AE90EBF9CB7E6F2EAC0E7C3E8DA2DA94B868962164A3AF9D54BEE8",
		Coin:   coin.BNB,
		From:   "bnb1nm4n03x00gw0x6v784jzryyp6wxnjaxswr3xm8",
		To:     "bnb1eff4hzx4lfsun3px5walchdy4vek4n0njcdzyn",
		Fee:    "60000",
		Date:   1580128370,
		Block:  63591484,
		Status: blockatlas.StatusCompleted,
		Memo:   "Trust Wallet Redeem",
		Meta: blockatlas.Transfer{
			Value:    "2",
			Decimals: 8,
			Symbol:   "BNB",
		},
	}
)

type testTx struct {
	name        string
	apiResponse string
	expected    blockatlas.Tx
	token       string
	wantError   bool
}

func TestNormalizeTx(t *testing.T) {
	testTxList := []testTx{
		{
			name:        "bnb transfer",
			apiResponse: transferTransaction,
			expected:    transferDst,
			token:       "BNB",
			wantError:   false,
		},
		{
			name:        "native token transfer",
			apiResponse: tokenTransferTransaction,
			expected:    tokenTransferDst,
			token:       "YLC-D8B",
			wantError:   false,
		},
		{
			name:        "multiple addresses token transfer",
			apiResponse: multipleTx,
			expected:    multipleTxDst,
			token:       "AERGO-46B",
			wantError:   false,
		},
		{
			name:        "multiple addresses with two transfers",
			apiResponse: multipleTwiceTx,
			expected:    multipleTwiceTxDst,
			token:       "BNB",
			wantError:   false,
		},
		{
			name:        "new order transfer",
			apiResponse: newOrderTransaction,
			expected:    newOrderTransferDst,
			token:       "AWC-986",
			wantError:   true,
		},
		{
			name:        "cancel order transfer",
			apiResponse: cancelOrderTransaction,
			expected:    cancelOrdeTransferDst,
			token:       "GTO-908",
			wantError:   true,
		},
		{
			name:        "new order transfer",
			apiResponse: newOrderTransaction,
			expected:    metaFreeNewOrderTransferDst,
			token:       "AWC-986",
			wantError:   true,
		},
		{
			name:        "cancel order transfer",
			apiResponse: cancelOrderTransaction,
			expected:    metaFreeCancelOrdeTransferDst,
			token:       "GTO-908",
			wantError:   true,
		},
		{
			name:        "normalize error transfer",
			apiResponse: tokenTransferTransaction,
			token:       "GTO-908",
			wantError:   true,
		},
	}

	for _, testTxInstance := range testTxList {
		t.Run(testTxInstance.name, func(t *testing.T) {
			var srcTx Tx
			err := json.Unmarshal([]byte(testTxInstance.apiResponse), &srcTx)
			assert.Nil(t, err)
			tx, ok := NormalizeTx(srcTx, testTxInstance.token, "")
			if testTxInstance.wantError {
				assert.False(t, ok, "transfer: tx could be normalized")
				return
			}
			assert.True(t, ok, "transfer: tx could not be normalized")
			assert.Equal(t, blockatlas.TxPage{testTxInstance.expected}, tx, "transfer: tx don't equal")
		})
	}
}

type testTxs struct {
	name        string
	apiResponse string
	expected    []blockatlas.Tx
	token       string
}

func convertJsonToArray(jsonString string) string {
	return "[" + jsonString + "]"
}

func TestNormalizeTxs(t *testing.T) {
	testTxsList := []testTxs{
		{
			name:        "bnb transfer",
			apiResponse: convertJsonToArray(transferTransaction),
			expected:    []blockatlas.Tx{transferDst},
			token:       "BNB",
		},
		{
			name:        "native token transfer",
			apiResponse: convertJsonToArray(tokenTransferTransaction),
			expected:    []blockatlas.Tx{tokenTransferDst},
			token:       "YLC-D8B",
		},
		//{
		//	name:        "all transfers",
		//	apiResponse: AllTransfersType,
		//	expected:    []blockatlas.Tx{transferDst, tokenTransferDst, newOrderTransferDst, cancelOrdeTransferDst},
		//	token:       "",
		//},
		//{
		//	name:        "new order transfer",
		//	apiResponse: convertJsonToArray(newOrderTransaction),
		//	expected:    []blockatlas.Tx{newOrderTransferDst},
		//	token:       "AWC-986",
		//},
		//{
		//	name:        "cancel order transfer",
		//	apiResponse: convertJsonToArray(cancelOrderTransaction),
		//	expected:    []blockatlas.Tx{cancelOrdeTransferDst},
		//	token:       "GTO-908",
		//}
	}

	for _, testTxsInstance := range testTxsList {
		t.Run(testTxsInstance.name, func(t *testing.T) {
			var srcTxs []Tx
			err := json.Unmarshal([]byte(testTxsInstance.apiResponse), &srcTxs)
			assert.Nil(t, err)
			txs := NormalizeTxs(srcTxs, testTxsInstance.token, "")
			assert.Equal(t, testTxsInstance.expected, txs, "transfer: tx don't equal")
		})
	}
}

func TestTokenSymbol(t *testing.T) {
	assert.Equal(t, "UGAS", TokenSymbol("UGAS"))
	assert.Equal(t, "UGAS", TokenSymbol("UGAS-B0C"))
}

var (
	metaTx = blockatlas.Transfer{
		Value:    "100000000",
		Symbol:   "BNB",
		Decimals: 8,
	}
	metaTx2 = blockatlas.Transfer{
		Value:    "2",
		Symbol:   "BNB",
		Decimals: 8,
	}
	metaTokenTx = blockatlas.NativeTokenTransfer{
		Value:    "326900000000",
		TokenID:  "AERGO-46B",
		Symbol:   "AERGO",
		From:     "bnb1nm4n03x00gw0x6v784jzryyp6wxnjaxswr3xm8",
		To:       "bnb1eff4hzx4lfsun3px5walchdy4vek4n0njcdzyn",
		Decimals: 8,
	}
)

func Test_normalizeTransfer(t *testing.T) {
	testTx := baseTransferTx
	testTx.Meta = metaTx
	testTx2 := baseTransferTx
	testTx2.Meta = metaTx2
	testTokenTx := baseTransferTx
	testTokenTx.Meta = metaTokenTx
	type args struct {
		tx      blockatlas.Tx
		srcTx   string
		token   string
		address string
	}
	tests := []struct {
		name  string
		args  args
		want  blockatlas.TxPage
		want1 bool
	}{
		{"test multiple tx 1", args{
			tx:      baseTransferTx,
			srcTx:   multipleTx,
			token:   "BNB",
			address: "",
		}, blockatlas.TxPage{testTx}, true},
		{"test multiple tx 2", args{
			tx:      baseTransferTx,
			srcTx:   multipleTwiceTx,
			token:   "BNB",
			address: "",
		}, blockatlas.TxPage{testTx2}, true},
		{"tx multiple token tx", args{
			tx:      baseTransferTx,
			srcTx:   multipleTx,
			token:   "AERGO-46B",
			address: "",
		}, blockatlas.TxPage{testTokenTx}, true},
		{"test multiple tx fail", args{
			tx:      baseTransferTx,
			srcTx:   multipleTwiceTx,
			token:   "AERGO-46B",
			address: "",
		}, blockatlas.TxPage{}, false},
		{"test multiple tx address fail", args{
			tx:      baseTransferTx,
			srcTx:   multipleTwiceTx,
			token:   "AERGO-46B",
			address: "tbnb1qxm48ndhmh7su0r7zgwmwkltuqgly57jdf8yf8",
		}, blockatlas.TxPage{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var srcTx Tx
			err := json.Unmarshal([]byte(tt.args.srcTx), &srcTx)
			assert.Nil(t, err)
			got, got1 := normalizeTransfer(tt.args.tx, srcTx, tt.args.token, tt.args.address)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
