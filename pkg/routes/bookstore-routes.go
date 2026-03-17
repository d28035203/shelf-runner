// Package routes maps URL paths to controller handlers (routing layer).
package routes

import (
	"github.com/d28035203/shelf-runner/pkg/controllers"
	"github.com/gorilla/mux"
)

// RegisterBookStoreRoutes attaches CRUD endpoints for books onto the given router.
//
//	GET    /book/           → list
//	GET    /book/{bookId}   → get one
//	POST   /book/           → create
//	PUT    /book/{bookId}   → update
//	DELETE /book/{bookId}   → delete
var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
