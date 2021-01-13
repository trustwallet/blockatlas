package filecoin

import (
	"reflect"
	"testing"

	"github.com/trustwallet/blockatlas/platform/filecoin/explorer"
	"github.com/trustwallet/golibs/types"
)

func TestPlatform_NormalizeMessage(t *testing.T) {
	type args struct {
		message explorer.Message
		address string
	}
	tests := []struct {
		name string
		args args
		want types.Tx
	}{
		{
			name: "Test transfer",
			args: args{
				message: explorer.Message{
					Cid:       "bafy2bzacectkidmsel5gn5qamrqhcgqgqefhdzlqukry3rc2ase4yqdbxazqi",
					Height:    305641,
					Timestamp: 1607475630,
					From:      "f1mdseaz4gkbz2cq2kf4q3xqbws5ongyt7vdbvzoa",
					To:        "f16hhfi2xkkmpsi4c5mmgwbkgk5wslfcaflefsqpy",
					Nonce:     6,
					Value:     "298040241318833792",
					Method:    "Send",
					Receipt: explorer.Receipt{
						ExitCode: 0,
					},
				},
				address: "f16hhfi2xkkmpsi4c5mmgwbkgk5wslfcaflefsqpy",
			},
			want: types.Tx{
				ID:        "bafy2bzacectkidmsel5gn5qamrqhcgqgqefhdzlqukry3rc2ase4yqdbxazqi",
				Coin:      461,
				From:      "f1mdseaz4gkbz2cq2kf4q3xqbws5ongyt7vdbvzoa",
				To:        "f16hhfi2xkkmpsi4c5mmgwbkgk5wslfcaflefsqpy",
				Date:      1607475630,
				Block:     305641,
				Status:    "completed",
				Sequence:  6,
				Type:      "transfer",
				Direction: "incoming",
				Meta: types.Transfer{
					Value:    "298040241318833792",
					Symbol:   "FIL",
					Decimals: 18,
				},
			},
		},
		{
			name: "Test transfer failed",
			args: args{
				message: explorer.Message{
					Cid:       "bafy2bzacebgbszawdnrl7flgrb2ixt4f6lsb4jjqywv3bsekplpsiswexepha",
					Height:    310103,
					Timestamp: 1607609490,
					From:      "f16hhfi2xkkmpsi4c5mmgwbkgk5wslfcaflefsqpy",
					To:        "f1elvpho4c6iba6xrlok3773cxbvq32thlofezw5y",
					Nonce:     23,
					Value:     "1000000000000000",
					Method:    "Send",
					Receipt: explorer.Receipt{
						ExitCode: 7,
					},
				},
				address: "f16hhfi2xkkmpsi4c5mmgwbkgk5wslfcaflefsqpy",
			},
			want: types.Tx{
				ID:        "bafy2bzacebgbszawdnrl7flgrb2ixt4f6lsb4jjqywv3bsekplpsiswexepha",
				Coin:      461,
				From:      "f16hhfi2xkkmpsi4c5mmgwbkgk5wslfcaflefsqpy",
				To:        "f1elvpho4c6iba6xrlok3773cxbvq32thlofezw5y",
				Date:      1607609490,
				Block:     310103,
				Status:    "error",
				Sequence:  23,
				Type:      "transfer",
				Direction: "outgoing",
				Meta: types.Transfer{
					Value:    "1000000000000000",
					Symbol:   "FIL",
					Decimals: 18,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Platform{}
			if got := p.NormalizeMessage(tt.args.message, tt.args.address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Platform.NormalizeMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
