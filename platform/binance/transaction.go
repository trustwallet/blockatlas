package binance

//const emptyToken = ""
//
//func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
//	explorerResponse, err := p.GetTokenTxsByAddress(address, emptyToken)
//	if err != nil {
//		return nil, err
//	}
//	return filterTxsByType(explorerResponse, blockatlas.TxTransfer), nil
//}
//
//func (p *Platform) GetTokenTxsByAddress(address, token string) (blockatlas.TxPage, error) {
//	explorerResponse, err := p.explorerClient.getTxsOfAddress(address, token)
//	if err != nil {
//		return nil, err
//	}
//
//	explorerTxs, err := p.addTxDetails(explorerResponse.Txs)
//	if err != nil {
//		return nil, err
//	}
//
//	return normalizeTxs(explorerTxs, address), nil
//}
//
//func normalizeTxs(explorerTxs []ExplorerTxs, address string) []blockatlas.Tx {
//	var txs []blockatlas.Tx
//	for _, tx := range explorerTxs {
//		normalizedTxs := normalizeTx(tx, address)
//		if normalizedTxs == nil {
//			continue
//		}
//		txs = append(txs, normalizedTxs...)
//	}
//	return txs
//}
//
//func filterTxsByType(txs []blockatlas.Tx, txType blockatlas.TransactionType) []blockatlas.Tx {
//	var result = make([]blockatlas.Tx, 0, len(txs))
//	for _, tx := range txs {
//		if tx.Type == txType {
//			result = append(result, tx)
//		}
//	}
//	return result
//}
//
//func normalizeTx(srcTx, address string) []blockatlas.Tx {
//	explorerTxType := srcTx.getTransactionType()
//	switch explorerTxType {
//	case SingleTransferOperation:
//		return normalizeSingleTransfer(srcTx, address)
//	case MultiTransferOperation:
//		return normalizeMultiTransfer(srcTx, address)
//	default:
//		return nil
//	}
//}
//
//// Construct base Tx out of explorer transfer using common fields for all type of Blockatlas transfers
//func getBase(srcTx ExplorerTxs) blockatlas.Tx {
//	base := blockatlas.Tx{
//		ID:    srcTx.TxHash,
//		Coin:  coin.BNB,
//		From:  srcTx.FromAddr,
//		Fee:   srcTx.getDexFee(),
//		Date:  srcTx.Timestamp / 1000,
//		Block: srcTx.BlockHeight,
//		Memo:  srcTx.Memo,
//		To:    srcTx.ToAddr,
//	}
//
//	status := srcTx.getStatus()
//	base.Status = status
//	if status == blockatlas.StatusError {
//		base.Error = srcTx.getError()
//	}
//
//	return base
//}
//
//// Extract BEP2 token symbol from asset name e.g: TWT-8C2 => TWT
//func tokenSymbol(asset string) string {
//	s := strings.Split(asset, "-")
//	if len(s) > 1 {
//		return s[0]
//	}
//	return asset
//}
