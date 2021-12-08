package main

import (
	"theway2meal/controller"
	"theway2meal/service"
)

func main() {
	//test loading data
	db := service.GetDefaultDB()
	service.DB_loadTestingData(db, true, true, true)

	r := controller.MapRoutes()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
