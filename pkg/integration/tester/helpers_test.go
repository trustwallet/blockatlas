// +build integration

package tester

import (
	"reflect"
	"testing"
)

func TestGetUrl(t *testing.T) {
	url := "http://blockatlas.com/v1/{{.Coin}}/{{.Address}}"
	coin := "ethereum"
	address := "0xfc10cab6a50a1ab10c56983c80cc82afc6559cf1"
	expected := "http://blockatlas.com/v1/ethereum/0xfc10cab6a50a1ab10c56983c80cc82afc6559cf1"
	r, err := getParameters(url, coin, address)
	if err != nil {
		t.Errorf("getBaseUrl() failed, error: %s", err)
	}
	if !reflect.DeepEqual(r, expected) {
		t.Errorf("result (%s) is not the expected: %s", r, err)
	}
}
