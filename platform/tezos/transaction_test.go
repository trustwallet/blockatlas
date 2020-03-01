package tezos

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

const transferSrc1 = `
{
  "tx": {
    "destination": "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK",
    "amount": "1560",
    "gasLimit": "15385",
    "kind": "transaction",
    "operationResultStatus": "applied",
    "blockHash": "BMJYDJk9wRpxQuuqcFS7MZivqShtrgG18eY5K6rSDBpx5vcJLB2",
    "fee": "0",
    "operationResultConsumedGas": "10207",
    "counter": "1383819",
    "blockLevel": 791441,
    "operationResultErrors": null,
    "blockTimestamp": "2020-01-22T16:34:22Z",
    "parameters": null,
    "source": "tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93",
    "insertedTimestamp": "2020-01-22 16:34:56.015281 UTC"
  },
  "op": {
    "opHash": "ooBAC2ynR5LfU9L2KEon8Z3ujmwEVyB9si1rrppBCmmEn4Mr3bJ",
    "chainId": "NetXdQprcVkpaWU",
    "blockHash": "BMJYDJk9wRpxQuuqcFS7MZivqShtrgG18eY5K6rSDBpx5vcJLB2",
    "blockLevel": 791441,
    "blockTimestamp": "2020-01-22T16:34:22Z",
    "insertedTimestamp": "2020-01-22 16:34:49.405793 UTC"
  }
}`

const transferSrc2 = `
{
  "tx": {
    "destination": "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK",
    "amount": "1751",
    "gasLimit": "15385",
    "kind": "transaction",
    "operationResultStatus": "applied",
    "blockHash": "BLJKc6f6SpFs3Jr7WMp2ekx5jQQyzHWN6SWHDo2AJ41HJoPKTF2",
    "fee": "0",
    "operationResultConsumedGas": "10207",
    "counter": "1383094",
    "blockLevel": 788725,
    "operationResultErrors": null,
    "blockTimestamp": "2020-01-20T18:54:52Z",
    "parameters": null,
    "source": "tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93",
    "insertedTimestamp": "2020-01-20 18:56:34.828193 UTC"
  },
  "op": {
    "opHash": "ookuQhVYopxrg8FtnfNASxpMhmnNhBqYaK3PJQXDpP7sDCJAZwf",
    "chainId": "NetXdQprcVkpaWU",
    "blockHash": "BLJKc6f6SpFs3Jr7WMp2ekx5jQQyzHWN6SWHDo2AJ41HJoPKTF2",
    "blockLevel": 788725,
    "blockTimestamp": "2020-01-20T18:54:52Z",
    "insertedTimestamp": "2020-01-20 18:56:28.193751 UTC"
  }
}`

const transferSrc3 = `
{
  "tx": {
    "destination": "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK",
    "amount": "1751",
    "gasLimit": "15385",
    "kind": "transaction",
    "operationResultStatus": "backtracked",
    "blockHash": "BMDYrXJ7GSwztzsy3ykJb43VXciNk1WY8EAaSoGbcUE7mA7HUzj",
    "fee": "0",
    "operationResultConsumedGas": "10207",
    "counter": "1382930",
    "blockLevel": 788568,
    "operationResultErrors": null,
    "blockTimestamp": "2020-01-20T16:16:32Z",
    "parameters": null,
    "source": "tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93",
    "insertedTimestamp": "2020-01-20 16:19:06.938114 UTC"
  },
  "op": {
    "opHash": "ooXN845juCMcQuqeodwBJWNhY17A5HKyWRxbcwS1m6TfqCqjM8q",
    "chainId": "NetXdQprcVkpaWU",
    "blockHash": "BMDYrXJ7GSwztzsy3ykJb43VXciNk1WY8EAaSoGbcUE7mA7HUzj",
    "blockLevel": 788568,
    "blockTimestamp": "2020-01-20T16:16:32Z",
    "insertedTimestamp": "2020-01-20 16:18:59.855515 UTC"
  }
}`

const transferSrc4 = `
{
  "op": {
    "signature": "sigjQxDjvom9zqhNbi4cbwPLXn4fnVHiYSwsxDyXdPiFCjDqgS2ukHeWqmq61jkboiXTE3UJUqiXGqtmDVDXTRr5EKhaAeWF",
    "blockUuid": "0ce492ed-ea9e-43c2-bf14-f17c44f099a4",
    "opHash": "opUx5394wcy8BdoaYCM32kVyLB27F9ruT4vnKQccgQ463KsFPy3",
    "uuid": "d5e2df82-9149-4d3e-9cac-16a425c79f67",
    "chainId": "NetXdQprcVkpaWU",
    "blockHash": "BMCkZ9NM5v3LnRPgfXYd4F9DHQ3bdTmmE8KjWhA5X982Bfz7ueC",
    "protocol": "PsBabyM1eUXZseaJdmXFApDSBqj8YBfwELoxZHHW77EMcAbbwAS",
    "branch": "BKybkjYtw3Cv7zdECiQd9W9WmWkoa25eJLwFiVQ4S3wB5mdcw4i",
    "blockLevel": 809102,
    "blockTimestamp": "2020-02-04T00:29:14Z",
    "insertedTimestamp": "2020-02-04 00:29:28.291063 UTC"
  },
  "endor": {
    "slots": [
      1
    ],
    "delegate": "tz1WCd2jm4uSt4vntk4vSuUWoZQGhLcDuR9q",
    "opUuid": "d5e2df82-9149-4d3e-9cac-16a425c79f67",
    "uuid": "7e1d6f45-b148-4e24-9c36-d7a0deb6ff72",
    "kind": "endorsement",
    "blockHash": "BMCkZ9NM5v3LnRPgfXYd4F9DHQ3bdTmmE8KjWhA5X982Bfz7ueC",
    "blockLevel": 809102,
    "blockTimestamp": "2020-02-04T00:29:14Z",
    "insertedTimestamp": "2020-02-04 00:29:31.168224 UTC",
    "metadataUuid": "36380127-30af-4bd0-88dd-18b49860d43c",
    "level": 809101
  }
}`

const delegationsSrc1 = `
{
  "op": {
    "signature": "sigvHd2YBByFXU8nL4CZKSTYXNdMapMsJw1f239YRRjgz9NvrTyA6iGnpBDhi9kCB4zMHysrg9H4jxcpPH975WiQtEmkMjb5",
    "blockUuid": "4b292c55-41ba-4383-a1d6-03fb71b88f41",
    "opHash": "opGphHGNEZZN5rF78yxwe9BJydxYA2yqxECnZR6s6HcxXtCg8Sj",
    "uuid": "e4ec0e07-1601-4da3-bd92-090a820ed369",
    "chainId": "NetXdQprcVkpaWU",
    "blockHash": "BLkscXpE63gajVzmgBS7fQx63hERKQRCZFGtMXdYY6WPThHyji7",
    "protocol": "PsBabyM1eUXZseaJdmXFApDSBqj8YBfwELoxZHHW77EMcAbbwAS",
    "branch": "BKqtLegZfdPR3USyYYcMpedB59W5eUBuFZAVpVMPpFgEvMcZjr1",
    "blockLevel": 791778,
    "blockTimestamp": "2020-01-22T22:13:38Z",
    "insertedTimestamp": "2020-01-22 22:14:05.937406 UTC"
  },
  "delegation": {
    "storageLimit": "257",
    "delegate": "tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93",
    "opUuid": "e4ec0e07-1601-4da3-bd92-090a820ed369",
    "uuid": "6459fcd9-5eee-4999-ac4d-92330b9eaab3",
    "gasLimit": "10600",
    "kind": "delegation",
    "operationResultStatus": "applied",
    "fee": "1500",
    "operationResultUuid": "791f6ec7-ecec-43d5-82ca-a1497be0188c",
    "operationResultConsumedGas": "10000",
    "counter": "2409130",
    "operationResultErrors": null,
    "source": "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK",
    "insertedTimestamp": "2020-01-22 22:15:12.038586 UTC",
    "metadataUuid": "85b42f50-1a89-421c-bcb0-a06926941bc4"
  }
}`

const delegationsSrc2 = `
{
  "op": {
    "signature": "sigbhvzA37wDGhfYT21LQyoYfzU4KvhbaTwrLMXMDMDUJt69HNAxSK44wv8S1576ZADcEwMLNUbSXyBEFgFkikbNRn51ozhg",
    "blockUuid": "785389b9-356d-4a98-9202-39f9c94b5b73",
    "opHash": "onfvhheWth5GCbKKAc9KLZoqQ9AFzAS27vfPCN2DGcGuPA31jmc",
    "uuid": "766932d0-fc9b-451d-834d-5feb48178580",
    "chainId": "NetXdQprcVkpaWU",
    "blockHash": "BMQShgncuUDtY2WbXz9aEqebKfztSL3VxVzgVRoMweb72fzS9MD",
    "protocol": "PsBabyM1eUXZseaJdmXFApDSBqj8YBfwELoxZHHW77EMcAbbwAS",
    "branch": "BLzBVG42EMvipNaWEPPTW9G6otFqwAEhSjsgmDMSNrVKComW4Tb",
    "blockLevel": 809061,
    "blockTimestamp": "2020-02-03T23:48:14Z",
    "insertedTimestamp": "2020-02-03 23:48:43.650483 UTC"
  },
  "delegation": {
    "storageLimit": "257",
    "delegate": "tz1V3yg82mcrPJbegqVCPn6bC8w1CSTRp3f8",
    "opUuid": "766932d0-fc9b-451d-834d-5feb48178580",
    "uuid": "ad5b8aca-86ae-4463-82b1-a05350c6899b",
    "gasLimit": "10600",
    "kind": "delegation",
    "operationResultStatus": "applied",
    "fee": "1500",
    "operationResultUuid": "a87226fd-4983-49d7-b65f-93378056fda7",
    "operationResultConsumedGas": "10000",
    "counter": "3032194",
    "operationResultErrors": null,
    "source": "tz1hpUTmafsAEH8vQTirnZsvPWZ3wZ6oEkUw",
    "insertedTimestamp": "2020-02-03 23:49:59.947433 UTC",
    "metadataUuid": "325bfab8-6db3-4f6e-beb3-783031c556de"
  }
}`

const delegationsSrc3 = `
{
  "op": {
    "signature": "signaVwXLttpBqZVKUwbdUN92TevU26CRUWvYBx5areFAb4uAaeKDVFacWPfLc2WwfpMhXgbC4PpwgMuZd1RfXNJF7w7T7Fn",
    "blockUuid": "ac319dd4-78ee-44c4-8652-6bf5be68a666",
    "opHash": "oneqVXTkJqDvCDfxqggnzveo6U28Bp5sByguaqSZ7seBgQgzZSG",
    "uuid": "c3bab716-ed01-4357-8f41-b75df6ea1708",
    "chainId": "NetXdQprcVkpaWU",
    "blockHash": "BLSN7G7UBSUys7njxo8Xm699L4U3xhhbkZyYYr5e8jJrDtnqfj8",
    "protocol": "PsBabyM1eUXZseaJdmXFApDSBqj8YBfwELoxZHHW77EMcAbbwAS",
    "branch": "BLAWNqLHLQWfSLwgQE8gqUWFPx3DPzxN2jrUxtz9q1zvisaCUwi",
    "blockLevel": 809007,
    "blockTimestamp": "2020-02-03T22:53:34Z",
    "insertedTimestamp": "2020-02-03 22:53:48.734959 UTC"
  },
  "delegation": {
    "storageLimit": "257",
    "delegate": null,
    "opUuid": "c3bab716-ed01-4357-8f41-b75df6ea1708",
    "uuid": "979e75b2-a64c-4844-a187-d5e59a3c7728",
    "gasLimit": "10600",
    "kind": "delegation",
    "operationResultStatus": "applied",
    "fee": "1500",
    "operationResultUuid": "a4c26049-8169-4804-b074-d201fac4bfd2",
    "operationResultConsumedGas": "10000",
    "counter": "3032193",
    "operationResultErrors": null,
    "source": "tz1hpUTmafsAEH8vQTirnZsvPWZ3wZ6oEkUw",
    "insertedTimestamp": "2020-02-03 22:54:50.993844 UTC",
    "metadataUuid": "edd476c1-88e0-40d0-ac29-2c32e6d16feb"
  }
}`

var transfer1 = blockatlas.Tx{
	ID:    "ooBAC2ynR5LfU9L2KEon8Z3ujmwEVyB9si1rrppBCmmEn4Mr3bJ",
	Coin:  coin.XTZ,
	Date:  1579710862,
	From:  "tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93",
	To:    "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK",
	Fee:   "0",
	Block: 791441,
	Meta: blockatlas.Transfer{
		Value:    blockatlas.Amount("1560"),
		Symbol:   coin.Coins[coin.XTZ].Symbol,
		Decimals: coin.Coins[coin.XTZ].Decimals,
	},
	Status: blockatlas.StatusCompleted,
}

var transfer2 = blockatlas.Tx{
	ID:    "ookuQhVYopxrg8FtnfNASxpMhmnNhBqYaK3PJQXDpP7sDCJAZwf",
	Coin:  coin.XTZ,
	Date:  1579546492,
	From:  "tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93",
	To:    "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK",
	Fee:   "0",
	Block: 788725,
	Meta: blockatlas.Transfer{
		Value:    blockatlas.Amount("1751"),
		Symbol:   coin.Coins[coin.XTZ].Symbol,
		Decimals: coin.Coins[coin.XTZ].Decimals,
	},
	Status: blockatlas.StatusCompleted,
}

var transfer3 = blockatlas.Tx{
	ID:    "ooXN845juCMcQuqeodwBJWNhY17A5HKyWRxbcwS1m6TfqCqjM8q",
	Coin:  coin.XTZ,
	Date:  1579536992,
	From:  "tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93",
	To:    "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK",
	Fee:   "0",
	Block: 788568,
	Meta: blockatlas.Transfer{
		Value:    blockatlas.Amount("1751"),
		Symbol:   coin.Coins[coin.XTZ].Symbol,
		Decimals: coin.Coins[coin.XTZ].Decimals,
	},
	Status: blockatlas.StatusError,
	Error:  "transaction failed",
}

var delegation1 = blockatlas.Tx{
	ID:     "opGphHGNEZZN5rF78yxwe9BJydxYA2yqxECnZR6s6HcxXtCg8Sj",
	Coin:   coin.XTZ,
	Date:   1579731218,
	From:   "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK",
	To:     "tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93",
	Fee:    "1500",
	Block:  791778,
	Status: blockatlas.StatusCompleted,
	Meta: blockatlas.AnyAction{
		Coin:     coin.Tezos().ID,
		Title:    blockatlas.AnyActionDelegation,
		Key:      blockatlas.KeyStakeDelegate,
		Name:     coin.Tezos().Name,
		Symbol:   coin.Tezos().Symbol,
		Decimals: coin.Tezos().Decimals,
	},
}

var delegation2 = blockatlas.Tx{
	ID:     "onfvhheWth5GCbKKAc9KLZoqQ9AFzAS27vfPCN2DGcGuPA31jmc",
	Coin:   coin.XTZ,
	Date:   1580773694,
	From:   "tz1hpUTmafsAEH8vQTirnZsvPWZ3wZ6oEkUw",
	To:     "tz1V3yg82mcrPJbegqVCPn6bC8w1CSTRp3f8",
	Fee:    "1500",
	Block:  809061,
	Status: blockatlas.StatusCompleted,
	Meta: blockatlas.AnyAction{
		Coin:     coin.Tezos().ID,
		Title:    blockatlas.AnyActionDelegation,
		Key:      blockatlas.KeyStakeDelegate,
		Name:     coin.Tezos().Name,
		Symbol:   coin.Tezos().Symbol,
		Decimals: coin.Tezos().Decimals,
	},
}

var delegation3 = blockatlas.Tx{
	ID:     "oneqVXTkJqDvCDfxqggnzveo6U28Bp5sByguaqSZ7seBgQgzZSG",
	Coin:   coin.XTZ,
	Date:   1580770414,
	From:   "tz1hpUTmafsAEH8vQTirnZsvPWZ3wZ6oEkUw",
	To:     "",
	Fee:    "1500",
	Block:  809007,
	Status: blockatlas.StatusCompleted,
	Meta: blockatlas.AnyAction{
		Coin:     coin.Tezos().ID,
		Title:    blockatlas.AnyActionUndelegation,
		Key:      blockatlas.KeyStakeDelegate,
		Name:     coin.Tezos().Name,
		Symbol:   coin.Tezos().Symbol,
		Decimals: coin.Tezos().Decimals,
	},
}

func TestNormalizeTx(t *testing.T) {
	tests := []struct {
		name   string
		srcTx  string
		wantTx blockatlas.Tx
		wantOk bool
	}{
		{"transfer 1", transferSrc1, transfer1, true},
		{"transfer 2", transferSrc2, transfer2, true},
		{"transfer 3", transferSrc3, transfer3, true},
		{"delegation 1", delegationsSrc1, delegation1, true},
		{"delegation 2", delegationsSrc2, delegation2, true},
		{"delegation 3", delegationsSrc3, delegation3, true},
		{"transfer 4", transferSrc4, blockatlas.Tx{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var transaction Transaction
			err := json.Unmarshal([]byte(tt.srcTx), &transaction)
			assert.Nil(t, err)
			gotTx, gotOk := NormalizeTx(transaction)
			assert.Equal(t, tt.wantOk, gotOk, "transfer ok result don't equal")
			assert.Equal(t, tt.wantTx, gotTx, "transfer don't equal")
		})
	}
}
