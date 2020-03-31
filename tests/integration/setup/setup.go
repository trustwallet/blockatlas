package setup

import (
	"github.com/jinzhu/gorm"
	"log"
)

func RunMQContainer() {
	if err := runMQContainer(); err != nil {
		log.Fatal(err)
	}
}

func StopMQContainer() {
	if err := stopMQContainer(); err != nil {
		log.Fatal(err)
	}
}

func RunPgContainer() *gorm.DB {
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
