package storage

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/models"
)

type Storage interface {
	Setup()
	List() []models.Observer
	Get(coin uint, address string) models.Observer
	Add(coin uint, address, webhook string) models.Observer
	Remove(coin uint, address string)
	Contains(coin uint, address string) bool
}

const (
	MemoryStorageKey = "memory"
)

var storageMap = map[string]Storage {
	MemoryStorageKey: new(MemoryStorage),
}

func init() {
	for _, storage := range storageMap {
		storage.Setup()
	}
}

func GetInstance() Storage {
	return storageMap[viper.GetString("storage")]
}
