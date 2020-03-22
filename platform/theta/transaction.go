package theta

import (
	"fmt"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"strconv"
	"strings"
)

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	// Endpoint supports queries without token query parameter
	return p.GetTokenTxsByAddress(address, "")
}

func (p *Platform) GetTokenTxsByAddress(address, token string) (blockatlas.TxPage, error) {
	trx, err := p.client.FetchAddressTransactions(address)
	if err != nil {
		return nil, err
	}

	var txs []blockatlas.Tx
	for _, tr := range trx {
		if tr.Type != SendTransaction {
			continue
		}
		tx, ok := Normalize(&tr, address, token)
		if !ok {
			continue
		}
		txs = append(txs, tx)
	}

	return txs, nil
}

func Normalize(trx *Tx, address, token string) (tx blockatlas.Tx, ok bool) {
	time, _ := strconv.ParseInt(trx.Timestamp, 10, 64)
	block, _ := strconv.ParseUint(trx.BlockHeight, 10, 64)

	tx = blockatlas.Tx{
		ID:       trx.Hash,
		Coin:     coin.THETA,
		Fee:      blockatlas.Amount(trx.Data.Fee.Tfuelwei),
		Date:     time,
		Block:    block,
		Sequence: block,
	}
	inputs := trx.Data.Inputs
	outputs := trx.Data.Outputs

	// Doesn't support multisend
	if len(inputs) == 0 && len(outputs) == 0 {
		return tx, false
	}

	input := inputs[0]
	output := outputs[0]
	sequence, _ := strconv.ParseUint(input.Sequence, 10, 64)

	// Condition for transfer THETA transfer
	if address != "" && token == "" && output.Coins.Tfuelwei == "0" {
		direction, err := getDirection(address, input, output)
		if err != nil {
			return tx, false
		}
		tx.From = input.Address
		tx.To = output.Address
		tx.Sequence = sequence
		tx.Direction = direction
		tx.Type = blockatlas.TxTransfer
		tx.Meta = blockatlas.Transfer{
			Value:    blockatlas.Amount(output.Coins.Thetawei),
			Symbol:   coin.Coins[coin.THETA].Symbol,
			Decimals: coin.Coins[coin.THETA].Decimals,
		}

		return tx, true
	}

	// Condition for transfer Theta Fuel (TFUEL)
	if address != "" && token == "tfuel" && output.Coins.Thetawei == "0" {
		from := input.Address
		to := output.Address
		direction, err := getDirection(address, input, output)
		if err != nil {
			return tx, false
		}

		tx.From = from
		tx.To = to
		tx.Sequence = sequence
		tx.Direction = direction
		tx.Type = blockatlas.TxNativeTokenTransfer
		tx.Meta = blockatlas.NativeTokenTransfer{
			Name:     "Theta Fuel",
			Symbol:   "TFUEL",
			TokenID:  "tfuel",
			Decimals: 18,
			Value:    blockatlas.Amount(output.Coins.Tfuelwei),
			From:     from,
			To:       to,
		}

		return tx, true
	}

	return tx, false
}

// Get transaction direction
func getDirection(a string, inputs Input, outputs Output) (dir blockatlas.Direction, err error) {
	address := strings.ToLower(a)
	inAddr := inputs.Address
	outAddr := outputs.Address

	switch {
	case inAddr == address && outAddr == address:
		return blockatlas.DirectionSelf, nil
	case inAddr == address && outAddr != address:
		return blockatlas.DirectionOutgoing, nil
	case inAddr != address && outAddr == address:
		return blockatlas.DirectionIncoming, nil
	}

	return "", fmt.Errorf("direction unknown")
}
