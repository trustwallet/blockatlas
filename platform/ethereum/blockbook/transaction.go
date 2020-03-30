package blockbook

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (c *Client) GetTransactions(address string, coinIndex uint) (blockatlas.TxPage, error) {
	page, err := c.GetTxs(address)
	if err != nil {
		return nil, err
	}
	return normalizePage(page, address, coinIndex), nil
}

func (c *Client) GetTokenTxs(address, token string, coinIndex uint) (blockatlas.TxPage, error) {
	page, err := c.GetTxsWithContract(address, token)
	if err != nil {
		return nil, err
	}
	return normalizePage(page, address, coinIndex), nil
}

func normalizePage(srcPage *Page, address string, coinIndex uint) blockatlas.TxPage {
	var txs []blockatlas.Tx
	for _, srcTx := range srcPage.Transactions {
		txs = AppendTxs(txs, &srcTx, coinIndex)
	}

	page := blockatlas.TxPage(txs)
	return page
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
	case blockatlas.TxContractCall:
		contractTx := baseTx
		contractTx.Meta = blockatlas.ContractCall{
			Input: baseTx.To,
			Value: srcTx.Value,
		}
		out = append(out, contractTx)
	}
	return
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
	if len(t.TokenTransfers) == 1 && t.TokenTransfers[0].Token == t.VOUT[0].Addresses[0] {
		return blockatlas.TxTokenTransfer, nil
	}
	if t.Value != "0" && t.EthereumSpecific.GasLimit.Uint64() == 21000 {
		return blockatlas.TxTransfer, nil
	}
	return blockatlas.TxContractCall, nil
}
