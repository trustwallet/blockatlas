package ethereum_blockbook

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"net/http"
	"sort"
)

func (p *Platform) RegisterRoutes(router gin.IRouter) {
	router.GET("/:address", func(c *gin.Context) {
		p.getTransactions(c)
	})
}

func (p *Platform) getTransactions(c *gin.Context) {
	token := c.Query("token")
	address := c.Param("address")
	var srcPage *Page
	var err error

	if token != "" {
		srcPage, err = p.client.GetTxsWithContract(address, token)
	} else {
		srcPage, err = p.client.GetTxs(address)
	}

	if apiError(c, err) {
		return
	}

	var txs []blockatlas.Tx
	for _, srcTx := range srcPage.Transactions {
		txs = AppendTxs(txs, &srcTx, p.CoinIndex)
	}

	page := blockatlas.TxPage(txs)
	sort.Sort(page)
	c.JSON(http.StatusOK, &page)
}

func extractBase(srcTx *Transaction, coinIndex uint) (base blockatlas.Tx, ok bool) {
	status, errReason := getStatus(srcTx.EthereumSpecific)
	from, err := getSenderOrRecipient(*srcTx)
	if err != nil {
		return
	}

	to, err := getSenderOrRecipient(*srcTx)
	if err != nil {
		return
	}

	base = blockatlas.Tx{
		ID:       srcTx.TxID,
		Coin:     coinIndex,
		From:     from,
		To:       to,
		Fee:      blockatlas.Amount(srcTx.Fees),
		Date:     srcTx.BlockTime,
		Block:    srcTx.BlockHeight,
		Status:   status,
		Error:    errReason,
		Sequence: srcTx.EthereumSpecific.Nonce,
	}
	return base, true
}

func AppendTxs(in []blockatlas.Tx, srcTx *Transaction, coinIndex uint) (out []blockatlas.Tx) {
	out = in
	baseTx, ok := extractBase(srcTx, coinIndex)
	if !ok {
		return
	}

	transferType, err := getTransferType(*srcTx)
	if err != nil {
		return
	}

	switch transferType {
	case blockatlas.TxTransfer:
		transfer := baseTx
		transfer.Meta = blockatlas.Transfer{
			Value:    blockatlas.Amount(srcTx.Value),
			Symbol:   coin.Coins[coinIndex].Symbol,
			Decimals: coin.Coins[coinIndex].Decimals,
		}
		out = append(out, transfer)
	case blockatlas.TxTokenTransfer:
		tokenTransfer := baseTx
		contractMeta := srcTx.TokenTransfers[0]

		tokenTransfer.Meta = blockatlas.TokenTransfer{
			Name:     contractMeta.Name,
			Symbol:   contractMeta.Symbol,
			TokenID:  contractMeta.Token,
			Decimals: contractMeta.Decimals,
			Value:    blockatlas.Amount(contractMeta.Value),
			From:     contractMeta.From,
			To:       contractMeta.To,
		}
		out = append(out, tokenTransfer)
	}

	// Smart Contract Call
	//if len(srcTx.TokenTransfers) > 0 && srcTx.Input != "0x" {
	//	contractTx := baseTx
	//	contractTx.Meta = blockatlas.ContractCall{
	//		Input: srcTx.Input,
	//		Value: srcTx.Value,
	//	}
	//	out = append(out, contractTx)
	//}

	return
}

func getStatus(specific *EthereumSpecific) (status blockatlas.Status, e string) {
	switch specific.Status {
	case -1:
		return blockatlas.StatusPending, ""
	case 0:
		return blockatlas.StatusError, "Error"
	case 1:
		return blockatlas.StatusCompleted, ""
	default:
		return blockatlas.StatusError, "Unable to define transaction status"
	}
}

func getSenderOrRecipient(t Transaction) (address string, err error) {
	if len(t.VIN) > 0 && len(t.VIN[0].Addresses) > 0 {
		return t.VIN[0].Addresses[0], nil
	}
	return "", nil
}

func getTransferType(t Transaction) (tt blockatlas.TransactionType, err error) {
	if len(t.TokenTransfers) == 0 {
		return blockatlas.TxTransfer, nil
	}
	if len(t.TokenTransfers) == 1 {
		return blockatlas.TxTokenTransfer, nil
	}
	// TODO blockatlas.TxContractCall
	return "", fmt.Errorf("can't define tranfer type")
}

func apiError(c *gin.Context, err error) bool {
	if err != nil {
		logger.Error(err, "Unhandled error")
		c.AbortWithStatus(http.StatusInternalServerError)
		return true
	}
	return false
}
