package opensea

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/mock"
	"github.com/trustwallet/golibs/types"
)

var (
	collectionsOwnerV4  = "0x0875BCab22dE3d02402bc38aEe4104e1239374a7"
	collectionsSrcV4, _ = mock.JsonStringFromFilePath("mocks/opensea_collections.json")
	collectibleSrcV4, _ = mock.JsonStringFromFilePath("mocks/opensea_collectible.json")

	collection1DstV4 = types.Collection{
		Name:         "CryptoKitties",
		ImageUrl:     "https://storage.opensea.io/0x06012c8cf97bead5deae237070f9587f8e7a266d-featured-1556588705.png",
		Description:  "CryptoKitties is a game centered around breedable, collectible, and oh-so-adorable creatures we call CryptoKitties! Each cat is one-of-a-kind and 100% owned by you; it cannot be replicated, taken away, or destroyed.",
		ExternalLink: "https://www.cryptokitties.co/",
		Total:        3,
		Id:           "cryptokitties",
		Address:      "0x0875BCab22dE3d02402bc38aEe4104e1239374a7",
		Coin:         60,
	}

	collection2DstV4 = types.Collection{
		Name:         "Age of Rust",
		ImageUrl:     "https://storage.opensea.io/age-of-rust-1561960816.jpg",
		Description:  "Year 4424: The search begins for new life on the other side of the galaxy. Explore abandoned space stations, mysterious caverns, and ruins on far away worlds in order to unlock puzzles and secrets! Beware the rogue machines!",
		ExternalLink: "https://www.ageofrust.games/",
		Total:        1,
		Id:           "age-of-rust",
		Address:      "0x0875BCab22dE3d02402bc38aEe4104e1239374a7",
		Coin:         60,
	}

	collection3DstV4 = types.Collection{
		Name:         "Gods Unchained",
		ImageUrl:     "https://lh3.googleusercontent.com/yArciVdcDv3O2R-O8XCxx3YEYZdzpiCMdossjUgv0kpLIluUQ1bYN_dyEk5xcvBEOgeq0zNIoWOh7TL9DvUEv--OLQ=s60",
		Description:  "Gods Unchained is a free-to-play, turn-based competitive trading card game in which cards can be bought and sold on the OpenSea marketplace. Players use their collection to build decks of cards, and select a God to play with at the start of each match. The goal of the game is to reduce your opponent's life to zero. Each deck contains exactly 30 cards. On OpenSea, cards can be sold for a fixed price, auctioned, or sold in bundles.",
		ExternalLink: "https://godsunchained.com/?refcode=0x5b3256965e7C3cF26E11FCAf296DfC8807C01073",
		Total:        535,
		Id:           "gods-unchained",
		Address:      "0x0875BCab22dE3d02402bc38aEe4104e1239374a7",
		Coin:         60,
	}

	collectibleDstV4 = types.Collectible{
		ID:              "0xfaafdc07907ff5120a76b34b731b278c38d6043c-54277541829991970107421667568590323026590803461896874578610080514640537714688",
		CollectionID:    "age-of-rust",
		TokenID:         "54277541829991970107421667568590323026590803461896874578610080514640537714688",
		ContractAddress: "0xfaafdc07907ff5120a76b34b731b278c38d6043c",
		Category:        "Age of Rust",
		ImageUrl:        "https://storage.opensea.io/0xfaafdc07907ff5120a76b34b731b278c38d6043c-preview/54277541829991970107421667568590323026590803461896874578610080514640537714688-1564858807.png",
		ExternalLink:    "https://opensea.io/",
		ProviderLink:    "https://opensea.io/assets/0xfaafdc07907ff5120a76b34b731b278c38d6043c/54277541829991970107421667568590323026590803461896874578610080514640537714688",
		Type:            "ERC1155",
		Description:     "Rustbits are the main token of use within the Age of Rust game universe. You need Rustbits to not only play Age of Rust, but also to purchase in-game cryptoitems as well. Rustbits are radioactive rust scraped off of hulls of abandoned ships that are in orbit around a hidden planet, which is also a gas giant. The planet is so radioactive, it damages ships and kills anyone that gets close to it. Getting bits of rust off of ships is highly rare and prized.",
		Coin:            60,
		Name:            "Rustbits",
	}
)

func TestNormalizeCollectionV4(t *testing.T) {
	var collections []Collection
	err := json.Unmarshal([]byte(collectionsSrcV4), &collections)
	assert.Nil(t, err)
	page := NormalizeCollections(collections, coin.ETH, collectionsOwnerV4)
	assert.Equal(t, 3, len(page), "collections could not be normalized")
	expected := types.CollectionPage{collection1DstV4, collection2DstV4, collection3DstV4}
	assert.Equal(t, page, expected, "collections don't equal")
}

func TestNormalizeCollectibleV4(t *testing.T) {
	var collectibles []Collectible
	err := json.Unmarshal([]byte(collectibleSrcV4), &collectibles)
	assert.Nil(t, err)
	page := NormalizeCollectiblePage(collectibles, coin.ETH)
	assert.Equal(t, len(page), 1, "collectible could not be normalized")
	expected := types.CollectiblePage{collectibleDstV4}
	assert.Equal(t, page, expected, "collectible don't equal")
}
