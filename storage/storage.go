package storage

import "github.com/trustwallet/blockatlas/models"

type Storage interface {
	List() []models.Observer
	Get(coin uint, address string) models.Observer
	Add(coin uint, address, webhook string)
	Remove(coin uint, address string)
}
