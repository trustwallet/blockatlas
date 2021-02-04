package ripple

import (
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/golibs/types"
)

func (p *Platform) GetTransactionsByAccount(account, token string, database *db.Instance) (page types.TxPage, err error) {
	txs, err := database.GetTransactionsByAccount(account, p.Coin().ID)
	if err != nil {
		return
	}
	return models.ToTxPage(txs)
}
