package ripple

import (
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/golibs/types"
)

func (p *Platform) GetTransactionsByAccount(account, token string, limit int, database *db.Instance) (page types.Txs, err error) {
	txs, err := database.GetTransactionsByAccount(account, p.Coin().ID, limit)
	if err != nil {
		return
	}
	return models.ToTxs(txs)
}
