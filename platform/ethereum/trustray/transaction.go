package trustray

import (
	"math/big"

	"github.com/trustwallet/golibs/address"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/txtype"
)

func (c *Client) GetTransactions(address string, coinIndex uint) (txtype.TxPage, error) {
	page, err := c.GetTxs(address)
	if err != nil {
		return nil, err
	}
	return normalizePage(page, address, coinIndex), nil
}

func (c *Client) GetTokenTxs(address, token string, coinIndex uint) (txtype.TxPage, error) {
	page, err := c.GetTxsWithContract(address, token)
	if err != nil {
		return nil, err
	}
	return normalizePage(page, address, coinIndex), nil
}

func normalizePage(srcPage *Page, address string, coinIndex uint) txtype.TxPage {
	var txs []txtype.Tx
	for i, srcTx := range srcPage.Docs {
		txs = AppendTxs(txs, &srcTx, coinIndex)
		txs[i].Direction = txs[i].GetTransactionDirection(address)
	}
	return txs
}

func AppendTxs(in []txtype.Tx, srcTx *Doc, coinIndex uint) (out []txtype.Tx) {
	out = in
	baseTx, ok := extractBase(srcTx, coinIndex)
	if !ok {
		return
	}

	// Native ETH transaction
	if len(srcTx.Ops) == 0 && srcTx.Input == "0x" {
		transferTx := baseTx
		transferTx.Meta = txtype.Transfer{
			Value:    txtype.Amount(srcTx.Value),
			Symbol:   coin.Coins[coinIndex].Symbol,
			Decimals: coin.Coins[coinIndex].Decimals,
		}
		out = append(out, transferTx)
		return
	}

	// Smart Contract Call
	if len(srcTx.Ops) == 0 && srcTx.Input != "0x" {
		contractTx := baseTx
		contractTx.Meta = txtype.ContractCall{
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
	if op.Type == txtype.TxTokenTransfer && op.Contract != nil {
		tokenTx := baseTx
		tokenId, err := address.ToEIP55ByCoinID(op.Contract.Address, coinIndex)
		if err != nil {
			return
		}
		tokenTx.Meta = txtype.TokenTransfer{
			Name:     op.Contract.Name,
			Symbol:   op.Contract.Symbol,
			TokenID:  tokenId,
			Decimals: op.Contract.Decimals,
			Value:    txtype.Amount(op.Value),
			From:     op.From,
			To:       op.To,
		}
		out = append(out, tokenTx)
		return
	}
	return
}

func extractBase(srcTx *Doc, coinIndex uint) (base txtype.Tx, ok bool) {
	var (
		status    txtype.Status
		errReason string
	)

	if srcTx.Error == "" {
		status = txtype.StatusCompleted
	} else {
		status = txtype.StatusError
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
	base = txtype.Tx{
		ID:       srcTx.ID,
		Coin:     coinIndex,
		Fee:      txtype.Amount(fee),
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
