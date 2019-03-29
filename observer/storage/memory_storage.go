package storage

import (
	"fmt"
	"github.com/trustwallet/blockatlas/models"
)

type MemoryStorage struct {
	observerMap map[string]models.Observer
}

func (m *MemoryStorage) Setup() {

}

func (m *MemoryStorage) Get(coin uint, address string) models.Observer {
	return m.observerMap[key(coin, address)]
}

func (m *MemoryStorage) Contains(coin uint, address string) bool {
	_, ok := m.observerMap[key(coin, address)]
	return ok
}

func (m *MemoryStorage) Add(coin uint, address, webhook string) models.Observer {
	value := models.Observer{
		Coin: coin,
		Address: address,
		Webhook: webhook,
	}
	m.init().observerMap[key(coin, address)] = value

	return  value
}

func (m *MemoryStorage) List() []models.Observer {
	var values []models.Observer
	for _, value := range m.observerMap {
		values = append(values, value)
	}
	return values
}

func (m *MemoryStorage) Remove(coin uint, address string ) {
	delete(m.init().observerMap, key(coin, address))
}

func (m *MemoryStorage) init() *MemoryStorage {
	if m.observerMap == nil {
		m.observerMap = make(map[string]models.Observer)
	}
	return m
}

func key(coin uint, address string) string {
	return fmt.Sprintf("%d-%s", coin, address)
}
