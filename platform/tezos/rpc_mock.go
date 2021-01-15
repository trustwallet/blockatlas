package tezos

type RpcClientMock struct {
	Balance string
}

func (r *RpcClientMock) GetAccountBalanceAtBlock(address string, block int64) (account AccountBalance, err error) {
	return AccountBalance{Balance: r.Balance}, nil
}
