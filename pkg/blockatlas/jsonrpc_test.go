package blockatlas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRpcRequests_fillDefaultValues(t *testing.T) {
	tests := []struct {
		name string
		rs   RpcRequests
		want RpcRequests
	}{
		{
			"test 1",
			RpcRequests{{Method: "method1", Params: "params1"}},
			RpcRequests{{Method: "method1", Params: "params1", JsonRpc: JsonRpcVersion, Id: 1}},
		}, {
			"test 2",
			RpcRequests{
				{Method: "method1", Params: "params1"}, {Method: "method2", Params: "params2"}},
			RpcRequests{
				{Method: "method1", Params: "params1", JsonRpc: JsonRpcVersion, Id: 2},
				{Method: "method2", Params: "params2", JsonRpc: JsonRpcVersion, Id: 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.rs.fillDefaultValues()
			assert.Equal(t, tt.want, got)
		})
	}
}
