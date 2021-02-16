package solana

const (
	StakeStateUninitialized StakeState = 0
	StakeStateInitialized   StakeState = 1
	StakeStateDelegated     StakeState = 2
	StakeStateRewardsPool   StakeState = 3
)

type StakeState uint32

type VoteAccount struct {
	NodePubkey       string     `json:"nodePubkey"`
	VotePubkey       string     `json:"votePubkey"`
	Commission       uint64     `json:"commission"`
	ActivatedStake   uint64     `json:"activatedStake"`
	RootSlot         uint64     `json:"rootSlot"`
	LastVote         uint64     `json:"lastVote"`
	EpochCredits     [][]uint64 `json:"epochCredits"`
	EpochVoteAccount bool       `json:"epochVoteAccount"`
}

type VoteAccounts struct {
	Current    []VoteAccount `json:"current"`
	Delinquent []VoteAccount `json:"delinquent"`
}

type Account struct {
	Data       string `json:"data"`
	Executable bool   `json:"executable"`
	Lamports   uint64 `json:"lamports"`
	Owner      string `json:"owner"`
	RentEpoch  uint64 `json:"rentEpoch"`
}

type KeyedAccount struct {
	Account Account `json:"account"`
	Pubkey  string  `json:"pubkey"`
}

type RpcAccount struct {
	Context RpcContext `json:"context"`
	Account Account    `json:"value"`
}

type RpcContext struct {
	Slot uint64 `json:"slot"`
}

type StakeData struct {
	State                StakeState
	RentExemptReserve    uint64
	AuthorizedStaker     [32]byte
	AuthorizedWithdrawer [32]byte
	UnixTimestamp        int64
	LockupEpoch          uint64
	Custodian            [32]byte
	VoterPubkey          [32]byte
	Stake                uint64
	ActivationEpoch      uint64
	DeactivationEpoch    uint64
	WarmupCooldownRate   float64
	CreditsObserved      uint64
}

type EpochInfo struct {
	AbsoluteSlot uint64 `json:"absoluteSlot"`
	BlockHeight  uint64 `json:"blockHeight"`
	Epoch        uint64 `json:"epoch"`
	SlotIndex    uint64 `json:"slotIndex"`
	SlotsInEpoch uint64 `json:"slotsInEpoch"`
}
type Block struct {
	BlockTime    int64                  `json:"blockTime"`
	Transactions []ConfirmedTransaction `json:"transactions"`
}

type ConfirmedSignature struct {
	Memo      string `json:"memo"`
	Signature string `json:"signature"`
	Slot      uint64 `json:"slot"`
}

type ConfirmedTransaction struct {
	Meta        Meta        `json:"meta"`
	BlockTime   int64       `json:"blockTime,omitempty"`
	Slot        uint64      `json:"slot,omitempty"`
	Transaction Transaction `json:"transaction"`
}

type Meta struct {
	Err interface{} `json:"err"`
	Fee uint64      `json:"fee"`
}

type TransferInfo struct {
	Destination string `json:"destination"`
	Lamports    uint64 `json:"lamports"`
	Source      string `json:"source"`
}

type Parsed struct {
	Info interface{} `json:"info"`
	Type string      `json:"type"`
}

type TokenTransferInfo struct {
	Destination string      `json:"destination"`
	Mint        string      `json:"mint"`
	Source      string      `json:"source"`
	TokenAmount TokenAmount `json:"tokenAmount"`
}

type TokenAmount struct {
	Amount   string `json:"amount"`
	Decimals uint   `json:"decimals"`
}

type Instruction struct {
	Parsed  interface{} `json:"parsed"`
	Program string      `json:"program"`
}

type Message struct {
	Instructions []Instruction `json:"instructions"`
}

type Transaction struct {
	Message    Message  `json:"message"`
	Signatures []string `json:"signatures"`
}
