package observer

import "github.com/trustwallet/blockatlas/pkg/blockatlas"

type BlockData struct {
	Block blockatlas.Block
	Coin  uint
}
