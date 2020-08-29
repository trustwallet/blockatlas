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
			coin.ETHEREUM,
			blockatlas.Token{
				Name:     "FusChain",
				Symbol:   "FUS",
				Decimals: 18,
				TokenID:  "0xa14839c9837657EFcDE754EbEAF5cbECDd801B2A",
				Coin:     coin.ETHEREUM,
				Type:     blockatlas.TokenTypeERC20,
			},
		},
		{"classic etc20",
			tokenSrc,
			coin.CLASSIC,
			blockatlas.Token{
				Name:     "FusChain",
				Symbol:   "FUS",
				Decimals: 18,
				TokenID:  "0xa14839c9837657EFcDE754EbEAF5cbECDd801B2A",
				Coin:     coin.CLASSIC,
				Type:     blockatlas.TokenTypeETC20,
			},
		},
		{"gochain go20",
			tokenSrc,
			coin.GOCHAIN,
			blockatlas.Token{
				Name:     "FusChain",
				Symbol:   "FUS",
				Decimals: 18,
				TokenID:  "0xa14839c9837657EFcDE754EbEAF5cbECDd801B2A",
				Coin:     coin.GOCHAIN,
				Type:     blockatlas.TokenTypeGO20,
			},
		},
		{"thudertoken tt20",
			tokenSrc,
			coin.THUNDERTOKEN,
			blockatlas.Token{
				Name:     "FusChain",
				Symbol:   "FUS",
				Decimals: 18,
				TokenID:  "0xa14839c9837657EFcDE754EbEAF5cbECDd801B2A",
				Coin:     coin.THUNDERTOKEN,
				Type:     blockatlas.TokenTypeTT20,
			},
		},
		{"wanchain wan20",
			tokenSrc,
			coin.WANCHAIN,
			blockatlas.Token{
				Name:     "FusChain",
				Symbol:   "FUS",
				Decimals: 18,
				TokenID:  "0xa14839c9837657EFcDE754EbEAF5cbECDd801B2A",
				Coin:     coin.WANCHAIN,
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
			coin.CALLISTO,
			blockatlas.Token{
				Name:     "FusChain",
				Symbol:   "FUS",
				Decimals: 18,
				TokenID:  "0xa14839c9837657EFcDE754EbEAF5cbECDd801B2A",
				Coin:     coin.CALLISTO,
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
