package setup

import (
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/pkg/servicerepo"
	"log"
)

func RunMQContainer(serviceRepo *servicerepo.ServiceRepo) {
	if err := runMQContainer(serviceRepo); err != nil {
		log.Fatal(err)
	}
}

func StopMQContainer() {
	if err := stopMQContainer(); err != nil {
		log.Fatal(err)
	}
}

func RunPgContainer() *db.Instance {
	dbConn, err := runPgContainerAndInitConnection()
	if err != nil {
		log.Fatal(err)
	}
	return dbConn
}

func StopPgContainer() {
	if err := stopPgContainer(); err != nil {
		log.Fatal(err)
	}
}
