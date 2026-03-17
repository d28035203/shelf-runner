// Command main starts the bookstore REST API server.
//
// It loads environment variables, builds a Gorilla Mux router, registers
// bookstore routes, and listens on SERVER_HOST:SERVER_PORT (defaults provided).
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/d28035203/shelf-runner/pkg/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Prefer .env for local development; container/cloud can inject env vars directly.
	if err := godotenv.Load(); err != nil {
		log.Println("warning: .env file not found, relying on process environment")
	}

	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")
	if host == "" {
		host = "0.0.0.0"
	}
	if port == "" {
		port = "9010"
	}

	r := mux.NewRouter()
	// Side-effect: importing models (via routes → controllers) connects DB and migrates.
	routes.RegisterBookStoreRoutes(r)

	addr := host + ":" + port
	log.Printf("bookstore API listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
