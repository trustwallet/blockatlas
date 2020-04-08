package trustray

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
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
	expected    *blockatlas.Token
	coin        int
}

func TestNormalizeToken(t *testing.T) {
	var tests = []struct {
		name     string
		tokenRaw string
		coin     int
		want     blockatlas.Token
	}{
		{
			"ethereum erc20",
			tokenSrc,
			coin.ETH,
			blockatlas.Token{
				Name:     "FusChain",
				Symbol:   "FUS",
				Decimals: 18,
				TokenID:  "0xa14839c9837657EFcDE754EbEAF5cbECDd801B2A",
				Coin:     coin.ETH,
				Type:     blockatlas.TokenTypeERC20,
			},
		},
		{"classic etc20",
			tokenSrc,
			coin.ETC,
			blockatlas.Token{
				Name:     "FusChain",
				Symbol:   "FUS",
				Decimals: 18,
				TokenID:  "0xa14839c9837657EFcDE754EbEAF5cbECDd801B2A",
				Coin:     coin.ETC,
				Type:     blockatlas.TokenTypeETC20,
			},
		},
		{"gochain go20",
			tokenSrc,
			coin.GO,
			blockatlas.Token{
				Name:     "FusChain",
				Symbol:   "FUS",
				Decimals: 18,
				TokenID:  "0xa14839c9837657EFcDE754EbEAF5cbECDd801B2A",
				Coin:     coin.GO,
				Type:     blockatlas.TokenTypeGO20,
			},
		},
		{"thudertoken tt20",
			tokenSrc,
			coin.TT,
			blockatlas.Token{
				Name:     "FusChain",
				Symbol:   "FUS",
				Decimals: 18,
				TokenID:  "0xa14839c9837657EFcDE754EbEAF5cbECDd801B2A",
				Coin:     coin.TT,
				Type:     blockatlas.TokenTypeTT20,
			},
		},
		{"wanchain wan20",
			tokenSrc,
			coin.WAN,
			blockatlas.Token{
				Name:     "FusChain",
				Symbol:   "FUS",
				Decimals: 18,
				TokenID:  "0xa14839c9837657EFcDE754EbEAF5cbECDd801B2A",
				Coin:     coin.WAN,
				Type:     blockatlas.TokenTypeWAN20,
			},
		},
		{"poa poa20",
			tokenSrc,
			coin.POA,
			blockatlas.Token{
				Name:     "FusChain",
				Symbol:   "FUS",
				Decimals: 18,
				TokenID:  "0xa14839c9837657EFcDE754EbEAF5cbECDd801B2A",
				Coin:     coin.POA,
				Type:     blockatlas.TokenTypePOA20,
			},
		},
		{"callisto clo20",
			tokenSrc,
			coin.CLO,
			blockatlas.Token{
				Name:     "FusChain",
				Symbol:   "FUS",
				Decimals: 18,
				TokenID:  "0xa14839c9837657EFcDE754EbEAF5cbECDd801B2A",
				Coin:     coin.CLO,
				Type:     blockatlas.TokenTypeCLO20,
			},
		},
		{"unkown",
			tokenSrc,
			1999,
			blockatlas.Token{
				Name:     "FusChain",
				Symbol:   "FUS",
				Decimals: 18,
				TokenID:  "0xa14839c9837657EFcDE754EbEAF5cbECDd801B2A",
				Coin:     1999,
				Type:     "unknown",
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

	if !bytes.Equal(resJSON, dstJSON) {
		println(string(resJSON))
		println(string(dstJSON))
		t.Error("token: token don't equal")
	}
}
