package observer

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

func Test_getDirection(t *testing.T) {
	type args struct {
		tx      blockatlas.Tx
		address string
	}
	tests := []struct {
		name string
		args args
		want blockatlas.Direction
	}{
		{"Test Direction Self",
			args{
				blockatlas.Tx{
					From: "0xfc10cab6a50a1ab10c56983c80cc82afc6559cf1", To: "0xfc10cab6a50a1ab10c56983c80cc82afc6559cf1"},
				"0xfc10cab6a50a1ab10c56983c80cc82afc6559cf1"}, blockatlas.DirectionSelf,
		},
		{"Test Direction Outgoing",
			args{
				blockatlas.Tx{
					From: "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB", To: "0x74c8199372c584DAB8b14c519bc8BC8C622F37b7"},
				"0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB"}, blockatlas.DirectionOutgoing,
		},
		{"Test Direction Incoming",
			args{
				blockatlas.Tx{
					From: "0x74c8199372c584DAB8b14c519bc8BC8C622F37b7", To: "0xfc10cab6a50a1ab10c56983c80cc82afc6559cf1"},
				"0xfc10cab6a50a1ab10c56983c80cc82afc6559cf1"}, blockatlas.DirectionIncoming,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDirection(tt.args.tx, tt.args.address); got != tt.want {
				t.Errorf("getDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}
