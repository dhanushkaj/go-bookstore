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

type TestBook struct {
	ID uint
}

var testBook TestBook

func TestCreateBook(t *testing.T) {

	router := mux.NewRouter()
	router.HandleFunc("/book", utils.MakeHttpHandler(controllers.CreateBook)).Methods("POST")
	r, _ := repository.InitDB()
	models.InitBook(r.DB)

	var book models.Book

	name := "Test Name2"
	author := "Test Author2"
	publication := "Test Publication2"

	book.Name = name
	book.Author = author
	book.Publication = publication

	postBook, _ := json.Marshal(book)
	reader := bytes.NewReader(postBook)

	req := httptest.NewRequest(http.MethodPost, "/book", reader)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	t.Log(string(body))
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var responseBook models.Book
	err := json.Unmarshal(body, &responseBook)
	responseName := responseBook.Name
	testBook.ID = *&responseBook.ID
	assert.Nil(t, err)
	assert.Equal(t, name, responseName)

}

func TestUpdateBook(t *testing.T) {

	router := mux.NewRouter()
	router.HandleFunc("/book/{bookId}", utils.MakeHttpHandler(controllers.UpdateBook)).Methods("PUT")
	r, _ := repository.InitDB()
	models.InitBook(r.DB)

	var book models.Book

	name := "Test Name3"
	author := "Test Author3"
	publication := "Test Publication3"

	book.Name = name
	book.Author = author
	book.Publication = publication

	postBook, _ := json.Marshal(book)
	reader := bytes.NewReader(postBook)

	req := httptest.NewRequest(http.MethodPut, "/book/"+fmt.Sprintf("%v", testBook.ID), reader)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	t.Log(string(body))
	assert.Equal(t, http.StatusOK, resp.StatusCode)

}

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

	req := httptest.NewRequest(http.MethodGet, "/book/"+fmt.Sprintf("%v", testBook.ID), nil)
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
	assert.Equal(t, *id, uint(testBook.ID))

}

func TestDeleteBook(t *testing.T) {

	router := mux.NewRouter()
	router.HandleFunc("/book/{bookId}", utils.MakeHttpHandler(controllers.DeleteBook)).Methods("DELETE")
	r, _ := repository.InitDB()
	models.InitBook(r.DB)

	req := httptest.NewRequest(http.MethodDelete, "/book/"+fmt.Sprintf("%v", testBook.ID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	t.Log(string(body))
	assert.Equal(t, http.StatusOK, resp.StatusCode)

}
