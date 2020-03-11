package setup

import "log"

func RunRedisContainer(){
	if err := runRedisContainer(); err != nil{
		log.Fatal(err)
	}
}

func StopRedisContainer(){
	if err := stopRedisContainer(); err != nil{
		log.Fatal(err)
	}
}

func RunMQContainer(){
	if err := runMQContainer(); err != nil{
		log.Fatal(err)
	}
}

func StopMQContainer(){
	if err := stopMQContainer(); err != nil{
		log.Fatal(err)
	}
}