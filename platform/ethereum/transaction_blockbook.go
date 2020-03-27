package ethereum

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/platform/ethereum/blockbook"
)

func (p *Platform) getTxsFromBlockbook(c *gin.Context) {
	token := c.Query("token")
	address := c.Param("address")
	var srcPage *blockbook.Page
	var err error

	if token != "" {
		srcPage, err = p.blockbook.GetTxsWithContract(address, token)
	} else {
		srcPage, err = p.blockbook.GetTxs(address)
	}

	if apiError(c, err) {
		return
	}

	var txs []blockatlas.Tx
	for _, srcTx := range srcPage.Transactions {
		txs = AppendTxsBlockbook(txs, &srcTx, p.CoinIndex)
	}

	page := blockatlas.TxPage(txs)
	sort.Sort(page)
	c.JSON(http.StatusOK, &page)
}

func extractBaseBlockbook(srcTx *blockbook.Transaction, coinIndex uint) (base blockatlas.Tx, ok bool) {
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

func AppendTxsBlockbook(in []blockatlas.Tx, srcTx *blockbook.Transaction, coinIndex uint) (out []blockatlas.Tx) {
	out = in
	baseTx, ok := extractBaseBlockbook(srcTx, coinIndex)
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

func getStatus(specific *blockbook.EthereumSpecific) (status blockatlas.Status, e string) {
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

func getSenderOrRecipient(t blockbook.Transaction) (address string, err error) {
	if len(t.VIN) > 0 && len(t.VIN[0].Addresses) > 0 {
		return t.VIN[0].Addresses[0], nil
	}
	return "", nil
}

func getTransferType(t blockbook.Transaction) (tt blockatlas.TransactionType, err error) {
	if len(t.TokenTransfers) == 0 {
		return blockatlas.TxTransfer, nil
	}
	if len(t.TokenTransfers) == 1 {
		return blockatlas.TxTokenTransfer, nil
	}
	// TODO blockatlas.TxContractCall
	return "", fmt.Errorf("can't define tranfer type")
}
