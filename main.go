package main

import (
	"log"

	route "github.com/Covid19_Vaccination_Register/http"
	"github.com/Covid19_Vaccination_Register/database"
	"github.com/Covid19_Vaccination_Register/server"
)

func main() {
	database.InitializeDatabase()
	r := route.InitializeRoutes()
	srv := server.InitializeServer(r)

	log.Fatal(srv.ListenAndServe())
}
