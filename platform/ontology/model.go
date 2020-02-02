package ontology

type AssetType string
type Transfers []Transfer

const (
	GovernanceContract = "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK"

	AssetONT AssetType = "ont"
	AssetONG AssetType = "ong"
)

type TxPage struct {
	Result Result `json:"Result"`
}

type Result struct {
	TxnList []Tx `json:"TxnList"`
}

type Transfer struct {
	Amount      string    `json:"Amount"`
	FromAddress string    `json:"FromAddress"`
	ToAddress   string    `json:"ToAddress"`
	AssetName   AssetType `json:"AssetName"`
}

type Tx struct {
	TxnHash      string    `json:"TxnHash"`
	ConfirmFlag  uint64    `json:"ConfirmFlag"`
	TxnType      uint64    `json:"TxnType"`
	TxnTime      int64     `json:"TxnTime"`
	Height       uint64    `json:"Height"`
	Fee          string    `json:"Fee"`
	BlockIndex   uint64    `json:"BlockIndex"`
	TransferList Transfers `json:"TransferList"`
}

type BlockResults struct {
	Error  int     `json:"Error"`
	Result []Block `json:"Result"`
}

type BlockResult struct {
	Error  int   `json:"Error"`
	Result Block `json:"Result"`
}

type Block struct {
	Height  int64  `json:"Height"`
	TxnList []Tx   `json:"TxnList"`
	Hash    string `json:"Hash"`
}

type TxResponse struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Result TxV2   `json:"Result"`
}

type TxV2 struct {
	Hash        string             `json:"tx_hash"`
	Type        int                `json:"tx_type"`
	Time        int64              `json:"tx_time"`
	BlockHeight uint64             `json:"block_height"`
	Fee         string             `json:"fee"`
	Description string             `json:"description"`
	BlockIndex  int                `json:"block_index"`
	ConfirmFlag int                `json:"confirm_flag"`
	EventType   int                `json:"event_type"`
	Details     TransactionDetails `json:"detail"`
}

type TransactionDetails struct {
	Transfers Transfers `json:"transfers"`
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

func (tfs Transfers) getTransfer() *Transfer {
	for _, tf := range tfs {
		if !tf.isFeeTransfer() {
			return &tf
		}
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
	// Verify if one of the transfers is a fee transfer.
	if !tfs.hasFeeTransfer() {
		return false
	}
	return true

}
