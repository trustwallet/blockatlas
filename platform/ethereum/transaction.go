package ethereum

import (
	"math/big"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/address"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/platform/ethereum/trustray"
)

func (p *Platform) RegisterRoutes(router gin.IRouter) {
	router.GET("/:address", func(c *gin.Context) {
		p.getTransactions(c)
	})
}

func (p *Platform) getTransactions(c *gin.Context) {
	token := c.Query("token")
	address := c.Param("address")
	var srcPage *trustray.Page
	var err error

	if token != "" {
		srcPage, err = p.trustray.GetTxsWithContract(address, token)
	} else {
		srcPage, err = p.trustray.GetTxs(address)
	}

	if apiError(c, err) {
		return
	}

	var txs []blockatlas.Tx
	for i, srcTx := range srcPage.Docs {
		txs = AppendTxs(txs, &srcTx, p.CoinIndex)
		txs[i].Direction = txs[i].GetTransactionDirection(address)
	}

	page := blockatlas.TxPage(txs)
	sort.Sort(page)
	c.JSON(http.StatusOK, &page)
}

func extractBase(srcTx *trustray.Doc, coinIndex uint) (base blockatlas.Tx, ok bool) {
	var status blockatlas.Status
	var errReason string
	if srcTx.Error == "" {
		status = blockatlas.StatusCompleted
	} else {
		status = blockatlas.StatusError
		errReason = srcTx.Error
	}

	fee := calcFee(srcTx.GasPrice, srcTx.GasUsed)

	base = blockatlas.Tx{
		ID:       srcTx.ID,
		Coin:     coinIndex,
		From:     srcTx.From,
		To:       srcTx.To,
		Fee:      blockatlas.Amount(fee),
		Date:     srcTx.Timestamp,
		Block:    srcTx.BlockNumber,
		Status:   status,
		Error:    errReason,
		Sequence: srcTx.Nonce,
	}
	return base, true
}

func AppendTxs(in []blockatlas.Tx, srcTx *trustray.Doc, coinIndex uint) (out []blockatlas.Tx) {
	out = in
	baseTx, ok := extractBase(srcTx, coinIndex)
	if !ok {
		return
	}

	// Native ETH transaction
	if len(srcTx.Ops) == 0 && srcTx.Input == "0x" {
		transferTx := baseTx
		transferTx.Meta = blockatlas.Transfer{
			Value:    blockatlas.Amount(srcTx.Value),
			Symbol:   coin.Coins[coinIndex].Symbol,
			Decimals: coin.Coins[coinIndex].Decimals,
		}
		out = append(out, transferTx)
	}

	// Smart Contract Call
	if len(srcTx.Ops) == 0 && srcTx.Input != "0x" {
		contractTx := baseTx
		contractTx.Meta = blockatlas.ContractCall{
			Input: srcTx.Input,
			Value: srcTx.Value,
		}
		out = append(out, contractTx)
	}

	if len(srcTx.Ops) == 0 {
		return
	}
	op := &srcTx.Ops[0]

	if op.Type == blockatlas.TxTokenTransfer && op.Contract != nil {
		tokenTx := baseTx

		tokenTx.Meta = blockatlas.TokenTransfer{
			Name:     op.Contract.Name,
			Symbol:   op.Contract.Symbol,
			TokenID:  address.EIP55Checksum(op.Contract.Address),
			Decimals: op.Contract.Decimals,
			Value:    blockatlas.Amount(op.Value),
			From:     op.From,
			To:       op.To,
		}
		out = append(out, tokenTx)
	}
	return
}

func calcFee(gasPrice string, gasUsed string) string {
	var gasPriceBig, gasUsedBig, feeBig big.Int

	gasPriceBig.SetString(gasPrice, 10)
	gasUsedBig.SetString(gasUsed, 10)

	feeBig.Mul(&gasPriceBig, &gasUsedBig)

	return feeBig.String()
}

func apiError(c *gin.Context, err error) bool {
	if err != nil {
		logger.Error(err, "Unhandled error")
		c.AbortWithStatus(http.StatusInternalServerError)
		return true
	}
	return false
}
