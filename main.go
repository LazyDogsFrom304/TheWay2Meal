package main

import (
	"flag"
	"log"
	"theway2meal/controller"
	"theway2meal/service"
)

var (
	p string
	r bool
)

func initFlags() {
	flag.StringVar(&p, "p", "8080", "port binds the service")
	flag.BoolVar(&r, "r", false, "reset database.")
}

func main() {
	// Command line process
	initFlags()
	flag.Parse()
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)

	// Loading data
	if service.DataBasePrepare(r) != nil {
		log.Fatalf("can't loading dataset, STOP")
		return
	}

	// Setup service
	r := controller.MapRoutes()
	r.Run(":" + p)
}
