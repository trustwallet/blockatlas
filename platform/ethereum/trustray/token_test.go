package trustray

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/golibs/types"

	"github.com/trustwallet/golibs/coin"
)

const tokenSrc = `
{
	"address": "0xa14839c9837657EFcDE754EbEAF5cbECDd801B2A",
	"name": "FusChain",
	"decimals": 18,
	"symbol": "FUS"
}`

type testToken struct {
	apiResponse string
	expected    *types.Token
	coin        int
}

func TestNormalizeToken(t *testing.T) {
	var tests = []struct {
		name     string
		tokenRaw string
		coin     int
		want     types.Token
	}{
		{
			"ethereum erc20",
			tokenSrc,
			coin.ETHEREUM,
			types.Token{
				Name:     "FusChain",
				Symbol:   "FUS",
				Decimals: 18,
				TokenID:  "0xa14839c9837657EFcDE754EbEAF5cbECDd801B2A",
				Coin:     coin.ETHEREUM,
				Type:     types.ERC20,
			},
		},
		{"classic etc20",
			tokenSrc,
			coin.CLASSIC,
			types.Token{
				Name:     "FusChain",
				Symbol:   "FUS",
				Decimals: 18,
				TokenID:  "0xa14839c9837657EFcDE754EbEAF5cbECDd801B2A",
				Coin:     coin.CLASSIC,
				Type:     types.ETC20,
			},
		},
		{"gochain go20",
			tokenSrc,
			coin.GOCHAIN,
			types.Token{
				Name:     "FusChain",
				Symbol:   "FUS",
				Decimals: 18,
				TokenID:  "0xa14839c9837657EFcDE754EbEAF5cbECDd801B2A",
				Coin:     coin.GOCHAIN,
				Type:     types.GO20,
			},
		},
		{"thudertoken tt20",
			tokenSrc,
			coin.THUNDERTOKEN,
			types.Token{
				Name:     "FusChain",
				Symbol:   "FUS",
				Decimals: 18,
				TokenID:  "0xa14839c9837657EFcDE754EbEAF5cbECDd801B2A",
				Coin:     coin.THUNDERTOKEN,
				Type:     types.TT20,
			},
		},
		{"wanchain wan20",
			tokenSrc,
			coin.WANCHAIN,
			types.Token{
				Name:     "FusChain",
				Symbol:   "FUS",
				Decimals: 18,
				TokenID:  "0xa14839c9837657EFcDE754EbEAF5cbECDd801B2A",
				Coin:     coin.WANCHAIN,
				Type:     types.WAN20,
			},
		},
		{"poa poa20",
			tokenSrc,
			coin.POA,
			types.Token{
				Name:     "FusChain",
				Symbol:   "FUS",
				Decimals: 18,
				TokenID:  "0xa14839c9837657EFcDE754EbEAF5cbECDd801B2A",
				Coin:     coin.POA,
				Type:     types.POA20,
			},
		},
		{"callisto clo20",
			tokenSrc,
			coin.CALLISTO,
			types.Token{
				Name:     "FusChain",
				Symbol:   "FUS",
				Decimals: 18,
				TokenID:  "0xa14839c9837657EFcDE754EbEAF5cbECDd801B2A",
				Coin:     coin.CALLISTO,
				Type:     types.CLO20,
			},
		},
		{"unknown",
			tokenSrc,
			1999,
			types.Token{
				Name:     "FusChain",
				Symbol:   "FUS",
				Decimals: 18,
				TokenID:  "0xa14839c9837657EFcDE754EbEAF5cbECDd801B2A",
				Coin:     1999,
				Type:     types.ERC20,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testNormalizeToken(t, &testToken{
				apiResponse: tt.tokenRaw,
				expected:    &tt.want,
				coin:        tt.coin,
			})
		})
	}

}

func testNormalizeToken(t *testing.T, _test *testToken) {
	var token Contract
	err := json.Unmarshal([]byte(_test.apiResponse), &token)
	if err != nil {
		t.Error(err)
		return
	}
	tk := NormalizeToken(&token, uint(_test.coin))

	resJSON, err := json.Marshal(&tk)
	if err != nil {
		t.Fatal(err)
	}

	dstJSON, err := json.Marshal(_test.expected)
	if err != nil {
		t.Fatal(err)
	}

	assert.JSONEq(t, string(resJSON), string(dstJSON))
}
