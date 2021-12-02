package main

import "theway2meal/controller"

func main() {
	r := controller.MapRoutes()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
