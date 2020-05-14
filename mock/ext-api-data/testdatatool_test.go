package main

import (
	"testing"
)

func TestMockURLFromFilename(t *testing.T) {
	tests := [][]string{
		[]string{
			"mock%2Fzcash-api%2Fv2%2Faddress%2Ft1LwLWo1Mo3s4RPtUpeyUD1eYd47inL3bwX%3Fdetails%3Dtxs.json",
			"mock/zcash-api/v2/address/t1LwLWo1Mo3s4RPtUpeyUD1eYd47inL3bwX?details=txs",
			"",
		},
		[]string{
			"mock%2Fcosmos-api%2Ftxs%3Flimit%3D25%26message.sender%3Dcosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq%26page%3D1.json",
			"mock/cosmos-api/txs?limit=25&message.sender=cosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq&page=1",
			"",
		},
		[]string{
			"mock%2Fzcash-api%2Fv2%2Faddress%2Ft1LwLWo1Mo3s4RPtUpeyUD1eYd47inL3bwX%3Fdetails%3Dtxs.0010.json",
			"mock/zcash-api/v2/address/t1LwLWo1Mo3s4RPtUpeyUD1eYd47inL3bwX?details=txs",
			"0010",
		},
	}
	for _, tt := range tests {
		expectedMockUrl := tt[1]
		expectedCounter := tt[2]
		mockUrl, counter, err := mockURLFromFilename(tt[0])
		if err != nil {
			t.Errorf("Error %v input %v", err, tt[0])
		}
		if mockUrl != expectedMockUrl {
			t.Errorf("Did not match URL, result %v expected %v", mockUrl, expectedMockUrl)
		}
		if counter != expectedCounter {
			t.Errorf("Did not match counter, result %v expected %v", counter, expectedCounter)
		}
	}
}

func TestFilenameFromMockURL(t *testing.T) {
	tests := [][]string{
		[]string{
			"mock/zcash-api/v2/address/t1LwLWo1Mo3s4RPtUpeyUD1eYd47inL3bwX?details=txs",
			"",
			"mock%2Fzcash-api%2Fv2%2Faddress%2Ft1LwLWo1Mo3s4RPtUpeyUD1eYd47inL3bwX%3Fdetails%3Dtxs.json",
		},
		[]string{
			"mock/cosmos-api/txs?limit=25&message.sender=cosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq&page=1",
			"",
			"mock%2Fcosmos-api%2Ftxs%3Flimit%3D25%26message.sender%3Dcosmos1dx27g0kzhwej0ekcf2k9hsktcxnmpl7fcehcvq%26page%3D1.json",
		},
		[]string{
			"mock/zcash-api/v2/address/t1LwLWo1Mo3s4RPtUpeyUD1eYd47inL3bwX?details=txs",
			"0010",
			"mock%2Fzcash-api%2Fv2%2Faddress%2Ft1LwLWo1Mo3s4RPtUpeyUD1eYd47inL3bwX%3Fdetails%3Dtxs.0010.json",
		},
	}
	for _, tt := range tests {
		expectedFN := tt[2]
		filename := filenameFromMockURL(tt[0], tt[1])
		if filename != expectedFN {
			t.Errorf("Did not match, result %v expected %v", filename, expectedFN)
		}
	}
}

func TestNormalizeURL(t *testing.T) {
	tests := [][]string{
		[]string{
			"mock/kava-api/txs?limit=25&message.sender=kava1l8va&page=1",
			"mock/kava-api/txs?limit=25&message.sender=kava1l8va&page=1",
		},
		[]string{
			"mock/kava-api/txs?page=1&message.sender=kava1l8va&limit=25",
			"mock/kava-api/txs?limit=25&message.sender=kava1l8va&page=1",
		},
		[]string{
			"mock/kava-api/txs?message.sender=kava1l8va&page=1&limit=25",
			"mock/kava-api/txs?limit=25&message.sender=kava1l8va&page=1",
		},
	}
	for _, tt := range tests {
		expected := tt[1]
		result := normalizeURL(tt[0])
		if result != expected {
			t.Errorf("Did not match, result %v expected %v", result, expected)
		}
	}
}

func TestURLMap(t *testing.T) {
	urlMap := URLMap{}
	urlMap["mock/tron-api"] = URLEntry{"https://api.trongrid.io"}
	urlMap["mock/binance-api"] = URLEntry{"https://explorer.binance.org/api/v1"}

	result := getRealURL("mock/tron-api/v1/assets/1002798", urlMap)
	expected := "https://api.trongrid.io/v1/assets/1002798"
	if result != expected {
		t.Errorf("Did not match, result %v expected %v", result, expected)
	}
	result = getRealURL("mock/nosuchmap-api/v1/assets/1002798", urlMap)
	expected = ""
	if result != expected {
		t.Errorf("Did not match, result %v expected %v", result, expected)
	}

	result = getMockURL("https://api.trongrid.io/v1/assets/1002798", urlMap)
	expected = "mock/tron-api/v1/assets/1002798"
	if result != expected {
		t.Errorf("Did not match, result %v expected %v", result, expected)
	}
	result = getMockURL("https://nosuchmap.io/v1/assets/1002798", urlMap)
	expected = ""
	if result != expected {
		t.Errorf("Did not match, result %v expected %v", result, expected)
	}
}
