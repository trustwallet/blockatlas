package ethereum

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"math/big"
	"testing"
)

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
	CategoryAddress: "cryptokitties",
	Id:              "cryptokitties",
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
	CategoryAddress: "age-of-rust",
	Id:              "age-of-rust",
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
	ID:               "0xfaafdc07907ff5120a76b34b731b278c38d6043c-54277541829991970107421667568590323026590803461896874578610080514640537714688",
	CollectionID:     "age-of-rust",
	TokenID:          "54277541829991970107421667568590323026590803461896874578610080514640537714688",
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

var c1 = Collection{
	Slug: "enjin",
	Contracts: []PrimaryAssetContract{{
		Address: "0xfaafdc07907ff5120a76b34b731b278c38d6043c",
	}},
}
var c2 = Collection{
	Slug: "cryptokitties",
	Contracts: []PrimaryAssetContract{{
		Address: "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
	}},
}
var c3 = Collection{
	Slug: "age-of-rust",
	Contracts: []PrimaryAssetContract{{
		Address: "0x0875BCab22dE3d02402bc38aEe4104e1239374a7",
	}},
}

func TestSearchCollection(t *testing.T) {
	var tests = []struct {
		collections   []Collection
		collectibleID string
		result        *Collection
	}{
		{[]Collection{c1, c2, c3}, "enjin", &c1},
		{[]Collection{c1, c2, c3}, "cryptokitties", &c2},
		{[]Collection{c1, c2}, "age-of-rust", nil},
		{[]Collection{c1, c2, c3}, "age-of-rust", &c3},
		{[]Collection{c1, c2}, "cryptokitties", &c2},
		{[]Collection{c1}, "age-of-rust", nil},
		{[]Collection{c1, c3}, "enjin", &c1},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("searchCollection %d", i), func(t *testing.T) {
			s := searchCollection(tt.collections, tt.collectibleID)
			assert.EqualValues(t, s, tt.result)
		})
	}

}

func TestNormalizeSupportedContracts(t *testing.T) {
	var contracts []PrimaryAssetContract
	err := json.Unmarshal([]byte(rawAssetContracts), &contracts)
	assert.Nil(t, err)
	var collection = Collection{}
	collection.Contracts = contracts
	normalizeSupportedContracts(&collection)
	assert.Equal(t, len(collection.Contracts), 2, "normalizeSupportedContracts with incorrect len")
	var expectedContracts []PrimaryAssetContract
	err = json.Unmarshal([]byte(rawAssetContractsExpected), &expectedContracts)
	assert.Nil(t, err)
	assert.Equal(t, collection.Contracts, expectedContracts, "normalizeSupportedContracts expectedContracts")
}

const rawAssetContracts = `[
    {
        "address": "0xee85966b4974d3c6f71a2779cc3b6f53afbc2b68",
        "asset_contract_type": "fungible",
        "created_date": "2019-10-16T07:36:16.102163",
        "name": "Rare Chest",
        "nft_version": null,
        "opensea_version": null,
        "owner": 1610615,
        "schema_name": "ERC20",
        "symbol": "",
        "total_supply": "1",
        "description": "Gods Unchained is a free-to-play, turn-based competitive trading card game in which cards can be bought and sold on the OpenSea marketplace. Players use their collection to build decks of cards, and select a God to play with at the start of each match. The goal of the game is to reduce your opponent's life to zero. Each deck contains exactly 30 cards. On OpenSea, cards can be sold for a fixed price, auctioned, or sold in bundles.",
        "external_link": "https://godsunchained.com/?refcode=0x5b3256965e7C3cF26E11FCAf296DfC8807C01073",
        "image_url": "https://lh3.googleusercontent.com/yArciVdcDv3O2R-O8XCxx3YEYZdzpiCMdossjUgv0kpLIluUQ1bYN_dyEk5xcvBEOgeq0zNIoWOh7TL9DvUEv--OLQ=s60",
        "default_to_fiat": false,
        "dev_buyer_fee_basis_points": 0,
        "dev_seller_fee_basis_points": 0,
        "only_proxied_transfers": false,
        "opensea_buyer_fee_basis_points": 0,
        "opensea_seller_fee_basis_points": 250,
        "buyer_fee_basis_points": 0,
        "seller_fee_basis_points": 250,
        "payout_address": null
    },
    {
        "address": "0x20d4cec36528e1c4563c1bfbe3de06aba70b22b4",
        "asset_contract_type": "fungible",
        "created_date": "2019-10-16T08:06:42.727997",
        "name": "Legendary Chest",
        "nft_version": null,
        "opensea_version": null,
        "owner": 1610615,
        "schema_name": "ERC20",
        "symbol": "",
        "total_supply": "1",
        "description": "Gods Unchained is a free-to-play, turn-based competitive trading card game in which cards can be bought and sold on the OpenSea marketplace. Players use their collection to build decks of cards, and select a God to play with at the start of each match. The goal of the game is to reduce your opponent's life to zero. Each deck contains exactly 30 cards. On OpenSea, cards can be sold for a fixed price, auctioned, or sold in bundles.",
        "external_link": "https://godsunchained.com/?refcode=0x5b3256965e7C3cF26E11FCAf296DfC8807C01073",
        "image_url": "https://lh3.googleusercontent.com/yArciVdcDv3O2R-O8XCxx3YEYZdzpiCMdossjUgv0kpLIluUQ1bYN_dyEk5xcvBEOgeq0zNIoWOh7TL9DvUEv--OLQ=s60",
        "default_to_fiat": false,
        "dev_buyer_fee_basis_points": 0,
        "dev_seller_fee_basis_points": 0,
        "only_proxied_transfers": false,
        "opensea_buyer_fee_basis_points": 0,
        "opensea_seller_fee_basis_points": 250,
        "buyer_fee_basis_points": 0,
        "seller_fee_basis_points": 250,
        "payout_address": null
    },
    {
        "address": "0x0e3a2a1f2146d86a604adc220b4967a898d7fe07",
        "asset_contract_type": "non-fungible",
        "created_date": "2019-11-01T06:39:04.363034",
        "name": "Gods Unchained Cards",
        "nft_version": "3.0",
        "opensea_version": null,
        "owner": 1691695,
        "schema_name": "ERC721",
        "symbol": "",
        "total_supply": "1",
        "description": "Gods Unchained is a free-to-play, turn-based competitive trading card game in which cards can be bought and sold on the OpenSea marketplace. Players use their collection to build decks of cards, and select a God to play with at the start of each match. The goal of the game is to reduce your opponent's life to zero. Each deck contains exactly 30 cards. On OpenSea, cards can be sold for a fixed price, auctioned, or sold in bundles.",
        "external_link": "https://godsunchained.com/?refcode=0x5b3256965e7C3cF26E11FCAf296DfC8807C01073",
        "image_url": "https://lh3.googleusercontent.com/yArciVdcDv3O2R-O8XCxx3YEYZdzpiCMdossjUgv0kpLIluUQ1bYN_dyEk5xcvBEOgeq0zNIoWOh7TL9DvUEv--OLQ=s60",
        "default_to_fiat": false,
        "dev_buyer_fee_basis_points": 0,
        "dev_seller_fee_basis_points": 0,
        "only_proxied_transfers": false,
        "opensea_buyer_fee_basis_points": 0,
        "opensea_seller_fee_basis_points": 250,
        "buyer_fee_basis_points": 0,
        "seller_fee_basis_points": 250,
        "payout_address": null
    },
    {
        "address": "0x564cb55c655f727b61d9baf258b547ca04e9e548",
        "asset_contract_type": "non-fungible",
        "created_date": "2019-10-29T12:28:37.643714",
        "name": "Gods Unchained",
        "nft_version": "3.0",
        "opensea_version": null,
        "owner": 1691695,
        "schema_name": "ERC721",
        "symbol": "",
        "total_supply": "205",
        "description": "Gods Unchained is a free-to-play, turn-based competitive trading card game in which cards can be bought and sold on the OpenSea marketplace. Players use their collection to build decks of cards, and select a God to play with at the start of each match. The goal of the game is to reduce your opponent's life to zero. Each deck contains exactly 30 cards. On OpenSea, cards can be sold for a fixed price, auctioned, or sold in bundles.",
        "external_link": "https://godsunchained.com/?refcode=0x5b3256965e7C3cF26E11FCAf296DfC8807C01073",
        "image_url": "https://lh3.googleusercontent.com/yArciVdcDv3O2R-O8XCxx3YEYZdzpiCMdossjUgv0kpLIluUQ1bYN_dyEk5xcvBEOgeq0zNIoWOh7TL9DvUEv--OLQ=s60",
        "default_to_fiat": false,
        "dev_buyer_fee_basis_points": 0,
        "dev_seller_fee_basis_points": 0,
        "only_proxied_transfers": false,
        "opensea_buyer_fee_basis_points": 0,
        "opensea_seller_fee_basis_points": 250,
        "buyer_fee_basis_points": 0,
        "seller_fee_basis_points": 250,
        "payout_address": null
    }
]`

const rawAssetContractsExpected = `[{
        "address": "0x0e3a2a1f2146d86a604adc220b4967a898d7fe07",
        "asset_contract_type": "non-fungible",
        "created_date": "2019-11-01T06:39:04.363034",
        "name": "Gods Unchained Cards",
        "nft_version": "3.0",
        "opensea_version": null,
        "owner": 1691695,
        "schema_name": "ERC721",
        "symbol": "",
        "total_supply": "1",
        "description": "Gods Unchained is a free-to-play, turn-based competitive trading card game in which cards can be bought and sold on the OpenSea marketplace. Players use their collection to build decks of cards, and select a God to play with at the start of each match. The goal of the game is to reduce your opponent's life to zero. Each deck contains exactly 30 cards. On OpenSea, cards can be sold for a fixed price, auctioned, or sold in bundles.",
        "external_link": "https://godsunchained.com/?refcode=0x5b3256965e7C3cF26E11FCAf296DfC8807C01073",
        "image_url": "https://lh3.googleusercontent.com/yArciVdcDv3O2R-O8XCxx3YEYZdzpiCMdossjUgv0kpLIluUQ1bYN_dyEk5xcvBEOgeq0zNIoWOh7TL9DvUEv--OLQ=s60",
        "default_to_fiat": false,
        "dev_buyer_fee_basis_points": 0,
        "dev_seller_fee_basis_points": 0,
        "only_proxied_transfers": false,
        "opensea_buyer_fee_basis_points": 0,
        "opensea_seller_fee_basis_points": 250,
        "buyer_fee_basis_points": 0,
        "seller_fee_basis_points": 250,
        "payout_address": null
    },
    {
        "address": "0x564cb55c655f727b61d9baf258b547ca04e9e548",
        "asset_contract_type": "non-fungible",
        "created_date": "2019-10-29T12:28:37.643714",
        "name": "Gods Unchained",
        "nft_version": "3.0",
        "opensea_version": null,
        "owner": 1691695,
        "schema_name": "ERC721",
        "symbol": "",
        "total_supply": "205",
        "description": "Gods Unchained is a free-to-play, turn-based competitive trading card game in which cards can be bought and sold on the OpenSea marketplace. Players use their collection to build decks of cards, and select a God to play with at the start of each match. The goal of the game is to reduce your opponent's life to zero. Each deck contains exactly 30 cards. On OpenSea, cards can be sold for a fixed price, auctioned, or sold in bundles.",
        "external_link": "https://godsunchained.com/?refcode=0x5b3256965e7C3cF26E11FCAf296DfC8807C01073",
        "image_url": "https://lh3.googleusercontent.com/yArciVdcDv3O2R-O8XCxx3YEYZdzpiCMdossjUgv0kpLIluUQ1bYN_dyEk5xcvBEOgeq0zNIoWOh7TL9DvUEv--OLQ=s60",
        "default_to_fiat": false,
        "dev_buyer_fee_basis_points": 0,
        "dev_seller_fee_basis_points": 0,
        "only_proxied_transfers": false,
        "opensea_buyer_fee_basis_points": 0,
        "opensea_seller_fee_basis_points": 250,
        "buyer_fee_basis_points": 0,
        "seller_fee_basis_points": 250,
        "payout_address": null
    }
]`
