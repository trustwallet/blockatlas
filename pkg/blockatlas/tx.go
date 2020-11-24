package blockatlas

import (
	"github.com/trustwallet/golibs/tokentype"
	"sort"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/golibs/asset"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/numbers"
)

const (
	StatusCompleted Status = "completed"
	StatusPending   Status = "pending"
	StatusError     Status = "error"

	DirectionOutgoing Direction = "outgoing"
	DirectionIncoming Direction = "incoming"
	DirectionSelf     Direction = "yourself"

	TxTransfer              TransactionType = "transfer"
	TxNativeTokenTransfer   TransactionType = "native_token_transfer"
	TxTokenTransfer         TransactionType = "token_transfer"
	TxCollectibleTransfer   TransactionType = "collectible_transfer"
	TxTokenSwap             TransactionType = "token_swap"
	TxContractCall          TransactionType = "contract_call"
	TxAnyAction             TransactionType = "any_action"
	TxMultiCurrencyTransfer TransactionType = "multi_currency_transfer"

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

	// TxPerPage says how many transactions to return per page
	TxPerPage = 25
)

type (
	// Types of transaction statuses
	Direction       string
	Status          string
	TransactionType string
	KeyType         string
	KeyTitle        string

	Block struct {
		Number int64  `json:"number"`
		ID     string `json:"id,omitempty"`
		Txs    []Tx   `json:"txs"`
	}

	// TxPage is a page of transactions
	TxPage []Tx

	// Amount is a positive decimal integer string.
	// It is written in the smallest possible unit (e.g. Wei, Satoshis)
	Amount string

	// Tx describes an on-chain transaction generically
	Tx struct {
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
		// Status of the transaction e.g: "completed", "pending", "error"
		Status Status `json:"status"`
		// Empty if the transaction "completed" or "pending", else error explaining why the transaction failed (optional)
		Error string `json:"error,omitempty"`
		// Transaction nonce or sequence
		Sequence uint64 `json:"sequence"`
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

	TxOutput struct {
		Address string `json:"address"`
		Value   Amount `json:"value"`
	}

	// Transfer describes the transfer of currency native to the platform
	Transfer struct {
		Value    Amount `json:"value"`
		Symbol   string `json:"symbol"`
		Decimals uint   `json:"decimals"`
	}

	// NativeTokenTransfer describes the transfer of native tokens.
	// Example: Stellar Tokens, TRC10
	NativeTokenTransfer struct {
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
	TokenTransfer struct {
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
	CollectibleTransfer struct {
		Name     string `json:"name"`
		Contract string `json:"contract"`
		ImageURL string `json:"image_url"`
	}

	// TokenSwap describes the exchange of two different tokens
	TokenSwap struct {
		Input  TokenTransfer `json:"input"`
		Output TokenTransfer `json:"output"`
	}

	// ContractCall describes a
	ContractCall struct {
		Input string `json:"input"`
		Value string `json:"value"`
	}

	// Currency describes currency information with its amount
	Currency struct {
		Token Token  `json:"token"`
		Value Amount `json:"value"`
	}

	// MultiCurrencyTransfer describes the transfer of multiple currency native to the platform
	MultiCurrencyTransfer struct {
		Currencies []Currency `json:"currencies"`
		Fees       []Currency `json:"fees"`
	}

	// AnyAction describes all other types
	AnyAction struct {
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
	TokenPage []Token

	// Token describes the non-native tokens.
	// Examples: ERC-20, TRC-20, BEP-2
	Token struct {
		Name     string         `json:"name"`
		Symbol   string         `json:"symbol"`
		Decimals uint           `json:"decimals"`
		TokenID  string         `json:"token_id"`
		Coin     uint           `json:"coin"`
		Type     tokentype.Type `json:"type"`
	}

	Txs []Tx
)

func (t Txs) FilterUniqueID() Txs {
	keys := make(map[string]bool)
	list := make(Txs, 0)
	for _, entry := range t {
		if _, value := keys[entry.ID]; !value {
			keys[entry.ID] = true
			list = append(list, entry)
		}
	}
	return list
}

func (txs TxPage) FilterTransactionsByMemo() TxPage {
	result := make(TxPage, 0)
	for _, tx := range txs {
		if !AllowMemo(tx.Memo) {
			tx.Memo = ""
		}
		result = append(result, tx)
	}
	return result
}

func AllowMemo(memo string) bool {
	// only allows numeric values
	_, err := strconv.ParseFloat(memo, 64)
	return err == nil
}

func (txs TxPage) FilterTransactionsByToken(token string) TxPage {
	result := make(TxPage, 0)
	for _, tx := range txs {
		switch tx.Meta.(type) {
		case TokenTransfer:
			if strings.EqualFold(tx.Meta.(TokenTransfer).TokenID, token) {
				result = append(result, tx)
			}
		case *TokenTransfer:
			if strings.EqualFold(tx.Meta.(*TokenTransfer).TokenID, token) {
				result = append(result, tx)
			}
		case NativeTokenTransfer:
			if strings.EqualFold(tx.Meta.(NativeTokenTransfer).TokenID, token) {
				result = append(result, tx)
			}
		case *NativeTokenTransfer:
			if strings.EqualFold(tx.Meta.(*NativeTokenTransfer).TokenID, token) {
				result = append(result, tx)
			}
		case AnyAction:
			if strings.EqualFold(tx.Meta.(AnyAction).TokenID, token) {
				result = append(result, tx)
			}
		case *AnyAction:
			if strings.EqualFold(tx.Meta.(*AnyAction).TokenID, token) {
				result = append(result, tx)
			}
		default:
			continue
		}
	}
	return result
}

func (t Txs) SortByDate() Txs {
	sort.Slice(t, func(i, j int) bool {
		return t[i].Date > t[j].Date
	})
	return t
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
	case Transfer, *Transfer, CollectibleTransfer, *CollectibleTransfer, ContractCall, *ContractCall, AnyAction, *AnyAction, MultiCurrencyTransfer, *MultiCurrencyTransfer:
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

func (t *Tx) TokenID() (string, bool) {
	var tokenID string
	switch t.Meta.(type) {
	case Transfer, *Transfer, CollectibleTransfer, *CollectibleTransfer, ContractCall, *ContractCall, MultiCurrencyTransfer, *MultiCurrencyTransfer:
		return "", false
	case NativeTokenTransfer:
		tokenID = t.Meta.(NativeTokenTransfer).TokenID
	case *NativeTokenTransfer:
		tokenID = t.Meta.(*NativeTokenTransfer).TokenID
	case TokenTransfer:
		tokenID = t.Meta.(TokenTransfer).TokenID
	case *TokenTransfer:
		tokenID = t.Meta.(*TokenTransfer).TokenID
	case AnyAction:
		tokenID = t.Meta.(AnyAction).TokenID
	case *AnyAction:
		tokenID = t.Meta.(*AnyAction).TokenID
	default:
		return "", false
	}
	return tokenID, true
}

func (t *Tx) GetTransactionDirection(address string) Direction {
	if t.Direction != "" {
		return t.Direction
	}
	if len(t.Inputs) > 0 && len(t.Outputs) > 0 {
		addressSet := mapset.NewSet(address)
		return InferDirection(t, addressSet)
	}
	switch meta := t.Meta.(type) {
	case *TokenTransfer:
		return determineTransactionDirection(address, meta.From, meta.To)
	case *NativeTokenTransfer:
		return determineTransactionDirection(address, meta.From, meta.To)
	case TokenTransfer:
		return determineTransactionDirection(address, meta.From, meta.To)
	case NativeTokenTransfer:
		return determineTransactionDirection(address, meta.From, meta.To)
	default:
		return determineTransactionDirection(address, t.From, t.To)
	}
}

func determineTransactionDirection(address, from, to string) Direction {
	if address == to {
		if from == to {
			return DirectionSelf
		}
		return DirectionIncoming
	}
	return DirectionOutgoing
}

func (t *Tx) InferUtxoValue(address string, coinIndex uint) {
	if len(t.Inputs) > 0 && len(t.Outputs) > 0 {
		addressSet := mapset.NewSet(address)
		value := InferValue(t, t.Direction, addressSet)
		t.Meta = Transfer{
			Value:    value,
			Symbol:   coin.Coins[coinIndex].Symbol,
			Decimals: coin.Coins[coinIndex].Decimals,
		}
	}
}

func InferDirection(tx *Tx, addressSet mapset.Set) Direction {
	inputSet := mapset.NewSet()
	for _, address := range tx.Inputs {
		inputSet.Add(address.Address)
	}
	outputSet := mapset.NewSet()
	for _, address := range tx.Outputs {
		outputSet.Add(address.Address)
	}
	intersect := addressSet.Intersect(inputSet)
	if intersect.Cardinality() == 0 {
		return DirectionIncoming
	}
	if outputSet.IsProperSubset(addressSet) || outputSet.Equal(inputSet) {
		return DirectionSelf
	}
	return DirectionOutgoing
}

func InferValue(tx *Tx, direction Direction, addressSet mapset.Set) Amount {
	value := Amount("0")
	if len(tx.Outputs) == 0 {
		return value
	}
	if direction == DirectionOutgoing || direction == DirectionSelf {
		value = tx.Outputs[0].Value
	} else if direction == DirectionIncoming {
		amount := value
		for _, output := range tx.Outputs {
			if !addressSet.Contains(output.Address) {
				continue
			}
			value := numbers.AddAmount(string(amount), string(output.Value))
			amount = Amount(value)
		}
		value = amount
	}
	return value
}

func (t Tx) AssetModel() (models.Asset, bool) {
	var a models.Asset
	switch t.Meta.(type) {
	case TokenTransfer:
		a.Asset = asset.BuildID(t.Coin, t.Meta.(TokenTransfer).TokenID)
		a.Decimals = t.Meta.(TokenTransfer).Decimals
		a.Name = t.Meta.(TokenTransfer).Name
		a.Symbol = t.Meta.(TokenTransfer).Symbol
		tp, ok := GetTokenType(t.Coin, t.Meta.(TokenTransfer).TokenID)
		if !ok {
			return models.Asset{}, false
		}
		a.Type = tp
	case *TokenTransfer:
		a.Asset = asset.BuildID(t.Coin, t.Meta.(*TokenTransfer).TokenID)
		a.Decimals = t.Meta.(*TokenTransfer).Decimals
		a.Name = t.Meta.(*TokenTransfer).Name
		a.Symbol = t.Meta.(*TokenTransfer).Symbol
		tp, ok := GetTokenType(t.Coin, t.Meta.(*TokenTransfer).TokenID)
		if !ok {
			return models.Asset{}, false
		}
		a.Type = tp
	case NativeTokenTransfer:
		a.Asset = asset.BuildID(t.Coin, t.Meta.(NativeTokenTransfer).TokenID)
		a.Decimals = t.Meta.(NativeTokenTransfer).Decimals
		a.Name = t.Meta.(NativeTokenTransfer).Name
		a.Symbol = t.Meta.(NativeTokenTransfer).Symbol
		tp, ok := GetTokenType(t.Coin, t.Meta.(NativeTokenTransfer).TokenID)
		if !ok {
			return models.Asset{}, false
		}
		a.Type = tp
	case *NativeTokenTransfer:
		a.Asset = asset.BuildID(t.Coin, t.Meta.(*NativeTokenTransfer).TokenID)
		a.Decimals = t.Meta.(*NativeTokenTransfer).Decimals
		a.Name = t.Meta.(*NativeTokenTransfer).Name
		a.Symbol = t.Meta.(*NativeTokenTransfer).Symbol
		tp, ok := GetTokenType(t.Coin, t.Meta.(*NativeTokenTransfer).TokenID)
		if !ok {
			return models.Asset{}, false
		}
		a.Type = tp
	case AnyAction:
		a.Asset = asset.BuildID(t.Coin, t.Meta.(AnyAction).TokenID)
		a.Decimals = t.Meta.(AnyAction).Decimals
		a.Name = t.Meta.(AnyAction).Name
		a.Symbol = t.Meta.(AnyAction).Symbol
		tp, ok := GetTokenType(t.Coin, t.Meta.(AnyAction).TokenID)
		if !ok {
			return models.Asset{}, false
		}
		a.Type = tp
	case *AnyAction:
		a.Asset = asset.BuildID(t.Coin, t.Meta.(*AnyAction).TokenID)
		a.Decimals = t.Meta.(*AnyAction).Decimals
		a.Name = t.Meta.(*AnyAction).Name
		a.Symbol = t.Meta.(*AnyAction).Symbol
		tp, ok := GetTokenType(t.Coin, t.Meta.(*AnyAction).TokenID)
		if !ok {
			return models.Asset{}, false
		}
		a.Type = tp
	default:
		return models.Asset{}, false
	}
	if a.Asset == "" {
		return models.Asset{}, false
	}
	a.Coin = t.Coin
	return a, true
}

func GetTokenType(c uint, tokenID string) (string, bool) {
	switch c {
	case coin.Ethereum().ID:
		return string(tokentype.ERC20), true
	case coin.Tron().ID:
		_, err := strconv.Atoi(tokenID)
		if err != nil {
			return string(tokentype.TRC20), true
		}
		return string(tokentype.TRC10), true
	case coin.Smartchain().ID:
		return string(tokentype.BEP20), true
	case coin.Binance().ID:
		return string(tokentype.BEP2), true
	default:
		return "", false
	}
}
