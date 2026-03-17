// Package controllers implements HTTP handlers for the bookstore API (controller layer).
package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/d28035203/shelf-runner/pkg/models"
	"github.com/d28035203/shelf-runner/pkg/utils"
	"github.com/gorilla/mux"
)

// writeJSON sets Content-Type and encodes payload as JSON with the given status.
func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

// writeError returns a uniform error object: { "error": "message" }.
func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

// GetBook lists all books (GET /book/).
func GetBook(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	writeJSON(w, http.StatusOK, books)
}

// GetBookById returns one book (GET /book/{bookId}).
func GetBookById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["bookId"], 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid book id")
		return
	}

	book, err := models.GetBookById(id)
	if err != nil {
		writeError(w, http.StatusNotFound, "book not found")
		return
	}

	writeJSON(w, http.StatusOK, book)
}

// CreateBook creates a book from JSON (POST /book/).
// Required fields: name, author.
func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	if err := utils.ParseBody(r, book); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	if book.Name == "" || book.Author == "" {
		writeError(w, http.StatusBadRequest, "name and author are required")
		return
	}

	created := book.CreateBook()
	writeJSON(w, http.StatusCreated, created)
}

// UpdateBook partially updates name/author/publication (PUT /book/{bookId}).
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	payload := &models.Book{}
	if err := utils.ParseBody(r, payload); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	id, err := strconv.ParseInt(mux.Vars(r)["bookId"], 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid book id")
		return
	}

	book, err := models.GetBookById(id)
	if err != nil {
		writeError(w, http.StatusNotFound, "book not found")
		return
	}

	// Only overwrite fields that the client actually sent (non-empty strings).
	if payload.Name != "" {
		book.Name = payload.Name
	}
	if payload.Author != "" {
		book.Author = payload.Author
	}
	if payload.Publication != "" {
		book.Publication = payload.Publication
	}

	if err := models.UpdateBook(book); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to update book")
		return
	}

	writeJSON(w, http.StatusOK, book)
}

// DeleteBook removes a book by id (DELETE /book/{bookId}).
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["bookId"], 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid book id")
		return
	}

	book := models.DeleteBook(id)
	writeJSON(w, http.StatusOK, book)
}
