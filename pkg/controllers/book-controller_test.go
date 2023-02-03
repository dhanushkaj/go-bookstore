package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dhanushkaj/go-bookstore/pkg/controllers"
	"github.com/dhanushkaj/go-bookstore/pkg/models"
	"github.com/dhanushkaj/go-bookstore/pkg/repository"
	"github.com/dhanushkaj/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetAllBooks(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/books", utils.MakeHttpHandler(controllers.GetAllBooks)).Methods("GET")
	r, _ := repository.InitDB()
	models.InitBook(r.DB)

	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	t.Log(string(body))

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	//assert.Nil(t, err)
	assert.NotEmpty(t, string(body))
	fmt.Println(string(body))
}

func TestBookById(t *testing.T) {

	router := mux.NewRouter()
	router.HandleFunc("/book/{bookId}", utils.MakeHttpHandler(controllers.GetBookById)).Methods("GET")
	r, _ := repository.InitDB()
	models.InitBook(r.DB)

	req := httptest.NewRequest(http.MethodGet, "/book/3", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	t.Log(string(body))
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var book models.Book
	err := json.Unmarshal(body, &book)
	id := &book.ID

	assert.Nil(t, err)
	assert.Equal(t, *id, uint(3))

}

func TestDeleteBook(t *testing.T) {

	router := mux.NewRouter()
	router.HandleFunc("/book/{bookId}", utils.MakeHttpHandler(controllers.DeleteBook)).Methods("DELETE")
	r, _ := repository.InitDB()
	models.InitBook(r.DB)

	req := httptest.NewRequest(http.MethodDelete, "/book/3", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	t.Log(string(body))
	assert.Equal(t, http.StatusOK, resp.StatusCode)

}

func TestUpdateBook(t *testing.T) {

	router := mux.NewRouter()
	router.HandleFunc("/book/{bookId}", utils.MakeHttpHandler(controllers.UpdateBook)).Methods("PUT")
	r, _ := repository.InitDB()
	models.InitBook(r.DB)

	var book models.Book

	name := "Test Name"
	author := "Test Author"
	publication := "Test Publication"

	book.Name = &name
	book.Author = &author
	book.Publication = &publication

	postBook, _ := json.Marshal(book)
	reader := bytes.NewReader(postBook)

	req := httptest.NewRequest(http.MethodPut, "/book/3", reader)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	t.Log(string(body))
	assert.Equal(t, http.StatusOK, resp.StatusCode)

}
