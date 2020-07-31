package stake

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"sort"
	"strconv"
)

func SortDelegations(delegations blockatlas.DelegationsPage) blockatlas.DelegationsPage {
	sort.Slice(delegations, func(i, j int) bool {
		iA, err := strconv.Atoi(delegations[i].Value)
		if err != nil {
			return false
		}
		jA, err := strconv.Atoi(delegations[j].Value)
		if err != nil {
			return false
		}
		return iA > jA
	})
	return delegations
}
