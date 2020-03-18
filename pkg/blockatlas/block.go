package blockatlas

type Block struct {
	Number int64  `json:"number"`
	ID     string `json:"id,omitempty"`
	Txs    []Tx   `json:"txs"`
}

type Subscription struct {
	Coin    uint   `json:"coin"`
	Address string `json:"address"`
	GUID    string `json:"guid"`
}

func (b *Block) GetTransactionsMap() map[string]*TxSet {
	txMap := make(map[string]*TxSet)
	for i := 0; i < len(b.Txs); i++ {
		addresses := b.Txs[i].GetAddresses()
		addresses = append(addresses, b.Txs[i].GetUtxoAddresses()...)
		for _, address := range addresses {
			if txMap[address] == nil {
				txMap[address] = new(TxSet)
			}
			txMap[address].Add(&b.Txs[i])
		}
	}
	return txMap
}
