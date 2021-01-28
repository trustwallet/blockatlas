package bounce

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/mock"
	"github.com/trustwallet/golibs/types"
)

func Test_normalizeCollectible(t *testing.T) {
	type args struct {
		filename  string
		coinIndex uint
		info      CollectionInfo
	}
	tests := []struct {
		name string
		args args
		want types.Collectible
	}{
		{
			name: "Test empty contract name",
			args: args{
				filename:  "artwork_nft.json",
				coinIndex: coin.SMARTCHAIN,
				info: CollectionInfo{
					Name:        "Hungry",
					Description: "Animal Series",
					Image:       "https://d3ggs2vjn5heyw.cloudfront.net/static/nfts/artworks/d9dc679ec0614eb78b479aed21694305.jpg",
				},
			},
			want: types.Collectible{
				ID:              "0x5Bc94e9347F3b9Be8415bDfd24af16666704E44f-450",
				CollectionID:    "0x5Bc94e9347F3b9Be8415bDfd24af16666704E44f",
				TokenID:         "450",
				ContractAddress: "0x5Bc94e9347F3b9Be8415bDfd24af16666704E44f",
				Category:        "Hungry",
				ImageUrl:        "https://d3ggs2vjn5heyw.cloudfront.net/static/nfts/artworks/d9dc679ec0614eb78b479aed21694305.jpg",
				ExternalLink:    "https://www.bakeryswap.org/api/v1/artworks/fb4576253e3d45cebf0ac4c8df8f1743",
				Type:            string(types.ERC721),
				Description:     "Animal Series",
				Coin:            coin.SMARTCHAIN,
				Name:            "Hungry",
				Version:         "3.0",
			},
		},
		{
			name: "Test pancake bunny",
			args: args{
				filename:  "pancake_bunny.json",
				coinIndex: coin.SMARTCHAIN,
				info: CollectionInfo{
					Name:        "Swapsies",
					Description: "These bunnies love nothing more than swapping pancakes. Especially on BSC.",
					Image:       "https://ipfs.io/ipfs/QmXdHqg3nywpNJWDevJQPtkz93vpfoHcZWQovFz2nmtPf5/swapsies.png",
				},
			},
			want: types.Collectible{
				ID:              "0xDf7952B35f24aCF7fC0487D01c8d5690a60DBa07-409",
				CollectionID:    "0xDf7952B35f24aCF7fC0487D01c8d5690a60DBa07",
				TokenID:         "409",
				ContractAddress: "0xDf7952B35f24aCF7fC0487D01c8d5690a60DBa07",
				Category:        "Pancake Bunnies",
				ImageUrl:        "https://ipfs.io/ipfs/QmXdHqg3nywpNJWDevJQPtkz93vpfoHcZWQovFz2nmtPf5/swapsies.png",
				ExternalLink:    "https://ipfs.io/ipfs/QmYu9WwPNKNSZQiTCDfRk7aCR472GURavR9M1qosDmqpev/swapsies.json",
				Type:            string(types.ERC721),
				Description:     "These bunnies love nothing more than swapping pancakes. Especially on BSC.",
				Coin:            coin.SMARTCHAIN,
				Name:            "Swapsies",
				Version:         "3.0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var c Collectible
			err := mock.JsonModelFromFilePath("mocks/"+tt.args.filename, &c)
			assert.Nil(t, err)
			if got := normalizeCollectible(c, tt.args.coinIndex, tt.args.info); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("normalizeCollectible() = %v, want %v", got, tt.want)
			}
		})
	}
}
