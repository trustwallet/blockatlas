package cache

type Backend interface {
	Init(host string) error
	IsReady() bool

	GetValue(key string, value interface{}) error
	Add(key string, value interface{}) error
	Delete(key string) error

	GetAllHM(entity string) (map[string]string, error)
	GetHMValue(entity, key string, value interface{}) error
	AddHM(entity, key string, value interface{}) error
	DeleteHM(entity, key string) error
}
