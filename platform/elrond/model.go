package elrond

import (
	"encoding/json"
	"math/big"
	"strings"
	"time"

	"github.com/trustwallet/golibs/types"
)

const roundDurationInSeconds = 6
const mainnetStartTime = 1596117600

type GenericResponse struct {
	Data  json.RawMessage `json:"data"`
	Code  string          `json:"code"`
	Error string          `json:"error"`
}

type NetworkStatus struct {
	Status StatusDetails `json:"status"`
}

type StatusDetails struct {
	Round float64 `json:"erd_current_round"`
	Epoch float64 `json:"erd_epoch_number"`
	Nonce float64 `json:"erd_nonce"`
}

type BlockResponse struct {
	Block Block `json:"hyperblock"`
}

type Block struct {
	Nonce        uint64        `json:"nonce"`
	Round        uint64        `json:"round"`
	Hash         string        `json:"hash"`
	Transactions []Transaction `json:"transactions"`
}

type TransactionsPage struct {
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Hash      string        `json:"hash"`
	Nonce     uint64        `json:"nonce"`
	Value     string        `json:"value"`
	Receiver  string        `json:"receiver"`
	Sender    string        `json:"sender"`
	Data      string        `json:"data"`
	Timestamp time.Duration `json:"timestamp"`
	Status    string        `json:"status"`
	Fee       string        `json:"fee"`
	GasPrice  uint64        `json:"gasPrice,omitempty"`
	GasLimit  uint64        `json:"gasLimit,omitempty"`
}

func (tx *Transaction) TxFee() types.Amount {
	if tx.Fee != "0" && tx.Fee != "" {
		return types.Amount(tx.Fee)
	}

	// Hyperblocks API V1 does not provide the transaction fees (nor "gasUsed"). Hyperblocks API V2 will provide this information, as well.
	// Until then, we can compute the fees deterministically (best-effort) in BlockAtlas, based on gasLimit and gasPrice. This logic will be soon removed from BlockAtlas, in a future PR, when the Hyperblocks API v2 becomes available.
	// Note: For Smart Contract transactions, the refunds will be incompletely provided by the API (until Hyperblocks V2 becomes available): e.g. intra-shard refunds are not visible etc.)

	txFee := big.NewInt(0).SetUint64(tx.GasPrice)
	txFee = txFee.Mul(txFee, big.NewInt(0).SetUint64(tx.GasLimit))

	return types.Amount(txFee.String())
}

func (tx *Transaction) TxStatus() types.Status {
	switch tx.Status {
	case "Success", "success":
		return types.StatusCompleted
	case "Pending", "pending":
		return types.StatusPending
	default:
		return types.StatusError
	}
}

func (tx *Transaction) Direction(address string) types.Direction {
	switch {
	case tx.Sender == address && tx.Receiver == address:
		return types.DirectionSelf
	case tx.Sender == address && tx.Receiver != address:
		return types.DirectionOutgoing
	default:
		return types.DirectionIncoming
	}
}

func (tx *Transaction) HasNegativeValue() bool {
	return strings.HasPrefix(tx.Value, "-")
}

func (tx *Transaction) TxTimestamp(blockRound uint64) time.Duration {
	if int64(tx.Timestamp) != 0 {
		return tx.Timestamp
	}

	// Minor issue (slight inconsistency) about the "timestamp" field:
	// The transactions fetched by querying the endpoint "address/erd1.../transactions" come from our central Elastic Search database, and their timestamp refers to the moment of *including* those transactions in shard blocks (detail: at destination).
	// However, the transactions fetched by querying the endpoint "hyperblock/by-nonce/..." come from another source - from our Observer Nodes - and their timestamp refers to the moment of *final acknowledgement*,
	// according to the Protocol (detail: the moment when the Metachain notarizes the destination shard block containing the transactions).
	// Note: the differences are small - e.g. a few seconds.
	// This inconsistency will be fixed in Hyperblocks API V2.

	txTimestamp := mainnetStartTime + blockRound*roundDurationInSeconds

	return time.Duration(txTimestamp)
}
