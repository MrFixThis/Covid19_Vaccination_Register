package main

import (
	"log"

	route "github.com/Covid19_Vaccination_Register/pkg/http"
	"github.com/Covid19_Vaccination_Register/pkg/server"
	"github.com/Covid19_Vaccination_Register/pkg/storage"
)

func main() {
	storage.InitializeDatabase()
	r := route.InitializeRoutes()
	srv := server.InitializeServer(r)

	log.Fatal(srv.ListenAndServe())
}
