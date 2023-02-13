package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dhanushkaj/go-bookstore/pkg/models"
	"github.com/dhanushkaj/go-bookstore/pkg/repository"
	"github.com/dhanushkaj/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	env := os.Getenv("APP_ENV")
	err := godotenv.Load("." + env + "env")

	if err != nil {
		log.Fatal(err)
	}

	r, _ := repository.InitDB()
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()

	models.InitBook(r.DB)
	routes.RegisterBookstoreRoutes(api)

	srv := &http.Server{
		Handler:      router,
		Addr:         "localhost:3000",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
