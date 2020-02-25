package ethereum

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

const collectionsOwnerV4 = "0x0875BCab22dE3d02402bc38aEe4104e1239374a7"

const collectionsSrcV4 = `
[
   {
      "primary_asset_contracts":[
         {
            "address":"0x06012c8cf97bead5deae237070f9587f8e7a266d",
            "name":"CryptoKitties",
            "symbol":"CKITTY",
            "description":"CryptoKitties is a game centered around breedable, collectible, and oh-so-adorable creatures we call CryptoKitties! Each cat is one-of-a-kind and 100% owned by you; it cannot be replicated, taken away, or destroyed.",
            "external_link":"https://www.cryptokitties.co/",
            "nft_version":"1.0",
            "schema_name":"ERC721",
            "display_data":{
               "images":[
                  "https://storage.googleapis.com/ck-kitty-image/0x06012c8cf97bead5deae237070f9587f8e7a266d/564155.svg",
                  "https://storage.googleapis.com/ck-kitty-image/0x06012c8cf97bead5deae237070f9587f8e7a266d/546630.svg",
                  "https://storage.googleapis.com/ck-kitty-image/0x06012c8cf97bead5deae237070f9587f8e7a266d/441529.svg",
                  "https://storage.googleapis.com/ck-kitty-image/0x06012c8cf97bead5deae237070f9587f8e7a266d/552435.svg",
                  "https://storage.googleapis.com/ck-kitty-image/0x06012c8cf97bead5deae237070f9587f8e7a266d/524748.png",
                  "https://storage.googleapis.com/ck-kitty-image/0x06012c8cf97bead5deae237070f9587f8e7a266d/540800.svg"
               ],
               "card_display_style":"padded"
            },
            "image_url":"https://storage.opensea.io/0x06012c8cf97bead5deae237070f9587f8e7a266d-featured-1556588705.png"
         }
      ],
      "name":"CryptoKitties",
      "slug":"cryptokitties",
      "image_url":"https://storage.opensea.io/0x06012c8cf97bead5deae237070f9587f8e7a266d-featured-1556588705.png",
      "description":"CryptoKitties is a game centered around breedable, collectible, and oh-so-adorable creatures we call CryptoKitties! Each cat is one-of-a-kind and 100% owned by you; it cannot be replicated, taken away, or destroyed.",
      "external_url":"https://www.cryptokitties.co/",
      "featured_image_url":"https://storage.opensea.io/0x06012c8cf97bead5deae237070f9587f8e7a266d-featured-1556589429.png",
      "created_date":"2019-04-26T22:13:04.207050",
      "owned_asset_count":3
   },
   {
      "primary_asset_contracts":[
         {
            "address":"0xfaafdc07907ff5120a76b34b731b278c38d6043c",
            "name":"Enjin",
            "symbol":"",
            "description":"",
            "external_link":null,
            "nft_version":null,
            "schema_name":"ERC1155",
            "display_data":{

            },
            "owner":null,
            "created_date":"2019-08-02T23:43:14.666153",
            "asset_contract_type":"semi-fungible"
         }
      ],
      "name":"Age of Rust",
      "slug":"age-of-rust",
      "image_url":"https://storage.opensea.io/age-of-rust-1561960816.jpg",
      "description":"Year 4424: The search begins for new life on the other side of the galaxy. Explore abandoned space stations, mysterious caverns, and ruins on far away worlds in order to unlock puzzles and secrets! Beware the rogue machines!",
      "external_url":"https://www.ageofrust.games/",
      "featured_image_url":null,
      "created_date":"2019-09-03T02:35:56.063685",
      "owned_asset_count":1
   },
   {
      "primary_asset_contracts":[
         {
            "address":"0xee85966b4974d3c6f71a2779cc3b6f53afbc2b68",
            "asset_contract_type":"fungible",
            "created_date":"2019-10-16T07:36:16.102163",
            "name":"Rare Chest",
            "nft_version":null,
            "opensea_version":null,
            "owner":1610615,
            "schema_name":"ERC20",
            "symbol":"",
            "total_supply":"1",
            "description":"Gods Unchained is a free-to-play, turn-based competitive trading card game in which cards can be bought and sold on the OpenSea marketplace. Players use their collection to build decks of cards, and select a God to play with at the start of each match. The goal of the game is to reduce your opponent's life to zero. Each deck contains exactly 30 cards. On OpenSea, cards can be sold for a fixed price, auctioned, or sold in bundles.",
            "external_link":"https://godsunchained.com/?refcode=0x5b3256965e7C3cF26E11FCAf296DfC8807C01073",
            "image_url":"https://lh3.googleusercontent.com/yArciVdcDv3O2R-O8XCxx3YEYZdzpiCMdossjUgv0kpLIluUQ1bYN_dyEk5xcvBEOgeq0zNIoWOh7TL9DvUEv--OLQ=s60",
            "default_to_fiat":false,
            "dev_buyer_fee_basis_points":0,
            "dev_seller_fee_basis_points":0,
            "only_proxied_transfers":false,
            "opensea_buyer_fee_basis_points":0,
            "opensea_seller_fee_basis_points":250,
            "buyer_fee_basis_points":0,
            "seller_fee_basis_points":250,
            "payout_address":null
         },
         {
            "address":"0x20d4cec36528e1c4563c1bfbe3de06aba70b22b4",
            "asset_contract_type":"fungible",
            "created_date":"2019-10-16T08:06:42.727997",
            "name":"Legendary Chest",
            "nft_version":null,
            "opensea_version":null,
            "owner":1610615,
            "schema_name":"ERC20",
            "symbol":"",
            "total_supply":"1",
            "description":"Gods Unchained is a free-to-play, turn-based competitive trading card game in which cards can be bought and sold on the OpenSea marketplace. Players use their collection to build decks of cards, and select a God to play with at the start of each match. The goal of the game is to reduce your opponent's life to zero. Each deck contains exactly 30 cards. On OpenSea, cards can be sold for a fixed price, auctioned, or sold in bundles.",
            "external_link":"https://godsunchained.com/?refcode=0x5b3256965e7C3cF26E11FCAf296DfC8807C01073",
            "image_url":"https://lh3.googleusercontent.com/yArciVdcDv3O2R-O8XCxx3YEYZdzpiCMdossjUgv0kpLIluUQ1bYN_dyEk5xcvBEOgeq0zNIoWOh7TL9DvUEv--OLQ=s60",
            "default_to_fiat":false,
            "dev_buyer_fee_basis_points":0,
            "dev_seller_fee_basis_points":0,
            "only_proxied_transfers":false,
            "opensea_buyer_fee_basis_points":0,
            "opensea_seller_fee_basis_points":250,
            "buyer_fee_basis_points":0,
            "seller_fee_basis_points":250,
            "payout_address":null
         },
         {
            "address":"0x0e3a2a1f2146d86a604adc220b4967a898d7fe07",
            "asset_contract_type":"non-fungible",
            "created_date":"2019-11-01T06:39:04.363034",
            "name":"Gods Unchained Cards",
            "nft_version":"3.0",
            "opensea_version":null,
            "owner":1691695,
            "schema_name":"ERC721",
            "symbol":"",
            "total_supply":"1",
            "description":"Gods Unchained is a free-to-play, turn-based competitive trading card game in which cards can be bought and sold on the OpenSea marketplace. Players use their collection to build decks of cards, and select a God to play with at the start of each match. The goal of the game is to reduce your opponent's life to zero. Each deck contains exactly 30 cards. On OpenSea, cards can be sold for a fixed price, auctioned, or sold in bundles.",
            "external_link":"https://godsunchained.com/?refcode=0x5b3256965e7C3cF26E11FCAf296DfC8807C01073",
            "image_url":"https://lh3.googleusercontent.com/yArciVdcDv3O2R-O8XCxx3YEYZdzpiCMdossjUgv0kpLIluUQ1bYN_dyEk5xcvBEOgeq0zNIoWOh7TL9DvUEv--OLQ=s60",
            "default_to_fiat":false,
            "dev_buyer_fee_basis_points":0,
            "dev_seller_fee_basis_points":0,
            "only_proxied_transfers":false,
            "opensea_buyer_fee_basis_points":0,
            "opensea_seller_fee_basis_points":250,
            "buyer_fee_basis_points":0,
            "seller_fee_basis_points":250,
            "payout_address":null
         },
         {
            "address":"0x564cb55c655f727b61d9baf258b547ca04e9e548",
            "asset_contract_type":"non-fungible",
            "created_date":"2019-10-29T12:28:37.643714",
            "name":"Gods Unchained",
            "nft_version":"3.0",
            "opensea_version":null,
            "owner":1691695,
            "schema_name":"ERC721",
            "symbol":"",
            "total_supply":"205",
            "description":"Gods Unchained is a free-to-play, turn-based competitive trading card game in which cards can be bought and sold on the OpenSea marketplace. Players use their collection to build decks of cards, and select a God to play with at the start of each match. The goal of the game is to reduce your opponent's life to zero. Each deck contains exactly 30 cards. On OpenSea, cards can be sold for a fixed price, auctioned, or sold in bundles.",
            "external_link":"https://godsunchained.com/?refcode=0x5b3256965e7C3cF26E11FCAf296DfC8807C01073",
            "image_url":"https://lh3.googleusercontent.com/yArciVdcDv3O2R-O8XCxx3YEYZdzpiCMdossjUgv0kpLIluUQ1bYN_dyEk5xcvBEOgeq0zNIoWOh7TL9DvUEv--OLQ=s60",
            "default_to_fiat":false,
            "dev_buyer_fee_basis_points":0,
            "dev_seller_fee_basis_points":0,
            "only_proxied_transfers":false,
            "opensea_buyer_fee_basis_points":0,
            "opensea_seller_fee_basis_points":250,
            "buyer_fee_basis_points":0,
            "seller_fee_basis_points":250,
            "payout_address":null
         }
      ],
      "traits":{
         "mana":{
            "min":1,
            "max":18
         },
         "health":{
            "min":1,
            "max":16
         },
         "attack":{
            "min":1,
            "max":16
         }
      },
      "stats":{
         "seven_day_volume":270.564697067863,
         "seven_day_change":-0.014179906699531201,
         "total_volume":11825.0384171324,
         "count":6835331,
         "num_owners":10791,
         "market_cap":150892.14138332763,
         "average_price":0.0396514694030496,
         "items_sold":298233
      },
      "banner_image_url":null,
      "chat_url":null,
      "created_date":"2019-11-13T03:01:42.051246",
      "default_to_fiat":false,
      "description":"Gods Unchained is a free-to-play, turn-based competitive trading card game in which cards can be bought and sold on the OpenSea marketplace. Players use their collection to build decks of cards, and select a God to play with at the start of each match. The goal of the game is to reduce your opponent's life to zero. Each deck contains exactly 30 cards. On OpenSea, cards can be sold for a fixed price, auctioned, or sold in bundles.",
      "dev_buyer_fee_basis_points":"0",
      "dev_seller_fee_basis_points":"0",
      "display_data":{
         "images":[
            "https://storage.googleapis.com/opensea-prod.appspot.com/0x6ebeaf8e8e946f0716e6533a6f2cefc83f60e8ab/25233.png",
            "https://storage.googleapis.com/opensea-prod.appspot.com/0x6ebeaf8e8e946f0716e6533a6f2cefc83f60e8ab/152875.png",
            "https://storage.googleapis.com/opensea-prod.appspot.com/0x6ebeaf8e8e946f0716e6533a6f2cefc83f60e8ab/25669.png",
            "https://storage.googleapis.com/opensea-prod.appspot.com/0x6ebeaf8e8e946f0716e6533a6f2cefc83f60e8ab/9237.png",
            "https://storage.googleapis.com/opensea-prod.appspot.com/0x6ebeaf8e8e946f0716e6533a6f2cefc83f60e8ab/9228.png",
            "https://storage.googleapis.com/opensea-prod.appspot.com/0x6ebeaf8e8e946f0716e6533a6f2cefc83f60e8ab/9231.png"
         ],
         "card_display_style":"contain"
      },
      "external_url":"https://godsunchained.com/?refcode=0x5b3256965e7C3cF26E11FCAf296DfC8807C01073",
      "featured":true,
      "featured_image_url":"https://storage.opensea.io/0x6ebeaf8e8e946f0716e6533a6f2cefc83f60e8ab-featured-1556589419.png",
      "hidden":false,
      "image_url":"https://lh3.googleusercontent.com/yArciVdcDv3O2R-O8XCxx3YEYZdzpiCMdossjUgv0kpLIluUQ1bYN_dyEk5xcvBEOgeq0zNIoWOh7TL9DvUEv--OLQ=s60",
      "is_subject_to_whitelist":false,
      "large_image_url":"https://lh3.googleusercontent.com/yArciVdcDv3O2R-O8XCxx3YEYZdzpiCMdossjUgv0kpLIluUQ1bYN_dyEk5xcvBEOgeq0zNIoWOh7TL9DvUEv--OLQ",
      "name":"Gods Unchained",
      "only_proxied_transfers":false,
      "opensea_buyer_fee_basis_points":"0",
      "opensea_seller_fee_basis_points":"250",
      "payout_address":null,
      "require_email":false,
      "short_description":null,
      "slug":"gods-unchained",
      "wiki_url":null,
      "owned_asset_count":535
   }
]
`

var collection1DstV4 = blockatlas.Collection{
	Name:         "CryptoKitties",
	Symbol:       "",
	Slug:         "cryptokitties",
	ImageUrl:     "https://storage.opensea.io/0x06012c8cf97bead5deae237070f9587f8e7a266d-featured-1556588705.png",
	Description:  "CryptoKitties is a game centered around breedable, collectible, and oh-so-adorable creatures we call CryptoKitties! Each cat is one-of-a-kind and 100% owned by you; it cannot be replicated, taken away, or destroyed.",
	ExternalLink: "https://www.cryptokitties.co/",
	Total:        3,
	Id:           "cryptokitties",
	Address:      "0x0875BCab22dE3d02402bc38aEe4104e1239374a7",
	Coin:         60,
}

var collection2DstV4 = blockatlas.Collection{
	Name:         "Age of Rust",
	Symbol:       "",
	Slug:         "age-of-rust",
	ImageUrl:     "https://storage.opensea.io/age-of-rust-1561960816.jpg",
	Description:  "Year 4424: The search begins for new life on the other side of the galaxy. Explore abandoned space stations, mysterious caverns, and ruins on far away worlds in order to unlock puzzles and secrets! Beware the rogue machines!",
	ExternalLink: "https://www.ageofrust.games/",
	Total:        1,
	Id:           "age-of-rust",
	Address:      "0x0875BCab22dE3d02402bc38aEe4104e1239374a7",
	Coin:         60,
}

var collection3DstV4 = blockatlas.Collection{
	Name:         "Gods Unchained",
	Symbol:       "",
	Slug:         "gods-unchained",
	ImageUrl:     "https://lh3.googleusercontent.com/yArciVdcDv3O2R-O8XCxx3YEYZdzpiCMdossjUgv0kpLIluUQ1bYN_dyEk5xcvBEOgeq0zNIoWOh7TL9DvUEv--OLQ=s60",
	Description:  "Gods Unchained is a free-to-play, turn-based competitive trading card game in which cards can be bought and sold on the OpenSea marketplace. Players use their collection to build decks of cards, and select a God to play with at the start of each match. The goal of the game is to reduce your opponent's life to zero. Each deck contains exactly 30 cards. On OpenSea, cards can be sold for a fixed price, auctioned, or sold in bundles.",
	ExternalLink: "https://godsunchained.com/?refcode=0x5b3256965e7C3cF26E11FCAf296DfC8807C01073",
	Total:        535,
	Id:           "gods-unchained",
	Address:      "0x0875BCab22dE3d02402bc38aEe4104e1239374a7",
	Coin:         60,
}

func TestNormalizeCollectionV4(t *testing.T) {
	var collections []Collection
	err := json.Unmarshal([]byte(collectionsSrcV4), &collections)
	assert.Nil(t, err)
	page := NormalizeCollectionsV4(collections, coin.ETH, collectionsOwnerV4)
	assert.Equal(t, 3, len(page), "collections could not be normalized")
	expected := blockatlas.CollectionPage{collection1DstV4, collection2DstV4, collection3DstV4}
	assert.Equal(t, page, expected, "collections don't equal")
}

const collectibleSrcV4 = `
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
	"collection": {
		"slug": "age-of-rust",
        "name": "Age of Rust",
        "external_url": "https://opensea.io/"
    },
    "permalink": "https://opensea.io/assets/0xfaafdc07907ff5120a76b34b731b278c38d6043c/54277541829991970107421667568590323026590803461896874578610080514640537714688"
  }
]
`

var collectibleDstV4 = blockatlas.Collectible{
	ID:               "0xfaafdc07907ff5120a76b34b731b278c38d6043c-54277541829991970107421667568590323026590803461896874578610080514640537714688",
	CollectionID:     "age-of-rust",
	TokenID:          "54277541829991970107421667568590323026590803461896874578610080514640537714688",
	CategoryContract: "",
	ContractAddress:  "0xfaafdc07907ff5120a76b34b731b278c38d6043c",
	Category:         "Age of Rust",
	ImageUrl:         "https://storage.opensea.io/0xfaafdc07907ff5120a76b34b731b278c38d6043c-preview/54277541829991970107421667568590323026590803461896874578610080514640537714688-1564858807.png",
	ExternalLink:     "https://opensea.io/",
	ProviderLink:     "https://opensea.io/assets/0xfaafdc07907ff5120a76b34b731b278c38d6043c/54277541829991970107421667568590323026590803461896874578610080514640537714688",
	Type:             "ERC1155",
	Description:      "Rustbits are the main token of use within the Age of Rust game universe. You need Rustbits to not only play Age of Rust, but also to purchase in-game cryptoitems as well. Rustbits are radioactive rust scraped off of hulls of abandoned ships that are in orbit around a hidden planet, which is also a gas giant. The planet is so radioactive, it damages ships and kills anyone that gets close to it. Getting bits of rust off of ships is highly rare and prized.",
	Coin:             60,
	Name:             "Rustbits",
}

func TestNormalizeCollectibleV4(t *testing.T) {
	var collectibles []Collectible
	err := json.Unmarshal([]byte(collectibleSrcV4), &collectibles)
	assert.Nil(t, err)
	page := NormalizeCollectiblePageV4(collectibles, coin.ETH)
	assert.Equal(t, len(page), 1, "collectible could not be normalized")
	expected := blockatlas.CollectiblePage{collectibleDstV4}
	assert.Equal(t, page, expected, "collectible don't equal")
}
