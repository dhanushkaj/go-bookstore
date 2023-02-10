package main

import (
	"log"
	"net/http"

	"github.com/dhanushkaj/go-bookstore/pkg/models"
	"github.com/dhanushkaj/go-bookstore/pkg/repository"
	"github.com/dhanushkaj/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r, _ := repository.InitDB()
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()

	models.InitBook(r.DB)
	routes.RegisterBookstoreRoutes(api)

	log.Fatal(http.ListenAndServe("localhost:3000", router))
}
