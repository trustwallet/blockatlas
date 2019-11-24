package api

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
)

func getDelegationResponse(p blockatlas.StakeAPI, address string) (blockatlas.DelegationResponse, error) {
	delegations, err := p.GetDelegations(address)
	if err != nil {
		return blockatlas.DelegationResponse{}, errors.E("Unable to fetch delegations list", err)
	}
	balance, err := p.UndelegatedBalance(address)
	if err != nil {
		return blockatlas.DelegationResponse{}, errors.E("Unable to fetch undelegated balance", err)
	}
	c := p.Coin()
	return blockatlas.DelegationResponse{
		Coin:        c.External(),
		Details:     p.GetDetails(),
		Address:     address,
		Balance:     balance,
		Delegations: delegations,
	}, nil
}
