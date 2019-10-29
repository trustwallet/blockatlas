package storage

import (
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/pkg/storage/sql"
)

func (s *Storage) GetBlockNumber(coin uint) (int64, error) {
	b, ok := s.blockHeights.GetBlock(int(coin))
	if ok {
		return b.BlockHeight, nil
	}

	b = &Block{Coin: int(coin)}
	err := sql.Get(s.Client, &b)
	if err != nil {
		return 0, nil
	}
	return b.BlockHeight, nil
}

func (s *Storage) SetBlockNumber(coin uint, num int64) {
	s.blockHeights.SetBlock(&Block{Coin: int(coin), BlockHeight: num})
}

func (s *Storage) SaveBlock(coin uint, num int64) error {
	b := &Block{Coin: int(coin), BlockHeight: num}
	logger.Info("Saving block", logger.Params{"Coin": b.Coin, "Height": b.BlockHeight})
	err := sql.Save(s.Client, b)
	if err != nil {
		return errors.E(err, errors.Params{"block": num, "coin": coin})
	}
	return nil
}

func (s *Storage) SaveAllBlocks() error {
	logger.Info("Saving cache blocks in database")
	h := s.blockHeights.GetHeights()
	for _, b := range h {
		logger.Info("Saving block", logger.Params{"Coin": b.Coin, "Height": b.BlockHeight})
		err := sql.Save(s.Client, b)
		if err != nil {
			logger.Error(err)
		}
	}
	return nil
}
