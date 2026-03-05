package main

import (
	"log"
	"net/http"
	"os"

	"github.com/d28035203/scaling-waddle/pkg/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	addr := os.Getenv("SERVER_HOST") + ":" + os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(addr, r))
}
