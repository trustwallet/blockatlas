package parser

import (
	"errors"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
	"time"
)

func getBlock(num int64) (*blockatlas.Block, error) {
	if num == 0 {
		return nil, errors.New("test")
	}
	return &blockatlas.Block{}, nil
}

func TestRetry(t *testing.T) {
	block, err := getBlockByNumberWithRetry(3, time.Millisecond*1, getBlock, 1)
	if err != nil {
		t.Error(err)
	}

	if block == nil {
		t.Error("block is nil")
	}
}

func TestRetryError(t *testing.T) {
	now := time.Now()
	block, err := getBlockByNumberWithRetry(3, time.Millisecond*1, getBlock, 0)
	elapsed := time.Since(now)
	if err == nil {
		t.Error("getBlockByNumberWithRetry method need fail")
	}

	if block != nil {
		t.Error("block object need be nil")
	}

	if elapsed > time.Millisecond*6 {
		t.Error("Thundering Herd prevent doesn't work")
	}
}
