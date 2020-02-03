package ontology

import (
	"reflect"
	"testing"
)

func TestTransfer_isFeeTransfer(t *testing.T) {
	type fields struct {
		ToAddress string
		AssetName AssetType
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"test fee transfer", fields{ToAddress: GovernanceContract, AssetName: AssetONG}, true},
		{"test non fee transfer 1", fields{ToAddress: GovernanceContract, AssetName: AssetONT}, false},
		{"test non fee transfer 2", fields{ToAddress: "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7", AssetName: AssetONG}, false},
		{"test non fee transfer 3", fields{ToAddress: "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7", AssetName: AssetONT}, false},
		{"test invalid asset", fields{ToAddress: GovernanceContract, AssetName: "BNB"}, false},
		{"test empty", fields{ToAddress: "", AssetName: ""}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tf := &Transfer{
				ToAddress: tt.fields.ToAddress,
				AssetName: tt.fields.AssetName,
			}
			if got := tf.isFeeTransfer(); got != tt.want {
				t.Errorf("isFeeTransfer() = %v, want %v", got, tt.want)
			}
		})
	}
}

var (
	transferFee = Transfer{
		Amount:      "0.01",
		FromAddress: "ARFXGXSmgFT2h9EiS4D5fen127Lzi48Eij",
		ToAddress:   "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK",
		AssetName:   "ong",
	}
	transferOng = Transfer{
		Amount:      "0.03534404",
		FromAddress: "AFmseVrdL9f9oyCzZefL9tG6UbvhUMqNMV",
		ToAddress:   "ARFXGXSmgFT2h9EiS4D5fen127Lzi48Eij",
		AssetName:   "ong",
	}
	transferOnt = Transfer{
		Amount:      "58",
		FromAddress: "ARFXGXSmgFT2h9EiS4D5fen127Lzi48Eij",
		ToAddress:   "ARncJn1rr9hivokUWxzr915vS3usR6xdgJ",
		AssetName:   "ont",
	}
	transfersClaims = Transfers{transferOng, transferFee}
	transfersOnt    = Transfers{transferOnt, transferFee}
	transfersOng    = Transfers{transferOng}
	transfersFee    = Transfers{transferFee}
)

func TestTransfers_getTransfer(t *testing.T) {
	tests := []struct {
		name  string
		tfs   Transfers
		asset AssetType
		want  *Transfer
	}{
		{"Transfer Claims Asset Ong", transfersClaims, AssetONG, &transferOng},
		{"Transfer Claims Asset Ont", transfersClaims, AssetONT, nil},
		{"Transfer Claims Asset All", transfersClaims, AssetAll, &transferOng},
		{"Transfer Ont Asset Ong", transfersOnt, AssetONG, nil},
		{"Transfer Ont Asset Ont", transfersOnt, AssetONT, &transferOnt},
		{"Transfer Ont Asset All", transfersOnt, AssetAll, &transferOnt},
		{"Transfer Ong Asset Ong", transfersOng, AssetONG, &transferOng},
		{"Transfer Ong Asset Ont", transfersOng, AssetONT, nil},
		{"Transfer Ong Asset All", transfersOng, AssetAll, &transferOng},
		{"Transfer Fee Asset Ong", transfersFee, AssetONG, nil},
		{"Transfer Fee Asset Ont", transfersFee, AssetONT, nil},
		{"Transfer Fee Asset All", transfersFee, AssetAll, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tfs.getTransfer(tt.asset); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTransfer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransfers_hasFeeTransfer(t *testing.T) {
	tests := []struct {
		name string
		tfs  Transfers
		want bool
	}{
		{"Transfer Claims", transfersClaims, true},
		{"Transfer Ont", transfersOnt, true},
		{"Transfer Ong", transfersOng, false},
		{"Transfer Fee", transfersFee, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tfs.hasFeeTransfer(); got != tt.want {
				t.Errorf("hasFeeTransfer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransfers_isClaimReward(t *testing.T) {
	tests := []struct {
		name string
		tfs  Transfers
		want bool
	}{
		{"Transfer Claims", transfersClaims, true},
		{"Transfer Ont", transfersOnt, false},
		{"Transfer Ong", transfersOng, false},
		{"Transfer Fee", transfersFee, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tfs.isClaimReward(); got != tt.want {
				t.Errorf("isClaimReward() = %v, want %v", got, tt.want)
			}
		})
	}
}
