package polkadot

import (
	"reflect"
	"testing"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/coin"
)

func TestNormalizeTransfer(t *testing.T) {
	platform := Platform{CoinIndex: coin.KSM}
	type args struct {
		srcTx *Transfer
	}
	tests := []struct {
		name   string
		args   args
		wantTx blockatlas.Tx
	}{
		{
			name: "Receive 1",
			args: args{
				srcTx: &Transfer{
					From:        "HKtMPUSoTC8Hts2uqcQVzPAuPRpecBt4XJ5Q1AT1GM3tp2r",
					To:          "CtwdfrhECFs3FpvCGoiE4hwRC4UsSiM8WL899HjRdQbfYZY",
					Amount:      "0.01",
					Module:      "balances",
					Hash:        "0x20cfbba19817e4b7a61e718d269de47e7067a24860fa978c2a8ead4c96a827c4",
					Timestamp:   1577176992,
					BlockNumber: 360298,
					Success:     true,
				},
			},
			wantTx: blockatlas.Tx{
				ID:     "0x20cfbba19817e4b7a61e718d269de47e7067a24860fa978c2a8ead4c96a827c4",
				Coin:   434,
				Date:   1577176992,
				From:   "HKtMPUSoTC8Hts2uqcQVzPAuPRpecBt4XJ5Q1AT1GM3tp2r",
				To:     "CtwdfrhECFs3FpvCGoiE4hwRC4UsSiM8WL899HjRdQbfYZY",
				Block:  360298,
				Status: "completed",
				Fee:    "100000000",
				Meta: blockatlas.Transfer{
					Value:    blockatlas.Amount("10000000000"),
					Symbol:   "KSM",
					Decimals: 12,
				},
			},
		},
		{
			name: "Receive 2",
			args: args{
				srcTx: &Transfer{
					From:        "DbCNECPna3k6MXFWWNZa5jGsuWycqEE6zcUxZYkxhVofrFk",
					To:          "CtwdfrhECFs3FpvCGoiE4hwRC4UsSiM8WL899HjRdQbfYZY",
					Amount:      "210",
					Module:      "balances",
					Hash:        "0xe0be47f6ce0e62a218f197dd68599989376ee7e951d54c3c6146e2c5c5eacd1f",
					Timestamp:   1575525432,
					BlockNumber: 90672,
					Success:     true,
				},
			},
			wantTx: blockatlas.Tx{
				ID:     "0xe0be47f6ce0e62a218f197dd68599989376ee7e951d54c3c6146e2c5c5eacd1f",
				Coin:   434,
				Date:   1575525432,
				From:   "DbCNECPna3k6MXFWWNZa5jGsuWycqEE6zcUxZYkxhVofrFk",
				To:     "CtwdfrhECFs3FpvCGoiE4hwRC4UsSiM8WL899HjRdQbfYZY",
				Block:  90672,
				Status: "completed",
				Fee:    "100000000",
				Meta: blockatlas.Transfer{
					Value:    blockatlas.Amount("210000000000000"),
					Symbol:   "KSM",
					Decimals: 12,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTx := platform.NormalizeTransfer(tt.args.srcTx); !reflect.DeepEqual(gotTx, tt.wantTx) {
				t.Errorf("Normalize() = %v\n Want = %v", gotTx, tt.wantTx)
			}
		})
	}
}

func TestNormalizeExtrinsic(t *testing.T) {
	platform := Platform{CoinIndex: coin.KSM}
	type args struct {
		srcTx *Extrinsic
	}
	tests := []struct {
		name   string
		args   args
		wantTx *blockatlas.Tx
	}{
		{
			name: "Transfer",
			args: args{
				srcTx: &Extrinsic{
					Timestamp:          1577176992,
					BlockNumber:        360298,
					CallModuleFunction: "transfer",
					CallModule:         "balances",
					Params:             "[{\"name\":\"dest\",\"type\":\"Address\",\"value\":\"CtwdfrhECFs3FpvCGoiE4hwRC4UsSiM8WL899HjRdQbfYZY\",\"valueRaw\":\"ff0e33fdfb980e4499e5c3576e742a563b6a4fc0f6f598b1917fd7a6fe393ffc72\"},{\"name\":\"value\",\"type\":\"Compact\\u003cBalance\\u003e\",\"value\":10000000000,\"valueRaw\":\"0700e40b5402\"}]",
					AccountId:          "HKtMPUSoTC8Hts2uqcQVzPAuPRpecBt4XJ5Q1AT1GM3tp2r",
					Nonce:              0,
					Hash:               "0x20cfbba19817e4b7a61e718d269de47e7067a24860fa978c2a8ead4c96a827c4",
					Success:            true,
				},
			},
			wantTx: &blockatlas.Tx{
				ID:     "0x20cfbba19817e4b7a61e718d269de47e7067a24860fa978c2a8ead4c96a827c4",
				Coin:   434,
				Date:   1577176992,
				From:   "HKtMPUSoTC8Hts2uqcQVzPAuPRpecBt4XJ5Q1AT1GM3tp2r",
				To:     "CtwdfrhECFs3FpvCGoiE4hwRC4UsSiM8WL899HjRdQbfYZY",
				Block:  360298,
				Status: "completed",
				Fee:    "100000000",
				Meta: blockatlas.Transfer{
					Value:    blockatlas.Amount("10000000000"),
					Symbol:   "KSM",
					Decimals: 12,
				},
			},
		},
		{
			name: "Bond",
			args: args{
				srcTx: &Extrinsic{
					Timestamp:          1577712822,
					BlockNumber:        447444,
					CallModuleFunction: "bond",
					CallModule:         "staking",
					Params:             "[{\"name\":\"controller\",\"type\":\"Address\",\"value\":\"b44024b9ac73ae8e2f6f6f72b5021a41963b2bc06f67181a040c40bcafb4127b\",\"valueRaw\":\"ffb44024b9ac73ae8e2f6f6f72b5021a41963b2bc06f67181a040c40bcafb4127b\"},{\"name\":\"value\",\"type\":\"Compact\\u003cBalanceOf\\u003e\",\"value\":100000000000,\"valueRaw\":\"0700e8764817\"},{\"name\":\"payee\",\"type\":\"RewardDestination\",\"value\":\"Staked\",\"valueRaw\":\"00\"}]",
					AccountId:          "b44024b9ac73ae8e2f6f6f72b5021a41963b2bc06f67181a040c40bcafb4127b",
					Nonce:              1,
					Hash:               "0xeaaa9ca1a93854be0d3cccc7d7a36272e5663a40a296ab9b0451e0d43ee376ce",
					Success:            true,
				},
			},
			wantTx: nil,
		},
		{
			name: "Error Params",
			args: args{
				srcTx: &Extrinsic{
					Timestamp:          1577712822,
					BlockNumber:        447444,
					CallModuleFunction: "set",
					CallModule:         "timestamp",
					Params:             "[{\"name\":\"now\",\"type\":\"Compact\\u003cMoment\\u003e\",\"value\":1580348178,\"valueRaw\":\"0b507e17f46f01\"}]",
					AccountId:          "b44024b9ac73ae8e2f6f6f72b5021a41963b2bc06f67181a040c40bcafb4127b",
					Nonce:              1,
					Hash:               "0xeaaa9ca1a93854be0d3cccc7d7a36272e5663a40a296ab9b0451e0d43ee376ce",
					Success:            true,
				},
			},
			wantTx: nil,
		},
		{
			name: "set_heads",
			args: args{
				srcTx: &Extrinsic{
					Timestamp:          1577712822,
					BlockNumber:        447444,
					CallModuleFunction: "set_heads",
					CallModule:         "parachains",
					Params:             "[{\"name\":\"heads\",\"type\":\"Vec\\u003cAttestedCandidate\\u003e\",\"value\":[],\"valueRaw\":\"\"}]",
					AccountId:          "b44024b9ac73ae8e2f6f6f72b5021a41963b2bc06f67181a040c40bcafb4127b",
					Nonce:              1,
					Hash:               "0xeaaa9ca1a93854be0d3cccc7d7a36272e5663a40a296ab9b0451e0d43ee376ce",
					Success:            true,
				},
			},
			wantTx: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTx := platform.NormalizeExtrinsic(tt.args.srcTx); !reflect.DeepEqual(gotTx, tt.wantTx) {
				t.Errorf("Normalize() = %v\n Want = %v", gotTx, tt.wantTx)
			}
		})
	}
}

func TestNormalizeAddress(t *testing.T) {
	platform := Platform{CoinIndex: coin.KSM}
	type args struct {
		valueRaw string
	}
	tests := []struct {
		name        string
		args        args
		wantAddress string
	}{
		{
			name: "KSM address 1",
			args: args{
				valueRaw: "ffe8e1b8de72651640e302b62dad1f643ec8b65a3647a7409b2896634db599ed60",
			},
			wantAddress: "HqfgRXDgCQcV8KAuTAPGuA1r91iEzinmmNBPkR9kiKhifJq",
		},
		{
			name: "KSM address 2",
			args: args{
				valueRaw: "ffe0b3fcccfe0283cc0f8c105c68b5690aab8c5c1692a868e55eaca836c8779085",
			},
			wantAddress: "HewiDTQv92L2bVtkziZC8ASxrFUxr6ajQ62RXAnwQ8FDVmg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if address := platform.NormalizeAddress(tt.args.valueRaw); address != tt.wantAddress {
				t.Errorf("Normalize() = %v\n Want = %v", address, tt.wantAddress)
			}
		})
	}
}
