package trustray

import (
	"math/big"

	"github.com/trustwallet/golibs/address"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/types"
)

func (c *Client) GetTransactions(address string, coinIndex uint) (types.TxPage, error) {
	page, err := c.GetTxs(address)
	if err != nil {
		return nil, err
	}
	return normalizePage(page, address, coinIndex), nil
}

func (c *Client) GetTokenTxs(address, token string, coinIndex uint) (types.TxPage, error) {
	page, err := c.GetTxsWithContract(address, token)
	if err != nil {
		return nil, err
	}
	return normalizePage(page, address, coinIndex), nil
}

func normalizePage(srcPage *Page, address string, coinIndex uint) types.TxPage {
	var txs []types.Tx
	for i, srcTx := range srcPage.Docs {
		txs = AppendTxs(txs, &srcTx, coinIndex)
		txs[i].Direction = txs[i].GetTransactionDirection(address)
	}
	return txs
}

func AppendTxs(in []types.Tx, srcTx *Doc, coinIndex uint) (out []types.Tx) {
	out = in
	baseTx, ok := extractBase(srcTx, coinIndex)
	if !ok {
		return
	}

	// Native ETH transaction
	if len(srcTx.Ops) == 0 && srcTx.Input == "0x" {
		transferTx := baseTx
		transferTx.Meta = types.Transfer{
			Value:    types.Amount(srcTx.Value),
			Symbol:   coin.Coins[coinIndex].Symbol,
			Decimals: coin.Coins[coinIndex].Decimals,
		}
		out = append(out, transferTx)
		return
	}

	// Smart Contract Call
	if len(srcTx.Ops) == 0 && srcTx.Input != "0x" {
		contractTx := baseTx
		contractTx.Meta = types.ContractCall{
			Input: srcTx.Input,
			Value: srcTx.Value,
		}
		out = append(out, contractTx)
		return
	}

	if len(srcTx.Ops) == 0 {
		return
	}
	op := &srcTx.Ops[0]
	// Token transfer transaction
	if op.Type == types.TxTokenTransfer && op.Contract != nil {
		tokenTx := baseTx
		tokenId, err := address.ToEIP55ByCoinID(op.Contract.Address, coinIndex)
		if err != nil {
			return
		}
		tokenTx.Meta = types.TokenTransfer{
			Name:     op.Contract.Name,
			Symbol:   op.Contract.Symbol,
			TokenID:  tokenId,
			Decimals: op.Contract.Decimals,
			Value:    types.Amount(op.Value),
			From:     op.From,
			To:       op.To,
		}
		out = append(out, tokenTx)
		return
	}
	return
}

func extractBase(srcTx *Doc, coinIndex uint) (base types.Tx, ok bool) {
	var (
		status    types.Status
		errReason string
	)

	if srcTx.Error == "" {
		status = types.StatusCompleted
	} else {
		status = types.StatusError
		errReason = srcTx.Error
	}

	fee := calcFee(srcTx.GasPrice, srcTx.GasUsed)
	from, err := address.ToEIP55ByCoinID(srcTx.From, coinIndex)
	if err != nil {
		return base, false
	}
	to, err := address.ToEIP55ByCoinID(srcTx.To, coinIndex)
	if err != nil {
		return base, false
	}
	base = types.Tx{
		ID:       srcTx.ID,
		Coin:     coinIndex,
		Fee:      types.Amount(fee),
		From:     from,
		To:       to,
		Date:     srcTx.Timestamp,
		Block:    srcTx.BlockNumber,
		Status:   status,
		Error:    errReason,
		Sequence: srcTx.Nonce,
	}

	return base, true
}

func calcFee(gasPrice string, gasUsed string) string {
	var gasPriceBig, gasUsedBig, feeBig big.Int

	gasPriceBig.SetString(gasPrice, 10)
	gasUsedBig.SetString(gasUsed, 10)

	feeBig.Mul(&gasPriceBig, &gasUsedBig)

	return feeBig.String()
}
