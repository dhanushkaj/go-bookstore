package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dhanushkaj/go-bookstore/pkg/models"
	"github.com/dhanushkaj/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetAllBooks(w http.ResponseWriter, r *http.Request) error {
	newBooks, _ := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	return utils.WriteJSON(w, res, http.StatusOK, nil)
}

func GetBookById(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		return utils.ApiError{Err: "ierror while parding book id", Status: http.StatusForbidden}
	}
	book, _, err := models.GetBookById(ID)
	res, _ := json.Marshal(book)
	return utils.WriteJSON(w, res, http.StatusOK, res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) error {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b, _ := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	return utils.WriteJSON(w, res, http.StatusOK, nil)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) error {

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parding book id")
	}
	book, err := models.DeleteBook(ID)

	res, _ := json.Marshal(book)
	return utils.WriteJSON(w, res, http.StatusOK, nil)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) error {

	var UpdateBook = &models.Book{}
	utils.ParseBody(r, UpdateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parding book id")
	}

	bookDetail, db, err := models.GetBookById(ID)

	if bookDetail.Name != nil {
		bookDetail.Name = UpdateBook.Name
	}
	if bookDetail.Author != nil {
		bookDetail.Author = UpdateBook.Author
	}
	if bookDetail.Publication != nil {
		bookDetail.Publication = UpdateBook.Publication
	}
	db.Save(&bookDetail)
	res, _ := json.Marshal(bookDetail)
	return utils.WriteJSON(w, res, http.StatusOK, nil)
}
