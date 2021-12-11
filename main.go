package main

import (
	"log"
	"theway2meal/controller"
	"theway2meal/service"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)

	//test loading data
	if service.DataBasePrepare() != nil {
		log.Fatalf("can't loading dataset, STOP")
		return
	}

	r := controller.MapRoutes()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
