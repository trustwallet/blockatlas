package ethereum

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas"
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
        "image_url": null,
        "large_image_url": null,
        "featured_image_url": null,
        "featured": false,
        "description": "",
        "external_link": null,
        "wiki_link": null,
        "hidden": true,
        "nft_version": null,
        "schema_name": "ERC1155",
        "display_data": {},
        "short_description": null,
        "total_supply": null,
        "owner": null,
        "dev_buyer_fee_basis_points": 0,
        "dev_seller_fee_basis_points": 0,
        "opensea_buyer_fee_basis_points": 0,
        "opensea_seller_fee_basis_points": 250,
        "buyer_fee_basis_points": 0,
        "seller_fee_basis_points": 250,
        "payout_address": null,
        "require_email": false,
        "require_whitelist": false,
        "only_proxied_transfers": false,
        "default_to_fiat": false,
        "created_date": "2019-08-02T23:43:14.666153",
        "opensea_version": null,
        "asset_contract_type": "semi-fungible"
      }
    ],
    "traits": {
      "Stage": {
        "min": 1,
        "max": 1
      },
      "Experience": {
        "min": 0,
        "max": 0
      },
      "property": {
        "min": 1,
        "max": 250
      }
    },
    "name": "Enjin",
    "slug": "enjin",
    "image_url": "https://storage.opensea.io/0x8562c38485b1e8ccd82e44f89823da76c98eb0ab-featured-1556588805.png",
    "short_description": null,
    "description": "Enjin assets are unique digital ERC1155 assets used in a variety of games in the Enjin multiverse.",
    "external_url": "https://enj1155.com",
    "chat_url": null,
    "wiki_url": null,
    "large_image_url": null,
    "featured_image_url": null,
    "featured": false,
    "banner_image_url": null,
    "display_data": {
      "card_display_style": "contain",
      "images": [
        "https://storage.opensea.io/0x8562c38485b1e8ccd82e44f89823da76c98eb0ab/11081664790290029703914123501093412959851005990698421485335279882236843786250-1550517150.png",
        "https://storage.opensea.io/0x8562c38485b1e8ccd82e44f89823da76c98eb0ab/11081664790290028134638689654423222000903650188781817459746418766228215563559-1552773943.png",
        "https://storage.opensea.io/0x8562c38485b1e8ccd82e44f89823da76c98eb0ab/50885195465617471177731255234521635734804147820253726770435014604985116655636-1551524003.jpg",
        "https://storage.opensea.io/0x8562c38485b1e8ccd82e44f89823da76c98eb0ab/11081664790290028178578401802129987347754176151235482372462906877476457152527-1552773783.png",
        "https://storage.opensea.io/0x8562c38485b1e8ccd82e44f89823da76c98eb0ab/7237005577332265873523498293477879557094607771672055840139323123026691620864-1551250671.jpg",
        "https://storage.opensea.io/0x8562c38485b1e8ccd82e44f89823da76c98eb0ab/10855508365998398882471917397163648119753490024396247545386072295876434329600-1552747248.jpg"
      ]
    },
    "hidden": false,
    "created_date": "2019-08-02T23:43:14.650920",
    "owned_asset_count": 1
  }
]
`

var collectionsDst = blockatlas.Collection{
	Name:            "Enjin",
	Symbol:          "",
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

func TestNormalizeCollection(t *testing.T) {
	var collections []Collection
	err := json.Unmarshal([]byte(collectionsSrc), &collections)
	assert.Nil(t, err)
	page := NormalizeCollectionPage(collections, coin.ETH, collectionsOwner)
	assert.Equal(t, len(page), 1, "collections could not be normalized")
	expected := blockatlas.CollectionPage{collectionsDst}
	assert.Equal(t, page, expected, "collections don't equal")
}
