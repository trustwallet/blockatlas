package setup

import (
	"fmt"
	"github.com/ory/dockertest"
	"github.com/trustwallet/blockatlas/storage"
	"log"
)

var (
	Cache *storage.Storage
	redisResource     *dockertest.Resource
)

func runRedisContainer() error {
	var err error
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	redisResource, err = pool.Run("redis", "latest", nil)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err = pool.Retry(func() error {
		var err error
		Cache = storage.New()
		err = Cache.Init(fmt.Sprintf("redis://localhost:%s", redisResource.GetPort("6379/tcp")))
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func stopRedisContainer() error {
	return redisResource.Close()
}