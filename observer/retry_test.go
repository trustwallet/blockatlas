package observer

import (
	"errors"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
	"time"
)

func getBlock(num int64) (*blockatlas.Block, error) {
	if num == 0 {
		return nil, errors.New("teste")
	}
	return &blockatlas.Block{}, nil
}

func TestRetry(t *testing.T) {
	block, err := retry(3, time.Second*1, getBlock, 1)
	if err != nil {
		t.Error(err)
	}

	if block == nil {
		t.Error("block is nil")
	}
}

func TestRetryError(t *testing.T) {
	now := time.Now()
	block, err := retry(3, time.Second*1, getBlock, 0)
	elapsed := time.Since(now)
	if err == nil {
		t.Error("retry method need fail")
	}

	if block != nil {
		t.Error("block object need be nil")
	}

	if elapsed > time.Second*6 {
		t.Error("Thundering Herd prevent doesn't work")
	}
}
