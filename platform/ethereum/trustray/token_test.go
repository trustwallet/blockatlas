package trustray

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/golibs/tokentype"
	"github.com/trustwallet/golibs/txtype"

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
	expected    *txtype.Token
	coin        int
}

func TestNormalizeToken(t *testing.T) {
	var tests = []struct {
		name     string
		tokenRaw string
		coin     int
		want     txtype.Token
	}{
		{
			"ethereum erc20",
			tokenSrc,
			coin.ETH,
			txtype.Token{
				Name:     "FusChain",
				Symbol:   "FUS",
				Decimals: 18,
				TokenID:  "0xa14839c9837657EFcDE754EbEAF5cbECDd801B2A",
				Coin:     coin.ETH,
				Type:     tokentype.ERC20,
			},
		},
		{"classic etc20",
			tokenSrc,
			coin.ETC,
			txtype.Token{
				Name:     "FusChain",
				Symbol:   "FUS",
				Decimals: 18,
				TokenID:  "0xa14839c9837657EFcDE754EbEAF5cbECDd801B2A",
				Coin:     coin.ETC,
				Type:     tokentype.ETC20,
			},
		},
		{"gochain go20",
			tokenSrc,
			coin.GO,
			txtype.Token{
				Name:     "FusChain",
				Symbol:   "FUS",
				Decimals: 18,
				TokenID:  "0xa14839c9837657EFcDE754EbEAF5cbECDd801B2A",
				Coin:     coin.GO,
				Type:     tokentype.GO20,
			},
		},
		{"thudertoken tt20",
			tokenSrc,
			coin.TT,
			txtype.Token{
				Name:     "FusChain",
				Symbol:   "FUS",
				Decimals: 18,
				TokenID:  "0xa14839c9837657EFcDE754EbEAF5cbECDd801B2A",
				Coin:     coin.TT,
				Type:     tokentype.TT20,
			},
		},
		{"wanchain wan20",
			tokenSrc,
			coin.WAN,
			txtype.Token{
				Name:     "FusChain",
				Symbol:   "FUS",
				Decimals: 18,
				TokenID:  "0xa14839c9837657EFcDE754EbEAF5cbECDd801B2A",
				Coin:     coin.WAN,
				Type:     tokentype.WAN20,
			},
		},
		{"poa poa20",
			tokenSrc,
			coin.POA,
			txtype.Token{
				Name:     "FusChain",
				Symbol:   "FUS",
				Decimals: 18,
				TokenID:  "0xa14839c9837657EFcDE754EbEAF5cbECDd801B2A",
				Coin:     coin.POA,
				Type:     tokentype.POA20,
			},
		},
		{"callisto clo20",
			tokenSrc,
			coin.CLO,
			txtype.Token{
				Name:     "FusChain",
				Symbol:   "FUS",
				Decimals: 18,
				TokenID:  "0xa14839c9837657EFcDE754EbEAF5cbECDd801B2A",
				Coin:     coin.CLO,
				Type:     tokentype.CLO20,
			},
		},
		{"unknown",
			tokenSrc,
			1999,
			txtype.Token{
				Name:     "FusChain",
				Symbol:   "FUS",
				Decimals: 18,
				TokenID:  "0xa14839c9837657EFcDE754EbEAF5cbECDd801B2A",
				Coin:     1999,
				Type:     tokentype.ERC20,
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
