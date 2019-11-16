package api

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
)

func getDelegationResponse(p blockatlas.StakeAPI, address string) (response blockatlas.DelegationResponse) {
	c := p.Coin()
	response.Coin = c.External()
	response.Address = address

	delegations, err := p.GetDelegations(address)
	balance, err := p.UndelegatedBalance(address)

	if err != nil {
		response.Error = errors.E(err, "Unable to fetch delegations list from the registry").Error()
	} else {
		response.Delegations = delegations
		response.Details = p.GetDetails()
		response.Balance = balance
	}

	return
}
