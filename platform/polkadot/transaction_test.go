package polkadot

import (
	"reflect"
	"testing"

	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/types"
)

func TestNormalizeTransfer(t *testing.T) {
	platform := Platform{CoinIndex: coin.KUSAMA}
	type args struct {
		srcTx *Transfer
	}
	tests := []struct {
		name   string
		args   args
		wantTx types.Tx
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
			wantTx: types.Tx{
				ID:     "0x20cfbba19817e4b7a61e718d269de47e7067a24860fa978c2a8ead4c96a827c4",
				Coin:   434,
				Date:   1577176992,
				From:   "HKtMPUSoTC8Hts2uqcQVzPAuPRpecBt4XJ5Q1AT1GM3tp2r",
				To:     "CtwdfrhECFs3FpvCGoiE4hwRC4UsSiM8WL899HjRdQbfYZY",
				Block:  360298,
				Status: "completed",
				Fee:    "100000000",
				Meta: types.Transfer{
					Value:    types.Amount("10000000000"),
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
			wantTx: types.Tx{
				ID:     "0xe0be47f6ce0e62a218f197dd68599989376ee7e951d54c3c6146e2c5c5eacd1f",
				Coin:   434,
				Date:   1575525432,
				From:   "DbCNECPna3k6MXFWWNZa5jGsuWycqEE6zcUxZYkxhVofrFk",
				To:     "CtwdfrhECFs3FpvCGoiE4hwRC4UsSiM8WL899HjRdQbfYZY",
				Block:  90672,
				Status: "completed",
				Fee:    "100000000",
				Meta: types.Transfer{
					Value:    types.Amount("210000000000000"),
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
	type args struct {
		platform Platform
		srcTx    *Extrinsic
	}
	tests := []struct {
		name   string
		args   args
		wantTx *types.Tx
	}{
		{
			name: "Transfer KSM",
			args: args{
				platform: Platform{CoinIndex: coin.KUSAMA},
				srcTx: &Extrinsic{
					Timestamp:          1577176992,
					BlockNumber:        360298,
					CallModuleFunction: "transfer",
					CallModule:         "balances",
					Params:             "[{\"name\":\"dest\",\"type\":\"Address\",\"value\":\"0e33fdfb980e4499e5c3576e742a563b6a4fc0f6f598b1917fd7a6fe393ffc72\",\"value_raw\":\"\"},{\"name\":\"value\",\"type\":\"Compact\\u003cBalance\\u003e\",\"value\":\"10000000000\",\"value_raw\":\"\"}]",
					AccountId:          "HKtMPUSoTC8Hts2uqcQVzPAuPRpecBt4XJ5Q1AT1GM3tp2r",
					Nonce:              0,
					Hash:               "0x20cfbba19817e4b7a61e718d269de47e7067a24860fa978c2a8ead4c96a827c4",
					Success:            true,
					Fee:                "100000000",
				},
			},
			wantTx: &types.Tx{
				ID:       "0x20cfbba19817e4b7a61e718d269de47e7067a24860fa978c2a8ead4c96a827c4",
				Coin:     434,
				Date:     1577176992,
				From:     "HKtMPUSoTC8Hts2uqcQVzPAuPRpecBt4XJ5Q1AT1GM3tp2r",
				To:       "CtwdfrhECFs3FpvCGoiE4hwRC4UsSiM8WL899HjRdQbfYZY",
				Fee:      "100000000",
				Block:    360298,
				Status:   "completed",
				Sequence: 0,
				Meta: types.Transfer{
					Value:    types.Amount("10000000000"),
					Symbol:   "KSM",
					Decimals: 12,
				},
			},
		},
		{
			name: "Transfer DOT",
			args: args{
				platform: Platform{CoinIndex: coin.POLKADOT},
				srcTx: &Extrinsic{
					Timestamp:          1607035338,
					BlockNumber:        2742892,
					CallModuleFunction: "transfer",
					CallModule:         "balances",
					Params:             "[{\"name\":\"dest\",\"type\":\"Address\",\"value\":\"deb1bb215d2188d0934e803473f02b61b7990ffaea63a533a9917d5707f74d35\",\"value_raw\":\"\"},{\"name\":\"value\",\"type\":\"Compact\\u003cBalance\\u003e\",\"value\":\"16694000000\",\"value_raw\":\"\"}]",
					AccountId:          "13VELdVkrHf1UUH6TRYSuofJAHykL3MAw3sXG9HGR8YpgrBH",
					Nonce:              1,
					Hash:               "0x17f92e09994e6885007bfdf4a0a5026f667e69ae94547c5c89b03d647541025e",
					Success:            true,
					Fee:                "153000000",
				},
			},
			wantTx: &types.Tx{
				ID:       "0x17f92e09994e6885007bfdf4a0a5026f667e69ae94547c5c89b03d647541025e",
				Coin:     354,
				Date:     1607035338,
				From:     "13VELdVkrHf1UUH6TRYSuofJAHykL3MAw3sXG9HGR8YpgrBH",
				To:       "162zRoDpEXQVCpPgJcoMMKZK6P8M9Ntu4myZ9ZBbDz6mqBzt",
				Fee:      "153000000",
				Block:    2742892,
				Status:   "completed",
				Sequence: 1,
				Meta: types.Transfer{
					Value:    types.Amount("16694000000"),
					Symbol:   "DOT",
					Decimals: 10,
				},
			},
		},
		{
			name: "Bond",
			args: args{
				platform: Platform{CoinIndex: coin.KUSAMA},
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
				platform: Platform{CoinIndex: coin.KUSAMA},
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
				platform: Platform{CoinIndex: coin.KUSAMA},
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
			if gotTx := tt.args.platform.NormalizeExtrinsic(tt.args.srcTx); !reflect.DeepEqual(gotTx, tt.wantTx) {
				t.Errorf("Normalize() = %v\n Want = %v", gotTx, tt.wantTx)
			}
		})
	}
}

func TestNormalizeAddress(t *testing.T) {

	type args struct {
		platform Platform
		value    string
	}
	tests := []struct {
		name        string
		args        args
		wantAddress string
	}{
		{
			name: "KSM address 1",
			args: args{
				platform: Platform{CoinIndex: coin.KUSAMA},
				value:    "e8e1b8de72651640e302b62dad1f643ec8b65a3647a7409b2896634db599ed60",
			},
			wantAddress: "HqfgRXDgCQcV8KAuTAPGuA1r91iEzinmmNBPkR9kiKhifJq",
		},
		{
			name: "KSM address 2",
			args: args{
				platform: Platform{CoinIndex: coin.KUSAMA},
				value:    "e0b3fcccfe0283cc0f8c105c68b5690aab8c5c1692a868e55eaca836c8779085",
			},
			wantAddress: "HewiDTQv92L2bVtkziZC8ASxrFUxr6ajQ62RXAnwQ8FDVmg",
		},
		{
			name: "DOT address",
			args: args{
				platform: Platform{CoinIndex: coin.POLKADOT},
				value:    "e0b3fcccfe0283cc0f8c105c68b5690aab8c5c1692a868e55eaca836c8779085",
			},
			wantAddress: "165dCENc9ZGsiUgxwvxWSKdbfsxtrUqYMWymC9tC1gwGfATj",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if address := tt.args.platform.NormalizeAddress(tt.args.value); address != tt.wantAddress {
				t.Errorf("Normalize() = %v\n Want = %v", address, tt.wantAddress)
			}
		})
	}
}
