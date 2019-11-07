package ethereum

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"math/big"
	"testing"

	"github.com/trustwallet/blockatlas/coin"
)

const tokenTransferSrc = `
{
    "operations": [
        {
            "transactionId": "0x7777854580f273df61e0162e1a41b3e1e05ab8b9f553036fa9329a90dd7e9ab2-0",
            "contract": {
                "address": "0xf3586684107ce0859c44aa2b2e0fb8cd8731a15a",
                "symbol": "KBC",
                "decimals": 7,
                "totalSupply": "120000000000000000",
                "name": "KaratBank Coin"
            },
            "from": "0xd35f30d194684a391c63a6deced7d3dd5207c265",
            "to": "0xaa4d790076f1bf7511a0a0ac498c89e13e1efe17",
            "type": "token_transfer",
            "value": "4291000000",
            "coin": 60
        }
    ],
    "contract": null,
    "_id": "0x7777854580f273df61e0162e1a41b3e1e05ab8b9f553036fa9329a90dd7e9ab2",
    "blockNumber": 7491945,
    "time": 1554248437,
    "nonce": 88,
    "from": "0xd35f30d194684a391c63a6deced7d3dd5207c265",
    "to": "0xf3586684107ce0859c44aa2b2e0fb8cd8731a15a",
    "value": "0",
    "gas": "67497",
    "gasPrice": "6900000256",
    "gasUsed": "51921",
    "input": "0xa9059cbb000000000000000000000000aa4d790076f1bf7511a0a0ac498c89e13e1efe1700000000000000000000000000000000000000000000000000000000ffc376c0",
    "error": "",
    "id": "0x7777854580f273df61e0162e1a41b3e1e05ab8b9f553036fa9329a90dd7e9ab2",
    "coin": 60
}`

const contractCallSrc = `
{
	"addresses": [
		"0x09862ed5908c0a336f9f92e5ffeb9768deac6091"
	],
	"operations": [],
	"contract": "0xe4dc0f23b1a3f2c47dc288a22f72e100e9b1cd70",
	"_id": "0x34ab0028a9aa794d5cc12887e7b813cec17889948276b301028f24a408da6da4",
	"blockNumber": 7522627,
	"time": 1554661737,
	"nonce": 534,
	"from": "0xc9a16a82c284efc3cb0fe8c891ab85d6eba0eefb",
	"to": "0xc67f9c909c4d185e4a5d21d642c27d05a145a76c",
	"value": "1800000000000000000",
	"gas": "1000000",
	"gasPrice": "2000000000",
	"gasUsed": "21340",
	"input": "0xfffdefefed",
	"error": "",
	"id": "0x34ab0028a9aa794d5cc12887e7b813cec17889948276b301028f24a408da6da4",
	"coin": 60
}
`

const transferSrc = `
{
	"operations": [],
	"contract": null,
	"_id": "0x77f8a3b2203933493d103a1637de814b4853410b1fb2981c4d2cff4d7a3071ab",
	"blockNumber": 7522781,
	"time": 1554663642,
	"nonce": 88,
	"from": "0xf5aea47e57c058881b31ee8fce1002c409188f06",
	"to": "0x0ae933a89d9e249d0873cfc7ca022fcb3f1280ce",
	"value": "1999895000000000000",
	"gas": "21000",
	"gasPrice": "5000000000",
	"gasUsed": "21000",
	"input": "0x",
	"error": "",
	"id": "0x77f8a3b2203933493d103a1637de814b4853410b1fb2981c4d2cff4d7a3071ab",
	"coin": 60
}`

const failedSrc = `
{
	"operations": [],
	"contract": null,
	"_id": "0x8dfe7e859f7bdcea4e6f4ada18567d96a51c3aa29e618ef09b80ae99385e191e",
	"blockNumber": 7522678,
	"time": 1554662399,
	"nonce": 1,
	"from": "0x4b55af7ce28a113d794f9a9940fe1506f37aa619",
	"to": "0xe65f787c8561a4b15771111bb427274dedfe85d7",
	"value": "59859820000000000",
	"gas": "21000",
	"gasPrice": "3000000000",
	"gasUsed": "21000",
	"input": "0x",
	"error": "Error",
	"id": "0x8dfe7e859f7bdcea4e6f4ada18567d96a51c3aa29e618ef09b80ae99385e191e",
	"coin": 60
}`

var tokenTransferDst = blockatlas.Tx{
	ID:       "0x7777854580f273df61e0162e1a41b3e1e05ab8b9f553036fa9329a90dd7e9ab2",
	Coin:     coin.ETH,
	From:     "0xd35f30d194684a391c63a6deced7d3dd5207c265",
	To:       "0xf3586684107ce0859c44aa2b2e0fb8cd8731a15a",
	Fee:      "358254913291776",
	Date:     1554248437,
	Block:    7491945,
	Sequence: 88,
	Status:   blockatlas.StatusCompleted,
	Meta: blockatlas.TokenTransfer{
		Name:     "KaratBank Coin",
		Symbol:   "KBC",
		TokenID:  "0xf3586684107CE0859c44aa2b2E0fB8cd8731a15a",
		Decimals: 7,
		Value:    "4291000000",
		From:     "0xd35f30d194684a391c63a6deced7d3dd5207c265",
		To:       "0xaa4d790076f1bf7511a0a0ac498c89e13e1efe17",
	},
}

var contractCallDst = blockatlas.Tx{
	ID:       "0x34ab0028a9aa794d5cc12887e7b813cec17889948276b301028f24a408da6da4",
	Coin:     coin.ETH,
	From:     "0xc9a16a82c284efc3cb0fe8c891ab85d6eba0eefb",
	To:       "0xc67f9c909c4d185e4a5d21d642c27d05a145a76c",
	Fee:      "42680000000000",
	Date:     1554661737,
	Block:    7522627,
	Sequence: 534,
	Status:   blockatlas.StatusCompleted,
	Meta: blockatlas.ContractCall{
		Input: "0xfffdefefed",
		Value: "1800000000000000000",
	},
}

var transferDst = blockatlas.Tx{
	ID:       "0x77f8a3b2203933493d103a1637de814b4853410b1fb2981c4d2cff4d7a3071ab",
	Coin:     coin.ETH,
	From:     "0xf5aea47e57c058881b31ee8fce1002c409188f06",
	To:       "0x0ae933a89d9e249d0873cfc7ca022fcb3f1280ce",
	Fee:      "105000000000000",
	Date:     1554663642,
	Block:    7522781,
	Sequence: 88,
	Status:   blockatlas.StatusCompleted,
	Meta: blockatlas.Transfer{
		Value:    "1999895000000000000",
		Symbol:   "ETH",
		Decimals: 18,
	},
}

var transferContractDst = blockatlas.Tx{
	ID:       "0x77f8a3b2203933493d103a1637de814b4853410b1fb2981c4d2cff4d7a3071ab",
	Coin:     coin.ETH,
	From:     "0xf5aea47e57c058881b31ee8fce1002c409188f06",
	To:       "0x0ae933a89d9e249d0873cfc7ca022fcb3f1280ce",
	Fee:      "105000000000000",
	Date:     1554663642,
	Block:    7522781,
	Sequence: 88,
	Status:   blockatlas.StatusCompleted,
	Meta: blockatlas.Transfer{
		Value:    "1999895000000000000",
		Symbol:   "ETH",
		Decimals: 18,
	},
}

var failedDst = blockatlas.Tx{
	ID:       "0x8dfe7e859f7bdcea4e6f4ada18567d96a51c3aa29e618ef09b80ae99385e191e",
	Coin:     coin.ETH,
	From:     "0x4b55af7ce28a113d794f9a9940fe1506f37aa619",
	To:       "0xe65f787c8561a4b15771111bb427274dedfe85d7",
	Fee:      "63000000000000",
	Date:     1554662399,
	Block:    7522678,
	Sequence: 1,
	Status:   blockatlas.StatusFailed,
	Error:    "Error",
	Meta: blockatlas.Transfer{
		Value:    "59859820000000000",
		Symbol:   "ETH",
		Decimals: 18,
	},
}

type test struct {
	name        string
	apiResponse string
	expected    *blockatlas.Tx
	token       bool
}

func TestNormalize(t *testing.T) {
	testNormalize(t, &test{
		name:        "transfer",
		apiResponse: transferSrc,
		expected:    &transferDst,
	})
	testNormalize(t, &test{
		name:        "token transfer",
		apiResponse: tokenTransferSrc,
		expected:    &tokenTransferDst,
		token:       true,
	})
	testNormalize(t, &test{
		name:        "contract call",
		apiResponse: contractCallSrc,
		expected:    &contractCallDst,
	})
	testNormalize(t, &test{
		name:        "failed transaction",
		apiResponse: failedSrc,
		expected:    &failedDst,
	})
}

func testNormalize(t *testing.T, _test *test) {
	var doc Doc
	err := json.Unmarshal([]byte(_test.apiResponse), &doc)
	if err != nil {
		t.Error(err)
		return
	}
	res := AppendTxs(nil, &doc, coin.ETH)

	resJSON, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}

	dstJSON, err := json.Marshal([]blockatlas.Tx{*_test.expected})
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(resJSON, dstJSON) {
		println(string(resJSON))
		println(string(dstJSON))
		t.Error(_test.name + ": tx don't equal")
	}
}

const tokenSrc = `
{
	"balance": "0",
	"contract": {
		"contract": "0xa14839c9837657efcde754ebeaf5cbecdd801b2a",
		"address": "0xa14839c9837657efcde754ebeaf5cbecdd801b2a",
		"name": "FusChain",
		"decimals": 18,
		"symbol": "FUS"
	}
}`

var tokenDst = blockatlas.Token{
	Name:     "FusChain",
	Symbol:   "FUS",
	Decimals: 18,
	TokenID:  "0xa14839c9837657efcde754ebeaf5cbecdd801b2a",
	Coin:     coin.ETH,
	Type:     blockatlas.TokenTypeERC20,
}

type testToken struct {
	name        string
	apiResponse string
	expected    *blockatlas.Token
	token       bool
}

func TestNormalizeToken(t *testing.T) {
	testNormalizeToken(t, &testToken{
		name:        "token",
		apiResponse: tokenSrc,
		expected:    &tokenDst,
	})
}

func testNormalizeToken(t *testing.T, _test *testToken) {
	var token Token
	err := json.Unmarshal([]byte(_test.apiResponse), &token)
	if err != nil {
		t.Error(err)
		return
	}
	tk, ok := NormalizeToken(&token, coin.ETH)
	if !ok {
		t.Errorf("token: token could not be normalized")
	}

	resJSON, err := json.Marshal(&tk)
	if err != nil {
		t.Fatal(err)
	}

	dstJSON, err := json.Marshal(_test.expected)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(resJSON, dstJSON) {
		println(string(resJSON))
		println(string(dstJSON))
		t.Error("token: token don't equal")
	}
}

const collectionsOwner = "0x0875BCab22dE3d02402bc38aEe4104e1239374a7"

const collectionsSrc = `
[
  {
    "primary_asset_contracts": [
      {
        "address": "0x06012c8cf97bead5deae237070f9587f8e7a266d",
        "name": "CryptoKitties",
        "symbol": "CKITTY",
        "description": "CryptoKitties is a game centered around breedable, collectible, and oh-so-adorable creatures we call CryptoKitties! Each cat is one-of-a-kind and 100% owned by you; it cannot be replicated, taken away, or destroyed.",
        "external_link": "https://www.cryptokitties.co/",
        "nft_version": "1.0",
        "schema_name": "ERC721",
        "display_data": {
          "images": [
            "https://storage.googleapis.com/ck-kitty-image/0x06012c8cf97bead5deae237070f9587f8e7a266d/564155.svg",
            "https://storage.googleapis.com/ck-kitty-image/0x06012c8cf97bead5deae237070f9587f8e7a266d/546630.svg",
            "https://storage.googleapis.com/ck-kitty-image/0x06012c8cf97bead5deae237070f9587f8e7a266d/441529.svg",
            "https://storage.googleapis.com/ck-kitty-image/0x06012c8cf97bead5deae237070f9587f8e7a266d/552435.svg",
            "https://storage.googleapis.com/ck-kitty-image/0x06012c8cf97bead5deae237070f9587f8e7a266d/524748.png",
            "https://storage.googleapis.com/ck-kitty-image/0x06012c8cf97bead5deae237070f9587f8e7a266d/540800.svg"
          ],
          "card_display_style": "padded"
        },
        "image_url": "https://storage.opensea.io/0x06012c8cf97bead5deae237070f9587f8e7a266d-featured-1556588705.png"
      }
    ],
    "name": "CryptoKitties",
    "slug": "cryptokitties",
    "image_url": "https://storage.opensea.io/0x06012c8cf97bead5deae237070f9587f8e7a266d-featured-1556588705.png",
    "description": "CryptoKitties is a game centered around breedable, collectible, and oh-so-adorable creatures we call CryptoKitties! Each cat is one-of-a-kind and 100% owned by you; it cannot be replicated, taken away, or destroyed.",
    "external_url": "https://www.cryptokitties.co/",
    "featured_image_url": "https://storage.opensea.io/0x06012c8cf97bead5deae237070f9587f8e7a266d-featured-1556589429.png",
    "created_date": "2019-04-26T22:13:04.207050",
    "owned_asset_count": 3
  },
  {
    "primary_asset_contracts": [
      {
        "address": "0xf629cbd94d3791c9250152bd8dfbdf380e2a3b9c",
        "name": "Coin",
        "symbol": "",
        "description": null,
        "external_link": null,
        "nft_version": null,
        "schema_name": "ERC20",
        "display_data": {}
      }
    ],
    "name": "Enjin Token",
    "slug": "enjin-token",
    "image_url": "https://storage.googleapis.com/opensea-static/tokens-high-res/ENJ.png",
    "description": "This is the collection of owners of EnjinCoin",
    "external_url": null,
    "owned_asset_count": 20000000000000000000
  },
  {
    "primary_asset_contracts": [
      {
        "address": "0xfaafdc07907ff5120a76b34b731b278c38d6043c",
        "name": "Enjin",
        "symbol": "",
        "description": "",
        "external_link": null,
        "nft_version": null,
        "schema_name": "ERC1155",
        "display_data": {},
        "owner": null,
        "created_date": "2019-08-02T23:43:14.666153",
        "asset_contract_type": "semi-fungible"
      }
    ],
    "name": "Age of Rust",
    "slug": "age-of-rust",
    "image_url": "https://storage.opensea.io/age-of-rust-1561960816.jpg",
    "description": "Year 4424: The search begins for new life on the other side of the galaxy. Explore abandoned space stations, mysterious caverns, and ruins on far away worlds in order to unlock puzzles and secrets! Beware the rogue machines!",
    "external_url": "https://www.ageofrust.games/",
    "featured_image_url": null,
    "created_date": "2019-09-03T02:35:56.063685",
    "owned_asset_count": 1
  }
]
`

var collection1Dst = blockatlas.Collection{
	Name:            "CryptoKitties",
	Symbol:          "CKITTY",
	Slug:            "cryptokitties",
	ImageUrl:        "https://storage.opensea.io/0x06012c8cf97bead5deae237070f9587f8e7a266d-featured-1556588705.png",
	Description:     "CryptoKitties is a game centered around breedable, collectible, and oh-so-adorable creatures we call CryptoKitties! Each cat is one-of-a-kind and 100% owned by you; it cannot be replicated, taken away, or destroyed.",
	ExternalLink:    "https://www.cryptokitties.co/",
	Total:           3,
	CategoryAddress: "0x06012c8cf97bead5deae237070f9587f8e7a266d",
	Id:              "0x06012c8cf97bead5deae237070f9587f8e7a266d",
	Address:         "0x0875BCab22dE3d02402bc38aEe4104e1239374a7",
	Version:         "1.0",
	Coin:            60,
	Type:            "ERC721",
}

var collection2Dst = blockatlas.Collection{
	Name:            "Age of Rust",
	Symbol:          "",
	Slug:            "age-of-rust",
	ImageUrl:        "https://storage.opensea.io/age-of-rust-1561960816.jpg",
	Description:     "Year 4424: The search begins for new life on the other side of the galaxy. Explore abandoned space stations, mysterious caverns, and ruins on far away worlds in order to unlock puzzles and secrets! Beware the rogue machines!",
	ExternalLink:    "https://www.ageofrust.games/",
	Total:           1,
	CategoryAddress: "0xfaafdc07907ff5120a76b34b731b278c38d6043c---age-of-rust",
	Id:              "0xfaafdc07907ff5120a76b34b731b278c38d6043c---age-of-rust",
	Address:         "0x0875BCab22dE3d02402bc38aEe4104e1239374a7",
	Version:         "",
	Coin:            60,
	Type:            "ERC1155",
}

func TestNormalizeCollection(t *testing.T) {
	var collections []Collection
	err := json.Unmarshal([]byte(collectionsSrc), &collections)
	assert.Nil(t, err)
	page := NormalizeCollectionPage(collections, coin.ETH, collectionsOwner)
	assert.Equal(t, len(page), 2, "collections could not be normalized")
	expected := blockatlas.CollectionPage{collection1Dst, collection2Dst}
	assert.Equal(t, page, expected, "collections don't equal")
}

const collectibleSrc = `
[
  {
    "token_id": "54277541829991970107421667568590323026590803461896874578610080514640537714688",
    "image_url": "https://storage.opensea.io/0xfaafdc07907ff5120a76b34b731b278c38d6043c/54277541829991970107421667568590323026590803461896874578610080514640537714688-1564858806.png",
    "image_preview_url": "https://storage.opensea.io/0xfaafdc07907ff5120a76b34b731b278c38d6043c-preview/54277541829991970107421667568590323026590803461896874578610080514640537714688-1564858807.png",
    "name": "Rustbits",
    "description": "Rustbits are the main token of use within the Age of Rust game universe. You need Rustbits to not only play Age of Rust, but also to purchase in-game cryptoitems as well. Rustbits are radioactive rust scraped off of hulls of abandoned ships that are in orbit around a hidden planet, which is also a gas giant. The planet is so radioactive, it damages ships and kills anyone that gets close to it. Getting bits of rust off of ships is highly rare and prized.",
    "external_link": "",
    "asset_contract": {
      "address": "0xfaafdc07907ff5120a76b34b731b278c38d6043c",
      "name": "Enjin",
      "external_link": null,
      "nft_version": null,
      "schema_name": "ERC1155",
	  "display_data": {}
    },
    "permalink": "https://opensea.io/assets/0xfaafdc07907ff5120a76b34b731b278c38d6043c/54277541829991970107421667568590323026590803461896874578610080514640537714688"
  }
]
`

var collectibleCollectionDst = Collection{
	Name:        "Age of Rust",
	ImageUrl:    "https://storage.opensea.io/age-of-rust-1561960816.jpg",
	Description: "Year 4424: The search begins for new life on the other side of the galaxy. Explore abandoned space stations, mysterious caverns, and ruins on far away worlds in order to unlock puzzles and secrets! Beware the rogue machines!",
	ExternalUrl: "https://www.ageofrust.games/",
	Total:       big.NewInt(1),
	Slug:        "age-of-rust",
	Contracts: []PrimaryAssetContract{
		{
			Name:        "Age of Rust",
			Address:     "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
			NftVersion:  "",
			Symbol:      "",
			Description: "",
			Type:        "ERC1155",
			Url:         "",
		},
	},
}

var collectibleDst = blockatlas.Collectible{
	CollectionID:     "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB---age-of-rust",
	TokenID:          "0xfaafdc07907ff5120a76b34b731b278c38d6043c-54277541829991970107421667568590323026590803461896874578610080514640537714688",
	CategoryContract: "0xfaafdc07907ff5120a76b34b731b278c38d6043c",
	ContractAddress:  "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
	Category:         "Age of Rust",
	ImageUrl:         "https://storage.opensea.io/0xfaafdc07907ff5120a76b34b731b278c38d6043c-preview/54277541829991970107421667568590323026590803461896874578610080514640537714688-1564858807.png",
	ExternalLink:     "",
	ProviderLink:     "https://opensea.io/assets/0xfaafdc07907ff5120a76b34b731b278c38d6043c/54277541829991970107421667568590323026590803461896874578610080514640537714688",
	Type:             "ERC1155",
	Description:      "Rustbits are the main token of use within the Age of Rust game universe. You need Rustbits to not only play Age of Rust, but also to purchase in-game cryptoitems as well. Rustbits are radioactive rust scraped off of hulls of abandoned ships that are in orbit around a hidden planet, which is also a gas giant. The planet is so radioactive, it damages ships and kills anyone that gets close to it. Getting bits of rust off of ships is highly rare and prized.",
	Coin:             60,
	Name:             "Rustbits",
}

func TestNormalizeCollectible(t *testing.T) {
	var collectible []Collectible
	err := json.Unmarshal([]byte(collectibleSrc), &collectible)
	assert.Nil(t, err)
	page := NormalizeCollectiblePage(&collectibleCollectionDst, collectible, coin.ETH)
	assert.Equal(t, len(page), 1, "collectible could not be normalized")
	expected := blockatlas.CollectiblePage{collectibleDst}
	assert.Equal(t, page, expected, "collectible don't equal")
}
