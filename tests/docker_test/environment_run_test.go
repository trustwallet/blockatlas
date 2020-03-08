// +build integration

package docker_test

import (
	"github.com/trustwallet/blockatlas/tests/docker_test/setup"
	"os"
	"testing"
)

func TestMain(m *testing.M){
	setup.RunMQContainer()
	setup.RunRedisContainer()
	code := m.Run()
	setup.StopMQContainer()
	setup.StopRedisContainer()
	os.Exit(code)
}