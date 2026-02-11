package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/d28035203/scaling-waddle/pkg/models"
	"github.com/d28035203/scaling-waddle/pkg/utils"
	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(mux.Vars(r)["bookId"], 10, 64)
	book, err := models.GetBookById(id)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	_ = utils.ParseBody(r, book)
	created := book.CreateBook()
	res, _ := json.Marshal(created)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(mux.Vars(r)["bookId"], 10, 64)
	book := models.DeleteBook(id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
