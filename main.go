package main

import (
	"github.com/Covid19_Vaccination_Register/database"
)

func main() {
	db := database.ConnectDatabase()
	database.MigrateSchema()

	// r := route.InitializeRoutes()
	// s := server.InitializeServer(r)
	//
	// log.Fatal(s.ListenAndServe())
}
