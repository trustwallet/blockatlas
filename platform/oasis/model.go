package oasis

import "math/big"

type Block struct {
	Height    int64  `json:"height"`
	Hash      string `json:"hash"`
	Timestamp int64  `json:"timestamp"`
}

type BlockRequest struct {
	BlockIdentifier int64 `json:"block_identifier"`
}

type DelegationsForRequest struct {
	Owner string `json:"owner"`
}
type DebondingDelegationsForRequest struct {
	Owner string `json:"owner"`
}

type Transaction struct {
	Hash     string `json:"tx_hash"`
	From     string `json:"from"`
	To       string `json:"to"`
	Amount   string `json:"amount"`
	Fee      string `json:"fee"`
	Date     int64  `json:"date"`
	Block    uint64 `json:"block"`
	Success  bool   `json:"success"`
	ErrorMsg string `json:"error_message,omitempty"`
	Sequence uint64 `json:"sequence"`
}

type Validator struct {
	ID          string `json:"id"`
	VotingPower int64  `json:"voting_power"`
	EffectiveAnnualReward float64 `json:"effective_annual_reward"`
}

type ConsensusParams struct {
	DebondingInterval                 uint64                    `json:"debonding_interval"`
	MinDelegationAmount               uint64                   `json:"min_delegation"`
}

type DelegationsFor struct {
	List map[string]Delegation `json:"delegations"`
}

type DebondingDelegationsFor struct {
	List map[string][]DebondingDelegation `json:"debonding_delegations"`
}

// Delegation is a delegation descriptor.
type Delegation struct {
	Shares big.Int `json:"shares"`
}

// DebondingDelegation is a debonding delegation descriptor.
type DebondingDelegation struct {
	Shares        big.Int `json:"shares"`
	DebondEndTime uint64  `json:"debond_end"`
}

type TransactionsByAddressRequest struct {
	Address string `json:"address"`
}
