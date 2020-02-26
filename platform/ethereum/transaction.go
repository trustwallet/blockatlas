package ethereum

import (
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/address"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/ginutils"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"math/big"
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
	for _, srcTx := range srcPage.Docs {
		txs = AppendTxs(txs, &srcTx, p.CoinIndex)
	}

	page := blockatlas.TxPage(txs)
	sort.Sort(page)
	c.JSON(http.StatusOK, &page)
}

func extractBase(srcTx *Doc, coinIndex uint) (base blockatlas.Tx, ok bool) {
	var status blockatlas.Status
	var errReason string
	if srcTx.Error == "" {
		status = blockatlas.StatusCompleted
	} else {
		status = blockatlas.StatusFailed
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

func AppendTxs(in []blockatlas.Tx, srcTx *Doc, coinIndex uint) (out []blockatlas.Tx) {
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
		ginutils.RenderError(c, http.StatusInternalServerError, err.Error())
		return true
	}
	return false
}
