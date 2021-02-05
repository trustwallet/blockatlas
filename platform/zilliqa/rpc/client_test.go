package rpc

import (
	"reflect"
	"testing"

	"github.com/trustwallet/golibs/client"
)

func Test_makeRequests(t *testing.T) {
	type args struct {
		hashes   []string
		perGroup int
	}
	tests := []struct {
		name string
		args args
		want []client.RpcRequests
	}{
		{
			name: "test group size 1",
			args: args{
				hashes: []string{
					"0x1", "0x2", "0x3",
				},
				perGroup: 1,
			},
			want: []client.RpcRequests{
				{
					&client.RpcRequest{
						Method: "GetTransaction",
						Params: []string{"0x1"},
					},
				},
				{
					&client.RpcRequest{
						Method: "GetTransaction",
						Params: []string{"0x2"},
					},
				},
				{
					&client.RpcRequest{
						Method: "GetTransaction",
						Params: []string{"0x3"},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeBatchRequests(tt.args.hashes, tt.args.perGroup); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeBatchRequests() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeBatches(t *testing.T) {
	type args struct {
		hashes    []string
		batchSize int
	}
	tests := []struct {
		name        string
		args        args
		wantBatches [][]string
	}{
		{
			name: "Test batch size 4",
			args: args{
				hashes: []string{
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11",
				},
				batchSize: 4,
			},
			wantBatches: [][]string{
				{"1", "2", "3", "4"},
				{"5", "6", "7", "8"},
				{"9", "10", "11"},
			},
		},
		{
			name: "Test batch size 10",
			args: args{
				hashes: []string{
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11",
				},
				batchSize: 10,
			},
			wantBatches: [][]string{
				{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
				{"11"},
			},
		},
		{
			name: "Test batch size 11",
			args: args{
				hashes: []string{
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11",
				},
				batchSize: 11,
			},
			wantBatches: [][]string{
				{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotBatches := makeBatches(tt.args.hashes, tt.args.batchSize); !reflect.DeepEqual(gotBatches, tt.wantBatches) {
				t.Errorf("makeBatches() = %v, want %v", gotBatches, tt.wantBatches)
			}
		})
	}
}
