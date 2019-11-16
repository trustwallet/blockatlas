package api

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func getDelegationResponse(p blockatlas.StakeAPI, address string) (response blockatlas.DelegationResponse) {
	c := p.Coin()
	response.Coin = c.External()
	response.Address = address

	delegations, err := p.GetDelegations(address)
	balance, err := p.UndelegatedBalance(address)

	if err != nil {
		response.Error = err.Error()
	} else {
		response.Delegations = delegations
		response.Details = p.GetDetails()
		response.Balance = balance
	}

	return
}
