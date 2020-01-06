package terra

import (
	"encoding/json"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type TxType string
type EventType string
type AttributeKey string
type DenomType string

// Types of messages
const (
	MsgSend                     TxType = "bank/MsgSend"
	MsgDelegate                 TxType = "staking/MsgDelegate"
	MsgUndelegate               TxType = "staking/MsgUndelegate"
	MsgBeginRedelegate          TxType = "staking/MsgBeginRedelegate"
	MsgWithdrawDelegationReward TxType = "distribution/MsgWithdrawDelegationReward"

	EventTransfer        EventType = "transfer"
	EventWithdrawRewards EventType = "withdraw_rewards"

	AttributeAmount    AttributeKey = "amount"
	AttributeValidator AttributeKey = "validator"

	DenomLuna DenomType = "uluna"
)

// Mapping info for internal and external denoms
var (
	DenomMap = map[string]string{
		"uluna": "LUNA",
		"ukrw":  "KRT",
		"usdr":  "SDT",
		"uusd":  "UST",
		"umnt":  "MNT",
	}
)

// Tx - Base transaction object. Always returned as part of an array
type Tx struct {
	Block  string `json:"height"`
	Code   int    `json:"code"`
	Date   string `json:"timestamp"`
	ID     string `json:"txhash"`
	Data   Data   `json:"tx"`
	Events Events `json:"events"`
}

type TxPage struct {
	Txs []Tx `json:"txs"`
}

// Events
type Event struct {
	Type       EventType
	Attributes Attributes `json:"Attributes"`
}

type Events []*Event

// GetWithdrawRewardValue returns withdrawn rewards as an array of blockatlas.Transfer
func (e Events) GetWithdrawRewardValue() (rewards Amounts) {

	coinMap := make(map[string]int64)
	for _, att := range e {
		if att.Type == EventWithdrawRewards {
			coinMap = att.Attributes.GetWithdrawRewardValue(coinMap)
		}
	}

	var keys []string
	for key := range coinMap {
		keys = append(keys, key)
	}

	// zero fee tx
	if len(keys) == 0 {
		return
	}

	sort.Strings(keys)

	for _, key := range keys {
		rewards = append(rewards, Amount{key, strconv.FormatInt(coinMap[key], 10)})
	}

	return
}

type Attribute struct {
	Key   AttributeKey `json:"key"`
	Value string       `json:"value"`
}

type Attributes []Attribute

// GetWithdrawRewardValue returns the summed coin map
func (a Attributes) GetWithdrawRewardValue(coinMap map[string]int64) map[string]int64 {
	for _, att := range a {
		if att.Key == AttributeAmount {
			coins := strings.Split(att.Value, ",")
			for _, coin := range coins {
				idx := strings.IndexByte(coin, 'u')
				if idx < 0 {
					continue
				}

				denom := coin[idx:]
				amount, err := strconv.ParseInt(coin[:idx], 10, 64)

				if err != nil {
					continue
				}

				if amt, ok := coinMap[denom]; ok {
					coinMap[denom] = amt + amount
				} else {
					coinMap[denom] = amount
				}
			}
		}
	}

	return coinMap
}

// Data - "tx" sub object
type Data struct {
	Contents Contents `json:"value"`
}

// Contents - amount, fee, and memo
type Contents struct {
	Message []Message `json:"msg"`
	Fee     Fee       `json:"fee"`
	Memo    string    `json:"memo"`
}

// Message - an array that holds multiple 'particulars' entries. Possibly used for multiple transfers in one transaction?
type Message struct {
	Type  TxType
	Value interface{}
}

// MessageValueTransfer - from, to, and amount
type MessageValueTransfer struct {
	FromAddr string   `json:"from_address"`
	ToAddr   string   `json:"to_address"`
	Amount   []Amount `json:"amount,omitempty"`
}

// MessageValueDelegate - from, to, and amount
type MessageValueDelegate struct {
	DelegatorAddr string `json:"delegator_address"`
	ValidatorAddr string `json:"validator_address"`
	Amount        Amount `json:"amount,omitempty"`
}

// Fee - also references the "amount" struct
type Fee struct {
	FeeAmount []Amount `json:"amount"`
}

// Amount - the asset & quantity. Always seems to be enclosed in an array/list for some reason.
// Perhaps used for multiple tokens transferred in a single sender/reciever transfer?
type Amount struct {
	Denom    string `json:"denom"`
	Quantity string `json:"amount"`
}

// Amounts - the array of Amount
type Amounts []Amount

func (amounts Amounts) toCoins() (coins []blockatlas.Transfer) {
	for _, amt := range amounts {
		coins = append(coins, blockatlas.Transfer{
			Decimals: coin.Terra().Decimals,
			Symbol:   DenomMap[amt.Denom],
			Value:    blockatlas.Amount(amt.Quantity),
		})
	}
	return
}

// # Staking

type TerraCommission struct {
	Rate string `json:"rate"`
}

type ValidatorsResult struct {
	Validators []Validator `json:"validators"`
}

type Validator struct {
	Status        string          `json:"status"`
	Address       string          `json:"operatorAddress"`
	Commission    TerraCommission `json:"commissionInfo"`
	StakingReturn string          `json:"stakingReturn"`
}

type Delegations struct {
	List []Delegation `json:"result"`
}

type Delegation struct {
	DelegatorAddress string `json:"delegator_address"`
	ValidatorAddress string `json:"validator_address"`
	Shares           string `json:"shares,omitempty"`
}

func (d *Delegation) Value() string {
	shares := strings.Split(d.Shares, ".")
	if len(shares) > 0 {
		return shares[0]
	}
	return d.Shares
}

type UnbondingDelegations struct {
	List []UnbondingDelegation `json:"result"`
}

type UnbondingDelegation struct {
	Delegation
	Entries []UnbondingDelegationEntry `json:"entries"`
}

type UnbondingDelegationEntry struct {
	DelegatorAddress string `json:"creation_height"`
	CompletionTime   string `json:"completion_time"`
	Balance          string `json:"balance"`
}

type StakingPool struct {
	Pool Pool `json:"result"`
}

type Pool struct {
	NotBondedTokens string `json:"not_bonded_tokens"`
	BondedTokens    string `json:"bonded_tokens"`
}

// Block - top object of get las block request
type Block struct {
	Meta BlockMeta `json:"block_meta"`
}

//BlockMeta - "Block" sub object
type BlockMeta struct {
	Header BlockHeader `json:"header"`
}

//BlockHeader - "BlockMeta" sub object, height
type BlockHeader struct {
	Height string `json:"height"`
}

//UnmarshalJSON reads different message types
func (m *Message) UnmarshalJSON(buf []byte) error {
	var messageInternal struct {
		Type  TxType          `json:"type"`
		Value json.RawMessage `json:"value"`
	}

	err := json.Unmarshal(buf, &messageInternal)
	if err != nil {
		return err
	}

	m.Type = messageInternal.Type

	switch messageInternal.Type {
	case MsgUndelegate, MsgDelegate, MsgWithdrawDelegationReward:
		var msgDelegate MessageValueDelegate
		err = json.Unmarshal(messageInternal.Value, &msgDelegate)
		m.Value = msgDelegate
	case MsgSend:
		var msgTransfer MessageValueTransfer
		err = json.Unmarshal(messageInternal.Value, &msgTransfer)
		m.Value = msgTransfer
	}
	return err
}

type AuthAccount struct {
	Account Account `json:"result"`
}

type Account struct {
	Value AccountValue `json:"value"`
}

type AccountValue struct {
	Coins []Balance `json:"coins"`
}

type Balance struct {
	Denom  DenomType `json:"denom"`
	Amount string    `json:"amount"`
}

// StakingReturn defines annualized staking return data
type StakingReturn struct {
	Datetime         time.Time `json:"datetime"`
	DailyReturn      string    `json:"dailyReturn"`
	AnnualizedReturn string    `json:"annualizedReturn"`
}

// StakingReturns is array of StakingReturn
type StakingReturns []StakingReturn
