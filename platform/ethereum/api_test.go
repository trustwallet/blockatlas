package ethereum

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas"
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
    "timeStamp": "1554248437",
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
	"timeStamp": "1554661737",
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
	"timeStamp": "1554663642",
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
	"timeStamp": "1554662399",
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
		TokenID:  "0xf3586684107ce0859c44aa2b2e0fb8cd8731a15a",
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
		Value: "1999895000000000000",
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
		Value: "1999895000000000000",
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
		Value: "59859820000000000",
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
        "address": "0xfaafdc07907ff5120a76b34b731b278c38d6043c",
        "name": "Enjin",
        "symbol": "",
        "description": "",
        "external_link": null,
        "nft_version": null,
        "schema_name": "ERC1155",
        "display_data": {}
      }
    ],
    "name": "Enjin",
    "slug": "enjin",
    "image_url": "https://storage.opensea.io/0x8562c38485b1e8ccd82e44f89823da76c98eb0ab-featured-1556588805.png",
    "description": "Enjin assets are unique digital ERC1155 assets used in a variety of games in the Enjin multiverse.",
    "external_url": "https://enj1155.com",
    "owned_asset_count": 1
  },
  {
    "primary_asset_contracts": [
      {
        "address": "0xf629cbd94d3791c9250152bd8dfbdf380e2a3b9c",
        "name": "EnjinCoin",
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
    "primary_asset_contracts": [],
    "name": "Dissolution",
    "slug": "dissolution",
    "image_url": "https://storage.opensea.io/dissolution-1566503734.png",
    "description": "tactical FPS combat in a cutthroat universe ravaged by an ongoing war of extinction between humanity and AI. Fight for loot backed by blockchain in competitive game modes.",
    "external_url": "https://dissolution.online",
    "owned_asset_count": 20
  }
]
`

var collection1Dst = blockatlas.Collection{
	Name:            "Enjin",
	Symbol:          "",
	Slug:            "enjin",
	ImageUrl:        "https://storage.opensea.io/0x8562c38485b1e8ccd82e44f89823da76c98eb0ab-featured-1556588805.png",
	Description:     "Enjin assets are unique digital ERC1155 assets used in a variety of games in the Enjin multiverse.",
	ExternalLink:    "https://enj1155.com",
	Total:           1,
	CategoryAddress: "0xfaafdc07907ff5120a76b34b731b278c38d6043c",
	Address:         "0x0875BCab22dE3d02402bc38aEe4104e1239374a7",
	Version:         "",
	Coin:            60,
	Type:            "ERC1155",
}

var collection2Dst = blockatlas.Collection{
	Name:            "Dissolution",
	Symbol:          "",
	Slug:            "dissolution",
	ImageUrl:        "https://storage.opensea.io/dissolution-1566503734.png",
	Description:     "tactical FPS combat in a cutthroat universe ravaged by an ongoing war of extinction between humanity and AI. Fight for loot backed by blockchain in competitive game modes.",
	ExternalLink:    "https://dissolution.online",
	Total:           20,
	CategoryAddress: "dissolution",
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
    "token_id": "36185027886661312632864264926498481399258436586721613871817000674017446723584",
    "image_url": "https://storage.opensea.io/0xfaafdc07907ff5120a76b34b731b278c38d6043c/36185027886661312632864264926498481399258436586721613871817000674017446723584-1565973585.jpg",
    "image_preview_url": "https://storage.opensea.io/0xfaafdc07907ff5120a76b34b731b278c38d6043c-preview/36185027886661312632864264926498481399258436586721613871817000674017446723584-1565973586.png",
    "name": "Aeonclipse Key",
    "description": "Forged by unknown, mystical entities at the very beginning of the multiverse, countless Aeonclipse keys were taken by a group of Architects and dispersed through their creations—entire universes. The keys are said to unlock Primythical Chests, legendary vaults hiding immense treasures.",
    "external_link": "",
    "asset_contract": {
      "address": "0xfaafdc07907ff5120a76b34b731b278c38d6043c",
      "name": "Enjin",
      "external_link": null,
      "nft_version": null,
      "schema_name": "ERC1155"
    },
    "permalink": "https://opensea.io/assets/0xfaafdc07907ff5120a76b34b731b278c38d6043c/36185027886661312632864264926498481399258436586721613871817000674017446723584"
  }
]
`

var collectibleCollectionDst = Collection{
	Name:        "Enjin",
	ImageUrl:    "https://storage.opensea.io/0x8562c38485b1e8ccd82e44f89823da76c98eb0ab-featured-1556588805.png",
	Description: "Enjin assets are unique digital ERC1155 assets used in a variety of games in the Enjin multiverse.",
	ExternalUrl: "https://enj1155.com",
	Total:       big.NewInt(1),
	Contracts: []PrimaryAssetContract{
		{
			Name:        "Enjin",
			Address:     "0xfaafdc07907ff5120a76b34b731b278c38d6043c",
			NftVersion:  "",
			Symbol:      "",
			Description: "",
			Type:        "ERC1155",
			Url:         "",
		},
	},
}

var collectibleDst = blockatlas.Collectible{
	CollectionID:     "0xfaafdc07907ff5120a76b34b731b278c38d6043c",
	TokenID:          "36185027886661312632864264926498481399258436586721613871817000674017446723584",
	CategoryContract: "0xfaafdc07907ff5120a76b34b731b278c38d6043c",
	ContractAddress:  "0xfaafdc07907ff5120a76b34b731b278c38d6043c",
	Category:         "Enjin",
	ImageUrl:         "https://storage.opensea.io/0xfaafdc07907ff5120a76b34b731b278c38d6043c-preview/36185027886661312632864264926498481399258436586721613871817000674017446723584-1565973586.png",
	ExternalLink:     "",
	ProviderLink:     "https://opensea.io/assets/0xfaafdc07907ff5120a76b34b731b278c38d6043c/36185027886661312632864264926498481399258436586721613871817000674017446723584",
	Type:             "ERC1155",
	Description:      "Forged by unknown, mystical entities at the very beginning of the multiverse, countless Aeonclipse keys were taken by a group of Architects and dispersed through their creations—entire universes. The keys are said to unlock Primythical Chests, legendary vaults hiding immense treasures.",
	Coin:             60,
	Name:             "Aeonclipse Key",
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
