package routes

import (
	"github.com/dhanushkaj/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"

	"github.com/dhanushkaj/go-bookstore/pkg/utils"
)

var RegisterBookstoreRoutes = func(router *mux.Router) {

	router.HandleFunc("/book", utils.MakeHttpHandler(controllers.CreateBook)).Methods("POST")
	router.HandleFunc("/books", utils.MakeHttpHandler(controllers.GetAllBooks)).Methods("GET")
	router.HandleFunc("/book/{bookId}", utils.MakeHttpHandler(controllers.GetBookById)).Methods("GET")
	router.HandleFunc("/book/{bookId}", utils.MakeHttpHandler(controllers.UpdateBook)).Methods(("PUT"))
	router.HandleFunc("/book/{bookId}", utils.MakeHttpHandler(controllers.DeleteBook)).Methods(("DELETE"))
	router.HandleFunc("/books/{name}/{author}", utils.MakeHttpHandler(controllers.GetBookByNameAndAuthor)).Methods("GET")

}
