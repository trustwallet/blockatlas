package blockatlas

type Direction string
type Status string
type TokenType string
type TransactionType string
type KeyType string
type KeyTitle string

// Types of transaction statuses
const (
	StatusCompleted Status = "completed"
	StatusPending   Status = "pending"
	StatusFailed    Status = "failed"

	DirectionOutgoing Direction = "outgoing"
	DirectionIncoming Direction = "incoming"
	DirectionSelf     Direction = "yourself"

	TokenTypeERC20 TokenType = "ERC20"
	TokenTypeBEP2  TokenType = "BEP2"
	TokenTypeTRC10 TokenType = "TRC10"

	TxTransfer            TransactionType = "transfer"
	TxMultiCoinTransfer   TransactionType = "multi_coin_transfer"
	TxNativeTokenTransfer TransactionType = "native_token_transfer"
	TxTokenTransfer       TransactionType = "token_transfer"
	TxCollectibleTransfer TransactionType = "collectible_transfer"
	TxTokenSwap           TransactionType = "token_swap"
	TxContractCall        TransactionType = "contract_call"
	TxAnyAction           TransactionType = "any_action"
	TxMultiCoinAnyAction  TransactionType = "multi_coin_any_action"

	KeyPlaceOrder        KeyType = "place_order"
	KeyCancelOrder       KeyType = "cancel_order"
	KeyIssueToken        KeyType = "issue_token"
	KeyBurnToken         KeyType = "burn_token"
	KeyMintToken         KeyType = "mint_token"
	KeyApproveToken      KeyType = "approve_token"
	KeyStakeDelegate     KeyType = "stake_delegate"
	KeyStakeClaimRewards KeyType = "stake_claim_rewards"

	KeyTitlePlaceOrder    KeyTitle = "Place Order"
	KeyTitleCancelOrder   KeyTitle = "Cancel Order"
	AnyActionDelegation   KeyTitle = "Delegation"
	AnyActionUndelegation KeyTitle = "Undelegation"
	AnyActionClaimRewards KeyTitle = "Claim Rewards"
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
	Status Status `json:"status"`
	// Empty if the transaction was successful,
	// else error explaining why the transaction failed (optional)
	Error string `json:"error,omitempty"`
	// Transaction nonce or sequence
	Sequence uint64 `json:"sequence,omitempty"`
	// Type of metadata
	Type TransactionType `json:"type"`
	// Input addresses
	Inputs []TxOutput `json:"inputs,omitempty"`
	// Output addresses
	Outputs []TxOutput `json:"outputs,omitempty"`
	// Transaction Direction
	Direction Direction `json:"direction,omitempty"`
	// Meta data object
	Memo string      `json:"memo"`
	Meta interface{} `json:"metadata"`
}

// TxOutput describes input and output of tx
// Excample: bitcoin UTXO
type TxOutput struct {
	Address string `json:"address"`
	Value   Amount `json:"value"`
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

// MultiCoinTransfer describes the transfer of multiple currencys native to the platform
// Example: Cosmos, Terra
type MultiCoinTransfer struct {
	Coins []Transfer `json:"coins"`
	Fees  []Transfer `json:"fees"`
}

// MultiCoinAnyAction describes any action for multiple currencys
// Example: Cosmos, Terra
type MultiCoinAnyAction struct {
	Title KeyTitle   `json:"title"`
	Key   KeyType    `json:"key"`
	Fees  []Transfer `json:"fees"`
	Coins []Transfer `json:"coins"`
}

// AnyAction describes all other types
type AnyAction struct {
	Coin     uint     `json:"coin"`
	Title    KeyTitle `json:"title"`
	Key      KeyType  `json:"key"`
	TokenID  string   `json:"token_id"`
	Name     string   `json:"name"`
	Symbol   string   `json:"symbol"`
	Decimals uint     `json:"decimals"`
	Value    Amount   `json:"value"`
}

// TokenPage is a page of transactions.
type TokenPage []Token

// Token describes the non-native tokens.
// Examples: ERC-20, TRC-20, BEP-2
type Token struct {
	Name     string    `json:"name"`
	Symbol   string    `json:"symbol"`
	Decimals uint      `json:"decimals"`
	TokenID  string    `json:"token_id"`
	Coin     uint      `json:"coin"`
	Type     TokenType `json:"type"`
}

func (t *Tx) GetUtxoAddresses() (addresses []string) {
	for _, input := range t.Inputs {
		addresses = append(addresses, input.Address)
	}

	for _, output := range t.Outputs {
		addresses = append(addresses, output.Address)
	}

	return addresses
}

func (t *Tx) GetAddresses() []string {
	addresses := make([]string, 0)
	switch t.Meta.(type) {
	case Transfer, *Transfer, CollectibleTransfer, *CollectibleTransfer, ContractCall, *ContractCall, AnyAction, *AnyAction, MultiCoinTransfer, *MultiCoinTransfer, MultiCoinAnyAction, *MultiCoinAnyAction:
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
