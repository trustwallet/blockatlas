package setup

import "log"

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

func RunPgContainer() {
	if err := runPgContainerAndInitConnection(); err != nil {
		log.Fatal(err)
	}
}

func StopPgContainer() {
	if err := stopPgContainer(); err != nil {
		log.Fatal(err)
	}
}
