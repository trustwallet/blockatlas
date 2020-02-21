package ontology

type AssetType string
type MsgType string
type Transfers []Transfer
type Balances []Balance

const (
	GovernanceContract = "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK"
	ONGDecimals        = 9

	MsgSuccess MsgType = "SUCCESS"

	AssetONT AssetType = "ont"
	AssetONG AssetType = "ong"
	AssetAll AssetType = "all"
)

type BaseResponse struct {
	Code int     `json:"code"`
	Msg  MsgType `json:"msg"`
}

type BlockResults struct {
	BaseResponse
	Result Block `json:"result"`
}

type BlockResult struct {
	BaseResponse
	Result BlockRecords `json:"result"`
}

type TxsResult struct {
	BaseResponse
	Result []Tx `json:"result"`
}

type TxResult struct {
	BaseResponse
	Result Tx `json:"result"`
}

type BalancesResult struct {
	BaseResponse
	Result Balances `json:"result"`
}

type BlockRecords struct {
	Total   int64   `json:"total"`
	Records []Block `json:"records"`
}

type Block struct {
	Height int64  `json:"block_height"`
	Hash   string `json:"block_hash"`
	Txs    []Tx   `json:"txs"`
}

type Tx struct {
	Hash        string    `json:"tx_hash"`
	ConfirmFlag uint64    `json:"confirm_flag"`
	Time        int64     `json:"tx_time"`
	Height      uint64    `json:"block_height"`
	Fee         string    `json:"fee"`
	BlockIndex  uint64    `json:"block_index"`
	EventType   uint64    `json:"event_type,omitempty"`
	Description string    `json:"description,omitempty"`
	Details     Detail    `json:"detail,omitempty"`
	Transfers   Transfers `json:"transfers,omitempty"`
}

type Detail struct {
	Transfers Transfers `json:"transfers"`
}

type Transfer struct {
	Amount      string    `json:"amount"`
	FromAddress string    `json:"from_address"`
	ToAddress   string    `json:"to_address"`
	AssetName   AssetType `json:"asset_name"`
	Description string    `json:"description,omitempty"`
}

type Balance struct {
	Balance      string    `json:"balance"`
	AssetName    AssetType `json:"asset_name"`
	AssetType    string    `json:"asset_type"`
	ContractHash string    `json:"contract_hash"`
}

func (tf *Transfer) isFeeTransfer() bool {
	if tf.AssetName != AssetONG {
		return false
	}
	if tf.ToAddress != GovernanceContract {
		return false
	}
	return true
}

func (tfs Transfers) hasFeeTransfer() bool {
	for _, tf := range tfs {
		if tf.isFeeTransfer() {
			return true
		}
	}
	return false
}

func (tx *Tx) getTransfers() Transfers {
	return append(tx.Details.Transfers, tx.Transfers...)
}

func (tfs Transfers) getTransfer(assetType AssetType) *Transfer {
	for _, tf := range tfs {
		if tf.isFeeTransfer() {
			continue
		}
		if assetType != AssetAll && tf.AssetName != assetType {
			continue
		}
		return &tf
	}
	return nil
}

func (tfs Transfers) isClaimReward() bool {
	// Claim Reward needs to have two transfers.
	if len(tfs) < 2 {
		return false
	}
	// Both transfers need to be ONG, one for reward and another one.
	if tfs[0].AssetName != AssetONG || tfs[1].AssetName != AssetONG {
		return false
	}
	// Avoid transactions for yourself
	if tfs[0].ToAddress == tfs[0].FromAddress || tfs[1].ToAddress == tfs[1].FromAddress {
		return false
	}
	// Verify if one of the transfers is a fee transfer.
	if !tfs.hasFeeTransfer() {
		return false
	}
	// The user need to pay a fee to get his reward
	if tfs[1].isFeeTransfer() && tfs[0].ToAddress == tfs[1].FromAddress && tfs[1].ToAddress == GovernanceContract {
		return true
	}
	if tfs[0].isFeeTransfer() && tfs[1].ToAddress == tfs[0].FromAddress && tfs[0].ToAddress == GovernanceContract {
		return false
	}
	return false

}

func (bs Balances) getBalance(assetType AssetType) *Balance {
	for _, b := range bs {
		if b.AssetName == assetType {
			return &b
		}
	}
	return nil
}
