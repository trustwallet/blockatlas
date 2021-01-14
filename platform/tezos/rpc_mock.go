package tezos

type RpcClientMock struct {
	Balance string
}

func (r *RpcClientMock) GetAccountAtBlock(address string, block int64) (account Account, err error) {
	return Account{Balance: r.Balance, Delegate: ""}, nil
}
