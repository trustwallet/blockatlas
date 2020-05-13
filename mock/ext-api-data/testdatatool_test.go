package main

import (
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	expected := "mock/kava-api/txs?limit=25&message.sender=kava1l8va&page=1"
	result := normalizeURL("mock/kava-api/txs?limit=25&message.sender=kava1l8va&page=1")
	if result != expected {
		t.Errorf("Did not match, result %v expected %v", result, expected)
	}
	result = normalizeURL("mock/kava-api/txs?page=1&message.sender=kava1l8va&limit=25")
	if result != expected {
		t.Errorf("Did not match, result %v expected %v", result, expected)
	}
	result = normalizeURL("mock/kava-api/txs?message.sender=kava1l8va&page=1&limit=25")
	if result != expected {
		t.Errorf("Did not match, result %v expected %v", result, expected)
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
