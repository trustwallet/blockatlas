package blockatlas

// Types of transaction metadata
const (
	TxTransfer            = "transfer"
	TxNativeTokenTransfer = "native_token_transfer"
	TxTokenTransfer       = "token_transfer"
	TxCollectibleTransfer = "collectible_transfer"
	TxTokenSwap           = "token_swap"
	TxContractCall        = "contract_call"
	TxAnyAction           = "any_action"
)

// Types of transaction statuses
const (
	StatusCompleted = "completed"
	StatusFailed    = "failed"
)

// Titles of AnyAction meta
const (
	AnyActionDelegation   = "Delegation"
	AnyActionUndelegation = "Undelegation"
)

// Keys of AnyAction meta
const (
	KeyPlaceOrder    = "place_order"
	KeyCancelOrder   = "cancel_order"
	KeyIssueToken    = "issue_token"
	KeyBurnToken     = "burn_token"
	KeyMintToken     = "mint_token"
	KeyApproveToken  = "approve_token"
	KeyStakeDelegate = "stake_delegate"
)

// TxPerPage says how many transactions to return per page
const TxPerPage = 25

// TxPage is a page of transactions
type TxPage []Tx

// Amount is a positive decimal integer string.
// It is written in the smallest possible unit (e.g. Wei, Satoshis)
type Amount string

// Tx describes an on-chain transaction generically
type Tx struct {
	// Unique identifier
	ID string `json:"id"`
	// SLIP-44 coin index of the platform
	Coin uint `json:"coin"`
	// Address of the transaction sender
	From string `json:"from"`
	// Address of the transaction recipient
	To string `json:"to"`
	// Transaction fee (native currency)
	Fee Amount `json:"fee"`
	// Unix timestamp of the block the transaction was included in
	Date int64 `json:"date"`
	// Height of the block the transaction was included in
	Block uint64 `json:"block"`
	// Status of the transaction
	Status string `json:"status"`
	// Empty if the transaction was successful,
	// else error explaining why the transaction failed (optional)
	Error string `json:"error,omitempty"`
	// Transaction nonce or sequence
	Sequence uint64 `json:"sequence,omitempty"`
	// Type of metadata
	Type string `json:"type"`
	// Meta data object
	Memo string      `json:"memo"`
	Meta interface{} `json:"metadata"`
}

// Transfer describes the transfer of currency native to the platform
type Transfer struct {
	Value    Amount `json:"value"`
	Symbol   string `json:"symbol"`
	Decimals uint   `json:"decimals"`
}

// NativeTokenTransfer describes the transfer of native tokens.
// Example: Stellar Tokens, TRC10
type NativeTokenTransfer struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	TokenID  string `json:"token_id"`
	Decimals uint   `json:"decimals"`
	Value    Amount `json:"value"`
	From     string `json:"from"`
	To       string `json:"to"`
}

// TokenTransfer describes the transfer of non-native tokens.
// Examples: ERC-20, TRC20
type TokenTransfer struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	TokenID  string `json:"token_id"`
	Decimals uint   `json:"decimals"`
	Value    Amount `json:"value"`
	From     string `json:"from"`
	To       string `json:"to"`
}

// CollectibleTransfer describes the transfer of a
// "collectible", unique token.
type CollectibleTransfer struct {
	Name     string `json:"name"`
	Contract string `json:"contract"`
	ImageURL string `json:"image_url"`
}

// TokenSwap describes the exchange of two different tokens
type TokenSwap struct {
	Input  TokenTransfer `json:"input"`
	Output TokenTransfer `json:"output"`
}

// ContractCall describes a
type ContractCall struct {
	Input string `json:"input"`
	Value string `json:"value"`
}

// AnyAction describes all other types
type AnyAction struct {
	Coin     uint   `json:"coin"`
	Title    string `json:"title"`
	Key      string `json:"key"`
	TokenID  string `json:"tokenID,omitempty"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals uint   `json:"decimals"`
	Value    Amount `json:"value"`
}

// TokenPage is a page of transactions.
type TokenPage []Token

// Token describes the non-native tokens.
// Examples: ERC-20, TRC-20, BEP-2
type Token struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals uint   `json:"decimals"`
	TokenId  string `json:"tokenID"`
	Coin     uint   `json:"coin"`
}

func (t *Tx) GetAddresses() (addresses []string) {
	switch t.Meta.(type) {
	case Transfer, *Transfer, CollectibleTransfer, *CollectibleTransfer, ContractCall, *ContractCall, AnyAction, *AnyAction:
		return append(addresses, t.From, t.To)
	case NativeTokenTransfer:
		return append(addresses, t.Meta.(NativeTokenTransfer).From, t.Meta.(NativeTokenTransfer).To)
	case *NativeTokenTransfer:
		return append(addresses, t.Meta.(*NativeTokenTransfer).From, t.Meta.(*NativeTokenTransfer).To)
	case TokenTransfer:
		return append(addresses, t.Meta.(TokenTransfer).From, t.Meta.(TokenTransfer).To)
	case *TokenTransfer:
		return append(addresses, t.Meta.(*TokenTransfer).From, t.Meta.(*TokenTransfer).To)
	case TokenSwap:
		{
			m := t.Meta.(TokenSwap)
			return append(addresses, m.Input.From, m.Input.To, m.Output.From, m.Output.To)
		}
	case *TokenSwap:
		{
			m := t.Meta.(*TokenSwap)
			return append(addresses, m.Input.From, m.Input.To, m.Output.From, m.Output.To)
		}
	default:
		return addresses
	}
}
